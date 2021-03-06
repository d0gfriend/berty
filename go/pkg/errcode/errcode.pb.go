// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                                  ErrCode = 0
	TODO                                       ErrCode = 666
	ErrNotImplemented                          ErrCode = 777
	ErrInternal                                ErrCode = 888
	ErrInvalidInput                            ErrCode = 100
	ErrMissingInput                            ErrCode = 101
	ErrSerialization                           ErrCode = 102
	ErrDeserialization                         ErrCode = 103
	ErrStreamRead                              ErrCode = 104
	ErrStreamWrite                             ErrCode = 105
	ErrMissingMapKey                           ErrCode = 106
	ErrCryptoRandomGeneration                  ErrCode = 200
	ErrCryptoKeyGeneration                     ErrCode = 201
	ErrCryptoNonceGeneration                   ErrCode = 202
	ErrCryptoSignature                         ErrCode = 203
	ErrCryptoSignatureVerification             ErrCode = 204
	ErrCryptoDecrypt                           ErrCode = 205
	ErrCryptoEncrypt                           ErrCode = 206
	ErrCryptoKeyConversion                     ErrCode = 207
	ErrOrbitDBInit                             ErrCode = 1000
	ErrOrbitDBOpen                             ErrCode = 1001
	ErrOrbitDBAppend                           ErrCode = 1002
	ErrOrbitDBDeserialization                  ErrCode = 1003
	ErrOrbitDBStoreCast                        ErrCode = 1004
	ErrHandshakeOwnEphemeralKeyGenSend         ErrCode = 1100
	ErrHandshakePeerEphemeralKeyRecv           ErrCode = 1101
	ErrHandshakeRequesterAuthenticateBoxKeyGen ErrCode = 1102
	ErrHandshakeResponderAcceptBoxKeyGen       ErrCode = 1103
	ErrHandshakeRequesterHello                 ErrCode = 1104
	ErrHandshakeResponderHello                 ErrCode = 1105
	ErrHandshakeRequesterAuthenticate          ErrCode = 1106
	ErrHandshakeResponderAccept                ErrCode = 1107
	ErrHandshakeRequesterAcknowledge           ErrCode = 1108
	ErrGroupMemberLogEventOpen                 ErrCode = 1200
	ErrGroupMemberLogEventSignature            ErrCode = 1201
	ErrGroupMemberUnknownGroupID               ErrCode = 1202
	ErrGroupSecretOtherDestMember              ErrCode = 1203
	ErrGroupSecretAlreadySentToMember          ErrCode = 1204
	ErrGroupInvalidType                        ErrCode = 1205
	ErrGroupMissing                            ErrCode = 1206
	ErrMessageKeyPersistencePut                ErrCode = 1300
	ErrMessageKeyPersistenceGet                ErrCode = 1301
	ErrBridgeInterrupted                       ErrCode = 1400
	ErrBridgeNotRunning                        ErrCode = 1401
	ErrMessengerInvalidDeepLink                ErrCode = 2001
	ErrCLINoTermcaps                           ErrCode = 3001
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrMissingInput",
	102:  "ErrSerialization",
	103:  "ErrDeserialization",
	104:  "ErrStreamRead",
	105:  "ErrStreamWrite",
	106:  "ErrMissingMapKey",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoEncrypt",
	207:  "ErrCryptoKeyConversion",
	1000: "ErrOrbitDBInit",
	1001: "ErrOrbitDBOpen",
	1002: "ErrOrbitDBAppend",
	1003: "ErrOrbitDBDeserialization",
	1004: "ErrOrbitDBStoreCast",
	1100: "ErrHandshakeOwnEphemeralKeyGenSend",
	1101: "ErrHandshakePeerEphemeralKeyRecv",
	1102: "ErrHandshakeRequesterAuthenticateBoxKeyGen",
	1103: "ErrHandshakeResponderAcceptBoxKeyGen",
	1104: "ErrHandshakeRequesterHello",
	1105: "ErrHandshakeResponderHello",
	1106: "ErrHandshakeRequesterAuthenticate",
	1107: "ErrHandshakeResponderAccept",
	1108: "ErrHandshakeRequesterAcknowledge",
	1200: "ErrGroupMemberLogEventOpen",
	1201: "ErrGroupMemberLogEventSignature",
	1202: "ErrGroupMemberUnknownGroupID",
	1203: "ErrGroupSecretOtherDestMember",
	1204: "ErrGroupSecretAlreadySentToMember",
	1205: "ErrGroupInvalidType",
	1206: "ErrGroupMissing",
	1300: "ErrMessageKeyPersistencePut",
	1301: "ErrMessageKeyPersistenceGet",
	1400: "ErrBridgeInterrupted",
	1401: "ErrBridgeNotRunning",
	2001: "ErrMessengerInvalidDeepLink",
	3001: "ErrCLINoTermcaps",
}

var ErrCode_value = map[string]int32{
	"Undefined":                                  0,
	"TODO":                                       666,
	"ErrNotImplemented":                          777,
	"ErrInternal":                                888,
	"ErrInvalidInput":                            100,
	"ErrMissingInput":                            101,
	"ErrSerialization":                           102,
	"ErrDeserialization":                         103,
	"ErrStreamRead":                              104,
	"ErrStreamWrite":                             105,
	"ErrMissingMapKey":                           106,
	"ErrCryptoRandomGeneration":                  200,
	"ErrCryptoKeyGeneration":                     201,
	"ErrCryptoNonceGeneration":                   202,
	"ErrCryptoSignature":                         203,
	"ErrCryptoSignatureVerification":             204,
	"ErrCryptoDecrypt":                           205,
	"ErrCryptoEncrypt":                           206,
	"ErrCryptoKeyConversion":                     207,
	"ErrOrbitDBInit":                             1000,
	"ErrOrbitDBOpen":                             1001,
	"ErrOrbitDBAppend":                           1002,
	"ErrOrbitDBDeserialization":                  1003,
	"ErrOrbitDBStoreCast":                        1004,
	"ErrHandshakeOwnEphemeralKeyGenSend":         1100,
	"ErrHandshakePeerEphemeralKeyRecv":           1101,
	"ErrHandshakeRequesterAuthenticateBoxKeyGen": 1102,
	"ErrHandshakeResponderAcceptBoxKeyGen":       1103,
	"ErrHandshakeRequesterHello":                 1104,
	"ErrHandshakeResponderHello":                 1105,
	"ErrHandshakeRequesterAuthenticate":          1106,
	"ErrHandshakeResponderAccept":                1107,
	"ErrHandshakeRequesterAcknowledge":           1108,
	"ErrGroupMemberLogEventOpen":                 1200,
	"ErrGroupMemberLogEventSignature":            1201,
	"ErrGroupMemberUnknownGroupID":               1202,
	"ErrGroupSecretOtherDestMember":              1203,
	"ErrGroupSecretAlreadySentToMember":          1204,
	"ErrGroupInvalidType":                        1205,
	"ErrGroupMissing":                            1206,
	"ErrMessageKeyPersistencePut":                1300,
	"ErrMessageKeyPersistenceGet":                1301,
	"ErrBridgeInterrupted":                       1400,
	"ErrBridgeNotRunning":                        1401,
	"ErrMessengerInvalidDeepLink":                2001,
	"ErrCLINoTermcaps":                           3001,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func init() {
	proto.RegisterEnum("berty.errcode.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcb, 0x72, 0xdc, 0x44,
	0x14, 0xcd, 0x54, 0x81, 0x67, 0xd4, 0x94, 0x49, 0xe7, 0xda, 0x38, 0x2f, 0x62, 0x27, 0x21, 0x81,
	0x22, 0x55, 0x78, 0x16, 0x7c, 0x81, 0xed, 0x51, 0x39, 0x53, 0x7e, 0x8c, 0x6b, 0xc6, 0x81, 0x2a,
	0x76, 0x1a, 0xe9, 0x58, 0xd3, 0x58, 0xea, 0x16, 0x57, 0x2d, 0x87, 0xe1, 0x0f, 0xd8, 0xc3, 0x86,
	0xaf, 0xe0, 0x5d, 0x95, 0x3f, 0xe0, 0x91, 0x27, 0xf0, 0x03, 0xec, 0x78, 0x7d, 0x40, 0xd8, 0x51,
	0x52, 0xcb, 0xf6, 0x0c, 0x76, 0x39, 0x2b, 0x49, 0xe7, 0x9e, 0x3e, 0xf7, 0xe8, 0x76, 0xf7, 0x11,
	0xb3, 0x60, 0x0e, 0x4d, 0x84, 0xe5, 0x8c, 0x8d, 0x35, 0x34, 0x3b, 0x04, 0xdb, 0xf1, 0x72, 0x0d,
	0x5e, 0x79, 0x27, 0x56, 0x76, 0x54, 0x0c, 0x97, 0x43, 0x93, 0xb6, 0x63, 0x13, 0x9b, 0x76, 0xc5,
	0x1a, 0x16, 0x7b, 0xd5, 0x57, 0xf5, 0x51, 0xbd, 0xb9, 0xd5, 0x77, 0x1e, 0x08, 0xd1, 0xf4, 0x99,
	0xd7, 0x4c, 0x04, 0x9a, 0x15, 0xde, 0x3d, 0x1d, 0x61, 0x4f, 0x69, 0x44, 0xf2, 0x1c, 0x79, 0xe2,
	0xa5, 0xdd, 0x5e, 0xa7, 0x27, 0xbf, 0x78, 0x99, 0x16, 0xc4, 0x05, 0x9f, 0x79, 0xdb, 0xd8, 0x6e,
	0x9a, 0x25, 0x48, 0xa1, 0x2d, 0x22, 0xf9, 0xe9, 0x0c, 0x49, 0xf1, 0x8a, 0xcf, 0xdc, 0xd5, 0x16,
	0xac, 0x83, 0x44, 0x3e, 0x9f, 0xa1, 0x39, 0x71, 0xbe, 0x42, 0x0e, 0x82, 0x44, 0x45, 0x5d, 0x9d,
	0x15, 0x56, 0x46, 0x35, 0xb8, 0xa5, 0xf2, 0x5c, 0xe9, 0xd8, 0x81, 0xa0, 0x79, 0x21, 0x7d, 0xe6,
	0x01, 0x58, 0x05, 0x89, 0xfa, 0x24, 0xb0, 0xca, 0x68, 0xb9, 0x47, 0x0b, 0x82, 0x7c, 0xe6, 0x0e,
	0xf2, 0x29, 0x3c, 0xa6, 0x0b, 0x62, 0xb6, 0x64, 0x5b, 0x46, 0x90, 0xf6, 0x11, 0x44, 0x72, 0x44,
	0x24, 0x5e, 0x3d, 0x82, 0xde, 0x67, 0x65, 0x21, 0x55, 0x2d, 0x5a, 0x77, 0xda, 0x0a, 0xb2, 0x0d,
	0x8c, 0xe5, 0x87, 0xb4, 0x28, 0x2e, 0x97, 0xff, 0xc8, 0xe3, 0xcc, 0x9a, 0x7e, 0xa0, 0x23, 0x93,
	0xae, 0x43, 0x83, 0x9d, 0xf6, 0x0f, 0x0d, 0xba, 0x2a, 0x16, 0x8e, 0xea, 0x1b, 0x18, 0x4f, 0x14,
	0x7f, 0x6c, 0xd0, 0x35, 0x71, 0xe9, 0xa8, 0xb8, 0x6d, 0x74, 0x88, 0x89, 0xf2, 0x4f, 0x0d, 0xba,
	0x58, 0x19, 0x76, 0xe5, 0x81, 0x8a, 0x75, 0x60, 0x0b, 0x86, 0xfc, 0xb9, 0x41, 0x6f, 0x88, 0xc5,
	0x93, 0x85, 0xf7, 0xc0, 0x6a, 0x4f, 0x85, 0x6e, 0xf5, 0xc3, 0x06, 0xbd, 0x56, 0xf9, 0x75, 0xa4,
	0x0e, 0xc2, 0xf2, 0x29, 0x1f, 0x4d, 0xc3, 0xbe, 0x76, 0xf0, 0xe3, 0x13, 0x3e, 0xd7, 0x8c, 0x3e,
	0x00, 0xe7, 0xa5, 0xd4, 0x93, 0x06, 0xcd, 0x55, 0xe3, 0xe8, 0xf1, 0x50, 0xd9, 0xce, 0x6a, 0x57,
	0x2b, 0x2b, 0xff, 0x68, 0x4e, 0x83, 0xbd, 0x0c, 0x5a, 0xfe, 0xd9, 0xac, 0xd5, 0x6b, 0x70, 0x25,
	0xcb, 0xa0, 0x23, 0xf9, 0x57, 0xb3, 0x9e, 0x52, 0x0d, 0xff, 0x7f, 0x07, 0xfe, 0x6e, 0xd2, 0x25,
	0x31, 0x77, 0x5c, 0x1f, 0x58, 0xc3, 0x58, 0x0b, 0x72, 0x2b, 0xff, 0x69, 0xd2, 0x5b, 0xe2, 0xa6,
	0xcf, 0x7c, 0x37, 0xd0, 0x51, 0x3e, 0x0a, 0xf6, 0xd1, 0xbb, 0xaf, 0xfd, 0x6c, 0x84, 0x14, 0x1c,
	0x24, 0x6e, 0x9c, 0x83, 0xb2, 0xc5, 0xc3, 0x16, 0xdd, 0x16, 0xd7, 0x27, 0x89, 0x3b, 0x00, 0x4f,
	0x32, 0xfb, 0x08, 0x0f, 0xe4, 0xa3, 0x16, 0xb5, 0xc5, 0x9d, 0x49, 0x5a, 0x1f, 0x1f, 0x15, 0xc8,
	0x2d, 0x78, 0xa5, 0xb0, 0x23, 0x68, 0x5b, 0xce, 0x0f, 0xab, 0xe6, 0x63, 0xa7, 0x2d, 0x1f, 0xb7,
	0xe8, 0x6d, 0x71, 0x6b, 0x7a, 0x41, 0x9e, 0x19, 0x1d, 0x81, 0x57, 0xc2, 0x10, 0x99, 0x3d, 0xa6,
	0x3e, 0x69, 0xd1, 0x92, 0xb8, 0x72, 0xaa, 0xf6, 0x5d, 0x24, 0x89, 0x91, 0x4f, 0x4f, 0x21, 0xd4,
	0x5a, 0x8e, 0xf0, 0xac, 0x45, 0x6f, 0x8a, 0x1b, 0x2f, 0x74, 0x27, 0x7f, 0x69, 0xd1, 0x75, 0x71,
	0xf5, 0x0c, 0x53, 0xf2, 0xd7, 0x13, 0xe3, 0x38, 0x56, 0x0a, 0xf7, 0xb5, 0xb9, 0x9f, 0x20, 0x8a,
	0x21, 0x7f, 0x3b, 0x74, 0xb4, 0xce, 0xa6, 0xc8, 0xb6, 0x90, 0x0e, 0xc1, 0x9b, 0x26, 0xf6, 0x0f,
	0xa0, 0x6d, 0xb5, 0xa1, 0x5f, 0x7a, 0x74, 0x4b, 0x2c, 0x9d, 0x4e, 0x38, 0x3e, 0x90, 0x5f, 0x79,
	0x74, 0x43, 0xbc, 0x3e, 0xcd, 0xba, 0xa7, 0xcb, 0x36, 0xba, 0x42, 0xba, 0x1d, 0xf9, 0xb5, 0x47,
	0x37, 0xc5, 0xb5, 0x43, 0xca, 0x00, 0x21, 0xc3, 0xf6, 0xec, 0x08, 0xe5, 0x6d, 0xb4, 0x6e, 0x85,
	0xfc, 0xc6, 0xab, 0x7f, 0x7f, 0x82, 0xb3, 0x92, 0x30, 0x82, 0x68, 0x3c, 0x80, 0xb6, 0xbb, 0xa6,
	0xe6, 0x7d, 0xeb, 0xd5, 0xc7, 0xc5, 0x89, 0xbb, 0x38, 0xd8, 0x1d, 0x67, 0x90, 0xdf, 0x79, 0x34,
	0x5f, 0xc5, 0x81, 0x33, 0xe2, 0x6e, 0xaa, 0xfc, 0xde, 0xab, 0xc7, 0xb5, 0x85, 0x3c, 0x0f, 0x62,
	0x6c, 0x60, 0xbc, 0x53, 0x1e, 0xed, 0xdc, 0x42, 0x87, 0xd8, 0x29, 0xac, 0xfc, 0x4c, 0x9c, 0xc5,
	0x58, 0x87, 0x95, 0x9f, 0x0b, 0xba, 0x2c, 0xe6, 0x7d, 0xe6, 0x55, 0x56, 0x51, 0x8c, 0x2a, 0x95,
	0xb8, 0xc8, 0xca, 0xa8, 0x7a, 0x2e, 0x6a, 0x3b, 0xae, 0xb4, 0x6d, 0x6c, 0xbf, 0xd0, 0xba, 0x6c,
	0xfc, 0xef, 0xa4, 0x2c, 0x74, 0x8c, 0xc3, 0xec, 0xea, 0x00, 0xd9, 0xa6, 0xd2, 0xfb, 0xf2, 0xd9,
	0xf9, 0xc3, 0xeb, 0xb8, 0xd9, 0xdd, 0x36, 0xbb, 0xe0, 0x34, 0x0c, 0xb2, 0x5c, 0x3e, 0xb8, 0xb8,
	0x7a, 0xfb, 0xe9, 0xef, 0x8b, 0xe7, 0x3e, 0x58, 0x72, 0xf9, 0x6b, 0x11, 0x8e, 0xda, 0xd5, 0x6b,
	0xbb, 0x0c, 0xdd, 0xfd, 0xb8, 0x5d, 0x27, 0xf2, 0x70, 0xa6, 0x4a, 0xda, 0x77, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xbe, 0x73, 0x56, 0xb8, 0x05, 0x00, 0x00,
}
