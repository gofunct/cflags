// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/bigtable/admin/v2/bigtable_table_admin.proto

package google_bigtable_admin_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.CreateTable][google.bigtable.admin.v2.BigtableTableAdmin.CreateTable]
type CreateTableRequest struct {
	// The unique name of the instance in which to create the table.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The name by which the new table should be referred to within the parent
	// instance, e.g., `foobar` rather than `<parent>/tables/foobar`.
	TableId string `protobuf:"bytes,2,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	// The Table to create.
	Table *Table `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	// The optional list of row keys that will be used to initially split the
	// table into several tablets (tablets are similar to HBase regions).
	// Given two split keys, `s1` and `s2`, three tablets will be created,
	// spanning the key ranges: `[, s1), [s1, s2), [s2, )`.
	//
	// Example:
	//
	// * Row keys := `["a", "apple", "custom", "customer_1", "customer_2",`
	//                `"other", "zz"]`
	// * initial_split_keys := `["apple", "customer_1", "customer_2", "other"]`
	// * Key assignment:
	//     - Tablet 1 `[, apple)                => {"a"}.`
	//     - Tablet 2 `[apple, customer_1)      => {"apple", "custom"}.`
	//     - Tablet 3 `[customer_1, customer_2) => {"customer_1"}.`
	//     - Tablet 4 `[customer_2, other)      => {"customer_2"}.`
	//     - Tablet 5 `[other, )                => {"other", "zz"}.`
	InitialSplits []*CreateTableRequest_Split `protobuf:"bytes,4,rep,name=initial_splits,json=initialSplits" json:"initial_splits,omitempty"`
}

func (m *CreateTableRequest) Reset()         { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()    {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{0}
}

func (m *CreateTableRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateTableRequest) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *CreateTableRequest) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *CreateTableRequest) GetInitialSplits() []*CreateTableRequest_Split {
	if m != nil {
		return m.InitialSplits
	}
	return nil
}

// An initial split point for a newly created table.
type CreateTableRequest_Split struct {
	// Row key to use as an initial tablet boundary.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *CreateTableRequest_Split) Reset()         { *m = CreateTableRequest_Split{} }
func (m *CreateTableRequest_Split) String() string { return proto.CompactTextString(m) }
func (*CreateTableRequest_Split) ProtoMessage()    {}
func (*CreateTableRequest_Split) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{0, 0}
}

func (m *CreateTableRequest_Split) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange][google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange]
type DropRowRangeRequest struct {
	// The unique name of the table on which to drop a range of rows.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Delete all rows or by prefix.
	//
	// Types that are valid to be assigned to Target:
	//	*DropRowRangeRequest_RowKeyPrefix
	//	*DropRowRangeRequest_DeleteAllDataFromTable
	Target isDropRowRangeRequest_Target `protobuf_oneof:"target"`
}

func (m *DropRowRangeRequest) Reset()         { *m = DropRowRangeRequest{} }
func (m *DropRowRangeRequest) String() string { return proto.CompactTextString(m) }
func (*DropRowRangeRequest) ProtoMessage()    {}
func (*DropRowRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{1}
}

type isDropRowRangeRequest_Target interface {
	isDropRowRangeRequest_Target()
}

type DropRowRangeRequest_RowKeyPrefix struct {
	RowKeyPrefix []byte `protobuf:"bytes,2,opt,name=row_key_prefix,json=rowKeyPrefix,proto3,oneof"`
}
type DropRowRangeRequest_DeleteAllDataFromTable struct {
	DeleteAllDataFromTable bool `protobuf:"varint,3,opt,name=delete_all_data_from_table,json=deleteAllDataFromTable,proto3,oneof"`
}

func (*DropRowRangeRequest_RowKeyPrefix) isDropRowRangeRequest_Target()           {}
func (*DropRowRangeRequest_DeleteAllDataFromTable) isDropRowRangeRequest_Target() {}

func (m *DropRowRangeRequest) GetTarget() isDropRowRangeRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DropRowRangeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DropRowRangeRequest) GetRowKeyPrefix() []byte {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_RowKeyPrefix); ok {
		return x.RowKeyPrefix
	}
	return nil
}

func (m *DropRowRangeRequest) GetDeleteAllDataFromTable() bool {
	if x, ok := m.GetTarget().(*DropRowRangeRequest_DeleteAllDataFromTable); ok {
		return x.DeleteAllDataFromTable
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DropRowRangeRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DropRowRangeRequest_OneofMarshaler, _DropRowRangeRequest_OneofUnmarshaler, _DropRowRangeRequest_OneofSizer, []interface{}{
		(*DropRowRangeRequest_RowKeyPrefix)(nil),
		(*DropRowRangeRequest_DeleteAllDataFromTable)(nil),
	}
}

func _DropRowRangeRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		t := uint64(0)
		if x.DeleteAllDataFromTable {
			t = 1
		}
		_ = b.EncodeVarint(3<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("DropRowRangeRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _DropRowRangeRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DropRowRangeRequest)
	switch tag {
	case 2: // target.row_key_prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &DropRowRangeRequest_RowKeyPrefix{x}
		return true, err
	case 3: // target.delete_all_data_from_table
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Target = &DropRowRangeRequest_DeleteAllDataFromTable{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _DropRowRangeRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DropRowRangeRequest)
	// target
	switch x := m.Target.(type) {
	case *DropRowRangeRequest_RowKeyPrefix:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RowKeyPrefix)))
		n += len(x.RowKeyPrefix)
	case *DropRowRangeRequest_DeleteAllDataFromTable:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesRequest struct {
	// The unique name of the instance for which tables should be listed.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The view to be applied to the returned tables' fields.
	// Defaults to `NAME_ONLY` if unspecified; no others are currently supported.
	View Table_View `protobuf:"varint,2,opt,name=view,proto3,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListTablesRequest) Reset()         { *m = ListTablesRequest{} }
func (m *ListTablesRequest) String() string { return proto.CompactTextString(m) }
func (*ListTablesRequest) ProtoMessage()    {}
func (*ListTablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{2}
}

func (m *ListTablesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListTablesRequest) GetView() Table_View {
	if m != nil {
		return m.View
	}
	return Table_VIEW_UNSPECIFIED
}

func (m *ListTablesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ListTables][google.bigtable.admin.v2.BigtableTableAdmin.ListTables]
type ListTablesResponse struct {
	// The tables present in the requested instance.
	Tables []*Table `protobuf:"bytes,1,rep,name=tables" json:"tables,omitempty"`
	// Set if not all tables could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListTablesResponse) Reset()         { *m = ListTablesResponse{} }
func (m *ListTablesResponse) String() string { return proto.CompactTextString(m) }
func (*ListTablesResponse) ProtoMessage()    {}
func (*ListTablesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{3}
}

func (m *ListTablesResponse) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *ListTablesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.GetTable][google.bigtable.admin.v2.BigtableTableAdmin.GetTable]
type GetTableRequest struct {
	// The unique name of the requested table.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The view to be applied to the returned table's fields.
	// Defaults to `SCHEMA_VIEW` if unspecified.
	View Table_View `protobuf:"varint,2,opt,name=view,proto3,enum=google.bigtable.admin.v2.Table_View" json:"view,omitempty"`
}

func (m *GetTableRequest) Reset()         { *m = GetTableRequest{} }
func (m *GetTableRequest) String() string { return proto.CompactTextString(m) }
func (*GetTableRequest) ProtoMessage()    {}
func (*GetTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{4}
}

func (m *GetTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTableRequest) GetView() Table_View {
	if m != nil {
		return m.View
	}
	return Table_VIEW_UNSPECIFIED
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable][google.bigtable.admin.v2.BigtableTableAdmin.DeleteTable]
type DeleteTableRequest struct {
	// The unique name of the table to be deleted.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteTableRequest) Reset()         { *m = DeleteTableRequest{} }
func (m *DeleteTableRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTableRequest) ProtoMessage()    {}
func (*DeleteTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{5}
}

func (m *DeleteTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for
// [google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies][google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies]
type ModifyColumnFamiliesRequest struct {
	// The unique name of the table whose families should be modified.
	// Values are of the form
	// `projects/<project>/instances/<instance>/tables/<table>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Modifications to be atomically applied to the specified table's families.
	// Entries are applied in order, meaning that earlier modifications can be
	// masked by later ones (in the case of repeated updates to the same family,
	// for example).
	Modifications []*ModifyColumnFamiliesRequest_Modification `protobuf:"bytes,2,rep,name=modifications" json:"modifications,omitempty"`
}

func (m *ModifyColumnFamiliesRequest) Reset()         { *m = ModifyColumnFamiliesRequest{} }
func (m *ModifyColumnFamiliesRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest) ProtoMessage()    {}
func (*ModifyColumnFamiliesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{6}
}

func (m *ModifyColumnFamiliesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModifyColumnFamiliesRequest) GetModifications() []*ModifyColumnFamiliesRequest_Modification {
	if m != nil {
		return m.Modifications
	}
	return nil
}

// A create, update, or delete of a particular column family.
type ModifyColumnFamiliesRequest_Modification struct {
	// The ID of the column family to be modified.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Column familiy modifications.
	//
	// Types that are valid to be assigned to Mod:
	//	*ModifyColumnFamiliesRequest_Modification_Create
	//	*ModifyColumnFamiliesRequest_Modification_Update
	//	*ModifyColumnFamiliesRequest_Modification_Drop
	Mod isModifyColumnFamiliesRequest_Modification_Mod `protobuf_oneof:"mod"`
}

func (m *ModifyColumnFamiliesRequest_Modification) Reset() {
	*m = ModifyColumnFamiliesRequest_Modification{}
}
func (m *ModifyColumnFamiliesRequest_Modification) String() string { return proto.CompactTextString(m) }
func (*ModifyColumnFamiliesRequest_Modification) ProtoMessage()    {}
func (*ModifyColumnFamiliesRequest_Modification) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableTableAdmin, []int{6, 0}
}

type isModifyColumnFamiliesRequest_Modification_Mod interface {
	isModifyColumnFamiliesRequest_Modification_Mod()
}

type ModifyColumnFamiliesRequest_Modification_Create struct {
	Create *ColumnFamily `protobuf:"bytes,2,opt,name=create,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Update struct {
	Update *ColumnFamily `protobuf:"bytes,3,opt,name=update,oneof"`
}
type ModifyColumnFamiliesRequest_Modification_Drop struct {
	Drop bool `protobuf:"varint,4,opt,name=drop,proto3,oneof"`
}

func (*ModifyColumnFamiliesRequest_Modification_Create) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Update) isModifyColumnFamiliesRequest_Modification_Mod() {
}
func (*ModifyColumnFamiliesRequest_Modification_Drop) isModifyColumnFamiliesRequest_Modification_Mod() {
}

func (m *ModifyColumnFamiliesRequest_Modification) GetMod() isModifyColumnFamiliesRequest_Modification_Mod {
	if m != nil {
		return m.Mod
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ModifyColumnFamiliesRequest_Modification) GetCreate() *ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Create); ok {
		return x.Create
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetUpdate() *ColumnFamily {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Update); ok {
		return x.Update
	}
	return nil
}

func (m *ModifyColumnFamiliesRequest_Modification) GetDrop() bool {
	if x, ok := m.GetMod().(*ModifyColumnFamiliesRequest_Modification_Drop); ok {
		return x.Drop
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModifyColumnFamiliesRequest_Modification) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModifyColumnFamiliesRequest_Modification_OneofMarshaler, _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler, _ModifyColumnFamiliesRequest_Modification_OneofSizer, []interface{}{
		(*ModifyColumnFamiliesRequest_Modification_Create)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Update)(nil),
		(*ModifyColumnFamiliesRequest_Modification_Drop)(nil),
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Update:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		t := uint64(0)
		if x.Drop {
			t = 1
		}
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ModifyColumnFamiliesRequest_Modification.Mod has unexpected type %T", x)
	}
	return nil
}

func _ModifyColumnFamiliesRequest_Modification_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	switch tag {
	case 2: // mod.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Create{msg}
		return true, err
	case 3: // mod.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ColumnFamily)
		err := b.DecodeMessage(msg)
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Update{msg}
		return true, err
	case 4: // mod.drop
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Mod = &ModifyColumnFamiliesRequest_Modification_Drop{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ModifyColumnFamiliesRequest_Modification_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModifyColumnFamiliesRequest_Modification)
	// mod
	switch x := m.Mod.(type) {
	case *ModifyColumnFamiliesRequest_Modification_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModifyColumnFamiliesRequest_Modification_Drop:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "google.bigtable.admin.v2.CreateTableRequest")
	proto.RegisterType((*CreateTableRequest_Split)(nil), "google.bigtable.admin.v2.CreateTableRequest.Split")
	proto.RegisterType((*DropRowRangeRequest)(nil), "google.bigtable.admin.v2.DropRowRangeRequest")
	proto.RegisterType((*ListTablesRequest)(nil), "google.bigtable.admin.v2.ListTablesRequest")
	proto.RegisterType((*ListTablesResponse)(nil), "google.bigtable.admin.v2.ListTablesResponse")
	proto.RegisterType((*GetTableRequest)(nil), "google.bigtable.admin.v2.GetTableRequest")
	proto.RegisterType((*DeleteTableRequest)(nil), "google.bigtable.admin.v2.DeleteTableRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest")
	proto.RegisterType((*ModifyColumnFamiliesRequest_Modification)(nil), "google.bigtable.admin.v2.ModifyColumnFamiliesRequest.Modification")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/v2/bigtable_table_admin.proto", fileDescriptorBigtableTableAdmin)
}

var fileDescriptorBigtableTableAdmin = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xcf, 0xec, 0x6e, 0xd2, 0xe4, 0xed, 0x26, 0x85, 0xa1, 0x0a, 0xdb, 0x6d, 0x69, 0x23, 0x53,
	0x45, 0x61, 0x09, 0xb6, 0xe4, 0x2a, 0x2a, 0x8a, 0x8a, 0xa0, 0x9b, 0xd0, 0x86, 0x7f, 0x52, 0x64,
	0xaa, 0x48, 0x48, 0x91, 0xac, 0xc9, 0x7a, 0x62, 0x86, 0xd8, 0x33, 0xc6, 0x1e, 0x67, 0xbb, 0xaa,
	0x7a, 0x41, 0x48, 0x48, 0x5c, 0x7b, 0xaa, 0xf8, 0x06, 0x7c, 0x02, 0xc4, 0x85, 0x3b, 0x57, 0xce,
	0xdc, 0xf8, 0x08, 0x7c, 0x80, 0x6a, 0x66, 0xec, 0x64, 0x93, 0x8d, 0xb3, 0xdd, 0x5c, 0xac, 0x99,
	0xf7, 0x7e, 0xef, 0xbd, 0xdf, 0xbc, 0x37, 0x3f, 0x0f, 0xdc, 0x0f, 0x85, 0x08, 0x23, 0xea, 0x1c,
	0xb0, 0x50, 0x92, 0x83, 0x88, 0x3a, 0x24, 0x88, 0x19, 0x77, 0x8e, 0xdd, 0x13, 0x8b, 0x6f, 0xbe,
	0xda, 0x6e, 0x27, 0xa9, 0x90, 0x02, 0xb7, 0x4d, 0x90, 0x5d, 0x42, 0x6c, 0xe3, 0x3c, 0x76, 0x3b,
	0xb7, 0x8b, 0x74, 0x24, 0x61, 0x0e, 0xe1, 0x5c, 0x48, 0x22, 0x99, 0xe0, 0x99, 0x89, 0xeb, 0xdc,
	0xab, 0x2c, 0x66, 0xd2, 0x18, 0xd4, 0xfb, 0x05, 0x2a, 0x12, 0x3c, 0x4c, 0x73, 0xce, 0x19, 0x0f,
	0x1d, 0x91, 0xd0, 0xf4, 0x4c, 0xaa, 0x3b, 0x05, 0x48, 0xef, 0x0e, 0xf2, 0x43, 0x27, 0xc8, 0x0d,
	0xa0, 0xf0, 0xdf, 0x3a, 0xef, 0xa7, 0x71, 0x22, 0x87, 0x85, 0xf3, 0xee, 0x79, 0xa7, 0x64, 0x31,
	0xcd, 0x24, 0x89, 0x13, 0x03, 0xb0, 0xfe, 0x47, 0x80, 0xb7, 0x52, 0x4a, 0x24, 0x7d, 0xaa, 0x88,
	0x79, 0xf4, 0xc7, 0x9c, 0x66, 0x12, 0x2f, 0xc3, 0x5c, 0x42, 0x52, 0xca, 0x65, 0x1b, 0xad, 0xa0,
	0xb5, 0x05, 0xaf, 0xd8, 0xe1, 0x9b, 0x30, 0x6f, 0x9a, 0xc4, 0x82, 0x76, 0x4d, 0x7b, 0xae, 0xe9,
	0xfd, 0x17, 0x01, 0xde, 0x80, 0x59, 0xbd, 0x6c, 0xd7, 0x57, 0xd0, 0x5a, 0xd3, 0xbd, 0x6b, 0x57,
	0xb5, 0xce, 0x36, 0x95, 0x0c, 0x1a, 0x7f, 0x07, 0x4b, 0x8c, 0x33, 0xc9, 0x48, 0xe4, 0x67, 0x49,
	0xc4, 0x64, 0xd6, 0x6e, 0xac, 0xd4, 0xd7, 0x9a, 0xae, 0x5b, 0x1d, 0x3f, 0xce, 0xd7, 0xfe, 0x56,
	0x85, 0x7a, 0x8b, 0x45, 0x26, 0xbd, 0xcb, 0x3a, 0x37, 0x61, 0x56, 0xaf, 0xf0, 0x5b, 0x50, 0x3f,
	0xa2, 0x43, 0x7d, 0x94, 0x96, 0xa7, 0x96, 0xd6, 0x2b, 0x04, 0xef, 0x6c, 0xa7, 0x22, 0xf1, 0xc4,
	0xc0, 0x23, 0x3c, 0x3c, 0x39, 0x37, 0x86, 0x06, 0x27, 0x31, 0x2d, 0x4e, 0xad, 0xd7, 0x78, 0x15,
	0x96, 0x52, 0x31, 0xf0, 0x8f, 0xe8, 0xd0, 0x4f, 0x52, 0x7a, 0xc8, 0x9e, 0xe9, 0x93, 0xb7, 0x76,
	0x66, 0xbc, 0x56, 0x2a, 0x06, 0x5f, 0xd1, 0xe1, 0xae, 0xb6, 0xe2, 0x87, 0xd0, 0x09, 0x68, 0x44,
	0x25, 0xf5, 0x49, 0x14, 0xf9, 0x01, 0x91, 0xc4, 0x3f, 0x4c, 0x45, 0xec, 0x9f, 0x76, 0x65, 0x7e,
	0x67, 0xc6, 0x5b, 0x36, 0x98, 0x47, 0x51, 0xb4, 0x4d, 0x24, 0x79, 0x9c, 0x8a, 0x58, 0x1f, 0xa4,
	0x37, 0x0f, 0x73, 0x92, 0xa4, 0x21, 0x95, 0xd6, 0xcf, 0x08, 0xde, 0xfe, 0x9a, 0x65, 0x52, 0xdb,
	0xb3, 0x49, 0x13, 0xf9, 0x18, 0x1a, 0xc7, 0x8c, 0x0e, 0x34, 0xa7, 0x25, 0xf7, 0xde, 0x84, 0xae,
	0xdb, 0x7b, 0x8c, 0x0e, 0x3c, 0x1d, 0x81, 0xdf, 0x03, 0x48, 0x48, 0x48, 0x7d, 0x29, 0x8e, 0x28,
	0xd7, 0xfc, 0x16, 0xbc, 0x05, 0x65, 0x79, 0xaa, 0x0c, 0x56, 0x0e, 0x78, 0x94, 0x45, 0x96, 0x08,
	0x9e, 0x51, 0xfc, 0x40, 0xd1, 0x54, 0x96, 0x36, 0xd2, 0x63, 0x9a, 0x38, 0xe6, 0x02, 0x8e, 0x57,
	0xe1, 0x3a, 0xa7, 0xcf, 0xa4, 0x3f, 0x52, 0xd2, 0x5c, 0xa0, 0x45, 0x65, 0xde, 0x3d, 0x29, 0xeb,
	0xc3, 0xf5, 0x27, 0x54, 0x9e, 0xb9, 0x8c, 0x17, 0x0d, 0xe5, 0xca, 0xc7, 0xb6, 0xd6, 0x00, 0x6f,
	0xeb, 0x11, 0x4c, 0xaa, 0x61, 0xfd, 0x5b, 0x83, 0x5b, 0xdf, 0x88, 0x80, 0x1d, 0x0e, 0xb7, 0x44,
	0x94, 0xc7, 0xfc, 0x31, 0x89, 0x59, 0xc4, 0x4e, 0x47, 0x72, 0x11, 0xaf, 0xef, 0x61, 0x31, 0x56,
	0x21, 0xac, 0x6f, 0x44, 0xdc, 0xae, 0xe9, 0x36, 0xf5, 0xaa, 0x09, 0x5e, 0x52, 0xc1, 0xf8, 0x8a,
	0x54, 0xde, 0xd9, 0xc4, 0x9d, 0xbf, 0x10, 0xb4, 0x46, 0xfd, 0x78, 0x09, 0x6a, 0x2c, 0x28, 0xc8,
	0xd4, 0x58, 0x80, 0x3f, 0x83, 0xb9, 0xbe, 0x56, 0x8a, 0x6e, 0x52, 0xd3, 0x5d, 0xbd, 0x44, 0x51,
	0xa7, 0xd5, 0x87, 0x3b, 0x33, 0x5e, 0x11, 0xa7, 0x32, 0xe4, 0x49, 0xa0, 0x32, 0xd4, 0xa7, 0xcd,
	0x60, 0xe2, 0xf0, 0x0d, 0x68, 0x04, 0xa9, 0x48, 0xda, 0x8d, 0xe2, 0xf6, 0xeb, 0x5d, 0x6f, 0x16,
	0xea, 0xb1, 0x08, 0xdc, 0x3f, 0xae, 0x01, 0xee, 0x15, 0x99, 0xf4, 0x30, 0x1e, 0xa9, 0x6c, 0xf8,
	0x25, 0x82, 0xe6, 0x88, 0xc4, 0xf1, 0xfa, 0x34, 0x7f, 0x82, 0xce, 0xa4, 0x0b, 0x69, 0x6d, 0xfc,
	0xf4, 0xcf, 0x7f, 0x2f, 0x6b, 0x8e, 0xd5, 0x55, 0x7f, 0xe3, 0xe7, 0x46, 0x45, 0x9f, 0x24, 0xa9,
	0xf8, 0x81, 0xf6, 0x65, 0xe6, 0x74, 0x1d, 0xc6, 0x33, 0x49, 0x78, 0x9f, 0x66, 0x4e, 0xf7, 0x85,
	0xf9, 0x5b, 0x67, 0x9b, 0xa8, 0x8b, 0x7f, 0x43, 0x00, 0xa7, 0x7a, 0xc0, 0x1f, 0x56, 0x97, 0x19,
	0xd3, 0x6e, 0x67, 0xfd, 0xcd, 0xc0, 0x46, 0x62, 0x96, 0xab, 0x09, 0xae, 0xe3, 0x29, 0x08, 0xe2,
	0x5f, 0x11, 0xcc, 0x97, 0xb2, 0xc1, 0x1f, 0x54, 0x97, 0x3b, 0x27, 0xad, 0xc9, 0xdd, 0x3a, 0x4b,
	0x46, 0x5d, 0xf1, 0x0a, 0x2a, 0x05, 0x13, 0xa7, 0xfb, 0x02, 0xff, 0x82, 0xa0, 0x39, 0x22, 0xb1,
	0xcb, 0x06, 0x38, 0xae, 0xc4, 0xce, 0x72, 0x89, 0x2e, 0xdf, 0x2c, 0xfb, 0x73, 0xf5, 0xa0, 0x95,
	0x4c, 0xba, 0xd3, 0x30, 0xf9, 0x13, 0xc1, 0x8d, 0x8b, 0xf4, 0x85, 0x37, 0xae, 0xa4, 0xc7, 0xc9,
	0xed, 0xfa, 0x52, 0x93, 0xdc, 0xb6, 0x3e, 0x7d, 0x73, 0x92, 0x9b, 0xf1, 0x05, 0x05, 0xd5, 0x8d,
	0x7b, 0x85, 0xa0, 0x35, 0xfa, 0x46, 0xe1, 0x8f, 0x2e, 0xe9, 0xe3, 0xf8, 0x5b, 0x56, 0xd9, 0xc8,
	0x9e, 0xe6, 0xf8, 0xd0, 0x7a, 0x30, 0x05, 0xc7, 0x60, 0x24, 0xff, 0x26, 0xea, 0xf6, 0x9e, 0xc3,
	0xed, 0xbe, 0x88, 0x2b, 0xf9, 0xf4, 0xde, 0x1d, 0xd7, 0xf5, 0xae, 0x62, 0xb1, 0x8b, 0x7e, 0xaf,
	0xdd, 0x79, 0x62, 0xa2, 0xb6, 0x22, 0x91, 0x07, 0x76, 0x09, 0xb4, 0x35, 0xc8, 0xde, 0x73, 0xff,
	0x2e, 0x01, 0xfb, 0x1a, 0xb0, 0x5f, 0x02, 0xf6, 0x35, 0x60, 0x7f, 0xcf, 0x3d, 0x98, 0xd3, 0x07,
	0xba, 0xff, 0x3a, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xeb, 0x72, 0x58, 0xd2, 0x09, 0x00, 0x00,
}
