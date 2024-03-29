// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailpositionrounddatachildren"
)

// Rent591HomeDetailPositionRoundDataChildrenDelete is the builder for deleting a Rent591HomeDetailPositionRoundDataChildren entity.
type Rent591HomeDetailPositionRoundDataChildrenDelete struct {
	config
	hooks    []Hook
	mutation *Rent591HomeDetailPositionRoundDataChildrenMutation
}

// Where appends a list predicates to the Rent591HomeDetailPositionRoundDataChildrenDelete builder.
func (rdprdcd *Rent591HomeDetailPositionRoundDataChildrenDelete) Where(ps ...predicate.Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataChildrenDelete {
	rdprdcd.mutation.Where(ps...)
	return rdprdcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rdprdcd *Rent591HomeDetailPositionRoundDataChildrenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rdprdcd.sqlExec, rdprdcd.mutation, rdprdcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprdcd *Rent591HomeDetailPositionRoundDataChildrenDelete) ExecX(ctx context.Context) int {
	n, err := rdprdcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rdprdcd *Rent591HomeDetailPositionRoundDataChildrenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rent591homedetailpositionrounddatachildren.Table, sqlgraph.NewFieldSpec(rent591homedetailpositionrounddatachildren.FieldID, field.TypeInt))
	if ps := rdprdcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rdprdcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rdprdcd.mutation.done = true
	return affected, err
}

// Rent591HomeDetailPositionRoundDataChildrenDeleteOne is the builder for deleting a single Rent591HomeDetailPositionRoundDataChildren entity.
type Rent591HomeDetailPositionRoundDataChildrenDeleteOne struct {
	rdprdcd *Rent591HomeDetailPositionRoundDataChildrenDelete
}

// Where appends a list predicates to the Rent591HomeDetailPositionRoundDataChildrenDelete builder.
func (rdprdcdo *Rent591HomeDetailPositionRoundDataChildrenDeleteOne) Where(ps ...predicate.Rent591HomeDetailPositionRoundDataChildren) *Rent591HomeDetailPositionRoundDataChildrenDeleteOne {
	rdprdcdo.rdprdcd.mutation.Where(ps...)
	return rdprdcdo
}

// Exec executes the deletion query.
func (rdprdcdo *Rent591HomeDetailPositionRoundDataChildrenDeleteOne) Exec(ctx context.Context) error {
	n, err := rdprdcdo.rdprdcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rent591homedetailpositionrounddatachildren.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdprdcdo *Rent591HomeDetailPositionRoundDataChildrenDeleteOne) ExecX(ctx context.Context) {
	if err := rdprdcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
