package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rent591HomeDetail holds the schema definition for the Rent591HomeDetail entity.
type Rent591HomeDetail struct {
	ent.Schema
}

// Fields of the Rent591HomeDetail.
func (Rent591HomeDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("deposit"),
		field.Int("kind"),
		field.Int("relieved"),
		field.Int("region_id"),
		field.Int("section_id"),
		field.String("deal_text").Optional(),
		field.Int("deal_time"),
		field.Int("price"),
		field.String("price_unit"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Rent591HomeDetail.
func (Rent591HomeDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rent591homes", Rent591Home.Type).
			Ref("rent591home_details").
			Unique(),
		edge.From("rent591home_detail_breadcrumbs", Rent591HomeDetailBreadcrumb.Type).
			Ref("rent591home_details"),
		edge.To("rent591home_detail_shareinfos", Rent591HomeDetailShareInfo.Type),
		edge.To("rent591home_detail_browses", Rent591HomeDetailBrowse.Type),
		edge.From("rent591home_detail_tags", Rent591HomeDetailTag.Type).
			Ref("rent591home_details"),
		edge.From("rent591home_detail_navdatas", Rent591HomeDetailNavData.Type).
			Ref("rent591home_details"),
		edge.From("rent591home_detail_infos", Rent591HomeDetailInfo.Type).
			Ref("rent591home_details"),
		edge.To("rent591home_detail_publishs", Rent591HomeDetailPublish.Type),
		edge.To("rent591home_detail_position_rounds", Rent591HomeDetailPositionRound.Type).
			Unique(),
	}
}
