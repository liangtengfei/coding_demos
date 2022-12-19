// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: laptop_service.proto

package pb

import (
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

type Laptop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand     string                 `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Laptop) Reset() {
	*x = Laptop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Laptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Laptop) ProtoMessage() {}

func (x *Laptop) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Laptop.ProtoReflect.Descriptor instead.
func (*Laptop) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{0}
}

func (x *Laptop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Laptop) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Laptop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Laptop) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type LaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LaptopRequest) Reset() {
	*x = LaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaptopRequest) ProtoMessage() {}

func (x *LaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaptopRequest.ProtoReflect.Descriptor instead.
func (*LaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{1}
}

func (x *LaptopRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_laptop_service_proto protoreflect.FileDescriptor

var file_laptop_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x33, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x06, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x4c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x60, 0x0a, 0x0d, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x21, 0x2e, 0x71,
	0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x33, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x33, 0x2e, 0x70, 0x63,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x00, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_service_proto_rawDescOnce sync.Once
	file_laptop_service_proto_rawDescData = file_laptop_service_proto_rawDesc
)

func file_laptop_service_proto_rawDescGZIP() []byte {
	file_laptop_service_proto_rawDescOnce.Do(func() {
		file_laptop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_service_proto_rawDescData)
	})
	return file_laptop_service_proto_rawDescData
}

var file_laptop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_laptop_service_proto_goTypes = []interface{}{
	(*Laptop)(nil),                // 0: quickstart3.pcbook.Laptop
	(*LaptopRequest)(nil),         // 1: quickstart3.pcbook.LaptopRequest
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_laptop_service_proto_depIdxs = []int32{
	2, // 0: quickstart3.pcbook.Laptop.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: quickstart3.pcbook.LaptopService.DetailLaptop:input_type -> quickstart3.pcbook.LaptopRequest
	0, // 2: quickstart3.pcbook.LaptopService.DetailLaptop:output_type -> quickstart3.pcbook.Laptop
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_laptop_service_proto_init() }
func file_laptop_service_proto_init() {
	if File_laptop_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_laptop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Laptop); i {
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
		file_laptop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaptopRequest); i {
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
			RawDescriptor: file_laptop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_service_proto_goTypes,
		DependencyIndexes: file_laptop_service_proto_depIdxs,
		MessageInfos:      file_laptop_service_proto_msgTypes,
	}.Build()
	File_laptop_service_proto = out.File
	file_laptop_service_proto_rawDesc = nil
	file_laptop_service_proto_goTypes = nil
	file_laptop_service_proto_depIdxs = nil
}
