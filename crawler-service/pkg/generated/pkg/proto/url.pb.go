// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/proto/url.proto

package crawlerv1

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

type Url struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Method        string                 `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Queue         string                 `protobuf:"bytes,5,opt,name=queue,proto3" json:"queue,omitempty"`
	Domain        string                 `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	IsActive      bool                   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Url) Reset() {
	*x = Url{}
	mi := &file_pkg_proto_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Url) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Url) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Url) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *Url) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Url) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Url) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Url) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateUrlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           *Url                   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUrlRequest) Reset() {
	*x = CreateUrlRequest{}
	mi := &file_pkg_proto_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlRequest) ProtoMessage() {}

func (x *CreateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateUrlRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUrlRequest) GetUrl() *Url {
	if x != nil {
		return x.Url
	}
	return nil
}

type CreateUrlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUrlResponse) Reset() {
	*x = CreateUrlResponse{}
	mi := &file_pkg_proto_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlResponse) ProtoMessage() {}

func (x *CreateUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlResponse.ProtoReflect.Descriptor instead.
func (*CreateUrlResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUrlResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateUrlResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetUrlsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUrlsRequest) Reset() {
	*x = GetUrlsRequest{}
	mi := &file_pkg_proto_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlsRequest) ProtoMessage() {}

func (x *GetUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlsRequest.ProtoReflect.Descriptor instead.
func (*GetUrlsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{3}
}

func (x *GetUrlsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUrlsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetUrlsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Urls          []*Url                 `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUrlsResponse) Reset() {
	*x = GetUrlsResponse{}
	mi := &file_pkg_proto_url_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlsResponse) ProtoMessage() {}

func (x *GetUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlsResponse.ProtoReflect.Descriptor instead.
func (*GetUrlsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{4}
}

func (x *GetUrlsResponse) GetUrls() []*Url {
	if x != nil {
		return x.Urls
	}
	return nil
}

type UpdateUrlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           *Url                   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUrlRequest) Reset() {
	*x = UpdateUrlRequest{}
	mi := &file_pkg_proto_url_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUrlRequest) ProtoMessage() {}

func (x *UpdateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUrlRequest.ProtoReflect.Descriptor instead.
func (*UpdateUrlRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUrlRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUrlRequest) GetUrl() *Url {
	if x != nil {
		return x.Url
	}
	return nil
}

type UpdateUrlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUrlResponse) Reset() {
	*x = UpdateUrlResponse{}
	mi := &file_pkg_proto_url_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUrlResponse) ProtoMessage() {}

func (x *UpdateUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_url_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUrlResponse.ProtoReflect.Descriptor instead.
func (*UpdateUrlResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_url_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUrlResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUrlResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_pkg_proto_url_proto protoreflect.FileDescriptor

const file_pkg_proto_url_proto_rawDesc = "" +
	"\n" +
	"\x13pkg/proto/url.proto\x12\n" +
	"crawler.v1\x1a\x1cgoogle/api/annotations.proto\"\xea\x01\n" +
	"\x03Url\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12\x16\n" +
	"\x06method\x18\x03 \x01(\tR\x06method\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x14\n" +
	"\x05queue\x18\x05 \x01(\tR\x05queue\x12\x16\n" +
	"\x06domain\x18\x06 \x01(\tR\x06domain\x12\x1b\n" +
	"\tis_active\x18\a \x01(\bR\bisActive\x12\x1d\n" +
	"\n" +
	"created_at\x18\b \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\t \x01(\tR\tupdatedAt\"5\n" +
	"\x10CreateUrlRequest\x12!\n" +
	"\x03url\x18\x01 \x01(\v2\x0f.crawler.v1.UrlR\x03url\";\n" +
	"\x11CreateUrlResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\">\n" +
	"\x0eGetUrlsRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x05R\x06offset\"6\n" +
	"\x0fGetUrlsResponse\x12#\n" +
	"\x04urls\x18\x01 \x03(\v2\x0f.crawler.v1.UrlR\x04urls\"E\n" +
	"\x10UpdateUrlRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12!\n" +
	"\x03url\x18\x02 \x01(\v2\x0f.crawler.v1.UrlR\x03url\";\n" +
	"\x11UpdateUrlResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status2\xb0\x02\n" +
	"\n" +
	"UrlService\x12`\n" +
	"\tCreateUrl\x12\x1c.crawler.v1.CreateUrlRequest\x1a\x1d.crawler.v1.CreateUrlResponse\"\x16\x82\xd3\xe4\x93\x02\x10:\x01*\"\v/api/v1/url\x12X\n" +
	"\aGetUrls\x12\x1a.crawler.v1.GetUrlsRequest\x1a\x1b.crawler.v1.GetUrlsResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/api/v1/urls\x12f\n" +
	"\tUpdateUrl\x12\x1c.crawler.v1.UpdateUrlRequest\x1a\x1d.crawler.v1.UpdateUrlResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\x1a\x11/api/v1/urls/{id}B\x88\x01\n" +
	"\x0ecom.crawler.v1B\bUrlProtoP\x01Z#crawler-service/pkg/proto;crawlerv1\xa2\x02\x03CXX\xaa\x02\n" +
	"Crawler.V1\xca\x02\n" +
	"Crawler\\V1\xe2\x02\x16Crawler\\V1\\GPBMetadata\xea\x02\vCrawler::V1b\x06proto3"

var (
	file_pkg_proto_url_proto_rawDescOnce sync.Once
	file_pkg_proto_url_proto_rawDescData []byte
)

func file_pkg_proto_url_proto_rawDescGZIP() []byte {
	file_pkg_proto_url_proto_rawDescOnce.Do(func() {
		file_pkg_proto_url_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_url_proto_rawDesc), len(file_pkg_proto_url_proto_rawDesc)))
	})
	return file_pkg_proto_url_proto_rawDescData
}

var file_pkg_proto_url_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_url_proto_goTypes = []any{
	(*Url)(nil),               // 0: crawler.v1.Url
	(*CreateUrlRequest)(nil),  // 1: crawler.v1.CreateUrlRequest
	(*CreateUrlResponse)(nil), // 2: crawler.v1.CreateUrlResponse
	(*GetUrlsRequest)(nil),    // 3: crawler.v1.GetUrlsRequest
	(*GetUrlsResponse)(nil),   // 4: crawler.v1.GetUrlsResponse
	(*UpdateUrlRequest)(nil),  // 5: crawler.v1.UpdateUrlRequest
	(*UpdateUrlResponse)(nil), // 6: crawler.v1.UpdateUrlResponse
}
var file_pkg_proto_url_proto_depIdxs = []int32{
	0, // 0: crawler.v1.CreateUrlRequest.url:type_name -> crawler.v1.Url
	0, // 1: crawler.v1.GetUrlsResponse.urls:type_name -> crawler.v1.Url
	0, // 2: crawler.v1.UpdateUrlRequest.url:type_name -> crawler.v1.Url
	1, // 3: crawler.v1.UrlService.CreateUrl:input_type -> crawler.v1.CreateUrlRequest
	3, // 4: crawler.v1.UrlService.GetUrls:input_type -> crawler.v1.GetUrlsRequest
	5, // 5: crawler.v1.UrlService.UpdateUrl:input_type -> crawler.v1.UpdateUrlRequest
	2, // 6: crawler.v1.UrlService.CreateUrl:output_type -> crawler.v1.CreateUrlResponse
	4, // 7: crawler.v1.UrlService.GetUrls:output_type -> crawler.v1.GetUrlsResponse
	6, // 8: crawler.v1.UrlService.UpdateUrl:output_type -> crawler.v1.UpdateUrlResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_url_proto_init() }
func file_pkg_proto_url_proto_init() {
	if File_pkg_proto_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_url_proto_rawDesc), len(file_pkg_proto_url_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_url_proto_goTypes,
		DependencyIndexes: file_pkg_proto_url_proto_depIdxs,
		MessageInfos:      file_pkg_proto_url_proto_msgTypes,
	}.Build()
	File_pkg_proto_url_proto = out.File
	file_pkg_proto_url_proto_goTypes = nil
	file_pkg_proto_url_proto_depIdxs = nil
}
