package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stat holds the schema definition for the Stat entity.
type Stat struct {
	ent.Schema
}

// Fields of the Stat.
func (Stat) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.String("title").Nillable().Optional(),
	}
}

// Edges of the Stat.
func (Stat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).
			Unique().
			Required(),
	}
}
