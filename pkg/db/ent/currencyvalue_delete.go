// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyvalue"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
)

// CurrencyValueDelete is the builder for deleting a CurrencyValue entity.
type CurrencyValueDelete struct {
	config
	hooks    []Hook
	mutation *CurrencyValueMutation
}

// Where appends a list predicates to the CurrencyValueDelete builder.
func (cvd *CurrencyValueDelete) Where(ps ...predicate.CurrencyValue) *CurrencyValueDelete {
	cvd.mutation.Where(ps...)
	return cvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cvd *CurrencyValueDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cvd.hooks) == 0 {
		affected, err = cvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyValueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cvd.mutation = mutation
			affected, err = cvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cvd.hooks) - 1; i >= 0; i-- {
			if cvd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvd *CurrencyValueDelete) ExecX(ctx context.Context) int {
	n, err := cvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cvd *CurrencyValueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: currencyvalue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currencyvalue.FieldID,
			},
		},
	}
	if ps := cvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CurrencyValueDeleteOne is the builder for deleting a single CurrencyValue entity.
type CurrencyValueDeleteOne struct {
	cvd *CurrencyValueDelete
}

// Exec executes the deletion query.
func (cvdo *CurrencyValueDeleteOne) Exec(ctx context.Context) error {
	n, err := cvdo.cvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{currencyvalue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cvdo *CurrencyValueDeleteOne) ExecX(ctx context.Context) {
	cvdo.cvd.ExecX(ctx)
}