package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetailPositionRound holds the schema definition for the Rent591HomeDetailPositionRound entity.
type Rent591HomeDetailPositionRound struct {
	ent.Schema
}

// Fields of the Rent591HomeDetailPositionRound.
func (Rent591HomeDetailPositionRound) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("key"),
		field.Int("active"),
		field.String("community_name"),
		field.Int("community_id"),
		field.String("address"),
		field.Float("lat"),
		field.Float("lng"),
	}
}

// Edges of the Rent591HomeDetailPositionRound.
func (Rent591HomeDetailPositionRound) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_details", Rent591HomeDetail.Type).
			Ref("rent591home_detail_position_rounds").
			Unique(),
		edge.To("rent591home_detail_position_round_datas", Rent591HomeDetailPositionRoundData.Type).
			StorageKey(edge.Table("rent591home_d_p_r_rent591home_d_p_r_ds")),
	}
}
