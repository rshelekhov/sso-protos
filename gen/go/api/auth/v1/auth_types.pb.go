// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/auth/v1/auth_types.proto

package authv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Platform options for user devices
type Platform int32

const (
	// Default unspecified value
	Platform_PLATFORM_UNSPECIFIED Platform = 0
	// Web platform
	Platform_PLATFORM_WEB Platform = 1
	// iOS platform
	Platform_PLATFORM_IOS Platform = 2
	// Android platform
	Platform_PLATFORM_ANDROID Platform = 3
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "PLATFORM_WEB",
		2: "PLATFORM_IOS",
		3: "PLATFORM_ANDROID",
	}
	Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"PLATFORM_WEB":         1,
		"PLATFORM_IOS":         2,
		"PLATFORM_ANDROID":     3,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v1_auth_types_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_api_auth_v1_auth_types_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_types_proto_rawDescGZIP(), []int{0}
}

// Device fingerprint for security and session management
type UserDeviceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User agent string from the device
	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// User's IP address
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// Platform type (Web, iOS, Android)
	Platform Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=api.auth.v1.Platform" json:"platform,omitempty"`
	// Only one of app_version or browser_version should be set based on the platform
	//
	// Types that are assignable to Version:
	//
	//	*UserDeviceData_AppVersion
	//	*UserDeviceData_BrowserVersion
	Version isUserDeviceData_Version `protobuf_oneof:"version"`
}

func (x *UserDeviceData) Reset() {
	*x = UserDeviceData{}
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDeviceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeviceData) ProtoMessage() {}

func (x *UserDeviceData) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeviceData.ProtoReflect.Descriptor instead.
func (*UserDeviceData) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_types_proto_rawDescGZIP(), []int{0}
}

func (x *UserDeviceData) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *UserDeviceData) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UserDeviceData) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_PLATFORM_UNSPECIFIED
}

func (m *UserDeviceData) GetVersion() isUserDeviceData_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *UserDeviceData) GetAppVersion() string {
	if x, ok := x.GetVersion().(*UserDeviceData_AppVersion); ok {
		return x.AppVersion
	}
	return ""
}

func (x *UserDeviceData) GetBrowserVersion() string {
	if x, ok := x.GetVersion().(*UserDeviceData_BrowserVersion); ok {
		return x.BrowserVersion
	}
	return ""
}

type isUserDeviceData_Version interface {
	isUserDeviceData_Version()
}

type UserDeviceData_AppVersion struct {
	// Mobile app version (for iOS/Android)
	AppVersion string `protobuf:"bytes,4,opt,name=app_version,json=appVersion,proto3,oneof"`
}

type UserDeviceData_BrowserVersion struct {
	// Browser version (for web platform)
	BrowserVersion string `protobuf:"bytes,5,opt,name=browser_version,json=browserVersion,proto3,oneof"`
}

func (*UserDeviceData_AppVersion) isUserDeviceData_Version() {}

func (*UserDeviceData_BrowserVersion) isUserDeviceData_Version() {}

// Token information for authentication
type TokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JWT access token for API authentication
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Refresh token for obtaining new access tokens
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// Cookie domain for web clients
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	// Cookie path for web clients
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// When access token expires
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Whether cookies should be HTTP-only
	HttpOnly bool `protobuf:"varint,6,opt,name=http_only,json=httpOnly,proto3" json:"http_only,omitempty"`
	// Extra fields for application-specific data
	AdditionalFields map[string]string `protobuf:"bytes,7,rep,name=additional_fields,json=additionalFields,proto3" json:"additional_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TokenData) Reset() {
	*x = TokenData{}
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenData) ProtoMessage() {}

func (x *TokenData) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenData.ProtoReflect.Descriptor instead.
func (*TokenData) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_types_proto_rawDescGZIP(), []int{1}
}

func (x *TokenData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenData) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *TokenData) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *TokenData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TokenData) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *TokenData) GetHttpOnly() bool {
	if x != nil {
		return x.HttpOnly
	}
	return false
}

func (x *TokenData) GetAdditionalFields() map[string]string {
	if x != nil {
		return x.AdditionalFields
	}
	return nil
}

// JSON Web Key for JWT signature verification
type JWK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptographic algorithm (e.g., "RS256")
	Alg string `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	// Key type (e.g., "RSA")
	Kty string `protobuf:"bytes,2,opt,name=kty,proto3" json:"kty,omitempty"`
	// Key usage (e.g., "sig" for signature)
	Use string `protobuf:"bytes,3,opt,name=use,proto3" json:"use,omitempty"`
	// Key identifier for key rotation
	Kid string `protobuf:"bytes,4,opt,name=kid,proto3" json:"kid,omitempty"`
	// RSA modulus (base64url encoded)
	N string `protobuf:"bytes,5,opt,name=n,proto3" json:"n,omitempty"`
	// RSA exponent (base64url encoded)
	E string `protobuf:"bytes,6,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *JWK) Reset() {
	*x = JWK{}
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWK) ProtoMessage() {}

func (x *JWK) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWK.ProtoReflect.Descriptor instead.
func (*JWK) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_types_proto_rawDescGZIP(), []int{2}
}

func (x *JWK) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *JWK) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *JWK) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

func (x *JWK) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JWK) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *JWK) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

var File_api_auth_v1_auth_types_proto protoreflect.FileDescriptor

var file_api_auth_v1_auth_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x72, 0x02, 0x70, 0x01, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x3e, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x21, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x62,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22,
	0xf7, 0x02, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x68, 0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x11, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x43, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x69, 0x0a, 0x03, 0x4a, 0x57, 0x4b,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x6c, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x65, 0x2a, 0x5e, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f,
	0x49, 0x44, 0x10, 0x03, 0x42, 0x92, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x72, 0x73, 0x68, 0x65,
	0x6c, 0x65, 0x6b, 0x68, 0x6f, 0x76, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x41, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_auth_v1_auth_types_proto_rawDescOnce sync.Once
	file_api_auth_v1_auth_types_proto_rawDescData = file_api_auth_v1_auth_types_proto_rawDesc
)

func file_api_auth_v1_auth_types_proto_rawDescGZIP() []byte {
	file_api_auth_v1_auth_types_proto_rawDescOnce.Do(func() {
		file_api_auth_v1_auth_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v1_auth_types_proto_rawDescData)
	})
	return file_api_auth_v1_auth_types_proto_rawDescData
}

var file_api_auth_v1_auth_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_auth_v1_auth_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_auth_v1_auth_types_proto_goTypes = []any{
	(Platform)(0),                 // 0: api.auth.v1.Platform
	(*UserDeviceData)(nil),        // 1: api.auth.v1.UserDeviceData
	(*TokenData)(nil),             // 2: api.auth.v1.TokenData
	(*JWK)(nil),                   // 3: api.auth.v1.JWK
	nil,                           // 4: api.auth.v1.TokenData.AdditionalFieldsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_auth_v1_auth_types_proto_depIdxs = []int32{
	0, // 0: api.auth.v1.UserDeviceData.platform:type_name -> api.auth.v1.Platform
	5, // 1: api.auth.v1.TokenData.expires_at:type_name -> google.protobuf.Timestamp
	4, // 2: api.auth.v1.TokenData.additional_fields:type_name -> api.auth.v1.TokenData.AdditionalFieldsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_auth_v1_auth_types_proto_init() }
func file_api_auth_v1_auth_types_proto_init() {
	if File_api_auth_v1_auth_types_proto != nil {
		return
	}
	file_api_auth_v1_auth_types_proto_msgTypes[0].OneofWrappers = []any{
		(*UserDeviceData_AppVersion)(nil),
		(*UserDeviceData_BrowserVersion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_auth_v1_auth_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_v1_auth_types_proto_goTypes,
		DependencyIndexes: file_api_auth_v1_auth_types_proto_depIdxs,
		EnumInfos:         file_api_auth_v1_auth_types_proto_enumTypes,
		MessageInfos:      file_api_auth_v1_auth_types_proto_msgTypes,
	}.Build()
	File_api_auth_v1_auth_types_proto = out.File
	file_api_auth_v1_auth_types_proto_rawDesc = nil
	file_api_auth_v1_auth_types_proto_goTypes = nil
	file_api_auth_v1_auth_types_proto_depIdxs = nil
}
