package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UploadImage holds the schema definition for the UploadImage entity.
type UploadImage struct {
	ent.Schema
}

// Fields of the UploadImage.
func (UploadImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").NotEmpty().Immutable(),
		field.String("destination_path").NotEmpty().Immutable(),
		field.String("original_filename").NotEmpty().Immutable(),
		field.String("type").NotEmpty().Immutable(),
	}
}

// Edges of the UploadImage.
func (UploadImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("imageprocess", ImageProcess.Type),
	}
}
