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
		field.JSON("processes", []types.ImageProcess{}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("path_prefix").Optional(),
		field.String("error").Optional(),
	}
}

// Edges of the ImageProcess.
func (ImageProcess) Edges() []ent.Edge {
	return nil
}
