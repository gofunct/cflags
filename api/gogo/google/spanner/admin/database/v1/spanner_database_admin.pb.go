// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/spanner/admin/database/v1/spanner_database_admin.proto

package google_spanner_admin_database_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import _ "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Indicates the current state of the database.
type Database_State int32

const (
	// Not specified.
	Database_STATE_UNSPECIFIED Database_State = 0
	// The database is still being created. Operations on the database may fail
	// with `FAILED_PRECONDITION` in this state.
	Database_CREATING Database_State = 1
	// The database is fully created and ready for use.
	Database_READY Database_State = 2
)

var Database_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
}
var Database_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"CREATING":          1,
	"READY":             2,
}

func (x Database_State) String() string {
	return proto.EnumName(Database_State_name, int32(x))
}
func (Database_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{0, 0}
}

// A Cloud Spanner database.
type Database struct {
	// Required. The name of the database. Values are of the form
	// `projects/<project>/instances/<instance>/databases/<database>`,
	// where `<database>` is as specified in the `CREATE DATABASE`
	// statement. This name can be passed to other API methods to
	// identify the database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The current database state.
	State Database_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.spanner.admin.database.v1.Database_State" json:"state,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptorSpannerDatabaseAdmin, []int{0} }

func (m *Database) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Database) GetState() Database_State {
	if m != nil {
		return m.State
	}
	return Database_STATE_UNSPECIFIED
}

// The request for [ListDatabases][google.spanner.admin.database.v1.DatabaseAdmin.ListDatabases].
type ListDatabasesRequest struct {
	// Required. The instance whose databases should be listed.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Number of databases to be returned in the response. If 0 or less,
	// defaults to the server's maximum allowed page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If non-empty, `page_token` should contain a
	// [next_page_token][google.spanner.admin.database.v1.ListDatabasesResponse.next_page_token] from a
	// previous [ListDatabasesResponse][google.spanner.admin.database.v1.ListDatabasesResponse].
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListDatabasesRequest) Reset()         { *m = ListDatabasesRequest{} }
func (m *ListDatabasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesRequest) ProtoMessage()    {}
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{1}
}

func (m *ListDatabasesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListDatabasesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDatabasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response for [ListDatabases][google.spanner.admin.database.v1.DatabaseAdmin.ListDatabases].
type ListDatabasesResponse struct {
	// Databases that matched the request.
	Databases []*Database `protobuf:"bytes,1,rep,name=databases" json:"databases,omitempty"`
	// `next_page_token` can be sent in a subsequent
	// [ListDatabases][google.spanner.admin.database.v1.DatabaseAdmin.ListDatabases] call to fetch more
	// of the matching databases.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListDatabasesResponse) Reset()         { *m = ListDatabasesResponse{} }
func (m *ListDatabasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesResponse) ProtoMessage()    {}
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{2}
}

func (m *ListDatabasesResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

func (m *ListDatabasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request for [CreateDatabase][google.spanner.admin.database.v1.DatabaseAdmin.CreateDatabase].
type CreateDatabaseRequest struct {
	// Required. The name of the instance that will serve the new database.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. A `CREATE DATABASE` statement, which specifies the ID of the
	// new database.  The database ID must conform to the regular expression
	// `[a-z][a-z0-9_\-]*[a-z0-9]` and be between 2 and 30 characters in length.
	// If the database ID is a reserved word or if it contains a hyphen, the
	// database ID must be enclosed in backticks (`` ` ``).
	CreateStatement string `protobuf:"bytes,2,opt,name=create_statement,json=createStatement,proto3" json:"create_statement,omitempty"`
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These
	// statements execute atomically with the creation of the database:
	// if there is an error in any statement, the database is not created.
	ExtraStatements []string `protobuf:"bytes,3,rep,name=extra_statements,json=extraStatements" json:"extra_statements,omitempty"`
}

func (m *CreateDatabaseRequest) Reset()         { *m = CreateDatabaseRequest{} }
func (m *CreateDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseRequest) ProtoMessage()    {}
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{3}
}

func (m *CreateDatabaseRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateDatabaseRequest) GetCreateStatement() string {
	if m != nil {
		return m.CreateStatement
	}
	return ""
}

func (m *CreateDatabaseRequest) GetExtraStatements() []string {
	if m != nil {
		return m.ExtraStatements
	}
	return nil
}

// Metadata type for the operation returned by
// [CreateDatabase][google.spanner.admin.database.v1.DatabaseAdmin.CreateDatabase].
type CreateDatabaseMetadata struct {
	// The database being created.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (m *CreateDatabaseMetadata) Reset()         { *m = CreateDatabaseMetadata{} }
func (m *CreateDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseMetadata) ProtoMessage()    {}
func (*CreateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{4}
}

func (m *CreateDatabaseMetadata) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

// The request for [GetDatabase][google.spanner.admin.database.v1.DatabaseAdmin.GetDatabase].
type GetDatabaseRequest struct {
	// Required. The name of the requested database. Values are of the form
	// `projects/<project>/instances/<instance>/databases/<database>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetDatabaseRequest) Reset()         { *m = GetDatabaseRequest{} }
func (m *GetDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseRequest) ProtoMessage()    {}
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{5}
}

func (m *GetDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Enqueues the given DDL statements to be applied, in order but not
// necessarily all at once, to the database schema at some point (or
// points) in the future. The server checks that the statements
// are executable (syntactically valid, name tables that exist, etc.)
// before enqueueing them, but they may still fail upon
// later execution (e.g., if a statement from another batch of
// statements is applied first and it conflicts in some way, or if
// there is some data-related problem like a `NULL` value in a column to
// which `NOT NULL` would be added). If a statement fails, all
// subsequent statements in the batch are automatically cancelled.
//
// Each batch of statements is assigned a name which can be used with
// the [Operations][google.longrunning.Operations] API to monitor
// progress. See the
// [operation_id][google.spanner.admin.database.v1.UpdateDatabaseDdlRequest.operation_id] field for more
// details.
type UpdateDatabaseDdlRequest struct {
	// Required. The database to update.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// DDL statements to be applied to the database.
	Statements []string `protobuf:"bytes,2,rep,name=statements" json:"statements,omitempty"`
	// If empty, the new update request is assigned an
	// automatically-generated operation ID. Otherwise, `operation_id`
	// is used to construct the name of the resulting
	// [Operation][google.longrunning.Operation].
	//
	// Specifying an explicit operation ID simplifies determining
	// whether the statements were executed in the event that the
	// [UpdateDatabaseDdl][google.spanner.admin.database.v1.DatabaseAdmin.UpdateDatabaseDdl] call is replayed,
	// or the return value is otherwise lost: the [database][google.spanner.admin.database.v1.UpdateDatabaseDdlRequest.database] and
	// `operation_id` fields can be combined to form the
	// [name][google.longrunning.Operation.name] of the resulting
	// [longrunning.Operation][google.longrunning.Operation]: `<database>/operations/<operation_id>`.
	//
	// `operation_id` should be unique within the database, and must be
	// a valid identifier: `[a-z][a-z0-9_]*`. Note that
	// automatically-generated operation IDs always begin with an
	// underscore. If the named operation already exists,
	// [UpdateDatabaseDdl][google.spanner.admin.database.v1.DatabaseAdmin.UpdateDatabaseDdl] returns
	// `ALREADY_EXISTS`.
	OperationId string `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (m *UpdateDatabaseDdlRequest) Reset()         { *m = UpdateDatabaseDdlRequest{} }
func (m *UpdateDatabaseDdlRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDatabaseDdlRequest) ProtoMessage()    {}
func (*UpdateDatabaseDdlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{6}
}

func (m *UpdateDatabaseDdlRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *UpdateDatabaseDdlRequest) GetStatements() []string {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *UpdateDatabaseDdlRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

// Metadata type for the operation returned by
// [UpdateDatabaseDdl][google.spanner.admin.database.v1.DatabaseAdmin.UpdateDatabaseDdl].
type UpdateDatabaseDdlMetadata struct {
	// The database being modified.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// For an update this list contains all the statements. For an
	// individual statement, this list contains only that statement.
	Statements []string `protobuf:"bytes,2,rep,name=statements" json:"statements,omitempty"`
	// Reports the commit timestamps of all statements that have
	// succeeded so far, where `commit_timestamps[i]` is the commit
	// timestamp for the statement `statements[i]`.
	CommitTimestamps []*google_protobuf3.Timestamp `protobuf:"bytes,3,rep,name=commit_timestamps,json=commitTimestamps" json:"commit_timestamps,omitempty"`
}

func (m *UpdateDatabaseDdlMetadata) Reset()         { *m = UpdateDatabaseDdlMetadata{} }
func (m *UpdateDatabaseDdlMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateDatabaseDdlMetadata) ProtoMessage()    {}
func (*UpdateDatabaseDdlMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{7}
}

func (m *UpdateDatabaseDdlMetadata) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *UpdateDatabaseDdlMetadata) GetStatements() []string {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *UpdateDatabaseDdlMetadata) GetCommitTimestamps() []*google_protobuf3.Timestamp {
	if m != nil {
		return m.CommitTimestamps
	}
	return nil
}

// The request for [DropDatabase][google.spanner.admin.database.v1.DatabaseAdmin.DropDatabase].
type DropDatabaseRequest struct {
	// Required. The database to be dropped.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (m *DropDatabaseRequest) Reset()         { *m = DropDatabaseRequest{} }
func (m *DropDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*DropDatabaseRequest) ProtoMessage()    {}
func (*DropDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{8}
}

func (m *DropDatabaseRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

// The request for [GetDatabaseDdl][google.spanner.admin.database.v1.DatabaseAdmin.GetDatabaseDdl].
type GetDatabaseDdlRequest struct {
	// Required. The database whose schema we wish to get.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (m *GetDatabaseDdlRequest) Reset()         { *m = GetDatabaseDdlRequest{} }
func (m *GetDatabaseDdlRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseDdlRequest) ProtoMessage()    {}
func (*GetDatabaseDdlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{9}
}

func (m *GetDatabaseDdlRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

// The response for [GetDatabaseDdl][google.spanner.admin.database.v1.DatabaseAdmin.GetDatabaseDdl].
type GetDatabaseDdlResponse struct {
	// A list of formatted DDL statements defining the schema of the database
	// specified in the request.
	Statements []string `protobuf:"bytes,1,rep,name=statements" json:"statements,omitempty"`
}

func (m *GetDatabaseDdlResponse) Reset()         { *m = GetDatabaseDdlResponse{} }
func (m *GetDatabaseDdlResponse) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseDdlResponse) ProtoMessage()    {}
func (*GetDatabaseDdlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerDatabaseAdmin, []int{10}
}

func (m *GetDatabaseDdlResponse) GetStatements() []string {
	if m != nil {
		return m.Statements
	}
	return nil
}

func init() {
	proto.RegisterType((*Database)(nil), "google.spanner.admin.database.v1.Database")
	proto.RegisterType((*ListDatabasesRequest)(nil), "google.spanner.admin.database.v1.ListDatabasesRequest")
	proto.RegisterType((*ListDatabasesResponse)(nil), "google.spanner.admin.database.v1.ListDatabasesResponse")
	proto.RegisterType((*CreateDatabaseRequest)(nil), "google.spanner.admin.database.v1.CreateDatabaseRequest")
	proto.RegisterType((*CreateDatabaseMetadata)(nil), "google.spanner.admin.database.v1.CreateDatabaseMetadata")
	proto.RegisterType((*GetDatabaseRequest)(nil), "google.spanner.admin.database.v1.GetDatabaseRequest")
	proto.RegisterType((*UpdateDatabaseDdlRequest)(nil), "google.spanner.admin.database.v1.UpdateDatabaseDdlRequest")
	proto.RegisterType((*UpdateDatabaseDdlMetadata)(nil), "google.spanner.admin.database.v1.UpdateDatabaseDdlMetadata")
	proto.RegisterType((*DropDatabaseRequest)(nil), "google.spanner.admin.database.v1.DropDatabaseRequest")
	proto.RegisterType((*GetDatabaseDdlRequest)(nil), "google.spanner.admin.database.v1.GetDatabaseDdlRequest")
	proto.RegisterType((*GetDatabaseDdlResponse)(nil), "google.spanner.admin.database.v1.GetDatabaseDdlResponse")
	proto.RegisterEnum("google.spanner.admin.database.v1.Database_State", Database_State_name, Database_State_value)
}

func init() {
	proto.RegisterFile("google/spanner/admin/database/v1/spanner_database_admin.proto", fileDescriptorSpannerDatabaseAdmin)
}

var fileDescriptorSpannerDatabaseAdmin = []byte{
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x19, 0xa7, 0xa9, 0xe2, 0x17, 0x27, 0x75, 0x06, 0x1c, 0xb9, 0x5b, 0x5a, 0xcc, 0x82,
	0x2a, 0xd7, 0x12, 0xbb, 0xd8, 0x69, 0x48, 0xb0, 0x14, 0x44, 0x6a, 0xbb, 0xc6, 0x12, 0xb4, 0x96,
	0xed, 0x56, 0x42, 0xb2, 0x64, 0x4d, 0xec, 0xc1, 0xda, 0xe2, 0xfd, 0xc1, 0xce, 0xb8, 0x6a, 0x8b,
	0x7a, 0x41, 0xe2, 0xc0, 0x19, 0x2a, 0x71, 0x03, 0x71, 0xe0, 0xc0, 0x89, 0x1b, 0x12, 0x47, 0x8e,
	0x5c, 0xf9, 0x17, 0xf8, 0x43, 0xd0, 0xcc, 0xee, 0xd8, 0xeb, 0x75, 0x1a, 0xaf, 0x39, 0xf4, 0xe6,
	0x7d, 0xef, 0xfb, 0xe6, 0x7d, 0xe6, 0xed, 0x7c, 0xc7, 0x0b, 0x27, 0x63, 0xd7, 0x1d, 0x4f, 0xa8,
	0xc9, 0x3c, 0xe2, 0x38, 0xd4, 0x37, 0xc9, 0xc8, 0xb6, 0x1c, 0x73, 0x44, 0x38, 0x39, 0x23, 0x8c,
	0x9a, 0x8f, 0xcb, 0x2a, 0x33, 0x50, 0xb1, 0x81, 0x94, 0x18, 0x9e, 0xef, 0x72, 0x17, 0x17, 0x82,
	0x72, 0x23, 0x14, 0x19, 0x41, 0x4e, 0x49, 0x8d, 0xc7, 0x65, 0xed, 0xcd, 0xb0, 0x01, 0xf1, 0x2c,
	0x93, 0x38, 0x8e, 0xcb, 0x09, 0xb7, 0x5c, 0x87, 0x05, 0xf5, 0xda, 0x8d, 0x30, 0x6b, 0x11, 0x5b,
	0xf4, 0xb2, 0x88, 0x3d, 0xf0, 0xdc, 0x89, 0x35, 0x7c, 0x1a, 0xe6, 0xb5, 0xc5, 0xfc, 0x42, 0xee,
	0x9d, 0x30, 0x37, 0x71, 0x9d, 0xb1, 0x3f, 0x75, 0x1c, 0xcb, 0x19, 0x9b, 0xae, 0x47, 0xfd, 0x85,
	0x06, 0xd7, 0x42, 0x91, 0x7c, 0x3a, 0x9b, 0x7e, 0x61, 0x52, 0xdb, 0xe3, 0x6a, 0x85, 0xb7, 0xe2,
	0x49, 0x6e, 0xd9, 0x94, 0x71, 0x62, 0x7b, 0x81, 0x40, 0xff, 0x09, 0xc1, 0x56, 0x3d, 0xdc, 0x0c,
	0xc6, 0x70, 0xc9, 0x21, 0x36, 0xcd, 0xa3, 0x02, 0x2a, 0xa6, 0x3b, 0xf2, 0x37, 0xbe, 0x0b, 0x9b,
	0x8c, 0x13, 0x4e, 0xf3, 0xa9, 0x02, 0x2a, 0xee, 0x56, 0xde, 0x37, 0x56, 0xcd, 0xc3, 0x50, 0xcb,
	0x19, 0x5d, 0x51, 0xd7, 0x09, 0xca, 0xf5, 0x23, 0xd8, 0x94, 0xcf, 0x38, 0x07, 0x7b, 0xdd, 0xde,
	0x69, 0xaf, 0x31, 0x78, 0x70, 0xaf, 0xdb, 0x6e, 0xd4, 0x5a, 0x77, 0x5b, 0x8d, 0x7a, 0xf6, 0x35,
	0x9c, 0x81, 0xad, 0x5a, 0xa7, 0x71, 0xda, 0x6b, 0xdd, 0x6b, 0x66, 0x11, 0x4e, 0xc3, 0x66, 0xa7,
	0x71, 0x5a, 0xff, 0x3c, 0x9b, 0xd2, 0x1f, 0xc1, 0x1b, 0x9f, 0x5a, 0x8c, 0xab, 0x55, 0x59, 0x87,
	0x7e, 0x35, 0xa5, 0x8c, 0xe3, 0x7d, 0xb8, 0xec, 0x11, 0x9f, 0x3a, 0x3c, 0xc4, 0x0d, 0x9f, 0xf0,
	0x35, 0x48, 0x7b, 0x64, 0x4c, 0x07, 0xcc, 0x7a, 0x46, 0xf3, 0x1b, 0x05, 0x54, 0xdc, 0xec, 0x6c,
	0x89, 0x40, 0xd7, 0x7a, 0x46, 0xf1, 0x75, 0x00, 0x99, 0xe4, 0xee, 0x97, 0xd4, 0xc9, 0x5f, 0x92,
	0x85, 0x52, 0xde, 0x13, 0x01, 0xfd, 0x3b, 0x04, 0xb9, 0x58, 0x33, 0xe6, 0xb9, 0x0e, 0xa3, 0xf8,
	0x13, 0x48, 0xab, 0x3d, 0xb2, 0x3c, 0x2a, 0x6c, 0x14, 0xb7, 0x2b, 0xa5, 0xe4, 0xa3, 0xe8, 0xcc,
	0x8b, 0xf1, 0x4d, 0xb8, 0xe2, 0xd0, 0x27, 0x7c, 0x10, 0xe1, 0x48, 0x49, 0x8e, 0x1d, 0x11, 0x6e,
	0xcf, 0x58, 0xbe, 0x45, 0x90, 0xab, 0xf9, 0x94, 0x70, 0x3a, 0x5b, 0x65, 0xc5, 0xce, 0x6f, 0x41,
	0x76, 0x28, 0x0b, 0x06, 0x72, 0xe4, 0xb6, 0x50, 0x04, 0x4b, 0x5f, 0x09, 0xe2, 0x5d, 0x15, 0x16,
	0x52, 0xfa, 0x84, 0xfb, 0x64, 0xae, 0x64, 0xf9, 0x8d, 0xc2, 0x86, 0x90, 0xca, 0xf8, 0x4c, 0xc9,
	0xf4, 0xdb, 0xb0, 0xbf, 0x88, 0xf1, 0x19, 0xe5, 0x44, 0x6c, 0x07, 0x6b, 0xb0, 0xa5, 0xb6, 0x15,
	0x92, 0xcc, 0x9e, 0xf5, 0x22, 0xe0, 0x26, 0xe5, 0x71, 0xf2, 0x73, 0x0e, 0x98, 0xfe, 0x14, 0xf2,
	0x0f, 0xbc, 0x51, 0x64, 0xfd, 0xfa, 0x68, 0xa2, 0xf4, 0x17, 0x74, 0xc0, 0x37, 0x00, 0x22, 0xf0,
	0x29, 0x09, 0x1f, 0x89, 0xe0, 0xb7, 0x21, 0x33, 0xf3, 0xca, 0xc0, 0x1a, 0xc9, 0xa3, 0x90, 0xee,
	0x6c, 0xcf, 0x62, 0xad, 0x91, 0xfe, 0x33, 0x82, 0xab, 0x4b, 0xbd, 0x93, 0x6c, 0x6f, 0x65, 0xf3,
	0x26, 0xec, 0x0d, 0x5d, 0xdb, 0xb6, 0xf8, 0x60, 0x66, 0xb8, 0x60, 0xc0, 0xdb, 0x15, 0x4d, 0x1d,
	0x1b, 0xe5, 0x49, 0xa3, 0xa7, 0x24, 0x9d, 0x6c, 0x50, 0x34, 0x0b, 0x30, 0xbd, 0x0c, 0xaf, 0xd7,
	0x7d, 0xd7, 0x8b, 0x0f, 0xf2, 0xa2, 0xd1, 0x1f, 0x40, 0x2e, 0x32, 0xfa, 0x64, 0xd3, 0xd4, 0x8f,
	0x61, 0x3f, 0x5e, 0x14, 0x9e, 0xfc, 0xc5, 0xad, 0xa2, 0xf8, 0x56, 0x2b, 0x2f, 0x32, 0xb0, 0xa3,
	0xea, 0x4e, 0x85, 0x03, 0xf0, 0x1f, 0x08, 0x76, 0x16, 0x5c, 0x84, 0x3f, 0x58, 0x6d, 0x95, 0xf3,
	0x3c, 0xae, 0x1d, 0xad, 0x5d, 0x17, 0x40, 0xeb, 0x87, 0xdf, 0xfc, 0xf3, 0xef, 0xf7, 0x29, 0x13,
	0xbf, 0x27, 0xee, 0xd4, 0xaf, 0x03, 0x7f, 0x9c, 0x78, 0xbe, 0xfb, 0x88, 0x0e, 0x39, 0x33, 0x4b,
	0xa6, 0xe5, 0x30, 0x4e, 0x9c, 0x21, 0x65, 0x66, 0xe9, 0xb9, 0x39, 0xf7, 0xe6, 0x2f, 0x08, 0x76,
	0x17, 0x0f, 0x3b, 0x4e, 0x80, 0x70, 0xae, 0x4b, 0xb5, 0xeb, 0xaa, 0x30, 0x72, 0x7b, 0x1b, 0xf7,
	0xd5, 0xe9, 0xd3, 0x8f, 0x25, 0x61, 0x45, 0x5f, 0x8f, 0xb0, 0x8a, 0x4a, 0xf8, 0x57, 0x04, 0xdb,
	0x91, 0x77, 0x85, 0x6f, 0xaf, 0x26, 0x5c, 0xb6, 0xa2, 0xb6, 0xc6, 0xed, 0x15, 0x9b, 0xa6, 0x70,
	0xed, 0x4b, 0x48, 0xe7, 0xa0, 0x66, 0xe9, 0x39, 0xfe, 0x1d, 0xc1, 0xde, 0x92, 0xbd, 0x70, 0x75,
	0x75, 0xe3, 0x97, 0xdd, 0x07, 0xab, 0x66, 0xfa, 0xb1, 0xe4, 0xac, 0x56, 0x0e, 0x25, 0xa7, 0x5a,
	0x31, 0x09, 0xab, 0x39, 0x1a, 0x4d, 0xc4, 0x6c, 0x7f, 0x44, 0x90, 0x89, 0xfa, 0x0d, 0x1f, 0x26,
	0x18, 0xd3, 0xb2, 0x3f, 0xb5, 0xfd, 0x25, 0x93, 0x37, 0xc4, 0xbf, 0xb2, 0xfe, 0xa1, 0x24, 0x3c,
	0x28, 0x95, 0xd7, 0x26, 0xc4, 0x7f, 0x21, 0xd8, 0x5d, 0xb4, 0x68, 0x92, 0xb3, 0x79, 0xee, 0x4d,
	0xa0, 0x1d, 0xaf, 0x5f, 0x18, 0x1a, 0xeb, 0x44, 0x6e, 0xe0, 0x08, 0xff, 0xbf, 0x11, 0xe3, 0x1f,
	0x10, 0x64, 0xba, 0x94, 0xb7, 0x88, 0xdd, 0x96, 0x1f, 0x3a, 0x58, 0x57, 0x24, 0x16, 0xb1, 0x45,
	0xdb, 0x68, 0x52, 0xd1, 0xe6, 0x62, 0x9a, 0x20, 0xab, 0xb7, 0x24, 0x4a, 0x4d, 0xff, 0x48, 0xa2,
	0xf8, 0x94, 0xb9, 0x53, 0x7f, 0x98, 0x08, 0xa5, 0xca, 0x22, 0x5d, 0xc4, 0x6b, 0x17, 0x58, 0xcd,
	0x8b, 0xb0, 0x9a, 0xaf, 0x04, 0x6b, 0x1c, 0xc3, 0xfa, 0x13, 0x01, 0xee, 0x51, 0x26, 0x83, 0xd4,
	0xb7, 0x2d, 0xc6, 0xc4, 0x77, 0x1f, 0x2e, 0xc6, 0x1a, 0x2f, 0x4b, 0x14, 0xe2, 0xad, 0x04, 0xca,
	0xf0, 0xc5, 0xde, 0x97, 0xd8, 0x2d, 0xbd, 0xbe, 0x3e, 0x36, 0x5f, 0x5a, 0xb5, 0x8a, 0x4a, 0x77,
	0x5e, 0x20, 0x78, 0x77, 0xe8, 0xda, 0x2b, 0x4f, 0xda, 0x9d, 0xab, 0xdd, 0x20, 0xb5, 0xf0, 0x27,
	0xd2, 0x16, 0xbe, 0x69, 0xa3, 0xdf, 0x52, 0x37, 0x9b, 0x41, 0x7d, 0x6d, 0xe2, 0x4e, 0x47, 0x46,
	0x28, 0x35, 0xa4, 0x66, 0xfe, 0xa5, 0xf9, 0xb0, 0xfc, 0xb7, 0x12, 0xf6, 0xa5, 0xb0, 0x1f, 0x0a,
	0xfb, 0x52, 0xd8, 0x57, 0xc2, 0xfe, 0xc3, 0xf2, 0xd9, 0x65, 0x69, 0xc9, 0x83, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x74, 0xd2, 0x7f, 0xbe, 0x18, 0x0c, 0x00, 0x00,
}
