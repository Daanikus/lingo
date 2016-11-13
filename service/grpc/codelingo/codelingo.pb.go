// Code generated by protoc-gen-go.
// source: codelingo.proto
// DO NOT EDIT!

/*
Package codelingo is a generated protocol buffer package.

It is generated from these files:
	codelingo.proto

It has these top-level messages:
	ListFactsRequest
	FactList
	Fact
	ListLexiconsRequest
	ListLexiconsReply
	SessionRequest
	SessionReply
	QueryRequest
	QueryReply
	ReviewRequest
	ReviewReply
	Issue
	IssueRange
	Position
*/
package codelingo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListFactsRequest struct {
	Lexicon string `protobuf:"bytes,1,opt,name=lexicon" json:"lexicon,omitempty"`
}

func (m *ListFactsRequest) Reset()                    { *m = ListFactsRequest{} }
func (m *ListFactsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFactsRequest) ProtoMessage()               {}
func (*ListFactsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FactList struct {
	Facts []*Fact `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty"`
}

func (m *FactList) Reset()                    { *m = FactList{} }
func (m *FactList) String() string            { return proto.CompactTextString(m) }
func (*FactList) ProtoMessage()               {}
func (*FactList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FactList) GetFacts() []*Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

type Fact struct {
	Facts      []*Fact  `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty"`
	Properties []string `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
	Kind       string   `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
}

func (m *Fact) Reset()                    { *m = Fact{} }
func (m *Fact) String() string            { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()               {}
func (*Fact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Fact) GetFacts() []*Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

type ListLexiconsRequest struct {
}

func (m *ListLexiconsRequest) Reset()                    { *m = ListLexiconsRequest{} }
func (m *ListLexiconsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLexiconsRequest) ProtoMessage()               {}
func (*ListLexiconsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListLexiconsReply struct {
	Lexicons []string `protobuf:"bytes,1,rep,name=lexicons" json:"lexicons,omitempty"`
}

func (m *ListLexiconsReply) Reset()                    { *m = ListLexiconsReply{} }
func (m *ListLexiconsReply) String() string            { return proto.CompactTextString(m) }
func (*ListLexiconsReply) ProtoMessage()               {}
func (*ListLexiconsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SessionRequest struct {
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SessionReply struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *SessionReply) Reset()                    { *m = SessionReply{} }
func (m *SessionReply) String() string            { return proto.CompactTextString(m) }
func (*SessionReply) ProtoMessage()               {}
func (*SessionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// The request message containing the CLQL source code.
type QueryRequest struct {
	Clql string `protobuf:"bytes,1,opt,name=clql" json:"clql,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// The query response.
type QueryReply struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *QueryReply) Reset()                    { *m = QueryReply{} }
func (m *QueryReply) String() string            { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()               {}
func (*QueryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// The request message containing the files or directories to review.
type ReviewRequest struct {
	Key           string   `protobuf:"bytes,12,opt,name=key" json:"key,omitempty"`
	Host          string   `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Owner         string   `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Repo          string   `protobuf:"bytes,3,opt,name=repo" json:"repo,omitempty"`
	Sha           string   `protobuf:"bytes,4,opt,name=sha" json:"sha,omitempty"`
	FilesAndDirs  []string `protobuf:"bytes,5,rep,name=filesAndDirs" json:"filesAndDirs,omitempty"`
	Recursive     bool     `protobuf:"varint,6,opt,name=recursive" json:"recursive,omitempty"`
	Patches       []string `protobuf:"bytes,7,rep,name=Patches" json:"Patches,omitempty"`
	IsPullRequest bool     `protobuf:"varint,8,opt,name=isPullRequest" json:"isPullRequest,omitempty"`
	PullRequestID int64    `protobuf:"varint,9,opt,name=pullRequestID" json:"pullRequestID,omitempty"`
	Vcs           string   `protobuf:"bytes,10,opt,name=vcs" json:"vcs,omitempty"`
	Dotlingo      string   `protobuf:"bytes,11,opt,name=dotlingo" json:"dotlingo,omitempty"`
}

func (m *ReviewRequest) Reset()                    { *m = ReviewRequest{} }
func (m *ReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*ReviewRequest) ProtoMessage()               {}
func (*ReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Empty, we use the control queue once a review is started
type ReviewReply struct {
}

func (m *ReviewReply) Reset()                    { *m = ReviewReply{} }
func (m *ReviewReply) String() string            { return proto.CompactTextString(m) }
func (*ReviewReply) ProtoMessage()               {}
func (*ReviewReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Issue returned from a review.
type Issue struct {
	// The name of the issue.
	Name          string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Position      *IssueRange       `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Comment       string            `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	CtxBefore     string            `protobuf:"bytes,4,opt,name=ctxBefore" json:"ctxBefore,omitempty"`
	LineText      string            `protobuf:"bytes,5,opt,name=lineText" json:"lineText,omitempty"`
	CtxAfter      string            `protobuf:"bytes,6,opt,name=ctxAfter" json:"ctxAfter,omitempty"`
	Metrics       map[string]string `protobuf:"bytes,7,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tags          []string          `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
	Link          string            `protobuf:"bytes,9,opt,name=link" json:"link,omitempty"`
	NewCode       bool              `protobuf:"varint,10,opt,name=newCode" json:"newCode,omitempty"`
	Patch         string            `protobuf:"bytes,11,opt,name=patch" json:"patch,omitempty"`
	Err           string            `protobuf:"bytes,12,opt,name=err" json:"err,omitempty"`
	Discard       bool              `protobuf:"varint,13,opt,name=discard" json:"discard,omitempty"`
	DiscardReason string            `protobuf:"bytes,14,opt,name=discardReason" json:"discardReason,omitempty"`
}

func (m *Issue) Reset()                    { *m = Issue{} }
func (m *Issue) String() string            { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()               {}
func (*Issue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Issue) GetPosition() *IssueRange {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Issue) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type IssueRange struct {
	Start *Position `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *Position `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *IssueRange) Reset()                    { *m = IssueRange{} }
func (m *IssueRange) String() string            { return proto.CompactTextString(m) }
func (*IssueRange) ProtoMessage()               {}
func (*IssueRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IssueRange) GetStart() *Position {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *IssueRange) GetEnd() *Position {
	if m != nil {
		return m.End
	}
	return nil
}

type Position struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Line     int64  `protobuf:"varint,3,opt,name=Line" json:"Line,omitempty"`
	Column   int64  `protobuf:"varint,4,opt,name=Column" json:"Column,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*ListFactsRequest)(nil), "codelingo.ListFactsRequest")
	proto.RegisterType((*FactList)(nil), "codelingo.FactList")
	proto.RegisterType((*Fact)(nil), "codelingo.fact")
	proto.RegisterType((*ListLexiconsRequest)(nil), "codelingo.ListLexiconsRequest")
	proto.RegisterType((*ListLexiconsReply)(nil), "codelingo.ListLexiconsReply")
	proto.RegisterType((*SessionRequest)(nil), "codelingo.SessionRequest")
	proto.RegisterType((*SessionReply)(nil), "codelingo.SessionReply")
	proto.RegisterType((*QueryRequest)(nil), "codelingo.QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "codelingo.QueryReply")
	proto.RegisterType((*ReviewRequest)(nil), "codelingo.ReviewRequest")
	proto.RegisterType((*ReviewReply)(nil), "codelingo.ReviewReply")
	proto.RegisterType((*Issue)(nil), "codelingo.Issue")
	proto.RegisterType((*IssueRange)(nil), "codelingo.IssueRange")
	proto.RegisterType((*Position)(nil), "codelingo.Position")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CodeLingo service

type CodeLingoClient interface {
	// Initialise session
	Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error)
	// Reviews files
	Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error)
	// Queries source code
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
	// Lists available lexicons
	ListLexicons(ctx context.Context, in *ListLexiconsRequest, opts ...grpc.CallOption) (*ListLexiconsReply, error)
	ListFacts(ctx context.Context, in *ListFactsRequest, opts ...grpc.CallOption) (*FactList, error)
}

type codeLingoClient struct {
	cc *grpc.ClientConn
}

func NewCodeLingoClient(cc *grpc.ClientConn) CodeLingoClient {
	return &codeLingoClient{cc}
}

func (c *codeLingoClient) Session(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionReply, error) {
	out := new(SessionReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Session", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewReply, error) {
	out := new(ReviewReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Review", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) ListLexicons(ctx context.Context, in *ListLexiconsRequest, opts ...grpc.CallOption) (*ListLexiconsReply, error) {
	out := new(ListLexiconsReply)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/ListLexicons", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeLingoClient) ListFacts(ctx context.Context, in *ListFactsRequest, opts ...grpc.CallOption) (*FactList, error) {
	out := new(FactList)
	err := grpc.Invoke(ctx, "/codelingo.CodeLingo/ListFacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CodeLingo service

type CodeLingoServer interface {
	// Initialise session
	Session(context.Context, *SessionRequest) (*SessionReply, error)
	// Reviews files
	Review(context.Context, *ReviewRequest) (*ReviewReply, error)
	// Queries source code
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	// Lists available lexicons
	ListLexicons(context.Context, *ListLexiconsRequest) (*ListLexiconsReply, error)
	ListFacts(context.Context, *ListFactsRequest) (*FactList, error)
}

func RegisterCodeLingoServer(s *grpc.Server, srv CodeLingoServer) {
	s.RegisterService(&_CodeLingo_serviceDesc, srv)
}

func _CodeLingo_Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Session(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_Review_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Review(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Review",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Review(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_ListLexicons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLexiconsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).ListLexicons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/ListLexicons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).ListLexicons(ctx, req.(*ListLexiconsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeLingo_ListFacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeLingoServer).ListFacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/codelingo.CodeLingo/ListFacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeLingoServer).ListFacts(ctx, req.(*ListFactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeLingo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "codelingo.CodeLingo",
	HandlerType: (*CodeLingoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Session",
			Handler:    _CodeLingo_Session_Handler,
		},
		{
			MethodName: "Review",
			Handler:    _CodeLingo_Review_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _CodeLingo_Query_Handler,
		},
		{
			MethodName: "ListLexicons",
			Handler:    _CodeLingo_ListLexicons_Handler,
		},
		{
			MethodName: "ListFacts",
			Handler:    _CodeLingo_ListFacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "codelingo.proto",
}

func init() { proto.RegisterFile("codelingo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xef, 0x6e, 0x1b, 0x45,
	0x10, 0xaf, 0x7d, 0x71, 0xe2, 0x1b, 0x3b, 0x69, 0xd8, 0x36, 0x65, 0x31, 0xa5, 0xb2, 0x4e, 0xad,
	0x14, 0x24, 0x14, 0x54, 0xf3, 0x01, 0xa8, 0x90, 0xaa, 0x26, 0x05, 0xa9, 0x92, 0x01, 0xb3, 0x20,
	0xf1, 0x0d, 0xe9, 0x38, 0x8f, 0x93, 0x25, 0xe7, 0xdd, 0xeb, 0xee, 0xda, 0x89, 0x5f, 0x87, 0x97,
	0xe1, 0x0d, 0xf8, 0xc6, 0xbb, 0xa0, 0xd9, 0xdb, 0x3d, 0x9f, 0xdb, 0x20, 0xf1, 0x6d, 0x7e, 0xf3,
	0x6f, 0x67, 0x7f, 0x33, 0x3b, 0x0b, 0xf7, 0x0b, 0x3d, 0xc7, 0x52, 0xaa, 0x4b, 0x7d, 0x56, 0x19,
	0xed, 0x34, 0x4b, 0x1b, 0x45, 0xf6, 0x19, 0x1c, 0x4f, 0xa5, 0x75, 0xdf, 0xe5, 0x85, 0xb3, 0x02,
	0xdf, 0xae, 0xd0, 0x3a, 0xc6, 0xe1, 0xa0, 0xc4, 0x5b, 0x59, 0x68, 0xc5, 0x3b, 0xe3, 0xce, 0x69,
	0x2a, 0x22, 0xcc, 0x9e, 0x43, 0x9f, 0x3c, 0x29, 0x82, 0x3d, 0x83, 0xde, 0x82, 0xa2, 0x78, 0x67,
	0x9c, 0x9c, 0x0e, 0x26, 0xf7, 0xcf, 0xb6, 0xa7, 0x90, 0x5e, 0xd4, 0xd6, 0x2c, 0x87, 0x3d, 0x12,
	0xfe, 0xa7, 0x3b, 0x7b, 0x02, 0x50, 0x19, 0x5d, 0xa1, 0x71, 0x12, 0x2d, 0xef, 0x8e, 0x93, 0xd3,
	0x54, 0xb4, 0x34, 0x8c, 0xc1, 0xde, 0xb5, 0x54, 0x73, 0x9e, 0xf8, 0xc2, 0xbc, 0x9c, 0x9d, 0xc0,
	0x03, 0xaa, 0x68, 0x5a, 0x17, 0x19, 0xaf, 0x91, 0x7d, 0x0e, 0x1f, 0xec, 0xaa, 0xab, 0x72, 0xc3,
	0x46, 0xd0, 0x0f, 0x97, 0xa9, 0x2b, 0x49, 0x45, 0x83, 0xb3, 0x63, 0x38, 0xfa, 0x19, 0xad, 0x95,
	0x5a, 0xc5, 0x14, 0x63, 0x18, 0x36, 0x1a, 0x8a, 0x3e, 0x86, 0xe4, 0x1a, 0x37, 0x81, 0x15, 0x12,
	0xb3, 0x0c, 0x86, 0x3f, 0xad, 0xd0, 0x6c, 0x22, 0x77, 0x0c, 0xf6, 0x8a, 0xf2, 0x6d, 0x19, 0x5c,
	0xbc, 0x9c, 0x3d, 0x05, 0x08, 0x3e, 0x94, 0xe3, 0x11, 0xec, 0x1b, 0xb4, 0xab, 0xd2, 0x05, 0x9f,
	0x80, 0xb2, 0xbf, 0xba, 0x70, 0x28, 0x70, 0x2d, 0xf1, 0x26, 0xe6, 0x0a, 0xa7, 0x0d, 0x9b, 0xd3,
	0x28, 0xfb, 0x95, 0xb6, 0x31, 0xd2, 0xcb, 0xec, 0x21, 0xf4, 0xf4, 0x8d, 0x42, 0xc3, 0xbb, 0x5e,
	0x59, 0x03, 0xf2, 0x34, 0x58, 0xe9, 0xc8, 0x13, 0xc9, 0x94, 0xcf, 0x5e, 0xe5, 0x7c, 0xaf, 0xce,
	0x67, 0xaf, 0x72, 0x96, 0xc1, 0x70, 0x21, 0x4b, 0xb4, 0xaf, 0xd4, 0xfc, 0xb5, 0x34, 0x96, 0xf7,
	0x3c, 0x23, 0x3b, 0x3a, 0xf6, 0x18, 0x52, 0x83, 0xc5, 0xca, 0x58, 0xb9, 0x46, 0xbe, 0x3f, 0xee,
	0x9c, 0xf6, 0xc5, 0x56, 0x41, 0xb3, 0x32, 0xcb, 0x5d, 0x71, 0x85, 0x96, 0x1f, 0xf8, 0xe0, 0x08,
	0xd9, 0x53, 0x38, 0x94, 0x76, 0xb6, 0x2a, 0xcb, 0x70, 0x1d, 0xde, 0xf7, 0xb1, 0xbb, 0x4a, 0xf2,
	0xaa, 0xb6, 0xf0, 0xcd, 0x6b, 0x9e, 0x8e, 0x3b, 0xa7, 0x89, 0xd8, 0x55, 0x52, 0xe5, 0xeb, 0xc2,
	0x72, 0xa8, 0x2b, 0x5f, 0x17, 0x96, 0xfa, 0x38, 0xd7, 0xce, 0xcf, 0x0f, 0x1f, 0x78, 0x75, 0x83,
	0xb3, 0x43, 0x18, 0x44, 0x22, 0xab, 0x72, 0x93, 0xfd, 0x93, 0x40, 0xef, 0x8d, 0xb5, 0x2b, 0x24,
	0x52, 0x54, 0xbe, 0xc4, 0x48, 0x1f, 0xc9, 0xec, 0x39, 0xf4, 0x2b, 0x6d, 0xa5, 0x93, 0x5a, 0x79,
	0x06, 0x07, 0x93, 0x93, 0xd6, 0x68, 0xfa, 0x38, 0x91, 0xab, 0x4b, 0x14, 0x8d, 0x1b, 0xdd, 0xb9,
	0xd0, 0xcb, 0x25, 0x2a, 0x17, 0xe8, 0x8d, 0x90, 0xb8, 0x2a, 0xdc, 0xed, 0x39, 0x2e, 0xb4, 0xc1,
	0xc0, 0xf3, 0x56, 0xe1, 0x67, 0x4f, 0x2a, 0xfc, 0x05, 0x6f, 0x1d, 0xef, 0xd5, 0x35, 0x47, 0x4c,
	0xb6, 0xc2, 0xdd, 0xbe, 0x5a, 0x38, 0x34, 0x9e, 0xe4, 0x54, 0x34, 0x98, 0x7d, 0x09, 0x07, 0x4b,
	0x74, 0x46, 0x16, 0x35, 0xc7, 0x83, 0xc9, 0x27, 0xef, 0x56, 0x78, 0xf6, 0x7d, 0x6d, 0xff, 0x56,
	0x39, 0xb3, 0x11, 0xd1, 0x9b, 0xee, 0xeb, 0xf2, 0x4b, 0xcb, 0xfb, 0xbe, 0x33, 0x5e, 0x26, 0x5d,
	0x29, 0xd5, 0xb5, 0xe7, 0x39, 0x15, 0x5e, 0xa6, 0x0b, 0x29, 0xbc, 0xb9, 0xd0, 0x73, 0xf4, 0x14,
	0xf7, 0x45, 0x84, 0x34, 0x5c, 0x15, 0xf5, 0x33, 0x70, 0x5c, 0x03, 0x6a, 0x07, 0x1a, 0x13, 0x07,
	0x13, 0x8d, 0xa1, 0x0c, 0x73, 0x69, 0x8b, 0xdc, 0xcc, 0xf9, 0x61, 0x9d, 0x21, 0x40, 0x6a, 0x70,
	0x10, 0x05, 0xe6, 0x56, 0x2b, 0x7e, 0xe4, 0xa3, 0x76, 0x95, 0xa3, 0x17, 0x30, 0x6c, 0x5f, 0xe1,
	0xfd, 0x87, 0x46, 0x95, 0xac, 0xf3, 0x72, 0x85, 0x71, 0xcc, 0x3d, 0x78, 0xd1, 0xfd, 0xaa, 0x93,
	0xfd, 0x06, 0xb0, 0x6d, 0x13, 0xfb, 0x14, 0x7a, 0xd6, 0xe5, 0xa6, 0x7e, 0x23, 0x83, 0xc9, 0x83,
	0x16, 0x55, 0xb3, 0xd0, 0x40, 0x51, 0x7b, 0xb0, 0x67, 0x90, 0xa0, 0x9a, 0x87, 0xae, 0xdf, 0xe9,
	0x48, 0xf6, 0xec, 0x0f, 0xe8, 0x47, 0x05, 0xb5, 0x89, 0x1e, 0x47, 0x6b, 0x8a, 0x1a, 0x4c, 0x0f,
	0xfb, 0xc7, 0xc5, 0xc2, 0xa2, 0xf3, 0x19, 0x13, 0x11, 0x10, 0x31, 0x3e, 0x95, 0x0a, 0xfd, 0xac,
	0x24, 0xc2, 0xcb, 0xe4, 0x7b, 0xa1, 0xcb, 0xd5, 0x52, 0xf9, 0x29, 0x49, 0x44, 0x40, 0x93, 0xbf,
	0xbb, 0x90, 0x12, 0xf1, 0x53, 0xaa, 0x83, 0xbd, 0x84, 0x83, 0xb0, 0x7e, 0xd8, 0x47, 0xad, 0xf2,
	0x76, 0x97, 0xd4, 0xe8, 0xc3, 0xbb, 0x4c, 0x34, 0xf8, 0xf7, 0xd8, 0x37, 0xb0, 0x5f, 0xbf, 0x04,
	0xc6, 0x5b, 0x4e, 0x3b, 0x5b, 0x66, 0xf4, 0xe8, 0x0e, 0x4b, 0x1d, 0xfd, 0x35, 0xf4, 0xfc, 0xde,
	0x62, 0xed, 0x13, 0xda, 0xdb, 0x6e, 0x74, 0xf2, 0xbe, 0xa1, 0x0e, 0xfd, 0x01, 0x86, 0xed, 0xdd,
	0xcb, 0x9e, 0xb4, 0x1c, 0xef, 0xd8, 0xd5, 0xa3, 0xc7, 0xff, 0x69, 0xaf, 0xf3, 0xbd, 0x84, 0xb4,
	0xf9, 0xa6, 0xd8, 0xc7, 0xef, 0x38, 0xb7, 0x3f, 0xaf, 0x51, 0xbb, 0x8f, 0xf1, 0xaf, 0xca, 0xee,
	0x9d, 0x9f, 0xc1, 0x43, 0xa9, 0x5b, 0xa6, 0x4b, 0x53, 0x15, 0x67, 0x45, 0x79, 0x7e, 0xd4, 0xb0,
	0x3d, 0xa3, 0xaf, 0x71, 0xd6, 0xf9, 0xb3, 0x9b, 0x5c, 0x4c, 0x7f, 0xfd, 0x7d, 0xdf, 0xff, 0x94,
	0x5f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xa9, 0x7b, 0x93, 0x3c, 0x07, 0x00, 0x00,
}
