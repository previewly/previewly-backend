// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wsw/backend/ent/predicate"
	"wsw/backend/ent/url"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// URLDelete is the builder for deleting a Url entity.
type URLDelete struct {
	config
	hooks    []Hook
	mutation *URLMutation
}

// Where appends a list predicates to the URLDelete builder.
func (ud *URLDelete) Where(ps ...predicate.Url) *URLDelete {
	ud.mutation.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *URLDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ud.sqlExec, ud.mutation, ud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *URLDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *URLDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(url.Table, sqlgraph.NewFieldSpec(url.FieldID, field.TypeInt))
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ud.mutation.done = true
	return affected, err
}

// URLDeleteOne is the builder for deleting a single Url entity.
type URLDeleteOne struct {
	ud *URLDelete
}

// Where appends a list predicates to the URLDelete builder.
func (udo *URLDeleteOne) Where(ps ...predicate.Url) *URLDeleteOne {
	udo.ud.mutation.Where(ps...)
	return udo
}

// Exec executes the deletion query.
func (udo *URLDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{url.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *URLDeleteOne) ExecX(ctx context.Context) {
	if err := udo.Exec(ctx); err != nil {
		panic(err)
	}
}
