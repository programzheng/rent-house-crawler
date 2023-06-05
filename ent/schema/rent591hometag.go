package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeTag holds the schema definition for the Rent591HomeTag entity.
type Rent591HomeTag struct {
	ent.Schema
}

// Fields of the Rent591HomeTag.
func (Rent591HomeTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Rent591HomeTag.
func (Rent591HomeTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rent591homes", Rent591Home.Type),
	}
}
