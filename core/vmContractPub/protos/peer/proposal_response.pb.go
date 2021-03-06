// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proposal_response.proto

package peer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement          *Endorsement `protobuf:"bytes,6,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProposalResponse) Reset()         { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()    {}
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f373cd368f08c5, []int{0}
}

func (m *ProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponse.Unmarshal(m, b)
}
func (m *ProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponse.Marshal(b, m, deterministic)
}
func (m *ProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponse.Merge(m, src)
}
func (m *ProposalResponse) XXX_Size() int {
	return xxx_messageInfo_ProposalResponse.Size(m)
}
func (m *ProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponse proto.InternalMessageInfo

func (m *ProposalResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProposalResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f373cd368f08c5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte). However this implies that the hash can only be verified
	// if the entire proposal message is available when ProposalResponsePayload is
	// included in a transaction or stored in the ledger. For confidentiality
	// reasons, with chaincodes it might be undesirable to store the proposal
	// payload in the ledger.  If the type is CHAINCODE, this is handled by
	// separating the proposal's header and
	// the payload: the header is always hashed in its entirety whereas the
	// payload can either be hashed fully, or only its hash may be hashed, or
	// nothing from the payload can be hashed. The PayloadVisibility field in the
	// Header's extension controls to which extent the proposal payload is
	// "visible" in the sense that was just explained.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension            []byte   `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalResponsePayload) Reset()         { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()    {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f373cd368f08c5, []int{2}
}

func (m *ProposalResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponsePayload.Unmarshal(m, b)
}
func (m *ProposalResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponsePayload.Marshal(b, m, deterministic)
}
func (m *ProposalResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponsePayload.Merge(m, src)
}
func (m *ProposalResponsePayload) XXX_Size() int {
	return xxx_messageInfo_ProposalResponsePayload.Size(m)
}
func (m *ProposalResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponsePayload proto.InternalMessageInfo

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endorsement) Reset()         { *m = Endorsement{} }
func (m *Endorsement) String() string { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()    {}
func (*Endorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f373cd368f08c5, []int{3}
}

func (m *Endorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endorsement.Unmarshal(m, b)
}
func (m *Endorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endorsement.Marshal(b, m, deterministic)
}
func (m *Endorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endorsement.Merge(m, src)
}
func (m *Endorsement) XXX_Size() int {
	return xxx_messageInfo_Endorsement.Size(m)
}
func (m *Endorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_Endorsement.DiscardUnknown(m)
}

var xxx_messageInfo_Endorsement proto.InternalMessageInfo

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*ProposalResponsePayload)(nil), "protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
}

func init() { proto.RegisterFile("proposal_response.proto", fileDescriptor_b4f373cd368f08c5) }

var fileDescriptor_b4f373cd368f08c5 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0xd9, 0x92, 0x25, 0x72, 0x06, 0xc1, 0x83, 0xc5, 0x84, 0xc1, 0x82, 0x77, 0xc9, 0x61,
	0xb3, 0x61, 0xa5, 0xd0, 0x63, 0x49, 0x29, 0xed, 0x31, 0x88, 0xd2, 0x43, 0x29, 0x14, 0xd9, 0x79,
	0xb5, 0x4d, 0x6d, 0x49, 0xe8, 0xc9, 0xa1, 0xfd, 0x83, 0xfb, 0x7f, 0x14, 0xcb, 0x92, 0x9d, 0x96,
	0x9e, 0xcc, 0xf7, 0xfc, 0xe9, 0xfb, 0xf1, 0x24, 0xb2, 0x94, 0x4a, 0x48, 0x81, 0xac, 0x7a, 0x50,
	0x80, 0x52, 0x70, 0x84, 0x58, 0x2a, 0xa1, 0x45, 0x30, 0x31, 0x1f, 0x5c, 0xfd, 0xce, 0x85, 0xc8,
	0x2b, 0x48, 0x0c, 0x4c, 0x9b, 0xc7, 0x44, 0x97, 0x35, 0xa0, 0x66, 0xb5, 0xec, 0x88, 0xd1, 0xab,
	0x47, 0x16, 0x3b, 0x2b, 0x42, 0xad, 0x46, 0x10, 0x92, 0x6f, 0x07, 0x50, 0x58, 0x0a, 0x1e, 0x7a,
	0x6b, 0x6f, 0x33, 0xa6, 0x0e, 0x06, 0x67, 0x64, 0xd6, 0x2b, 0x84, 0xa3, 0xb5, 0xb7, 0xf1, 0xff,
	0xaf, 0xe2, 0xce, 0x23, 0x76, 0x1e, 0xf1, 0x8d, 0x63, 0xd0, 0x81, 0x1c, 0xfc, 0x25, 0x53, 0x97,
	0x31, 0xfc, 0x6a, 0x0e, 0x2e, 0xba, 0x13, 0x18, 0x3b, 0x5f, 0xda, 0x33, 0xda, 0x04, 0x92, 0xbd,
	0x54, 0x82, 0xed, 0xc3, 0xf1, 0xda, 0xdb, 0xcc, 0xa9, 0x83, 0xc1, 0x29, 0xf1, 0x81, 0xef, 0x85,
	0x42, 0xa8, 0x81, 0xeb, 0x70, 0x62, 0xa4, 0x7e, 0x38, 0xa9, 0xcb, 0xe1, 0x17, 0x3d, 0xe6, 0x45,
	0xb7, 0x64, 0xda, 0xd7, 0xfb, 0x49, 0x26, 0xa8, 0x99, 0x6e, 0xd0, 0xb6, 0xb3, 0xa8, 0x35, 0xad,
	0x01, 0x91, 0xe5, 0x60, 0xaa, 0xcd, 0xa8, 0x83, 0xc7, 0x71, 0xbe, 0xbc, 0x8b, 0x13, 0xdd, 0x93,
	0xe5, 0xc7, 0xf5, 0xed, 0x6c, 0xd2, 0x3f, 0xe4, 0x7b, 0x7f, 0x3d, 0x05, 0xc3, 0xc2, 0xb8, 0xcd,
	0xe9, 0xdc, 0x0d, 0xaf, 0x19, 0x16, 0xc1, 0x2f, 0x32, 0x83, 0x67, 0x0d, 0xdc, 0x2c, 0x7b, 0x64,
	0x08, 0xc3, 0x20, 0xba, 0x22, 0xfe, 0x51, 0xa3, 0x60, 0x45, 0xa6, 0xb6, 0x93, 0xb2, 0x62, 0x3d,
	0x6e, 0x85, 0xb0, 0xcc, 0x39, 0xd3, 0x8d, 0x02, 0x27, 0xd4, 0x0f, 0xb6, 0x05, 0xf1, 0xed, 0x86,
	0x24, 0x80, 0xda, 0x7e, 0x92, 0x39, 0x7b, 0x62, 0x39, 0xdc, 0x9d, 0xe7, 0xa5, 0x2e, 0x9a, 0x34,
	0xce, 0x44, 0x9d, 0x48, 0x56, 0x55, 0xa0, 0x05, 0x87, 0x24, 0x17, 0xff, 0x06, 0x90, 0x09, 0x05,
	0xc9, 0xa1, 0xbe, 0x10, 0x5c, 0x2b, 0x96, 0xe9, 0x5d, 0x93, 0x76, 0xaf, 0x0b, 0x93, 0x56, 0x3a,
	0xed, 0x5e, 0xde, 0xc9, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x2d, 0x5a, 0x00, 0x9b, 0x02,
	0x00, 0x00,
}
