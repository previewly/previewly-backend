// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/types"
	"wsw/backend/ent/uploadimage"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageProcessCreate is the builder for creating a ImageProcess entity.
type ImageProcessCreate struct {
	config
	mutation *ImageProcessMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (ipc *ImageProcessCreate) SetStatus(te types.StatusEnum) *ImageProcessCreate {
	ipc.mutation.SetStatus(te)
	return ipc
}

// SetProcessHash sets the "process_hash" field.
func (ipc *ImageProcessCreate) SetProcessHash(s string) *ImageProcessCreate {
	ipc.mutation.SetProcessHash(s)
	return ipc
}

// SetProcesses sets the "processes" field.
func (ipc *ImageProcessCreate) SetProcesses(tp []types.ImageProcess) *ImageProcessCreate {
	ipc.mutation.SetProcesses(tp)
	return ipc
}

// SetCreatedAt sets the "created_at" field.
func (ipc *ImageProcessCreate) SetCreatedAt(t time.Time) *ImageProcessCreate {
	ipc.mutation.SetCreatedAt(t)
	return ipc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableCreatedAt(t *time.Time) *ImageProcessCreate {
	if t != nil {
		ipc.SetCreatedAt(*t)
	}
	return ipc
}

// SetUpdatedAt sets the "updated_at" field.
func (ipc *ImageProcessCreate) SetUpdatedAt(t time.Time) *ImageProcessCreate {
	ipc.mutation.SetUpdatedAt(t)
	return ipc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableUpdatedAt(t *time.Time) *ImageProcessCreate {
	if t != nil {
		ipc.SetUpdatedAt(*t)
	}
	return ipc
}

// SetPathPrefix sets the "path_prefix" field.
func (ipc *ImageProcessCreate) SetPathPrefix(s string) *ImageProcessCreate {
	ipc.mutation.SetPathPrefix(s)
	return ipc
}

// SetNillablePathPrefix sets the "path_prefix" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillablePathPrefix(s *string) *ImageProcessCreate {
	if s != nil {
		ipc.SetPathPrefix(*s)
	}
	return ipc
}

// SetError sets the "error" field.
func (ipc *ImageProcessCreate) SetError(s string) *ImageProcessCreate {
	ipc.mutation.SetError(s)
	return ipc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableError(s *string) *ImageProcessCreate {
	if s != nil {
		ipc.SetError(*s)
	}
	return ipc
}

// SetUploadimageID sets the "uploadimage" edge to the UploadImage entity by ID.
func (ipc *ImageProcessCreate) SetUploadimageID(id int) *ImageProcessCreate {
	ipc.mutation.SetUploadimageID(id)
	return ipc
}

// SetNillableUploadimageID sets the "uploadimage" edge to the UploadImage entity by ID if the given value is not nil.
func (ipc *ImageProcessCreate) SetNillableUploadimageID(id *int) *ImageProcessCreate {
	if id != nil {
		ipc = ipc.SetUploadimageID(*id)
	}
	return ipc
}

// SetUploadimage sets the "uploadimage" edge to the UploadImage entity.
func (ipc *ImageProcessCreate) SetUploadimage(u *UploadImage) *ImageProcessCreate {
	return ipc.SetUploadimageID(u.ID)
}

// Mutation returns the ImageProcessMutation object of the builder.
func (ipc *ImageProcessCreate) Mutation() *ImageProcessMutation {
	return ipc.mutation
}

// Save creates the ImageProcess in the database.
func (ipc *ImageProcessCreate) Save(ctx context.Context) (*ImageProcess, error) {
	ipc.defaults()
	return withHooks(ctx, ipc.sqlSave, ipc.mutation, ipc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ipc *ImageProcessCreate) SaveX(ctx context.Context) *ImageProcess {
	v, err := ipc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipc *ImageProcessCreate) Exec(ctx context.Context) error {
	_, err := ipc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipc *ImageProcessCreate) ExecX(ctx context.Context) {
	if err := ipc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ipc *ImageProcessCreate) defaults() {
	if _, ok := ipc.mutation.CreatedAt(); !ok {
		v := imageprocess.DefaultCreatedAt()
		ipc.mutation.SetCreatedAt(v)
	}
	if _, ok := ipc.mutation.UpdatedAt(); !ok {
		v := imageprocess.DefaultUpdatedAt()
		ipc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ipc *ImageProcessCreate) check() error {
	if _, ok := ipc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ImageProcess.status"`)}
	}
	if v, ok := ipc.mutation.Status(); ok {
		if err := imageprocess.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ImageProcess.status": %w`, err)}
		}
	}
	if _, ok := ipc.mutation.ProcessHash(); !ok {
		return &ValidationError{Name: "process_hash", err: errors.New(`ent: missing required field "ImageProcess.process_hash"`)}
	}
	if _, ok := ipc.mutation.Processes(); !ok {
		return &ValidationError{Name: "processes", err: errors.New(`ent: missing required field "ImageProcess.processes"`)}
	}
	if _, ok := ipc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ImageProcess.created_at"`)}
	}
	if _, ok := ipc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ImageProcess.updated_at"`)}
	}
	return nil
}

func (ipc *ImageProcessCreate) sqlSave(ctx context.Context) (*ImageProcess, error) {
	if err := ipc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ipc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ipc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ipc.mutation.id = &_node.ID
	ipc.mutation.done = true
	return _node, nil
}

func (ipc *ImageProcessCreate) createSpec() (*ImageProcess, *sqlgraph.CreateSpec) {
	var (
		_node = &ImageProcess{config: ipc.config}
		_spec = sqlgraph.NewCreateSpec(imageprocess.Table, sqlgraph.NewFieldSpec(imageprocess.FieldID, field.TypeInt))
	)
	if value, ok := ipc.mutation.Status(); ok {
		_spec.SetField(imageprocess.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ipc.mutation.ProcessHash(); ok {
		_spec.SetField(imageprocess.FieldProcessHash, field.TypeString, value)
		_node.ProcessHash = value
	}
	if value, ok := ipc.mutation.Processes(); ok {
		_spec.SetField(imageprocess.FieldProcesses, field.TypeJSON, value)
		_node.Processes = value
	}
	if value, ok := ipc.mutation.CreatedAt(); ok {
		_spec.SetField(imageprocess.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ipc.mutation.UpdatedAt(); ok {
		_spec.SetField(imageprocess.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ipc.mutation.PathPrefix(); ok {
		_spec.SetField(imageprocess.FieldPathPrefix, field.TypeString, value)
		_node.PathPrefix = value
	}
	if value, ok := ipc.mutation.Error(); ok {
		_spec.SetField(imageprocess.FieldError, field.TypeString, value)
		_node.Error = value
	}
	if nodes := ipc.mutation.UploadimageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imageprocess.UploadimageTable,
			Columns: []string{imageprocess.UploadimageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(uploadimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.upload_image_imageprocess = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageProcessCreateBulk is the builder for creating many ImageProcess entities in bulk.
type ImageProcessCreateBulk struct {
	config
	err      error
	builders []*ImageProcessCreate
}

// Save creates the ImageProcess entities in the database.
func (ipcb *ImageProcessCreateBulk) Save(ctx context.Context) ([]*ImageProcess, error) {
	if ipcb.err != nil {
		return nil, ipcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ipcb.builders))
	nodes := make([]*ImageProcess, len(ipcb.builders))
	mutators := make([]Mutator, len(ipcb.builders))
	for i := range ipcb.builders {
		func(i int, root context.Context) {
			builder := ipcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageProcessMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ipcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ipcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ipcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ipcb *ImageProcessCreateBulk) SaveX(ctx context.Context) []*ImageProcess {
	v, err := ipcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ipcb *ImageProcessCreateBulk) Exec(ctx context.Context) error {
	_, err := ipcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipcb *ImageProcessCreateBulk) ExecX(ctx context.Context) {
	if err := ipcb.Exec(ctx); err != nil {
		panic(err)
	}
}
