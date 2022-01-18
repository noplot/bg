// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"back/ent/migrate"

	"back/ent/biouser"
	"back/ent/enttiktokuser"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// BioUser is the client for interacting with the BioUser builders.
	BioUser *BioUserClient
	// EntTikTokUser is the client for interacting with the EntTikTokUser builders.
	EntTikTokUser *EntTikTokUserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.BioUser = NewBioUserClient(c.config)
	c.EntTikTokUser = NewEntTikTokUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		BioUser:       NewBioUserClient(cfg),
		EntTikTokUser: NewEntTikTokUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:        cfg,
		BioUser:       NewBioUserClient(cfg),
		EntTikTokUser: NewEntTikTokUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		BioUser.
//		Query().
//		Count(ctx)
//
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
	c.BioUser.Use(hooks...)
	c.EntTikTokUser.Use(hooks...)
}

// BioUserClient is a client for the BioUser schema.
type BioUserClient struct {
	config
}

// NewBioUserClient returns a client for the BioUser from the given config.
func NewBioUserClient(c config) *BioUserClient {
	return &BioUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `biouser.Hooks(f(g(h())))`.
func (c *BioUserClient) Use(hooks ...Hook) {
	c.hooks.BioUser = append(c.hooks.BioUser, hooks...)
}

// Create returns a create builder for BioUser.
func (c *BioUserClient) Create() *BioUserCreate {
	mutation := newBioUserMutation(c.config, OpCreate)
	return &BioUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BioUser entities.
func (c *BioUserClient) CreateBulk(builders ...*BioUserCreate) *BioUserCreateBulk {
	return &BioUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BioUser.
func (c *BioUserClient) Update() *BioUserUpdate {
	mutation := newBioUserMutation(c.config, OpUpdate)
	return &BioUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BioUserClient) UpdateOne(bu *BioUser) *BioUserUpdateOne {
	mutation := newBioUserMutation(c.config, OpUpdateOne, withBioUser(bu))
	return &BioUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BioUserClient) UpdateOneID(id int) *BioUserUpdateOne {
	mutation := newBioUserMutation(c.config, OpUpdateOne, withBioUserID(id))
	return &BioUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BioUser.
func (c *BioUserClient) Delete() *BioUserDelete {
	mutation := newBioUserMutation(c.config, OpDelete)
	return &BioUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BioUserClient) DeleteOne(bu *BioUser) *BioUserDeleteOne {
	return c.DeleteOneID(bu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BioUserClient) DeleteOneID(id int) *BioUserDeleteOne {
	builder := c.Delete().Where(biouser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BioUserDeleteOne{builder}
}

// Query returns a query builder for BioUser.
func (c *BioUserClient) Query() *BioUserQuery {
	return &BioUserQuery{
		config: c.config,
	}
}

// Get returns a BioUser entity by its id.
func (c *BioUserClient) Get(ctx context.Context, id int) (*BioUser, error) {
	return c.Query().Where(biouser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BioUserClient) GetX(ctx context.Context, id int) *BioUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BioUserClient) Hooks() []Hook {
	return c.hooks.BioUser
}

// EntTikTokUserClient is a client for the EntTikTokUser schema.
type EntTikTokUserClient struct {
	config
}

// NewEntTikTokUserClient returns a client for the EntTikTokUser from the given config.
func NewEntTikTokUserClient(c config) *EntTikTokUserClient {
	return &EntTikTokUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `enttiktokuser.Hooks(f(g(h())))`.
func (c *EntTikTokUserClient) Use(hooks ...Hook) {
	c.hooks.EntTikTokUser = append(c.hooks.EntTikTokUser, hooks...)
}

// Create returns a create builder for EntTikTokUser.
func (c *EntTikTokUserClient) Create() *EntTikTokUserCreate {
	mutation := newEntTikTokUserMutation(c.config, OpCreate)
	return &EntTikTokUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EntTikTokUser entities.
func (c *EntTikTokUserClient) CreateBulk(builders ...*EntTikTokUserCreate) *EntTikTokUserCreateBulk {
	return &EntTikTokUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EntTikTokUser.
func (c *EntTikTokUserClient) Update() *EntTikTokUserUpdate {
	mutation := newEntTikTokUserMutation(c.config, OpUpdate)
	return &EntTikTokUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EntTikTokUserClient) UpdateOne(ettu *EntTikTokUser) *EntTikTokUserUpdateOne {
	mutation := newEntTikTokUserMutation(c.config, OpUpdateOne, withEntTikTokUser(ettu))
	return &EntTikTokUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EntTikTokUserClient) UpdateOneID(id int) *EntTikTokUserUpdateOne {
	mutation := newEntTikTokUserMutation(c.config, OpUpdateOne, withEntTikTokUserID(id))
	return &EntTikTokUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EntTikTokUser.
func (c *EntTikTokUserClient) Delete() *EntTikTokUserDelete {
	mutation := newEntTikTokUserMutation(c.config, OpDelete)
	return &EntTikTokUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EntTikTokUserClient) DeleteOne(ettu *EntTikTokUser) *EntTikTokUserDeleteOne {
	return c.DeleteOneID(ettu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EntTikTokUserClient) DeleteOneID(id int) *EntTikTokUserDeleteOne {
	builder := c.Delete().Where(enttiktokuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EntTikTokUserDeleteOne{builder}
}

// Query returns a query builder for EntTikTokUser.
func (c *EntTikTokUserClient) Query() *EntTikTokUserQuery {
	return &EntTikTokUserQuery{
		config: c.config,
	}
}

// Get returns a EntTikTokUser entity by its id.
func (c *EntTikTokUserClient) Get(ctx context.Context, id int) (*EntTikTokUser, error) {
	return c.Query().Where(enttiktokuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EntTikTokUserClient) GetX(ctx context.Context, id int) *EntTikTokUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EntTikTokUserClient) Hooks() []Hook {
	return c.hooks.EntTikTokUser
}
