package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the UploadImage.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").NotEmpty().Immutable(),
		field.String("destination_path").NotEmpty().Immutable(),
		field.String("original_filename").NotEmpty().Immutable(),
		field.String("type").NotEmpty().Immutable(),
		field.String("extra_value").Nillable().Optional().Immutable(),
	}
}

// Edges of the UploadImage.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("imageprocess", ImageProcess.Type),
	}
}
