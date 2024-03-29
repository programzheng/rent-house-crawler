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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailtag"
)

// Rent591HomeDetailTagCreate is the builder for creating a Rent591HomeDetailTag entity.
type Rent591HomeDetailTagCreate struct {
	config
	mutation *Rent591HomeDetailTagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPostID sets the "post_id" field.
func (rdtc *Rent591HomeDetailTagCreate) SetPostID(i int) *Rent591HomeDetailTagCreate {
	rdtc.mutation.SetPostID(i)
	return rdtc
}

// SetValue sets the "value" field.
func (rdtc *Rent591HomeDetailTagCreate) SetValue(s string) *Rent591HomeDetailTagCreate {
	rdtc.mutation.SetValue(s)
	return rdtc
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdtc *Rent591HomeDetailTagCreate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailTagCreate {
	rdtc.mutation.AddRent591homeDetailIDs(ids...)
	return rdtc
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdtc *Rent591HomeDetailTagCreate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailTagCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdtc.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailTagMutation object of the builder.
func (rdtc *Rent591HomeDetailTagCreate) Mutation() *Rent591HomeDetailTagMutation {
	return rdtc.mutation
}

// Save creates the Rent591HomeDetailTag in the database.
func (rdtc *Rent591HomeDetailTagCreate) Save(ctx context.Context) (*Rent591HomeDetailTag, error) {
	return withHooks(ctx, rdtc.sqlSave, rdtc.mutation, rdtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdtc *Rent591HomeDetailTagCreate) SaveX(ctx context.Context) *Rent591HomeDetailTag {
	v, err := rdtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdtc *Rent591HomeDetailTagCreate) Exec(ctx context.Context) error {
	_, err := rdtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdtc *Rent591HomeDetailTagCreate) ExecX(ctx context.Context) {
	if err := rdtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdtc *Rent591HomeDetailTagCreate) check() error {
	if _, ok := rdtc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Rent591HomeDetailTag.post_id"`)}
	}
	if _, ok := rdtc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Rent591HomeDetailTag.value"`)}
	}
	return nil
}

func (rdtc *Rent591HomeDetailTagCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailTag, error) {
	if err := rdtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdtc.mutation.id = &_node.ID
	rdtc.mutation.done = true
	return _node, nil
}

func (rdtc *Rent591HomeDetailTagCreate) createSpec() (*Rent591HomeDetailTag, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailTag{config: rdtc.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailtag.Table, sqlgraph.NewFieldSpec(rent591homedetailtag.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdtc.conflict
	if value, ok := rdtc.mutation.PostID(); ok {
		_spec.SetField(rent591homedetailtag.FieldPostID, field.TypeInt, value)
		_node.PostID = value
	}
	if value, ok := rdtc.mutation.Value(); ok {
		_spec.SetField(rent591homedetailtag.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := rdtc.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailtag.Rent591homeDetailsTable,
			Columns: rent591homedetailtag.Rent591homeDetailsPrimaryKey,
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
//	client.Rent591HomeDetailTag.Create().
//		SetPostID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailTagUpsert) {
//			SetPostID(v+v).
//		}).
//		Exec(ctx)
func (rdtc *Rent591HomeDetailTagCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailTagUpsertOne {
	rdtc.conflict = opts
	return &Rent591HomeDetailTagUpsertOne{
		create: rdtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdtc *Rent591HomeDetailTagCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailTagUpsertOne {
	rdtc.conflict = append(rdtc.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailTagUpsertOne{
		create: rdtc,
	}
}

type (
	// Rent591HomeDetailTagUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailTag node.
	Rent591HomeDetailTagUpsertOne struct {
		create *Rent591HomeDetailTagCreate
	}

	// Rent591HomeDetailTagUpsert is the "OnConflict" setter.
	Rent591HomeDetailTagUpsert struct {
		*sql.UpdateSet
	}
)

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailTagUpsert) SetPostID(v int) *Rent591HomeDetailTagUpsert {
	u.Set(rent591homedetailtag.FieldPostID, v)
	return u
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsert) UpdatePostID() *Rent591HomeDetailTagUpsert {
	u.SetExcluded(rent591homedetailtag.FieldPostID)
	return u
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailTagUpsert) AddPostID(v int) *Rent591HomeDetailTagUpsert {
	u.Add(rent591homedetailtag.FieldPostID, v)
	return u
}

// SetValue sets the "value" field.
func (u *Rent591HomeDetailTagUpsert) SetValue(v string) *Rent591HomeDetailTagUpsert {
	u.Set(rent591homedetailtag.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsert) UpdateValue() *Rent591HomeDetailTagUpsert {
	u.SetExcluded(rent591homedetailtag.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailTagUpsertOne) UpdateNewValues() *Rent591HomeDetailTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailTagUpsertOne) Ignore() *Rent591HomeDetailTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailTagUpsertOne) DoNothing() *Rent591HomeDetailTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailTagCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailTagUpsertOne) Update(set func(*Rent591HomeDetailTagUpsert)) *Rent591HomeDetailTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailTagUpsertOne) SetPostID(v int) *Rent591HomeDetailTagUpsertOne {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.SetPostID(v)
	})
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailTagUpsertOne) AddPostID(v int) *Rent591HomeDetailTagUpsertOne {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.AddPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsertOne) UpdatePostID() *Rent591HomeDetailTagUpsertOne {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.UpdatePostID()
	})
}

// SetValue sets the "value" field.
func (u *Rent591HomeDetailTagUpsertOne) SetValue(v string) *Rent591HomeDetailTagUpsertOne {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsertOne) UpdateValue() *Rent591HomeDetailTagUpsertOne {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailTagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailTagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailTagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailTagUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailTagUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailTagCreateBulk is the builder for creating many Rent591HomeDetailTag entities in bulk.
type Rent591HomeDetailTagCreateBulk struct {
	config
	builders []*Rent591HomeDetailTagCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailTag entities in the database.
func (rdtcb *Rent591HomeDetailTagCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdtcb.builders))
	nodes := make([]*Rent591HomeDetailTag, len(rdtcb.builders))
	mutators := make([]Mutator, len(rdtcb.builders))
	for i := range rdtcb.builders {
		func(i int, root context.Context) {
			builder := rdtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailTagMutation)
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
					_, err = mutators[i+1].Mutate(root, rdtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rdtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdtcb *Rent591HomeDetailTagCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailTag {
	v, err := rdtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdtcb *Rent591HomeDetailTagCreateBulk) Exec(ctx context.Context) error {
	_, err := rdtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdtcb *Rent591HomeDetailTagCreateBulk) ExecX(ctx context.Context) {
	if err := rdtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailTag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailTagUpsert) {
//			SetPostID(v+v).
//		}).
//		Exec(ctx)
func (rdtcb *Rent591HomeDetailTagCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailTagUpsertBulk {
	rdtcb.conflict = opts
	return &Rent591HomeDetailTagUpsertBulk{
		create: rdtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdtcb *Rent591HomeDetailTagCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailTagUpsertBulk {
	rdtcb.conflict = append(rdtcb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailTagUpsertBulk{
		create: rdtcb,
	}
}

// Rent591HomeDetailTagUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailTag nodes.
type Rent591HomeDetailTagUpsertBulk struct {
	create *Rent591HomeDetailTagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailTagUpsertBulk) UpdateNewValues() *Rent591HomeDetailTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailTag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailTagUpsertBulk) Ignore() *Rent591HomeDetailTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailTagUpsertBulk) DoNothing() *Rent591HomeDetailTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailTagCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailTagUpsertBulk) Update(set func(*Rent591HomeDetailTagUpsert)) *Rent591HomeDetailTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailTagUpsertBulk) SetPostID(v int) *Rent591HomeDetailTagUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.SetPostID(v)
	})
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailTagUpsertBulk) AddPostID(v int) *Rent591HomeDetailTagUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.AddPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsertBulk) UpdatePostID() *Rent591HomeDetailTagUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.UpdatePostID()
	})
}

// SetValue sets the "value" field.
func (u *Rent591HomeDetailTagUpsertBulk) SetValue(v string) *Rent591HomeDetailTagUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *Rent591HomeDetailTagUpsertBulk) UpdateValue() *Rent591HomeDetailTagUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailTagUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailTagUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailTagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailTagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailTagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
