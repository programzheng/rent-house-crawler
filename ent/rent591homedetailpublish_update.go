// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetail"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpublish"
)

// Rent591HomeDetailPublishUpdate is the builder for updating Rent591HomeDetailPublish entities.
type Rent591HomeDetailPublishUpdate struct {
	config
	hooks    []Hook
	mutation *Rent591HomeDetailPublishMutation
}

// Where appends a list predicates to the Rent591HomeDetailPublishUpdate builder.
func (rdpu *Rent591HomeDetailPublishUpdate) Where(ps ...predicate.Rent591HomeDetailPublish) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.Where(ps...)
	return rdpu
}

// SetPostID sets the "post_id" field.
func (rdpu *Rent591HomeDetailPublishUpdate) SetPostID(i int) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.ResetPostID()
	rdpu.mutation.SetPostID(i)
	return rdpu
}

// AddPostID adds i to the "post_id" field.
func (rdpu *Rent591HomeDetailPublishUpdate) AddPostID(i int) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.AddPostID(i)
	return rdpu
}

// SetName sets the "name" field.
func (rdpu *Rent591HomeDetailPublishUpdate) SetName(s string) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.SetName(s)
	return rdpu
}

// SetKey sets the "key" field.
func (rdpu *Rent591HomeDetailPublishUpdate) SetKey(s string) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.SetKey(s)
	return rdpu
}

// SetPostTime sets the "post_time" field.
func (rdpu *Rent591HomeDetailPublishUpdate) SetPostTime(s string) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.SetPostTime(s)
	return rdpu
}

// SetUpdateTime sets the "update_time" field.
func (rdpu *Rent591HomeDetailPublishUpdate) SetUpdateTime(s string) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.SetUpdateTime(s)
	return rdpu
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdpu *Rent591HomeDetailPublishUpdate) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.AddRent591homeDetailIDs(ids...)
	return rdpu
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdpu *Rent591HomeDetailPublishUpdate) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailPublishUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdpu.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPublishMutation object of the builder.
func (rdpu *Rent591HomeDetailPublishUpdate) Mutation() *Rent591HomeDetailPublishMutation {
	return rdpu.mutation
}

// ClearRent591homeDetails clears all "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdpu *Rent591HomeDetailPublishUpdate) ClearRent591homeDetails() *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.ClearRent591homeDetails()
	return rdpu
}

// RemoveRent591homeDetailIDs removes the "rent591home_details" edge to Rent591HomeDetail entities by IDs.
func (rdpu *Rent591HomeDetailPublishUpdate) RemoveRent591homeDetailIDs(ids ...int) *Rent591HomeDetailPublishUpdate {
	rdpu.mutation.RemoveRent591homeDetailIDs(ids...)
	return rdpu
}

// RemoveRent591homeDetails removes "rent591home_details" edges to Rent591HomeDetail entities.
func (rdpu *Rent591HomeDetailPublishUpdate) RemoveRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailPublishUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdpu.RemoveRent591homeDetailIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rdpu *Rent591HomeDetailPublishUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rdpu.sqlSave, rdpu.mutation, rdpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdpu *Rent591HomeDetailPublishUpdate) SaveX(ctx context.Context) int {
	affected, err := rdpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rdpu *Rent591HomeDetailPublishUpdate) Exec(ctx context.Context) error {
	_, err := rdpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdpu *Rent591HomeDetailPublishUpdate) ExecX(ctx context.Context) {
	if err := rdpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rdpu *Rent591HomeDetailPublishUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591homedetailpublish.Table, rent591homedetailpublish.Columns, sqlgraph.NewFieldSpec(rent591homedetailpublish.FieldID, field.TypeInt))
	if ps := rdpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdpu.mutation.PostID(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostID, field.TypeInt, value)
	}
	if value, ok := rdpu.mutation.AddedPostID(); ok {
		_spec.AddField(rent591homedetailpublish.FieldPostID, field.TypeInt, value)
	}
	if value, ok := rdpu.mutation.Name(); ok {
		_spec.SetField(rent591homedetailpublish.FieldName, field.TypeString, value)
	}
	if value, ok := rdpu.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpublish.FieldKey, field.TypeString, value)
	}
	if value, ok := rdpu.mutation.PostTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostTime, field.TypeString, value)
	}
	if value, ok := rdpu.mutation.UpdateTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldUpdateTime, field.TypeString, value)
	}
	if rdpu.mutation.Rent591homeDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdpu.mutation.RemovedRent591homeDetailsIDs(); len(nodes) > 0 && !rdpu.mutation.Rent591homeDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdpu.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rdpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591homedetailpublish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rdpu.mutation.done = true
	return n, nil
}

// Rent591HomeDetailPublishUpdateOne is the builder for updating a single Rent591HomeDetailPublish entity.
type Rent591HomeDetailPublishUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *Rent591HomeDetailPublishMutation
}

// SetPostID sets the "post_id" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SetPostID(i int) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.ResetPostID()
	rdpuo.mutation.SetPostID(i)
	return rdpuo
}

// AddPostID adds i to the "post_id" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) AddPostID(i int) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.AddPostID(i)
	return rdpuo
}

// SetName sets the "name" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SetName(s string) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.SetName(s)
	return rdpuo
}

// SetKey sets the "key" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SetKey(s string) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.SetKey(s)
	return rdpuo
}

// SetPostTime sets the "post_time" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SetPostTime(s string) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.SetPostTime(s)
	return rdpuo
}

// SetUpdateTime sets the "update_time" field.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SetUpdateTime(s string) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.SetUpdateTime(s)
	return rdpuo
}

// AddRent591homeDetailIDs adds the "rent591home_details" edge to the Rent591HomeDetail entity by IDs.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) AddRent591homeDetailIDs(ids ...int) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.AddRent591homeDetailIDs(ids...)
	return rdpuo
}

// AddRent591homeDetails adds the "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) AddRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailPublishUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdpuo.AddRent591homeDetailIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPublishMutation object of the builder.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) Mutation() *Rent591HomeDetailPublishMutation {
	return rdpuo.mutation
}

// ClearRent591homeDetails clears all "rent591home_details" edges to the Rent591HomeDetail entity.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) ClearRent591homeDetails() *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.ClearRent591homeDetails()
	return rdpuo
}

// RemoveRent591homeDetailIDs removes the "rent591home_details" edge to Rent591HomeDetail entities by IDs.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) RemoveRent591homeDetailIDs(ids ...int) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.RemoveRent591homeDetailIDs(ids...)
	return rdpuo
}

// RemoveRent591homeDetails removes "rent591home_details" edges to Rent591HomeDetail entities.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) RemoveRent591homeDetails(r ...*Rent591HomeDetail) *Rent591HomeDetailPublishUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdpuo.RemoveRent591homeDetailIDs(ids...)
}

// Where appends a list predicates to the Rent591HomeDetailPublishUpdate builder.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) Where(ps ...predicate.Rent591HomeDetailPublish) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.mutation.Where(ps...)
	return rdpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) Select(field string, fields ...string) *Rent591HomeDetailPublishUpdateOne {
	rdpuo.fields = append([]string{field}, fields...)
	return rdpuo
}

// Save executes the query and returns the updated Rent591HomeDetailPublish entity.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) Save(ctx context.Context) (*Rent591HomeDetailPublish, error) {
	return withHooks(ctx, rdpuo.sqlSave, rdpuo.mutation, rdpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) SaveX(ctx context.Context) *Rent591HomeDetailPublish {
	node, err := rdpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) Exec(ctx context.Context) error {
	_, err := rdpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdpuo *Rent591HomeDetailPublishUpdateOne) ExecX(ctx context.Context) {
	if err := rdpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rdpuo *Rent591HomeDetailPublishUpdateOne) sqlSave(ctx context.Context) (_node *Rent591HomeDetailPublish, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591homedetailpublish.Table, rent591homedetailpublish.Columns, sqlgraph.NewFieldSpec(rent591homedetailpublish.FieldID, field.TypeInt))
	id, ok := rdpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rent591HomeDetailPublish.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rdpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rent591homedetailpublish.FieldID)
		for _, f := range fields {
			if !rent591homedetailpublish.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rent591homedetailpublish.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rdpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdpuo.mutation.PostID(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostID, field.TypeInt, value)
	}
	if value, ok := rdpuo.mutation.AddedPostID(); ok {
		_spec.AddField(rent591homedetailpublish.FieldPostID, field.TypeInt, value)
	}
	if value, ok := rdpuo.mutation.Name(); ok {
		_spec.SetField(rent591homedetailpublish.FieldName, field.TypeString, value)
	}
	if value, ok := rdpuo.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpublish.FieldKey, field.TypeString, value)
	}
	if value, ok := rdpuo.mutation.PostTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldPostTime, field.TypeString, value)
	}
	if value, ok := rdpuo.mutation.UpdateTime(); ok {
		_spec.SetField(rent591homedetailpublish.FieldUpdateTime, field.TypeString, value)
	}
	if rdpuo.mutation.Rent591homeDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdpuo.mutation.RemovedRent591homeDetailsIDs(); len(nodes) > 0 && !rdpuo.mutation.Rent591homeDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdpuo.mutation.Rent591homeDetailsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rent591HomeDetailPublish{config: rdpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rdpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591homedetailpublish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rdpuo.mutation.done = true
	return _node, nil
}
