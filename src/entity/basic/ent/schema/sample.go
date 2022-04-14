package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Sample holds the schema definition for the Sample entity.
type Sample struct {
	ent.Schema
}

// Fields of the Sample.
func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("name").Default("sample"),
	}
}

// Edges of the Sample.
func (Sample) Edges() []ent.Edge {
	return nil
}
