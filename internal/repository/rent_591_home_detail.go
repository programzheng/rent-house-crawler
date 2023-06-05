package repository

import (
	"context"

	"github.com/programzheng/rent-house-crawler/ent"
)

func (rp *Repository) UpsertRent591HomeDetailByRent591Home(ctx context.Context, r5h *ent.Rent591Home, r5hd *ent.Rent591HomeDetail) (int, error) {
	return rp.client.Rent591HomeDetail.
		Create().
		SetTitle(r5hd.Title).
		SetDeposit(r5hd.Deposit).
		SetKind(r5hd.Kind).
		SetRelieved(r5hd.Relieved).
		SetRegionID(r5hd.RegionID).
		SetSectionID(r5hd.SectionID).
		SetDealText(r5hd.DealText).
		SetDealTime(r5hd.DealTime).
		SetPrice(r5hd.Price).
		SetPriceUnit(r5hd.PriceUnit).
		SetRent591homes(r5h).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
}

func (rp *Repository) UpsertRent591HomesDetailByRent591Home(ctx context.Context, r5h *ent.Rent591Home, r5hds *ent.Rent591HomeDetails) error {
	r5hdcs := []*ent.Rent591HomeDetailCreate{}
	for _, r5hd := range *r5hds {
		create := rp.client.Rent591HomeDetail.Create().
			SetTitle(r5hd.Title).
			SetDeposit(r5hd.Deposit).
			SetKind(r5hd.Kind).
			SetRelieved(r5hd.Relieved).
			SetRegionID(r5hd.RegionID).
			SetSectionID(r5hd.SectionID).
			SetDealText(r5hd.DealText).
			SetDealTime(r5hd.DealTime).
			SetPrice(r5hd.Price).
			SetPriceUnit(r5hd.PriceUnit).
			SetRent591homes(r5h)
		r5hdcs = append(r5hdcs, create)
	}

	return rp.client.Rent591HomeDetail.
		CreateBulk(r5hdcs...).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
}
