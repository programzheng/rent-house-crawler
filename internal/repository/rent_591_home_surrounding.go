package repository

import (
	"context"

	"github.com/programzheng/rent-house-crawler/ent"
)

func (rp *Repository) UpsertRent591SurroundingByRent591Home(ctx context.Context, r5h *ent.Rent591Home, r5si ent.Rent591HomeSurrounding) (*ent.Rent591HomeSurrounding, error) {
	id, err := rp.client.Rent591HomeSurrounding.
		Create().
		SetType(r5si.Type).
		SetDesc(r5si.Desc).
		SetDistance(r5si.Distance).
		SetRent591homes(r5h).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, err
	}

	r5s, err := rp.client.Rent591HomeSurrounding.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return r5s, nil
}
