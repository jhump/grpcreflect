// Code generated by protoc-gen-gosrcinfo. DO NOT EDIT.
// source: desc_test1.proto

package testprotos

import (
	sourceinfo "github.com/jhump/protoreflect/v2/sourceinfo"
)

func init() {
	srcInfo := []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x94, 0x58, 0xff, 0x4f, 0x1b, 0xc7,
		0x12, 0xbf, 0xbd, 0xb9, 0x3b, 0x2f, 0x03, 0x14, 0x7b, 0x20, 0x60, 0x9f, 0x31, 0x1c, 0x18, 0xc7,
		0x5f, 0x12, 0xa0, 0x36, 0x04, 0x02, 0x84, 0x06, 0x68, 0x89, 0x78, 0x29, 0x21, 0x11, 0xe5, 0xf1,
		0x52, 0x55, 0x11, 0x0f, 0x92, 0x4b, 0x83, 0xc0, 0x76, 0x82, 0x4d, 0x95, 0xbc, 0xa8, 0xea, 0x93,
		0xaa, 0xea, 0xf5, 0xcf, 0x7e, 0x9a, 0xbb, 0x5b, 0xfb, 0xf6, 0xe4, 0xa8, 0xea, 0x6f, 0xfb, 0x19,
		0xcf, 0x7e, 0x3e, 0xb3, 0xb3, 0xb3, 0xb3, 0x7b, 0xc6, 0x14, 0xd9, 0x86, 0xf1, 0xa7, 0x10, 0x28,
		0x51, 0x8c, 0x10, 0x18, 0x06, 0xf1, 0x48, 0x12, 0x98, 0xc6, 0xf7, 0x38, 0x84, 0xa6, 0x1c, 0x0e,
		0x87, 0x12, 0x85, 0x49, 0x60, 0x19, 0xe3, 0x58, 0x42, 0xd3, 0x32, 0xc8, 0x4a, 0x19, 0xdf, 0x0a,
		0x37, 0xe7, 0x7d, 0xdb, 0x6e, 0x36, 0xfd, 0x56, 0xd7, 0x7b, 0xdb, 0xbe, 0xf1, 0x4e, 0xfc, 0x4e,
		0xf7, 0x99, 0xdf, 0xe9, 0x9c, 0xff, 0xec, 0x23, 0x22, 0x82, 0x65, 0x08, 0x82, 0x94, 0x1c, 0xc7,
		0x2a, 0x5a, 0x96, 0x01, 0x06, 0x59, 0x43, 0xb2, 0x31, 0xe4, 0xe6, 0xb5, 0x49, 0x47, 0x7e, 0xa7,
		0xeb, 0xbf, 0xe9, 0x4d, 0x1b, 0x41, 0x9b, 0x5d, 0x05, 0xc1, 0x50, 0xba, 0x80, 0x0d, 0x74, 0x18,
		0xf1, 0xd4, 0xe1, 0xf4, 0xbd, 0x8c, 0x3b, 0xaf, 0x4d, 0xdd, 0x6d, 0xb5, 0xbb, 0xef, 0xfc, 0x9b,
		0x04, 0xc3, 0x57, 0x98, 0x0a, 0xe7, 0x08, 0x82, 0xe1, 0xec, 0x7d, 0xdc, 0x51, 0xd8, 0x21, 0x6b,
		0x34, 0x9b, 0xce, 0xb9, 0xcb, 0x83, 0x48, 0x62, 0xb1, 0x7b, 0xfe, 0xc7, 0xae, 0xdf, 0xea, 0x5c,
		0xb6, 0x5b, 0x1d, 0xaf, 0x52, 0xaf, 0x22, 0x2e, 0xa0, 0x8c, 0x18, 0x0c, 0x82, 0x31, 0xef, 0xa9,
		0x3b, 0xae, 0x31, 0xbc, 0xbd, 0x3e, 0xff, 0xb9, 0x83, 0x98, 0xc6, 0x21, 0xe5, 0x65, 0x12, 0x8c,
		0xce, 0xd6, 0xe3, 0x16, 0x8b, 0x27, 0x56, 0xe2, 0x16, 0x9b, 0x60, 0xac, 0xba, 0x18, 0xb7, 0x08,
		0x82, 0xb1, 0xa5, 0x95, 0xb8, 0x05, 0x08, 0xc6, 0xd6, 0x36, 0xe2, 0x16, 0x49, 0x30, 0xb6, 0xf9,
		0x0f, 0xcc, 0x20, 0xf6, 0x2c, 0x26, 0xc1, 0xd8, 0xd6, 0x01, 0xae, 0xab, 0x20, 0x39, 0x59, 0x94,
		0xad, 0xe4, 0xdc, 0x92, 0x16, 0xe5, 0x8f, 0x7e, 0x77, 0x70, 0xbe, 0x7a, 0xec, 0x41, 0xc6, 0xc8,
		0x5b, 0xc7, 0x86, 0x62, 0x07, 0x83, 0x77, 0x7a, 0xc2, 0xcb, 0xcd, 0xb9, 0x33, 0x1a, 0xd9, 0x77,
		0xbe, 0xff, 0xfe, 0xfa, 0x53, 0x48, 0xb4, 0xdf, 0xba, 0x6d, 0x22, 0x12, 0x0e, 0xf7, 0xe7, 0x08,
		0x82, 0x89, 0xd2, 0x03, 0xac, 0xe2, 0x48, 0xdf, 0x66, 0x1a, 0x04, 0x93, 0x95, 0x15, 0x77, 0x42,
		0x23, 0x3a, 0xdd, 0x3d, 0xfc, 0xe7, 0x7e, 0x1d, 0x71, 0x02, 0x47, 0xe3, 0xae, 0x82, 0x7d, 0x97,
		0x92, 0x56, 0x93, 0x60, 0xb2, 0xde, 0x48, 0xd0, 0x0a, 0x82, 0xec, 0x60, 0xda, 0x46, 0x92, 0x56,
		0x04, 0xbe, 0x49, 0x5a, 0xae, 0xeb, 0x6c, 0xbd, 0x81, 0x0b, 0xfd, 0x55, 0x73, 0xac, 0x79, 0xef,
		0xa1, 0x9b, 0xd1, 0xf7, 0xb9, 0xdd, 0xd6, 0xd6, 0x69, 0xf2, 0xae, 0xe6, 0xbd, 0x8a, 0x6e, 0xb3,
		0x09, 0xf2, 0xd5, 0x65, 0xdd, 0x26, 0x08, 0xf2, 0x5f, 0xaf, 0xe8, 0x36, 0x20, 0xc8, 0xaf, 0xad,
		0x6b, 0xaa, 0x82, 0xa0, 0xe0, 0xad, 0x27, 0x54, 0x2f, 0xce, 0x6f, 0x74, 0x55, 0x61, 0xb1, 0x9b,
		0xae, 0x2a, 0x6c, 0x82, 0x42, 0x75, 0x49, 0xb7, 0x31, 0xdd, 0x72, 0x43, 0xb7, 0x01, 0x41, 0xe1,
		0xc1, 0x9a, 0xa6, 0x6a, 0x12, 0xcc, 0x0e, 0x50, 0xfd, 0x8f, 0xae, 0x6a, 0x5a, 0xec, 0xa6, 0xab,
		0x9a, 0x36, 0xc1, 0x6c, 0x42, 0x95, 0x17, 0x31, 0x9b, 0x50, 0x35, 0x81, 0x60, 0x36, 0xa1, 0x0a,
		0x04, 0x73, 0xde, 0x5e, 0x42, 0xf5, 0x4d, 0xcb, 0xd7, 0x55, 0xc1, 0x62, 0x37, 0x5d, 0x15, 0x1c,
		0x82, 0xb9, 0xea, 0x86, 0x6e, 0x13, 0x04, 0x73, 0x9b, 0xdb, 0xba, 0x8d, 0x25, 0x76, 0x76, 0x35,
		0x55, 0x8b, 0xa0, 0xe8, 0x3d, 0x49, 0xa8, 0x9e, 0xb7, 0xf4, 0xfa, 0x35, 0xad, 0xc0, 0x4d, 0x57,
		0xb5, 0x1c, 0x82, 0x62, 0x55, 0x57, 0xb0, 0x04, 0x41, 0xf1, 0x9b, 0x5d, 0xdd, 0x06, 0x04, 0xc5,
		0xef, 0xf6, 0xb1, 0x18, 0x53, 0xb5, 0x09, 0x4a, 0xde, 0x37, 0x6e, 0x5a, 0x53, 0x4d, 0x8a, 0xda,
		0x16, 0x7b, 0xe9, 0xa2, 0xb6, 0x43, 0x50, 0xaa, 0xae, 0xe9, 0x36, 0x41, 0x50, 0x5a, 0xd7, 0x97,
		0x6f, 0x03, 0x41, 0xe9, 0xd1, 0xb6, 0x26, 0xea, 0x10, 0x94, 0xbd, 0x47, 0x09, 0xd1, 0x6e, 0x42,
		0xd4, 0xb1, 0xd8, 0x4b, 0x17, 0x75, 0x78, 0x66, 0x75, 0x55, 0xb7, 0x09, 0x82, 0xf2, 0x83, 0x75,
		0xdd, 0x06, 0x04, 0xe5, 0xcd, 0x2d, 0x2c, 0xaa, 0xc6, 0xc3, 0xa7, 0xa6, 0x96, 0xdd, 0x73, 0x49,
		0x93, 0xfc, 0x14, 0xa4, 0xb7, 0xd7, 0x64, 0x82, 0x43, 0x53, 0xcb, 0x7a, 0x71, 0x8b, 0x43, 0x50,
		0x9b, 0x7b, 0x18, 0xb7, 0x08, 0x82, 0xda, 0xc6, 0x76, 0xdc, 0x02, 0x04, 0xb5, 0x9d, 0x5d, 0xf4,
		0xc2, 0x0b, 0x81, 0xa5, 0x16, 0xd3, 0x6b, 0x83, 0x36, 0x32, 0x6a, 0xff, 0x81, 0xce, 0x62, 0x3a,
		0xdb, 0xc7, 0x0e, 0xc1, 0x62, 0x6e, 0xb1, 0x8f, 0x05, 0xc1, 0xe2, 0x52, 0xbd, 0x8f, 0x81, 0x60,
		0x71, 0xf5, 0x01, 0xce, 0x47, 0x0a, 0x82, 0xac, 0xe5, 0x74, 0xbd, 0x30, 0x70, 0x35, 0x6a, 0x0e,
		0x1f, 0xc5, 0xe5, 0x74, 0x16, 0xc7, 0x14, 0x76, 0xc8, 0x5a, 0xce, 0x7d, 0xdd, 0x27, 0xe5, 0x33,
		0x58, 0x9f, 0xe8, 0x07, 0xc1, 0xe7, 0xaf, 0x9e, 0x9f, 0xc6, 0xbb, 0x7c, 0x21, 0x72, 0x6f, 0x5d,
		0x95, 0x1b, 0x43, 0x6e, 0x76, 0xc0, 0x85, 0x18, 0x76, 0xd5, 0xe0, 0x36, 0x0c, 0xfa, 0xe9, 0xea,
		0xe8, 0x14, 0x16, 0x39, 0xb4, 0xb0, 0x93, 0xae, 0xa5, 0xf3, 0x5f, 0xe8, 0xa4, 0x81, 0x54, 0xd4,
		0x43, 0xd7, 0xd2, 0x93, 0x7d, 0x6c, 0x12, 0xac, 0xe5, 0xdc, 0x1e, 0x89, 0x20, 0x78, 0x38, 0x98,
		0xa4, 0xd1, 0x27, 0x11, 0x81, 0x57, 0x9f, 0x84, 0x7b, 0xe5, 0xc3, 0x9c, 0x8b, 0x05, 0x8e, 0x9f,
		0xe3, 0xd8, 0x92, 0x77, 0x07, 0xd4, 0x75, 0x10, 0x76, 0xb0, 0x03, 0x5b, 0x32, 0xad, 0x90, 0x43,
		0xb0, 0x95, 0x99, 0x51, 0x48, 0x10, 0x6c, 0xcd, 0xce, 0x29, 0x04, 0x04, 0x5b, 0x0b, 0x25, 0x9c,
		0x09, 0x68, 0x05, 0xc1, 0xb6, 0x7c, 0x34, 0x68, 0x6f, 0x43, 0x6f, 0xce, 0xfa, 0x76, 0x8f, 0x57,
		0x38, 0x04, 0xdb, 0x99, 0x15, 0x85, 0x78, 0xf2, 0xea, 0xba, 0x42, 0x40, 0xb0, 0xbd, 0xb9, 0x85,
		0x5e, 0xc0, 0x6b, 0x92, 0xf5, 0x58, 0xee, 0x2c, 0x0d, 0xdc, 0xd1, 0xd0, 0x9f, 0xfb, 0xc3, 0x63,
		0x99, 0xc6, 0xd1, 0x10, 0x39, 0x64, 0x3d, 0xce, 0xec, 0x2c, 0xa8, 0x1f, 0x05, 0xc1, 0x4e, 0xa9,
		0xaa, 0x10, 0x10, 0xec, 0xdc, 0x5f, 0x8c, 0x32, 0x01, 0x04, 0x7b, 0xb2, 0x98, 0xcc, 0x84, 0x7a,
		0xce, 0x04, 0x6d, 0x6c, 0xaf, 0x17, 0x31, 0x37, 0xb0, 0xbd, 0x4c, 0x5e, 0x21, 0x41, 0xb0, 0x37,
		0xad, 0xf2, 0xc2, 0x4d, 0x6b, 0x6f, 0x6e, 0x1e, 0xef, 0x23, 0x77, 0x17, 0xeb, 0x89, 0xf1, 0x56,
		0xb8, 0xde, 0x5f, 0xbc, 0x57, 0xc2, 0xb7, 0x16, 0xaf, 0xfc, 0x89, 0x74, 0x83, 0x1c, 0x0a, 0xde,
		0x9a, 0x03, 0xf9, 0x61, 0x50, 0x7b, 0x65, 0x1d, 0x11, 0xec, 0xcd, 0x41, 0x14, 0x91, 0x08, 0xf6,
		0xe6, 0x20, 0xf3, 0x4e, 0x21, 0x41, 0x70, 0x70, 0x79, 0xad, 0x10, 0x10, 0x1c, 0xb4, 0xdf, 0x63,
		0x29, 0xe0, 0x15, 0x04, 0x4f, 0x65, 0x2d, 0x51, 0xb1, 0xcd, 0xf3, 0xf7, 0x67, 0x6f, 0x2f, 0xfd,
		0xeb, 0x37, 0x75, 0x45, 0xcf, 0x9b, 0xf2, 0x54, 0xba, 0x0a, 0xf1, 0xac, 0x7c, 0x49, 0x21, 0x20,
		0x78, 0x5a, 0xa9, 0x46, 0x84, 0x26, 0xc1, 0xa1, 0xac, 0x7e, 0x89, 0xb0, 0xa1, 0x08, 0xb9, 0x9d,
		0x1d, 0xca, 0x9c, 0x42, 0x82, 0xe0, 0xd0, 0x5d, 0x50, 0x08, 0x08, 0x0e, 0xcb, 0x95, 0x88, 0x10,
		0x08, 0x8e, 0xbe, 0x4c, 0xb8, 0xa2, 0x08, 0x79, 0x13, 0x8e, 0x7a, 0x84, 0xbc, 0x09, 0x47, 0x3d,
		0x42, 0xde, 0x84, 0xa3, 0x1e, 0xa1, 0x45, 0xf0, 0x42, 0xae, 0x7f, 0x89, 0x70, 0x55, 0x11, 0xf2,
		0x05, 0xf1, 0x42, 0x96, 0x15, 0x12, 0x04, 0x2f, 0x2a, 0x0d, 0x85, 0x80, 0xe0, 0xc5, 0x83, 0x35,
		0x1c, 0x09, 0x08, 0x6d, 0xb2, 0x8e, 0xe5, 0x8f, 0x43, 0xd1, 0x6f, 0xdc, 0xf7, 0x8f, 0x7b, 0x3b,
		0x61, 0xdb, 0x04, 0xc7, 0x99, 0x49, 0x85, 0x04, 0xc1, 0xf1, 0x94, 0xa7, 0x10, 0x10, 0x1c, 0x17,
		0x17, 0xc2, 0xb0, 0xc0, 0x0a, 0x59, 0xdc, 0x29, 0x2d, 0xae, 0xe3, 0xf6, 0xeb, 0xab, 0xa3, 0xe3,
		0xf6, 0xf5, 0x75, 0x14, 0x16, 0x58, 0x3a, 0x85, 0x13, 0xa2, 0x05, 0x74, 0xf8, 0x37, 0x2e, 0x93,
		0x93, 0xf4, 0x7d, 0xf7, 0x8e, 0x7e, 0xf7, 0xfb, 0xe7, 0xdd, 0x6b, 0xbf, 0x13, 0x76, 0x81, 0xc0,
		0xcb, 0x62, 0xb7, 0x6c, 0x1f, 0xdb, 0x04, 0x27, 0xb9, 0xd9, 0x3e, 0x16, 0x04, 0x27, 0x5e, 0xb9,
		0x8f, 0x81, 0xe0, 0xa4, 0x76, 0x2f, 0x68, 0x35, 0x8c, 0x05, 0xc1, 0x69, 0xfa, 0x5e, 0xa2, 0xd5,
		0x74, 0xba, 0xed, 0x56, 0x5c, 0x84, 0x4f, 0xf5, 0x69, 0x4c, 0x84, 0x9f, 0x34, 0xa7, 0x31, 0x11,
		0x2e, 0xa2, 0x53, 0xef, 0x6e, 0x1f, 0x03, 0xc1, 0x69, 0xb5, 0x16, 0xf4, 0x6b, 0xc6, 0x26, 0xc1,
		0xcb, 0x74, 0x2d, 0xf1, 0x34, 0x7f, 0xd3, 0x6e, 0xdf, 0xc4, 0x34, 0x78, 0x2f, 0x5f, 0xc6, 0x34,
		0xf8, 0x66, 0x7e, 0x19, 0xd3, 0xe0, 0x48, 0x5f, 0x7a, 0xa5, 0x3e, 0x06, 0x82, 0x97, 0x95, 0x2a,
		0xce, 0x71, 0xc6, 0xa5, 0x41, 0xd6, 0x4f, 0xf2, 0xdf, 0x43, 0x09, 0x89, 0xf3, 0x6e, 0x93, 0xdf,
		0x85, 0x41, 0x7e, 0x25, 0x67, 0xe2, 0xa7, 0xaf, 0xc6, 0xa3, 0x43, 0xe8, 0x10, 0xbc, 0x4a, 0xcf,
		0x26, 0x0e, 0x61, 0xa7, 0x7b, 0xa3, 0x4a, 0xc6, 0xb1, 0xd9, 0x41, 0x6d, 0x36, 0xdf, 0xb1, 0xaf,
		0xa6, 0xd4, 0x99, 0xe1, 0xdb, 0xf5, 0x55, 0x61, 0x26, 0x62, 0x4a, 0x11, 0x9c, 0xa5, 0x67, 0x12,
		0x4c, 0x97, 0xad, 0xae, 0x62, 0x4a, 0xd9, 0xec, 0x70, 0x47, 0x21, 0x41, 0x70, 0x36, 0xa9, 0x6a,
		0x3b, 0x05, 0x04, 0x67, 0xd3, 0x85, 0xa8, 0xf8, 0x24, 0x59, 0x17, 0xf2, 0xb5, 0x2a, 0x3e, 0x69,
		0x11, 0x5c, 0xf4, 0x8a, 0x4f, 0xda, 0x04, 0x17, 0xbd, 0xe2, 0x93, 0x82, 0xe0, 0x62, 0x6a, 0x5e,
		0x21, 0x20, 0xb8, 0x28, 0xdd, 0x55, 0x48, 0x12, 0x5c, 0x94, 0x37, 0x71, 0x94, 0x93, 0x6f, 0xca,
		0xe0, 0xc7, 0xca, 0x06, 0x96, 0x83, 0xca, 0xb4, 0x43, 0x89, 0xc4, 0xc7, 0xe1, 0xbf, 0x2e, 0xbb,
		0xef, 0x9e, 0xbf, 0xef, 0xf2, 0x57, 0x95, 0xaa, 0x4d, 0x5b, 0x57, 0x70, 0x42, 0x14, 0x34, 0x33,
		0x9b, 0xc0, 0x97, 0x33, 0x38, 0xcc, 0x84, 0xb6, 0x41, 0xe0, 0x8f, 0x17, 0x42, 0x37, 0x9b, 0x53,
		0xec, 0x8f, 0x4f, 0x2a, 0x64, 0x12, 0xf8, 0x6e, 0x01, 0x37, 0x51, 0xa4, 0xc8, 0xba, 0x34, 0x5a,
		0xe2, 0xef, 0x7d, 0xde, 0x35, 0xaa, 0x88, 0xd3, 0x68, 0xa6, 0x0c, 0x82, 0x2b, 0x59, 0x4e, 0x24,
		0xf8, 0x23, 0x3f, 0x97, 0x10, 0x21, 0xc5, 0x2a, 0x97, 0xa9, 0x5c, 0x38, 0xb6, 0xd8, 0x35, 0x1d,
		0x8e, 0x1d, 0x82, 0xab, 0xcc, 0x74, 0x38, 0x16, 0x04, 0x57, 0x05, 0x2f, 0x1c, 0x03, 0xc1, 0x55,
		0xf1, 0x2e, 0xe6, 0x31, 0xd8, 0x8b, 0xa6, 0x9c, 0x4b, 0x5c, 0x0d, 0x1f, 0x3b, 0x21, 0xb1, 0x88,
		0x11, 0xf3, 0x01, 0x68, 0x46, 0xc4, 0xbc, 0xfe, 0x66, 0x66, 0x2a, 0x1c, 0x33, 0x43, 0xd6, 0x0d,
		0xc7, 0x40, 0xd0, 0x2c, 0x78, 0xd1, 0x72, 0x3f, 0x18, 0xbf, 0xfc, 0xcd, 0xe5, 0xae, 0x54, 0x31,
		0x08, 0xca, 0x24, 0xe8, 0x48, 0x2f, 0x19, 0xd4, 0x65, 0x18, 0x14, 0x1f, 0xa5, 0x0f, 0x51, 0x50,
		0x7c, 0x62, 0x3a, 0x51, 0x50, 0x7c, 0x5a, 0x3a, 0x99, 0xc9, 0x70, 0x2c, 0x08, 0x3a, 0x53, 0x91,
		0x0f, 0x10, 0x74, 0xa6, 0x67, 0x83, 0x3c, 0x02, 0xc1, 0xad, 0x9c, 0x4f, 0xe6, 0xf1, 0x36, 0x62,
		0x86, 0x18, 0x33, 0xdf, 0x89, 0xb7, 0x11, 0x33, 0xd8, 0x04, 0xb7, 0xd1, 0x72, 0xb9, 0x15, 0xdf,
		0x66, 0xf3, 0xe1, 0x98, 0xe9, 0x66, 0xe6, 0xb0, 0x88, 0xa6, 0x6d, 0x90, 0xfd, 0xc9, 0xf8, 0xaf,
		0x10, 0xee, 0xa4, 0xc6, 0xfd, 0x43, 0xbb, 0xe9, 0x87, 0x6f, 0x25, 0x44, 0x08, 0x8a, 0xe3, 0x93,
		0x3d, 0x8a, 0x45, 0xb4, 0xec, 0xe0, 0x7d, 0xf2, 0x59, 0xde, 0x49, 0x4e, 0x78, 0xfe, 0x6c, 0xff,
		0xec, 0x74, 0xf7, 0x30, 0x28, 0x42, 0x3b, 0x7c, 0x89, 0x7c, 0x0e, 0x8f, 0x80, 0x1d, 0xbe, 0x92,
		0x3e, 0x8f, 0x4f, 0xf0, 0x03, 0xcd, 0x0e, 0x5e, 0x22, 0xbf, 0xca, 0x6c, 0xa2, 0x90, 0x77, 0x8f,
		0x9e, 0x9f, 0x1c, 0xec, 0x1f, 0xc7, 0x39, 0x44, 0xe0, 0x38, 0xae, 0x90, 0x49, 0xf0, 0xeb, 0xe4,
		0x14, 0xd6, 0x02, 0x0e, 0x93, 0xe0, 0x37, 0x39, 0xed, 0x16, 0xf4, 0x8f, 0xf1, 0xfd, 0x93, 0xb3,
		0x01, 0x3c, 0x2c, 0xf8, 0x9b, 0x9c, 0x52, 0x88, 0xa7, 0xba, 0x79, 0x2c, 0xa3, 0xe9, 0x18, 0xe4,
		0xfc, 0x2e, 0x8c, 0x3f, 0x45, 0xf2, 0x2f, 0x17, 0x5e, 0xfe, 0x0f, 0xfe, 0xcd, 0x2f, 0x97, 0xaf,
		0x7d, 0xc4, 0x61, 0x04, 0xc7, 0x10, 0x64, 0xfd, 0x2e, 0xe4, 0x38, 0xaf, 0xc0, 0xe1, 0x14, 0x58,
		0x7f, 0x08, 0xb9, 0x99, 0xb8, 0xbd, 0x78, 0xd6, 0x33, 0xbf, 0xfb, 0xae, 0xfd, 0x06, 0xf9, 0xad,
		0xe3, 0x04, 0x59, 0xb0, 0xfe, 0x10, 0x23, 0x93, 0x0a, 0x9a, 0x0c, 0xa7, 0xe6, 0x15, 0x04, 0x86,
		0x8b, 0x0f, 0xf1, 0x5e, 0x40, 0x2a, 0xc8, 0xfa, 0x9f, 0x90, 0xcf, 0x12, 0x6b, 0x62, 0xd2, 0xe7,
		0x5c, 0x7c, 0x3a, 0xb3, 0x08, 0xbc, 0x47, 0xf2, 0x0a, 0x9a, 0x0c, 0xa7, 0x97, 0x14, 0x04, 0x86,
		0x1b, 0xdf, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x86, 0xfb, 0x92, 0x32, 0x88, 0x12, 0x00, 0x00,
	}
	sourceinfo.Register("desc_test1.proto", srcInfo)
}