// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTran = "Tran"
)

// TranMutation represents an operation that mutates the Tran nodes in the graph.
type TranMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *uint32
	addcreated_at *int32
	updated_at    *uint32
	addupdated_at *int32
	deleted_at    *uint32
	adddeleted_at *int32
	app_id        *uuid.UUID
	user_id       *uuid.UUID
	coin_type_id  *uuid.UUID
	incoming      *decimal.Decimal
	locked        *decimal.Decimal
	outcoming     *decimal.Decimal
	spendable     *decimal.Decimal
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Tran, error)
	predicates    []predicate.Tran
}

var _ ent.Mutation = (*TranMutation)(nil)

// tranOption allows management of the mutation configuration using functional options.
type tranOption func(*TranMutation)

// newTranMutation creates new mutation for the Tran entity.
func newTranMutation(c config, op Op, opts ...tranOption) *TranMutation {
	m := &TranMutation{
		config:        c,
		op:            op,
		typ:           TypeTran,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTranID sets the ID field of the mutation.
func withTranID(id uuid.UUID) tranOption {
	return func(m *TranMutation) {
		var (
			err   error
			once  sync.Once
			value *Tran
		)
		m.oldValue = func(ctx context.Context) (*Tran, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tran.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTran sets the old Tran of the mutation.
func withTran(node *Tran) tranOption {
	return func(m *TranMutation) {
		m.oldValue = func(context.Context) (*Tran, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TranMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TranMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Tran entities.
func (m *TranMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TranMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TranMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tran.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TranMutation) SetCreatedAt(u uint32) {
	m.created_at = &u
	m.addcreated_at = nil
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TranMutation) CreatedAt() (r uint32, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldCreatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// AddCreatedAt adds u to the "created_at" field.
func (m *TranMutation) AddCreatedAt(u int32) {
	if m.addcreated_at != nil {
		*m.addcreated_at += u
	} else {
		m.addcreated_at = &u
	}
}

// AddedCreatedAt returns the value that was added to the "created_at" field in this mutation.
func (m *TranMutation) AddedCreatedAt() (r int32, exists bool) {
	v := m.addcreated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TranMutation) ResetCreatedAt() {
	m.created_at = nil
	m.addcreated_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TranMutation) SetUpdatedAt(u uint32) {
	m.updated_at = &u
	m.addupdated_at = nil
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TranMutation) UpdatedAt() (r uint32, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldUpdatedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// AddUpdatedAt adds u to the "updated_at" field.
func (m *TranMutation) AddUpdatedAt(u int32) {
	if m.addupdated_at != nil {
		*m.addupdated_at += u
	} else {
		m.addupdated_at = &u
	}
}

// AddedUpdatedAt returns the value that was added to the "updated_at" field in this mutation.
func (m *TranMutation) AddedUpdatedAt() (r int32, exists bool) {
	v := m.addupdated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TranMutation) ResetUpdatedAt() {
	m.updated_at = nil
	m.addupdated_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *TranMutation) SetDeletedAt(u uint32) {
	m.deleted_at = &u
	m.adddeleted_at = nil
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *TranMutation) DeletedAt() (r uint32, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldDeletedAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// AddDeletedAt adds u to the "deleted_at" field.
func (m *TranMutation) AddDeletedAt(u int32) {
	if m.adddeleted_at != nil {
		*m.adddeleted_at += u
	} else {
		m.adddeleted_at = &u
	}
}

// AddedDeletedAt returns the value that was added to the "deleted_at" field in this mutation.
func (m *TranMutation) AddedDeletedAt() (r int32, exists bool) {
	v := m.adddeleted_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *TranMutation) ResetDeletedAt() {
	m.deleted_at = nil
	m.adddeleted_at = nil
}

// SetAppID sets the "app_id" field.
func (m *TranMutation) SetAppID(u uuid.UUID) {
	m.app_id = &u
}

// AppID returns the value of the "app_id" field in the mutation.
func (m *TranMutation) AppID() (r uuid.UUID, exists bool) {
	v := m.app_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAppID returns the old "app_id" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldAppID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAppID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAppID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAppID: %w", err)
	}
	return oldValue.AppID, nil
}

// ClearAppID clears the value of the "app_id" field.
func (m *TranMutation) ClearAppID() {
	m.app_id = nil
	m.clearedFields[tran.FieldAppID] = struct{}{}
}

// AppIDCleared returns if the "app_id" field was cleared in this mutation.
func (m *TranMutation) AppIDCleared() bool {
	_, ok := m.clearedFields[tran.FieldAppID]
	return ok
}

// ResetAppID resets all changes to the "app_id" field.
func (m *TranMutation) ResetAppID() {
	m.app_id = nil
	delete(m.clearedFields, tran.FieldAppID)
}

// SetUserID sets the "user_id" field.
func (m *TranMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *TranMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ClearUserID clears the value of the "user_id" field.
func (m *TranMutation) ClearUserID() {
	m.user_id = nil
	m.clearedFields[tran.FieldUserID] = struct{}{}
}

// UserIDCleared returns if the "user_id" field was cleared in this mutation.
func (m *TranMutation) UserIDCleared() bool {
	_, ok := m.clearedFields[tran.FieldUserID]
	return ok
}

// ResetUserID resets all changes to the "user_id" field.
func (m *TranMutation) ResetUserID() {
	m.user_id = nil
	delete(m.clearedFields, tran.FieldUserID)
}

// SetCoinTypeID sets the "coin_type_id" field.
func (m *TranMutation) SetCoinTypeID(u uuid.UUID) {
	m.coin_type_id = &u
}

// CoinTypeID returns the value of the "coin_type_id" field in the mutation.
func (m *TranMutation) CoinTypeID() (r uuid.UUID, exists bool) {
	v := m.coin_type_id
	if v == nil {
		return
	}
	return *v, true
}

// OldCoinTypeID returns the old "coin_type_id" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldCoinTypeID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCoinTypeID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCoinTypeID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCoinTypeID: %w", err)
	}
	return oldValue.CoinTypeID, nil
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (m *TranMutation) ClearCoinTypeID() {
	m.coin_type_id = nil
	m.clearedFields[tran.FieldCoinTypeID] = struct{}{}
}

// CoinTypeIDCleared returns if the "coin_type_id" field was cleared in this mutation.
func (m *TranMutation) CoinTypeIDCleared() bool {
	_, ok := m.clearedFields[tran.FieldCoinTypeID]
	return ok
}

// ResetCoinTypeID resets all changes to the "coin_type_id" field.
func (m *TranMutation) ResetCoinTypeID() {
	m.coin_type_id = nil
	delete(m.clearedFields, tran.FieldCoinTypeID)
}

// SetIncoming sets the "incoming" field.
func (m *TranMutation) SetIncoming(d decimal.Decimal) {
	m.incoming = &d
}

// Incoming returns the value of the "incoming" field in the mutation.
func (m *TranMutation) Incoming() (r decimal.Decimal, exists bool) {
	v := m.incoming
	if v == nil {
		return
	}
	return *v, true
}

// OldIncoming returns the old "incoming" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldIncoming(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIncoming is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIncoming requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIncoming: %w", err)
	}
	return oldValue.Incoming, nil
}

// ClearIncoming clears the value of the "incoming" field.
func (m *TranMutation) ClearIncoming() {
	m.incoming = nil
	m.clearedFields[tran.FieldIncoming] = struct{}{}
}

// IncomingCleared returns if the "incoming" field was cleared in this mutation.
func (m *TranMutation) IncomingCleared() bool {
	_, ok := m.clearedFields[tran.FieldIncoming]
	return ok
}

// ResetIncoming resets all changes to the "incoming" field.
func (m *TranMutation) ResetIncoming() {
	m.incoming = nil
	delete(m.clearedFields, tran.FieldIncoming)
}

// SetLocked sets the "locked" field.
func (m *TranMutation) SetLocked(d decimal.Decimal) {
	m.locked = &d
}

// Locked returns the value of the "locked" field in the mutation.
func (m *TranMutation) Locked() (r decimal.Decimal, exists bool) {
	v := m.locked
	if v == nil {
		return
	}
	return *v, true
}

// OldLocked returns the old "locked" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldLocked(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocked is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocked requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocked: %w", err)
	}
	return oldValue.Locked, nil
}

// ClearLocked clears the value of the "locked" field.
func (m *TranMutation) ClearLocked() {
	m.locked = nil
	m.clearedFields[tran.FieldLocked] = struct{}{}
}

// LockedCleared returns if the "locked" field was cleared in this mutation.
func (m *TranMutation) LockedCleared() bool {
	_, ok := m.clearedFields[tran.FieldLocked]
	return ok
}

// ResetLocked resets all changes to the "locked" field.
func (m *TranMutation) ResetLocked() {
	m.locked = nil
	delete(m.clearedFields, tran.FieldLocked)
}

// SetOutcoming sets the "outcoming" field.
func (m *TranMutation) SetOutcoming(d decimal.Decimal) {
	m.outcoming = &d
}

// Outcoming returns the value of the "outcoming" field in the mutation.
func (m *TranMutation) Outcoming() (r decimal.Decimal, exists bool) {
	v := m.outcoming
	if v == nil {
		return
	}
	return *v, true
}

// OldOutcoming returns the old "outcoming" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldOutcoming(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOutcoming is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOutcoming requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOutcoming: %w", err)
	}
	return oldValue.Outcoming, nil
}

// ClearOutcoming clears the value of the "outcoming" field.
func (m *TranMutation) ClearOutcoming() {
	m.outcoming = nil
	m.clearedFields[tran.FieldOutcoming] = struct{}{}
}

// OutcomingCleared returns if the "outcoming" field was cleared in this mutation.
func (m *TranMutation) OutcomingCleared() bool {
	_, ok := m.clearedFields[tran.FieldOutcoming]
	return ok
}

// ResetOutcoming resets all changes to the "outcoming" field.
func (m *TranMutation) ResetOutcoming() {
	m.outcoming = nil
	delete(m.clearedFields, tran.FieldOutcoming)
}

// SetSpendable sets the "spendable" field.
func (m *TranMutation) SetSpendable(d decimal.Decimal) {
	m.spendable = &d
}

// Spendable returns the value of the "spendable" field in the mutation.
func (m *TranMutation) Spendable() (r decimal.Decimal, exists bool) {
	v := m.spendable
	if v == nil {
		return
	}
	return *v, true
}

// OldSpendable returns the old "spendable" field's value of the Tran entity.
// If the Tran object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TranMutation) OldSpendable(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSpendable is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSpendable requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpendable: %w", err)
	}
	return oldValue.Spendable, nil
}

// ClearSpendable clears the value of the "spendable" field.
func (m *TranMutation) ClearSpendable() {
	m.spendable = nil
	m.clearedFields[tran.FieldSpendable] = struct{}{}
}

// SpendableCleared returns if the "spendable" field was cleared in this mutation.
func (m *TranMutation) SpendableCleared() bool {
	_, ok := m.clearedFields[tran.FieldSpendable]
	return ok
}

// ResetSpendable resets all changes to the "spendable" field.
func (m *TranMutation) ResetSpendable() {
	m.spendable = nil
	delete(m.clearedFields, tran.FieldSpendable)
}

// Where appends a list predicates to the TranMutation builder.
func (m *TranMutation) Where(ps ...predicate.Tran) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TranMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Tran).
func (m *TranMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TranMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.created_at != nil {
		fields = append(fields, tran.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, tran.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, tran.FieldDeletedAt)
	}
	if m.app_id != nil {
		fields = append(fields, tran.FieldAppID)
	}
	if m.user_id != nil {
		fields = append(fields, tran.FieldUserID)
	}
	if m.coin_type_id != nil {
		fields = append(fields, tran.FieldCoinTypeID)
	}
	if m.incoming != nil {
		fields = append(fields, tran.FieldIncoming)
	}
	if m.locked != nil {
		fields = append(fields, tran.FieldLocked)
	}
	if m.outcoming != nil {
		fields = append(fields, tran.FieldOutcoming)
	}
	if m.spendable != nil {
		fields = append(fields, tran.FieldSpendable)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TranMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tran.FieldCreatedAt:
		return m.CreatedAt()
	case tran.FieldUpdatedAt:
		return m.UpdatedAt()
	case tran.FieldDeletedAt:
		return m.DeletedAt()
	case tran.FieldAppID:
		return m.AppID()
	case tran.FieldUserID:
		return m.UserID()
	case tran.FieldCoinTypeID:
		return m.CoinTypeID()
	case tran.FieldIncoming:
		return m.Incoming()
	case tran.FieldLocked:
		return m.Locked()
	case tran.FieldOutcoming:
		return m.Outcoming()
	case tran.FieldSpendable:
		return m.Spendable()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TranMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tran.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case tran.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case tran.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case tran.FieldAppID:
		return m.OldAppID(ctx)
	case tran.FieldUserID:
		return m.OldUserID(ctx)
	case tran.FieldCoinTypeID:
		return m.OldCoinTypeID(ctx)
	case tran.FieldIncoming:
		return m.OldIncoming(ctx)
	case tran.FieldLocked:
		return m.OldLocked(ctx)
	case tran.FieldOutcoming:
		return m.OldOutcoming(ctx)
	case tran.FieldSpendable:
		return m.OldSpendable(ctx)
	}
	return nil, fmt.Errorf("unknown Tran field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TranMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tran.FieldCreatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case tran.FieldUpdatedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case tran.FieldDeletedAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case tran.FieldAppID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAppID(v)
		return nil
	case tran.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case tran.FieldCoinTypeID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCoinTypeID(v)
		return nil
	case tran.FieldIncoming:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIncoming(v)
		return nil
	case tran.FieldLocked:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocked(v)
		return nil
	case tran.FieldOutcoming:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOutcoming(v)
		return nil
	case tran.FieldSpendable:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpendable(v)
		return nil
	}
	return fmt.Errorf("unknown Tran field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TranMutation) AddedFields() []string {
	var fields []string
	if m.addcreated_at != nil {
		fields = append(fields, tran.FieldCreatedAt)
	}
	if m.addupdated_at != nil {
		fields = append(fields, tran.FieldUpdatedAt)
	}
	if m.adddeleted_at != nil {
		fields = append(fields, tran.FieldDeletedAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TranMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case tran.FieldCreatedAt:
		return m.AddedCreatedAt()
	case tran.FieldUpdatedAt:
		return m.AddedUpdatedAt()
	case tran.FieldDeletedAt:
		return m.AddedDeletedAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TranMutation) AddField(name string, value ent.Value) error {
	switch name {
	case tran.FieldCreatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedAt(v)
		return nil
	case tran.FieldUpdatedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdatedAt(v)
		return nil
	case tran.FieldDeletedAt:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Tran numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TranMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(tran.FieldAppID) {
		fields = append(fields, tran.FieldAppID)
	}
	if m.FieldCleared(tran.FieldUserID) {
		fields = append(fields, tran.FieldUserID)
	}
	if m.FieldCleared(tran.FieldCoinTypeID) {
		fields = append(fields, tran.FieldCoinTypeID)
	}
	if m.FieldCleared(tran.FieldIncoming) {
		fields = append(fields, tran.FieldIncoming)
	}
	if m.FieldCleared(tran.FieldLocked) {
		fields = append(fields, tran.FieldLocked)
	}
	if m.FieldCleared(tran.FieldOutcoming) {
		fields = append(fields, tran.FieldOutcoming)
	}
	if m.FieldCleared(tran.FieldSpendable) {
		fields = append(fields, tran.FieldSpendable)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TranMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TranMutation) ClearField(name string) error {
	switch name {
	case tran.FieldAppID:
		m.ClearAppID()
		return nil
	case tran.FieldUserID:
		m.ClearUserID()
		return nil
	case tran.FieldCoinTypeID:
		m.ClearCoinTypeID()
		return nil
	case tran.FieldIncoming:
		m.ClearIncoming()
		return nil
	case tran.FieldLocked:
		m.ClearLocked()
		return nil
	case tran.FieldOutcoming:
		m.ClearOutcoming()
		return nil
	case tran.FieldSpendable:
		m.ClearSpendable()
		return nil
	}
	return fmt.Errorf("unknown Tran nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TranMutation) ResetField(name string) error {
	switch name {
	case tran.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case tran.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case tran.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case tran.FieldAppID:
		m.ResetAppID()
		return nil
	case tran.FieldUserID:
		m.ResetUserID()
		return nil
	case tran.FieldCoinTypeID:
		m.ResetCoinTypeID()
		return nil
	case tran.FieldIncoming:
		m.ResetIncoming()
		return nil
	case tran.FieldLocked:
		m.ResetLocked()
		return nil
	case tran.FieldOutcoming:
		m.ResetOutcoming()
		return nil
	case tran.FieldSpendable:
		m.ResetSpendable()
		return nil
	}
	return fmt.Errorf("unknown Tran field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TranMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TranMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TranMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TranMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TranMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TranMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TranMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Tran unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TranMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Tran edge %s", name)
}
