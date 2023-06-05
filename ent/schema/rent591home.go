package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591Home holds the schema definition for the Rent591Home entity.
type Rent591Home struct {
	ent.Schema
}

// Fields of the Rent591Home.
func (Rent591Home) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Int("type"),
		field.Int("post_id").Unique(),
		field.String("kind_name"),
		field.String("room_str"),
		field.String("floor_str"),
		field.String("community"),
		field.Int("price"),
		field.String("price_unit"),
		field.JSON("photo_list", []string{}).
			Optional(),
		// This data is optional because it is retrieved from another API later.
		field.String("region_name").Optional(),
		field.String("section_name"),
		field.String("street_name"),
		field.String("location"),
		field.String("area"),
		field.String("role_name"),
		field.String("contact"),
		field.String("refresh_time"),
		field.Int("yesterday_hit"),
		field.Int("is_vip"),
		field.Int("is_combine"),
		field.Int("hurry"),
		field.Int("is_social"),
		field.String("discount_price_str"),
		field.Int("cases_id"),
		field.Int("is_video"),
		field.Int("preferred"),
		field.Int("cid"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Rent591Home.
func (Rent591Home) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591home_tags", Rent591HomeTag.Type).
			Ref("rent591homes"),
		edge.To("rent591home_surroundings", Rent591HomeSurrounding.Type).
			Unique(),
		edge.To("rent591home_details", Rent591HomeDetail.Type).
			Unique(),
	}
}
