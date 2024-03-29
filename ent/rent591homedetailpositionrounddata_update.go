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
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionround"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionrounddata"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionrounddatachildren"
)

// Rent591HomeDetailPositionRoundDataUpdate is the builder for updating Rent591HomeDetailPositionRoundData entities.
type Rent591HomeDetailPositionRoundDataUpdate struct {
	config
	hooks    []Hook
	mutation *Rent591HomeDetailPositionRoundDataMutation
}

// Where appends a list predicates to the Rent591HomeDetailPositionRoundDataUpdate builder.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) Where(ps ...predicate.Rent591HomeDetailPositionRoundData) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.Where(ps...)
	return rdprdu
}

// SetName sets the "name" field.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) SetName(s string) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.SetName(s)
	return rdprdu
}

// SetKey sets the "key" field.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) SetKey(s string) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.SetKey(s)
	return rdprdu
}

// AddRent591homeDetailPositionRoundIDs adds the "rent591home_detail_position_rounds" edge to the Rent591HomeDetailPositionRound entity by IDs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) AddRent591homeDetailPositionRoundIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.AddRent591homeDetailPositionRoundIDs(ids...)
	return rdprdu
}

// AddRent591homeDetailPositionRounds adds the "rent591home_detail_position_rounds" edges to the Rent591HomeDetailPositionRound entity.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) AddRent591homeDetailPositionRounds(r ...*Rent591HomeDetailPositionRound) *Rent591HomeDetailPositionRoundDataUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprdu.AddRent591homeDetailPositionRoundIDs(ids...)
}

// AddRent591homeDetailPositionRoundDataChildrenIDs adds the "rent591home_detail_position_round_data_childrens" edge to the Rent591HomeDetailPositionRoundDataChildren entity by IDs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) AddRent591homeDetailPositionRoundDataChildrenIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.AddRent591homeDetailPositionRoundDataChildrenIDs(ids...)
	return rdprdu
}

// AddRent591homeDetailPositionRoundDataChildrens adds the "rent591home_detail_position_round_data_childrens" edges to the Rent591HomeDetailPositionRoundDataChildren entity.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) AddRent591homeDetailPositionRoundDataChildrens(r ...*Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprdu.AddRent591homeDetailPositionRoundDataChildrenIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPositionRoundDataMutation object of the builder.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) Mutation() *Rent591HomeDetailPositionRoundDataMutation {
	return rdprdu.mutation
}

// ClearRent591homeDetailPositionRounds clears all "rent591home_detail_position_rounds" edges to the Rent591HomeDetailPositionRound entity.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) ClearRent591homeDetailPositionRounds() *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.ClearRent591homeDetailPositionRounds()
	return rdprdu
}

// RemoveRent591homeDetailPositionRoundIDs removes the "rent591home_detail_position_rounds" edge to Rent591HomeDetailPositionRound entities by IDs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) RemoveRent591homeDetailPositionRoundIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.RemoveRent591homeDetailPositionRoundIDs(ids...)
	return rdprdu
}

// RemoveRent591homeDetailPositionRounds removes "rent591home_detail_position_rounds" edges to Rent591HomeDetailPositionRound entities.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) RemoveRent591homeDetailPositionRounds(r ...*Rent591HomeDetailPositionRound) *Rent591HomeDetailPositionRoundDataUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprdu.RemoveRent591homeDetailPositionRoundIDs(ids...)
}

// ClearRent591homeDetailPositionRoundDataChildrens clears all "rent591home_detail_position_round_data_childrens" edges to the Rent591HomeDetailPositionRoundDataChildren entity.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) ClearRent591homeDetailPositionRoundDataChildrens() *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.ClearRent591homeDetailPositionRoundDataChildrens()
	return rdprdu
}

// RemoveRent591homeDetailPositionRoundDataChildrenIDs removes the "rent591home_detail_position_round_data_childrens" edge to Rent591HomeDetailPositionRoundDataChildren entities by IDs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdate {
	rdprdu.mutation.RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids...)
	return rdprdu
}

// RemoveRent591homeDetailPositionRoundDataChildrens removes "rent591home_detail_position_round_data_childrens" edges to Rent591HomeDetailPositionRoundDataChildren entities.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) RemoveRent591homeDetailPositionRoundDataChildrens(r ...*Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprdu.RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rdprdu.sqlSave, rdprdu.mutation, rdprdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) SaveX(ctx context.Context) int {
	affected, err := rdprdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) Exec(ctx context.Context) error {
	_, err := rdprdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) ExecX(ctx context.Context) {
	if err := rdprdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rdprdu *Rent591HomeDetailPositionRoundDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591homedetailpositionrounddata.Table, rent591homedetailpositionrounddata.Columns, sqlgraph.NewFieldSpec(rent591homedetailpositionrounddata.FieldID, field.TypeInt))
	if ps := rdprdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdprdu.mutation.Name(); ok {
		_spec.SetField(rent591homedetailpositionrounddata.FieldName, field.TypeString, value)
	}
	if value, ok := rdprdu.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpositionrounddata.FieldKey, field.TypeString, value)
	}
	if rdprdu.mutation.Rent591homeDetailPositionRoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprdu.mutation.RemovedRent591homeDetailPositionRoundsIDs(); len(nodes) > 0 && !rdprdu.mutation.Rent591homeDetailPositionRoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprdu.mutation.Rent591homeDetailPositionRoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdprdu.mutation.Rent591homeDetailPositionRoundDataChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprdu.mutation.RemovedRent591homeDetailPositionRoundDataChildrensIDs(); len(nodes) > 0 && !rdprdu.mutation.Rent591homeDetailPositionRoundDataChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprdu.mutation.Rent591homeDetailPositionRoundDataChildrensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rdprdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591homedetailpositionrounddata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rdprdu.mutation.done = true
	return n, nil
}

// Rent591HomeDetailPositionRoundDataUpdateOne is the builder for updating a single Rent591HomeDetailPositionRoundData entity.
type Rent591HomeDetailPositionRoundDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *Rent591HomeDetailPositionRoundDataMutation
}

// SetName sets the "name" field.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) SetName(s string) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.SetName(s)
	return rdprduo
}

// SetKey sets the "key" field.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) SetKey(s string) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.SetKey(s)
	return rdprduo
}

// AddRent591homeDetailPositionRoundIDs adds the "rent591home_detail_position_rounds" edge to the Rent591HomeDetailPositionRound entity by IDs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) AddRent591homeDetailPositionRoundIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.AddRent591homeDetailPositionRoundIDs(ids...)
	return rdprduo
}

// AddRent591homeDetailPositionRounds adds the "rent591home_detail_position_rounds" edges to the Rent591HomeDetailPositionRound entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) AddRent591homeDetailPositionRounds(r ...*Rent591HomeDetailPositionRound) *Rent591HomeDetailPositionRoundDataUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprduo.AddRent591homeDetailPositionRoundIDs(ids...)
}

// AddRent591homeDetailPositionRoundDataChildrenIDs adds the "rent591home_detail_position_round_data_childrens" edge to the Rent591HomeDetailPositionRoundDataChildren entity by IDs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) AddRent591homeDetailPositionRoundDataChildrenIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.AddRent591homeDetailPositionRoundDataChildrenIDs(ids...)
	return rdprduo
}

// AddRent591homeDetailPositionRoundDataChildrens adds the "rent591home_detail_position_round_data_childrens" edges to the Rent591HomeDetailPositionRoundDataChildren entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) AddRent591homeDetailPositionRoundDataChildrens(r ...*Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprduo.AddRent591homeDetailPositionRoundDataChildrenIDs(ids...)
}

// Mutation returns the Rent591HomeDetailPositionRoundDataMutation object of the builder.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) Mutation() *Rent591HomeDetailPositionRoundDataMutation {
	return rdprduo.mutation
}

// ClearRent591homeDetailPositionRounds clears all "rent591home_detail_position_rounds" edges to the Rent591HomeDetailPositionRound entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) ClearRent591homeDetailPositionRounds() *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.ClearRent591homeDetailPositionRounds()
	return rdprduo
}

// RemoveRent591homeDetailPositionRoundIDs removes the "rent591home_detail_position_rounds" edge to Rent591HomeDetailPositionRound entities by IDs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) RemoveRent591homeDetailPositionRoundIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.RemoveRent591homeDetailPositionRoundIDs(ids...)
	return rdprduo
}

// RemoveRent591homeDetailPositionRounds removes "rent591home_detail_position_rounds" edges to Rent591HomeDetailPositionRound entities.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) RemoveRent591homeDetailPositionRounds(r ...*Rent591HomeDetailPositionRound) *Rent591HomeDetailPositionRoundDataUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprduo.RemoveRent591homeDetailPositionRoundIDs(ids...)
}

// ClearRent591homeDetailPositionRoundDataChildrens clears all "rent591home_detail_position_round_data_childrens" edges to the Rent591HomeDetailPositionRoundDataChildren entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) ClearRent591homeDetailPositionRoundDataChildrens() *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.ClearRent591homeDetailPositionRoundDataChildrens()
	return rdprduo
}

// RemoveRent591homeDetailPositionRoundDataChildrenIDs removes the "rent591home_detail_position_round_data_childrens" edge to Rent591HomeDetailPositionRoundDataChildren entities by IDs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids ...int) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids...)
	return rdprduo
}

// RemoveRent591homeDetailPositionRoundDataChildrens removes "rent591home_detail_position_round_data_childrens" edges to Rent591HomeDetailPositionRoundDataChildren entities.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) RemoveRent591homeDetailPositionRoundDataChildrens(r ...*Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rdprduo.RemoveRent591homeDetailPositionRoundDataChildrenIDs(ids...)
}

// Where appends a list predicates to the Rent591HomeDetailPositionRoundDataUpdate builder.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) Where(ps ...predicate.Rent591HomeDetailPositionRoundData) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.mutation.Where(ps...)
	return rdprduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) Select(field string, fields ...string) *Rent591HomeDetailPositionRoundDataUpdateOne {
	rdprduo.fields = append([]string{field}, fields...)
	return rdprduo
}

// Save executes the query and returns the updated Rent591HomeDetailPositionRoundData entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) Save(ctx context.Context) (*Rent591HomeDetailPositionRoundData, error) {
	return withHooks(ctx, rdprduo.sqlSave, rdprduo.mutation, rdprduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) SaveX(ctx context.Context) *Rent591HomeDetailPositionRoundData {
	node, err := rdprduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) Exec(ctx context.Context) error {
	_, err := rdprduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) ExecX(ctx context.Context) {
	if err := rdprduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rdprduo *Rent591HomeDetailPositionRoundDataUpdateOne) sqlSave(ctx context.Context) (_node *Rent591HomeDetailPositionRoundData, err error) {
	_spec := sqlgraph.NewUpdateSpec(rent591homedetailpositionrounddata.Table, rent591homedetailpositionrounddata.Columns, sqlgraph.NewFieldSpec(rent591homedetailpositionrounddata.FieldID, field.TypeInt))
	id, ok := rdprduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rent591HomeDetailPositionRoundData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rdprduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rent591homedetailpositionrounddata.FieldID)
		for _, f := range fields {
			if !rent591homedetailpositionrounddata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rent591homedetailpositionrounddata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rdprduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rdprduo.mutation.Name(); ok {
		_spec.SetField(rent591homedetailpositionrounddata.FieldName, field.TypeString, value)
	}
	if value, ok := rdprduo.mutation.Key(); ok {
		_spec.SetField(rent591homedetailpositionrounddata.FieldKey, field.TypeString, value)
	}
	if rdprduo.mutation.Rent591homeDetailPositionRoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprduo.mutation.RemovedRent591homeDetailPositionRoundsIDs(); len(nodes) > 0 && !rdprduo.mutation.Rent591homeDetailPositionRoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprduo.mutation.Rent591homeDetailPositionRoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionround.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rdprduo.mutation.Rent591homeDetailPositionRoundDataChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprduo.mutation.RemovedRent591homeDetailPositionRoundDataChildrensIDs(); len(nodes) > 0 && !rdprduo.mutation.Rent591homeDetailPositionRoundDataChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rdprduo.mutation.Rent591homeDetailPositionRoundDataChildrensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensTable,
			Columns: rent591homedetailpositionrounddata.Rent591homeDetailPositionRoundDataChildrensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rent591HomeDetailPositionRoundData{config: rdprduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rdprduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent591homedetailpositionrounddata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rdprduo.mutation.done = true
	return _node, nil
}
