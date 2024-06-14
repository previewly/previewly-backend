package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").
			Unique(),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return nil
}
