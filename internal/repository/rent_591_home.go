package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent"
	"github.com/programzheng/rent-house-crawler/ent/rent591home"
)

func (rp *Repository) UpsertRent591Home(ctx context.Context, r5hi ent.Rent591Home) (*ent.Rent591Home, error) {

	id, err := rp.client.Rent591Home.
		Create().
		SetTitle(r5hi.Title).
		SetType(r5hi.Type).
		SetPostID(r5hi.PostID).
		SetKindName(r5hi.KindName).
		SetRoomStr(r5hi.RoomStr).
		SetFloorStr(r5hi.FloorStr).
		SetCommunity(r5hi.Community).
		SetPrice(r5hi.Price).
		SetPriceUnit(r5hi.PriceUnit).
		SetPhotoList(r5hi.PhotoList).
		SetRegionName(r5hi.RegionName).
		SetSectionName(r5hi.SectionName).
		SetStreetName(r5hi.StreetName).
		SetLocation(r5hi.Location).
		SetArea(r5hi.Area).
		SetRoleName(r5hi.RoleName).
		SetContact(r5hi.Contact).
		SetRefreshTime(r5hi.RefreshTime).
		SetYesterdayHit(r5hi.YesterdayHit).
		SetIsVip(r5hi.IsVip).
		SetIsCombine(r5hi.IsCombine).
		SetHurry(r5hi.Hurry).
		SetIsSocial(r5hi.IsSocial).
		SetDiscountPriceStr(r5hi.DiscountPriceStr).
		SetCasesID(r5hi.CasesID).
		SetIsVideo(r5hi.IsVideo).
		SetPreferred(r5hi.Preferred).
		SetCid(r5hi.Cid).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, err
	}

	r5h, err := rp.client.Rent591Home.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return r5h, nil
}

func (rp *Repository) GetRent591HomePostIDs(ctx context.Context) ([]*ent.Rent591Home, error) {
	r5hs, err := rp.client.Rent591Home.Query().
		Select(rent591home.FieldPostID).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return r5hs, nil
}

func (rp *Repository) SetRent591HomeRegionByID(ctx context.Context, id int, regionName string) error {
	err := rp.client.Rent591Home.
		UpdateOneID(id).
		SetRegionName(regionName).
		Exec(ctx)

	return err
}

func (rp *Repository) DeleteRent591HomeByID(ctx context.Context, id int) error {
	return rp.client.Rent591Home.
		DeleteOneID(id).
		Exec(ctx)
}

func (rp *Repository) DeleteRent591HomeByPostID(ctx context.Context, postID int) (int, error) {
	return rp.client.Rent591Home.
		Delete().
		Where(
			sql.FieldEQ(rent591home.FieldPostID, postID),
		).
		Exec(ctx)
}
