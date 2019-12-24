// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// User holds the value of the "user" field.
	User *string `json:"user,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullInt64{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(values ...interface{}) error {
	if m, n := len(values), len(file.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field size", values[0])
	} else if value.Valid {
		f.Size = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		f.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user", values[2])
	} else if value.Valid {
		f.User = new(string)
		*f.User = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field group", values[3])
	} else if value.Valid {
		f.Group = value.String
	}
	return nil
}

// QueryOwner queries the owner edge of the File.
func (f *File) QueryOwner() *UserQuery {
	return (&FileClient{f.config}).QueryOwner(f)
}

// QueryType queries the type edge of the File.
func (f *File) QueryType() *FileTypeQuery {
	return (&FileClient{f.config}).QueryType(f)
}

// Update returns a builder for updating this File.
// Note that, you need to call File.Unwrap() before calling this method, if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	if v := f.User; v != nil {
		builder.WriteString(", user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", group=")
	builder.WriteString(f.Group)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (f *File) id() int {
	id, _ := strconv.Atoi(f.ID)
	return id
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
