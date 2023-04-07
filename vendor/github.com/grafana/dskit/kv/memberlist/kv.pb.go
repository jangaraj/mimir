// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv.proto

package memberlist

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// KV Store is just a series of key-value pairs.
type KeyValueStore struct {
	Pairs []*KeyValuePair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (m *KeyValueStore) Reset()      { *m = KeyValueStore{} }
func (*KeyValueStore) ProtoMessage() {}
func (*KeyValueStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{0}
}
func (m *KeyValueStore) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValueStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValueStore.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValueStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValueStore.Merge(m, src)
}
func (m *KeyValueStore) XXX_Size() int {
	return m.Size()
}
func (m *KeyValueStore) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValueStore.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValueStore proto.InternalMessageInfo

func (m *KeyValueStore) GetPairs() []*KeyValuePair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

// Single Key-Value pair. Key must be non-empty.
type KeyValuePair struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// ID of the codec used to write the value
	Codec string `protobuf:"bytes,3,opt,name=codec,proto3" json:"codec,omitempty"`
}

func (m *KeyValuePair) Reset()      { *m = KeyValuePair{} }
func (*KeyValuePair) ProtoMessage() {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{1}
}
func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return m.Size()
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValuePair) GetCodec() string {
	if m != nil {
		return m.Codec
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyValueStore)(nil), "memberlist.KeyValueStore")
	proto.RegisterType((*KeyValuePair)(nil), "memberlist.KeyValuePair")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_2216fe83c9c12408) }

var fileDescriptor_2216fe83c9c12408 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2e, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0xca, 0xc9, 0x2c, 0x2e,
	0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f,
	0xcf, 0xd7, 0x07, 0x2b, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x55, 0xc9,
	0x9e, 0x8b, 0xd7, 0x3b, 0xb5, 0x32, 0x2c, 0x31, 0xa7, 0x34, 0x35, 0xb8, 0x24, 0xbf, 0x28, 0x55,
	0x48, 0x8f, 0x8b, 0xb5, 0x20, 0x31, 0xb3, 0xa8, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48,
	0x42, 0x0f, 0x61, 0xb6, 0x1e, 0x4c, 0x65, 0x40, 0x62, 0x66, 0x51, 0x10, 0x44, 0x99, 0x92, 0x0f,
	0x17, 0x0f, 0xb2, 0xb0, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x06, 0x92, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x09, 0x82, 0x70, 0x40, 0xa2, 0xc9, 0xf9, 0x29, 0xa9, 0xc9, 0x12, 0xcc, 0x60, 0x95, 0x10, 0x8e,
	0x93, 0xc9, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0,
	0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0x5f, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x7a, 0x22, 0xdf, 0xec, 0x12, 0x01, 0x00, 0x00,
}

func (this *KeyValueStore) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KeyValueStore)
	if !ok {
		that2, ok := that.(KeyValueStore)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Pairs) != len(that1.Pairs) {
		return false
	}
	for i := range this.Pairs {
		if !this.Pairs[i].Equal(that1.Pairs[i]) {
			return false
		}
	}
	return true
}
func (this *KeyValuePair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KeyValuePair)
	if !ok {
		that2, ok := that.(KeyValuePair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	if this.Codec != that1.Codec {
		return false
	}
	return true
}
func (this *KeyValueStore) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&memberlist.KeyValueStore{")
	if this.Pairs != nil {
		s = append(s, "Pairs: "+fmt.Sprintf("%#v", this.Pairs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *KeyValuePair) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&memberlist.KeyValuePair{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Codec: "+fmt.Sprintf("%#v", this.Codec)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringKv(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *KeyValueStore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValueStore) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValueStore) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pairs) > 0 {
		for iNdEx := len(m.Pairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKv(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *KeyValuePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyValuePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyValuePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codec) > 0 {
		i -= len(m.Codec)
		copy(dAtA[i:], m.Codec)
		i = encodeVarintKv(dAtA, i, uint64(len(m.Codec)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintKv(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintKv(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKv(dAtA []byte, offset int, v uint64) int {
	offset -= sovKv(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyValueStore) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pairs) > 0 {
		for _, e := range m.Pairs {
			l = e.Size()
			n += 1 + l + sovKv(uint64(l))
		}
	}
	return n
}

func (m *KeyValuePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	l = len(m.Codec)
	if l > 0 {
		n += 1 + l + sovKv(uint64(l))
	}
	return n
}

func sovKv(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKv(x uint64) (n int) {
	return sovKv(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *KeyValueStore) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPairs := "[]*KeyValuePair{"
	for _, f := range this.Pairs {
		repeatedStringForPairs += strings.Replace(f.String(), "KeyValuePair", "KeyValuePair", 1) + ","
	}
	repeatedStringForPairs += "}"
	s := strings.Join([]string{`&KeyValueStore{`,
		`Pairs:` + repeatedStringForPairs + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyValuePair) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyValuePair{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Codec:` + fmt.Sprintf("%v", this.Codec) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringKv(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *KeyValueStore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
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
			return fmt.Errorf("proto: KeyValueStore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValueStore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pairs = append(m.Pairs, &KeyValuePair{})
			if err := m.Pairs[len(m.Pairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKv
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
func (m *KeyValuePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKv
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
			return fmt.Errorf("proto: KeyValuePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValuePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKv
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
				return ErrInvalidLengthKv
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKv
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKv(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKv
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKv
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
func skipKv(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKv
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
					return 0, ErrIntOverflowKv
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKv
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
				return 0, ErrInvalidLengthKv
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthKv
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKv
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipKv(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthKv
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthKv = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKv   = fmt.Errorf("proto: integer overflow")
)
