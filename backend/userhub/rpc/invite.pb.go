// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: invite.proto

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

type InviteCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friend string `protobuf:"bytes,1,opt,name=friend,proto3" json:"friend,omitempty"`
}

func (x *InviteCreateRequest) Reset() {
	*x = InviteCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteCreateRequest) ProtoMessage() {}

func (x *InviteCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteCreateRequest.ProtoReflect.Descriptor instead.
func (*InviteCreateRequest) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{0}
}

func (x *InviteCreateRequest) GetFriend() string {
	if x != nil {
		return x.Friend
	}
	return ""
}

type InviteAcceptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviterId string `protobuf:"bytes,1,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
}

func (x *InviteAcceptRequest) Reset() {
	*x = InviteAcceptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteAcceptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteAcceptRequest) ProtoMessage() {}

func (x *InviteAcceptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteAcceptRequest.ProtoReflect.Descriptor instead.
func (*InviteAcceptRequest) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{1}
}

func (x *InviteAcceptRequest) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

type InviteRejectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviterId string `protobuf:"bytes,1,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
}

func (x *InviteRejectRequest) Reset() {
	*x = InviteRejectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRejectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRejectRequest) ProtoMessage() {}

func (x *InviteRejectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRejectRequest.ProtoReflect.Descriptor instead.
func (*InviteRejectRequest) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{2}
}

func (x *InviteRejectRequest) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

type InviteFriendInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId     string `protobuf:"bytes,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	FriendName   string `protobuf:"bytes,2,opt,name=friend_name,json=friendName,proto3" json:"friend_name,omitempty"`
	FriendAvatar string `protobuf:"bytes,3,opt,name=friend_avatar,json=friendAvatar,proto3" json:"friend_avatar,omitempty"`
	CreatedAt    string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AcceptedAt   string `protobuf:"bytes,5,opt,name=accepted_at,json=acceptedAt,proto3" json:"accepted_at,omitempty"`
	RejectedAt   string `protobuf:"bytes,6,opt,name=rejected_at,json=rejectedAt,proto3" json:"rejected_at,omitempty"`
}

func (x *InviteFriendInvite) Reset() {
	*x = InviteFriendInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteFriendInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteFriendInvite) ProtoMessage() {}

func (x *InviteFriendInvite) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteFriendInvite.ProtoReflect.Descriptor instead.
func (*InviteFriendInvite) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{3}
}

func (x *InviteFriendInvite) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

func (x *InviteFriendInvite) GetFriendName() string {
	if x != nil {
		return x.FriendName
	}
	return ""
}

func (x *InviteFriendInvite) GetFriendAvatar() string {
	if x != nil {
		return x.FriendAvatar
	}
	return ""
}

func (x *InviteFriendInvite) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *InviteFriendInvite) GetAcceptedAt() string {
	if x != nil {
		return x.AcceptedAt
	}
	return ""
}

func (x *InviteFriendInvite) GetRejectedAt() string {
	if x != nil {
		return x.RejectedAt
	}
	return ""
}

type InviteFromMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invites []*InviteFriendInvite `protobuf:"bytes,1,rep,name=invites,proto3" json:"invites,omitempty"`
}

func (x *InviteFromMeResponse) Reset() {
	*x = InviteFromMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteFromMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteFromMeResponse) ProtoMessage() {}

func (x *InviteFromMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteFromMeResponse.ProtoReflect.Descriptor instead.
func (*InviteFromMeResponse) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{4}
}

func (x *InviteFromMeResponse) GetInvites() []*InviteFriendInvite {
	if x != nil {
		return x.Invites
	}
	return nil
}

type InviteForMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invites []*InviteFriendInvite `protobuf:"bytes,1,rep,name=invites,proto3" json:"invites,omitempty"`
}

func (x *InviteForMeResponse) Reset() {
	*x = InviteForMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invite_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteForMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteForMeResponse) ProtoMessage() {}

func (x *InviteForMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invite_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteForMeResponse.ProtoReflect.Descriptor instead.
func (*InviteForMeResponse) Descriptor() ([]byte, []int) {
	return file_invite_proto_rawDescGZIP(), []int{5}
}

func (x *InviteForMeResponse) GetInvites() []*InviteFriendInvite {
	if x != nil {
		return x.Invites
	}
	return nil
}

var File_invite_proto protoreflect.FileDescriptor

var file_invite_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x22,
	0x34, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd8, 0x01, 0x0a, 0x12,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x22, 0x48, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x4d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x32, 0xff, 0x01, 0x0a, 0x0d,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a,
	0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a,
	0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2f, 0x0a,
	0x06, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x05, 0x46, 0x6f, 0x72, 0x4d, 0x65, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invite_proto_rawDescOnce sync.Once
	file_invite_proto_rawDescData = file_invite_proto_rawDesc
)

func file_invite_proto_rawDescGZIP() []byte {
	file_invite_proto_rawDescOnce.Do(func() {
		file_invite_proto_rawDescData = protoimpl.X.CompressGZIP(file_invite_proto_rawDescData)
	})
	return file_invite_proto_rawDescData
}

var file_invite_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_invite_proto_goTypes = []interface{}{
	(*InviteCreateRequest)(nil),  // 0: rpc.InviteCreateRequest
	(*InviteAcceptRequest)(nil),  // 1: rpc.InviteAcceptRequest
	(*InviteRejectRequest)(nil),  // 2: rpc.InviteRejectRequest
	(*InviteFriendInvite)(nil),   // 3: rpc.InviteFriendInvite
	(*InviteFromMeResponse)(nil), // 4: rpc.InviteFromMeResponse
	(*InviteForMeResponse)(nil),  // 5: rpc.InviteForMeResponse
	(*Empty)(nil),                // 6: rpc.Empty
}
var file_invite_proto_depIdxs = []int32{
	3, // 0: rpc.InviteFromMeResponse.invites:type_name -> rpc.InviteFriendInvite
	3, // 1: rpc.InviteForMeResponse.invites:type_name -> rpc.InviteFriendInvite
	0, // 2: rpc.InviteService.Create:input_type -> rpc.InviteCreateRequest
	1, // 3: rpc.InviteService.Accept:input_type -> rpc.InviteAcceptRequest
	2, // 4: rpc.InviteService.Reject:input_type -> rpc.InviteRejectRequest
	6, // 5: rpc.InviteService.FromMe:input_type -> rpc.Empty
	6, // 6: rpc.InviteService.ForMe:input_type -> rpc.Empty
	6, // 7: rpc.InviteService.Create:output_type -> rpc.Empty
	6, // 8: rpc.InviteService.Accept:output_type -> rpc.Empty
	6, // 9: rpc.InviteService.Reject:output_type -> rpc.Empty
	4, // 10: rpc.InviteService.FromMe:output_type -> rpc.InviteFromMeResponse
	5, // 11: rpc.InviteService.ForMe:output_type -> rpc.InviteForMeResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_invite_proto_init() }
func file_invite_proto_init() {
	if File_invite_proto != nil {
		return
	}
	file_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_invite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteCreateRequest); i {
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
		file_invite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteAcceptRequest); i {
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
		file_invite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRejectRequest); i {
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
		file_invite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteFriendInvite); i {
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
		file_invite_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteFromMeResponse); i {
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
		file_invite_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteForMeResponse); i {
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
			RawDescriptor: file_invite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invite_proto_goTypes,
		DependencyIndexes: file_invite_proto_depIdxs,
		MessageInfos:      file_invite_proto_msgTypes,
	}.Build()
	File_invite_proto = out.File
	file_invite_proto_rawDesc = nil
	file_invite_proto_goTypes = nil
	file_invite_proto_depIdxs = nil
}
