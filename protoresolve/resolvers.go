package protoresolve

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var (
	// GlobalDescriptors provides a view of protoregistry.GlobalFiles and protoregistry.GlobalTypes
	// as a Resolver.
	GlobalDescriptors = ResolverFromPools(protoregistry.GlobalFiles, protoregistry.GlobalTypes)

	// ErrNotFound is a sentinel error that is returned from resolvers to indicate that the named
	// element is not known to the registry. It is the same as protoregistry.NotFound.
	ErrNotFound = protoregistry.NotFound
)

// FileResolver can resolve file descriptors by path.
type FileResolver interface {
	FindFileByPath(string) (protoreflect.FileDescriptor, error)
}

// FilePool is a FileResolver that also allows iteration over the known file descriptors.
type FilePool interface {
	FileResolver
	NumFiles() int
	RangeFiles(fn func(protoreflect.FileDescriptor) bool)
	NumFilesByPackage(name protoreflect.FullName) int
	RangeFilesByPackage(name protoreflect.FullName, fn func(protoreflect.FileDescriptor) bool)
}

// DescriptorResolver can resolve descriptors by full name.
type DescriptorResolver interface {
	FindDescriptorByName(protoreflect.FullName) (protoreflect.Descriptor, error)
}

// DescriptorPool is a FilePool that also functions as a DescriptorResolver.
type DescriptorPool interface {
	FilePool
	DescriptorResolver
}

var _ DescriptorPool = (*Registry)(nil)
var _ DescriptorPool = (*protoregistry.Files)(nil)

// DescriptorRegistry is a file and descriptor resolver that allows the caller to add files
// (and their contained descriptors) to the set of files and descriptors it can resolve.
type DescriptorRegistry interface {
	DescriptorPool
	RegisterFile(protoreflect.FileDescriptor) error
}

var _ DescriptorRegistry = (*Registry)(nil)
var _ DescriptorRegistry = (*protoregistry.Files)(nil)

// ExtensionResolver can resolve extensions based on the containing message name and field number.
type ExtensionResolver interface {
	FindExtensionByName(protoreflect.FullName) (protoreflect.ExtensionDescriptor, error)
	FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionDescriptor, error)
}

// ExtensionPool is an ExtensionResolver that also allows iteration over all extensions for a message.
type ExtensionPool interface {
	ExtensionResolver
	RangeExtensionsByMessage(message protoreflect.FullName, fn func(protoreflect.ExtensionDescriptor) bool)
}

// MessageResolver can resolve messages based on their name or a type URL. URLs must include the
// fully-qualified type name as the last URI path component.
type MessageResolver interface {
	FindMessageByName(protoreflect.FullName) (protoreflect.MessageDescriptor, error)
	FindMessageByURL(url string) (protoreflect.MessageDescriptor, error)
}

// ExtensionTypeResolver can resolve extension types. An extension type includes more than just
// a descriptor but also refers to runtime types (which could be static types generated by the
// protoc-gen-go plugin).
//
// This interface is the same as protoregistry.ExtensionTypeResolver.
type ExtensionTypeResolver interface {
	FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)
	FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
}

var _ protoregistry.ExtensionTypeResolver = ExtensionTypeResolver(nil)
var _ ExtensionTypeResolver = protoregistry.ExtensionTypeResolver(nil)

// MessageTypeResolver can resolve message types. A message type includes more than just a
// descriptor but also refers to runtime types (which could be static types generated by the
// protoc-gen-go plugin).
//
// This interface is the same as protoregistry.MessageTypeResolver.
type MessageTypeResolver interface {
	FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error)
	FindMessageByURL(url string) (protoreflect.MessageType, error)
}

var _ protoregistry.MessageTypeResolver = MessageTypeResolver(nil)
var _ MessageTypeResolver = protoregistry.MessageTypeResolver(nil)

// EnumTypeResolver can resolve enum types. An enum type includes more than just a descriptor
// but also refers to runtime types (which could be static types generated by the
// protoc-gen-go plugin).
type EnumTypeResolver interface {
	FindEnumByName(enum protoreflect.FullName) (protoreflect.EnumType, error)
}

// SerializationResolver is a named interface that can be used as a resolver for
// various marshalling and unmarshalling operations. For example it can be used
// to recognize extensions when unmarshalling the binary format. And it can be used
// for recognizing extensions and the contents of Any messages when marshalling and
// unmarshalling the JSON and text formats.
//
// This type can be assigned to the following fields:
//   - proto.UnmarshalOptions.resolver
//   - protojson.MarshalOptions.resolver
//   - protojson.UnmarshalOptions.resolver
//   - prototext.MarshalOptions.resolver
//   - prototext.UnmarshalOptions.resolver
type SerializationResolver interface {
	ExtensionTypeResolver
	MessageTypeResolver
}

// TypeResolver can resolve all types: extensions, messages, and enums.
type TypeResolver interface {
	ExtensionTypeResolver
	MessageTypeResolver
	EnumTypeResolver
}

// DependencyResolver can resolve dependencies, which is needed when
// constructing a [protoreflect.FileDescriptor] from a FileDescriptorProto.
//
// This interface is the same as [protodesc.Resolver].
type DependencyResolver interface {
	FileResolver
	DescriptorResolver
}

var _ protodesc.Resolver = DependencyResolver(nil)
var _ DependencyResolver = protodesc.Resolver(nil)

// TODO: Other sometimes-used interfaces that might warrant a named type:
// 			interface { DependencyResolver; ExtensionResolver }
//			interface { DescriptorResolver; ExtensionResolver }
//			interface { RangeExtensionsByMessage(protoreflect.FullName, func(protoreflect.ExtensionType) bool) }
//			interface { Resolver; AsTypePool() TypePool }
//		Admittedly, we've got a LOT of named types already, so maybe best to leave these out...

// Resolver is a comprehensive resolver interface with methods for resolving all kinds
// of descriptors.
//
// The AsTypeResolver method returns a view of the resolver as a TypeResolver. In most
// cases, the returned types will be dynamic types constructed using the resolver's
// descriptors and the [google.golang.org/protobuf/types/dynamicpb] package.
type Resolver interface {
	DescriptorPool
	ExtensionPool
	MessageResolver
	AsTypeResolver() TypeResolver
}

// DescriptorKind represents the kind of a descriptor. Unlike other
// descriptor-related APIs, DescriptorKind distinguishes between
// extension fields (DescriptorKindExtension) and "regular", non-extension
// fields (DescriptorKindField).
type DescriptorKind int

// The various supported DescriptorKind values.
const (
	DescriptorKindUnknown = DescriptorKind(iota)
	DescriptorKindFile
	DescriptorKindMessage
	DescriptorKindField
	DescriptorKindOneof
	DescriptorKindEnum
	DescriptorKindEnumValue
	DescriptorKindExtension
	DescriptorKindService
	DescriptorKindMethod
)

// KindOf returns the DescriptorKind of the given descriptor d.
func KindOf(d protoreflect.Descriptor) DescriptorKind {
	switch d := d.(type) {
	case protoreflect.FileDescriptor:
		return DescriptorKindFile
	case protoreflect.MessageDescriptor:
		return DescriptorKindMessage
	case protoreflect.FieldDescriptor:
		if d.IsExtension() {
			return DescriptorKindExtension
		}
		return DescriptorKindField
	case protoreflect.OneofDescriptor:
		return DescriptorKindOneof
	case protoreflect.EnumDescriptor:
		return DescriptorKindEnum
	case protoreflect.EnumValueDescriptor:
		return DescriptorKindEnumValue
	case protoreflect.ServiceDescriptor:
		return DescriptorKindService
	case protoreflect.MethodDescriptor:
		return DescriptorKindMethod
	default:
		return DescriptorKindUnknown
	}
}

// String returns a textual representation of k.
func (k DescriptorKind) String() string {
	switch k {
	case DescriptorKindFile:
		return "file"
	case DescriptorKindMessage:
		return "message"
	case DescriptorKindField:
		return "field"
	case DescriptorKindOneof:
		return "oneof"
	case DescriptorKindEnum:
		return "enum"
	case DescriptorKindEnumValue:
		return "enum value"
	case DescriptorKindExtension:
		return "extension"
	case DescriptorKindService:
		return "service"
	case DescriptorKindMethod:
		return "method"
	case DescriptorKindUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unknown kind (%d)", k)
	}
}

func (k DescriptorKind) withArticle() string {
	switch k {
	case DescriptorKindFile:
		return "a file"
	case DescriptorKindMessage:
		return "a message"
	case DescriptorKindField:
		return "a field"
	case DescriptorKindOneof:
		return "a oneof"
	case DescriptorKindEnum:
		return "an enum"
	case DescriptorKindEnumValue:
		return "an enum value"
	case DescriptorKindExtension:
		return "an extension"
	case DescriptorKindService:
		return "a service"
	case DescriptorKindMethod:
		return "a method"
	case DescriptorKindUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("unknown kind (%d)", k)
	}
}

// NewNotFoundError returns an error that wraps ErrNotFound with context
// indicating the given name as the element that could not be found.
//
// The parameter is generic so that it will accept both plain strings
// and named string types like protoreflect.FullName.
func NewNotFoundError[T ~string](name T) error {
	return fmt.Errorf("%s: %w", name, ErrNotFound)
}

// ErrUnexpectedType is an error that indicates a descriptor was resolved for
// a given URL or name, but it is of the wrong type. So a query may have been
// expecting a service descriptor, for example, but instead the queried name
// resolved to an extension descriptor.
//
// See NewUnexpectedTypeError.
type ErrUnexpectedType struct {
	// Only one of URL or Name will be set, depending on whether
	// the type was looked up by URL or name. These fields indicate
	// the query that resulted in a descriptor of the wrong type.
	URL  string
	Name protoreflect.FullName
	// The kind of descriptor that was expected.
	Expecting DescriptorKind
	// The kind of descriptor that was actually found.
	Actual DescriptorKind
	// Optional: the descriptor that was actually found. This may
	// be nil. If non-nil, this is the descriptor instance that
	// was resolved whose kind is Actual instead of Expecting.
	Descriptor protoreflect.Descriptor
}

// NewUnexpectedTypeError constructs a new *ErrUnexpectedType based
// on the given properties. The last parameter, url, is optional. If
// empty, the returned error will indicate that the given descriptor's
// full name as the query.
func NewUnexpectedTypeError(expecting DescriptorKind, got protoreflect.Descriptor, url string) *ErrUnexpectedType {
	var name protoreflect.FullName
	if url == "" {
		name = got.FullName()
	}
	return &ErrUnexpectedType{
		URL:        url,
		Name:       name,
		Expecting:  expecting,
		Descriptor: got,
	}
}

// Error implements the error interface.
func (e *ErrUnexpectedType) Error() string {
	var queryKind, query string
	if e.URL != "" {
		queryKind = "URL"
		query = e.URL
	} else {
		queryKind = "name"
		query = string(e.Name)
	}
	return fmt.Sprintf("wrong kind of descriptor for %s %q: expected %s, got %s", queryKind, query, e.Expecting.withArticle(), e.Actual.withArticle())
}

// FindExtensionByNumber searches the given descriptor pool for the requested extension.
// This performs an inefficient search through all files and extensions in the pool.
// It returns nil if the extension is not found in the file.
func FindExtensionByNumber(res DescriptorPool, message protoreflect.FullName, field protoreflect.FieldNumber) protoreflect.ExtensionDescriptor {
	var ext protoreflect.ExtensionDescriptor
	res.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		ext = FindExtensionByNumberInFile(fd, message, field)
		return ext == nil
	})
	return ext
}

// FindExtensionByNumberInFile searches all extension in the given file for the requested
// extension. It returns nil if the extension is not found in the file.
func FindExtensionByNumberInFile(file protoreflect.FileDescriptor, message protoreflect.FullName, field protoreflect.FieldNumber) protoreflect.ExtensionDescriptor {
	return findExtension(file, message, field)
}

func findExtension(container TypeContainer, message protoreflect.FullName, field protoreflect.FieldNumber) protoreflect.FieldDescriptor {
	// search extensions in this scope
	exts := container.Extensions()
	for i, length := 0, exts.Len(); i < length; i++ {
		ext := exts.Get(i)
		if ext.Number() == field && ext.ContainingMessage().FullName() == message {
			return ext
		}
	}

	// if not found, search nested scopes
	msgs := container.Messages()
	for i, length := 0, msgs.Len(); i < length; i++ {
		msg := msgs.Get(i)
		ext := findExtension(msg, message, field)
		if ext != nil {
			return ext
		}
	}
	return nil
}

// RangeExtensionsByMessage enumerates all extensions in the given descriptor pool that
// extend the given message. It stops early if the given function returns false.
func RangeExtensionsByMessage(res DescriptorPool, message protoreflect.FullName, fn func(descriptor protoreflect.ExtensionDescriptor) bool) {
	var rangeInContext func(container TypeContainer, fn func(protoreflect.ExtensionDescriptor) bool) bool
	rangeInContext = func(container TypeContainer, fn func(protoreflect.ExtensionDescriptor) bool) bool {
		exts := container.Extensions()
		for i, length := 0, exts.Len(); i < length; i++ {
			ext := exts.Get(i)
			if ext.ContainingMessage().FullName() == message {
				if !fn(ext) {
					return false
				}
			}
		}
		msgs := container.Messages()
		for i, length := 0, msgs.Len(); i < length; i++ {
			msg := msgs.Get(i)
			if !rangeInContext(msg, fn) {
				return false
			}
		}
		return true
	}
	res.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		return rangeInContext(file, fn)
	})
}

// FindDescriptorByNameInFile searches the given file for the element with the given
// fully-qualified name. This could be used to implement the
// [DescriptorResolver.FindDescriptorByName] method for a resolver that doesn't want
// to create an index of all descriptors. This returns nil if no element with the
// given name belongs to this file.
//
// This does not perform a brute-force search of all elements to find the given name.
// It breaks up the given name into components and then descends the descriptor
// hierarchy one element at a time. If the given name does not start with the file's
// package, it immediately returns nil.
func FindDescriptorByNameInFile(file protoreflect.FileDescriptor, sym protoreflect.FullName) protoreflect.Descriptor {
	symNoPkg := string(sym)
	if file.Package() != "" {
		symNoPkg = strings.TrimPrefix(string(sym), string(file.Package())+".")
		if symNoPkg == string(sym) {
			// symbol is not in this file's package
			return nil
		}
	}
	parts := strings.Split(symNoPkg, ".")
	return findSymbolInFile(parts, file)
}

func findSymbolInFile(symbolParts []string, fd protoreflect.FileDescriptor) protoreflect.Descriptor {
	// ==1 name means it's a direct child of this file
	if len(symbolParts) == 1 {
		n := protoreflect.Name(symbolParts[0])
		if d := fd.Messages().ByName(n); d != nil {
			return d
		}
		if d := fd.Enums().ByName(n); d != nil {
			return d
		}
		if d := fd.Extensions().ByName(n); d != nil {
			return d
		}
		if d := fd.Services().ByName(n); d != nil {
			return d
		}
		// enum values are defined in the scope that encloses the enum, so
		// we have to look in all enums to find top-level enum values
		enums := fd.Enums()
		for i, length := 0, enums.Len(); i < length; i++ {
			enum := enums.Get(i)
			if d := enum.Values().ByName(n); d != nil {
				return d
			}
		}
		// not in this file
		return nil
	}

	// >1 name means it's inside a message or (if ==2) a method inside a service
	first := protoreflect.Name(symbolParts[0])
	if len(symbolParts) == 2 {
		second := protoreflect.Name(symbolParts[1])
		if svc := fd.Services().ByName(first); svc != nil {
			if d := svc.Methods().ByName(second); d != nil {
				return d
			}
			return nil
		}
	}
	rest := symbolParts[1:]
	if msg := fd.Messages().ByName(first); msg != nil {
		return findSymbolInMessage(rest, msg)
	}

	// no other option; can't be in this file
	return nil
}

func findSymbolInMessage(symbolParts []string, md protoreflect.MessageDescriptor) protoreflect.Descriptor {
	// ==1 name means it's a direct child of this message
	if len(symbolParts) == 1 {
		n := protoreflect.Name(symbolParts[0])
		if d := md.Fields().ByName(n); d != nil {
			return d
		}
		if d := md.Oneofs().ByName(n); d != nil {
			return d
		}
		if d := md.Messages().ByName(n); d != nil {
			return d
		}
		if d := md.Enums().ByName(n); d != nil {
			return d
		}
		if d := md.Extensions().ByName(n); d != nil {
			return d
		}
		// enum values are defined in the scope that encloses the enum, so
		// we have to look in all enums to find enum values at this level
		enums := md.Enums()
		for i, length := 0, enums.Len(); i < length; i++ {
			enum := enums.Get(i)
			if d := enum.Values().ByName(n); d != nil {
				return d
			}
		}
		// not in this file
		return nil
	}

	// >1 name means it's inside a nested message
	first := protoreflect.Name(symbolParts[0])
	rest := symbolParts[1:]
	if nested := md.Messages().ByName(first); nested != nil {
		return findSymbolInMessage(rest, nested)
	}

	// no other option; can't be in this message
	return nil
}

// ResolverFromPool implements the full Resolver interface on top of the
// given DescriptorPool. This can be used to upgrade a *[protoregistry.Files]
// to the Resolver interface. The AsTypeResolver method uses TypesFromResolver,
// so it returns dynamic types.
//
// See also ResolverFromPools.
func ResolverFromPool(pool DescriptorPool) Resolver {
	return &resolverFromPool{DescriptorPool: pool}
}

// ResolverFromPools (plural) is just like ResolverFromPool (singular) except it
// also accepts a TypePool that is used to implement the AsTypeResolver method.
// So instead of always returning dynamic types based on the given DescriptorPool,
// it uses the given TypePool.
func ResolverFromPools(descPool DescriptorPool, typePool TypePool) interface {
	Resolver
	AsTypePool() TypePool
} {
	return &resolverWithTypes{Resolver: ResolverFromPool(descPool), types: typePool}
}

type resolverFromPool struct {
	DescriptorPool
}

func (r *resolverFromPool) FindMessageByName(name protoreflect.FullName) (protoreflect.MessageDescriptor, error) {
	d, err := r.DescriptorPool.FindDescriptorByName(name)
	if err != nil {
		return nil, err
	}
	msg, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, NewUnexpectedTypeError(DescriptorKindMessage, d, "")
	}
	return msg, nil
}

func (r *resolverFromPool) FindExtensionByName(name protoreflect.FullName) (protoreflect.ExtensionDescriptor, error) {
	d, err := r.DescriptorPool.FindDescriptorByName(name)
	if err != nil {
		return nil, err
	}
	field, ok := d.(protoreflect.FieldDescriptor)
	if !ok {
		return nil, NewUnexpectedTypeError(DescriptorKindExtension, d, "")
	}
	if !field.IsExtension() {
		return nil, NewUnexpectedTypeError(DescriptorKindExtension, field, "")
	}
	return field, nil
}

func (r *resolverFromPool) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionDescriptor, error) {
	extd := FindExtensionByNumber(r.DescriptorPool, message, field)
	if extd == nil {
		return nil, ErrNotFound
	}
	return extd, nil
}

func (r *resolverFromPool) RangeExtensionsByMessage(message protoreflect.FullName, fn func(protoreflect.ExtensionDescriptor) bool) {
	RangeExtensionsByMessage(r.DescriptorPool, message, fn)
}

func (r *resolverFromPool) FindMessageByURL(url string) (protoreflect.MessageDescriptor, error) {
	return r.FindMessageByName(TypeNameFromURL(url))
}

func (r *resolverFromPool) AsTypeResolver() TypeResolver {
	return TypesFromResolver(r)
}

type resolverWithTypes struct {
	Resolver
	types TypePool
}

func (r *resolverWithTypes) AsTypeResolver() TypeResolver {
	return r.types
}

func (r *resolverWithTypes) AsTypePool() TypePool {
	return r.types
}