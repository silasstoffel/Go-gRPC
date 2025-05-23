// Command to generate the Go code:
// protoc --go_out=. --go-grpc_out=. proto/course_category.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/course_category.proto

package pb

import (
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

type Blank struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blank) Reset() {
	*x = Blank{}
	mi := &file_proto_course_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_proto_course_category_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CategoryOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      []*Category            `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryOutput) Reset() {
	*x = CategoryOutput{}
	mi := &file_proto_course_category_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryOutput) ProtoMessage() {}

func (x *CategoryOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryOutput.ProtoReflect.Descriptor instead.
func (*CategoryOutput) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryOutput) GetCategory() []*Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type GetCategoriesOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*Category            `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoriesOutput) Reset() {
	*x = GetCategoriesOutput{}
	mi := &file_proto_course_category_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoriesOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesOutput) ProtoMessage() {}

func (x *GetCategoriesOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesOutput.ProtoReflect.Descriptor instead.
func (*GetCategoriesOutput) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetCategoriesOutput) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CreateCategoryInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryInput) Reset() {
	*x = CreateCategoryInput{}
	mi := &file_proto_course_category_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryInput) ProtoMessage() {}

func (x *CreateCategoryInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryInput.ProtoReflect.Descriptor instead.
func (*CreateCategoryInput) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCategoryInput) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetCategoryInput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryInput) Reset() {
	*x = GetCategoryInput{}
	mi := &file_proto_course_category_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryInput) ProtoMessage() {}

func (x *GetCategoryInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_category_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryInput.ProtoReflect.Descriptor instead.
func (*GetCategoryInput) Descriptor() ([]byte, []int) {
	return file_proto_course_category_proto_rawDescGZIP(), []int{5}
}

func (x *GetCategoryInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_course_category_proto protoreflect.FileDescriptor

const file_proto_course_category_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/course_category.proto\x12\x02pb\"\a\n" +
	"\x05Blank\"P\n" +
	"\bCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\":\n" +
	"\x0eCategoryOutput\x12(\n" +
	"\bcategory\x18\x01 \x03(\v2\f.pb.CategoryR\bcategory\"C\n" +
	"\x13GetCategoriesOutput\x12,\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\f.pb.CategoryR\n" +
	"categories\"K\n" +
	"\x13CreateCategoryInput\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"\"\n" +
	"\x10GetCategoryInput\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xd8\x02\n" +
	"\x0fCategoryService\x129\n" +
	"\x0eCreateCategory\x12\x17.pb.CreateCategoryInput\x1a\f.pb.Category\"\x00\x12L\n" +
	"\x14CreateCategoryStream\x12\x17.pb.CreateCategoryInput\x1a\x17.pb.GetCategoriesOutput\"\x00(\x01\x12P\n" +
	"!CreateCategoryStreamBidirectional\x12\x17.pb.CreateCategoryInput\x1a\f.pb.Category\"\x00(\x010\x01\x125\n" +
	"\rGetCategories\x12\t.pb.Blank\x1a\x17.pb.GetCategoriesOutput\"\x00\x123\n" +
	"\vGetCategory\x12\x14.pb.GetCategoryInput\x1a\f.pb.Category\"\x00B\rZ\vinternal/pbb\x06proto3"

var (
	file_proto_course_category_proto_rawDescOnce sync.Once
	file_proto_course_category_proto_rawDescData []byte
)

func file_proto_course_category_proto_rawDescGZIP() []byte {
	file_proto_course_category_proto_rawDescOnce.Do(func() {
		file_proto_course_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_course_category_proto_rawDesc), len(file_proto_course_category_proto_rawDesc)))
	})
	return file_proto_course_category_proto_rawDescData
}

var file_proto_course_category_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_course_category_proto_goTypes = []any{
	(*Blank)(nil),               // 0: pb.Blank
	(*Category)(nil),            // 1: pb.Category
	(*CategoryOutput)(nil),      // 2: pb.CategoryOutput
	(*GetCategoriesOutput)(nil), // 3: pb.GetCategoriesOutput
	(*CreateCategoryInput)(nil), // 4: pb.CreateCategoryInput
	(*GetCategoryInput)(nil),    // 5: pb.GetCategoryInput
}
var file_proto_course_category_proto_depIdxs = []int32{
	1, // 0: pb.CategoryOutput.category:type_name -> pb.Category
	1, // 1: pb.GetCategoriesOutput.categories:type_name -> pb.Category
	4, // 2: pb.CategoryService.CreateCategory:input_type -> pb.CreateCategoryInput
	4, // 3: pb.CategoryService.CreateCategoryStream:input_type -> pb.CreateCategoryInput
	4, // 4: pb.CategoryService.CreateCategoryStreamBidirectional:input_type -> pb.CreateCategoryInput
	0, // 5: pb.CategoryService.GetCategories:input_type -> pb.Blank
	5, // 6: pb.CategoryService.GetCategory:input_type -> pb.GetCategoryInput
	1, // 7: pb.CategoryService.CreateCategory:output_type -> pb.Category
	3, // 8: pb.CategoryService.CreateCategoryStream:output_type -> pb.GetCategoriesOutput
	1, // 9: pb.CategoryService.CreateCategoryStreamBidirectional:output_type -> pb.Category
	3, // 10: pb.CategoryService.GetCategories:output_type -> pb.GetCategoriesOutput
	1, // 11: pb.CategoryService.GetCategory:output_type -> pb.Category
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_course_category_proto_init() }
func file_proto_course_category_proto_init() {
	if File_proto_course_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_course_category_proto_rawDesc), len(file_proto_course_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_course_category_proto_goTypes,
		DependencyIndexes: file_proto_course_category_proto_depIdxs,
		MessageInfos:      file_proto_course_category_proto_msgTypes,
	}.Build()
	File_proto_course_category_proto = out.File
	file_proto_course_category_proto_goTypes = nil
	file_proto_course_category_proto_depIdxs = nil
}
