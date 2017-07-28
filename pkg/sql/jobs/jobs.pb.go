// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/sql/jobs/jobs.proto
// DO NOT EDIT!

/*
	Package jobs is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/sql/jobs/jobs.proto

	It has these top-level messages:
		BackupDetails
		RestoreDetails
		SchemaChangeDetails
		Payload
*/
package jobs

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_sql_sqlbase "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BackupDetails struct {
}

func (m *BackupDetails) Reset()                    { *m = BackupDetails{} }
func (m *BackupDetails) String() string            { return proto.CompactTextString(m) }
func (*BackupDetails) ProtoMessage()               {}
func (*BackupDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{0} }

type RestoreDetails struct {
	LowWaterMark []byte `protobuf:"bytes,1,opt,name=low_water_mark,json=lowWaterMark,proto3" json:"low_water_mark,omitempty"`
}

func (m *RestoreDetails) Reset()                    { *m = RestoreDetails{} }
func (m *RestoreDetails) String() string            { return proto.CompactTextString(m) }
func (*RestoreDetails) ProtoMessage()               {}
func (*RestoreDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{1} }

type SchemaChangeDetails struct {
	ReadAsOf cockroach_util_hlc.Timestamp `protobuf:"bytes,1,opt,name=read_as_of,json=readAsOf" json:"read_as_of"`
}

func (m *SchemaChangeDetails) Reset()                    { *m = SchemaChangeDetails{} }
func (m *SchemaChangeDetails) String() string            { return proto.CompactTextString(m) }
func (*SchemaChangeDetails) ProtoMessage()               {}
func (*SchemaChangeDetails) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{2} }

type Payload struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// For consistency with the SQL timestamp type, which has microsecond
	// precision, we avoid the timestamp.Timestamp WKT, which has nanosecond
	// precision, and use microsecond integers directly.
	StartedMicros     int64                                                 `protobuf:"varint,3,opt,name=started_micros,json=startedMicros,proto3" json:"started_micros,omitempty"`
	FinishedMicros    int64                                                 `protobuf:"varint,4,opt,name=finished_micros,json=finishedMicros,proto3" json:"finished_micros,omitempty"`
	ModifiedMicros    int64                                                 `protobuf:"varint,5,opt,name=modified_micros,json=modifiedMicros,proto3" json:"modified_micros,omitempty"`
	DescriptorIDs     []github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID `protobuf:"varint,6,rep,packed,name=descriptor_ids,json=descriptorIds,casttype=github.com/cockroachdb/cockroach/pkg/sql/sqlbase.ID" json:"descriptor_ids,omitempty"`
	FractionCompleted float32                                               `protobuf:"fixed32,7,opt,name=fraction_completed,json=fractionCompleted,proto3" json:"fraction_completed,omitempty"`
	Error             string                                                `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are valid to be assigned to Details:
	//	*Payload_Backup
	//	*Payload_Restore
	//	*Payload_SchemaChange
	Details isPayload_Details `protobuf_oneof:"details"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorJobs, []int{3} }

type isPayload_Details interface {
	isPayload_Details()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Payload_Backup struct {
	Backup *BackupDetails `protobuf:"bytes,10,opt,name=backup,oneof"`
}
type Payload_Restore struct {
	Restore *RestoreDetails `protobuf:"bytes,11,opt,name=restore,oneof"`
}
type Payload_SchemaChange struct {
	SchemaChange *SchemaChangeDetails `protobuf:"bytes,12,opt,name=schemaChange,oneof"`
}

func (*Payload_Backup) isPayload_Details()       {}
func (*Payload_Restore) isPayload_Details()      {}
func (*Payload_SchemaChange) isPayload_Details() {}

func (m *Payload) GetDetails() isPayload_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Payload) GetBackup() *BackupDetails {
	if x, ok := m.GetDetails().(*Payload_Backup); ok {
		return x.Backup
	}
	return nil
}

func (m *Payload) GetRestore() *RestoreDetails {
	if x, ok := m.GetDetails().(*Payload_Restore); ok {
		return x.Restore
	}
	return nil
}

func (m *Payload) GetSchemaChange() *SchemaChangeDetails {
	if x, ok := m.GetDetails().(*Payload_SchemaChange); ok {
		return x.SchemaChange
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Payload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Payload_OneofMarshaler, _Payload_OneofUnmarshaler, _Payload_OneofSizer, []interface{}{
		(*Payload_Backup)(nil),
		(*Payload_Restore)(nil),
		(*Payload_SchemaChange)(nil),
	}
}

func _Payload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Payload)
	// details
	switch x := m.Details.(type) {
	case *Payload_Backup:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Backup); err != nil {
			return err
		}
	case *Payload_Restore:
		_ = b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Restore); err != nil {
			return err
		}
	case *Payload_SchemaChange:
		_ = b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SchemaChange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Payload.Details has unexpected type %T", x)
	}
	return nil
}

func _Payload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Payload)
	switch tag {
	case 10: // details.backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BackupDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_Backup{msg}
		return true, err
	case 11: // details.restore
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RestoreDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_Restore{msg}
		return true, err
	case 12: // details.schemaChange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SchemaChangeDetails)
		err := b.DecodeMessage(msg)
		m.Details = &Payload_SchemaChange{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Payload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Payload)
	// details
	switch x := m.Details.(type) {
	case *Payload_Backup:
		s := proto.Size(x.Backup)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Payload_Restore:
		s := proto.Size(x.Restore)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Payload_SchemaChange:
		s := proto.Size(x.SchemaChange)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BackupDetails)(nil), "cockroach.sql.jobs.BackupDetails")
	proto.RegisterType((*RestoreDetails)(nil), "cockroach.sql.jobs.RestoreDetails")
	proto.RegisterType((*SchemaChangeDetails)(nil), "cockroach.sql.jobs.SchemaChangeDetails")
	proto.RegisterType((*Payload)(nil), "cockroach.sql.jobs.Payload")
}
func (m *BackupDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *RestoreDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RestoreDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LowWaterMark) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.LowWaterMark)))
		i += copy(dAtA[i:], m.LowWaterMark)
	}
	return i, nil
}

func (m *SchemaChangeDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaChangeDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintJobs(dAtA, i, uint64(m.ReadAsOf.Size()))
	n1, err := m.ReadAsOf.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.StartedMicros != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		dAtA3 := make([]byte, len(m.DescriptorIDs)*10)
		var j2 int
		for _, num := range m.DescriptorIDs {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x32
		i++
		i = encodeVarintJobs(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if m.FractionCompleted != 0 {
		dAtA[i] = 0x3d
		i++
		i = encodeFixed32Jobs(dAtA, i, uint32(math.Float32bits(float32(m.FractionCompleted))))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJobs(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.Details != nil {
		nn4, err := m.Details.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *Payload_Backup) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Backup != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Backup.Size()))
		n5, err := m.Backup.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Payload_Restore) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Restore != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.Restore.Size()))
		n6, err := m.Restore.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *Payload_SchemaChange) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.SchemaChange != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintJobs(dAtA, i, uint64(m.SchemaChange.Size()))
		n7, err := m.SchemaChange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func encodeFixed64Jobs(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Jobs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintJobs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BackupDetails) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *RestoreDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.LowWaterMark)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}

func (m *SchemaChangeDetails) Size() (n int) {
	var l int
	_ = l
	l = m.ReadAsOf.Size()
	n += 1 + l + sovJobs(uint64(l))
	return n
}

func (m *Payload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.StartedMicros != 0 {
		n += 1 + sovJobs(uint64(m.StartedMicros))
	}
	if m.FinishedMicros != 0 {
		n += 1 + sovJobs(uint64(m.FinishedMicros))
	}
	if m.ModifiedMicros != 0 {
		n += 1 + sovJobs(uint64(m.ModifiedMicros))
	}
	if len(m.DescriptorIDs) > 0 {
		l = 0
		for _, e := range m.DescriptorIDs {
			l += sovJobs(uint64(e))
		}
		n += 1 + sovJobs(uint64(l)) + l
	}
	if m.FractionCompleted != 0 {
		n += 5
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovJobs(uint64(l))
	}
	if m.Details != nil {
		n += m.Details.Size()
	}
	return n
}

func (m *Payload_Backup) Size() (n int) {
	var l int
	_ = l
	if m.Backup != nil {
		l = m.Backup.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}
func (m *Payload_Restore) Size() (n int) {
	var l int
	_ = l
	if m.Restore != nil {
		l = m.Restore.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}
func (m *Payload_SchemaChange) Size() (n int) {
	var l int
	_ = l
	if m.SchemaChange != nil {
		l = m.SchemaChange.Size()
		n += 1 + l + sovJobs(uint64(l))
	}
	return n
}

func sovJobs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJobs(x uint64) (n int) {
	return sovJobs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BackupDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func (m *RestoreDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RestoreDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RestoreDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowWaterMark", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LowWaterMark = append(m.LowWaterMark[:0], dAtA[iNdEx:postIndex]...)
			if m.LowWaterMark == nil {
				m.LowWaterMark = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func (m *SchemaChangeDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaChangeDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaChangeDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadAsOf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReadAsOf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJobs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedMicros", wireType)
			}
			m.StartedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinishedMicros", wireType)
			}
			m.FinishedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinishedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedMicros", wireType)
			}
			m.ModifiedMicros = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedMicros |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DescriptorIDs = append(m.DescriptorIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowJobs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthJobs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowJobs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (github_com_cockroachdb_cockroach_pkg_sql_sqlbase.ID(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DescriptorIDs = append(m.DescriptorIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorIDs", wireType)
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FractionCompleted", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(dAtA[iNdEx-4])
			v |= uint32(dAtA[iNdEx-3]) << 8
			v |= uint32(dAtA[iNdEx-2]) << 16
			v |= uint32(dAtA[iNdEx-1]) << 24
			m.FractionCompleted = float32(math.Float32frombits(v))
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BackupDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_Backup{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Restore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RestoreDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_Restore{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaChange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJobs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJobs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SchemaChangeDetails{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Details = &Payload_SchemaChange{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJobs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJobs
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
func skipJobs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJobs
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
					return 0, ErrIntOverflowJobs
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
					return 0, ErrIntOverflowJobs
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthJobs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJobs
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
				next, err := skipJobs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthJobs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJobs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/sql/jobs/jobs.proto", fileDescriptorJobs) }

var fileDescriptorJobs = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x8f, 0xd2, 0x4c,
	0x1c, 0x6f, 0x17, 0x96, 0x97, 0xe1, 0x65, 0x9f, 0x67, 0xdc, 0x43, 0x43, 0x62, 0x61, 0x89, 0xba,
	0x5c, 0x6c, 0x13, 0x37, 0xf1, 0x62, 0x62, 0xb2, 0x2c, 0x07, 0x30, 0x21, 0x9a, 0x6a, 0xa2, 0xf1,
	0xd2, 0x4c, 0xa7, 0x03, 0x1d, 0x99, 0x32, 0x65, 0x66, 0x08, 0xf1, 0x5b, 0xf8, 0xb1, 0x38, 0x9a,
	0x78, 0xf1, 0x44, 0x14, 0xbf, 0x85, 0x27, 0xd3, 0x29, 0x85, 0x25, 0xcb, 0xa5, 0xe9, 0xfc, 0xde,
	0xf2, 0xef, 0xbf, 0xbf, 0x01, 0x57, 0x98, 0xe3, 0x99, 0xe0, 0x08, 0x47, 0x6e, 0x32, 0x9b, 0xba,
	0x72, 0xc1, 0xdc, 0x2f, 0x3c, 0x90, 0xfa, 0xe1, 0x24, 0x82, 0x2b, 0x0e, 0xe1, 0x5e, 0xe2, 0xc8,
	0x05, 0x73, 0x52, 0xa6, 0x75, 0x39, 0xe5, 0x53, 0xae, 0x69, 0x37, 0x7d, 0xcb, 0x94, 0xad, 0x67,
	0xc7, 0x61, 0x4b, 0x45, 0x99, 0x1b, 0x31, 0xec, 0x2a, 0x1a, 0x13, 0xa9, 0x50, 0x9c, 0x64, 0xba,
	0xee, 0x05, 0x68, 0xf4, 0x11, 0x9e, 0x2d, 0x93, 0x01, 0x51, 0x88, 0x32, 0xd9, 0x7d, 0x09, 0x9a,
	0x1e, 0x91, 0x8a, 0x0b, 0xb2, 0x43, 0xe0, 0x13, 0xd0, 0x64, 0x7c, 0xe5, 0xaf, 0x90, 0x22, 0xc2,
	0x8f, 0x91, 0x98, 0x59, 0x66, 0xc7, 0xec, 0xd5, 0xbd, 0x3a, 0xe3, 0xab, 0x8f, 0x29, 0x38, 0x46,
	0x62, 0xd6, 0xfd, 0x04, 0x1e, 0xbd, 0xc7, 0x11, 0x89, 0xd1, 0x5d, 0x84, 0xe6, 0xd3, 0xbd, 0xf9,
	0x16, 0x00, 0x41, 0x50, 0xe8, 0x23, 0xe9, 0xf3, 0x89, 0x36, 0xd6, 0x5e, 0x3c, 0x76, 0x0e, 0x9f,
	0x91, 0x0e, 0xe6, 0x44, 0x0c, 0x3b, 0x1f, 0xf2, 0xc1, 0xfa, 0xc5, 0xf5, 0xa6, 0x6d, 0x78, 0x95,
	0xd4, 0x76, 0x2b, 0xdf, 0x4e, 0xba, 0x3f, 0x8a, 0xa0, 0xfc, 0x0e, 0x7d, 0x65, 0x1c, 0x85, 0xb0,
	0x03, 0x6a, 0x21, 0x91, 0x58, 0xd0, 0x44, 0x51, 0x3e, 0xd7, 0x79, 0x55, 0xef, 0x3e, 0x04, 0x5b,
	0xa0, 0xb2, 0x94, 0x44, 0xcc, 0x51, 0x4c, 0xac, 0x33, 0x4d, 0xef, 0xcf, 0xf0, 0x29, 0x68, 0x4a,
	0x85, 0x84, 0x22, 0xa1, 0x1f, 0x53, 0x2c, 0xb8, 0xb4, 0x0a, 0x1d, 0xb3, 0x57, 0xf0, 0x1a, 0x3b,
	0x74, 0xac, 0x41, 0x78, 0x0d, 0x2e, 0x26, 0x74, 0x4e, 0x65, 0x74, 0xd0, 0x15, 0xb5, 0xae, 0x99,
	0xc3, 0x07, 0x61, 0xcc, 0x43, 0x3a, 0xa1, 0x07, 0xe1, 0x79, 0x26, 0xcc, 0xe1, 0x9d, 0x90, 0x83,
	0x66, 0x3e, 0x23, 0x17, 0x3e, 0x0d, 0xa5, 0x55, 0xea, 0x14, 0x7a, 0x8d, 0xfe, 0x70, 0xbb, 0x69,
	0x37, 0x06, 0x7b, 0x66, 0x34, 0x90, 0x7f, 0x37, 0xed, 0x9b, 0x29, 0x55, 0xd1, 0x32, 0x70, 0x30,
	0x8f, 0xdd, 0xfd, 0xa2, 0xc2, 0xc0, 0x7d, 0x58, 0x0f, 0xb9, 0x60, 0x01, 0x92, 0xc4, 0x19, 0x0d,
	0xbc, 0xc6, 0x21, 0x7f, 0x14, 0x4a, 0xf8, 0x1c, 0xc0, 0x89, 0x40, 0x38, 0xdd, 0x88, 0x8f, 0x79,
	0x9c, 0x30, 0xa2, 0x48, 0x68, 0x95, 0x3b, 0x66, 0xef, 0xcc, 0xfb, 0x3f, 0x67, 0xee, 0x72, 0x02,
	0x5e, 0x82, 0x73, 0x22, 0x04, 0x17, 0x56, 0x45, 0x6f, 0x2c, 0x3b, 0xc0, 0x57, 0xa0, 0x14, 0xe8,
	0x6e, 0x58, 0x40, 0xff, 0xb7, 0x2b, 0xe7, 0x61, 0xfd, 0x9c, 0xa3, 0xf6, 0x0c, 0x0d, 0x6f, 0x67,
	0x81, 0xaf, 0x41, 0x59, 0x64, 0x3d, 0xb2, 0x6a, 0xda, 0xdd, 0x3d, 0xe5, 0x3e, 0xae, 0xda, 0xd0,
	0xf0, 0x72, 0x13, 0x1c, 0x83, 0xba, 0xbc, 0xd7, 0x27, 0xab, 0xae, 0x43, 0xae, 0x4f, 0x85, 0x9c,
	0xe8, 0xdd, 0xd0, 0xf0, 0x8e, 0xec, 0xfd, 0x2a, 0x28, 0x87, 0x19, 0xf5, 0xa6, 0x58, 0xa9, 0xfe,
	0x07, 0xfa, 0xf6, 0xfa, 0xb7, 0x6d, 0xac, 0xb7, 0xb6, 0xf9, 0x7d, 0x6b, 0x9b, 0x3f, 0xb7, 0xb6,
	0xf9, 0x6b, 0x6b, 0x9b, 0xdf, 0xfe, 0xd8, 0xc6, 0xe7, 0x62, 0x1a, 0x1a, 0x94, 0xf4, 0xfd, 0xb8,
	0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x20, 0xf6, 0xf0, 0x96, 0x03, 0x00, 0x00,
}
