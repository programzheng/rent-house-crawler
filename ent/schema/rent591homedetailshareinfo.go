package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailShareInfo holds the schema definition for the Rent591HomeDetailShareInfo entity.
type Rent591HomeDetailShareInfo struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailShareInfo.
func (Rent591HomeDetailShareInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("From"),
		field.String("title"),
	}
}

// Edges of the Rent591HomeDetailShareInfo.
func (Rent591HomeDetailShareInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_details", Rent591HomeDetail.Type).
			Ref("rent591home_detail_shareinfos"),
	}
}
