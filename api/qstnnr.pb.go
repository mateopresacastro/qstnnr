// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: api/qstnnr.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetQuestionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Questions     []*Question            `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionsResponse) Reset() {
	*x = GetQuestionsResponse{}
	mi := &file_api_qstnnr_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsResponse) ProtoMessage() {}

func (x *GetQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{0}
}

func (x *GetQuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

type Question struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Options       []*Option              `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Question) Reset() {
	*x = Question{}
	mi := &file_api_qstnnr_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Option) Reset() {
	*x = Option{}
	mi := &file_api_qstnnr_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{2}
}

func (x *Option) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Option) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SubmitAnswersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Answers       []*Answer              `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitAnswersRequest) Reset() {
	*x = SubmitAnswersRequest{}
	mi := &file_api_qstnnr_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitAnswersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAnswersRequest) ProtoMessage() {}

func (x *SubmitAnswersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAnswersRequest.ProtoReflect.Descriptor instead.
func (*SubmitAnswersRequest) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitAnswersRequest) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	QuestionId    int32                  `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	OptionId      int32                  `protobuf:"varint,2,opt,name=option_id,json=optionId,proto3" json:"option_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Answer) Reset() {
	*x = Answer{}
	mi := &file_api_qstnnr_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{4}
}

func (x *Answer) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answer) GetOptionId() int32 {
	if x != nil {
		return x.OptionId
	}
	return 0
}

type SubmitAnswersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Solutions     []*Solution            `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
	Correct       int32                  `protobuf:"varint,2,opt,name=correct,proto3" json:"correct,omitempty"`
	BetterThan    int32                  `protobuf:"varint,3,opt,name=better_than,json=betterThan,proto3" json:"better_than,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitAnswersResponse) Reset() {
	*x = SubmitAnswersResponse{}
	mi := &file_api_qstnnr_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitAnswersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAnswersResponse) ProtoMessage() {}

func (x *SubmitAnswersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAnswersResponse.ProtoReflect.Descriptor instead.
func (*SubmitAnswersResponse) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitAnswersResponse) GetSolutions() []*Solution {
	if x != nil {
		return x.Solutions
	}
	return nil
}

func (x *SubmitAnswersResponse) GetCorrect() int32 {
	if x != nil {
		return x.Correct
	}
	return 0
}

func (x *SubmitAnswersResponse) GetBetterThan() int32 {
	if x != nil {
		return x.BetterThan
	}
	return 0
}

type Solution struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Question          *Question              `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	CorrectOptionId   int32                  `protobuf:"varint,2,opt,name=correct_option_id,json=correctOptionId,proto3" json:"correct_option_id,omitempty"`
	CorrectOptionText string                 `protobuf:"bytes,3,opt,name=correct_option_text,json=correctOptionText,proto3" json:"correct_option_text,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Solution) Reset() {
	*x = Solution{}
	mi := &file_api_qstnnr_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{6}
}

func (x *Solution) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

func (x *Solution) GetCorrectOptionId() int32 {
	if x != nil {
		return x.CorrectOptionId
	}
	return 0
}

func (x *Solution) GetCorrectOptionText() string {
	if x != nil {
		return x.CorrectOptionText
	}
	return ""
}

type GetSolutionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Solutions     []*Solution            `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSolutionsResponse) Reset() {
	*x = GetSolutionsResponse{}
	mi := &file_api_qstnnr_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSolutionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSolutionsResponse) ProtoMessage() {}

func (x *GetSolutionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_qstnnr_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSolutionsResponse.ProtoReflect.Descriptor instead.
func (*GetSolutionsResponse) Descriptor() ([]byte, []int) {
	return file_api_qstnnr_proto_rawDescGZIP(), []int{7}
}

func (x *GetSolutionsResponse) GetSolutions() []*Solution {
	if x != nil {
		return x.Solutions
	}
	return nil
}

var File_api_qstnnr_proto protoreflect.FileDescriptor

var file_api_qstnnr_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x58, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x28, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f,
	0x74, 0x68, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x22, 0x46, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e,
	0x72, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe9, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x6e, 0x61, 0x69, 0x72, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1c, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1c,
	0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x71,
	0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x61, 0x74, 0x65, 0x6f, 0x70, 0x72, 0x65, 0x73, 0x61, 0x63, 0x61, 0x73, 0x74, 0x72, 0x6f,
	0x2f, 0x71, 0x73, 0x74, 0x6e, 0x6e, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_qstnnr_proto_rawDescOnce sync.Once
	file_api_qstnnr_proto_rawDescData = file_api_qstnnr_proto_rawDesc
)

func file_api_qstnnr_proto_rawDescGZIP() []byte {
	file_api_qstnnr_proto_rawDescOnce.Do(func() {
		file_api_qstnnr_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_qstnnr_proto_rawDescData)
	})
	return file_api_qstnnr_proto_rawDescData
}

var file_api_qstnnr_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_qstnnr_proto_goTypes = []any{
	(*GetQuestionsResponse)(nil),  // 0: qstnnr.GetQuestionsResponse
	(*Question)(nil),              // 1: qstnnr.Question
	(*Option)(nil),                // 2: qstnnr.Option
	(*SubmitAnswersRequest)(nil),  // 3: qstnnr.SubmitAnswersRequest
	(*Answer)(nil),                // 4: qstnnr.Answer
	(*SubmitAnswersResponse)(nil), // 5: qstnnr.SubmitAnswersResponse
	(*Solution)(nil),              // 6: qstnnr.Solution
	(*GetSolutionsResponse)(nil),  // 7: qstnnr.GetSolutionsResponse
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_qstnnr_proto_depIdxs = []int32{
	1, // 0: qstnnr.GetQuestionsResponse.questions:type_name -> qstnnr.Question
	2, // 1: qstnnr.Question.options:type_name -> qstnnr.Option
	4, // 2: qstnnr.SubmitAnswersRequest.answers:type_name -> qstnnr.Answer
	6, // 3: qstnnr.SubmitAnswersResponse.solutions:type_name -> qstnnr.Solution
	1, // 4: qstnnr.Solution.question:type_name -> qstnnr.Question
	6, // 5: qstnnr.GetSolutionsResponse.solutions:type_name -> qstnnr.Solution
	8, // 6: qstnnr.Questionnaire.GetQuestions:input_type -> google.protobuf.Empty
	3, // 7: qstnnr.Questionnaire.SubmitAnswers:input_type -> qstnnr.SubmitAnswersRequest
	8, // 8: qstnnr.Questionnaire.GetSolutions:input_type -> google.protobuf.Empty
	0, // 9: qstnnr.Questionnaire.GetQuestions:output_type -> qstnnr.GetQuestionsResponse
	5, // 10: qstnnr.Questionnaire.SubmitAnswers:output_type -> qstnnr.SubmitAnswersResponse
	7, // 11: qstnnr.Questionnaire.GetSolutions:output_type -> qstnnr.GetSolutionsResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_qstnnr_proto_init() }
func file_api_qstnnr_proto_init() {
	if File_api_qstnnr_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_qstnnr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_qstnnr_proto_goTypes,
		DependencyIndexes: file_api_qstnnr_proto_depIdxs,
		MessageInfos:      file_api_qstnnr_proto_msgTypes,
	}.Build()
	File_api_qstnnr_proto = out.File
	file_api_qstnnr_proto_rawDesc = nil
	file_api_qstnnr_proto_goTypes = nil
	file_api_qstnnr_proto_depIdxs = nil
}
