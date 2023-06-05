package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailBreadcrumb holds the schema definition for the Rent591HomeDetailBreadcrumb entity.
type Rent591HomeDetailBreadcrumb struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailBreadcrumb.
func (Rent591HomeDetailBreadcrumb) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("post_id"),
		field.String("query"),
		field.String("link"),
	}
}

// Edges of the Rent591HomeDetailBreadcrumb.
func (Rent591HomeDetailBreadcrumb) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rent591home_details", Rent591HomeDetail.Type),
	}
}
