// Code generated by protoc-gen-gosrcinfo. DO NOT EDIT.
// source: desc_test_options.proto

package testdata

import (
	sourceinfo "github.com/jhump/protoreflect/v2/sourceinfo"
)

func init() {
	srcInfo := []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x7c, 0x95, 0x4b, 0x53, 0x23, 0xb7,
		0x17, 0xc5, 0xa5, 0x3e, 0x92, 0xad, 0x96, 0x9f, 0x7d, 0xda, 0x2f, 0x8c, 0x6d, 0x8c, 0x8d, 0xb1,
		0x61, 0xc0, 0xff, 0x1a, 0x98, 0x1a, 0xf8, 0xd7, 0x40, 0xf6, 0xd9, 0xe7, 0x0b, 0xe4, 0x41, 0xa5,
		0xb2, 0x60, 0x48, 0x95, 0x61, 0x91, 0x6f, 0x9f, 0x52, 0xdf, 0xeb, 0xa9, 0x5e, 0x65, 0x77, 0x74,
		0xac, 0xfe, 0xf9, 0xe8, 0xea, 0xf6, 0xed, 0xd8, 0xa0, 0x33, 0xe6, 0x27, 0x1b, 0x43, 0xb4, 0x6d,
		0xc2, 0x18, 0x26, 0x15, 0x88, 0xcc, 0xfc, 0x1c, 0xf3, 0x98, 0x85, 0x96, 0xc8, 0x10, 0x6d, 0x46,
		0x38, 0x53, 0x26, 0x13, 0x86, 0x68, 0x98, 0xeb, 0x98, 0x47, 0xdb, 0xa4, 0x0b, 0x26, 0xda, 0xe4,
		0x36, 0x0d, 0x91, 0x87, 0x4d, 0x8c, 0x11, 0x4d, 0x93, 0x11, 0xa1, 0xa9, 0xda, 0x25, 0xbf, 0x2f,
		0xda, 0x13, 0x79, 0x31, 0x14, 0x6d, 0x89, 0x7c, 0x34, 0x13, 0x0d, 0x22, 0x3f, 0xbb, 0x50, 0x62,
		0xdb, 0xf4, 0x84, 0x68, 0x89, 0x4e, 0xd8, 0x56, 0x3b, 0xd2, 0xdf, 0xb7, 0x9b, 0x6b, 0xd1, 0x2e,
		0xf9, 0x42, 0xb4, 0x9e, 0xe8, 0x14, 0x63, 0xd1, 0x69, 0xff, 0x64, 0x21, 0x1a, 0x44, 0xe7, 0xfc,
		0xb2, 0xc2, 0x64, 0x44, 0x57, 0x31, 0x59, 0x0d, 0x93, 0xb9, 0xe4, 0x0b, 0x26, 0xf3, 0x44, 0xb7,
		0x18, 0x89, 0xb6, 0x44, 0x77, 0x2c, 0x98, 0x0c, 0x44, 0xb7, 0xc2, 0xa4, 0x60, 0x85, 0x19, 0x4b,
		0x30, 0x10, 0x0c, 0x97, 0xd5, 0x0e, 0x64, 0x44, 0xd1, 0x5c, 0x89, 0x76, 0xc9, 0x17, 0x22, 0x3c,
		0x41, 0x25, 0xc2, 0x12, 0x1c, 0xcf, 0x45, 0xa7, 0x67, 0x97, 0x9b, 0x0a, 0xe3, 0x88, 0x32, 0xec,
		0x2a, 0xdb, 0xd5, 0x30, 0xae, 0xf2, 0x05, 0xe3, 0x3c, 0x51, 0xea, 0xf9, 0x9c, 0x25, 0xca, 0xc9,
		0x99, 0x68, 0x10, 0xe5, 0x6a, 0x5b, 0x61, 0x3c, 0x31, 0x08, 0x9f, 0x2a, 0xdb, 0xd7, 0x30, 0xde,
		0x25, 0x5f, 0x30, 0x3e, 0xed, 0x29, 0x4e, 0x44, 0x5b, 0x62, 0x30, 0xd5, 0x3d, 0x20, 0x06, 0x9b,
		0xeb, 0x0a, 0xd3, 0x20, 0x86, 0x9a, 0xa6, 0x51, 0xc3, 0x34, 0x5c, 0xf2, 0x05, 0xd3, 0xf0, 0xc4,
		0x50, 0xd3, 0x34, 0x2c, 0x31, 0xd4, 0x34, 0x0d, 0x10, 0x43, 0x4d, 0xd3, 0x24, 0x46, 0xe1, 0xba,
		0xb2, 0x9b, 0x35, 0x4c, 0x3a, 0xec, 0x48, 0x31, 0x29, 0xf1, 0xa8, 0x98, 0x88, 0xb6, 0xc4, 0xe8,
		0xe4, 0x5c, 0x34, 0x88, 0xd1, 0xc5, 0x95, 0x56, 0xfb, 0xc4, 0x9c, 0x49, 0xb5, 0x03, 0x31, 0xd5,
		0xfb, 0x0b, 0x19, 0x71, 0xd2, 0x54, 0xed, 0x92, 0x2f, 0xc4, 0xe0, 0x89, 0xa9, 0x56, 0x3b, 0x58,
		0x62, 0xaa, 0xf7, 0x17, 0x40, 0x4c, 0xb5, 0x0d, 0x72, 0xe2, 0x34, 0x5c, 0x55, 0x76, 0x5e, 0xc3,
		0xe4, 0x2e, 0xf9, 0x82, 0xc9, 0x3d, 0x71, 0xaa, 0xe7, 0xcb, 0x2d, 0x71, 0x3a, 0x59, 0x8a, 0x06,
		0x71, 0xba, 0xde, 0x55, 0x98, 0x48, 0xcc, 0xc2, 0x4d, 0x65, 0xc7, 0x1a, 0x26, 0xba, 0xe4, 0x0b,
		0x26, 0x7a, 0x62, 0xa6, 0xd5, 0x8e, 0x96, 0x98, 0x4d, 0xa5, 0xe3, 0x22, 0x88, 0xd9, 0xe5, 0xa7,
		0x0a, 0xd3, 0x22, 0xe6, 0x9a, 0xa6, 0x55, 0xc3, 0xb4, 0x5c, 0xf2, 0x05, 0xd3, 0xf2, 0xc4, 0x5c,
		0xd3, 0xb4, 0x2c, 0x31, 0xd7, 0x34, 0x2d, 0x10, 0x73, 0x4d, 0xd3, 0x26, 0x16, 0x7a, 0xf7, 0xed,
		0x1a, 0xa6, 0xed, 0x92, 0x2f, 0x98, 0xb6, 0x27, 0x16, 0x5a, 0xed, 0xb6, 0x25, 0x16, 0x27, 0x72,
		0x23, 0x6d, 0x10, 0x8b, 0xcd, 0xf1, 0x35, 0x3e, 0x37, 0x17, 0x52, 0xed, 0x0e, 0xb1, 0x0a, 0x5f,
		0xaa, 0x1d, 0x9d, 0x8c, 0x38, 0xd7, 0xd7, 0xb8, 0xe3, 0x92, 0x2f, 0xc4, 0x4e, 0x83, 0x58, 0x15,
		0x17, 0xa2, 0x2d, 0xb1, 0xda, 0x48, 0x82, 0x0e, 0x88, 0xd5, 0xfe, 0xbe, 0xc2, 0x74, 0x89, 0x75,
		0xb8, 0xab, 0xec, 0x6e, 0x0d, 0xd3, 0x75, 0xc9, 0x17, 0x4c, 0xb7, 0x41, 0xac, 0x0b, 0xb9, 0xfa,
		0xae, 0x25, 0xd6, 0x2b, 0xa9, 0x47, 0x17, 0xc4, 0xfa, 0xe6, 0xb3, 0x06, 0xbb, 0x34, 0x57, 0x12,
		0xac, 0x47, 0x6c, 0xb5, 0x0d, 0x7a, 0x19, 0x71, 0xd9, 0x94, 0x00, 0x3d, 0x97, 0x7c, 0x21, 0xf6,
		0x3c, 0xb1, 0xd5, 0x36, 0xe8, 0x59, 0x62, 0xab, 0x6d, 0xd0, 0x03, 0xb1, 0xd5, 0x36, 0xe8, 0x13,
		0x3b, 0x2d, 0x7c, 0xbf, 0x86, 0xe9, 0xbb, 0xe4, 0x0b, 0xa6, 0xef, 0x89, 0x9d, 0x16, 0xbe, 0x6f,
		0x89, 0x9d, 0x16, 0xbe, 0x0f, 0x62, 0xb7, 0xde, 0xc5, 0xcf, 0x31, 0x73, 0x86, 0xee, 0xc6, 0xfc,
		0xcf, 0x4e, 0x37, 0xcb, 0x5f, 0x5e, 0x0e, 0xef, 0xcb, 0xd7, 0x97, 0xc3, 0xe1, 0xd7, 0x3f, 0x5f,
		0x96, 0x1f, 0x87, 0x97, 0x3f, 0x96, 0xbf, 0xfd, 0xb3, 0xfc, 0xfd, 0xe3, 0xf0, 0xfe, 0xf6, 0xba,
		0x7c, 0xfb, 0xfb, 0xfd, 0xaf, 0xb7, 0xef, 0x87, 0x98, 0x1e, 0x77, 0x69, 0xe2, 0xdd, 0x84, 0xd3,
		0xd8, 0x8a, 0xce, 0x99, 0xcc, 0x10, 0xb7, 0xe1, 0x2c, 0xb6, 0xa3, 0x4f, 0x0b, 0x97, 0x56, 0xfd,
		0xe3, 0xca, 0x13, 0xb7, 0xc5, 0xf8, 0xb8, 0xb2, 0xc4, 0xed, 0x64, 0x7a, 0x5c, 0x81, 0xb8, 0x9d,
		0x2f, 0x14, 0x62, 0x89, 0x7d, 0x38, 0xd7, 0x9f, 0xd2, 0x5c, 0xdc, 0xff, 0x80, 0xa4, 0xc9, 0xb8,
		0xff, 0x01, 0x49, 0xb3, 0x71, 0x3f, 0x99, 0x1d, 0x57, 0x20, 0xf6, 0x67, 0xcb, 0xb8, 0x8f, 0x99,
		0x37, 0x74, 0x77, 0xe6, 0x8b, 0x9d, 0xae, 0xe4, 0x20, 0x2f, 0xdf, 0x3f, 0x5e, 0xff, 0xeb, 0x14,
		0x3e, 0xc5, 0xb9, 0xf3, 0xc3, 0x14, 0xc0, 0x57, 0xa7, 0xb8, 0x0f, 0x4c, 0x58, 0x2f, 0x49, 0xef,
		0x43, 0xe7, 0xb8, 0xca, 0x88, 0xfb, 0x7e, 0xa1, 0xd7, 0xf8, 0xd5, 0xfc, 0x5f, 0xae, 0xb1, 0x20,
		0x1e, 0x74, 0xcc, 0x14, 0x19, 0xf1, 0xb5, 0x29, 0xef, 0x52, 0xe1, 0x92, 0x2f, 0xf5, 0x2f, 0x3c,
		0xf1, 0xa0, 0xf5, 0x2f, 0x2c, 0xf1, 0xa0, 0x63, 0xa6, 0x00, 0xf1, 0xa0, 0x63, 0x86, 0xc4, 0xa3,
		0x62, 0x58, 0xc3, 0xd0, 0x25, 0x5f, 0x30, 0xf4, 0xc4, 0xa3, 0x76, 0x03, 0x2d, 0xf1, 0x38, 0x16,
		0x0c, 0x41, 0x3c, 0x56, 0x98, 0x14, 0xec, 0x5b, 0xfa, 0xfc, 0x25, 0x62, 0x49, 0x3c, 0x29, 0xb1,
		0xcc, 0x88, 0x6f, 0xfa, 0x99, 0x28, 0x5d, 0xf2, 0x85, 0x58, 0x7a, 0xe2, 0x49, 0x83, 0x95, 0x96,
		0x78, 0xd2, 0x60, 0x25, 0x88, 0x27, 0x0d, 0x36, 0x20, 0x9e, 0x15, 0x33, 0xa8, 0x61, 0x06, 0x2e,
		0xf9, 0x82, 0x19, 0x78, 0xe2, 0x59, 0x83, 0x0d, 0x2c, 0xf1, 0xac, 0xc1, 0x06, 0x20, 0x9e, 0x57,
		0xdb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x10, 0x7a, 0xa7, 0x95, 0x07, 0x00, 0x00,
	}
	sourceinfo.Register("desc_test_options.proto", srcInfo)
}