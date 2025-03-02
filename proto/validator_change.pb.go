// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator_change.proto

package libra

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// This is used to prove validator changes.  When a validator is changing, it
// triggers an event on /validator_change_account/events/sent.  To tell the
// client about validator changes, we query
// /validator_change_account/events/sent to get all versions that contain
// validator changes after the version that we are trying to update from. For
// each of these versions, the old validator set would have signed the ledger
// info at that version.  The client needs this as well as the event results +
// proof.  The client can then verify that these events were under the current
// tree and that the changes were signed by the old validators (and that the
// events correctly show which validators are the new validators).
//
// This message represents a single validator change event and the proof that
// corresponds to it
type ValidatorChangeEventWithProof struct {
	LedgerInfoWithSigs   *LedgerInfoWithSignatures `protobuf:"bytes,1,opt,name=ledger_info_with_sigs,json=ledgerInfoWithSigs,proto3" json:"ledger_info_with_sigs,omitempty"`
	EventWithProof       *EventWithProof           `protobuf:"bytes,2,opt,name=event_with_proof,json=eventWithProof,proto3" json:"event_with_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ValidatorChangeEventWithProof) Reset()         { *m = ValidatorChangeEventWithProof{} }
func (m *ValidatorChangeEventWithProof) String() string { return proto.CompactTextString(m) }
func (*ValidatorChangeEventWithProof) ProtoMessage()    {}
func (*ValidatorChangeEventWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_484a7446a4b3dd15, []int{0}
}

func (m *ValidatorChangeEventWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Unmarshal(m, b)
}
func (m *ValidatorChangeEventWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Marshal(b, m, deterministic)
}
func (m *ValidatorChangeEventWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorChangeEventWithProof.Merge(m, src)
}
func (m *ValidatorChangeEventWithProof) XXX_Size() int {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Size(m)
}
func (m *ValidatorChangeEventWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorChangeEventWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorChangeEventWithProof proto.InternalMessageInfo

func (m *ValidatorChangeEventWithProof) GetLedgerInfoWithSigs() *LedgerInfoWithSignatures {
	if m != nil {
		return m.LedgerInfoWithSigs
	}
	return nil
}

func (m *ValidatorChangeEventWithProof) GetEventWithProof() *EventWithProof {
	if m != nil {
		return m.EventWithProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorChangeEventWithProof)(nil), "types.ValidatorChangeEventWithProof")
}

func init() { proto.RegisterFile("validator_change.proto", fileDescriptor_484a7446a4b3dd15) }

var fileDescriptor_484a7446a4b3dd15 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x8a, 0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xe2, 0x49, 0x2d, 0x4b, 0xcd, 0x2b,
	0x29, 0x86, 0x08, 0x4a, 0x09, 0xe6, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xc5, 0x67, 0xe6, 0xa5, 0xe5,
	0x43, 0x84, 0x94, 0xb6, 0x30, 0x72, 0xc9, 0x86, 0xc1, 0x8c, 0x70, 0x06, 0x9b, 0xe0, 0x0a, 0xd2,
	0x12, 0x9e, 0x59, 0x92, 0x11, 0x50, 0x94, 0x9f, 0x9f, 0x26, 0x14, 0xc4, 0x25, 0x8a, 0xa4, 0x2d,
	0xbe, 0x3c, 0xb3, 0x24, 0x23, 0xbe, 0x38, 0x33, 0xbd, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb,
	0x48, 0x5e, 0x0f, 0x6c, 0x93, 0x9e, 0x0f, 0x58, 0x8d, 0x67, 0x5e, 0x5a, 0x3e, 0x48, 0x6b, 0x70,
	0x66, 0x7a, 0x5e, 0x62, 0x49, 0x69, 0x51, 0x6a, 0x71, 0x90, 0x50, 0x0e, 0xba, 0x4c, 0xb1, 0x90,
	0x3d, 0x97, 0x00, 0xd8, 0x61, 0x10, 0xd3, 0x0a, 0x40, 0xf6, 0x48, 0x30, 0x81, 0x8d, 0x13, 0x85,
	0x1a, 0x87, 0xea, 0x88, 0x20, 0xbe, 0x54, 0x14, 0x7e, 0x12, 0x1b, 0xd8, 0xf5, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x54, 0x40, 0x94, 0xff, 0x00, 0x00, 0x00,
}
