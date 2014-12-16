// Code generated by protoc-gen-go.
// source: snapshot_request.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SnapshotRequest struct {
	LeaderName       *string `protobuf:"bytes,1,req" json:"LeaderName,omitempty"`
	LastIndex        *uint64 `protobuf:"varint,2,req" json:"LastIndex,omitempty"`
	LastTerm         *uint64 `protobuf:"varint,3,req" json:"LastTerm,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SnapshotRequest) Reset()         { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()    {}

func (m *SnapshotRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *SnapshotRequest) GetLastIndex() uint64 {
	if m != nil && m.LastIndex != nil {
		return *m.LastIndex
	}
	return 0
}

func (m *SnapshotRequest) GetLastTerm() uint64 {
	if m != nil && m.LastTerm != nil {
		return *m.LastTerm
	}
	return 0
}

func init() {
}
