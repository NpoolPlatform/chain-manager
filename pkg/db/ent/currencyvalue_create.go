// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyvalue"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CurrencyValueCreate is the builder for creating a CurrencyValue entity.
type CurrencyValueCreate struct {
	config
	mutation *CurrencyValueMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cvc *CurrencyValueCreate) SetCreatedAt(u uint32) *CurrencyValueCreate {
	cvc.mutation.SetCreatedAt(u)
	return cvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableCreatedAt(u *uint32) *CurrencyValueCreate {
	if u != nil {
		cvc.SetCreatedAt(*u)
	}
	return cvc
}

// SetUpdatedAt sets the "updated_at" field.
func (cvc *CurrencyValueCreate) SetUpdatedAt(u uint32) *CurrencyValueCreate {
	cvc.mutation.SetUpdatedAt(u)
	return cvc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableUpdatedAt(u *uint32) *CurrencyValueCreate {
	if u != nil {
		cvc.SetUpdatedAt(*u)
	}
	return cvc
}

// SetDeletedAt sets the "deleted_at" field.
func (cvc *CurrencyValueCreate) SetDeletedAt(u uint32) *CurrencyValueCreate {
	cvc.mutation.SetDeletedAt(u)
	return cvc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableDeletedAt(u *uint32) *CurrencyValueCreate {
	if u != nil {
		cvc.SetDeletedAt(*u)
	}
	return cvc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cvc *CurrencyValueCreate) SetCoinTypeID(u uuid.UUID) *CurrencyValueCreate {
	cvc.mutation.SetCoinTypeID(u)
	return cvc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableCoinTypeID(u *uuid.UUID) *CurrencyValueCreate {
	if u != nil {
		cvc.SetCoinTypeID(*u)
	}
	return cvc
}

// SetFeedSourceID sets the "feed_source_id" field.
func (cvc *CurrencyValueCreate) SetFeedSourceID(u uuid.UUID) *CurrencyValueCreate {
	cvc.mutation.SetFeedSourceID(u)
	return cvc
}

// SetNillableFeedSourceID sets the "feed_source_id" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableFeedSourceID(u *uuid.UUID) *CurrencyValueCreate {
	if u != nil {
		cvc.SetFeedSourceID(*u)
	}
	return cvc
}

// SetMarketValueHigh sets the "market_value_high" field.
func (cvc *CurrencyValueCreate) SetMarketValueHigh(d decimal.Decimal) *CurrencyValueCreate {
	cvc.mutation.SetMarketValueHigh(d)
	return cvc
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableMarketValueHigh(d *decimal.Decimal) *CurrencyValueCreate {
	if d != nil {
		cvc.SetMarketValueHigh(*d)
	}
	return cvc
}

// SetMarketValueLow sets the "market_value_low" field.
func (cvc *CurrencyValueCreate) SetMarketValueLow(d decimal.Decimal) *CurrencyValueCreate {
	cvc.mutation.SetMarketValueLow(d)
	return cvc
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableMarketValueLow(d *decimal.Decimal) *CurrencyValueCreate {
	if d != nil {
		cvc.SetMarketValueLow(*d)
	}
	return cvc
}

// SetID sets the "id" field.
func (cvc *CurrencyValueCreate) SetID(u uuid.UUID) *CurrencyValueCreate {
	cvc.mutation.SetID(u)
	return cvc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cvc *CurrencyValueCreate) SetNillableID(u *uuid.UUID) *CurrencyValueCreate {
	if u != nil {
		cvc.SetID(*u)
	}
	return cvc
}

// Mutation returns the CurrencyValueMutation object of the builder.
func (cvc *CurrencyValueCreate) Mutation() *CurrencyValueMutation {
	return cvc.mutation
}

// Save creates the CurrencyValue in the database.
func (cvc *CurrencyValueCreate) Save(ctx context.Context) (*CurrencyValue, error) {
	var (
		err  error
		node *CurrencyValue
	)
	if err := cvc.defaults(); err != nil {
		return nil, err
	}
	if len(cvc.hooks) == 0 {
		if err = cvc.check(); err != nil {
			return nil, err
		}
		node, err = cvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyValueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cvc.check(); err != nil {
				return nil, err
			}
			cvc.mutation = mutation
			if node, err = cvc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cvc.hooks) - 1; i >= 0; i-- {
			if cvc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cvc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cvc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CurrencyValue)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CurrencyValueMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cvc *CurrencyValueCreate) SaveX(ctx context.Context) *CurrencyValue {
	v, err := cvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cvc *CurrencyValueCreate) Exec(ctx context.Context) error {
	_, err := cvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvc *CurrencyValueCreate) ExecX(ctx context.Context) {
	if err := cvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cvc *CurrencyValueCreate) defaults() error {
	if _, ok := cvc.mutation.CreatedAt(); !ok {
		if currencyvalue.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultCreatedAt()
		cvc.mutation.SetCreatedAt(v)
	}
	if _, ok := cvc.mutation.UpdatedAt(); !ok {
		if currencyvalue.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultUpdatedAt()
		cvc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cvc.mutation.DeletedAt(); !ok {
		if currencyvalue.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultDeletedAt()
		cvc.mutation.SetDeletedAt(v)
	}
	if _, ok := cvc.mutation.CoinTypeID(); !ok {
		if currencyvalue.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultCoinTypeID()
		cvc.mutation.SetCoinTypeID(v)
	}
	if _, ok := cvc.mutation.FeedSourceID(); !ok {
		if currencyvalue.DefaultFeedSourceID == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultFeedSourceID (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultFeedSourceID()
		cvc.mutation.SetFeedSourceID(v)
	}
	if _, ok := cvc.mutation.MarketValueHigh(); !ok {
		v := currencyvalue.DefaultMarketValueHigh
		cvc.mutation.SetMarketValueHigh(v)
	}
	if _, ok := cvc.mutation.MarketValueLow(); !ok {
		v := currencyvalue.DefaultMarketValueLow
		cvc.mutation.SetMarketValueLow(v)
	}
	if _, ok := cvc.mutation.ID(); !ok {
		if currencyvalue.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized currencyvalue.DefaultID (forgotten import ent/runtime?)")
		}
		v := currencyvalue.DefaultID()
		cvc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cvc *CurrencyValueCreate) check() error {
	if _, ok := cvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CurrencyValue.created_at"`)}
	}
	if _, ok := cvc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CurrencyValue.updated_at"`)}
	}
	if _, ok := cvc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CurrencyValue.deleted_at"`)}
	}
	return nil
}

func (cvc *CurrencyValueCreate) sqlSave(ctx context.Context) (*CurrencyValue, error) {
	_node, _spec := cvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cvc *CurrencyValueCreate) createSpec() (*CurrencyValue, *sqlgraph.CreateSpec) {
	var (
		_node = &CurrencyValue{config: cvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: currencyvalue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currencyvalue.FieldID,
			},
		}
	)
	_spec.OnConflict = cvc.conflict
	if id, ok := cvc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cvc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyvalue.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cvc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyvalue.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cvc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyvalue.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cvc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyvalue.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := cvc.mutation.FeedSourceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyvalue.FieldFeedSourceID,
		})
		_node.FeedSourceID = value
	}
	if value, ok := cvc.mutation.MarketValueHigh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: currencyvalue.FieldMarketValueHigh,
		})
		_node.MarketValueHigh = value
	}
	if value, ok := cvc.mutation.MarketValueLow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: currencyvalue.FieldMarketValueLow,
		})
		_node.MarketValueLow = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyValue.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyValueUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cvc *CurrencyValueCreate) OnConflict(opts ...sql.ConflictOption) *CurrencyValueUpsertOne {
	cvc.conflict = opts
	return &CurrencyValueUpsertOne{
		create: cvc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyValue.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cvc *CurrencyValueCreate) OnConflictColumns(columns ...string) *CurrencyValueUpsertOne {
	cvc.conflict = append(cvc.conflict, sql.ConflictColumns(columns...))
	return &CurrencyValueUpsertOne{
		create: cvc,
	}
}

type (
	// CurrencyValueUpsertOne is the builder for "upsert"-ing
	//  one CurrencyValue node.
	CurrencyValueUpsertOne struct {
		create *CurrencyValueCreate
	}

	// CurrencyValueUpsert is the "OnConflict" setter.
	CurrencyValueUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyValueUpsert) SetCreatedAt(v uint32) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateCreatedAt() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyValueUpsert) AddCreatedAt(v uint32) *CurrencyValueUpsert {
	u.Add(currencyvalue.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyValueUpsert) SetUpdatedAt(v uint32) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateUpdatedAt() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyValueUpsert) AddUpdatedAt(v uint32) *CurrencyValueUpsert {
	u.Add(currencyvalue.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyValueUpsert) SetDeletedAt(v uint32) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateDeletedAt() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyValueUpsert) AddDeletedAt(v uint32) *CurrencyValueUpsert {
	u.Add(currencyvalue.FieldDeletedAt, v)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyValueUpsert) SetCoinTypeID(v uuid.UUID) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateCoinTypeID() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyValueUpsert) ClearCoinTypeID() *CurrencyValueUpsert {
	u.SetNull(currencyvalue.FieldCoinTypeID)
	return u
}

// SetFeedSourceID sets the "feed_source_id" field.
func (u *CurrencyValueUpsert) SetFeedSourceID(v uuid.UUID) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldFeedSourceID, v)
	return u
}

// UpdateFeedSourceID sets the "feed_source_id" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateFeedSourceID() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldFeedSourceID)
	return u
}

// ClearFeedSourceID clears the value of the "feed_source_id" field.
func (u *CurrencyValueUpsert) ClearFeedSourceID() *CurrencyValueUpsert {
	u.SetNull(currencyvalue.FieldFeedSourceID)
	return u
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyValueUpsert) SetMarketValueHigh(v decimal.Decimal) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldMarketValueHigh, v)
	return u
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateMarketValueHigh() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldMarketValueHigh)
	return u
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyValueUpsert) ClearMarketValueHigh() *CurrencyValueUpsert {
	u.SetNull(currencyvalue.FieldMarketValueHigh)
	return u
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyValueUpsert) SetMarketValueLow(v decimal.Decimal) *CurrencyValueUpsert {
	u.Set(currencyvalue.FieldMarketValueLow, v)
	return u
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyValueUpsert) UpdateMarketValueLow() *CurrencyValueUpsert {
	u.SetExcluded(currencyvalue.FieldMarketValueLow)
	return u
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyValueUpsert) ClearMarketValueLow() *CurrencyValueUpsert {
	u.SetNull(currencyvalue.FieldMarketValueLow)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CurrencyValue.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyvalue.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyValueUpsertOne) UpdateNewValues() *CurrencyValueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(currencyvalue.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CurrencyValue.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CurrencyValueUpsertOne) Ignore() *CurrencyValueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyValueUpsertOne) DoNothing() *CurrencyValueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyValueCreate.OnConflict
// documentation for more info.
func (u *CurrencyValueUpsertOne) Update(set func(*CurrencyValueUpsert)) *CurrencyValueUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyValueUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyValueUpsertOne) SetCreatedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyValueUpsertOne) AddCreatedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateCreatedAt() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyValueUpsertOne) SetUpdatedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyValueUpsertOne) AddUpdatedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateUpdatedAt() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyValueUpsertOne) SetDeletedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyValueUpsertOne) AddDeletedAt(v uint32) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateDeletedAt() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyValueUpsertOne) SetCoinTypeID(v uuid.UUID) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateCoinTypeID() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyValueUpsertOne) ClearCoinTypeID() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedSourceID sets the "feed_source_id" field.
func (u *CurrencyValueUpsertOne) SetFeedSourceID(v uuid.UUID) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetFeedSourceID(v)
	})
}

// UpdateFeedSourceID sets the "feed_source_id" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateFeedSourceID() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateFeedSourceID()
	})
}

// ClearFeedSourceID clears the value of the "feed_source_id" field.
func (u *CurrencyValueUpsertOne) ClearFeedSourceID() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearFeedSourceID()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyValueUpsertOne) SetMarketValueHigh(v decimal.Decimal) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateMarketValueHigh() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyValueUpsertOne) ClearMarketValueHigh() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearMarketValueHigh()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyValueUpsertOne) SetMarketValueLow(v decimal.Decimal) *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyValueUpsertOne) UpdateMarketValueLow() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyValueUpsertOne) ClearMarketValueLow() *CurrencyValueUpsertOne {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearMarketValueLow()
	})
}

// Exec executes the query.
func (u *CurrencyValueUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyValueCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyValueUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CurrencyValueUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CurrencyValueUpsertOne.ID is not supported by MySQL driver. Use CurrencyValueUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CurrencyValueUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CurrencyValueCreateBulk is the builder for creating many CurrencyValue entities in bulk.
type CurrencyValueCreateBulk struct {
	config
	builders []*CurrencyValueCreate
	conflict []sql.ConflictOption
}

// Save creates the CurrencyValue entities in the database.
func (cvcb *CurrencyValueCreateBulk) Save(ctx context.Context) ([]*CurrencyValue, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cvcb.builders))
	nodes := make([]*CurrencyValue, len(cvcb.builders))
	mutators := make([]Mutator, len(cvcb.builders))
	for i := range cvcb.builders {
		func(i int, root context.Context) {
			builder := cvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CurrencyValueMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cvcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cvcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, cvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cvcb *CurrencyValueCreateBulk) SaveX(ctx context.Context) []*CurrencyValue {
	v, err := cvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cvcb *CurrencyValueCreateBulk) Exec(ctx context.Context) error {
	_, err := cvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cvcb *CurrencyValueCreateBulk) ExecX(ctx context.Context) {
	if err := cvcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyValue.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyValueUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cvcb *CurrencyValueCreateBulk) OnConflict(opts ...sql.ConflictOption) *CurrencyValueUpsertBulk {
	cvcb.conflict = opts
	return &CurrencyValueUpsertBulk{
		create: cvcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyValue.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cvcb *CurrencyValueCreateBulk) OnConflictColumns(columns ...string) *CurrencyValueUpsertBulk {
	cvcb.conflict = append(cvcb.conflict, sql.ConflictColumns(columns...))
	return &CurrencyValueUpsertBulk{
		create: cvcb,
	}
}

// CurrencyValueUpsertBulk is the builder for "upsert"-ing
// a bulk of CurrencyValue nodes.
type CurrencyValueUpsertBulk struct {
	create *CurrencyValueCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CurrencyValue.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyvalue.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyValueUpsertBulk) UpdateNewValues() *CurrencyValueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(currencyvalue.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CurrencyValue.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CurrencyValueUpsertBulk) Ignore() *CurrencyValueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyValueUpsertBulk) DoNothing() *CurrencyValueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyValueCreateBulk.OnConflict
// documentation for more info.
func (u *CurrencyValueUpsertBulk) Update(set func(*CurrencyValueUpsert)) *CurrencyValueUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyValueUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyValueUpsertBulk) SetCreatedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyValueUpsertBulk) AddCreatedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateCreatedAt() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyValueUpsertBulk) SetUpdatedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyValueUpsertBulk) AddUpdatedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateUpdatedAt() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyValueUpsertBulk) SetDeletedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyValueUpsertBulk) AddDeletedAt(v uint32) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateDeletedAt() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyValueUpsertBulk) SetCoinTypeID(v uuid.UUID) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateCoinTypeID() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyValueUpsertBulk) ClearCoinTypeID() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedSourceID sets the "feed_source_id" field.
func (u *CurrencyValueUpsertBulk) SetFeedSourceID(v uuid.UUID) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetFeedSourceID(v)
	})
}

// UpdateFeedSourceID sets the "feed_source_id" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateFeedSourceID() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateFeedSourceID()
	})
}

// ClearFeedSourceID clears the value of the "feed_source_id" field.
func (u *CurrencyValueUpsertBulk) ClearFeedSourceID() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearFeedSourceID()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyValueUpsertBulk) SetMarketValueHigh(v decimal.Decimal) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateMarketValueHigh() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyValueUpsertBulk) ClearMarketValueHigh() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearMarketValueHigh()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyValueUpsertBulk) SetMarketValueLow(v decimal.Decimal) *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyValueUpsertBulk) UpdateMarketValueLow() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyValueUpsertBulk) ClearMarketValueLow() *CurrencyValueUpsertBulk {
	return u.Update(func(s *CurrencyValueUpsert) {
		s.ClearMarketValueLow()
	})
}

// Exec executes the query.
func (u *CurrencyValueUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CurrencyValueCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyValueCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyValueUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}