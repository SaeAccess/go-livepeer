// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/lp_rpc.proto

package server

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

// This request is sent by the broadcaster in `GetTranscoder` to request
// information on which transcoder to use.
type TranscoderRequest struct {
	// ID of the job that the broadcaster needs a transcoder for
	JobId int64 `protobuf:"varint,1,opt,name=jobId" json:"jobId,omitempty"`
	// Broadcaster's signature over the jobId
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranscoderRequest) Reset()         { *m = TranscoderRequest{} }
func (m *TranscoderRequest) String() string { return proto.CompactTextString(m) }
func (*TranscoderRequest) ProtoMessage()    {}
func (*TranscoderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{0}
}
func (m *TranscoderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscoderRequest.Unmarshal(m, b)
}
func (m *TranscoderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscoderRequest.Marshal(b, m, deterministic)
}
func (dst *TranscoderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscoderRequest.Merge(dst, src)
}
func (m *TranscoderRequest) XXX_Size() int {
	return xxx_messageInfo_TranscoderRequest.Size(m)
}
func (m *TranscoderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscoderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranscoderRequest proto.InternalMessageInfo

func (m *TranscoderRequest) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *TranscoderRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// The orchestrator sends this in response to `GetTranscoder`, containing the
// transcoder URI, associated credentials authorizing the broadcaster to
// use the transcoder, and miscellaneous data related to the job.
type TranscoderInfo struct {
	// URI of the transcoder to use for submitting segments
	Transcoder string `protobuf:"bytes,1,opt,name=transcoder" json:"transcoder,omitempty"`
	// Signals the authentication method to expect within `credentials`. This
	// field is opaque to the broadcaster, and should be passed to the transcoder.
	AuthType string `protobuf:"bytes,2,opt,name=authType" json:"authType,omitempty"`
	// Credentials to verify the request has been authorized by an orchestrator.
	// This field is opaque to the broadcaster.
	Credentials string `protobuf:"bytes,3,opt,name=credentials" json:"credentials,omitempty"`
	// Transcoded streamId list to update the master manifest on the broadcaster.
	StreamIds            map[string]string `protobuf:"bytes,16,rep,name=streamIds" json:"streamIds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TranscoderInfo) Reset()         { *m = TranscoderInfo{} }
func (m *TranscoderInfo) String() string { return proto.CompactTextString(m) }
func (*TranscoderInfo) ProtoMessage()    {}
func (*TranscoderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{1}
}
func (m *TranscoderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscoderInfo.Unmarshal(m, b)
}
func (m *TranscoderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscoderInfo.Marshal(b, m, deterministic)
}
func (dst *TranscoderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscoderInfo.Merge(dst, src)
}
func (m *TranscoderInfo) XXX_Size() int {
	return xxx_messageInfo_TranscoderInfo.Size(m)
}
func (m *TranscoderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscoderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TranscoderInfo proto.InternalMessageInfo

func (m *TranscoderInfo) GetTranscoder() string {
	if m != nil {
		return m.Transcoder
	}
	return ""
}

func (m *TranscoderInfo) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

func (m *TranscoderInfo) GetCredentials() string {
	if m != nil {
		return m.Credentials
	}
	return ""
}

func (m *TranscoderInfo) GetStreamIds() map[string]string {
	if m != nil {
		return m.StreamIds
	}
	return nil
}

// AuthToken is sent by the orchestrator and encoded in the `credentials` field
// This record is opaque to the broadcaster and is only relevant between the
// orchestrator and the transcoder.
type AuthToken struct {
	// Signature of the orchestrator over the remaining fields
	Sig                  []byte   `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	JobId                int64    `protobuf:"varint,16,opt,name=jobId" json:"jobId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthToken) Reset()         { *m = AuthToken{} }
func (m *AuthToken) String() string { return proto.CompactTextString(m) }
func (*AuthToken) ProtoMessage()    {}
func (*AuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{2}
}
func (m *AuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthToken.Unmarshal(m, b)
}
func (m *AuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthToken.Marshal(b, m, deterministic)
}
func (dst *AuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthToken.Merge(dst, src)
}
func (m *AuthToken) XXX_Size() int {
	return xxx_messageInfo_AuthToken.Size(m)
}
func (m *AuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_AuthToken proto.InternalMessageInfo

func (m *AuthToken) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *AuthToken) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

// Data included by the broadcaster when submitting a segment for transcoding.
type SegData struct {
	// Sequence number of the segment to be transcoded
	Seq int64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	// Hash of the segment data to be transcoded
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Broadcaster signature for the segment. Corresponds to:
	// broadcaster.sign(streamId | seqNo | dataHash)
	// where streamId is derived from the jobId
	Sig                  []byte   `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegData) Reset()         { *m = SegData{} }
func (m *SegData) String() string { return proto.CompactTextString(m) }
func (*SegData) ProtoMessage()    {}
func (*SegData) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{3}
}
func (m *SegData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegData.Unmarshal(m, b)
}
func (m *SegData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegData.Marshal(b, m, deterministic)
}
func (dst *SegData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegData.Merge(dst, src)
}
func (m *SegData) XXX_Size() int {
	return xxx_messageInfo_SegData.Size(m)
}
func (m *SegData) XXX_DiscardUnknown() {
	xxx_messageInfo_SegData.DiscardUnknown(m)
}

var xxx_messageInfo_SegData proto.InternalMessageInfo

func (m *SegData) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *SegData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SegData) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// Individual transcoded segment data.
type TranscodedSegmentData struct {
	// URL where the transcoded data can be downloaded from
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranscodedSegmentData) Reset()         { *m = TranscodedSegmentData{} }
func (m *TranscodedSegmentData) String() string { return proto.CompactTextString(m) }
func (*TranscodedSegmentData) ProtoMessage()    {}
func (*TranscodedSegmentData) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{4}
}
func (m *TranscodedSegmentData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodedSegmentData.Unmarshal(m, b)
}
func (m *TranscodedSegmentData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodedSegmentData.Marshal(b, m, deterministic)
}
func (dst *TranscodedSegmentData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodedSegmentData.Merge(dst, src)
}
func (m *TranscodedSegmentData) XXX_Size() int {
	return xxx_messageInfo_TranscodedSegmentData.Size(m)
}
func (m *TranscodedSegmentData) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodedSegmentData.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodedSegmentData proto.InternalMessageInfo

func (m *TranscodedSegmentData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A set of transcoded segments following the profiles specified in the jo .
type TranscodeData struct {
	// Transcoded data, in the order specified in the job options
	Segments []*TranscodedSegmentData `protobuf:"bytes,1,rep,name=segments" json:"segments,omitempty"`
	// Signature of the hash of the concatenated hashes
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranscodeData) Reset()         { *m = TranscodeData{} }
func (m *TranscodeData) String() string { return proto.CompactTextString(m) }
func (*TranscodeData) ProtoMessage()    {}
func (*TranscodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{5}
}
func (m *TranscodeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeData.Unmarshal(m, b)
}
func (m *TranscodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeData.Marshal(b, m, deterministic)
}
func (dst *TranscodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeData.Merge(dst, src)
}
func (m *TranscodeData) XXX_Size() int {
	return xxx_messageInfo_TranscodeData.Size(m)
}
func (m *TranscodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeData.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeData proto.InternalMessageInfo

func (m *TranscodeData) GetSegments() []*TranscodedSegmentData {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *TranscodeData) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// Response that a transcoder sends after transcoding a segment.
type TranscodeResult struct {
	// Sequence number of the transcoded results.
	Seq int64 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	// Result of transcoding can be an error, or successful with more info
	//
	// Types that are valid to be assigned to Result:
	//	*TranscodeResult_Error
	//	*TranscodeResult_Data
	Result               isTranscodeResult_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TranscodeResult) Reset()         { *m = TranscodeResult{} }
func (m *TranscodeResult) String() string { return proto.CompactTextString(m) }
func (*TranscodeResult) ProtoMessage()    {}
func (*TranscodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lp_rpc_f4f0a90333c683ed, []int{6}
}
func (m *TranscodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeResult.Unmarshal(m, b)
}
func (m *TranscodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeResult.Marshal(b, m, deterministic)
}
func (dst *TranscodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeResult.Merge(dst, src)
}
func (m *TranscodeResult) XXX_Size() int {
	return xxx_messageInfo_TranscodeResult.Size(m)
}
func (m *TranscodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeResult proto.InternalMessageInfo

type isTranscodeResult_Result interface {
	isTranscodeResult_Result()
}

type TranscodeResult_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,oneof"`
}
type TranscodeResult_Data struct {
	Data *TranscodeData `protobuf:"bytes,3,opt,name=data,oneof"`
}

func (*TranscodeResult_Error) isTranscodeResult_Result() {}
func (*TranscodeResult_Data) isTranscodeResult_Result()  {}

func (m *TranscodeResult) GetResult() isTranscodeResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *TranscodeResult) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *TranscodeResult) GetError() string {
	if x, ok := m.GetResult().(*TranscodeResult_Error); ok {
		return x.Error
	}
	return ""
}

func (m *TranscodeResult) GetData() *TranscodeData {
	if x, ok := m.GetResult().(*TranscodeResult_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TranscodeResult) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TranscodeResult_OneofMarshaler, _TranscodeResult_OneofUnmarshaler, _TranscodeResult_OneofSizer, []interface{}{
		(*TranscodeResult_Error)(nil),
		(*TranscodeResult_Data)(nil),
	}
}

func _TranscodeResult_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TranscodeResult)
	// result
	switch x := m.Result.(type) {
	case *TranscodeResult_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Error)
	case *TranscodeResult_Data:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TranscodeResult.Result has unexpected type %T", x)
	}
	return nil
}

func _TranscodeResult_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TranscodeResult)
	switch tag {
	case 2: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Result = &TranscodeResult_Error{x}
		return true, err
	case 3: // result.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranscodeData)
		err := b.DecodeMessage(msg)
		m.Result = &TranscodeResult_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TranscodeResult_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TranscodeResult)
	// result
	switch x := m.Result.(type) {
	case *TranscodeResult_Error:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Error)))
		n += len(x.Error)
	case *TranscodeResult_Data:
		s := proto.Size(x.Data)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*TranscoderRequest)(nil), "server.TranscoderRequest")
	proto.RegisterType((*TranscoderInfo)(nil), "server.TranscoderInfo")
	proto.RegisterMapType((map[string]string)(nil), "server.TranscoderInfo.StreamIdsEntry")
	proto.RegisterType((*AuthToken)(nil), "server.AuthToken")
	proto.RegisterType((*SegData)(nil), "server.SegData")
	proto.RegisterType((*TranscodedSegmentData)(nil), "server.TranscodedSegmentData")
	proto.RegisterType((*TranscodeData)(nil), "server.TranscodeData")
	proto.RegisterType((*TranscodeResult)(nil), "server.TranscodeResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrchestratorClient interface {
	// Called by the broadcaster to request transcoder info from an orchestrator.
	GetTranscoder(ctx context.Context, in *TranscoderRequest, opts ...grpc.CallOption) (*TranscoderInfo, error)
}

type orchestratorClient struct {
	cc *grpc.ClientConn
}

func NewOrchestratorClient(cc *grpc.ClientConn) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) GetTranscoder(ctx context.Context, in *TranscoderRequest, opts ...grpc.CallOption) (*TranscoderInfo, error) {
	out := new(TranscoderInfo)
	err := c.cc.Invoke(ctx, "/server.Orchestrator/GetTranscoder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
type OrchestratorServer interface {
	// Called by the broadcaster to request transcoder info from an orchestrator.
	GetTranscoder(context.Context, *TranscoderRequest) (*TranscoderInfo, error)
}

func RegisterOrchestratorServer(s *grpc.Server, srv OrchestratorServer) {
	s.RegisterService(&_Orchestrator_serviceDesc, srv)
}

func _Orchestrator_GetTranscoder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscoderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetTranscoder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Orchestrator/GetTranscoder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetTranscoder(ctx, req.(*TranscoderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orchestrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTranscoder",
			Handler:    _Orchestrator_GetTranscoder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/lp_rpc.proto",
}

func init() { proto.RegisterFile("server/lp_rpc.proto", fileDescriptor_lp_rpc_f4f0a90333c683ed) }

var fileDescriptor_lp_rpc_f4f0a90333c683ed = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0xba, 0x0d, 0xf1, 0xa4, 0x2d, 0x61, 0xa1, 0x95, 0x89, 0x04, 0xb2, 0x2c, 0x21,
	0x05, 0x21, 0x19, 0x29, 0xbd, 0xf0, 0x75, 0x69, 0x01, 0xd1, 0x9c, 0x90, 0x36, 0x3d, 0x22, 0xa1,
	0x8d, 0x3d, 0xc4, 0x21, 0x8e, 0xd7, 0x99, 0x5d, 0x47, 0xca, 0xcf, 0xe6, 0x1f, 0x20, 0xaf, 0x3f,
	0xa3, 0xe4, 0x36, 0x33, 0x9e, 0x77, 0xf6, 0x9d, 0x67, 0xd7, 0xf0, 0x5c, 0x21, 0x6d, 0x91, 0xde,
	0x27, 0xd9, 0x6f, 0xca, 0xc2, 0x20, 0x23, 0xa9, 0x25, 0xeb, 0x95, 0x45, 0xff, 0x33, 0x3c, 0x7b,
	0x24, 0x91, 0xaa, 0x50, 0x46, 0x48, 0x1c, 0x37, 0x39, 0x2a, 0xcd, 0x5e, 0xc0, 0xf9, 0x5f, 0x39,
	0x9f, 0x46, 0xae, 0xe5, 0x59, 0x63, 0x9b, 0x97, 0x09, 0x1b, 0x82, 0xad, 0x96, 0x0b, 0xf7, 0xd4,
	0xb3, 0xc6, 0x17, 0xbc, 0x08, 0xfd, 0x7f, 0x16, 0x5c, 0xb5, 0xea, 0x69, 0xfa, 0x47, 0xb2, 0xd7,
	0x00, 0xba, 0xa9, 0x18, 0xbd, 0xc3, 0x3b, 0x15, 0x36, 0x82, 0xbe, 0xc8, 0x75, 0xfc, 0xb8, 0xcb,
	0xd0, 0x4c, 0x72, 0x78, 0x93, 0x33, 0x0f, 0x06, 0x21, 0x61, 0x84, 0xa9, 0x5e, 0x8a, 0x44, 0xb9,
	0xb6, 0xf9, 0xdc, 0x2d, 0xb1, 0xaf, 0xe0, 0x28, 0x4d, 0x28, 0xd6, 0xd3, 0x48, 0xb9, 0x43, 0xcf,
	0x1e, 0x0f, 0x26, 0x6f, 0x82, 0x72, 0x93, 0x60, 0xdf, 0x48, 0x30, 0xab, 0xfb, 0xbe, 0xa7, 0x9a,
	0x76, 0xbc, 0xd5, 0x8d, 0xbe, 0xc0, 0xd5, 0xfe, 0xc7, 0x62, 0xb3, 0x15, 0xee, 0x2a, 0xb7, 0x45,
	0x58, 0x10, 0xd8, 0x8a, 0x24, 0xaf, 0x3d, 0x96, 0xc9, 0xa7, 0xd3, 0x0f, 0x96, 0x7f, 0x0b, 0xce,
	0x5d, 0x61, 0x58, 0xae, 0x30, 0xad, 0x91, 0x58, 0x0d, 0x92, 0x16, 0xdd, 0xb0, 0x83, 0xce, 0xbf,
	0x83, 0x27, 0x33, 0x5c, 0x7c, 0x13, 0x5a, 0x18, 0x09, 0x6e, 0x2a, 0xb2, 0x45, 0xc8, 0x18, 0x9c,
	0xc5, 0x42, 0xc5, 0x15, 0x58, 0x13, 0xd7, 0x83, 0xed, 0x96, 0xf5, 0x5b, 0xb8, 0x6e, 0x36, 0x8c,
	0x66, 0xb8, 0x58, 0x63, 0xaa, 0xeb, 0x81, 0x39, 0x25, 0xb5, 0xf9, 0x9c, 0x12, 0xff, 0x17, 0x5c,
	0x36, 0xad, 0xa6, 0xe5, 0x23, 0xf4, 0x55, 0xa9, 0x50, 0xae, 0x65, 0xa8, 0xbd, 0x3a, 0xa0, 0xd6,
	0x9d, 0xc9, 0x9b, 0xf6, 0x23, 0x97, 0x4e, 0xf0, 0xb4, 0x11, 0x71, 0x54, 0x79, 0xa2, 0x8f, 0xec,
	0x74, 0x03, 0xe7, 0x48, 0x24, 0xa9, 0xe4, 0xf7, 0x70, 0xc2, 0xcb, 0x94, 0xbd, 0x83, 0xb3, 0x48,
	0x68, 0x61, 0x16, 0x1b, 0x4c, 0xae, 0x0f, 0x5c, 0x14, 0xa7, 0x3f, 0x9c, 0x70, 0xd3, 0x74, 0xdf,
	0x87, 0x1e, 0x99, 0x03, 0x26, 0x1c, 0x2e, 0x7e, 0x52, 0x18, 0xa3, 0xd2, 0x24, 0xb4, 0x24, 0x76,
	0x0f, 0x97, 0x3f, 0x50, 0xb7, 0x37, 0xce, 0x5e, 0x1e, 0xbe, 0x82, 0xea, 0x31, 0x8f, 0x6e, 0x8e,
	0x3f, 0x90, 0x79, 0xcf, 0xfc, 0x08, 0xb7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x96, 0xcf,
	0x30, 0x1f, 0x03, 0x00, 0x00,
}
