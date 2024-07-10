package schema

import (
	"wsw/backend/domain/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Unique(),
		field.Enum("status").GoType(url.Status("pending")),
		field.Int("api_url_id").Nillable().Optional(),
		field.String("image"),
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return nil
}
