// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: code/person.proto

package person

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

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=Ids,proto3" json:"Ids,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_code_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_code_person_proto_rawDescGZIP(), []int{0}
}

func (x *Ids) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_code_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_code_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_code_person_proto protoreflect.FileDescriptor

var file_code_person_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x17, 0x0a, 0x03, 0x49, 0x64, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x49,
	0x64, 0x73, 0x22, 0x3e, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41,
	0x67, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_person_proto_rawDescOnce sync.Once
	file_code_person_proto_rawDescData = file_code_person_proto_rawDesc
)

func file_code_person_proto_rawDescGZIP() []byte {
	file_code_person_proto_rawDescOnce.Do(func() {
		file_code_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_person_proto_rawDescData)
	})
	return file_code_person_proto_rawDescData
}

var file_code_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_code_person_proto_goTypes = []interface{}{
	(*Ids)(nil),    // 0: code.Ids
	(*Person)(nil), // 1: code.Person
}
var file_code_person_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_person_proto_init() }
func file_code_person_proto_init() {
	if File_code_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_code_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
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
		file_code_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_code_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_code_person_proto_goTypes,
		DependencyIndexes: file_code_person_proto_depIdxs,
		MessageInfos:      file_code_person_proto_msgTypes,
	}.Build()
	File_code_person_proto = out.File
	file_code_person_proto_rawDesc = nil
	file_code_person_proto_goTypes = nil
	file_code_person_proto_depIdxs = nil
}
