// Code generated by entc, DO NOT EDIT.

package ent

import (
	"back/ent/enttiktokuser"
	"back/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntTikTokUserUpdate is the builder for updating EntTikTokUser entities.
type EntTikTokUserUpdate struct {
	config
	hooks    []Hook
	mutation *EntTikTokUserMutation
}

// Where appends a list predicates to the EntTikTokUserUpdate builder.
func (ettuu *EntTikTokUserUpdate) Where(ps ...predicate.EntTikTokUser) *EntTikTokUserUpdate {
	ettuu.mutation.Where(ps...)
	return ettuu
}

// SetName sets the "name" field.
func (ettuu *EntTikTokUserUpdate) SetName(s string) *EntTikTokUserUpdate {
	ettuu.mutation.SetName(s)
	return ettuu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ettuu *EntTikTokUserUpdate) SetNillableName(s *string) *EntTikTokUserUpdate {
	if s != nil {
		ettuu.SetName(*s)
	}
	return ettuu
}

// Mutation returns the EntTikTokUserMutation object of the builder.
func (ettuu *EntTikTokUserUpdate) Mutation() *EntTikTokUserMutation {
	return ettuu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ettuu *EntTikTokUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ettuu.hooks) == 0 {
		affected, err = ettuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntTikTokUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ettuu.mutation = mutation
			affected, err = ettuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ettuu.hooks) - 1; i >= 0; i-- {
			if ettuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ettuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ettuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ettuu *EntTikTokUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ettuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ettuu *EntTikTokUserUpdate) Exec(ctx context.Context) error {
	_, err := ettuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ettuu *EntTikTokUserUpdate) ExecX(ctx context.Context) {
	if err := ettuu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ettuu *EntTikTokUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   enttiktokuser.Table,
			Columns: enttiktokuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enttiktokuser.FieldID,
			},
		},
	}
	if ps := ettuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ettuu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enttiktokuser.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ettuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enttiktokuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EntTikTokUserUpdateOne is the builder for updating a single EntTikTokUser entity.
type EntTikTokUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntTikTokUserMutation
}

// SetName sets the "name" field.
func (ettuuo *EntTikTokUserUpdateOne) SetName(s string) *EntTikTokUserUpdateOne {
	ettuuo.mutation.SetName(s)
	return ettuuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ettuuo *EntTikTokUserUpdateOne) SetNillableName(s *string) *EntTikTokUserUpdateOne {
	if s != nil {
		ettuuo.SetName(*s)
	}
	return ettuuo
}

// Mutation returns the EntTikTokUserMutation object of the builder.
func (ettuuo *EntTikTokUserUpdateOne) Mutation() *EntTikTokUserMutation {
	return ettuuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ettuuo *EntTikTokUserUpdateOne) Select(field string, fields ...string) *EntTikTokUserUpdateOne {
	ettuuo.fields = append([]string{field}, fields...)
	return ettuuo
}

// Save executes the query and returns the updated EntTikTokUser entity.
func (ettuuo *EntTikTokUserUpdateOne) Save(ctx context.Context) (*EntTikTokUser, error) {
	var (
		err  error
		node *EntTikTokUser
	)
	if len(ettuuo.hooks) == 0 {
		node, err = ettuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntTikTokUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ettuuo.mutation = mutation
			node, err = ettuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ettuuo.hooks) - 1; i >= 0; i-- {
			if ettuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ettuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ettuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ettuuo *EntTikTokUserUpdateOne) SaveX(ctx context.Context) *EntTikTokUser {
	node, err := ettuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ettuuo *EntTikTokUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ettuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ettuuo *EntTikTokUserUpdateOne) ExecX(ctx context.Context) {
	if err := ettuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ettuuo *EntTikTokUserUpdateOne) sqlSave(ctx context.Context) (_node *EntTikTokUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   enttiktokuser.Table,
			Columns: enttiktokuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enttiktokuser.FieldID,
			},
		},
	}
	id, ok := ettuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EntTikTokUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ettuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enttiktokuser.FieldID)
		for _, f := range fields {
			if !enttiktokuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != enttiktokuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ettuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ettuuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enttiktokuser.FieldName,
		})
	}
	_node = &EntTikTokUser{config: ettuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ettuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enttiktokuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
