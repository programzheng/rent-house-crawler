package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailPositionRoundDataChildren holds the schema definition for the Rent591HomeDetailPositionRoundDataChildren entity.
type Rent591HomeDetailPositionRoundDataChildren struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailPositionRoundDataChildren.
func (Rent591HomeDetailPositionRoundDataChildren) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("key"),
	}
}

// Edges of the Rent591HomeDetailPositionRoundDataChildren.
func (Rent591HomeDetailPositionRoundDataChildren) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_detail_position_round_datas", Rent591HomeDetailPositionRoundData.Type).
			Ref("rent591home_detail_position_round_data_childrens"),
	}
}
