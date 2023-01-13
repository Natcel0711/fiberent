package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Usuario holds the schema definition for the Usuario entity.
type Usuario struct {
	ent.Schema
}

// Fields of the Usuario.
func (Usuario) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("email"),
		field.String("password"),
	}
}

// Edges of the Usuario.
func (Usuario) Edges() []ent.Edge {
	return nil
}
