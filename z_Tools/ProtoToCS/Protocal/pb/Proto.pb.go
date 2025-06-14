// [START declaration]

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.1
// source: Proto.proto

package pb

import (
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

type PlaceHolder3333 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dec     string  `protobuf:"bytes,1,opt,name=dec,proto3" json:"dec,omitempty"`
	Account string  `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Flag    bool    `protobuf:"varint,3,opt,name=flag,proto3" json:"flag,omitempty"`
	FunID   int32   `protobuf:"varint,4,opt,name=funID,proto3" json:"funID,omitempty"`
	Cards   []int32 `protobuf:"varint,5,rep,packed,name=cards,proto3" json:"cards,omitempty"`
}

func (x *PlaceHolder3333) Reset() {
	*x = PlaceHolder3333{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceHolder3333) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceHolder3333) ProtoMessage() {}

func (x *PlaceHolder3333) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceHolder3333.ProtoReflect.Descriptor instead.
func (*PlaceHolder3333) Descriptor() ([]byte, []int) {
	return file_Proto_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceHolder3333) GetDec() string {
	if x != nil {
		return x.Dec
	}
	return ""
}

func (x *PlaceHolder3333) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PlaceHolder3333) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

func (x *PlaceHolder3333) GetFunID() int32 {
	if x != nil {
		return x.FunID
	}
	return 0
}

func (x *PlaceHolder3333) GetCards() []int32 {
	if x != nil {
		return x.Cards
	}
	return nil
}

type AvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Account string  `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Money   int32   `protobuf:"varint,3,opt,name=money,proto3" json:"money,omitempty"`
	TeamId  int32   `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Cards   []int32 `protobuf:"varint,5,rep,packed,name=cards,proto3" json:"cards,omitempty"`
}

func (x *AvatarInfo) Reset() {
	*x = AvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarInfo) ProtoMessage() {}

func (x *AvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarInfo.ProtoReflect.Descriptor instead.
func (*AvatarInfo) Descriptor() ([]byte, []int) {
	return file_Proto_proto_rawDescGZIP(), []int{1}
}

func (x *AvatarInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AvatarInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AvatarInfo) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *AvatarInfo) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *AvatarInfo) GetCards() []int32 {
	if x != nil {
		return x.Cards
	}
	return nil
}

type MatchingSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarInfos map[int32]*AvatarInfo `protobuf:"bytes,1,rep,name=AvatarInfos,proto3" json:"AvatarInfos,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatchingSuccess) Reset() {
	*x = MatchingSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingSuccess) ProtoMessage() {}

func (x *MatchingSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingSuccess.ProtoReflect.Descriptor instead.
func (*MatchingSuccess) Descriptor() ([]byte, []int) {
	return file_Proto_proto_rawDescGZIP(), []int{2}
}

func (x *MatchingSuccess) GetAvatarInfos() map[int32]*AvatarInfo {
	if x != nil {
		return x.AvatarInfos
	}
	return nil
}

type PlayedCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId       int32                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Account      string                `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	CardValue    int32                 `protobuf:"varint,3,opt,name=card_value,json=cardValue,proto3" json:"card_value,omitempty"`
	OutCardValue int32                 `protobuf:"varint,4,opt,name=out_card_value,json=outCardValue,proto3" json:"out_card_value,omitempty"`
	CardsLeftNum int32                 `protobuf:"varint,5,opt,name=cards_left_num,json=cardsLeftNum,proto3" json:"cards_left_num,omitempty"`
	AvatarInfos  map[int32]*AvatarInfo `protobuf:"bytes,6,rep,name=AvatarInfos,proto3" json:"AvatarInfos,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlayedCardInfo) Reset() {
	*x = PlayedCardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayedCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayedCardInfo) ProtoMessage() {}

func (x *PlayedCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayedCardInfo.ProtoReflect.Descriptor instead.
func (*PlayedCardInfo) Descriptor() ([]byte, []int) {
	return file_Proto_proto_rawDescGZIP(), []int{3}
}

func (x *PlayedCardInfo) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *PlayedCardInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PlayedCardInfo) GetCardValue() int32 {
	if x != nil {
		return x.CardValue
	}
	return 0
}

func (x *PlayedCardInfo) GetOutCardValue() int32 {
	if x != nil {
		return x.OutCardValue
	}
	return 0
}

func (x *PlayedCardInfo) GetCardsLeftNum() int32 {
	if x != nil {
		return x.CardsLeftNum
	}
	return 0
}

func (x *PlayedCardInfo) GetAvatarInfos() map[int32]*AvatarInfo {
	if x != nil {
		return x.AvatarInfos
	}
	return nil
}

var File_Proto_proto protoreflect.FileDescriptor

var file_Proto_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x50,
	0x62, 0x22, 0x7d, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x33, 0x33, 0x33, 0x33, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x75, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x75, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x7f, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x50, 0x62, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x4e, 0x0a,
	0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x62, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x02,
	0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x73, 0x4c, 0x65, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x45,
	0x0a, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x50, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x4e, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x62, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_proto_rawDescOnce sync.Once
	file_Proto_proto_rawDescData = file_Proto_proto_rawDesc
)

func file_Proto_proto_rawDescGZIP() []byte {
	file_Proto_proto_rawDescOnce.Do(func() {
		file_Proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_proto_rawDescData)
	})
	return file_Proto_proto_rawDescData
}

var file_Proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Proto_proto_goTypes = []interface{}{
	(*PlaceHolder3333)(nil), // 0: Pb.PlaceHolder3333
	(*AvatarInfo)(nil),      // 1: Pb.AvatarInfo
	(*MatchingSuccess)(nil), // 2: Pb.MatchingSuccess
	(*PlayedCardInfo)(nil),  // 3: Pb.PlayedCardInfo
	nil,                     // 4: Pb.MatchingSuccess.AvatarInfosEntry
	nil,                     // 5: Pb.PlayedCardInfo.AvatarInfosEntry
}
var file_Proto_proto_depIdxs = []int32{
	4, // 0: Pb.MatchingSuccess.AvatarInfos:type_name -> Pb.MatchingSuccess.AvatarInfosEntry
	5, // 1: Pb.PlayedCardInfo.AvatarInfos:type_name -> Pb.PlayedCardInfo.AvatarInfosEntry
	1, // 2: Pb.MatchingSuccess.AvatarInfosEntry.value:type_name -> Pb.AvatarInfo
	1, // 3: Pb.PlayedCardInfo.AvatarInfosEntry.value:type_name -> Pb.AvatarInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Proto_proto_init() }
func file_Proto_proto_init() {
	if File_Proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceHolder3333); i {
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
		file_Proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarInfo); i {
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
		file_Proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingSuccess); i {
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
		file_Proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayedCardInfo); i {
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
			RawDescriptor: file_Proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Proto_proto_goTypes,
		DependencyIndexes: file_Proto_proto_depIdxs,
		MessageInfos:      file_Proto_proto_msgTypes,
	}.Build()
	File_Proto_proto = out.File
	file_Proto_proto_rawDesc = nil
	file_Proto_proto_goTypes = nil
	file_Proto_proto_depIdxs = nil
}
