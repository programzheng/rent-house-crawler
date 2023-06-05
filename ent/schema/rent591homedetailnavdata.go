package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailNavData holds the schema definition for the Rent591HomeDetailNavData entity.
type Rent591HomeDetailNavData struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailNavData.
func (Rent591HomeDetailNavData) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("key"),
		field.Int("active"),
	}
}

// Edges of the Rent591HomeDetailNavData.
func (Rent591HomeDetailNavData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rent591home_details", Rent591HomeDetail.Type),
	}
}
