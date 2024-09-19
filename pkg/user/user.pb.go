// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: pkg/user/user.proto

package user

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username       string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	FirstName      string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age            int32  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	PhoneNumber    string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Address        string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	City           string `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	Country        string `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	PostalCode     string `protobuf:"bytes,11,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	CreatedAt      int64  `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastLoginAt    int64  `protobuf:"varint,13,opt,name=last_login_at,json=lastLoginAt,proto3" json:"last_login_at,omitempty"`
	IsActive       bool   `protobuf:"varint,14,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	ProfilePicture string `protobuf:"bytes,15,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Occupation     string `protobuf:"bytes,16,opt,name=occupation,proto3" json:"occupation,omitempty"`
	Company        string `protobuf:"bytes,17,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_user_proto_msgTypes[0]
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
	return file_pkg_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Person) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Person) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Person) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Person) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Person) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Person) GetLastLoginAt() int64 {
	if x != nil {
		return x.LastLoginAt
	}
	return 0
}

func (x *Person) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Person) GetProfilePicture() string {
	if x != nil {
		return x.ProfilePicture
	}
	return ""
}

func (x *Person) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *Person) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

var File_pkg_user_user_proto protoreflect.FileDescriptor

var file_pkg_user_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xe7, 0x03, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f,
	0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_user_user_proto_rawDescOnce sync.Once
	file_pkg_user_user_proto_rawDescData = file_pkg_user_user_proto_rawDesc
)

func file_pkg_user_user_proto_rawDescGZIP() []byte {
	file_pkg_user_user_proto_rawDescOnce.Do(func() {
		file_pkg_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_user_user_proto_rawDescData)
	})
	return file_pkg_user_user_proto_rawDescData
}

var file_pkg_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_user_user_proto_goTypes = []any{
	(*Person)(nil), // 0: user.Person
}
var file_pkg_user_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_user_user_proto_init() }
func file_pkg_user_user_proto_init() {
	if File_pkg_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_user_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_pkg_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_user_user_proto_goTypes,
		DependencyIndexes: file_pkg_user_user_proto_depIdxs,
		MessageInfos:      file_pkg_user_user_proto_msgTypes,
	}.Build()
	File_pkg_user_user_proto = out.File
	file_pkg_user_user_proto_rawDesc = nil
	file_pkg_user_user_proto_goTypes = nil
	file_pkg_user_user_proto_depIdxs = nil
}
