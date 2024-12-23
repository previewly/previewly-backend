// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wsw/backend/ent/types"
	"wsw/backend/ent/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Url is the model entity for the Url schema.
type Url struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Status holds the value of the "status" field.
	Status types.StatusEnum `json:"status,omitempty"`
	// RelativePath holds the value of the "relative_path" field.
	RelativePath *string `json:"relative_path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UrlQuery when eager-loading is set.
	Edges        UrlEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UrlEdges holds the relations/edges for other nodes in the graph.
type UrlEdges struct {
	// Errorresult holds the value of the errorresult edge.
	Errorresult []*ErrorResult `json:"errorresult,omitempty"`
	// Stat holds the value of the stat edge.
	Stat []*Stat `json:"stat,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ErrorresultOrErr returns the Errorresult value or an error if the edge
// was not loaded in eager-loading.
func (e UrlEdges) ErrorresultOrErr() ([]*ErrorResult, error) {
	if e.loadedTypes[0] {
		return e.Errorresult, nil
	}
	return nil, &NotLoadedError{edge: "errorresult"}
}

// StatOrErr returns the Stat value or an error if the edge
// was not loaded in eager-loading.
func (e UrlEdges) StatOrErr() ([]*Stat, error) {
	if e.loadedTypes[1] {
		return e.Stat, nil
	}
	return nil, &NotLoadedError{edge: "stat"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Url) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			values[i] = new(sql.NullInt64)
		case url.FieldURL, url.FieldStatus, url.FieldRelativePath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Url fields.
func (u *Url) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case url.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		case url.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = types.StatusEnum(value.String)
			}
		case url.FieldRelativePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relative_path", values[i])
			} else if value.Valid {
				u.RelativePath = new(string)
				*u.RelativePath = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Url.
// This includes values selected through modifiers, order, etc.
func (u *Url) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryErrorresult queries the "errorresult" edge of the Url entity.
func (u *Url) QueryErrorresult() *ErrorResultQuery {
	return NewURLClient(u.config).QueryErrorresult(u)
}

// QueryStat queries the "stat" edge of the Url entity.
func (u *Url) QueryStat() *StatQuery {
	return NewURLClient(u.config).QueryStat(u)
}

// Update returns a builder for updating this Url.
// Note that you need to call Url.Unwrap() before calling this method if this Url
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Url) Update() *URLUpdateOne {
	return NewURLClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Url entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Url) Unwrap() *Url {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Url is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Url) String() string {
	var builder strings.Builder
	builder.WriteString("Url(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("url=")
	builder.WriteString(u.URL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	if v := u.RelativePath; v != nil {
		builder.WriteString("relative_path=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Urls is a parsable slice of Url.
type Urls []*Url
