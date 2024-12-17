// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/types"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ImageProcess is the model entity for the ImageProcess schema.
type ImageProcess struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status types.StatusEnum `json:"status,omitempty"`
	// Process holds the value of the "process" field.
	Process types.ImageProcess `json:"process,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	upload_image_imageprocess *int
	selectValues              sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImageProcess) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case imageprocess.FieldID:
			values[i] = new(sql.NullInt64)
		case imageprocess.FieldStatus, imageprocess.FieldProcess:
			values[i] = new(sql.NullString)
		case imageprocess.FieldCreatedAt, imageprocess.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case imageprocess.ForeignKeys[0]: // upload_image_imageprocess
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImageProcess fields.
func (ip *ImageProcess) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case imageprocess.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ip.ID = int(value.Int64)
		case imageprocess.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ip.Status = types.StatusEnum(value.String)
			}
		case imageprocess.FieldProcess:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field process", values[i])
			} else if value.Valid {
				ip.Process = types.ImageProcess(value.String)
			}
		case imageprocess.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ip.CreatedAt = value.Time
			}
		case imageprocess.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ip.UpdatedAt = value.Time
			}
		case imageprocess.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field upload_image_imageprocess", value)
			} else if value.Valid {
				ip.upload_image_imageprocess = new(int)
				*ip.upload_image_imageprocess = int(value.Int64)
			}
		default:
			ip.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ImageProcess.
// This includes values selected through modifiers, order, etc.
func (ip *ImageProcess) Value(name string) (ent.Value, error) {
	return ip.selectValues.Get(name)
}

// Update returns a builder for updating this ImageProcess.
// Note that you need to call ImageProcess.Unwrap() before calling this method if this ImageProcess
// was returned from a transaction, and the transaction was committed or rolled back.
func (ip *ImageProcess) Update() *ImageProcessUpdateOne {
	return NewImageProcessClient(ip.config).UpdateOne(ip)
}

// Unwrap unwraps the ImageProcess entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ip *ImageProcess) Unwrap() *ImageProcess {
	_tx, ok := ip.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImageProcess is not a transactional entity")
	}
	ip.config.driver = _tx.drv
	return ip
}

// String implements the fmt.Stringer.
func (ip *ImageProcess) String() string {
	var builder strings.Builder
	builder.WriteString("ImageProcess(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ip.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ip.Status))
	builder.WriteString(", ")
	builder.WriteString("process=")
	builder.WriteString(fmt.Sprintf("%v", ip.Process))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ip.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ip.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ImageProcesses is a parsable slice of ImageProcess.
type ImageProcesses []*ImageProcess
