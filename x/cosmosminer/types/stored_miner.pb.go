// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmosminer/stored_miner.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type StoredMiner struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	LastClaimed   uint64 `protobuf:"varint,2,opt,name=lastClaimed,proto3" json:"lastClaimed,omitempty"`
	Creator       string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	TemplateIndex uint64 `protobuf:"varint,4,opt,name=templateIndex,proto3" json:"templateIndex,omitempty"`
}

func (m *StoredMiner) Reset()         { *m = StoredMiner{} }
func (m *StoredMiner) String() string { return proto.CompactTextString(m) }
func (*StoredMiner) ProtoMessage()    {}
func (*StoredMiner) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3730a938628d7, []int{0}
}
func (m *StoredMiner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredMiner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredMiner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredMiner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredMiner.Merge(m, src)
}
func (m *StoredMiner) XXX_Size() int {
	return m.Size()
}
func (m *StoredMiner) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredMiner.DiscardUnknown(m)
}

var xxx_messageInfo_StoredMiner proto.InternalMessageInfo

func (m *StoredMiner) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredMiner) GetLastClaimed() uint64 {
	if m != nil {
		return m.LastClaimed
	}
	return 0
}

func (m *StoredMiner) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StoredMiner) GetTemplateIndex() uint64 {
	if m != nil {
		return m.TemplateIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*StoredMiner)(nil), "luktechlabs.cosmosminer.cosmosminer.StoredMiner")
}

func init() { proto.RegisterFile("cosmosminer/stored_miner.proto", fileDescriptor_5fe3730a938628d7) }

var fileDescriptor_5fe3730a938628d7 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xcd, 0xcc, 0x4b, 0x2d, 0xd2, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4d, 0x89, 0x07,
	0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x94, 0x73, 0x4a, 0xb3, 0x4b, 0x52, 0x93, 0x33,
	0x72, 0x12, 0x93, 0x8a, 0xf5, 0x90, 0xd4, 0x22, 0xb3, 0x95, 0x5a, 0x19, 0xb9, 0xb8, 0x83, 0xc1,
	0x7a, 0x7d, 0x41, 0x7c, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x81, 0x8b, 0x3b, 0x27, 0xb1, 0xb8, 0xc4, 0x39, 0x27,
	0x31, 0x33, 0x37, 0x35, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0x59, 0x48, 0x48, 0x82,
	0x8b, 0x3d, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x19, 0xac, 0x13, 0xc6, 0x15, 0x52,
	0xe1, 0xe2, 0x2d, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xf5, 0x04, 0x9b, 0xcc, 0x02, 0xd6,
	0x8d, 0x2a, 0xe8, 0x14, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x16,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x1f, 0xe9, 0x82, 0xbc,
	0xa4, 0x0f, 0xf1, 0x86, 0x2e, 0xc4, 0xff, 0x15, 0xfa, 0xc8, 0xa1, 0x51, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0x0e, 0x07, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xaa, 0xc2, 0x5b,
	0x29, 0x01, 0x00, 0x00,
}

func (m *StoredMiner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredMiner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredMiner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TemplateIndex != 0 {
		i = encodeVarintStoredMiner(dAtA, i, uint64(m.TemplateIndex))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStoredMiner(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LastClaimed != 0 {
		i = encodeVarintStoredMiner(dAtA, i, uint64(m.LastClaimed))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredMiner(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredMiner(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredMiner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredMiner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredMiner(uint64(l))
	}
	if m.LastClaimed != 0 {
		n += 1 + sovStoredMiner(uint64(m.LastClaimed))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStoredMiner(uint64(l))
	}
	if m.TemplateIndex != 0 {
		n += 1 + sovStoredMiner(uint64(m.TemplateIndex))
	}
	return n
}

func sovStoredMiner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredMiner(x uint64) (n int) {
	return sovStoredMiner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredMiner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredMiner
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
			return fmt.Errorf("proto: StoredMiner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredMiner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredMiner
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
				return ErrInvalidLengthStoredMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaimed", wireType)
			}
			m.LastClaimed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredMiner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastClaimed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredMiner
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
				return ErrInvalidLengthStoredMiner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredMiner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemplateIndex", wireType)
			}
			m.TemplateIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredMiner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TemplateIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStoredMiner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredMiner
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
func skipStoredMiner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredMiner
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
					return 0, ErrIntOverflowStoredMiner
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
					return 0, ErrIntOverflowStoredMiner
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
				return 0, ErrInvalidLengthStoredMiner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredMiner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredMiner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredMiner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredMiner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredMiner = fmt.Errorf("proto: unexpected end of group")
)