package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailPublish holds the schema definition for the Rent591HomeDetailPublish entity.
type Rent591HomeDetailPublish struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailPublish.
func (Rent591HomeDetailPublish) Fields() []ent.Field {
	return []ent.Field{
		// original id
		field.Int("post_id"),
		field.String("name"),
		field.String("key"),
		field.String("post_time"),
		field.String("update_time"),
	}
}

// Edges of the Rent591HomeDetailPublish.
func (Rent591HomeDetailPublish) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_details", Rent591HomeDetail.Type).
			Ref("rent591home_detail_publishs"),
	}
}
