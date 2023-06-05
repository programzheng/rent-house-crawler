// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetail"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionround"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionrounddata"
)

// Rent591HomeDetailPositionRoundCreate is the builder for creating a Rent591HomeDetailPositionRound entity.
type Rent591HomeDetailPositionRoundCreate struct {
	config
	mutation *Rent591HomeDetailPositionRoundMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetTitle(s string) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetTitle(s)
	return rdprc
}

// SetKey sets the "key" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetKey(s string) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetKey(s)
	return rdprc
}

// SetActive sets the "active" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetActive(i int) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetActive(i)
	return rdprc
}

// SetCommunityName sets the "community_name" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetCommunityName(s string) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetCommunityName(s)
	return rdprc
}

// SetCommunityID sets the "community_id" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetCommunityID(i int) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetCommunityID(i)
	return rdprc
}

// SetAddress sets the "address" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetAddress(s string) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetAddress(s)
	return rdprc
}

// SetLat sets the "lat" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetLat(f float64) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetLat(f)
	return rdprc
}

// SetLng sets the "lng" field.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetLng(f float64) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetLng(f)
	return rdprc
}

// SetRent591homeDetailsID sets the "rent591home_details" edge to the Rent591HomeDetail entity by ID.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetRent591homeDetailsID(id int) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.SetRent591homeDetailsID(id)
	return rdprc
}

// SetNillableRent591homeDetailsID sets the "rent591home_details" edge to the Rent591HomeDetail entity by ID if the given value is not nil.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetNillableRent591homeDetailsID(id *int) *Rent591HomeDetailPositionRoundCreate {
	if id != nil {
		rdprc = rdprc.SetRent591homeDetailsID(*id)
	}
	return rdprc
}

// SetRent591homeDetails sets the "rent591home_details" edge to the Rent591HomeDetail entity.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SetRent591homeDetails(r *Rent591HomeDetail) *Rent591HomeDetailPositionRoundCreate {
	return rdprc.SetRent591homeDetailsID(r.ID)
}

// AddRent591homeDetailPositionRoundDataIDs adds the "rent591home_detail_position_round_datas" edge to the Rent591HomeDetailPositionRoundData entity by IDs.
func (rdprc *Rent591HomeDetailPositionRoundCreate) AddRent591homeDetailPositionRoundDataIDs(ids ...int) *Rent591HomeDetailPositionRoundCreate {
	rdprc.mutation.AddRent591homeDetailPositionRoundDataIDs(ids...)
	return rdprc
}

// AddRent591homeDetailPositionRoundDatas adds the "rent591home_detail_position_round_datas" edges to the Rent591HomeDetailPositionRoundData entity.
func (rdprc *Rent591HomeDetailPositionRoundCreate) AddRent591homeDetailPositionRoundDatas(r ...*Rent591HomeDetailPositionRoundData) *Rent591HomeDetailPositionRoundCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprc.AddRent591homeDetailPositionRoundDataIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPositionRoundMutation object of the builder.
func (rdprc *Rent591HomeDetailPositionRoundCreate) Mutation() *Rent591HomeDetailPositionRoundMutation {
	return rdprc.mutation
}

// Save creates the Rent591HomeDetailPositionRound in the database.
func (rdprc *Rent591HomeDetailPositionRoundCreate) Save(ctx context.Context) (*Rent591HomeDetailPositionRound, error) {
	return withHooks(ctx, rdprc.sqlSave, rdprc.mutation, rdprc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdprc *Rent591HomeDetailPositionRoundCreate) SaveX(ctx context.Context) *Rent591HomeDetailPositionRound {
	v, err := rdprc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdprc *Rent591HomeDetailPositionRoundCreate) Exec(ctx context.Context) error {
	_, err := rdprc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprc *Rent591HomeDetailPositionRoundCreate) ExecX(ctx context.Context) {
	if err := rdprc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdprc *Rent591HomeDetailPositionRoundCreate) check() error {
	if _, ok := rdprc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.title"`)}
	}
	if _, ok := rdprc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.key"`)}
	}
	if _, ok := rdprc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.active"`)}
	}
	if _, ok := rdprc.mutation.CommunityName(); !ok {
		return &ValidationError{Name: "community_name", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.community_name"`)}
	}
	if _, ok := rdprc.mutation.CommunityID(); !ok {
		return &ValidationError{Name: "community_id", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.community_id"`)}
	}
	if _, ok := rdprc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.address"`)}
	}
	if _, ok := rdprc.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.lat"`)}
	}
	if _, ok := rdprc.mutation.Lng(); !ok {
		return &ValidationError{Name: "lng", err: errors.New(`ent: missing required field "Rent591HomeDetailPositionRound.lng"`)}
	}
	return nil
}

func (rdprc *Rent591HomeDetailPositionRoundCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailPositionRound, error) {
	if err := rdprc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdprc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdprc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdprc.mutation.id = &_node.ID
	rdprc.mutation.done = true
	return _node, nil
}

func (rdprc *Rent591HomeDetailPositionRoundCreate) createSpec() (*Rent591HomeDetailPositionRound, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailPositionRound{config: rdprc.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailpositionround.Table, sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdprc.conflict
	if value, ok := rdprc.mutation.Title(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := rdprc.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := rdprc.mutation.Active(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldActive, field.TypeInt, value)
		_node.Active = value
	}
	if value, ok := rdprc.mutation.CommunityName(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldCommunityName, field.TypeString, value)
		_node.CommunityName = value
	}
	if value, ok := rdprc.mutation.CommunityID(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldCommunityID, field.TypeInt, value)
		_node.CommunityID = value
	}
	if value, ok := rdprc.mutation.Address(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := rdprc.mutation.Lat(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldLat, field.TypeFloat64, value)
		_node.Lat = value
	}
	if value, ok := rdprc.mutation.Lng(); ok {
		_spec.SetField(rent591homedetailpositionround.FieldLng, field.TypeFloat64, value)
		_node.Lng = value
	}
	if nodes := rdprc.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent591homedetailpositionround.Rent591homeDetailsTable,
			Columns: []string{rent591homedetailpositionround.Rent591homeDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.rent591home_detail_rent591home_detail_position_rounds = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rdprc.mutation.Rent591homeDetailPositionRoundDatasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionround.Rent591homeDetailPositionRoundDatasTable,
			Columns: rent591homedetailpositionround.Rent591homeDetailPositionRoundDatasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailPositionRoundUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (rdprc *Rent591HomeDetailPositionRoundCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailPositionRoundUpsertOne {
	rdprc.conflict = opts
	return &Rent591HomeDetailPositionRoundUpsertOne{
		create: rdprc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdprc *Rent591HomeDetailPositionRoundCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailPositionRoundUpsertOne {
	rdprc.conflict = append(rdprc.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailPositionRoundUpsertOne{
		create: rdprc,
	}
}

type (
	// Rent591HomeDetailPositionRoundUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailPositionRound node.
	Rent591HomeDetailPositionRoundUpsertOne struct {
		create *Rent591HomeDetailPositionRoundCreate
	}

	// Rent591HomeDetailPositionRoundUpsert is the "OnConflict" setter.
	Rent591HomeDetailPositionRoundUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetTitle(v string) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateTitle() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldTitle)
	return u
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetKey(v string) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateKey() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldKey)
	return u
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetActive(v int) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateActive() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldActive)
	return u
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsert) AddActive(v int) *Rent591HomeDetailPositionRoundUpsert {
	u.Add(rent591homedetailpositionround.FieldActive, v)
	return u
}

// SetCommunityName sets the "community_name" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetCommunityName(v string) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldCommunityName, v)
	return u
}

// UpdateCommunityName sets the "community_name" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateCommunityName() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldCommunityName)
	return u
}

// SetCommunityID sets the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetCommunityID(v int) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldCommunityID, v)
	return u
}

// UpdateCommunityID sets the "community_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateCommunityID() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldCommunityID)
	return u
}

// AddCommunityID adds v to the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsert) AddCommunityID(v int) *Rent591HomeDetailPositionRoundUpsert {
	u.Add(rent591homedetailpositionround.FieldCommunityID, v)
	return u
}

// SetAddress sets the "address" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetAddress(v string) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateAddress() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldAddress)
	return u
}

// SetLat sets the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetLat(v float64) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldLat, v)
	return u
}

// UpdateLat sets the "lat" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateLat() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldLat)
	return u
}

// AddLat adds v to the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsert) AddLat(v float64) *Rent591HomeDetailPositionRoundUpsert {
	u.Add(rent591homedetailpositionround.FieldLat, v)
	return u
}

// SetLng sets the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsert) SetLng(v float64) *Rent591HomeDetailPositionRoundUpsert {
	u.Set(rent591homedetailpositionround.FieldLng, v)
	return u
}

// UpdateLng sets the "lng" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsert) UpdateLng() *Rent591HomeDetailPositionRoundUpsert {
	u.SetExcluded(rent591homedetailpositionround.FieldLng)
	return u
}

// AddLng adds v to the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsert) AddLng(v float64) *Rent591HomeDetailPositionRoundUpsert {
	u.Add(rent591homedetailpositionround.FieldLng, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateNewValues() *Rent591HomeDetailPositionRoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailPositionRoundUpsertOne) Ignore() *Rent591HomeDetailPositionRoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailPositionRoundUpsertOne) DoNothing() *Rent591HomeDetailPositionRoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailPositionRoundCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailPositionRoundUpsertOne) Update(set func(*Rent591HomeDetailPositionRoundUpsert)) *Rent591HomeDetailPositionRoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailPositionRoundUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetTitle(v string) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateTitle() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetKey(v string) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateKey() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateKey()
	})
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetActive(v int) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetActive(v)
	})
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) AddActive(v int) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateActive() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateActive()
	})
}

// SetCommunityName sets the "community_name" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetCommunityName(v string) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetCommunityName(v)
	})
}

// UpdateCommunityName sets the "community_name" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateCommunityName() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateCommunityName()
	})
}

// SetCommunityID sets the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetCommunityID(v int) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetCommunityID(v)
	})
}

// AddCommunityID adds v to the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) AddCommunityID(v int) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddCommunityID(v)
	})
}

// UpdateCommunityID sets the "community_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateCommunityID() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateCommunityID()
	})
}

// SetAddress sets the "address" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetAddress(v string) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateAddress() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateAddress()
	})
}

// SetLat sets the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetLat(v float64) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetLat(v)
	})
}

// AddLat adds v to the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) AddLat(v float64) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddLat(v)
	})
}

// UpdateLat sets the "lat" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateLat() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateLat()
	})
}

// SetLng sets the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) SetLng(v float64) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetLng(v)
	})
}

// AddLng adds v to the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsertOne) AddLng(v float64) *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddLng(v)
	})
}

// UpdateLng sets the "lng" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertOne) UpdateLng() *Rent591HomeDetailPositionRoundUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateLng()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailPositionRoundUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailPositionRoundCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailPositionRoundUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailPositionRoundUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailPositionRoundUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailPositionRoundCreateBulk is the builder for creating many Rent591HomeDetailPositionRound entities in bulk.
type Rent591HomeDetailPositionRoundCreateBulk struct {
	config
	builders []*Rent591HomeDetailPositionRoundCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailPositionRound entities in the database.
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailPositionRound, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdprcb.builders))
	nodes := make([]*Rent591HomeDetailPositionRound, len(rdprcb.builders))
	mutators := make([]Mutator, len(rdprcb.builders))
	for i := range rdprcb.builders {
		func(i int, root context.Context) {
			builder := rdprcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailPositionRoundMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rdprcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdprcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdprcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rdprcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailPositionRound {
	v, err := rdprcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) Exec(ctx context.Context) error {
	_, err := rdprcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) ExecX(ctx context.Context) {
	if err := rdprcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailPositionRound.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailPositionRoundUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailPositionRoundUpsertBulk {
	rdprcb.conflict = opts
	return &Rent591HomeDetailPositionRoundUpsertBulk{
		create: rdprcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdprcb *Rent591HomeDetailPositionRoundCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailPositionRoundUpsertBulk {
	rdprcb.conflict = append(rdprcb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailPositionRoundUpsertBulk{
		create: rdprcb,
	}
}

// Rent591HomeDetailPositionRoundUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailPositionRound nodes.
type Rent591HomeDetailPositionRoundUpsertBulk struct {
	create *Rent591HomeDetailPositionRoundCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateNewValues() *Rent591HomeDetailPositionRoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPositionRound.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailPositionRoundUpsertBulk) Ignore() *Rent591HomeDetailPositionRoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) DoNothing() *Rent591HomeDetailPositionRoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailPositionRoundCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) Update(set func(*Rent591HomeDetailPositionRoundUpsert)) *Rent591HomeDetailPositionRoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailPositionRoundUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetTitle(v string) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateTitle() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetKey(v string) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateKey() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateKey()
	})
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetActive(v int) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetActive(v)
	})
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) AddActive(v int) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateActive() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateActive()
	})
}

// SetCommunityName sets the "community_name" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetCommunityName(v string) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetCommunityName(v)
	})
}

// UpdateCommunityName sets the "community_name" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateCommunityName() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateCommunityName()
	})
}

// SetCommunityID sets the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetCommunityID(v int) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetCommunityID(v)
	})
}

// AddCommunityID adds v to the "community_id" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) AddCommunityID(v int) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddCommunityID(v)
	})
}

// UpdateCommunityID sets the "community_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateCommunityID() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateCommunityID()
	})
}

// SetAddress sets the "address" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetAddress(v string) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateAddress() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateAddress()
	})
}

// SetLat sets the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetLat(v float64) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetLat(v)
	})
}

// AddLat adds v to the "lat" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) AddLat(v float64) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddLat(v)
	})
}

// UpdateLat sets the "lat" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateLat() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateLat()
	})
}

// SetLng sets the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) SetLng(v float64) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.SetLng(v)
	})
}

// AddLng adds v to the "lng" field.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) AddLng(v float64) *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.AddLng(v)
	})
}

// UpdateLng sets the "lng" field to the value that was provided on create.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) UpdateLng() *Rent591HomeDetailPositionRoundUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPositionRoundUpsert) {
		s.UpdateLng()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailPositionRoundCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailPositionRoundCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailPositionRoundUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
