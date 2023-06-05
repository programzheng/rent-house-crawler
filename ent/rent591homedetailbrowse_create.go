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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailbrowse"
)

// Rent591HomeDetailBrowseCreate is the builder for creating a Rent591HomeDetailBrowse entity.
type Rent591HomeDetailBrowseCreate struct {
	config
	mutation *Rent591HomeDetailBrowseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPc sets the "pc" field.
func (rdbc *Rent591HomeDetailBrowseCreate) SetPc(i int) *Rent591HomeDetailBrowseCreate {
	rdbc.mutation.SetPc(i)
	return rdbc
}

// SetMobile sets the "mobile" field.
func (rdbc *Rent591HomeDetailBrowseCreate) SetMobile(i int) *Rent591HomeDetailBrowseCreate {
	rdbc.mutation.SetMobile(i)
	return rdbc
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdbc *Rent591HomeDetailBrowseCreate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailBrowseCreate {
	rdbc.mutation.AddRent591homeDetailIDs(ids...)
	return rdbc
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdbc *Rent591HomeDetailBrowseCreate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailBrowseCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdbc.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailBrowseMutation object of the builder.
func (rdbc *Rent591HomeDetailBrowseCreate) Mutation() *Rent591HomeDetailBrowseMutation {
	return rdbc.mutation
}

// Save creates the Rent591HomeDetailBrowse in the database.
func (rdbc *Rent591HomeDetailBrowseCreate) Save(ctx context.Context) (*Rent591HomeDetailBrowse, error) {
	return withHooks(ctx, rdbc.sqlSave, rdbc.mutation, rdbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdbc *Rent591HomeDetailBrowseCreate) SaveX(ctx context.Context) *Rent591HomeDetailBrowse {
	v, err := rdbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdbc *Rent591HomeDetailBrowseCreate) Exec(ctx context.Context) error {
	_, err := rdbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdbc *Rent591HomeDetailBrowseCreate) ExecX(ctx context.Context) {
	if err := rdbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdbc *Rent591HomeDetailBrowseCreate) check() error {
	if _, ok := rdbc.mutation.Pc(); !ok {
		return &ValidationError{Name: "pc", err: errors.New(`ent: missing required field "Rent591HomeDetailBrowse.pc"`)}
	}
	if _, ok := rdbc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "Rent591HomeDetailBrowse.mobile"`)}
	}
	return nil
}

func (rdbc *Rent591HomeDetailBrowseCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailBrowse, error) {
	if err := rdbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdbc.mutation.id = &_node.ID
	rdbc.mutation.done = true
	return _node, nil
}

func (rdbc *Rent591HomeDetailBrowseCreate) createSpec() (*Rent591HomeDetailBrowse, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailBrowse{config: rdbc.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailbrowse.Table, sqlgraph.NewFieldSpec(rent591homedetailbrowse.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdbc.conflict
	if value, ok := rdbc.mutation.Pc(); ok {
		_spec.SetField(rent591homedetailbrowse.FieldPc, field.TypeInt, value)
		_node.Pc = value
	}
	if value, ok := rdbc.mutation.Mobile(); ok {
		_spec.SetField(rent591homedetailbrowse.FieldMobile, field.TypeInt, value)
		_node.Mobile = value
	}
	if nodes := rdbc.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailbrowse.Rent591homeDetailsTable,
			Columns: rent591homedetailbrowse.Rent591homeDetailsPrimaryKey,
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
//	client.Rent591HomeDetailBrowse.Create().
//		SetPc(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailBrowseUpsert) {
//			SetPc(v+v).
//		}).
//		Exec(ctx)
func (rdbc *Rent591HomeDetailBrowseCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailBrowseUpsertOne {
	rdbc.conflict = opts
	return &Rent591HomeDetailBrowseUpsertOne{
		create: rdbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdbc *Rent591HomeDetailBrowseCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailBrowseUpsertOne {
	rdbc.conflict = append(rdbc.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailBrowseUpsertOne{
		create: rdbc,
	}
}

type (
	// Rent591HomeDetailBrowseUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailBrowse node.
	Rent591HomeDetailBrowseUpsertOne struct {
		create *Rent591HomeDetailBrowseCreate
	}

	// Rent591HomeDetailBrowseUpsert is the "OnConflict" setter.
	Rent591HomeDetailBrowseUpsert struct {
		*sql.UpdateSet
	}
)

// SetPc sets the "pc" field.
func (u *Rent591HomeDetailBrowseUpsert) SetPc(v int) *Rent591HomeDetailBrowseUpsert {
	u.Set(rent591homedetailbrowse.FieldPc, v)
	return u
}

// UpdatePc sets the "pc" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsert) UpdatePc() *Rent591HomeDetailBrowseUpsert {
	u.SetExcluded(rent591homedetailbrowse.FieldPc)
	return u
}

// AddPc adds v to the "pc" field.
func (u *Rent591HomeDetailBrowseUpsert) AddPc(v int) *Rent591HomeDetailBrowseUpsert {
	u.Add(rent591homedetailbrowse.FieldPc, v)
	return u
}

// SetMobile sets the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsert) SetMobile(v int) *Rent591HomeDetailBrowseUpsert {
	u.Set(rent591homedetailbrowse.FieldMobile, v)
	return u
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsert) UpdateMobile() *Rent591HomeDetailBrowseUpsert {
	u.SetExcluded(rent591homedetailbrowse.FieldMobile)
	return u
}

// AddMobile adds v to the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsert) AddMobile(v int) *Rent591HomeDetailBrowseUpsert {
	u.Add(rent591homedetailbrowse.FieldMobile, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailBrowseUpsertOne) UpdateNewValues() *Rent591HomeDetailBrowseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailBrowseUpsertOne) Ignore() *Rent591HomeDetailBrowseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailBrowseUpsertOne) DoNothing() *Rent591HomeDetailBrowseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailBrowseCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailBrowseUpsertOne) Update(set func(*Rent591HomeDetailBrowseUpsert)) *Rent591HomeDetailBrowseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailBrowseUpsert{UpdateSet: update})
	}))
	return u
}

// SetPc sets the "pc" field.
func (u *Rent591HomeDetailBrowseUpsertOne) SetPc(v int) *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.SetPc(v)
	})
}

// AddPc adds v to the "pc" field.
func (u *Rent591HomeDetailBrowseUpsertOne) AddPc(v int) *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.AddPc(v)
	})
}

// UpdatePc sets the "pc" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsertOne) UpdatePc() *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.UpdatePc()
	})
}

// SetMobile sets the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsertOne) SetMobile(v int) *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.SetMobile(v)
	})
}

// AddMobile adds v to the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsertOne) AddMobile(v int) *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.AddMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsertOne) UpdateMobile() *Rent591HomeDetailBrowseUpsertOne {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.UpdateMobile()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailBrowseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailBrowseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailBrowseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailBrowseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailBrowseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailBrowseCreateBulk is the builder for creating many Rent591HomeDetailBrowse entities in bulk.
type Rent591HomeDetailBrowseCreateBulk struct {
	config
	builders []*Rent591HomeDetailBrowseCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailBrowse entities in the database.
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailBrowse, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdbcb.builders))
	nodes := make([]*Rent591HomeDetailBrowse, len(rdbcb.builders))
	mutators := make([]Mutator, len(rdbcb.builders))
	for i := range rdbcb.builders {
		func(i int, root context.Context) {
			builder := rdbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailBrowseMutation)
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
					_, err = mutators[i+1].Mutate(root, rdbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rdbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailBrowse {
	v, err := rdbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) Exec(ctx context.Context) error {
	_, err := rdbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) ExecX(ctx context.Context) {
	if err := rdbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailBrowse.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailBrowseUpsert) {
//			SetPc(v+v).
//		}).
//		Exec(ctx)
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailBrowseUpsertBulk {
	rdbcb.conflict = opts
	return &Rent591HomeDetailBrowseUpsertBulk{
		create: rdbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdbcb *Rent591HomeDetailBrowseCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailBrowseUpsertBulk {
	rdbcb.conflict = append(rdbcb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailBrowseUpsertBulk{
		create: rdbcb,
	}
}

// Rent591HomeDetailBrowseUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailBrowse nodes.
type Rent591HomeDetailBrowseUpsertBulk struct {
	create *Rent591HomeDetailBrowseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailBrowseUpsertBulk) UpdateNewValues() *Rent591HomeDetailBrowseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailBrowse.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailBrowseUpsertBulk) Ignore() *Rent591HomeDetailBrowseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailBrowseUpsertBulk) DoNothing() *Rent591HomeDetailBrowseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailBrowseCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailBrowseUpsertBulk) Update(set func(*Rent591HomeDetailBrowseUpsert)) *Rent591HomeDetailBrowseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailBrowseUpsert{UpdateSet: update})
	}))
	return u
}

// SetPc sets the "pc" field.
func (u *Rent591HomeDetailBrowseUpsertBulk) SetPc(v int) *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.SetPc(v)
	})
}

// AddPc adds v to the "pc" field.
func (u *Rent591HomeDetailBrowseUpsertBulk) AddPc(v int) *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.AddPc(v)
	})
}

// UpdatePc sets the "pc" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsertBulk) UpdatePc() *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.UpdatePc()
	})
}

// SetMobile sets the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsertBulk) SetMobile(v int) *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.SetMobile(v)
	})
}

// AddMobile adds v to the "mobile" field.
func (u *Rent591HomeDetailBrowseUpsertBulk) AddMobile(v int) *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.AddMobile(v)
	})
}

// UpdateMobile sets the "mobile" field to the value that was provided on create.
func (u *Rent591HomeDetailBrowseUpsertBulk) UpdateMobile() *Rent591HomeDetailBrowseUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailBrowseUpsert) {
		s.UpdateMobile()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailBrowseUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailBrowseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailBrowseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailBrowseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}