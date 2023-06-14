// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"swapcore/Core/DB/ent/migrate"

	"swapcore/Core/DB/ent/network"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Network is the client for interacting with the Network builders.
	Network *NetworkClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Network = NewNetworkClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Network: NewNetworkClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Network: NewNetworkClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Network.
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
	c.Network.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Network.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *NetworkMutation:
		return c.Network.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// NetworkClient is a client for the Network schema.
type NetworkClient struct {
	config
}

// NewNetworkClient returns a client for the Network from the given config.
func NewNetworkClient(c config) *NetworkClient {
	return &NetworkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `network.Hooks(f(g(h())))`.
func (c *NetworkClient) Use(hooks ...Hook) {
	c.hooks.Network = append(c.hooks.Network, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `network.Intercept(f(g(h())))`.
func (c *NetworkClient) Intercept(interceptors ...Interceptor) {
	c.inters.Network = append(c.inters.Network, interceptors...)
}

// Create returns a builder for creating a Network entity.
func (c *NetworkClient) Create() *NetworkCreate {
	mutation := newNetworkMutation(c.config, OpCreate)
	return &NetworkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Network entities.
func (c *NetworkClient) CreateBulk(builders ...*NetworkCreate) *NetworkCreateBulk {
	return &NetworkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Network.
func (c *NetworkClient) Update() *NetworkUpdate {
	mutation := newNetworkMutation(c.config, OpUpdate)
	return &NetworkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NetworkClient) UpdateOne(n *Network) *NetworkUpdateOne {
	mutation := newNetworkMutation(c.config, OpUpdateOne, withNetwork(n))
	return &NetworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NetworkClient) UpdateOneID(id int) *NetworkUpdateOne {
	mutation := newNetworkMutation(c.config, OpUpdateOne, withNetworkID(id))
	return &NetworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Network.
func (c *NetworkClient) Delete() *NetworkDelete {
	mutation := newNetworkMutation(c.config, OpDelete)
	return &NetworkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NetworkClient) DeleteOne(n *Network) *NetworkDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NetworkClient) DeleteOneID(id int) *NetworkDeleteOne {
	builder := c.Delete().Where(network.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NetworkDeleteOne{builder}
}

// Query returns a query builder for Network.
func (c *NetworkClient) Query() *NetworkQuery {
	return &NetworkQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNetwork},
		inters: c.Interceptors(),
	}
}

// Get returns a Network entity by its id.
func (c *NetworkClient) Get(ctx context.Context, id int) (*Network, error) {
	return c.Query().Where(network.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NetworkClient) GetX(ctx context.Context, id int) *Network {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NetworkClient) Hooks() []Hook {
	return c.hooks.Network
}

// Interceptors returns the client interceptors.
func (c *NetworkClient) Interceptors() []Interceptor {
	return c.inters.Network
}

func (c *NetworkClient) mutate(ctx context.Context, m *NetworkMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NetworkCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NetworkUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NetworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NetworkDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Network mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Network []ent.Hook
	}
	inters struct {
		Network []ent.Interceptor
	}
)
