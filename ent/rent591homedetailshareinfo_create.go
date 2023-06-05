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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailshareinfo"
)

// Rent591HomeDetailShareInfoCreate is the builder for creating a Rent591HomeDetailShareInfo entity.
type Rent591HomeDetailShareInfoCreate struct {
	config
	mutation *Rent591HomeDetailShareInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetURL sets the "url" field.
func (rdsic *Rent591HomeDetailShareInfoCreate) SetURL(s string) *Rent591HomeDetailShareInfoCreate {
	rdsic.mutation.SetURL(s)
	return rdsic
}

// SetFrom sets the "From" field.
func (rdsic *Rent591HomeDetailShareInfoCreate) SetFrom(s string) *Rent591HomeDetailShareInfoCreate {
	rdsic.mutation.SetFrom(s)
	return rdsic
}

// SetTitle sets the "title" field.
func (rdsic *Rent591HomeDetailShareInfoCreate) SetTitle(s string) *Rent591HomeDetailShareInfoCreate {
	rdsic.mutation.SetTitle(s)
	return rdsic
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdsic *Rent591HomeDetailShareInfoCreate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailShareInfoCreate {
	rdsic.mutation.AddRent591homeDetailIDs(ids...)
	return rdsic
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdsic *Rent591HomeDetailShareInfoCreate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailShareInfoCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdsic.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailShareInfoMutation object of the builder.
func (rdsic *Rent591HomeDetailShareInfoCreate) Mutation() *Rent591HomeDetailShareInfoMutation {
	return rdsic.mutation
}

// Save creates the Rent591HomeDetailShareInfo in the database.
func (rdsic *Rent591HomeDetailShareInfoCreate) Save(ctx context.Context) (*Rent591HomeDetailShareInfo, error) {
	return withHooks(ctx, rdsic.sqlSave, rdsic.mutation, rdsic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdsic *Rent591HomeDetailShareInfoCreate) SaveX(ctx context.Context) *Rent591HomeDetailShareInfo {
	v, err := rdsic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdsic *Rent591HomeDetailShareInfoCreate) Exec(ctx context.Context) error {
	_, err := rdsic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdsic *Rent591HomeDetailShareInfoCreate) ExecX(ctx context.Context) {
	if err := rdsic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdsic *Rent591HomeDetailShareInfoCreate) check() error {
	if _, ok := rdsic.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Rent591HomeDetailShareInfo.url"`)}
	}
	if _, ok := rdsic.mutation.From(); !ok {
		return &ValidationError{Name: "From", err: errors.New(`ent: missing required field "Rent591HomeDetailShareInfo.From"`)}
	}
	if _, ok := rdsic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Rent591HomeDetailShareInfo.title"`)}
	}
	return nil
}

func (rdsic *Rent591HomeDetailShareInfoCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailShareInfo, error) {
	if err := rdsic.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdsic.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdsic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdsic.mutation.id = &_node.ID
	rdsic.mutation.done = true
	return _node, nil
}

func (rdsic *Rent591HomeDetailShareInfoCreate) createSpec() (*Rent591HomeDetailShareInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailShareInfo{config: rdsic.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailshareinfo.Table, sqlgraph.NewFieldSpec(rent591homedetailshareinfo.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdsic.conflict
	if value, ok := rdsic.mutation.URL(); ok {
		_spec.SetField(rent591homedetailshareinfo.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := rdsic.mutation.From(); ok {
		_spec.SetField(rent591homedetailshareinfo.FieldFrom, field.TypeString, value)
		_node.From = value
	}
	if value, ok := rdsic.mutation.Title(); ok {
		_spec.SetField(rent591homedetailshareinfo.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if nodes := rdsic.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailshareinfo.Rent591homeDetailsTable,
			Columns: rent591homedetailshareinfo.Rent591homeDetailsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetail.FieldID, field.TypeInt),
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
//	client.Rent591HomeDetailShareInfo.Create().
//		SetURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailShareInfoUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (rdsic *Rent591HomeDetailShareInfoCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailShareInfoUpsertOne {
	rdsic.conflict = opts
	return &Rent591HomeDetailShareInfoUpsertOne{
		create: rdsic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdsic *Rent591HomeDetailShareInfoCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailShareInfoUpsertOne {
	rdsic.conflict = append(rdsic.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailShareInfoUpsertOne{
		create: rdsic,
	}
}

type (
	// Rent591HomeDetailShareInfoUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailShareInfo node.
	Rent591HomeDetailShareInfoUpsertOne struct {
		create *Rent591HomeDetailShareInfoCreate
	}

	// Rent591HomeDetailShareInfoUpsert is the "OnConflict" setter.
	Rent591HomeDetailShareInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetURL sets the "url" field.
func (u *Rent591HomeDetailShareInfoUpsert) SetURL(v string) *Rent591HomeDetailShareInfoUpsert {
	u.Set(rent591homedetailshareinfo.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsert) UpdateURL() *Rent591HomeDetailShareInfoUpsert {
	u.SetExcluded(rent591homedetailshareinfo.FieldURL)
	return u
}

// SetFrom sets the "From" field.
func (u *Rent591HomeDetailShareInfoUpsert) SetFrom(v string) *Rent591HomeDetailShareInfoUpsert {
	u.Set(rent591homedetailshareinfo.FieldFrom, v)
	return u
}

// UpdateFrom sets the "From" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsert) UpdateFrom() *Rent591HomeDetailShareInfoUpsert {
	u.SetExcluded(rent591homedetailshareinfo.FieldFrom)
	return u
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailShareInfoUpsert) SetTitle(v string) *Rent591HomeDetailShareInfoUpsert {
	u.Set(rent591homedetailshareinfo.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsert) UpdateTitle() *Rent591HomeDetailShareInfoUpsert {
	u.SetExcluded(rent591homedetailshareinfo.FieldTitle)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailShareInfoUpsertOne) UpdateNewValues() *Rent591HomeDetailShareInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailShareInfoUpsertOne) Ignore() *Rent591HomeDetailShareInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailShareInfoUpsertOne) DoNothing() *Rent591HomeDetailShareInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailShareInfoCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailShareInfoUpsertOne) Update(set func(*Rent591HomeDetailShareInfoUpsert)) *Rent591HomeDetailShareInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailShareInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *Rent591HomeDetailShareInfoUpsertOne) SetURL(v string) *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertOne) UpdateURL() *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateURL()
	})
}

// SetFrom sets the "From" field.
func (u *Rent591HomeDetailShareInfoUpsertOne) SetFrom(v string) *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "From" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertOne) UpdateFrom() *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateFrom()
	})
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailShareInfoUpsertOne) SetTitle(v string) *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertOne) UpdateTitle() *Rent591HomeDetailShareInfoUpsertOne {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateTitle()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailShareInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailShareInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailShareInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailShareInfoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailShareInfoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailShareInfoCreateBulk is the builder for creating many Rent591HomeDetailShareInfo entities in bulk.
type Rent591HomeDetailShareInfoCreateBulk struct {
	config
	builders []*Rent591HomeDetailShareInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailShareInfo entities in the database.
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailShareInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdsicb.builders))
	nodes := make([]*Rent591HomeDetailShareInfo, len(rdsicb.builders))
	mutators := make([]Mutator, len(rdsicb.builders))
	for i := range rdsicb.builders {
		func(i int, root context.Context) {
			builder := rdsicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailShareInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, rdsicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdsicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdsicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rdsicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailShareInfo {
	v, err := rdsicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := rdsicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) ExecX(ctx context.Context) {
	if err := rdsicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailShareInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailShareInfoUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailShareInfoUpsertBulk {
	rdsicb.conflict = opts
	return &Rent591HomeDetailShareInfoUpsertBulk{
		create: rdsicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdsicb *Rent591HomeDetailShareInfoCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailShareInfoUpsertBulk {
	rdsicb.conflict = append(rdsicb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailShareInfoUpsertBulk{
		create: rdsicb,
	}
}

// Rent591HomeDetailShareInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailShareInfo nodes.
type Rent591HomeDetailShareInfoUpsertBulk struct {
	create *Rent591HomeDetailShareInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailShareInfoUpsertBulk) UpdateNewValues() *Rent591HomeDetailShareInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailShareInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailShareInfoUpsertBulk) Ignore() *Rent591HomeDetailShareInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailShareInfoUpsertBulk) DoNothing() *Rent591HomeDetailShareInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailShareInfoCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailShareInfoUpsertBulk) Update(set func(*Rent591HomeDetailShareInfoUpsert)) *Rent591HomeDetailShareInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailShareInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *Rent591HomeDetailShareInfoUpsertBulk) SetURL(v string) *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertBulk) UpdateURL() *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateURL()
	})
}

// SetFrom sets the "From" field.
func (u *Rent591HomeDetailShareInfoUpsertBulk) SetFrom(v string) *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "From" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertBulk) UpdateFrom() *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateFrom()
	})
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailShareInfoUpsertBulk) SetTitle(v string) *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailShareInfoUpsertBulk) UpdateTitle() *Rent591HomeDetailShareInfoUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailShareInfoUpsert) {
		s.UpdateTitle()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailShareInfoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailShareInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailShareInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailShareInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}