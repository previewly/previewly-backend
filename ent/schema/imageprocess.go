package schema

import (
	"time"

	"wsw/backend/ent/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ImageProcess holds the schema definition for the ImageProcess entity.
type ImageProcess struct {
	ent.Schema
}

// Fields of the ImageProcess.
func (ImageProcess) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").GoType(types.StatusEnum("pending")),
		field.Enum("process").GoType(types.ImageProcess("resize")),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ImageProcess.
func (ImageProcess) Edges() []ent.Edge {
	return nil
}
