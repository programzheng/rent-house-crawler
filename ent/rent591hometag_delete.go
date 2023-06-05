// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591hometag"
)

// Rent591HomeTagDelete is the builder for deleting a Rent591HomeTag entity.
type Rent591HomeTagDelete struct {
	config
	hooks    []Hook
	mutation *Rent591HomeTagMutation
}

// Where appends a list predicates to the Rent591HomeTagDelete builder.
func (rtd *Rent591HomeTagDelete) Where(ps ...predicate.Rent591HomeTag) *Rent591HomeTagDelete {
	rtd.mutation.Where(ps...)
	return rtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rtd *Rent591HomeTagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rtd.sqlExec, rtd.mutation, rtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rtd *Rent591HomeTagDelete) ExecX(ctx context.Context) int {
	n, err := rtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rtd *Rent591HomeTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rent591hometag.Table, sqlgraph.NewFieldSpec(rent591hometag.FieldID, field.TypeInt))
	if ps := rtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rtd.mutation.done = true
	return affected, err
}

// Rent591HomeTagDeleteOne is the builder for deleting a single Rent591HomeTag entity.
type Rent591HomeTagDeleteOne struct {
	rtd *Rent591HomeTagDelete
}

// Where appends a list predicates to the Rent591HomeTagDelete builder.
func (rtdo *Rent591HomeTagDeleteOne) Where(ps ...predicate.Rent591HomeTag) *Rent591HomeTagDeleteOne {
	rtdo.rtd.mutation.Where(ps...)
	return rtdo
}

// Exec executes the deletion query.
func (rtdo *Rent591HomeTagDeleteOne) Exec(ctx context.Context) error {
	n, err := rtdo.rtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rent591hometag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rtdo *Rent591HomeTagDeleteOne) ExecX(ctx context.Context) {
	if err := rtdo.Exec(ctx); err != nil {
		panic(err)
	}
}