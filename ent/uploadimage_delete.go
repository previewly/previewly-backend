// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"wsw/backend/ent/predicate"
	"wsw/backend/ent/uploadimage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UploadImageDelete is the builder for deleting a UploadImage entity.
type UploadImageDelete struct {
	config
	hooks    []Hook
	mutation *UploadImageMutation
}

// Where appends a list predicates to the UploadImageDelete builder.
func (uid *UploadImageDelete) Where(ps ...predicate.UploadImage) *UploadImageDelete {
	uid.mutation.Where(ps...)
	return uid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uid *UploadImageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uid.sqlExec, uid.mutation, uid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uid *UploadImageDelete) ExecX(ctx context.Context) int {
	n, err := uid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uid *UploadImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(uploadimage.Table, sqlgraph.NewFieldSpec(uploadimage.FieldID, field.TypeInt))
	if ps := uid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uid.mutation.done = true
	return affected, err
}

// UploadImageDeleteOne is the builder for deleting a single UploadImage entity.
type UploadImageDeleteOne struct {
	uid *UploadImageDelete
}

// Where appends a list predicates to the UploadImageDelete builder.
func (uido *UploadImageDeleteOne) Where(ps ...predicate.UploadImage) *UploadImageDeleteOne {
	uido.uid.mutation.Where(ps...)
	return uido
}

// Exec executes the deletion query.
func (uido *UploadImageDeleteOne) Exec(ctx context.Context) error {
	n, err := uido.uid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{uploadimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uido *UploadImageDeleteOne) ExecX(ctx context.Context) {
	if err := uido.Exec(ctx); err != nil {
		panic(err)
	}
}
