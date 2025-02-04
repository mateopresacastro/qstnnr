// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/qstnnr.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Questionnaire_GetQuestions_FullMethodName  = "/qstnnr.Questionnaire/GetQuestions"
	Questionnaire_SubmitAnswers_FullMethodName = "/qstnnr.Questionnaire/SubmitAnswers"
	Questionnaire_GetSolutions_FullMethodName  = "/qstnnr.Questionnaire/GetSolutions"
)

// QuestionnaireClient is the client API for Questionnaire service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionnaireClient interface {
	// GetQuestions gets all the questions and options for each.
	GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (*GetQuestionsResponse, error)
	// SubmitAnswers submits the answers to be evaluated.
	SubmitAnswers(ctx context.Context, in *SubmitAnswersRequest, opts ...grpc.CallOption) (*SubmitAnswersResponse, error)
	// GetSolutons gets all solutions if the user wants to check them in isolation.
	GetSolutions(ctx context.Context, in *GetSolutionsRequest, opts ...grpc.CallOption) (*GetSolutionsResponse, error)
}

type questionnaireClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionnaireClient(cc grpc.ClientConnInterface) QuestionnaireClient {
	return &questionnaireClient{cc}
}

func (c *questionnaireClient) GetQuestions(ctx context.Context, in *GetQuestionsRequest, opts ...grpc.CallOption) (*GetQuestionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuestionsResponse)
	err := c.cc.Invoke(ctx, Questionnaire_GetQuestions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) SubmitAnswers(ctx context.Context, in *SubmitAnswersRequest, opts ...grpc.CallOption) (*SubmitAnswersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswersResponse)
	err := c.cc.Invoke(ctx, Questionnaire_SubmitAnswers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionnaireClient) GetSolutions(ctx context.Context, in *GetSolutionsRequest, opts ...grpc.CallOption) (*GetSolutionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSolutionsResponse)
	err := c.cc.Invoke(ctx, Questionnaire_GetSolutions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionnaireServer is the server API for Questionnaire service.
// All implementations must embed UnimplementedQuestionnaireServer
// for forward compatibility.
type QuestionnaireServer interface {
	// GetQuestions gets all the questions and options for each.
	GetQuestions(context.Context, *GetQuestionsRequest) (*GetQuestionsResponse, error)
	// SubmitAnswers submits the answers to be evaluated.
	SubmitAnswers(context.Context, *SubmitAnswersRequest) (*SubmitAnswersResponse, error)
	// GetSolutons gets all solutions if the user wants to check them in isolation.
	GetSolutions(context.Context, *GetSolutionsRequest) (*GetSolutionsResponse, error)
	mustEmbedUnimplementedQuestionnaireServer()
}

// UnimplementedQuestionnaireServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuestionnaireServer struct{}

func (UnimplementedQuestionnaireServer) GetQuestions(context.Context, *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestions not implemented")
}
func (UnimplementedQuestionnaireServer) SubmitAnswers(context.Context, *SubmitAnswersRequest) (*SubmitAnswersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswers not implemented")
}
func (UnimplementedQuestionnaireServer) GetSolutions(context.Context, *GetSolutionsRequest) (*GetSolutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSolutions not implemented")
}
func (UnimplementedQuestionnaireServer) mustEmbedUnimplementedQuestionnaireServer() {}
func (UnimplementedQuestionnaireServer) testEmbeddedByValue()                       {}

// UnsafeQuestionnaireServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionnaireServer will
// result in compilation errors.
type UnsafeQuestionnaireServer interface {
	mustEmbedUnimplementedQuestionnaireServer()
}

func RegisterQuestionnaireServer(s grpc.ServiceRegistrar, srv QuestionnaireServer) {
	// If the following call pancis, it indicates UnimplementedQuestionnaireServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Questionnaire_ServiceDesc, srv)
}

func _Questionnaire_GetQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).GetQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_GetQuestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).GetQuestions(ctx, req.(*GetQuestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_SubmitAnswers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).SubmitAnswers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_SubmitAnswers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).SubmitAnswers(ctx, req.(*SubmitAnswersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Questionnaire_GetSolutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSolutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionnaireServer).GetSolutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Questionnaire_GetSolutions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionnaireServer).GetSolutions(ctx, req.(*GetSolutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Questionnaire_ServiceDesc is the grpc.ServiceDesc for Questionnaire service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Questionnaire_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qstnnr.Questionnaire",
	HandlerType: (*QuestionnaireServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestions",
			Handler:    _Questionnaire_GetQuestions_Handler,
		},
		{
			MethodName: "SubmitAnswers",
			Handler:    _Questionnaire_SubmitAnswers_Handler,
		},
		{
			MethodName: "GetSolutions",
			Handler:    _Questionnaire_GetSolutions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qstnnr.proto",
}
