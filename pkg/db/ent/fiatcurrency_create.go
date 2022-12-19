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
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrency"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FiatCurrencyCreate is the builder for creating a FiatCurrency entity.
type FiatCurrencyCreate struct {
	config
	mutation *FiatCurrencyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fcc *FiatCurrencyCreate) SetCreatedAt(u uint32) *FiatCurrencyCreate {
	fcc.mutation.SetCreatedAt(u)
	return fcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableCreatedAt(u *uint32) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetCreatedAt(*u)
	}
	return fcc
}

// SetUpdatedAt sets the "updated_at" field.
func (fcc *FiatCurrencyCreate) SetUpdatedAt(u uint32) *FiatCurrencyCreate {
	fcc.mutation.SetUpdatedAt(u)
	return fcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableUpdatedAt(u *uint32) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetUpdatedAt(*u)
	}
	return fcc
}

// SetDeletedAt sets the "deleted_at" field.
func (fcc *FiatCurrencyCreate) SetDeletedAt(u uint32) *FiatCurrencyCreate {
	fcc.mutation.SetDeletedAt(u)
	return fcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableDeletedAt(u *uint32) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetDeletedAt(*u)
	}
	return fcc
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (fcc *FiatCurrencyCreate) SetFiatCurrencyTypeID(u uuid.UUID) *FiatCurrencyCreate {
	fcc.mutation.SetFiatCurrencyTypeID(u)
	return fcc
}

// SetNillableFiatCurrencyTypeID sets the "fiat_currency_type_id" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableFiatCurrencyTypeID(u *uuid.UUID) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetFiatCurrencyTypeID(*u)
	}
	return fcc
}

// SetFeedType sets the "feed_type" field.
func (fcc *FiatCurrencyCreate) SetFeedType(s string) *FiatCurrencyCreate {
	fcc.mutation.SetFeedType(s)
	return fcc
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableFeedType(s *string) *FiatCurrencyCreate {
	if s != nil {
		fcc.SetFeedType(*s)
	}
	return fcc
}

// SetMarketValueLow sets the "market_value_low" field.
func (fcc *FiatCurrencyCreate) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyCreate {
	fcc.mutation.SetMarketValueLow(d)
	return fcc
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyCreate {
	if d != nil {
		fcc.SetMarketValueLow(*d)
	}
	return fcc
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fcc *FiatCurrencyCreate) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyCreate {
	fcc.mutation.SetMarketValueHigh(d)
	return fcc
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyCreate {
	if d != nil {
		fcc.SetMarketValueHigh(*d)
	}
	return fcc
}

// SetID sets the "id" field.
func (fcc *FiatCurrencyCreate) SetID(u uuid.UUID) *FiatCurrencyCreate {
	fcc.mutation.SetID(u)
	return fcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fcc *FiatCurrencyCreate) SetNillableID(u *uuid.UUID) *FiatCurrencyCreate {
	if u != nil {
		fcc.SetID(*u)
	}
	return fcc
}

// Mutation returns the FiatCurrencyMutation object of the builder.
func (fcc *FiatCurrencyCreate) Mutation() *FiatCurrencyMutation {
	return fcc.mutation
}

// Save creates the FiatCurrency in the database.
func (fcc *FiatCurrencyCreate) Save(ctx context.Context) (*FiatCurrency, error) {
	var (
		err  error
		node *FiatCurrency
	)
	if err := fcc.defaults(); err != nil {
		return nil, err
	}
	if len(fcc.hooks) == 0 {
		if err = fcc.check(); err != nil {
			return nil, err
		}
		node, err = fcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fcc.check(); err != nil {
				return nil, err
			}
			fcc.mutation = mutation
			if node, err = fcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fcc.hooks) - 1; i >= 0; i-- {
			if fcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FiatCurrency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FiatCurrencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fcc *FiatCurrencyCreate) SaveX(ctx context.Context) *FiatCurrency {
	v, err := fcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcc *FiatCurrencyCreate) Exec(ctx context.Context) error {
	_, err := fcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcc *FiatCurrencyCreate) ExecX(ctx context.Context) {
	if err := fcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcc *FiatCurrencyCreate) defaults() error {
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		if fiatcurrency.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.DefaultCreatedAt()
		fcc.mutation.SetCreatedAt(v)
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		if fiatcurrency.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.DefaultUpdatedAt()
		fcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fcc.mutation.DeletedAt(); !ok {
		if fiatcurrency.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.DefaultDeletedAt()
		fcc.mutation.SetDeletedAt(v)
	}
	if _, ok := fcc.mutation.FiatCurrencyTypeID(); !ok {
		if fiatcurrency.DefaultFiatCurrencyTypeID == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.DefaultFiatCurrencyTypeID (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.DefaultFiatCurrencyTypeID()
		fcc.mutation.SetFiatCurrencyTypeID(v)
	}
	if _, ok := fcc.mutation.FeedType(); !ok {
		v := fiatcurrency.DefaultFeedType
		fcc.mutation.SetFeedType(v)
	}
	if _, ok := fcc.mutation.MarketValueLow(); !ok {
		v := fiatcurrency.DefaultMarketValueLow
		fcc.mutation.SetMarketValueLow(v)
	}
	if _, ok := fcc.mutation.MarketValueHigh(); !ok {
		v := fiatcurrency.DefaultMarketValueHigh
		fcc.mutation.SetMarketValueHigh(v)
	}
	if _, ok := fcc.mutation.ID(); !ok {
		if fiatcurrency.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.DefaultID (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.DefaultID()
		fcc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fcc *FiatCurrencyCreate) check() error {
	if _, ok := fcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FiatCurrency.created_at"`)}
	}
	if _, ok := fcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FiatCurrency.updated_at"`)}
	}
	if _, ok := fcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "FiatCurrency.deleted_at"`)}
	}
	return nil
}

func (fcc *FiatCurrencyCreate) sqlSave(ctx context.Context) (*FiatCurrency, error) {
	_node, _spec := fcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fcc.driver, _spec); err != nil {
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

func (fcc *FiatCurrencyCreate) createSpec() (*FiatCurrency, *sqlgraph.CreateSpec) {
	var (
		_node = &FiatCurrency{config: fcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fiatcurrency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiatcurrency.FieldID,
			},
		}
	)
	_spec.OnConflict = fcc.conflict
	if id, ok := fcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := fcc.mutation.FiatCurrencyTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrency.FieldFiatCurrencyTypeID,
		})
		_node.FiatCurrencyTypeID = value
	}
	if value, ok := fcc.mutation.FeedType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrency.FieldFeedType,
		})
		_node.FeedType = value
	}
	if value, ok := fcc.mutation.MarketValueLow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueLow,
		})
		_node.MarketValueLow = value
	}
	if value, ok := fcc.mutation.MarketValueHigh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueHigh,
		})
		_node.MarketValueHigh = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrency.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fcc *FiatCurrencyCreate) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyUpsertOne {
	fcc.conflict = opts
	return &FiatCurrencyUpsertOne{
		create: fcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fcc *FiatCurrencyCreate) OnConflictColumns(columns ...string) *FiatCurrencyUpsertOne {
	fcc.conflict = append(fcc.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyUpsertOne{
		create: fcc,
	}
}

type (
	// FiatCurrencyUpsertOne is the builder for "upsert"-ing
	//  one FiatCurrency node.
	FiatCurrencyUpsertOne struct {
		create *FiatCurrencyCreate
	}

	// FiatCurrencyUpsert is the "OnConflict" setter.
	FiatCurrencyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyUpsert) SetCreatedAt(v uint32) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateCreatedAt() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyUpsert) AddCreatedAt(v uint32) *FiatCurrencyUpsert {
	u.Add(fiatcurrency.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyUpsert) SetUpdatedAt(v uint32) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateUpdatedAt() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyUpsert) AddUpdatedAt(v uint32) *FiatCurrencyUpsert {
	u.Add(fiatcurrency.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyUpsert) SetDeletedAt(v uint32) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateDeletedAt() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyUpsert) AddDeletedAt(v uint32) *FiatCurrencyUpsert {
	u.Add(fiatcurrency.FieldDeletedAt, v)
	return u
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsert) SetFiatCurrencyTypeID(v uuid.UUID) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldFiatCurrencyTypeID, v)
	return u
}

// UpdateFiatCurrencyTypeID sets the "fiat_currency_type_id" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateFiatCurrencyTypeID() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldFiatCurrencyTypeID)
	return u
}

// ClearFiatCurrencyTypeID clears the value of the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsert) ClearFiatCurrencyTypeID() *FiatCurrencyUpsert {
	u.SetNull(fiatcurrency.FieldFiatCurrencyTypeID)
	return u
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyUpsert) SetFeedType(v string) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldFeedType, v)
	return u
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateFeedType() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldFeedType)
	return u
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyUpsert) ClearFeedType() *FiatCurrencyUpsert {
	u.SetNull(fiatcurrency.FieldFeedType)
	return u
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyUpsert) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldMarketValueLow, v)
	return u
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateMarketValueLow() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldMarketValueLow)
	return u
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyUpsert) ClearMarketValueLow() *FiatCurrencyUpsert {
	u.SetNull(fiatcurrency.FieldMarketValueLow)
	return u
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyUpsert) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyUpsert {
	u.Set(fiatcurrency.FieldMarketValueHigh, v)
	return u
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyUpsert) UpdateMarketValueHigh() *FiatCurrencyUpsert {
	u.SetExcluded(fiatcurrency.FieldMarketValueHigh)
	return u
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyUpsert) ClearMarketValueHigh() *FiatCurrencyUpsert {
	u.SetNull(fiatcurrency.FieldMarketValueHigh)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FiatCurrency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatCurrencyUpsertOne) UpdateNewValues() *FiatCurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(fiatcurrency.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.FiatCurrency.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FiatCurrencyUpsertOne) Ignore() *FiatCurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyUpsertOne) DoNothing() *FiatCurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyCreate.OnConflict
// documentation for more info.
func (u *FiatCurrencyUpsertOne) Update(set func(*FiatCurrencyUpsert)) *FiatCurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyUpsertOne) SetCreatedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyUpsertOne) AddCreatedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateCreatedAt() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyUpsertOne) SetUpdatedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyUpsertOne) AddUpdatedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateUpdatedAt() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyUpsertOne) SetDeletedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyUpsertOne) AddDeletedAt(v uint32) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateDeletedAt() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsertOne) SetFiatCurrencyTypeID(v uuid.UUID) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetFiatCurrencyTypeID(v)
	})
}

// UpdateFiatCurrencyTypeID sets the "fiat_currency_type_id" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateFiatCurrencyTypeID() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateFiatCurrencyTypeID()
	})
}

// ClearFiatCurrencyTypeID clears the value of the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsertOne) ClearFiatCurrencyTypeID() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearFiatCurrencyTypeID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyUpsertOne) SetFeedType(v string) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateFeedType() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyUpsertOne) ClearFeedType() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyUpsertOne) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateMarketValueLow() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyUpsertOne) ClearMarketValueLow() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearMarketValueLow()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyUpsertOne) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyUpsertOne) UpdateMarketValueHigh() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyUpsertOne) ClearMarketValueHigh() *FiatCurrencyUpsertOne {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearMarketValueHigh()
	})
}

// Exec executes the query.
func (u *FiatCurrencyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCurrencyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FiatCurrencyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FiatCurrencyUpsertOne.ID is not supported by MySQL driver. Use FiatCurrencyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FiatCurrencyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FiatCurrencyCreateBulk is the builder for creating many FiatCurrency entities in bulk.
type FiatCurrencyCreateBulk struct {
	config
	builders []*FiatCurrencyCreate
	conflict []sql.ConflictOption
}

// Save creates the FiatCurrency entities in the database.
func (fccb *FiatCurrencyCreateBulk) Save(ctx context.Context) ([]*FiatCurrency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fccb.builders))
	nodes := make([]*FiatCurrency, len(fccb.builders))
	mutators := make([]Mutator, len(fccb.builders))
	for i := range fccb.builders {
		func(i int, root context.Context) {
			builder := fccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiatCurrencyMutation)
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
					_, err = mutators[i+1].Mutate(root, fccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fccb *FiatCurrencyCreateBulk) SaveX(ctx context.Context) []*FiatCurrency {
	v, err := fccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fccb *FiatCurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := fccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fccb *FiatCurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := fccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrency.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fccb *FiatCurrencyCreateBulk) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyUpsertBulk {
	fccb.conflict = opts
	return &FiatCurrencyUpsertBulk{
		create: fccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fccb *FiatCurrencyCreateBulk) OnConflictColumns(columns ...string) *FiatCurrencyUpsertBulk {
	fccb.conflict = append(fccb.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyUpsertBulk{
		create: fccb,
	}
}

// FiatCurrencyUpsertBulk is the builder for "upsert"-ing
// a bulk of FiatCurrency nodes.
type FiatCurrencyUpsertBulk struct {
	create *FiatCurrencyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FiatCurrency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatCurrencyUpsertBulk) UpdateNewValues() *FiatCurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(fiatcurrency.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FiatCurrency.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FiatCurrencyUpsertBulk) Ignore() *FiatCurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyUpsertBulk) DoNothing() *FiatCurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyCreateBulk.OnConflict
// documentation for more info.
func (u *FiatCurrencyUpsertBulk) Update(set func(*FiatCurrencyUpsert)) *FiatCurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyUpsertBulk) SetCreatedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyUpsertBulk) AddCreatedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateCreatedAt() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyUpsertBulk) SetUpdatedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyUpsertBulk) AddUpdatedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateUpdatedAt() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyUpsertBulk) SetDeletedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyUpsertBulk) AddDeletedAt(v uint32) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateDeletedAt() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsertBulk) SetFiatCurrencyTypeID(v uuid.UUID) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetFiatCurrencyTypeID(v)
	})
}

// UpdateFiatCurrencyTypeID sets the "fiat_currency_type_id" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateFiatCurrencyTypeID() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateFiatCurrencyTypeID()
	})
}

// ClearFiatCurrencyTypeID clears the value of the "fiat_currency_type_id" field.
func (u *FiatCurrencyUpsertBulk) ClearFiatCurrencyTypeID() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearFiatCurrencyTypeID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyUpsertBulk) SetFeedType(v string) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateFeedType() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyUpsertBulk) ClearFeedType() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyUpsertBulk) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateMarketValueLow() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyUpsertBulk) ClearMarketValueLow() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearMarketValueLow()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyUpsertBulk) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyUpsertBulk) UpdateMarketValueHigh() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyUpsertBulk) ClearMarketValueHigh() *FiatCurrencyUpsertBulk {
	return u.Update(func(s *FiatCurrencyUpsert) {
		s.ClearMarketValueHigh()
	})
}

// Exec executes the query.
func (u *FiatCurrencyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FiatCurrencyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCurrencyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}