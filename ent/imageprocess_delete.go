// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageProcessDelete is the builder for deleting a ImageProcess entity.
type ImageProcessDelete struct {
	config
	hooks    []Hook
	mutation *ImageProcessMutation
}

// Where appends a list predicates to the ImageProcessDelete builder.
func (ipd *ImageProcessDelete) Where(ps ...predicate.ImageProcess) *ImageProcessDelete {
	ipd.mutation.Where(ps...)
	return ipd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ipd *ImageProcessDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ipd.sqlExec, ipd.mutation, ipd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ipd *ImageProcessDelete) ExecX(ctx context.Context) int {
	n, err := ipd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ipd *ImageProcessDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(imageprocess.Table, sqlgraph.NewFieldSpec(imageprocess.FieldID, field.TypeInt))
	if ps := ipd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ipd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ipd.mutation.done = true
	return affected, err
}

// ImageProcessDeleteOne is the builder for deleting a single ImageProcess entity.
type ImageProcessDeleteOne struct {
	ipd *ImageProcessDelete
}

// Where appends a list predicates to the ImageProcessDelete builder.
func (ipdo *ImageProcessDeleteOne) Where(ps ...predicate.ImageProcess) *ImageProcessDeleteOne {
	ipdo.ipd.mutation.Where(ps...)
	return ipdo
}

// Exec executes the deletion query.
func (ipdo *ImageProcessDeleteOne) Exec(ctx context.Context) error {
	n, err := ipdo.ipd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{imageprocess.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ipdo *ImageProcessDeleteOne) ExecX(ctx context.Context) {
	if err := ipdo.Exec(ctx); err != nil {
		panic(err)
	}
}