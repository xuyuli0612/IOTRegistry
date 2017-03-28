// Code generated by protoc-gen-go.
// source: IOTRegistry.proto
// DO NOT EDIT!

/*
Package IOTRegistry is a generated protocol buffer package.

It is generated from these files:
	IOTRegistry.proto

It has these top-level messages:
	RegisterThingTX
	CreateRegistrantTX
	RegisterSpecTX
*/
package IOTRegistry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegisterThingTX struct {
	Nonce          []byte   `protobuf:"bytes,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Identities     []string `protobuf:"bytes,2,rep,name=Identities" json:"Identities,omitempty"`
	RegistrantName string   `protobuf:"bytes,3,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	Signature      []byte   `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Data           string   `protobuf:"bytes,5,opt,name=Data" json:"Data,omitempty"`
	Spec           string   `protobuf:"bytes,6,opt,name=Spec" json:"Spec,omitempty"`
}

func (m *RegisterThingTX) Reset()         { *m = RegisterThingTX{} }
func (m *RegisterThingTX) String() string { return proto.CompactTextString(m) }
func (*RegisterThingTX) ProtoMessage()    {}

type CreateRegistrantTX struct {
	RegistrantName string `protobuf:"bytes,1,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	PubKey         []byte `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	Signature      []byte `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Data           string `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
}

func (m *CreateRegistrantTX) Reset()         { *m = CreateRegistrantTX{} }
func (m *CreateRegistrantTX) String() string { return proto.CompactTextString(m) }
func (*CreateRegistrantTX) ProtoMessage()    {}

type RegisterSpecTX struct {
	SpecName       string `protobuf:"bytes,1,opt,name=SpecName" json:"SpecName,omitempty"`
	RegistrantName string `protobuf:"bytes,2,opt,name=RegistrantName" json:"RegistrantName,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Data           string `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *RegisterSpecTX) Reset()         { *m = RegisterSpecTX{} }
func (m *RegisterSpecTX) String() string { return proto.CompactTextString(m) }
func (*RegisterSpecTX) ProtoMessage()    {}
