package repository

import (
	"context"

	"github.com/programzheng/rent-house-crawler/ent"
)

func (rp *Repository) UpsertRent591HomeDetailPublishByRent591HomeDetailID(ctx context.Context, r5hdID int, r5hdp *ent.Rent591HomeDetailPublish) (int, error) {
	return rp.client.Rent591HomeDetailPublish.
		Create().
		SetPostID(r5hdp.PostID).
		SetName(r5hdp.Name).
		SetKey(r5hdp.Key).
		SetPostTime(r5hdp.PostTime).
		SetUpdateTime(r5hdp.UpdateTime).
		AddRent591homeDetailIDs(r5hdID).
		OnConflict().
		ID(ctx)
}
