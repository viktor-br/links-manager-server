package dao

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type sessionTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *sessionTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("sessions").
func (v *sessionTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *sessionTableType) Columns() []string {
	return []string{"id", "user_id", "ip", "created_at", "expires_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *sessionTableType) NewStruct() reform.Struct {
	return new(Session)
}

// NewRecord makes a new record for that table.
func (v *sessionTableType) NewRecord() reform.Record {
	return new(Session)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *sessionTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// SessionTable represents sessions view or table in SQL database.
var SessionTable = &sessionTableType{
	s: parse.StructInfo{Type: "Session", SQLSchema: "", SQLName: "sessions", Fields: []parse.FieldInfo{{Name: "ID", PKType: "string", Column: "id"}, {Name: "UserID", PKType: "", Column: "user_id"}, {Name: "RemoteAddr", PKType: "", Column: "ip"}, {Name: "CreatedAt", PKType: "", Column: "created_at"}, {Name: "ExpiresAt", PKType: "", Column: "expires_at"}}, PKFieldIndex: 0},
	z: new(Session).Values(),
}

// String returns a string representation of this struct or record.
func (s Session) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "UserID: " + reform.Inspect(s.UserID, true)
	res[2] = "RemoteAddr: " + reform.Inspect(s.RemoteAddr, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "ExpiresAt: " + reform.Inspect(s.ExpiresAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Session) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.UserID,
		s.RemoteAddr,
		s.CreatedAt,
		s.ExpiresAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Session) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.UserID,
		&s.RemoteAddr,
		&s.CreatedAt,
		&s.ExpiresAt,
	}
}

// View returns View object for that struct.
func (s *Session) View() reform.View {
	return SessionTable
}

// Table returns Table object for that record.
func (s *Session) Table() reform.Table {
	return SessionTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Session) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Session) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Session) HasPK() bool {
	return s.ID != SessionTable.z[SessionTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Session) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = SessionTable
	_ reform.Struct = new(Session)
	_ reform.Table  = SessionTable
	_ reform.Record = new(Session)
	_ fmt.Stringer  = new(Session)
)

func init() {
	parse.AssertUpToDate(&SessionTable.s, new(Session))
}
