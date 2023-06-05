package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailPositionRoundData holds the schema definition for the Rent591HomeDetailPositionRoundData entity.
type Rent591HomeDetailPositionRoundData struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailPositionRoundData.
func (Rent591HomeDetailPositionRoundData) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("key"),
	}
}

// Edges of the Rent591HomeDetailPositionRoundData.
func (Rent591HomeDetailPositionRoundData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_detail_position_rounds", Rent591HomeDetailPositionRound.Type).
			Ref("rent591home_detail_position_round_datas"),
		edge.To("rent591home_detail_position_round_data_childrens", Rent591HomeDetailPositionRoundDataChildren.Type).
			StorageKey(edge.Table("rent591home_d_p_r_ds_rent591home_d_p_r_d_cs")),
	}
}
