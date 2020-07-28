// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: user.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserFriendsFriendOfFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	InviteStatus string `protobuf:"bytes,2,opt,name=invite_status,json=inviteStatus,proto3" json:"invite_status,omitempty"`
}

func (x *UserFriendsFriendOfFriend) Reset() {
	*x = UserFriendsFriendOfFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsFriendOfFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsFriendOfFriend) ProtoMessage() {}

func (x *UserFriendsFriendOfFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsFriendOfFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsFriendOfFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserFriendsFriendOfFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsFriendOfFriend) GetInviteStatus() string {
	if x != nil {
		return x.InviteStatus
	}
	return ""
}

type UserFriendsFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User                        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Friends []*UserFriendsFriendOfFriend `protobuf:"bytes,2,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsFriend) Reset() {
	*x = UserFriendsFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsFriend) ProtoMessage() {}

func (x *UserFriendsFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserFriendsFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsFriend) GetFriends() []*UserFriendsFriendOfFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*UserFriendsFriend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsResponse) Reset() {
	*x = UserFriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsResponse) ProtoMessage() {}

func (x *UserFriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsResponse.ProtoReflect.Descriptor instead.
func (*UserFriendsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserFriendsResponse) GetFriends() []*UserFriendsFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsOfFriendsResponseFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	InviteStatus string  `protobuf:"bytes,2,opt,name=invite_status,json=inviteStatus,proto3" json:"invite_status,omitempty"`
	Friends      []*User `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsOfFriendsResponseFriend) Reset() {
	*x = UserFriendsOfFriendsResponseFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsOfFriendsResponseFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsOfFriendsResponseFriend) ProtoMessage() {}

func (x *UserFriendsOfFriendsResponseFriend) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsOfFriendsResponseFriend.ProtoReflect.Descriptor instead.
func (*UserFriendsOfFriendsResponseFriend) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserFriendsOfFriendsResponseFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserFriendsOfFriendsResponseFriend) GetInviteStatus() string {
	if x != nil {
		return x.InviteStatus
	}
	return ""
}

func (x *UserFriendsOfFriendsResponseFriend) GetFriends() []*User {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserFriendsOfFriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*UserFriendsOfFriendsResponseFriend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *UserFriendsOfFriendsResponse) Reset() {
	*x = UserFriendsOfFriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFriendsOfFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFriendsOfFriendsResponse) ProtoMessage() {}

func (x *UserFriendsOfFriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFriendsOfFriendsResponse.ProtoReflect.Descriptor instead.
func (*UserFriendsOfFriendsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserFriendsOfFriendsResponse) GetFriends() []*UserFriendsOfFriendsResponseFriend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type UserMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	IsAdmin bool  `protobuf:"varint,2,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *UserMeResponse) Reset() {
	*x = UserMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeResponse) ProtoMessage() {}

func (x *UserMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeResponse.ProtoReflect.Descriptor instead.
func (*UserMeResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserMeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserMeResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type UserEditProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailChanged    bool   `protobuf:"varint,1,opt,name=email_changed,json=emailChanged,proto3" json:"email_changed,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChanged bool   `protobuf:"varint,3,opt,name=password_changed,json=passwordChanged,proto3" json:"password_changed,omitempty"`
	Password        string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	AvatarChanged   bool   `protobuf:"varint,5,opt,name=avatar_changed,json=avatarChanged,proto3" json:"avatar_changed,omitempty"`
	AvatarId        string `protobuf:"bytes,6,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *UserEditProfileRequest) Reset() {
	*x = UserEditProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEditProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEditProfileRequest) ProtoMessage() {}

func (x *UserEditProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEditProfileRequest.ProtoReflect.Descriptor instead.
func (*UserEditProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserEditProfileRequest) GetEmailChanged() bool {
	if x != nil {
		return x.EmailChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserEditProfileRequest) GetPasswordChanged() bool {
	if x != nil {
		return x.PasswordChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserEditProfileRequest) GetAvatarChanged() bool {
	if x != nil {
		return x.AvatarChanged
	}
	return false
}

func (x *UserEditProfileRequest) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

type UserUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *UserUsersRequest) Reset() {
	*x = UserUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUsersRequest) ProtoMessage() {}

func (x *UserUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUsersRequest.ProtoReflect.Descriptor instead.
func (*UserUsersRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type UserUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserUsersResponse) Reset() {
	*x = UserUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUsersResponse) ProtoMessage() {}

func (x *UserUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUsersResponse.ProtoReflect.Descriptor instead.
func (*UserUsersResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x6c, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x47, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x22, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x45, 0x64,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x98, 0x02, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x66,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x02, 0x4d, 0x65, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []interface{}{
	(*UserFriendsFriendOfFriend)(nil),          // 0: rpc.UserFriendsFriendOfFriend
	(*UserFriendsFriend)(nil),                  // 1: rpc.UserFriendsFriend
	(*UserFriendsResponse)(nil),                // 2: rpc.UserFriendsResponse
	(*UserFriendsOfFriendsResponseFriend)(nil), // 3: rpc.UserFriendsOfFriendsResponseFriend
	(*UserFriendsOfFriendsResponse)(nil),       // 4: rpc.UserFriendsOfFriendsResponse
	(*UserMeResponse)(nil),                     // 5: rpc.UserMeResponse
	(*UserEditProfileRequest)(nil),             // 6: rpc.UserEditProfileRequest
	(*UserUsersRequest)(nil),                   // 7: rpc.UserUsersRequest
	(*UserUsersResponse)(nil),                  // 8: rpc.UserUsersResponse
	(*User)(nil),                               // 9: rpc.User
	(*Empty)(nil),                              // 10: rpc.Empty
}
var file_user_proto_depIdxs = []int32{
	9,  // 0: rpc.UserFriendsFriendOfFriend.user:type_name -> rpc.User
	9,  // 1: rpc.UserFriendsFriend.user:type_name -> rpc.User
	0,  // 2: rpc.UserFriendsFriend.friends:type_name -> rpc.UserFriendsFriendOfFriend
	1,  // 3: rpc.UserFriendsResponse.friends:type_name -> rpc.UserFriendsFriend
	9,  // 4: rpc.UserFriendsOfFriendsResponseFriend.user:type_name -> rpc.User
	9,  // 5: rpc.UserFriendsOfFriendsResponseFriend.friends:type_name -> rpc.User
	3,  // 6: rpc.UserFriendsOfFriendsResponse.friends:type_name -> rpc.UserFriendsOfFriendsResponseFriend
	9,  // 7: rpc.UserMeResponse.user:type_name -> rpc.User
	9,  // 8: rpc.UserUsersResponse.users:type_name -> rpc.User
	10, // 9: rpc.UserService.Friends:input_type -> rpc.Empty
	10, // 10: rpc.UserService.FriendsOfFriends:input_type -> rpc.Empty
	10, // 11: rpc.UserService.Me:input_type -> rpc.Empty
	6,  // 12: rpc.UserService.EditProfile:input_type -> rpc.UserEditProfileRequest
	7,  // 13: rpc.UserService.Users:input_type -> rpc.UserUsersRequest
	2,  // 14: rpc.UserService.Friends:output_type -> rpc.UserFriendsResponse
	4,  // 15: rpc.UserService.FriendsOfFriends:output_type -> rpc.UserFriendsOfFriendsResponse
	5,  // 16: rpc.UserService.Me:output_type -> rpc.UserMeResponse
	10, // 17: rpc.UserService.EditProfile:output_type -> rpc.Empty
	8,  // 18: rpc.UserService.Users:output_type -> rpc.UserUsersResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsFriendOfFriend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsFriend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsOfFriendsResponseFriend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFriendsOfFriendsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEditProfileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
