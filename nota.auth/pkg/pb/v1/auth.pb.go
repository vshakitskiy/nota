// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/auth.proto

package auth_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{5}
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	mi := &file_v1_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ValidateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	mi := &file_v1_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_v1_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_v1_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_v1_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_v1_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_v1_auth_proto protoreflect.FileDescriptor

const file_v1_auth_proto_rawDesc = "" +
	"\n" +
	"\rv1/auth.proto\x12\aauth_v1\x1a\x1cgoogle/api/annotations.proto\"_\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"+\n" +
	"\x10RegisterResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"W\n" +
	"\rLoginResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\"2\n" +
	"\rLogoutRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"\x10\n" +
	"\x0eLogoutResponse\"9\n" +
	"\x14ValidateTokenRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"0\n" +
	"\x15ValidateTokenResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\":\n" +
	"\x13RefreshTokenRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"9\n" +
	"\x14RefreshTokenResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"3\n" +
	"\x0eGetUserRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"\\\n" +
	"\x0fGetUserResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email2\xe7\x03\n" +
	"\vAuthService\x12a\n" +
	"\bRegister\x12\x18.auth_v1.RegisterRequest\x1a\x19.auth_v1.RegisterResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/api/v1/auth/register\x12U\n" +
	"\x05Login\x12\x15.auth_v1.LoginRequest\x1a\x16.auth_v1.LoginResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/api/v1/auth/login\x12Y\n" +
	"\x06Logout\x12\x16.auth_v1.LogoutRequest\x1a\x17.auth_v1.LogoutResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/api/v1/auth/logout\x12l\n" +
	"\fRefreshToken\x12\x1c.auth_v1.RefreshTokenRequest\x1a\x1d.auth_v1.RefreshTokenResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/api/v1/auth/refresh\x12U\n" +
	"\aGetUser\x12\x17.auth_v1.GetUserRequest\x1a\x18.auth_v1.GetUserResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/api/v1/auth/me2~\n" +
	"\rAccessService\x12m\n" +
	"\rValidateToken\x12\x1d.auth_v1.ValidateTokenRequest\x1a\x1e.auth_v1.ValidateTokenResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/api/v1/auth/checkB\x1dZ\x1bnota/pkg/pb/auth/v1;auth_v1b\x06proto3"

var (
	file_v1_auth_proto_rawDescOnce sync.Once
	file_v1_auth_proto_rawDescData []byte
)

func file_v1_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)))
	})
	return file_v1_auth_proto_rawDescData
}

var file_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_auth_proto_goTypes = []any{
	(*RegisterRequest)(nil),       // 0: auth_v1.RegisterRequest
	(*RegisterResponse)(nil),      // 1: auth_v1.RegisterResponse
	(*LoginRequest)(nil),          // 2: auth_v1.LoginRequest
	(*LoginResponse)(nil),         // 3: auth_v1.LoginResponse
	(*LogoutRequest)(nil),         // 4: auth_v1.LogoutRequest
	(*LogoutResponse)(nil),        // 5: auth_v1.LogoutResponse
	(*ValidateTokenRequest)(nil),  // 6: auth_v1.ValidateTokenRequest
	(*ValidateTokenResponse)(nil), // 7: auth_v1.ValidateTokenResponse
	(*RefreshTokenRequest)(nil),   // 8: auth_v1.RefreshTokenRequest
	(*RefreshTokenResponse)(nil),  // 9: auth_v1.RefreshTokenResponse
	(*GetUserRequest)(nil),        // 10: auth_v1.GetUserRequest
	(*GetUserResponse)(nil),       // 11: auth_v1.GetUserResponse
}
var file_v1_auth_proto_depIdxs = []int32{
	0,  // 0: auth_v1.AuthService.Register:input_type -> auth_v1.RegisterRequest
	2,  // 1: auth_v1.AuthService.Login:input_type -> auth_v1.LoginRequest
	4,  // 2: auth_v1.AuthService.Logout:input_type -> auth_v1.LogoutRequest
	8,  // 3: auth_v1.AuthService.RefreshToken:input_type -> auth_v1.RefreshTokenRequest
	10, // 4: auth_v1.AuthService.GetUser:input_type -> auth_v1.GetUserRequest
	6,  // 5: auth_v1.AccessService.ValidateToken:input_type -> auth_v1.ValidateTokenRequest
	1,  // 6: auth_v1.AuthService.Register:output_type -> auth_v1.RegisterResponse
	3,  // 7: auth_v1.AuthService.Login:output_type -> auth_v1.LoginResponse
	5,  // 8: auth_v1.AuthService.Logout:output_type -> auth_v1.LogoutResponse
	9,  // 9: auth_v1.AuthService.RefreshToken:output_type -> auth_v1.RefreshTokenResponse
	11, // 10: auth_v1.AuthService.GetUser:output_type -> auth_v1.GetUserResponse
	7,  // 11: auth_v1.AccessService.ValidateToken:output_type -> auth_v1.ValidateTokenResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_auth_proto_init() }
func file_v1_auth_proto_init() {
	if File_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_v1_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_proto_depIdxs,
		MessageInfos:      file_v1_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_proto = out.File
	file_v1_auth_proto_goTypes = nil
	file_v1_auth_proto_depIdxs = nil
}
