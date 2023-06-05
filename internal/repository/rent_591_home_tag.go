package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/programzheng/rent-house-crawler/ent"
	"github.com/programzheng/rent-house-crawler/ent/rent591hometag"
)

func (rp *Repository) UpdateOrCreateRent591TagsByRent591Home(ctx context.Context, r5h *ent.Rent591Home, r5tsi ent.Rent591HomeTags) ([]*ent.Rent591HomeTag, error) {

	r5tcs := []*ent.Rent591HomeTagCreate{}
	for _, r5ti := range r5tsi {
		r5t, err := rp.client.Rent591HomeTag.Query().
			Where(
				sql.FieldEQ(rent591hometag.FieldName, r5ti.Name),
			).Only(ctx)

		if ent.IsNotFound(err) {
			r5tcs = append(r5tcs, rp.client.Rent591HomeTag.Create().SetName(r5ti.Name).AddRent591homes(r5h))
		} else {
			r5t.Update().AddRent591homes(r5h).Save(ctx)
		}
	}

	r5ts, err := rp.client.Rent591HomeTag.CreateBulk(r5tcs...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return r5ts, nil
}
