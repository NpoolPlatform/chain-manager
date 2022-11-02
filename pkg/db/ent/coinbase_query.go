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
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinBaseQuery is the builder for querying CoinBase entities.
type CoinBaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinBaseQuery builder.
func (cbq *CoinBaseQuery) Where(ps ...predicate.CoinBase) *CoinBaseQuery {
	cbq.predicates = append(cbq.predicates, ps...)
	return cbq
}

// Limit adds a limit step to the query.
func (cbq *CoinBaseQuery) Limit(limit int) *CoinBaseQuery {
	cbq.limit = &limit
	return cbq
}

// Offset adds an offset step to the query.
func (cbq *CoinBaseQuery) Offset(offset int) *CoinBaseQuery {
	cbq.offset = &offset
	return cbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cbq *CoinBaseQuery) Unique(unique bool) *CoinBaseQuery {
	cbq.unique = &unique
	return cbq
}

// Order adds an order step to the query.
func (cbq *CoinBaseQuery) Order(o ...OrderFunc) *CoinBaseQuery {
	cbq.order = append(cbq.order, o...)
	return cbq
}

// First returns the first CoinBase entity from the query.
// Returns a *NotFoundError when no CoinBase was found.
func (cbq *CoinBaseQuery) First(ctx context.Context) (*CoinBase, error) {
	nodes, err := cbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coinbase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cbq *CoinBaseQuery) FirstX(ctx context.Context) *CoinBase {
	node, err := cbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinBase ID from the query.
// Returns a *NotFoundError when no CoinBase ID was found.
func (cbq *CoinBaseQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coinbase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cbq *CoinBaseQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoinBase entity is found.
// Returns a *NotFoundError when no CoinBase entities are found.
func (cbq *CoinBaseQuery) Only(ctx context.Context) (*CoinBase, error) {
	nodes, err := cbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coinbase.Label}
	default:
		return nil, &NotSingularError{coinbase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cbq *CoinBaseQuery) OnlyX(ctx context.Context) *CoinBase {
	node, err := cbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinBase ID in the query.
// Returns a *NotSingularError when more than one CoinBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (cbq *CoinBaseQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coinbase.Label}
	default:
		err = &NotSingularError{coinbase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cbq *CoinBaseQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinBases.
func (cbq *CoinBaseQuery) All(ctx context.Context) ([]*CoinBase, error) {
	if err := cbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cbq *CoinBaseQuery) AllX(ctx context.Context) []*CoinBase {
	nodes, err := cbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinBase IDs.
func (cbq *CoinBaseQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cbq.Select(coinbase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cbq *CoinBaseQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cbq *CoinBaseQuery) Count(ctx context.Context) (int, error) {
	if err := cbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cbq *CoinBaseQuery) CountX(ctx context.Context) int {
	count, err := cbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cbq *CoinBaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := cbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cbq *CoinBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := cbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cbq *CoinBaseQuery) Clone() *CoinBaseQuery {
	if cbq == nil {
		return nil
	}
	return &CoinBaseQuery{
		config:     cbq.config,
		limit:      cbq.limit,
		offset:     cbq.offset,
		order:      append([]OrderFunc{}, cbq.order...),
		predicates: append([]predicate.CoinBase{}, cbq.predicates...),
		// clone intermediate query.
		sql:    cbq.sql.Clone(),
		path:   cbq.path,
		unique: cbq.unique,
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
//	client.CoinBase.Query().
//		GroupBy(coinbase.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cbq *CoinBaseQuery) GroupBy(field string, fields ...string) *CoinBaseGroupBy {
	grbuild := &CoinBaseGroupBy{config: cbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cbq.sqlQuery(ctx), nil
	}
	grbuild.label = coinbase.Label
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
//	client.CoinBase.Query().
//		Select(coinbase.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (cbq *CoinBaseQuery) Select(fields ...string) *CoinBaseSelect {
	cbq.fields = append(cbq.fields, fields...)
	selbuild := &CoinBaseSelect{CoinBaseQuery: cbq}
	selbuild.label = coinbase.Label
	selbuild.flds, selbuild.scan = &cbq.fields, selbuild.Scan
	return selbuild
}

func (cbq *CoinBaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cbq.fields {
		if !coinbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cbq.path != nil {
		prev, err := cbq.path(ctx)
		if err != nil {
			return err
		}
		cbq.sql = prev
	}
	if coinbase.Policy == nil {
		return errors.New("ent: uninitialized coinbase.Policy (forgotten import ent/runtime?)")
	}
	if err := coinbase.Policy.EvalQuery(ctx, cbq); err != nil {
		return err
	}
	return nil
}

func (cbq *CoinBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CoinBase, error) {
	var (
		nodes = []*CoinBase{}
		_spec = cbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CoinBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CoinBase{config: cbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cbq *CoinBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cbq.querySpec()
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	_spec.Node.Columns = cbq.fields
	if len(cbq.fields) > 0 {
		_spec.Unique = cbq.unique != nil && *cbq.unique
	}
	return sqlgraph.CountNodes(ctx, cbq.driver, _spec)
}

func (cbq *CoinBaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cbq *CoinBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinbase.Table,
			Columns: coinbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinbase.FieldID,
			},
		},
		From:   cbq.sql,
		Unique: true,
	}
	if unique := cbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinbase.FieldID)
		for i := range fields {
			if fields[i] != coinbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cbq *CoinBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cbq.driver.Dialect())
	t1 := builder.Table(coinbase.Table)
	columns := cbq.fields
	if len(columns) == 0 {
		columns = coinbase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cbq.sql != nil {
		selector = cbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cbq.unique != nil && *cbq.unique {
		selector.Distinct()
	}
	for _, m := range cbq.modifiers {
		m(selector)
	}
	for _, p := range cbq.predicates {
		p(selector)
	}
	for _, p := range cbq.order {
		p(selector)
	}
	if offset := cbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cbq *CoinBaseQuery) ForUpdate(opts ...sql.LockOption) *CoinBaseQuery {
	if cbq.driver.Dialect() == dialect.Postgres {
		cbq.Unique(false)
	}
	cbq.modifiers = append(cbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cbq *CoinBaseQuery) ForShare(opts ...sql.LockOption) *CoinBaseQuery {
	if cbq.driver.Dialect() == dialect.Postgres {
		cbq.Unique(false)
	}
	cbq.modifiers = append(cbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cbq *CoinBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *CoinBaseSelect {
	cbq.modifiers = append(cbq.modifiers, modifiers...)
	return cbq.Select()
}

// CoinBaseGroupBy is the group-by builder for CoinBase entities.
type CoinBaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cbgb *CoinBaseGroupBy) Aggregate(fns ...AggregateFunc) *CoinBaseGroupBy {
	cbgb.fns = append(cbgb.fns, fns...)
	return cbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cbgb *CoinBaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cbgb.path(ctx)
	if err != nil {
		return err
	}
	cbgb.sql = query
	return cbgb.sqlScan(ctx, v)
}

func (cbgb *CoinBaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cbgb.fields {
		if !coinbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cbgb *CoinBaseGroupBy) sqlQuery() *sql.Selector {
	selector := cbgb.sql.Select()
	aggregation := make([]string, 0, len(cbgb.fns))
	for _, fn := range cbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cbgb.fields)+len(cbgb.fns))
		for _, f := range cbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cbgb.fields...)...)
}

// CoinBaseSelect is the builder for selecting fields of CoinBase entities.
type CoinBaseSelect struct {
	*CoinBaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cbs *CoinBaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cbs.prepareQuery(ctx); err != nil {
		return err
	}
	cbs.sql = cbs.CoinBaseQuery.sqlQuery(ctx)
	return cbs.sqlScan(ctx, v)
}

func (cbs *CoinBaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cbs.sql.Query()
	if err := cbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cbs *CoinBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *CoinBaseSelect {
	cbs.modifiers = append(cbs.modifiers, modifiers...)
	return cbs
}