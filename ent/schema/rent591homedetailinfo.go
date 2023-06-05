package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailInfo holds the schema definition for the Rent591HomeDetailInfo entity.
type Rent591HomeDetailInfo struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailInfo.
func (Rent591HomeDetailInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
		field.Int("key"),
	}
}

// Edges of the Rent591HomeDetailInfo.
func (Rent591HomeDetailInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rent591home_details", Rent591HomeDetail.Type),
	}
}
