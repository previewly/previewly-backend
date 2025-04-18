// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"wsw/backend/ent/image"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageProcessQuery is the builder for querying ImageProcess entities.
type ImageProcessQuery struct {
	config
	ctx             *QueryContext
	order           []imageprocess.OrderOption
	inters          []Interceptor
	predicates      []predicate.ImageProcess
	withUploadimage *ImageQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImageProcessQuery builder.
func (ipq *ImageProcessQuery) Where(ps ...predicate.ImageProcess) *ImageProcessQuery {
	ipq.predicates = append(ipq.predicates, ps...)
	return ipq
}

// Limit the number of records to be returned by this query.
func (ipq *ImageProcessQuery) Limit(limit int) *ImageProcessQuery {
	ipq.ctx.Limit = &limit
	return ipq
}

// Offset to start from.
func (ipq *ImageProcessQuery) Offset(offset int) *ImageProcessQuery {
	ipq.ctx.Offset = &offset
	return ipq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ipq *ImageProcessQuery) Unique(unique bool) *ImageProcessQuery {
	ipq.ctx.Unique = &unique
	return ipq
}

// Order specifies how the records should be ordered.
func (ipq *ImageProcessQuery) Order(o ...imageprocess.OrderOption) *ImageProcessQuery {
	ipq.order = append(ipq.order, o...)
	return ipq
}

// QueryUploadimage chains the current query on the "uploadimage" edge.
func (ipq *ImageProcessQuery) QueryUploadimage() *ImageQuery {
	query := (&ImageClient{config: ipq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ipq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(imageprocess.Table, imageprocess.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imageprocess.UploadimageTable, imageprocess.UploadimageColumn),
		)
		fromU = sqlgraph.SetNeighbors(ipq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ImageProcess entity from the query.
// Returns a *NotFoundError when no ImageProcess was found.
func (ipq *ImageProcessQuery) First(ctx context.Context) (*ImageProcess, error) {
	nodes, err := ipq.Limit(1).All(setContextOp(ctx, ipq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{imageprocess.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ipq *ImageProcessQuery) FirstX(ctx context.Context) *ImageProcess {
	node, err := ipq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImageProcess ID from the query.
// Returns a *NotFoundError when no ImageProcess ID was found.
func (ipq *ImageProcessQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(1).IDs(setContextOp(ctx, ipq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{imageprocess.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ipq *ImageProcessQuery) FirstIDX(ctx context.Context) int {
	id, err := ipq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImageProcess entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImageProcess entity is found.
// Returns a *NotFoundError when no ImageProcess entities are found.
func (ipq *ImageProcessQuery) Only(ctx context.Context) (*ImageProcess, error) {
	nodes, err := ipq.Limit(2).All(setContextOp(ctx, ipq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{imageprocess.Label}
	default:
		return nil, &NotSingularError{imageprocess.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ipq *ImageProcessQuery) OnlyX(ctx context.Context) *ImageProcess {
	node, err := ipq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImageProcess ID in the query.
// Returns a *NotSingularError when more than one ImageProcess ID is found.
// Returns a *NotFoundError when no entities are found.
func (ipq *ImageProcessQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(2).IDs(setContextOp(ctx, ipq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{imageprocess.Label}
	default:
		err = &NotSingularError{imageprocess.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ipq *ImageProcessQuery) OnlyIDX(ctx context.Context) int {
	id, err := ipq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImageProcesses.
func (ipq *ImageProcessQuery) All(ctx context.Context) ([]*ImageProcess, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryAll)
	if err := ipq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ImageProcess, *ImageProcessQuery]()
	return withInterceptors[[]*ImageProcess](ctx, ipq, qr, ipq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ipq *ImageProcessQuery) AllX(ctx context.Context) []*ImageProcess {
	nodes, err := ipq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImageProcess IDs.
func (ipq *ImageProcessQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ipq.ctx.Unique == nil && ipq.path != nil {
		ipq.Unique(true)
	}
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryIDs)
	if err = ipq.Select(imageprocess.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ipq *ImageProcessQuery) IDsX(ctx context.Context) []int {
	ids, err := ipq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ipq *ImageProcessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryCount)
	if err := ipq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ipq, querierCount[*ImageProcessQuery](), ipq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ipq *ImageProcessQuery) CountX(ctx context.Context) int {
	count, err := ipq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ipq *ImageProcessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryExist)
	switch _, err := ipq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ipq *ImageProcessQuery) ExistX(ctx context.Context) bool {
	exist, err := ipq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImageProcessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ipq *ImageProcessQuery) Clone() *ImageProcessQuery {
	if ipq == nil {
		return nil
	}
	return &ImageProcessQuery{
		config:          ipq.config,
		ctx:             ipq.ctx.Clone(),
		order:           append([]imageprocess.OrderOption{}, ipq.order...),
		inters:          append([]Interceptor{}, ipq.inters...),
		predicates:      append([]predicate.ImageProcess{}, ipq.predicates...),
		withUploadimage: ipq.withUploadimage.Clone(),
		// clone intermediate query.
		sql:  ipq.sql.Clone(),
		path: ipq.path,
	}
}

// WithUploadimage tells the query-builder to eager-load the nodes that are connected to
// the "uploadimage" edge. The optional arguments are used to configure the query builder of the edge.
func (ipq *ImageProcessQuery) WithUploadimage(opts ...func(*ImageQuery)) *ImageProcessQuery {
	query := (&ImageClient{config: ipq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ipq.withUploadimage = query
	return ipq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status types.StatusEnum `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ImageProcess.Query().
//		GroupBy(imageprocess.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ipq *ImageProcessQuery) GroupBy(field string, fields ...string) *ImageProcessGroupBy {
	ipq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImageProcessGroupBy{build: ipq}
	grbuild.flds = &ipq.ctx.Fields
	grbuild.label = imageprocess.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status types.StatusEnum `json:"status,omitempty"`
//	}
//
//	client.ImageProcess.Query().
//		Select(imageprocess.FieldStatus).
//		Scan(ctx, &v)
func (ipq *ImageProcessQuery) Select(fields ...string) *ImageProcessSelect {
	ipq.ctx.Fields = append(ipq.ctx.Fields, fields...)
	sbuild := &ImageProcessSelect{ImageProcessQuery: ipq}
	sbuild.label = imageprocess.Label
	sbuild.flds, sbuild.scan = &ipq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImageProcessSelect configured with the given aggregations.
func (ipq *ImageProcessQuery) Aggregate(fns ...AggregateFunc) *ImageProcessSelect {
	return ipq.Select().Aggregate(fns...)
}

func (ipq *ImageProcessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ipq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ipq); err != nil {
				return err
			}
		}
	}
	for _, f := range ipq.ctx.Fields {
		if !imageprocess.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ipq.path != nil {
		prev, err := ipq.path(ctx)
		if err != nil {
			return err
		}
		ipq.sql = prev
	}
	return nil
}

func (ipq *ImageProcessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ImageProcess, error) {
	var (
		nodes       = []*ImageProcess{}
		withFKs     = ipq.withFKs
		_spec       = ipq.querySpec()
		loadedTypes = [1]bool{
			ipq.withUploadimage != nil,
		}
	)
	if ipq.withUploadimage != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, imageprocess.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ImageProcess).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ImageProcess{config: ipq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ipq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ipq.withUploadimage; query != nil {
		if err := ipq.loadUploadimage(ctx, query, nodes, nil,
			func(n *ImageProcess, e *Image) { n.Edges.Uploadimage = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ipq *ImageProcessQuery) loadUploadimage(ctx context.Context, query *ImageQuery, nodes []*ImageProcess, init func(*ImageProcess), assign func(*ImageProcess, *Image)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ImageProcess)
	for i := range nodes {
		if nodes[i].image_imageprocess == nil {
			continue
		}
		fk := *nodes[i].image_imageprocess
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(image.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "image_imageprocess" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ipq *ImageProcessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ipq.querySpec()
	_spec.Node.Columns = ipq.ctx.Fields
	if len(ipq.ctx.Fields) > 0 {
		_spec.Unique = ipq.ctx.Unique != nil && *ipq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ipq.driver, _spec)
}

func (ipq *ImageProcessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(imageprocess.Table, imageprocess.Columns, sqlgraph.NewFieldSpec(imageprocess.FieldID, field.TypeInt))
	_spec.From = ipq.sql
	if unique := ipq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ipq.path != nil {
		_spec.Unique = true
	}
	if fields := ipq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, imageprocess.FieldID)
		for i := range fields {
			if fields[i] != imageprocess.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ipq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ipq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ipq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ipq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ipq *ImageProcessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ipq.driver.Dialect())
	t1 := builder.Table(imageprocess.Table)
	columns := ipq.ctx.Fields
	if len(columns) == 0 {
		columns = imageprocess.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ipq.sql != nil {
		selector = ipq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ipq.ctx.Unique != nil && *ipq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ipq.predicates {
		p(selector)
	}
	for _, p := range ipq.order {
		p(selector)
	}
	if offset := ipq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ipq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImageProcessGroupBy is the group-by builder for ImageProcess entities.
type ImageProcessGroupBy struct {
	selector
	build *ImageProcessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ipgb *ImageProcessGroupBy) Aggregate(fns ...AggregateFunc) *ImageProcessGroupBy {
	ipgb.fns = append(ipgb.fns, fns...)
	return ipgb
}

// Scan applies the selector query and scans the result into the given value.
func (ipgb *ImageProcessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ipgb.build.ctx, ent.OpQueryGroupBy)
	if err := ipgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageProcessQuery, *ImageProcessGroupBy](ctx, ipgb.build, ipgb, ipgb.build.inters, v)
}

func (ipgb *ImageProcessGroupBy) sqlScan(ctx context.Context, root *ImageProcessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ipgb.fns))
	for _, fn := range ipgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ipgb.flds)+len(ipgb.fns))
		for _, f := range *ipgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ipgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ipgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImageProcessSelect is the builder for selecting fields of ImageProcess entities.
type ImageProcessSelect struct {
	*ImageProcessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ips *ImageProcessSelect) Aggregate(fns ...AggregateFunc) *ImageProcessSelect {
	ips.fns = append(ips.fns, fns...)
	return ips
}

// Scan applies the selector query and scans the result into the given value.
func (ips *ImageProcessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ips.ctx, ent.OpQuerySelect)
	if err := ips.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImageProcessQuery, *ImageProcessSelect](ctx, ips.ImageProcessQuery, ips, ips.inters, v)
}

func (ips *ImageProcessSelect) sqlScan(ctx context.Context, root *ImageProcessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ips.fns))
	for _, fn := range ips.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ips.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ips.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
