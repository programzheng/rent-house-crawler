package repository

import (
	"context"

	"github.com/programzheng/rent-house-crawler/ent"
)

func (rp *Repository) UpsertRent591HomeDetailPositionRoundByRent591HomeDetailID(ctx context.Context, r5hdID int, r5hdpr *ent.Rent591HomeDetailPositionRound) (int, error) {
	return rp.client.Rent591HomeDetailPositionRound.
		Create().
		SetTitle(r5hdpr.Title).
		SetKey(r5hdpr.Key).
		SetActive(r5hdpr.Active).
		SetCommunityName(r5hdpr.CommunityName).
		SetCommunityID(r5hdpr.CommunityID).
		SetAddress(r5hdpr.Address).
		SetLat(r5hdpr.Lat).
		SetLng(r5hdpr.Lng).
		SetRent591homeDetailsID(r5hdID).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
}
