package dynamic

// Binary serialization and de-serialization for dynamic messages

import (
	"fmt"
	"io"
	"math"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/jhump/protoreflect/codec"
	"github.com/jhump/protoreflect/desc"
)

// defaultDeterminism, if true, will mean that calls to Marshal will produce
// deterministic output. This is used to make the output of proto.Marshal(...)
// deterministic (since there is no way to have that convey determinism intent).
// **This is only used from tests.**
var defaultDeterminism = false

// Marshal serializes this message to bytes, returning an error if the operation
// fails. The resulting bytes are in the standard protocol buffer binary format.
func (m *Message) Marshal() ([]byte, error) {
	var b codec.Buffer
	if err := m.marshal(&b, defaultDeterminism); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// MarshalAppend behaves exactly the same as Marshal, except instead of allocating a
// new byte slice to marshal into, it uses the provided byte slice. The backing array
// for the returned byte slice *may* be the same as the one that was passed in, but
// it's not guaranteed as a new backing array will automatically be allocated if
// more bytes need to be written than the provided buffer has capacity for.
func (m *Message) MarshalAppend(b []byte) ([]byte, error) {
	codedBuf := codec.NewBuffer(b)
	if err := m.marshal(codedBuf, defaultDeterminism); err != nil {
		return nil, err
	}
	return codedBuf.Bytes(), nil
}

// MarshalDeterministic serializes this message to bytes in a deterministic way,
// returning an error if the operation fails. This differs from Marshal in that
// map keys will be sorted before serializing to bytes. The protobuf spec does
// not define ordering for map entries, so Marshal will use standard Go map
// iteration order (which will be random). But for cases where determinism is
// more important than performance, use this method instead.
func (m *Message) MarshalDeterministic() ([]byte, error) {
	var b codec.Buffer
	if err := m.marshal(&b, true); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *Message) marshal(b *codec.Buffer, deterministic bool) error {
	if err := m.marshalKnownFields(b, deterministic); err != nil {
		return err
	}
	return m.marshalUnknownFields(b)
}

func (m *Message) marshalKnownFields(b *codec.Buffer, deterministic bool) error {
	for _, tag := range m.knownFieldTags() {
		itag := int32(tag)
		val := m.values[itag]
		fd := m.FindFieldDescriptor(itag)
		if fd == nil {
			panic(fmt.Sprintf("Couldn't find field for tag %d", itag))
		}
		if deterministic {
			if err := b.EncodeFieldValueDeterministic(fd, val); err != nil {
				return err
			}
		} else if err := b.EncodeFieldValue(fd, val); err != nil {
			return err
		}
	}
	return nil
}

func (m *Message) marshalUnknownFields(b *codec.Buffer) error {
	for _, tag := range m.unknownFieldTags() {
		itag := int32(tag)
		sl := m.unknownFields[itag]
		for _, u := range sl {
			if err := b.EncodeTagAndWireType(itag, u.Encoding); err != nil {
				return err
			}
			switch u.Encoding {
			case proto.WireBytes:
				if err := b.EncodeRawBytes(u.Contents); err != nil {
					return err
				}
			case proto.WireStartGroup:
				_, _ = b.Write(u.Contents)
				if err := b.EncodeTagAndWireType(itag, proto.WireEndGroup); err != nil {
					return err
				}
			case proto.WireFixed32:
				if err := b.EncodeFixed32(u.Value); err != nil {
					return err
				}
			case proto.WireFixed64:
				if err := b.EncodeFixed64(u.Value); err != nil {
					return err
				}
			case proto.WireVarint:
				if err := b.EncodeVarint(u.Value); err != nil {
					return err
				}
			default:
				return proto.ErrInternalBadWireType
			}
		}
	}
	return nil
}

// Unmarshal de-serializes the message that is present in the given bytes into
// this message. It first resets the current message. It returns an error if the
// given bytes do not contain a valid encoding of this message type.
func (m *Message) Unmarshal(b []byte) error {
	m.Reset()
	if err := m.UnmarshalMerge(b); err != nil {
		return err
	}
	return m.Validate()
}

// UnmarshalMerge de-serializes the message that is present in the given bytes
// into this message. Unlike Unmarshal, it does not first reset the message,
// instead merging the data in the given bytes into the existing data in this
// message.
func (m *Message) UnmarshalMerge(b []byte) error {
	return m.unmarshal(codec.NewBuffer(b), false)
}

func (m *Message) unmarshal(buf *codec.Buffer, isGroup bool) error {
	for !buf.EOF() {
		tagNumber, wireType, err := buf.DecodeTagAndWireType()
		if err != nil {
			return err
		}
		if wireType == proto.WireEndGroup {
			if isGroup {
				// finished parsing group
				return nil
			} else {
				return proto.ErrInternalBadWireType
			}
		}
		fd := m.FindFieldDescriptor(tagNumber)
		if fd == nil {
			err := m.unmarshalUnknownField(tagNumber, wireType, buf)
			if err != nil {
				return err
			}
		} else {
			err := m.unmarshalKnownField(fd, wireType, buf)
			if err != nil {
				return err
			}
		}
	}
	if isGroup {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func unmarshalSimpleField(fd *desc.FieldDescriptor, v uint64) (interface{}, error) {
	switch fd.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return v != 0, nil
	case descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32:
		if v > math.MaxUint32 {
			return nil, NumericOverflowError
		}
		return uint32(v), nil

	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		s := int64(v)
		if s > math.MaxInt32 || s < math.MinInt32 {
			return nil, NumericOverflowError
		}
		return int32(s), nil

	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		if v > math.MaxUint32 {
			return nil, NumericOverflowError
		}
		return int32(v), nil

	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		if v > math.MaxUint32 {
			return nil, NumericOverflowError
		}
		return codec.DecodeZigZag32(v), nil

	case descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return v, nil

	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return int64(v), nil

	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		return codec.DecodeZigZag64(v), nil

	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if v > math.MaxUint32 {
			return nil, NumericOverflowError
		}
		return math.Float32frombits(uint32(v)), nil

	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return math.Float64frombits(v), nil

	default:
		// bytes, string, message, and group cannot be represented as a simple numeric value
		return nil, fmt.Errorf("bad input; field %s requires length-delimited wire type", fd.GetFullyQualifiedName())
	}
}

func unmarshalLengthDelimitedField(fd *desc.FieldDescriptor, bytes []byte, mf *MessageFactory) (interface{}, error) {
	switch {
	case fd.GetType() == descriptor.FieldDescriptorProto_TYPE_BYTES:
		return bytes, nil

	case fd.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING:
		return string(bytes), nil

	case fd.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE ||
		fd.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP:
		msg := mf.NewMessage(fd.GetMessageType())
		err := proto.Unmarshal(bytes, msg)
		if err != nil {
			return nil, err
		} else {
			return msg, nil
		}

	default:
		// even if the field is not repeated or not packed, we still parse it as such for
		// backwards compatibility (e.g. message we are de-serializing could have been both
		// repeated and packed at the time of serialization)
		packedBuf := codec.NewBuffer(bytes)
		var slice []interface{}
		var val interface{}
		for !packedBuf.EOF() {
			var v uint64
			var err error
			if varintTypes[fd.GetType()] {
				v, err = packedBuf.DecodeVarint()
			} else if fixed32Types[fd.GetType()] {
				v, err = packedBuf.DecodeFixed32()
			} else if fixed64Types[fd.GetType()] {
				v, err = packedBuf.DecodeFixed64()
			} else {
				return nil, fmt.Errorf("bad input; cannot parse length-delimited wire type for field %s", fd.GetFullyQualifiedName())
			}
			if err != nil {
				return nil, err
			}
			val, err = unmarshalSimpleField(fd, v)
			if err != nil {
				return nil, err
			}
			if fd.IsRepeated() {
				slice = append(slice, val)
			}
		}
		if fd.IsRepeated() {
			return slice, nil
		} else {
			// if not a repeated field, last value wins
			return val, nil
		}
	}
}

func (m *Message) unmarshalKnownField(fd *desc.FieldDescriptor, encoding int8, b *codec.Buffer) error {
	var val interface{}
	var err error
	switch encoding {
	case proto.WireFixed32:
		var num uint64
		num, err = b.DecodeFixed32()
		if err == nil {
			val, err = unmarshalSimpleField(fd, num)
		}
	case proto.WireFixed64:
		var num uint64
		num, err = b.DecodeFixed64()
		if err == nil {
			val, err = unmarshalSimpleField(fd, num)
		}
	case proto.WireVarint:
		var num uint64
		num, err = b.DecodeVarint()
		if err == nil {
			val, err = unmarshalSimpleField(fd, num)
		}

	case proto.WireBytes:
		if fd.GetType() == descriptor.FieldDescriptorProto_TYPE_BYTES {
			val, err = b.DecodeRawBytes(true) // defensive copy
		} else if fd.GetType() == descriptor.FieldDescriptorProto_TYPE_STRING {
			var raw []byte
			raw, err = b.DecodeRawBytes(true) // defensive copy
			if err == nil {
				val = string(raw)
			}
		} else {
			var raw []byte
			raw, err = b.DecodeRawBytes(false)
			if err == nil {
				val, err = unmarshalLengthDelimitedField(fd, raw, m.mf)
			}
		}

	case proto.WireStartGroup:
		if fd.GetMessageType() == nil {
			return fmt.Errorf("cannot parse field %s from group-encoded wire type", fd.GetFullyQualifiedName())
		}
		msg := m.mf.NewMessage(fd.GetMessageType())
		if dm, ok := msg.(*Message); ok {
			err = dm.unmarshal(b, true)
			if err == nil {
				val = dm
			}
		} else {
			data, err := b.ReadGroup(false)
			if err == nil {
				err = proto.Unmarshal(data, msg)
				if err == nil {
					val = msg
				}
			}
		}

	default:
		return proto.ErrInternalBadWireType
	}
	if err != nil {
		return err
	}

	return mergeField(m, fd, val)
}

func (m *Message) unmarshalUnknownField(tagNumber int32, encoding int8, b *codec.Buffer) error {
	u := UnknownField{Encoding: encoding}
	var err error
	switch encoding {
	case proto.WireFixed32:
		u.Value, err = b.DecodeFixed32()
	case proto.WireFixed64:
		u.Value, err = b.DecodeFixed64()
	case proto.WireVarint:
		u.Value, err = b.DecodeVarint()
	case proto.WireBytes:
		u.Contents, err = b.DecodeRawBytes(true)
	case proto.WireStartGroup:
		u.Contents, err = b.ReadGroup(true)
	default:
		err = proto.ErrInternalBadWireType
	}
	if err != nil {
		return err
	}
	if m.unknownFields == nil {
		m.unknownFields = map[int32][]UnknownField{}
	}
	m.unknownFields[tagNumber] = append(m.unknownFields[tagNumber], u)
	return nil
}
