// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wsw/backend/ent/errorresult"
	"wsw/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ErrorResultDelete is the builder for deleting a ErrorResult entity.
type ErrorResultDelete struct {
	config
	hooks    []Hook
	mutation *ErrorResultMutation
}

// Where appends a list predicates to the ErrorResultDelete builder.
func (erd *ErrorResultDelete) Where(ps ...predicate.ErrorResult) *ErrorResultDelete {
	erd.mutation.Where(ps...)
	return erd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (erd *ErrorResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, erd.sqlExec, erd.mutation, erd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (erd *ErrorResultDelete) ExecX(ctx context.Context) int {
	n, err := erd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (erd *ErrorResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(errorresult.Table, sqlgraph.NewFieldSpec(errorresult.FieldID, field.TypeInt))
	if ps := erd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, erd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	erd.mutation.done = true
	return affected, err
}

// ErrorResultDeleteOne is the builder for deleting a single ErrorResult entity.
type ErrorResultDeleteOne struct {
	erd *ErrorResultDelete
}

// Where appends a list predicates to the ErrorResultDelete builder.
func (erdo *ErrorResultDeleteOne) Where(ps ...predicate.ErrorResult) *ErrorResultDeleteOne {
	erdo.erd.mutation.Where(ps...)
	return erdo
}

// Exec executes the deletion query.
func (erdo *ErrorResultDeleteOne) Exec(ctx context.Context) error {
	n, err := erdo.erd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{errorresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (erdo *ErrorResultDeleteOne) ExecX(ctx context.Context) {
	if err := erdo.Exec(ctx); err != nil {
		panic(err)
	}
}
