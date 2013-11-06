// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package testpb

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TestMessage struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	City             *string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}

func (m *TestMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TestMessage) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

type TestRequest struct {
	Lower            *string `protobuf:"bytes,1,req,name=lower" json:"lower,omitempty"`
	RepeatCount      *int32  `protobuf:"varint,2,opt,name=repeat_count,def=1" json:"repeat_count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}

const Default_TestRequest_RepeatCount int32 = 1

func (m *TestRequest) GetLower() string {
	if m != nil && m.Lower != nil {
		return *m.Lower
	}
	return ""
}

func (m *TestRequest) GetRepeatCount() int32 {
	if m != nil && m.RepeatCount != nil {
		return *m.RepeatCount
	}
	return Default_TestRequest_RepeatCount
}

type TestResponse struct {
	Value            *string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}

func (m *TestResponse) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CacheStats struct {
	Items            *int64 `protobuf:"varint,1,opt,name=items" json:"items,omitempty"`
	Bytes            *int64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	Gets             *int64 `protobuf:"varint,3,opt,name=gets" json:"gets,omitempty"`
	Hits             *int64 `protobuf:"varint,4,opt,name=hits" json:"hits,omitempty"`
	Evicts           *int64 `protobuf:"varint,5,opt,name=evicts" json:"evicts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CacheStats) Reset()         { *m = CacheStats{} }
func (m *CacheStats) String() string { return proto.CompactTextString(m) }
func (*CacheStats) ProtoMessage()    {}

func (m *CacheStats) GetItems() int64 {
	if m != nil && m.Items != nil {
		return *m.Items
	}
	return 0
}

func (m *CacheStats) GetBytes() int64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *CacheStats) GetGets() int64 {
	if m != nil && m.Gets != nil {
		return *m.Gets
	}
	return 0
}

func (m *CacheStats) GetHits() int64 {
	if m != nil && m.Hits != nil {
		return *m.Hits
	}
	return 0
}

func (m *CacheStats) GetEvicts() int64 {
	if m != nil && m.Evicts != nil {
		return *m.Evicts
	}
	return 0
}

type StatsResponse struct {
	Gets             *int64      `protobuf:"varint,1,opt,name=gets" json:"gets,omitempty"`
	CacheHits        *int64      `protobuf:"varint,12,opt,name=cache_hits" json:"cache_hits,omitempty"`
	Fills            *int64      `protobuf:"varint,2,opt,name=fills" json:"fills,omitempty"`
	TotalAlloc       *uint64     `protobuf:"varint,3,opt,name=total_alloc" json:"total_alloc,omitempty"`
	MainCache        *CacheStats `protobuf:"bytes,4,opt,name=main_cache" json:"main_cache,omitempty"`
	HotCache         *CacheStats `protobuf:"bytes,5,opt,name=hot_cache" json:"hot_cache,omitempty"`
	ServerIn         *int64      `protobuf:"varint,6,opt,name=server_in" json:"server_in,omitempty"`
	Loads            *int64      `protobuf:"varint,8,opt,name=loads" json:"loads,omitempty"`
	PeerLoads        *int64      `protobuf:"varint,9,opt,name=peer_loads" json:"peer_loads,omitempty"`
	PeerErrors       *int64      `protobuf:"varint,10,opt,name=peer_errors" json:"peer_errors,omitempty"`
	LocalLoads       *int64      `protobuf:"varint,11,opt,name=local_loads" json:"local_loads,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}

func (m *StatsResponse) GetGets() int64 {
	if m != nil && m.Gets != nil {
		return *m.Gets
	}
	return 0
}

func (m *StatsResponse) GetCacheHits() int64 {
	if m != nil && m.CacheHits != nil {
		return *m.CacheHits
	}
	return 0
}

func (m *StatsResponse) GetFills() int64 {
	if m != nil && m.Fills != nil {
		return *m.Fills
	}
	return 0
}

func (m *StatsResponse) GetTotalAlloc() uint64 {
	if m != nil && m.TotalAlloc != nil {
		return *m.TotalAlloc
	}
	return 0
}

func (m *StatsResponse) GetMainCache() *CacheStats {
	if m != nil {
		return m.MainCache
	}
	return nil
}

func (m *StatsResponse) GetHotCache() *CacheStats {
	if m != nil {
		return m.HotCache
	}
	return nil
}

func (m *StatsResponse) GetServerIn() int64 {
	if m != nil && m.ServerIn != nil {
		return *m.ServerIn
	}
	return 0
}

func (m *StatsResponse) GetLoads() int64 {
	if m != nil && m.Loads != nil {
		return *m.Loads
	}
	return 0
}

func (m *StatsResponse) GetPeerLoads() int64 {
	if m != nil && m.PeerLoads != nil {
		return *m.PeerLoads
	}
	return 0
}

func (m *StatsResponse) GetPeerErrors() int64 {
	if m != nil && m.PeerErrors != nil {
		return *m.PeerErrors
	}
	return 0
}

func (m *StatsResponse) GetLocalLoads() int64 {
	if m != nil && m.LocalLoads != nil {
		return *m.LocalLoads
	}
	return 0
}

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}

func init() {
}
