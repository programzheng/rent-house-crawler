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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpublish"
)

// Rent591HomeDetailPublishCreate is the builder for creating a Rent591HomeDetailPublish entity.
type Rent591HomeDetailPublishCreate struct {
	config
	mutation *Rent591HomeDetailPublishMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPostID sets the "post_id" field.
func (rdpc *Rent591HomeDetailPublishCreate) SetPostID(i int) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.SetPostID(i)
	return rdpc
}

// SetName sets the "name" field.
func (rdpc *Rent591HomeDetailPublishCreate) SetName(s string) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.SetName(s)
	return rdpc
}

// SetKey sets the "key" field.
func (rdpc *Rent591HomeDetailPublishCreate) SetKey(s string) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.SetKey(s)
	return rdpc
}

// SetPostTime sets the "post_time" field.
func (rdpc *Rent591HomeDetailPublishCreate) SetPostTime(s string) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.SetPostTime(s)
	return rdpc
}

// SetUpdateTime sets the "update_time" field.
func (rdpc *Rent591HomeDetailPublishCreate) SetUpdateTime(s string) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.SetUpdateTime(s)
	return rdpc
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdpc *Rent591HomeDetailPublishCreate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailPublishCreate {
	rdpc.mutation.AddRent591homeDetailIDs(ids...)
	return rdpc
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdpc *Rent591HomeDetailPublishCreate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailPublishCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdpc.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPublishMutation object of the builder.
func (rdpc *Rent591HomeDetailPublishCreate) Mutation() *Rent591HomeDetailPublishMutation {
	return rdpc.mutation
}

// Save creates the Rent591HomeDetailPublish in the database.
func (rdpc *Rent591HomeDetailPublishCreate) Save(ctx context.Context) (*Rent591HomeDetailPublish, error) {
	return withHooks(ctx, rdpc.sqlSave, rdpc.mutation, rdpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rdpc *Rent591HomeDetailPublishCreate) SaveX(ctx context.Context) *Rent591HomeDetailPublish {
	v, err := rdpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdpc *Rent591HomeDetailPublishCreate) Exec(ctx context.Context) error {
	_, err := rdpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdpc *Rent591HomeDetailPublishCreate) ExecX(ctx context.Context) {
	if err := rdpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rdpc *Rent591HomeDetailPublishCreate) check() error {
	if _, ok := rdpc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Rent591HomeDetailPublish.post_id"`)}
	}
	if _, ok := rdpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Rent591HomeDetailPublish.name"`)}
	}
	if _, ok := rdpc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "Rent591HomeDetailPublish.key"`)}
	}
	if _, ok := rdpc.mutation.PostTime(); !ok {
		return &ValidationError{Name: "post_time", err: errors.New(`ent: missing required field "Rent591HomeDetailPublish.post_time"`)}
	}
	if _, ok := rdpc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Rent591HomeDetailPublish.update_time"`)}
	}
	return nil
}

func (rdpc *Rent591HomeDetailPublishCreate) sqlSave(ctx context.Context) (*Rent591HomeDetailPublish, error) {
	if err := rdpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rdpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rdpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rdpc.mutation.id = &_node.ID
	rdpc.mutation.done = true
	return _node, nil
}

func (rdpc *Rent591HomeDetailPublishCreate) createSpec() (*Rent591HomeDetailPublish, *sqlgraph.CreateSpec) {
	var (
		_node = &Rent591HomeDetailPublish{config: rdpc.config}
		_spec = sqlgraph.NewCreateSpec(rent591homedetailpublish.Table, sqlgraph.NewFieldSpec(rent591homedetailpublish.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rdpc.conflict
	if value, ok := rdpc.mutation.PostID(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostID, field.TypeInt, value)
		_node.PostID = value
	}
	if value, ok := rdpc.mutation.Name(); ok {
		_spec.SetField(rent591homedetailpublish.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rdpc.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpublish.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := rdpc.mutation.PostTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostTime, field.TypeString, value)
		_node.PostTime = value
	}
	if value, ok := rdpc.mutation.UpdateTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldUpdateTime, field.TypeString, value)
		_node.UpdateTime = value
	}
	if nodes := rdpc.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpublish.Rent591homeDetailsTable,
			Columns: rent591homedetailpublish.Rent591homeDetailsPrimaryKey,
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
//	client.Rent591HomeDetailPublish.Create().
//		SetPostID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailPublishUpsert) {
//			SetPostID(v+v).
//		}).
//		Exec(ctx)
func (rdpc *Rent591HomeDetailPublishCreate) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailPublishUpsertOne {
	rdpc.conflict = opts
	return &Rent591HomeDetailPublishUpsertOne{
		create: rdpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdpc *Rent591HomeDetailPublishCreate) OnConflictColumns(columns ...string) *Rent591HomeDetailPublishUpsertOne {
	rdpc.conflict = append(rdpc.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailPublishUpsertOne{
		create: rdpc,
	}
}

type (
	// Rent591HomeDetailPublishUpsertOne is the builder for "upsert"-ing
	//  one Rent591HomeDetailPublish node.
	Rent591HomeDetailPublishUpsertOne struct {
		create *Rent591HomeDetailPublishCreate
	}

	// Rent591HomeDetailPublishUpsert is the "OnConflict" setter.
	Rent591HomeDetailPublishUpsert struct {
		*sql.UpdateSet
	}
)

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailPublishUpsert) SetPostID(v int) *Rent591HomeDetailPublishUpsert {
	u.Set(rent591homedetailpublish.FieldPostID, v)
	return u
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsert) UpdatePostID() *Rent591HomeDetailPublishUpsert {
	u.SetExcluded(rent591homedetailpublish.FieldPostID)
	return u
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailPublishUpsert) AddPostID(v int) *Rent591HomeDetailPublishUpsert {
	u.Add(rent591homedetailpublish.FieldPostID, v)
	return u
}

// SetName sets the "name" field.
func (u *Rent591HomeDetailPublishUpsert) SetName(v string) *Rent591HomeDetailPublishUpsert {
	u.Set(rent591homedetailpublish.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsert) UpdateName() *Rent591HomeDetailPublishUpsert {
	u.SetExcluded(rent591homedetailpublish.FieldName)
	return u
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPublishUpsert) SetKey(v string) *Rent591HomeDetailPublishUpsert {
	u.Set(rent591homedetailpublish.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsert) UpdateKey() *Rent591HomeDetailPublishUpsert {
	u.SetExcluded(rent591homedetailpublish.FieldKey)
	return u
}

// SetPostTime sets the "post_time" field.
func (u *Rent591HomeDetailPublishUpsert) SetPostTime(v string) *Rent591HomeDetailPublishUpsert {
	u.Set(rent591homedetailpublish.FieldPostTime, v)
	return u
}

// UpdatePostTime sets the "post_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsert) UpdatePostTime() *Rent591HomeDetailPublishUpsert {
	u.SetExcluded(rent591homedetailpublish.FieldPostTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *Rent591HomeDetailPublishUpsert) SetUpdateTime(v string) *Rent591HomeDetailPublishUpsert {
	u.Set(rent591homedetailpublish.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsert) UpdateUpdateTime() *Rent591HomeDetailPublishUpsert {
	u.SetExcluded(rent591homedetailpublish.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailPublishUpsertOne) UpdateNewValues() *Rent591HomeDetailPublishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *Rent591HomeDetailPublishUpsertOne) Ignore() *Rent591HomeDetailPublishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailPublishUpsertOne) DoNothing() *Rent591HomeDetailPublishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailPublishCreate.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailPublishUpsertOne) Update(set func(*Rent591HomeDetailPublishUpsert)) *Rent591HomeDetailPublishUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailPublishUpsert{UpdateSet: update})
	}))
	return u
}

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailPublishUpsertOne) SetPostID(v int) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetPostID(v)
	})
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailPublishUpsertOne) AddPostID(v int) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.AddPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertOne) UpdatePostID() *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdatePostID()
	})
}

// SetName sets the "name" field.
func (u *Rent591HomeDetailPublishUpsertOne) SetName(v string) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertOne) UpdateName() *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateName()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPublishUpsertOne) SetKey(v string) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertOne) UpdateKey() *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateKey()
	})
}

// SetPostTime sets the "post_time" field.
func (u *Rent591HomeDetailPublishUpsertOne) SetPostTime(v string) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetPostTime(v)
	})
}

// UpdatePostTime sets the "post_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertOne) UpdatePostTime() *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdatePostTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *Rent591HomeDetailPublishUpsertOne) SetUpdateTime(v string) *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertOne) UpdateUpdateTime() *Rent591HomeDetailPublishUpsertOne {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailPublishUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailPublishCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailPublishUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Rent591HomeDetailPublishUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Rent591HomeDetailPublishUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Rent591HomeDetailPublishCreateBulk is the builder for creating many Rent591HomeDetailPublish entities in bulk.
type Rent591HomeDetailPublishCreateBulk struct {
	config
	builders []*Rent591HomeDetailPublishCreate
	conflict []sql.ConflictOption
}

// Save creates the Rent591HomeDetailPublish entities in the database.
func (rdpcb *Rent591HomeDetailPublishCreateBulk) Save(ctx context.Context) ([]*Rent591HomeDetailPublish, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rdpcb.builders))
	nodes := make([]*Rent591HomeDetailPublish, len(rdpcb.builders))
	mutators := make([]Mutator, len(rdpcb.builders))
	for i := range rdpcb.builders {
		func(i int, root context.Context) {
			builder := rdpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Rent591HomeDetailPublishMutation)
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
					_, err = mutators[i+1].Mutate(root, rdpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rdpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rdpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rdpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rdpcb *Rent591HomeDetailPublishCreateBulk) SaveX(ctx context.Context) []*Rent591HomeDetailPublish {
	v, err := rdpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rdpcb *Rent591HomeDetailPublishCreateBulk) Exec(ctx context.Context) error {
	_, err := rdpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdpcb *Rent591HomeDetailPublishCreateBulk) ExecX(ctx context.Context) {
	if err := rdpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Rent591HomeDetailPublish.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Rent591HomeDetailPublishUpsert) {
//			SetPostID(v+v).
//		}).
//		Exec(ctx)
func (rdpcb *Rent591HomeDetailPublishCreateBulk) OnConflict(opts ...sql.ConflictOption) *Rent591HomeDetailPublishUpsertBulk {
	rdpcb.conflict = opts
	return &Rent591HomeDetailPublishUpsertBulk{
		create: rdpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rdpcb *Rent591HomeDetailPublishCreateBulk) OnConflictColumns(columns ...string) *Rent591HomeDetailPublishUpsertBulk {
	rdpcb.conflict = append(rdpcb.conflict, sql.ConflictColumns(columns...))
	return &Rent591HomeDetailPublishUpsertBulk{
		create: rdpcb,
	}
}

// Rent591HomeDetailPublishUpsertBulk is the builder for "upsert"-ing
// a bulk of Rent591HomeDetailPublish nodes.
type Rent591HomeDetailPublishUpsertBulk struct {
	create *Rent591HomeDetailPublishCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *Rent591HomeDetailPublishUpsertBulk) UpdateNewValues() *Rent591HomeDetailPublishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Rent591HomeDetailPublish.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *Rent591HomeDetailPublishUpsertBulk) Ignore() *Rent591HomeDetailPublishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Rent591HomeDetailPublishUpsertBulk) DoNothing() *Rent591HomeDetailPublishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Rent591HomeDetailPublishCreateBulk.OnConflict
// documentation for more info.
func (u *Rent591HomeDetailPublishUpsertBulk) Update(set func(*Rent591HomeDetailPublishUpsert)) *Rent591HomeDetailPublishUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Rent591HomeDetailPublishUpsert{UpdateSet: update})
	}))
	return u
}

// SetPostID sets the "post_id" field.
func (u *Rent591HomeDetailPublishUpsertBulk) SetPostID(v int) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetPostID(v)
	})
}

// AddPostID adds v to the "post_id" field.
func (u *Rent591HomeDetailPublishUpsertBulk) AddPostID(v int) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.AddPostID(v)
	})
}

// UpdatePostID sets the "post_id" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertBulk) UpdatePostID() *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdatePostID()
	})
}

// SetName sets the "name" field.
func (u *Rent591HomeDetailPublishUpsertBulk) SetName(v string) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertBulk) UpdateName() *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateName()
	})
}

// SetKey sets the "key" field.
func (u *Rent591HomeDetailPublishUpsertBulk) SetKey(v string) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertBulk) UpdateKey() *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateKey()
	})
}

// SetPostTime sets the "post_time" field.
func (u *Rent591HomeDetailPublishUpsertBulk) SetPostTime(v string) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetPostTime(v)
	})
}

// UpdatePostTime sets the "post_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertBulk) UpdatePostTime() *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdatePostTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *Rent591HomeDetailPublishUpsertBulk) SetUpdateTime(v string) *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *Rent591HomeDetailPublishUpsertBulk) UpdateUpdateTime() *Rent591HomeDetailPublishUpsertBulk {
	return u.Update(func(s *Rent591HomeDetailPublishUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *Rent591HomeDetailPublishUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Rent591HomeDetailPublishCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Rent591HomeDetailPublishCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Rent591HomeDetailPublishUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
