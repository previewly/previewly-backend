package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ErrorResult holds the schema definition for the ErrorResult entity.
type ErrorResult struct {
	ent.Schema
}

// Fields of the ErrorResult.
func (ErrorResult) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.String("message").Nillable().Optional(),
	}
}

// Edges of the ErrorResult.
func (ErrorResult) Edges() []ent.Edge {
	return nil
}
