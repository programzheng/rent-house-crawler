package repository

import (
	"context"

	"github.com/programzheng/rent-house-crawler/ent"
)

func (rp *Repository) UpsertRent591HomeDetailBreadcrumbsByRent591HomeDetailID(ctx context.Context, r5hdID int, r5hdbs *ent.Rent591HomeDetailBreadcrumbs) error {
	r5hdbcs := []*ent.Rent591HomeDetailBreadcrumbCreate{}
	for _, r5hdb := range *r5hdbs {
		create := rp.client.Rent591HomeDetailBreadcrumb.Create().
			SetName(r5hdb.Name).
			SetPostID(r5hdb.PostID).
			SetQuery(r5hdb.Query).
			SetLink(r5hdb.Link).
			AddRent591homeDetailIDs(r5hdID)
		r5hdbcs = append(r5hdbcs, create)
	}

	return rp.client.Rent591HomeDetailBreadcrumb.
		CreateBulk(r5hdbcs...).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
}
