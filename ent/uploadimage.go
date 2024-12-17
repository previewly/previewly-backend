// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wsw/backend/ent/uploadimage"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UploadImage is the model entity for the UploadImage schema.
type UploadImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// DestinationPath holds the value of the "destination_path" field.
	DestinationPath string `json:"destination_path,omitempty"`
	// OriginalFilename holds the value of the "original_filename" field.
	OriginalFilename string `json:"original_filename,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UploadImageQuery when eager-loading is set.
	Edges        UploadImageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UploadImageEdges holds the relations/edges for other nodes in the graph.
type UploadImageEdges struct {
	// Imageprocess holds the value of the imageprocess edge.
	Imageprocess []*ImageProcess `json:"imageprocess,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ImageprocessOrErr returns the Imageprocess value or an error if the edge
// was not loaded in eager-loading.
func (e UploadImageEdges) ImageprocessOrErr() ([]*ImageProcess, error) {
	if e.loadedTypes[0] {
		return e.Imageprocess, nil
	}
	return nil, &NotLoadedError{edge: "imageprocess"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UploadImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case uploadimage.FieldID:
			values[i] = new(sql.NullInt64)
		case uploadimage.FieldFilename, uploadimage.FieldDestinationPath, uploadimage.FieldOriginalFilename, uploadimage.FieldType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UploadImage fields.
func (ui *UploadImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case uploadimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ui.ID = int(value.Int64)
		case uploadimage.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				ui.Filename = value.String
			}
		case uploadimage.FieldDestinationPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_path", values[i])
			} else if value.Valid {
				ui.DestinationPath = value.String
			}
		case uploadimage.FieldOriginalFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_filename", values[i])
			} else if value.Valid {
				ui.OriginalFilename = value.String
			}
		case uploadimage.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ui.Type = value.String
			}
		default:
			ui.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UploadImage.
// This includes values selected through modifiers, order, etc.
func (ui *UploadImage) Value(name string) (ent.Value, error) {
	return ui.selectValues.Get(name)
}

// QueryImageprocess queries the "imageprocess" edge of the UploadImage entity.
func (ui *UploadImage) QueryImageprocess() *ImageProcessQuery {
	return NewUploadImageClient(ui.config).QueryImageprocess(ui)
}

// Update returns a builder for updating this UploadImage.
// Note that you need to call UploadImage.Unwrap() before calling this method if this UploadImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ui *UploadImage) Update() *UploadImageUpdateOne {
	return NewUploadImageClient(ui.config).UpdateOne(ui)
}

// Unwrap unwraps the UploadImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ui *UploadImage) Unwrap() *UploadImage {
	_tx, ok := ui.config.driver.(*txDriver)
	if !ok {
		panic("ent: UploadImage is not a transactional entity")
	}
	ui.config.driver = _tx.drv
	return ui
}

// String implements the fmt.Stringer.
func (ui *UploadImage) String() string {
	var builder strings.Builder
	builder.WriteString("UploadImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ui.ID))
	builder.WriteString("filename=")
	builder.WriteString(ui.Filename)
	builder.WriteString(", ")
	builder.WriteString("destination_path=")
	builder.WriteString(ui.DestinationPath)
	builder.WriteString(", ")
	builder.WriteString("original_filename=")
	builder.WriteString(ui.OriginalFilename)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ui.Type)
	builder.WriteByte(')')
	return builder.String()
}

// UploadImages is a parsable slice of UploadImage.
type UploadImages []*UploadImage
