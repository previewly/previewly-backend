// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wsw/backend/ent/predicate"
	"wsw/backend/ent/url"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// URLUpdate is the builder for updating Url entities.
type URLUpdate struct {
	config
	hooks    []Hook
	mutation *URLMutation
}

// Where appends a list predicates to the URLUpdate builder.
func (uu *URLUpdate) Where(ps ...predicate.Url) *URLUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetURL sets the "url" field.
func (uu *URLUpdate) SetURL(s string) *URLUpdate {
	uu.mutation.SetURL(s)
	return uu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uu *URLUpdate) SetNillableURL(s *string) *URLUpdate {
	if s != nil {
		uu.SetURL(*s)
	}
	return uu
}

// SetStatus sets the "status" field.
func (uu *URLUpdate) SetStatus(s string) *URLUpdate {
	uu.mutation.SetStatus(s)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *URLUpdate) SetNillableStatus(s *string) *URLUpdate {
	if s != nil {
		uu.SetStatus(*s)
	}
	return uu
}

// Mutation returns the URLMutation object of the builder.
func (uu *URLUpdate) Mutation() *URLMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *URLUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *URLUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *URLUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *URLUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *URLUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(url.Table, url.Columns, sqlgraph.NewFieldSpec(url.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeString, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(url.FieldStatus, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// URLUpdateOne is the builder for updating a single Url entity.
type URLUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *URLMutation
}

// SetURL sets the "url" field.
func (uuo *URLUpdateOne) SetURL(s string) *URLUpdateOne {
	uuo.mutation.SetURL(s)
	return uuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uuo *URLUpdateOne) SetNillableURL(s *string) *URLUpdateOne {
	if s != nil {
		uuo.SetURL(*s)
	}
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *URLUpdateOne) SetStatus(s string) *URLUpdateOne {
	uuo.mutation.SetStatus(s)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *URLUpdateOne) SetNillableStatus(s *string) *URLUpdateOne {
	if s != nil {
		uuo.SetStatus(*s)
	}
	return uuo
}

// Mutation returns the URLMutation object of the builder.
func (uuo *URLUpdateOne) Mutation() *URLMutation {
	return uuo.mutation
}

// Where appends a list predicates to the URLUpdate builder.
func (uuo *URLUpdateOne) Where(ps ...predicate.Url) *URLUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *URLUpdateOne) Select(field string, fields ...string) *URLUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Url entity.
func (uuo *URLUpdateOne) Save(ctx context.Context) (*Url, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *URLUpdateOne) SaveX(ctx context.Context) *Url {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *URLUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *URLUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *URLUpdateOne) sqlSave(ctx context.Context) (_node *Url, err error) {
	_spec := sqlgraph.NewUpdateSpec(url.Table, url.Columns, sqlgraph.NewFieldSpec(url.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Url.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, url.FieldID)
		for _, f := range fields {
			if !url.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != url.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(url.FieldStatus, field.TypeString, value)
	}
	_node = &Url{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}