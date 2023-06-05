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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailnavdata"
)

// Rent591HomeDetailNavDataCreate is the builder for creating a Rent591HomeDetailNavData entity.
type Rent591HomeDetailNavDataCreate struct {
	config
	mutation *Rent591HomeDetailNavDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (rdndc *Rent591HomeDetailNavDataCreate) SetTitle(s string) *Rent591HomeDetailNavDataCreate {
	rdndc.mutation.SetTitle(s)
	return rdndc
}

// SetKey sets the "key" field.
func (rdndc *Rent591HomeDetailNavDataCreate) SetKey(s string) *Rent591HomeDetailNavDataCreate {
	rdndc.mutation.SetKey(s)
	return rdndc
}

// SetActive sets the "active" field.
func (rdndc *Rent591HomeDetailNavDataCreate) SetActive(i int) *Rent591HomeDetailNavDataCreate {
	rdndc.mutation.SetActive(i)
	return rdndc
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdndc *Rent591HomeDetailNavDataCreate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailNavDataCreate {
	rdndc.mutation.AddRent591homeDetailIDs(ids...)
	return rdndc
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdndc *Rent591HomeDetailNavDataCreate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailNavDataCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdndc.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailNavDataMutation object of the builder.
func (rdndc *Rent591HomeDetailNavDataCreate) Mutation() *Rent591HomeDetailNavDataMutation {
	return rdndc.mutation
}

// Save creates the Rent591HomeDetailNavData in the database.
func (rdndc *Rent591HomeDetailNavDataCreate) Save(ctx context.Context) (*Rent591HomeDetailNavData, error) {
	return withHooks(ctx, rdndc.sqlSave, rdndc.mutation, rdndc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdndc *Rent591HomeDetailNavDataCreate) SaveX(ctx context.Context) *Rent591HomeDetailNavData {
	v, err := rdndc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdndc *Rent591HomeDetailNavDataCreate) Exec(ctx context.Context) error {
	_, err := rdndc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdndc *Rent591HomeDetailNavDataCreate) ExecX(ctx context.Context) {
	if err := rdndc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdndc *Rent591HomeDetailNavDataCreate) check() error {
	if _, ok := rdndc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Rent591HomeDetailNavData.title"`)}
	}
	if _, ok := rdndc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Rent591HomeDetailNavData.key"`)}
	}
	if _, ok := rdndc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Rent591HomeDetailNavData.active"`)}
	}
	return nil
}

func (rdndc *Rent591HomeDetailNavDataCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailNavData, error) {
	if err := rdndc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdndc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdndc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdndc.mutation.id = &_node.ID
	rdndc.mutation.done = true
	return _node, nil
}

func (rdndc *Rent591HomeDetailNavDataCreate) createSpec() (*Rent591HomeDetailNavData, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailNavData{config: rdndc.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailnavdata.Table, sqlgraph.NewFieldSpec(rent591homedetailnavdata.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdndc.conflict
	if value, ok := rdndc.mutation.Title(); ok {
		_spec.SetField(rent591homedetailnavdata.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := rdndc.mutation.Key(); ok {
		_spec.SetField(rent591homedetailnavdata.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := rdndc.mutation.Active(); ok {
		_spec.SetField(rent591homedetailnavdata.FieldActive, field.TypeInt, value)
		_node.Active = value
	}
	if nodes := rdndc.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailnavdata.Rent591homeDetailsTable,
			Columns: rent591homedetailnavdata.Rent591homeDetailsPrimaryKey,
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
//	client.Rent591HomeDetailNavData.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailNavDataUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (rdndc *Rent591HomeDetailNavDataCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailNavDataUpsertOne {
	rdndc.conflict = opts
	return &Rent591HomeDetailNavDataUpsertOne{
		create: rdndc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdndc *Rent591HomeDetailNavDataCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailNavDataUpsertOne {
	rdndc.conflict = append(rdndc.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailNavDataUpsertOne{
		create: rdndc,
	}
}

type (
	// Rent591HomeDetailNavDataUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailNavData node.
	Rent591HomeDetailNavDataUpsertOne struct {
		create *Rent591HomeDetailNavDataCreate
	}

	// Rent591HomeDetailNavDataUpsert is the "OnConflict" setter.
	Rent591HomeDetailNavDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailNavDataUpsert) SetTitle(v string) *Rent591HomeDetailNavDataUpsert {
	u.Set(rent591homedetailnavdata.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsert) UpdateTitle() *Rent591HomeDetailNavDataUpsert {
	u.SetExcluded(rent591homedetailnavdata.FieldTitle)
	return u
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailNavDataUpsert) SetKey(v string) *Rent591HomeDetailNavDataUpsert {
	u.Set(rent591homedetailnavdata.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsert) UpdateKey() *Rent591HomeDetailNavDataUpsert {
	u.SetExcluded(rent591homedetailnavdata.FieldKey)
	return u
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailNavDataUpsert) SetActive(v int) *Rent591HomeDetailNavDataUpsert {
	u.Set(rent591homedetailnavdata.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsert) UpdateActive() *Rent591HomeDetailNavDataUpsert {
	u.SetExcluded(rent591homedetailnavdata.FieldActive)
	return u
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailNavDataUpsert) AddActive(v int) *Rent591HomeDetailNavDataUpsert {
	u.Add(rent591homedetailnavdata.FieldActive, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailNavDataUpsertOne) UpdateNewValues() *Rent591HomeDetailNavDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailNavDataUpsertOne) Ignore() *Rent591HomeDetailNavDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailNavDataUpsertOne) DoNothing() *Rent591HomeDetailNavDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailNavDataCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailNavDataUpsertOne) Update(set func(*Rent591HomeDetailNavDataUpsert)) *Rent591HomeDetailNavDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailNavDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailNavDataUpsertOne) SetTitle(v string) *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertOne) UpdateTitle() *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailNavDataUpsertOne) SetKey(v string) *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertOne) UpdateKey() *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateKey()
	})
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailNavDataUpsertOne) SetActive(v int) *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetActive(v)
	})
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailNavDataUpsertOne) AddActive(v int) *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.AddActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertOne) UpdateActive() *Rent591HomeDetailNavDataUpsertOne {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailNavDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailNavDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailNavDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailNavDataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailNavDataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailNavDataCreateBulk is the builder for creating many Rent591HomeDetailNavData entities in bulk.
type Rent591HomeDetailNavDataCreateBulk struct {
	config
	builders []*Rent591HomeDetailNavDataCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailNavData entities in the database.
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailNavData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdndcb.builders))
	nodes := make([]*Rent591HomeDetailNavData, len(rdndcb.builders))
	mutators := make([]Mutator, len(rdndcb.builders))
	for i := range rdndcb.builders {
		func(i int, root context.Context) {
			builder := rdndcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailNavDataMutation)
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
					_, err = mutators[i+1].Mutate(root, rdndcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdndcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdndcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rdndcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailNavData {
	v, err := rdndcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) Exec(ctx context.Context) error {
	_, err := rdndcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) ExecX(ctx context.Context) {
	if err := rdndcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailNavData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailNavDataUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailNavDataUpsertBulk {
	rdndcb.conflict = opts
	return &Rent591HomeDetailNavDataUpsertBulk{
		create: rdndcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdndcb *Rent591HomeDetailNavDataCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailNavDataUpsertBulk {
	rdndcb.conflict = append(rdndcb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailNavDataUpsertBulk{
		create: rdndcb,
	}
}

// Rent591HomeDetailNavDataUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailNavData nodes.
type Rent591HomeDetailNavDataUpsertBulk struct {
	create *Rent591HomeDetailNavDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailNavDataUpsertBulk) UpdateNewValues() *Rent591HomeDetailNavDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailNavData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailNavDataUpsertBulk) Ignore() *Rent591HomeDetailNavDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailNavDataUpsertBulk) DoNothing() *Rent591HomeDetailNavDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailNavDataCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailNavDataUpsertBulk) Update(set func(*Rent591HomeDetailNavDataUpsert)) *Rent591HomeDetailNavDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailNavDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *Rent591HomeDetailNavDataUpsertBulk) SetTitle(v string) *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertBulk) UpdateTitle() *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateTitle()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailNavDataUpsertBulk) SetKey(v string) *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertBulk) UpdateKey() *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateKey()
	})
}

// SetActive sets the "active" field.
func (u *Rent591HomeDetailNavDataUpsertBulk) SetActive(v int) *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.SetActive(v)
	})
}

// AddActive adds v to the "active" field.
func (u *Rent591HomeDetailNavDataUpsertBulk) AddActive(v int) *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.AddActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *Rent591HomeDetailNavDataUpsertBulk) UpdateActive() *Rent591HomeDetailNavDataUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailNavDataUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailNavDataUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailNavDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailNavDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailNavDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}