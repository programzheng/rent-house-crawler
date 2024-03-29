// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591home"
	"github.com/programzheng/rent-house-crawler/ent/rent591hometag"
)

// Rent591HomeTagUpdate is the builder for updating Rent591HomeTag entities.
type Rent591HomeTagUpdate struct {
	config
	hooks    []Hook
	mutation *Rent591HomeTagMutation
}

// Where appends a list predicates to the Rent591HomeTagUpdate builder.
func (rtu *Rent591HomeTagUpdate) Where(ps ...predicate.Rent591HomeTag) *Rent591HomeTagUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetName sets the "name" field.
func (rtu *Rent591HomeTagUpdate) SetName(s string) *Rent591HomeTagUpdate {
	rtu.mutation.SetName(s)
	return rtu
}

// SetCreatedAt sets the "created_at" field.
func (rtu *Rent591HomeTagUpdate) SetCreatedAt(t time.Time) *Rent591HomeTagUpdate {
	rtu.mutation.SetCreatedAt(t)
	return rtu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtu *Rent591HomeTagUpdate) SetNillableCreatedAt(t *time.Time) *Rent591HomeTagUpdate {
	if t != nil {
		rtu.SetCreatedAt(*t)
	}
	return rtu
}

// AddRent591homeIDs adds the "rent591homes" edge to the Rent591Home entity by IDs.
func (rtu *Rent591HomeTagUpdate) AddRent591homeIDs(ids ...int) *Rent591HomeTagUpdate {
	rtu.mutation.AddRent591homeIDs(ids...)
	return rtu
}

// AddRent591homes adds the "rent591homes" edges to the Rent591Home entity.
func (rtu *Rent591HomeTagUpdate) AddRent591homes(r ...*Rent591Home) *Rent591HomeTagUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.AddRent591homeIDs(ids...)
}

// Mutation returns the Rent591HomeTagMutation object of the builder.
func (rtu *Rent591HomeTagUpdate) Mutation() *Rent591HomeTagMutation {
	return rtu.mutation
}

// ClearRent591homes clears all "rent591homes" edges to the Rent591Home entity.
func (rtu *Rent591HomeTagUpdate) ClearRent591homes() *Rent591HomeTagUpdate {
	rtu.mutation.ClearRent591homes()
	return rtu
}

// RemoveRent591homeIDs removes the "rent591homes" edge to Rent591Home entities by IDs.
func (rtu *Rent591HomeTagUpdate) RemoveRent591homeIDs(ids ...int) *Rent591HomeTagUpdate {
	rtu.mutation.RemoveRent591homeIDs(ids...)
	return rtu
}

// RemoveRent591homes removes "rent591homes" edges to Rent591Home entities.
func (rtu *Rent591HomeTagUpdate) RemoveRent591homes(r ...*Rent591Home) *Rent591HomeTagUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.RemoveRent591homeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *Rent591HomeTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rtu.sqlSave, rtu.mutation, rtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *Rent591HomeTagUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *Rent591HomeTagUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *Rent591HomeTagUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtu *Rent591HomeTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591hometag.Table, rent591hometag.Columns, sqlgraph.NewFieldSpec(rent591hometag.FieldID, field.TypeInt))
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.Name(); ok {
		_spec.SetField(rent591hometag.FieldName, field.TypeString, value)
	}
	if value, ok := rtu.mutation.CreatedAt(); ok {
		_spec.SetField(rent591hometag.FieldCreatedAt, field.TypeTime, value)
	}
	if rtu.mutation.Rent591homesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RemovedRent591homesIDs(); len(nodes) > 0 && !rtu.mutation.Rent591homesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.Rent591homesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591hometag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rtu.mutation.done = true
	return n, nil
}

// Rent591HomeTagUpdateOne is the builder for updating a single Rent591HomeTag entity.
type Rent591HomeTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *Rent591HomeTagMutation
}

// SetName sets the "name" field.
func (rtuo *Rent591HomeTagUpdateOne) SetName(s string) *Rent591HomeTagUpdateOne {
	rtuo.mutation.SetName(s)
	return rtuo
}

// SetCreatedAt sets the "created_at" field.
func (rtuo *Rent591HomeTagUpdateOne) SetCreatedAt(t time.Time) *Rent591HomeTagUpdateOne {
	rtuo.mutation.SetCreatedAt(t)
	return rtuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rtuo *Rent591HomeTagUpdateOne) SetNillableCreatedAt(t *time.Time) *Rent591HomeTagUpdateOne {
	if t != nil {
		rtuo.SetCreatedAt(*t)
	}
	return rtuo
}

// AddRent591homeIDs adds the "rent591homes" edge to the Rent591Home entity by IDs.
func (rtuo *Rent591HomeTagUpdateOne) AddRent591homeIDs(ids ...int) *Rent591HomeTagUpdateOne {
	rtuo.mutation.AddRent591homeIDs(ids...)
	return rtuo
}

// AddRent591homes adds the "rent591homes" edges to the Rent591Home entity.
func (rtuo *Rent591HomeTagUpdateOne) AddRent591homes(r ...*Rent591Home) *Rent591HomeTagUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.AddRent591homeIDs(ids...)
}

// Mutation returns the Rent591HomeTagMutation object of the builder.
func (rtuo *Rent591HomeTagUpdateOne) Mutation() *Rent591HomeTagMutation {
	return rtuo.mutation
}

// ClearRent591homes clears all "rent591homes" edges to the Rent591Home entity.
func (rtuo *Rent591HomeTagUpdateOne) ClearRent591homes() *Rent591HomeTagUpdateOne {
	rtuo.mutation.ClearRent591homes()
	return rtuo
}

// RemoveRent591homeIDs removes the "rent591homes" edge to Rent591Home entities by IDs.
func (rtuo *Rent591HomeTagUpdateOne) RemoveRent591homeIDs(ids ...int) *Rent591HomeTagUpdateOne {
	rtuo.mutation.RemoveRent591homeIDs(ids...)
	return rtuo
}

// RemoveRent591homes removes "rent591homes" edges to Rent591Home entities.
func (rtuo *Rent591HomeTagUpdateOne) RemoveRent591homes(r ...*Rent591Home) *Rent591HomeTagUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.RemoveRent591homeIDs(ids...)
}

// Where appends a list predicates to the Rent591HomeTagUpdate builder.
func (rtuo *Rent591HomeTagUpdateOne) Where(ps ...predicate.Rent591HomeTag) *Rent591HomeTagUpdateOne {
	rtuo.mutation.Where(ps...)
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *Rent591HomeTagUpdateOne) Select(field string, fields ...string) *Rent591HomeTagUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated Rent591HomeTag entity.
func (rtuo *Rent591HomeTagUpdateOne) Save(ctx context.Context) (*Rent591HomeTag, error) {
	return withHooks(ctx, rtuo.sqlSave, rtuo.mutation, rtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *Rent591HomeTagUpdateOne) SaveX(ctx context.Context) *Rent591HomeTag {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *Rent591HomeTagUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *Rent591HomeTagUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtuo *Rent591HomeTagUpdateOne) sqlSave(ctx context.Context) (_node *Rent591HomeTag, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591hometag.Table, rent591hometag.Columns, sqlgraph.NewFieldSpec(rent591hometag.FieldID, field.TypeInt))
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rent591HomeTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rent591hometag.FieldID)
		for _, f := range fields {
			if !rent591hometag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rent591hometag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.Name(); ok {
		_spec.SetField(rent591hometag.FieldName, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.CreatedAt(); ok {
		_spec.SetField(rent591hometag.FieldCreatedAt, field.TypeTime, value)
	}
	if rtuo.mutation.Rent591homesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RemovedRent591homesIDs(); len(nodes) > 0 && !rtuo.mutation.Rent591homesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.Rent591homesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591hometag.Rent591homesTable,
			Columns: rent591hometag.Rent591homesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rent591HomeTag{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591hometag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rtuo.mutation.done = true
	return _node, nil
}
