// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
)

// ExchangeRateDelete is the builder for deleting a ExchangeRate entity.
type ExchangeRateDelete struct {
	config
	hooks    []Hook
	mutation *ExchangeRateMutation
}

// Where appends a list predicates to the ExchangeRateDelete builder.
func (erd *ExchangeRateDelete) Where(ps ...predicate.ExchangeRate) *ExchangeRateDelete {
	erd.mutation.Where(ps...)
	return erd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (erd *ExchangeRateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(erd.hooks) == 0 {
		affected, err = erd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExchangeRateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			erd.mutation = mutation
			affected, err = erd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(erd.hooks) - 1; i >= 0; i-- {
			if erd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = erd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, erd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (erd *ExchangeRateDelete) ExecX(ctx context.Context) int {
	n, err := erd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (erd *ExchangeRateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: exchangerate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: exchangerate.FieldID,
			},
		},
	}
	if ps := erd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, erd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ExchangeRateDeleteOne is the builder for deleting a single ExchangeRate entity.
type ExchangeRateDeleteOne struct {
	erd *ExchangeRateDelete
}

// Exec executes the deletion query.
func (erdo *ExchangeRateDeleteOne) Exec(ctx context.Context) error {
	n, err := erdo.erd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exchangerate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (erdo *ExchangeRateDeleteOne) ExecX(ctx context.Context) {
	erdo.erd.ExecX(ctx)
}