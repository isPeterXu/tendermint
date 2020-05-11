// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/privval/types.proto

package privval

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	keys "github.com/tendermint/tendermint/proto/crypto/keys"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// FilePVKey stores the immutable part of PrivValidator.
type FilePVKey struct {
	Address  []byte          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey   keys.PublicKey  `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	PrivKey  keys.PrivateKey `protobuf:"bytes,3,opt,name=priv_key,json=privKey,proto3" json:"priv_key"`
	FilePath string          `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (m *FilePVKey) Reset()         { *m = FilePVKey{} }
func (m *FilePVKey) String() string { return proto.CompactTextString(m) }
func (*FilePVKey) ProtoMessage()    {}
func (*FilePVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{0}
}
func (m *FilePVKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilePVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilePVKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilePVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVKey.Merge(m, src)
}
func (m *FilePVKey) XXX_Size() int {
	return m.Size()
}
func (m *FilePVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVKey.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVKey proto.InternalMessageInfo

func (m *FilePVKey) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *FilePVKey) GetPubKey() keys.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return keys.PublicKey{}
}

func (m *FilePVKey) GetPrivKey() keys.PrivateKey {
	if m != nil {
		return m.PrivKey
	}
	return keys.PrivateKey{}
}

func (m *FilePVKey) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

// FilePVLastSignState stores the mutable part of PrivValidator.
type FilePVLastSignState struct {
	Height    int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round     int32  `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Step      int32  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	SignBytes []byte `protobuf:"bytes,5,opt,name=sign_bytes,json=signBytes,proto3" json:"sign_bytes,omitempty"`
	FilePath  string `protobuf:"bytes,6,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (m *FilePVLastSignState) Reset()         { *m = FilePVLastSignState{} }
func (m *FilePVLastSignState) String() string { return proto.CompactTextString(m) }
func (*FilePVLastSignState) ProtoMessage()    {}
func (*FilePVLastSignState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d74c406df3ad93, []int{1}
}
func (m *FilePVLastSignState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilePVLastSignState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilePVLastSignState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilePVLastSignState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePVLastSignState.Merge(m, src)
}
func (m *FilePVLastSignState) XXX_Size() int {
	return m.Size()
}
func (m *FilePVLastSignState) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePVLastSignState.DiscardUnknown(m)
}

var xxx_messageInfo_FilePVLastSignState proto.InternalMessageInfo

func (m *FilePVLastSignState) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *FilePVLastSignState) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *FilePVLastSignState) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *FilePVLastSignState) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FilePVLastSignState) GetSignBytes() []byte {
	if m != nil {
		return m.SignBytes
	}
	return nil
}

func (m *FilePVLastSignState) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func init() {
	proto.RegisterType((*FilePVKey)(nil), "tendermint.proto.privval.FilePVKey")
	proto.RegisterType((*FilePVLastSignState)(nil), "tendermint.proto.privval.FilePVLastSignState")
}

func init() { proto.RegisterFile("proto/privval/types.proto", fileDescriptor_a9d74c406df3ad93) }

var fileDescriptor_a9d74c406df3ad93 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6e, 0xd4, 0x30,
	0x14, 0x8e, 0xe9, 0x24, 0xd3, 0x98, 0xae, 0x0c, 0x42, 0xa1, 0xd0, 0x30, 0xea, 0x02, 0xb2, 0x4a,
	0x24, 0xb8, 0xc1, 0x2c, 0x2a, 0x50, 0x59, 0x8c, 0x52, 0x89, 0x05, 0x9b, 0xc8, 0x99, 0x3c, 0x12,
	0xab, 0x69, 0x62, 0xd9, 0x2f, 0x23, 0xf9, 0x16, 0x5c, 0x85, 0x5b, 0x74, 0xd9, 0x0d, 0x12, 0x2b,
	0x84, 0x66, 0x2e, 0x82, 0xec, 0x0c, 0xca, 0x8c, 0x58, 0x74, 0xf7, 0xbe, 0xcf, 0xcf, 0xdf, 0x8f,
	0x65, 0xfa, 0x52, 0xaa, 0x1e, 0xfb, 0x4c, 0x2a, 0xb1, 0xd9, 0xf0, 0x36, 0x43, 0x23, 0x41, 0xa7,
	0x8e, 0x63, 0x11, 0x42, 0x57, 0x81, 0xba, 0x13, 0x1d, 0x8e, 0x4c, 0xba, 0xdf, 0x3a, 0x7f, 0x8b,
	0x8d, 0x50, 0x55, 0x21, 0xb9, 0x42, 0x93, 0x8d, 0x02, 0x75, 0x5f, 0xf7, 0xd3, 0x34, 0xee, 0x9f,
	0x5f, 0x8c, 0xcc, 0x5a, 0x19, 0x89, 0x7d, 0x76, 0x0b, 0x46, 0x1f, 0x1a, 0x5c, 0xfe, 0x24, 0x34,
	0xbc, 0x12, 0x2d, 0xac, 0xbe, 0x5c, 0x83, 0x61, 0x11, 0x9d, 0xf3, 0xaa, 0x52, 0xa0, 0x75, 0x44,
	0x16, 0x24, 0x39, 0xcb, 0xff, 0x41, 0x76, 0x45, 0xe7, 0x72, 0x28, 0x8b, 0x5b, 0x30, 0xd1, 0x93,
	0x05, 0x49, 0x9e, 0xbe, 0x7f, 0x97, 0xfe, 0x17, 0x6d, 0xf4, 0x48, 0xad, 0x47, 0xba, 0x1a, 0xca,
	0x56, 0xac, 0xaf, 0xc1, 0x2c, 0x67, 0xf7, 0xbf, 0xdf, 0x78, 0x79, 0x20, 0x87, 0xd2, 0x3a, 0x7c,
	0xa2, 0xa7, 0xb6, 0x81, 0x13, 0x3a, 0x71, 0x42, 0xc9, 0x23, 0x42, 0x4a, 0x6c, 0x38, 0xc2, 0xa4,
	0x34, 0xb7, 0xf7, 0xad, 0xd4, 0x2b, 0x1a, 0x7e, 0x13, 0x2d, 0x14, 0x92, 0x63, 0x13, 0xcd, 0x16,
	0x24, 0x09, 0xf3, 0x53, 0x4b, 0xac, 0x38, 0x36, 0x97, 0x3f, 0x08, 0x7d, 0x36, 0xf6, 0xfa, 0xcc,
	0x35, 0xde, 0x88, 0xba, 0xbb, 0x41, 0x8e, 0xc0, 0x5e, 0xd0, 0xa0, 0x01, 0x51, 0x37, 0xe8, 0x0a,
	0x9e, 0xe4, 0x7b, 0xc4, 0x9e, 0x53, 0x5f, 0xf5, 0x43, 0x57, 0xb9, 0x76, 0x7e, 0x3e, 0x02, 0xc6,
	0xe8, 0x4c, 0x23, 0x48, 0x97, 0xd4, 0xcf, 0xdd, 0xcc, 0x5e, 0xd3, 0x50, 0x8b, 0xba, 0xe3, 0x38,
	0x28, 0x70, 0xb6, 0x67, 0xf9, 0x44, 0xb0, 0x0b, 0x4a, 0x2d, 0x28, 0x4a, 0x83, 0xa0, 0x23, 0x7f,
	0x3a, 0x5e, 0x5a, 0xe2, 0x38, 0x73, 0x70, 0x9c, 0x79, 0xf9, 0xf1, 0x7e, 0x1b, 0x93, 0x87, 0x6d,
	0x4c, 0xfe, 0x6c, 0x63, 0xf2, 0x7d, 0x17, 0x7b, 0x0f, 0xbb, 0xd8, 0xfb, 0xb5, 0x8b, 0xbd, 0xaf,
	0x69, 0x2d, 0xb0, 0x19, 0xca, 0x74, 0xdd, 0xdf, 0x65, 0xd3, 0x6b, 0x1d, 0x8e, 0x47, 0x5f, 0xa8,
	0x0c, 0x1c, 0xfc, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xac, 0x70, 0xf0, 0x5a, 0x02, 0x00,
	0x00,
}

func (m *FilePVKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilePVKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilePVKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.PrivKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FilePVLastSignState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilePVLastSignState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilePVLastSignState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SignBytes) > 0 {
		i -= len(m.SignBytes)
		copy(dAtA[i:], m.SignBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SignBytes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if m.Step != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilePVKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.PubKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.PrivKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *FilePVLastSignState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	if m.Step != 0 {
		n += 1 + sovTypes(uint64(m.Step))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SignBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilePVKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilePVKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilePVKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrivKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FilePVLastSignState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilePVLastSignState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilePVLastSignState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignBytes = append(m.SignBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.SignBytes == nil {
				m.SignBytes = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
