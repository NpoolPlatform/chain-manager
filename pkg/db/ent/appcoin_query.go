// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppCoinQuery is the builder for querying AppCoin entities.
type AppCoinQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppCoin
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppCoinQuery builder.
func (acq *AppCoinQuery) Where(ps ...predicate.AppCoin) *AppCoinQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit adds a limit step to the query.
func (acq *AppCoinQuery) Limit(limit int) *AppCoinQuery {
	acq.limit = &limit
	return acq
}

// Offset adds an offset step to the query.
func (acq *AppCoinQuery) Offset(offset int) *AppCoinQuery {
	acq.offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AppCoinQuery) Unique(unique bool) *AppCoinQuery {
	acq.unique = &unique
	return acq
}

// Order adds an order step to the query.
func (acq *AppCoinQuery) Order(o ...OrderFunc) *AppCoinQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first AppCoin entity from the query.
// Returns a *NotFoundError when no AppCoin was found.
func (acq *AppCoinQuery) First(ctx context.Context) (*AppCoin, error) {
	nodes, err := acq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appcoin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AppCoinQuery) FirstX(ctx context.Context) *AppCoin {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppCoin ID from the query.
// Returns a *NotFoundError when no AppCoin ID was found.
func (acq *AppCoinQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appcoin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AppCoinQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppCoin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppCoin entity is found.
// Returns a *NotFoundError when no AppCoin entities are found.
func (acq *AppCoinQuery) Only(ctx context.Context) (*AppCoin, error) {
	nodes, err := acq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appcoin.Label}
	default:
		return nil, &NotSingularError{appcoin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AppCoinQuery) OnlyX(ctx context.Context) *AppCoin {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppCoin ID in the query.
// Returns a *NotSingularError when more than one AppCoin ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AppCoinQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appcoin.Label}
	default:
		err = &NotSingularError{appcoin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AppCoinQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppCoins.
func (acq *AppCoinQuery) All(ctx context.Context) ([]*AppCoin, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acq *AppCoinQuery) AllX(ctx context.Context) []*AppCoin {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppCoin IDs.
func (acq *AppCoinQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := acq.Select(appcoin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AppCoinQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AppCoinQuery) Count(ctx context.Context) (int, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AppCoinQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AppCoinQuery) Exist(ctx context.Context) (bool, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AppCoinQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppCoinQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AppCoinQuery) Clone() *AppCoinQuery {
	if acq == nil {
		return nil
	}
	return &AppCoinQuery{
		config:     acq.config,
		limit:      acq.limit,
		offset:     acq.offset,
		order:      append([]OrderFunc{}, acq.order...),
		predicates: append([]predicate.AppCoin{}, acq.predicates...),
		// clone intermediate query.
		sql:    acq.sql.Clone(),
		path:   acq.path,
		unique: acq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppCoin.Query().
//		GroupBy(appcoin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acq *AppCoinQuery) GroupBy(field string, fields ...string) *AppCoinGroupBy {
	grbuild := &AppCoinGroupBy{config: acq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acq.sqlQuery(ctx), nil
	}
	grbuild.label = appcoin.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppCoin.Query().
//		Select(appcoin.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (acq *AppCoinQuery) Select(fields ...string) *AppCoinSelect {
	acq.fields = append(acq.fields, fields...)
	selbuild := &AppCoinSelect{AppCoinQuery: acq}
	selbuild.label = appcoin.Label
	selbuild.flds, selbuild.scan = &acq.fields, selbuild.Scan
	return selbuild
}

func (acq *AppCoinQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acq.fields {
		if !appcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	if appcoin.Policy == nil {
		return errors.New("ent: uninitialized appcoin.Policy (forgotten import ent/runtime?)")
	}
	if err := appcoin.Policy.EvalQuery(ctx, acq); err != nil {
		return err
	}
	return nil
}

func (acq *AppCoinQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppCoin, error) {
	var (
		nodes = []*AppCoin{}
		_spec = acq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppCoin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppCoin{config: acq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acq *AppCoinQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.fields
	if len(acq.fields) > 0 {
		_spec.Unique = acq.unique != nil && *acq.unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AppCoinQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acq *AppCoinQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcoin.Table,
			Columns: appcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcoin.FieldID,
			},
		},
		From:   acq.sql,
		Unique: true,
	}
	if unique := acq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcoin.FieldID)
		for i := range fields {
			if fields[i] != appcoin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AppCoinQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(appcoin.Table)
	columns := acq.fields
	if len(columns) == 0 {
		columns = appcoin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.unique != nil && *acq.unique {
		selector.Distinct()
	}
	for _, m := range acq.modifiers {
		m(selector)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (acq *AppCoinQuery) ForUpdate(opts ...sql.LockOption) *AppCoinQuery {
	if acq.driver.Dialect() == dialect.Postgres {
		acq.Unique(false)
	}
	acq.modifiers = append(acq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return acq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (acq *AppCoinQuery) ForShare(opts ...sql.LockOption) *AppCoinQuery {
	if acq.driver.Dialect() == dialect.Postgres {
		acq.Unique(false)
	}
	acq.modifiers = append(acq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return acq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acq *AppCoinQuery) Modify(modifiers ...func(s *sql.Selector)) *AppCoinSelect {
	acq.modifiers = append(acq.modifiers, modifiers...)
	return acq.Select()
}

// AppCoinGroupBy is the group-by builder for AppCoin entities.
type AppCoinGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AppCoinGroupBy) Aggregate(fns ...AggregateFunc) *AppCoinGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acgb *AppCoinGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acgb.path(ctx)
	if err != nil {
		return err
	}
	acgb.sql = query
	return acgb.sqlScan(ctx, v)
}

func (acgb *AppCoinGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acgb.fields {
		if !appcoin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acgb *AppCoinGroupBy) sqlQuery() *sql.Selector {
	selector := acgb.sql.Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acgb.fields)+len(acgb.fns))
		for _, f := range acgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acgb.fields...)...)
}

// AppCoinSelect is the builder for selecting fields of AppCoin entities.
type AppCoinSelect struct {
	*AppCoinQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AppCoinSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	acs.sql = acs.AppCoinQuery.sqlQuery(ctx)
	return acs.sqlScan(ctx, v)
}

func (acs *AppCoinSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acs.sql.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acs *AppCoinSelect) Modify(modifiers ...func(s *sql.Selector)) *AppCoinSelect {
	acs.modifiers = append(acs.modifiers, modifiers...)
	return acs
}