package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeSurrounding holds the schema definition for the Rent591HomeSurrounding entity.
type Rent591HomeSurrounding struct {
	ent.Schema
}

// Fields of the Rent591HomeSurrounding.
func (Rent591HomeSurrounding) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("desc"),
		field.String("distance"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Rent591HomeSurrounding.
func (Rent591HomeSurrounding) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591homes", Rent591Home.Type).
			Ref("rent591home_surroundings").
			Unique(),
	}
}
