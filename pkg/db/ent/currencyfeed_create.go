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
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"
	"github.com/google/uuid"
)

// CurrencyFeedCreate is the builder for creating a CurrencyFeed entity.
type CurrencyFeedCreate struct {
	config
	mutation *CurrencyFeedMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cfc *CurrencyFeedCreate) SetCreatedAt(u uint32) *CurrencyFeedCreate {
	cfc.mutation.SetCreatedAt(u)
	return cfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableCreatedAt(u *uint32) *CurrencyFeedCreate {
	if u != nil {
		cfc.SetCreatedAt(*u)
	}
	return cfc
}

// SetUpdatedAt sets the "updated_at" field.
func (cfc *CurrencyFeedCreate) SetUpdatedAt(u uint32) *CurrencyFeedCreate {
	cfc.mutation.SetUpdatedAt(u)
	return cfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableUpdatedAt(u *uint32) *CurrencyFeedCreate {
	if u != nil {
		cfc.SetUpdatedAt(*u)
	}
	return cfc
}

// SetDeletedAt sets the "deleted_at" field.
func (cfc *CurrencyFeedCreate) SetDeletedAt(u uint32) *CurrencyFeedCreate {
	cfc.mutation.SetDeletedAt(u)
	return cfc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableDeletedAt(u *uint32) *CurrencyFeedCreate {
	if u != nil {
		cfc.SetDeletedAt(*u)
	}
	return cfc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cfc *CurrencyFeedCreate) SetCoinTypeID(u uuid.UUID) *CurrencyFeedCreate {
	cfc.mutation.SetCoinTypeID(u)
	return cfc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableCoinTypeID(u *uuid.UUID) *CurrencyFeedCreate {
	if u != nil {
		cfc.SetCoinTypeID(*u)
	}
	return cfc
}

// SetFeedSource sets the "feed_source" field.
func (cfc *CurrencyFeedCreate) SetFeedSource(s string) *CurrencyFeedCreate {
	cfc.mutation.SetFeedSource(s)
	return cfc
}

// SetNillableFeedSource sets the "feed_source" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableFeedSource(s *string) *CurrencyFeedCreate {
	if s != nil {
		cfc.SetFeedSource(*s)
	}
	return cfc
}

// SetFeedType sets the "feed_type" field.
func (cfc *CurrencyFeedCreate) SetFeedType(s string) *CurrencyFeedCreate {
	cfc.mutation.SetFeedType(s)
	return cfc
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableFeedType(s *string) *CurrencyFeedCreate {
	if s != nil {
		cfc.SetFeedType(*s)
	}
	return cfc
}

// SetDisabled sets the "disabled" field.
func (cfc *CurrencyFeedCreate) SetDisabled(b bool) *CurrencyFeedCreate {
	cfc.mutation.SetDisabled(b)
	return cfc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableDisabled(b *bool) *CurrencyFeedCreate {
	if b != nil {
		cfc.SetDisabled(*b)
	}
	return cfc
}

// SetID sets the "id" field.
func (cfc *CurrencyFeedCreate) SetID(u uuid.UUID) *CurrencyFeedCreate {
	cfc.mutation.SetID(u)
	return cfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cfc *CurrencyFeedCreate) SetNillableID(u *uuid.UUID) *CurrencyFeedCreate {
	if u != nil {
		cfc.SetID(*u)
	}
	return cfc
}

// Mutation returns the CurrencyFeedMutation object of the builder.
func (cfc *CurrencyFeedCreate) Mutation() *CurrencyFeedMutation {
	return cfc.mutation
}

// Save creates the CurrencyFeed in the database.
func (cfc *CurrencyFeedCreate) Save(ctx context.Context) (*CurrencyFeed, error) {
	var (
		err  error
		node *CurrencyFeed
	)
	if err := cfc.defaults(); err != nil {
		return nil, err
	}
	if len(cfc.hooks) == 0 {
		if err = cfc.check(); err != nil {
			return nil, err
		}
		node, err = cfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyFeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cfc.check(); err != nil {
				return nil, err
			}
			cfc.mutation = mutation
			if node, err = cfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cfc.hooks) - 1; i >= 0; i-- {
			if cfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cfc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CurrencyFeed)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CurrencyFeedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cfc *CurrencyFeedCreate) SaveX(ctx context.Context) *CurrencyFeed {
	v, err := cfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cfc *CurrencyFeedCreate) Exec(ctx context.Context) error {
	_, err := cfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfc *CurrencyFeedCreate) ExecX(ctx context.Context) {
	if err := cfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfc *CurrencyFeedCreate) defaults() error {
	if _, ok := cfc.mutation.CreatedAt(); !ok {
		if currencyfeed.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := currencyfeed.DefaultCreatedAt()
		cfc.mutation.SetCreatedAt(v)
	}
	if _, ok := cfc.mutation.UpdatedAt(); !ok {
		if currencyfeed.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currencyfeed.DefaultUpdatedAt()
		cfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cfc.mutation.DeletedAt(); !ok {
		if currencyfeed.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := currencyfeed.DefaultDeletedAt()
		cfc.mutation.SetDeletedAt(v)
	}
	if _, ok := cfc.mutation.CoinTypeID(); !ok {
		if currencyfeed.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := currencyfeed.DefaultCoinTypeID()
		cfc.mutation.SetCoinTypeID(v)
	}
	if _, ok := cfc.mutation.FeedSource(); !ok {
		v := currencyfeed.DefaultFeedSource
		cfc.mutation.SetFeedSource(v)
	}
	if _, ok := cfc.mutation.FeedType(); !ok {
		v := currencyfeed.DefaultFeedType
		cfc.mutation.SetFeedType(v)
	}
	if _, ok := cfc.mutation.Disabled(); !ok {
		v := currencyfeed.DefaultDisabled
		cfc.mutation.SetDisabled(v)
	}
	if _, ok := cfc.mutation.ID(); !ok {
		if currencyfeed.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.DefaultID (forgotten import ent/runtime?)")
		}
		v := currencyfeed.DefaultID()
		cfc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cfc *CurrencyFeedCreate) check() error {
	if _, ok := cfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CurrencyFeed.created_at"`)}
	}
	if _, ok := cfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CurrencyFeed.updated_at"`)}
	}
	if _, ok := cfc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CurrencyFeed.deleted_at"`)}
	}
	return nil
}

func (cfc *CurrencyFeedCreate) sqlSave(ctx context.Context) (*CurrencyFeed, error) {
	_node, _spec := cfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cfc.driver, _spec); err != nil {
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

func (cfc *CurrencyFeedCreate) createSpec() (*CurrencyFeed, *sqlgraph.CreateSpec) {
	var (
		_node = &CurrencyFeed{config: cfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: currencyfeed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currencyfeed.FieldID,
			},
		}
	)
	_spec.OnConflict = cfc.conflict
	if id, ok := cfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cfc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cfc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cfc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cfc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyfeed.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := cfc.mutation.FeedSource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedSource,
		})
		_node.FeedSource = value
	}
	if value, ok := cfc.mutation.FeedType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedType,
		})
		_node.FeedType = value
	}
	if value, ok := cfc.mutation.Disabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: currencyfeed.FieldDisabled,
		})
		_node.Disabled = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyFeed.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyFeedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cfc *CurrencyFeedCreate) OnConflict(opts ...sql.ConflictOption) *CurrencyFeedUpsertOne {
	cfc.conflict = opts
	return &CurrencyFeedUpsertOne{
		create: cfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cfc *CurrencyFeedCreate) OnConflictColumns(columns ...string) *CurrencyFeedUpsertOne {
	cfc.conflict = append(cfc.conflict, sql.ConflictColumns(columns...))
	return &CurrencyFeedUpsertOne{
		create: cfc,
	}
}

type (
	// CurrencyFeedUpsertOne is the builder for "upsert"-ing
	//  one CurrencyFeed node.
	CurrencyFeedUpsertOne struct {
		create *CurrencyFeedCreate
	}

	// CurrencyFeedUpsert is the "OnConflict" setter.
	CurrencyFeedUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyFeedUpsert) SetCreatedAt(v uint32) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateCreatedAt() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyFeedUpsert) AddCreatedAt(v uint32) *CurrencyFeedUpsert {
	u.Add(currencyfeed.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyFeedUpsert) SetUpdatedAt(v uint32) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateUpdatedAt() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyFeedUpsert) AddUpdatedAt(v uint32) *CurrencyFeedUpsert {
	u.Add(currencyfeed.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyFeedUpsert) SetDeletedAt(v uint32) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateDeletedAt() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyFeedUpsert) AddDeletedAt(v uint32) *CurrencyFeedUpsert {
	u.Add(currencyfeed.FieldDeletedAt, v)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyFeedUpsert) SetCoinTypeID(v uuid.UUID) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateCoinTypeID() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyFeedUpsert) ClearCoinTypeID() *CurrencyFeedUpsert {
	u.SetNull(currencyfeed.FieldCoinTypeID)
	return u
}

// SetFeedSource sets the "feed_source" field.
func (u *CurrencyFeedUpsert) SetFeedSource(v string) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldFeedSource, v)
	return u
}

// UpdateFeedSource sets the "feed_source" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateFeedSource() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldFeedSource)
	return u
}

// ClearFeedSource clears the value of the "feed_source" field.
func (u *CurrencyFeedUpsert) ClearFeedSource() *CurrencyFeedUpsert {
	u.SetNull(currencyfeed.FieldFeedSource)
	return u
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyFeedUpsert) SetFeedType(v string) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldFeedType, v)
	return u
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateFeedType() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldFeedType)
	return u
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyFeedUpsert) ClearFeedType() *CurrencyFeedUpsert {
	u.SetNull(currencyfeed.FieldFeedType)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *CurrencyFeedUpsert) SetDisabled(v bool) *CurrencyFeedUpsert {
	u.Set(currencyfeed.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *CurrencyFeedUpsert) UpdateDisabled() *CurrencyFeedUpsert {
	u.SetExcluded(currencyfeed.FieldDisabled)
	return u
}

// ClearDisabled clears the value of the "disabled" field.
func (u *CurrencyFeedUpsert) ClearDisabled() *CurrencyFeedUpsert {
	u.SetNull(currencyfeed.FieldDisabled)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CurrencyFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyFeedUpsertOne) UpdateNewValues() *CurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(currencyfeed.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CurrencyFeed.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CurrencyFeedUpsertOne) Ignore() *CurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyFeedUpsertOne) DoNothing() *CurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyFeedCreate.OnConflict
// documentation for more info.
func (u *CurrencyFeedUpsertOne) Update(set func(*CurrencyFeedUpsert)) *CurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyFeedUpsertOne) SetCreatedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyFeedUpsertOne) AddCreatedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateCreatedAt() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyFeedUpsertOne) SetUpdatedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyFeedUpsertOne) AddUpdatedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateUpdatedAt() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyFeedUpsertOne) SetDeletedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyFeedUpsertOne) AddDeletedAt(v uint32) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateDeletedAt() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyFeedUpsertOne) SetCoinTypeID(v uuid.UUID) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateCoinTypeID() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyFeedUpsertOne) ClearCoinTypeID() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedSource sets the "feed_source" field.
func (u *CurrencyFeedUpsertOne) SetFeedSource(v string) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetFeedSource(v)
	})
}

// UpdateFeedSource sets the "feed_source" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateFeedSource() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateFeedSource()
	})
}

// ClearFeedSource clears the value of the "feed_source" field.
func (u *CurrencyFeedUpsertOne) ClearFeedSource() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearFeedSource()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyFeedUpsertOne) SetFeedType(v string) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateFeedType() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyFeedUpsertOne) ClearFeedType() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearFeedType()
	})
}

// SetDisabled sets the "disabled" field.
func (u *CurrencyFeedUpsertOne) SetDisabled(v bool) *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *CurrencyFeedUpsertOne) UpdateDisabled() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *CurrencyFeedUpsertOne) ClearDisabled() *CurrencyFeedUpsertOne {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearDisabled()
	})
}

// Exec executes the query.
func (u *CurrencyFeedUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyFeedCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyFeedUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CurrencyFeedUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CurrencyFeedUpsertOne.ID is not supported by MySQL driver. Use CurrencyFeedUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CurrencyFeedUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CurrencyFeedCreateBulk is the builder for creating many CurrencyFeed entities in bulk.
type CurrencyFeedCreateBulk struct {
	config
	builders []*CurrencyFeedCreate
	conflict []sql.ConflictOption
}

// Save creates the CurrencyFeed entities in the database.
func (cfcb *CurrencyFeedCreateBulk) Save(ctx context.Context) ([]*CurrencyFeed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cfcb.builders))
	nodes := make([]*CurrencyFeed, len(cfcb.builders))
	mutators := make([]Mutator, len(cfcb.builders))
	for i := range cfcb.builders {
		func(i int, root context.Context) {
			builder := cfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CurrencyFeedMutation)
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
					_, err = mutators[i+1].Mutate(root, cfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cfcb *CurrencyFeedCreateBulk) SaveX(ctx context.Context) []*CurrencyFeed {
	v, err := cfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cfcb *CurrencyFeedCreateBulk) Exec(ctx context.Context) error {
	_, err := cfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfcb *CurrencyFeedCreateBulk) ExecX(ctx context.Context) {
	if err := cfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyFeed.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyFeedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cfcb *CurrencyFeedCreateBulk) OnConflict(opts ...sql.ConflictOption) *CurrencyFeedUpsertBulk {
	cfcb.conflict = opts
	return &CurrencyFeedUpsertBulk{
		create: cfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cfcb *CurrencyFeedCreateBulk) OnConflictColumns(columns ...string) *CurrencyFeedUpsertBulk {
	cfcb.conflict = append(cfcb.conflict, sql.ConflictColumns(columns...))
	return &CurrencyFeedUpsertBulk{
		create: cfcb,
	}
}

// CurrencyFeedUpsertBulk is the builder for "upsert"-ing
// a bulk of CurrencyFeed nodes.
type CurrencyFeedUpsertBulk struct {
	create *CurrencyFeedCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CurrencyFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyFeedUpsertBulk) UpdateNewValues() *CurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(currencyfeed.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CurrencyFeed.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CurrencyFeedUpsertBulk) Ignore() *CurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyFeedUpsertBulk) DoNothing() *CurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyFeedCreateBulk.OnConflict
// documentation for more info.
func (u *CurrencyFeedUpsertBulk) Update(set func(*CurrencyFeedUpsert)) *CurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyFeedUpsertBulk) SetCreatedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyFeedUpsertBulk) AddCreatedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateCreatedAt() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyFeedUpsertBulk) SetUpdatedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyFeedUpsertBulk) AddUpdatedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateUpdatedAt() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyFeedUpsertBulk) SetDeletedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyFeedUpsertBulk) AddDeletedAt(v uint32) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateDeletedAt() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyFeedUpsertBulk) SetCoinTypeID(v uuid.UUID) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateCoinTypeID() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyFeedUpsertBulk) ClearCoinTypeID() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedSource sets the "feed_source" field.
func (u *CurrencyFeedUpsertBulk) SetFeedSource(v string) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetFeedSource(v)
	})
}

// UpdateFeedSource sets the "feed_source" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateFeedSource() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateFeedSource()
	})
}

// ClearFeedSource clears the value of the "feed_source" field.
func (u *CurrencyFeedUpsertBulk) ClearFeedSource() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearFeedSource()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyFeedUpsertBulk) SetFeedType(v string) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateFeedType() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyFeedUpsertBulk) ClearFeedType() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearFeedType()
	})
}

// SetDisabled sets the "disabled" field.
func (u *CurrencyFeedUpsertBulk) SetDisabled(v bool) *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *CurrencyFeedUpsertBulk) UpdateDisabled() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *CurrencyFeedUpsertBulk) ClearDisabled() *CurrencyFeedUpsertBulk {
	return u.Update(func(s *CurrencyFeedUpsert) {
		s.ClearDisabled()
	})
}

// Exec executes the query.
func (u *CurrencyFeedUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CurrencyFeedCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyFeedCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyFeedUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
