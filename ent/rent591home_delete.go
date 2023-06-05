// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591home"
)

// Rent591HomeDelete is the builder for deleting a Rent591Home entity.
type Rent591HomeDelete struct {
	config
	hooks    []Hook
	mutation *Rent591HomeMutation
}

// Where appends a list predicates to the Rent591HomeDelete builder.
func (rd *Rent591HomeDelete) Where(ps ...predicate.Rent591Home) *Rent591HomeDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *Rent591HomeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *Rent591HomeDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *Rent591HomeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rent591home.Table, sqlgraph.NewFieldSpec(rent591home.FieldID, field.TypeInt))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// Rent591HomeDeleteOne is the builder for deleting a single Rent591Home entity.
type Rent591HomeDeleteOne struct {
	rd *Rent591HomeDelete
}

// Where appends a list predicates to the Rent591HomeDelete builder.
func (rdo *Rent591HomeDeleteOne) Where(ps ...predicate.Rent591Home) *Rent591HomeDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *Rent591HomeDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rent591home.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *Rent591HomeDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}