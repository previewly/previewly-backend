package schema

import (
	"time"

	"wsw/backend/ent/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ImageProcess holds the schema definition for the ImageProcess entity.
type ImageProcess struct {
	ent.Schema
}

// Fields of the ImageProcess.
func (ImageProcess) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").GoType(types.StatusEnum("pending")),
		field.String("process_hash"),
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
	return []ent.Edge{
		edge.From("uploadimage", UploadImage.Type).Ref("imageprocess").Unique(),
	}
}

// Indexes of the Street.
func (ImageProcess) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("process_hash").
			Edges("uploadimage").
			Unique(),
	}
}
