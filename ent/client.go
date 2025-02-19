// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"wsw/backend/ent/migrate"

	"wsw/backend/ent/errorresult"
	"wsw/backend/ent/image"
	"wsw/backend/ent/imageprocess"
	"wsw/backend/ent/stat"
	"wsw/backend/ent/token"
	"wsw/backend/ent/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ErrorResult is the client for interacting with the ErrorResult builders.
	ErrorResult *ErrorResultClient
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// ImageProcess is the client for interacting with the ImageProcess builders.
	ImageProcess *ImageProcessClient
	// Stat is the client for interacting with the Stat builders.
	Stat *StatClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// Url is the client for interacting with the Url builders.
	Url *URLClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ErrorResult = NewErrorResultClient(c.config)
	c.Image = NewImageClient(c.config)
	c.ImageProcess = NewImageProcessClient(c.config)
	c.Stat = NewStatClient(c.config)
	c.Token = NewTokenClient(c.config)
	c.Url = NewURLClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		ErrorResult:  NewErrorResultClient(cfg),
		Image:        NewImageClient(cfg),
		ImageProcess: NewImageProcessClient(cfg),
		Stat:         NewStatClient(cfg),
		Token:        NewTokenClient(cfg),
		Url:          NewURLClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		ErrorResult:  NewErrorResultClient(cfg),
		Image:        NewImageClient(cfg),
		ImageProcess: NewImageProcessClient(cfg),
		Stat:         NewStatClient(cfg),
		Token:        NewTokenClient(cfg),
		Url:          NewURLClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ErrorResult.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	for _, n := range []interface{ Use(...Hook) }{
		c.ErrorResult, c.Image, c.ImageProcess, c.Stat, c.Token, c.Url,
	} {
		n.Use(hooks...)
	}
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	for _, n := range []interface{ Intercept(...Interceptor) }{
		c.ErrorResult, c.Image, c.ImageProcess, c.Stat, c.Token, c.Url,
	} {
		n.Intercept(interceptors...)
	}
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ErrorResultMutation:
		return c.ErrorResult.mutate(ctx, m)
	case *ImageMutation:
		return c.Image.mutate(ctx, m)
	case *ImageProcessMutation:
		return c.ImageProcess.mutate(ctx, m)
	case *StatMutation:
		return c.Stat.mutate(ctx, m)
	case *TokenMutation:
		return c.Token.mutate(ctx, m)
	case *URLMutation:
		return c.Url.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ErrorResultClient is a client for the ErrorResult schema.
type ErrorResultClient struct {
	config
}

// NewErrorResultClient returns a client for the ErrorResult from the given config.
func NewErrorResultClient(c config) *ErrorResultClient {
	return &ErrorResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `errorresult.Hooks(f(g(h())))`.
func (c *ErrorResultClient) Use(hooks ...Hook) {
	c.hooks.ErrorResult = append(c.hooks.ErrorResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `errorresult.Intercept(f(g(h())))`.
func (c *ErrorResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.ErrorResult = append(c.inters.ErrorResult, interceptors...)
}

// Create returns a builder for creating a ErrorResult entity.
func (c *ErrorResultClient) Create() *ErrorResultCreate {
	mutation := newErrorResultMutation(c.config, OpCreate)
	return &ErrorResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ErrorResult entities.
func (c *ErrorResultClient) CreateBulk(builders ...*ErrorResultCreate) *ErrorResultCreateBulk {
	return &ErrorResultCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ErrorResultClient) MapCreateBulk(slice any, setFunc func(*ErrorResultCreate, int)) *ErrorResultCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ErrorResultCreateBulk{err: fmt.Errorf("calling to ErrorResultClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ErrorResultCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ErrorResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ErrorResult.
func (c *ErrorResultClient) Update() *ErrorResultUpdate {
	mutation := newErrorResultMutation(c.config, OpUpdate)
	return &ErrorResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ErrorResultClient) UpdateOne(er *ErrorResult) *ErrorResultUpdateOne {
	mutation := newErrorResultMutation(c.config, OpUpdateOne, withErrorResult(er))
	return &ErrorResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ErrorResultClient) UpdateOneID(id int) *ErrorResultUpdateOne {
	mutation := newErrorResultMutation(c.config, OpUpdateOne, withErrorResultID(id))
	return &ErrorResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ErrorResult.
func (c *ErrorResultClient) Delete() *ErrorResultDelete {
	mutation := newErrorResultMutation(c.config, OpDelete)
	return &ErrorResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ErrorResultClient) DeleteOne(er *ErrorResult) *ErrorResultDeleteOne {
	return c.DeleteOneID(er.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ErrorResultClient) DeleteOneID(id int) *ErrorResultDeleteOne {
	builder := c.Delete().Where(errorresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ErrorResultDeleteOne{builder}
}

// Query returns a query builder for ErrorResult.
func (c *ErrorResultClient) Query() *ErrorResultQuery {
	return &ErrorResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeErrorResult},
		inters: c.Interceptors(),
	}
}

// Get returns a ErrorResult entity by its id.
func (c *ErrorResultClient) Get(ctx context.Context, id int) (*ErrorResult, error) {
	return c.Query().Where(errorresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ErrorResultClient) GetX(ctx context.Context, id int) *ErrorResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ErrorResultClient) Hooks() []Hook {
	return c.hooks.ErrorResult
}

// Interceptors returns the client interceptors.
func (c *ErrorResultClient) Interceptors() []Interceptor {
	return c.inters.ErrorResult
}

func (c *ErrorResultClient) mutate(ctx context.Context, m *ErrorResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ErrorResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ErrorResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ErrorResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ErrorResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ErrorResult mutation op: %q", m.Op())
	}
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `image.Intercept(f(g(h())))`.
func (c *ImageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Image = append(c.inters.Image, interceptors...)
}

// Create returns a builder for creating a Image entity.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ImageClient) MapCreateBulk(slice any, setFunc func(*ImageCreate, int)) *ImageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ImageCreateBulk{err: fmt.Errorf("calling to ImageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ImageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id int) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageClient) DeleteOneID(id int) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeImage},
		inters: c.Interceptors(),
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id int) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id int) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImageprocess queries the imageprocess edge of a Image.
func (c *ImageClient) QueryImageprocess(i *Image) *ImageProcessQuery {
	query := (&ImageProcessClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(imageprocess.Table, imageprocess.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, image.ImageprocessTable, image.ImageprocessColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// Interceptors returns the client interceptors.
func (c *ImageClient) Interceptors() []Interceptor {
	return c.inters.Image
}

func (c *ImageClient) mutate(ctx context.Context, m *ImageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ImageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ImageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Image mutation op: %q", m.Op())
	}
}

// ImageProcessClient is a client for the ImageProcess schema.
type ImageProcessClient struct {
	config
}

// NewImageProcessClient returns a client for the ImageProcess from the given config.
func NewImageProcessClient(c config) *ImageProcessClient {
	return &ImageProcessClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `imageprocess.Hooks(f(g(h())))`.
func (c *ImageProcessClient) Use(hooks ...Hook) {
	c.hooks.ImageProcess = append(c.hooks.ImageProcess, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `imageprocess.Intercept(f(g(h())))`.
func (c *ImageProcessClient) Intercept(interceptors ...Interceptor) {
	c.inters.ImageProcess = append(c.inters.ImageProcess, interceptors...)
}

// Create returns a builder for creating a ImageProcess entity.
func (c *ImageProcessClient) Create() *ImageProcessCreate {
	mutation := newImageProcessMutation(c.config, OpCreate)
	return &ImageProcessCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ImageProcess entities.
func (c *ImageProcessClient) CreateBulk(builders ...*ImageProcessCreate) *ImageProcessCreateBulk {
	return &ImageProcessCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ImageProcessClient) MapCreateBulk(slice any, setFunc func(*ImageProcessCreate, int)) *ImageProcessCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ImageProcessCreateBulk{err: fmt.Errorf("calling to ImageProcessClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ImageProcessCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ImageProcessCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ImageProcess.
func (c *ImageProcessClient) Update() *ImageProcessUpdate {
	mutation := newImageProcessMutation(c.config, OpUpdate)
	return &ImageProcessUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageProcessClient) UpdateOne(ip *ImageProcess) *ImageProcessUpdateOne {
	mutation := newImageProcessMutation(c.config, OpUpdateOne, withImageProcess(ip))
	return &ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageProcessClient) UpdateOneID(id int) *ImageProcessUpdateOne {
	mutation := newImageProcessMutation(c.config, OpUpdateOne, withImageProcessID(id))
	return &ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ImageProcess.
func (c *ImageProcessClient) Delete() *ImageProcessDelete {
	mutation := newImageProcessMutation(c.config, OpDelete)
	return &ImageProcessDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageProcessClient) DeleteOne(ip *ImageProcess) *ImageProcessDeleteOne {
	return c.DeleteOneID(ip.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageProcessClient) DeleteOneID(id int) *ImageProcessDeleteOne {
	builder := c.Delete().Where(imageprocess.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageProcessDeleteOne{builder}
}

// Query returns a query builder for ImageProcess.
func (c *ImageProcessClient) Query() *ImageProcessQuery {
	return &ImageProcessQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeImageProcess},
		inters: c.Interceptors(),
	}
}

// Get returns a ImageProcess entity by its id.
func (c *ImageProcessClient) Get(ctx context.Context, id int) (*ImageProcess, error) {
	return c.Query().Where(imageprocess.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageProcessClient) GetX(ctx context.Context, id int) *ImageProcess {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUploadimage queries the uploadimage edge of a ImageProcess.
func (c *ImageProcessClient) QueryUploadimage(ip *ImageProcess) *ImageQuery {
	query := (&ImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ip.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(imageprocess.Table, imageprocess.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imageprocess.UploadimageTable, imageprocess.UploadimageColumn),
		)
		fromV = sqlgraph.Neighbors(ip.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageProcessClient) Hooks() []Hook {
	return c.hooks.ImageProcess
}

// Interceptors returns the client interceptors.
func (c *ImageProcessClient) Interceptors() []Interceptor {
	return c.inters.ImageProcess
}

func (c *ImageProcessClient) mutate(ctx context.Context, m *ImageProcessMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ImageProcessCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ImageProcessUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ImageProcessUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ImageProcessDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ImageProcess mutation op: %q", m.Op())
	}
}

// StatClient is a client for the Stat schema.
type StatClient struct {
	config
}

// NewStatClient returns a client for the Stat from the given config.
func NewStatClient(c config) *StatClient {
	return &StatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stat.Hooks(f(g(h())))`.
func (c *StatClient) Use(hooks ...Hook) {
	c.hooks.Stat = append(c.hooks.Stat, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `stat.Intercept(f(g(h())))`.
func (c *StatClient) Intercept(interceptors ...Interceptor) {
	c.inters.Stat = append(c.inters.Stat, interceptors...)
}

// Create returns a builder for creating a Stat entity.
func (c *StatClient) Create() *StatCreate {
	mutation := newStatMutation(c.config, OpCreate)
	return &StatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Stat entities.
func (c *StatClient) CreateBulk(builders ...*StatCreate) *StatCreateBulk {
	return &StatCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StatClient) MapCreateBulk(slice any, setFunc func(*StatCreate, int)) *StatCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StatCreateBulk{err: fmt.Errorf("calling to StatClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StatCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Stat.
func (c *StatClient) Update() *StatUpdate {
	mutation := newStatMutation(c.config, OpUpdate)
	return &StatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatClient) UpdateOne(s *Stat) *StatUpdateOne {
	mutation := newStatMutation(c.config, OpUpdateOne, withStat(s))
	return &StatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatClient) UpdateOneID(id int) *StatUpdateOne {
	mutation := newStatMutation(c.config, OpUpdateOne, withStatID(id))
	return &StatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Stat.
func (c *StatClient) Delete() *StatDelete {
	mutation := newStatMutation(c.config, OpDelete)
	return &StatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatClient) DeleteOne(s *Stat) *StatDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatClient) DeleteOneID(id int) *StatDeleteOne {
	builder := c.Delete().Where(stat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatDeleteOne{builder}
}

// Query returns a query builder for Stat.
func (c *StatClient) Query() *StatQuery {
	return &StatQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStat},
		inters: c.Interceptors(),
	}
}

// Get returns a Stat entity by its id.
func (c *StatClient) Get(ctx context.Context, id int) (*Stat, error) {
	return c.Query().Where(stat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatClient) GetX(ctx context.Context, id int) *Stat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImage queries the image edge of a Stat.
func (c *StatClient) QueryImage(s *Stat) *ImageQuery {
	query := (&ImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(stat.Table, stat.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, stat.ImageTable, stat.ImageColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatClient) Hooks() []Hook {
	return c.hooks.Stat
}

// Interceptors returns the client interceptors.
func (c *StatClient) Interceptors() []Interceptor {
	return c.inters.Stat
}

func (c *StatClient) mutate(ctx context.Context, m *StatMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StatCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StatUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StatDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Stat mutation op: %q", m.Op())
	}
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `token.Intercept(f(g(h())))`.
func (c *TokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.Token = append(c.inters.Token, interceptors...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TokenClient) MapCreateBulk(slice any, setFunc func(*TokenCreate, int)) *TokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TokenCreateBulk{err: fmt.Errorf("calling to TokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id int) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id int) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeToken},
		inters: c.Interceptors(),
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id int) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id int) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// Interceptors returns the client interceptors.
func (c *TokenClient) Interceptors() []Interceptor {
	return c.inters.Token
}

func (c *TokenClient) mutate(ctx context.Context, m *TokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Token mutation op: %q", m.Op())
	}
}

// URLClient is a client for the Url schema.
type URLClient struct {
	config
}

// NewURLClient returns a client for the Url from the given config.
func NewURLClient(c config) *URLClient {
	return &URLClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `url.Hooks(f(g(h())))`.
func (c *URLClient) Use(hooks ...Hook) {
	c.hooks.Url = append(c.hooks.Url, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `url.Intercept(f(g(h())))`.
func (c *URLClient) Intercept(interceptors ...Interceptor) {
	c.inters.Url = append(c.inters.Url, interceptors...)
}

// Create returns a builder for creating a Url entity.
func (c *URLClient) Create() *URLCreate {
	mutation := newURLMutation(c.config, OpCreate)
	return &URLCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Url entities.
func (c *URLClient) CreateBulk(builders ...*URLCreate) *URLCreateBulk {
	return &URLCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *URLClient) MapCreateBulk(slice any, setFunc func(*URLCreate, int)) *URLCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &URLCreateBulk{err: fmt.Errorf("calling to URLClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*URLCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &URLCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Url.
func (c *URLClient) Update() *URLUpdate {
	mutation := newURLMutation(c.config, OpUpdate)
	return &URLUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *URLClient) UpdateOne(u *Url) *URLUpdateOne {
	mutation := newURLMutation(c.config, OpUpdateOne, withUrl(u))
	return &URLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *URLClient) UpdateOneID(id int) *URLUpdateOne {
	mutation := newURLMutation(c.config, OpUpdateOne, withUrlID(id))
	return &URLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Url.
func (c *URLClient) Delete() *URLDelete {
	mutation := newURLMutation(c.config, OpDelete)
	return &URLDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *URLClient) DeleteOne(u *Url) *URLDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *URLClient) DeleteOneID(id int) *URLDeleteOne {
	builder := c.Delete().Where(url.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &URLDeleteOne{builder}
}

// Query returns a query builder for Url.
func (c *URLClient) Query() *URLQuery {
	return &URLQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeURL},
		inters: c.Interceptors(),
	}
}

// Get returns a Url entity by its id.
func (c *URLClient) Get(ctx context.Context, id int) (*Url, error) {
	return c.Query().Where(url.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *URLClient) GetX(ctx context.Context, id int) *Url {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryErrorresult queries the errorresult edge of a Url.
func (c *URLClient) QueryErrorresult(u *Url) *ErrorResultQuery {
	query := (&ErrorResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(url.Table, url.FieldID, id),
			sqlgraph.To(errorresult.Table, errorresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, url.ErrorresultTable, url.ErrorresultColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStat queries the stat edge of a Url.
func (c *URLClient) QueryStat(u *Url) *StatQuery {
	query := (&StatClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(url.Table, url.FieldID, id),
			sqlgraph.To(stat.Table, stat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, url.StatTable, url.StatColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *URLClient) Hooks() []Hook {
	return c.hooks.Url
}

// Interceptors returns the client interceptors.
func (c *URLClient) Interceptors() []Interceptor {
	return c.inters.Url
}

func (c *URLClient) mutate(ctx context.Context, m *URLMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&URLCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&URLUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&URLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&URLDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Url mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ErrorResult, Image, ImageProcess, Stat, Token, Url []ent.Hook
	}
	inters struct {
		ErrorResult, Image, ImageProcess, Stat, Token, Url []ent.Interceptor
	}
)
