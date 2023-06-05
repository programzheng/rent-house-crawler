// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetail"
	"github.com/programzheng/rent-house-crawler/ent/rent591homedetailtag"
)

// Rent591HomeDetailTagQuery is the builder for querying Rent591HomeDetailTag entities.
type Rent591HomeDetailTagQuery struct {
	config
	ctx                    *QueryContext
	order                  []rent591homedetailtag.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Rent591HomeDetailTag
	withRent591homeDetails *Rent591HomeDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the Rent591HomeDetailTagQuery builder.
func (rdtq *Rent591HomeDetailTagQuery) Where(ps ...predicate.Rent591HomeDetailTag) *Rent591HomeDetailTagQuery {
	rdtq.predicates = append(rdtq.predicates, ps...)
	return rdtq
}

// Limit the number of records to be returned by this query.
func (rdtq *Rent591HomeDetailTagQuery) Limit(limit int) *Rent591HomeDetailTagQuery {
	rdtq.ctx.Limit = &limit
	return rdtq
}

// Offset to start from.
func (rdtq *Rent591HomeDetailTagQuery) Offset(offset int) *Rent591HomeDetailTagQuery {
	rdtq.ctx.Offset = &offset
	return rdtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rdtq *Rent591HomeDetailTagQuery) Unique(unique bool) *Rent591HomeDetailTagQuery {
	rdtq.ctx.Unique = &unique
	return rdtq
}

// Order specifies how the records should be ordered.
func (rdtq *Rent591HomeDetailTagQuery) Order(o ...rent591homedetailtag.OrderOption) *Rent591HomeDetailTagQuery {
	rdtq.order = append(rdtq.order, o...)
	return rdtq
}

// QueryRent591homeDetails chains the current query on the "rent591home_details" edge.
func (rdtq *Rent591HomeDetailTagQuery) QueryRent591homeDetails() *Rent591HomeDetailQuery {
	query := (&Rent591HomeDetailClient{config: rdtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rdtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rdtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rent591homedetailtag.Table, rent591homedetailtag.FieldID, selector),
			sqlgraph.To(rent591homedetail.Table, rent591homedetail.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, rent591homedetailtag.Rent591homeDetailsTable, rent591homedetailtag.Rent591homeDetailsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rdtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Rent591HomeDetailTag entity from the query.
// Returns a *NotFoundError when no Rent591HomeDetailTag was found.
func (rdtq *Rent591HomeDetailTagQuery) First(ctx context.Context) (*Rent591HomeDetailTag, error) {
	nodes, err := rdtq.Limit(1).All(setContextOp(ctx, rdtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rent591homedetailtag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) FirstX(ctx context.Context) *Rent591HomeDetailTag {
	node, err := rdtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Rent591HomeDetailTag ID from the query.
// Returns a *NotFoundError when no Rent591HomeDetailTag ID was found.
func (rdtq *Rent591HomeDetailTagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rdtq.Limit(1).IDs(setContextOp(ctx, rdtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rent591homedetailtag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) FirstIDX(ctx context.Context) int {
	id, err := rdtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Rent591HomeDetailTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Rent591HomeDetailTag entity is found.
// Returns a *NotFoundError when no Rent591HomeDetailTag entities are found.
func (rdtq *Rent591HomeDetailTagQuery) Only(ctx context.Context) (*Rent591HomeDetailTag, error) {
	nodes, err := rdtq.Limit(2).All(setContextOp(ctx, rdtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rent591homedetailtag.Label}
	default:
		return nil, &NotSingularError{rent591homedetailtag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) OnlyX(ctx context.Context) *Rent591HomeDetailTag {
	node, err := rdtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Rent591HomeDetailTag ID in the query.
// Returns a *NotSingularError when more than one Rent591HomeDetailTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (rdtq *Rent591HomeDetailTagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rdtq.Limit(2).IDs(setContextOp(ctx, rdtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rent591homedetailtag.Label}
	default:
		err = &NotSingularError{rent591homedetailtag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) OnlyIDX(ctx context.Context) int {
	id, err := rdtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rent591HomeDetailTags.
func (rdtq *Rent591HomeDetailTagQuery) All(ctx context.Context) ([]*Rent591HomeDetailTag, error) {
	ctx = setContextOp(ctx, rdtq.ctx, "All")
	if err := rdtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Rent591HomeDetailTag, *Rent591HomeDetailTagQuery]()
	return withInterceptors[[]*Rent591HomeDetailTag](ctx, rdtq, qr, rdtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) AllX(ctx context.Context) []*Rent591HomeDetailTag {
	nodes, err := rdtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Rent591HomeDetailTag IDs.
func (rdtq *Rent591HomeDetailTagQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rdtq.ctx.Unique == nil && rdtq.path != nil {
		rdtq.Unique(true)
	}
	ctx = setContextOp(ctx, rdtq.ctx, "IDs")
	if err = rdtq.Select(rent591homedetailtag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) IDsX(ctx context.Context) []int {
	ids, err := rdtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rdtq *Rent591HomeDetailTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rdtq.ctx, "Count")
	if err := rdtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rdtq, querierCount[*Rent591HomeDetailTagQuery](), rdtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) CountX(ctx context.Context) int {
	count, err := rdtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rdtq *Rent591HomeDetailTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rdtq.ctx, "Exist")
	switch _, err := rdtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rdtq *Rent591HomeDetailTagQuery) ExistX(ctx context.Context) bool {
	exist, err := rdtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the Rent591HomeDetailTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rdtq *Rent591HomeDetailTagQuery) Clone() *Rent591HomeDetailTagQuery {
	if rdtq == nil {
		return nil
	}
	return &Rent591HomeDetailTagQuery{
		config:                 rdtq.config,
		ctx:                    rdtq.ctx.Clone(),
		order:                  append([]rent591homedetailtag.OrderOption{}, rdtq.order...),
		inters:                 append([]Interceptor{}, rdtq.inters...),
		predicates:             append([]predicate.Rent591HomeDetailTag{}, rdtq.predicates...),
		withRent591homeDetails: rdtq.withRent591homeDetails.Clone(),
		// clone intermediate query.
		sql:  rdtq.sql.Clone(),
		path: rdtq.path,
	}
}

// WithRent591homeDetails tells the query-builder to eager-load the nodes that are connected to
// the "rent591home_details" edge. The optional arguments are used to configure the query builder of the edge.
func (rdtq *Rent591HomeDetailTagQuery) WithRent591homeDetails(opts ...func(*Rent591HomeDetailQuery)) *Rent591HomeDetailTagQuery {
	query := (&Rent591HomeDetailClient{config: rdtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rdtq.withRent591homeDetails = query
	return rdtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PostID int `json:"post_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Rent591HomeDetailTag.Query().
//		GroupBy(rent591homedetailtag.FieldPostID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rdtq *Rent591HomeDetailTagQuery) GroupBy(field string, fields ...string) *Rent591HomeDetailTagGroupBy {
	rdtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &Rent591HomeDetailTagGroupBy{build: rdtq}
	grbuild.flds = &rdtq.ctx.Fields
	grbuild.label = rent591homedetailtag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PostID int `json:"post_id,omitempty"`
//	}
//
//	client.Rent591HomeDetailTag.Query().
//		Select(rent591homedetailtag.FieldPostID).
//		Scan(ctx, &v)
func (rdtq *Rent591HomeDetailTagQuery) Select(fields ...string) *Rent591HomeDetailTagSelect {
	rdtq.ctx.Fields = append(rdtq.ctx.Fields, fields...)
	sbuild := &Rent591HomeDetailTagSelect{Rent591HomeDetailTagQuery: rdtq}
	sbuild.label = rent591homedetailtag.Label
	sbuild.flds, sbuild.scan = &rdtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a Rent591HomeDetailTagSelect configured with the given aggregations.
func (rdtq *Rent591HomeDetailTagQuery) Aggregate(fns ...AggregateFunc) *Rent591HomeDetailTagSelect {
	return rdtq.Select().Aggregate(fns...)
}

func (rdtq *Rent591HomeDetailTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rdtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rdtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rdtq.ctx.Fields {
		if !rent591homedetailtag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rdtq.path != nil {
		prev, err := rdtq.path(ctx)
		if err != nil {
			return err
		}
		rdtq.sql = prev
	}
	return nil
}

func (rdtq *Rent591HomeDetailTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Rent591HomeDetailTag, error) {
	var (
		nodes       = []*Rent591HomeDetailTag{}
		_spec       = rdtq.querySpec()
		loadedTypes = [1]bool{
			rdtq.withRent591homeDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Rent591HomeDetailTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Rent591HomeDetailTag{config: rdtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rdtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rdtq.withRent591homeDetails; query != nil {
		if err := rdtq.loadRent591homeDetails(ctx, query, nodes,
			func(n *Rent591HomeDetailTag) { n.Edges.Rent591homeDetails = []*Rent591HomeDetail{} },
			func(n *Rent591HomeDetailTag, e *Rent591HomeDetail) {
				n.Edges.Rent591homeDetails = append(n.Edges.Rent591homeDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rdtq *Rent591HomeDetailTagQuery) loadRent591homeDetails(ctx context.Context, query *Rent591HomeDetailQuery, nodes []*Rent591HomeDetailTag, init func(*Rent591HomeDetailTag), assign func(*Rent591HomeDetailTag, *Rent591HomeDetail)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Rent591HomeDetailTag)
	nids := make(map[int]map[*Rent591HomeDetailTag]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(rent591homedetailtag.Rent591homeDetailsTable)
		s.Join(joinT).On(s.C(rent591homedetail.FieldID), joinT.C(rent591homedetailtag.Rent591homeDetailsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(rent591homedetailtag.Rent591homeDetailsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(rent591homedetailtag.Rent591homeDetailsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Rent591HomeDetailTag]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Rent591HomeDetail](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "rent591home_details" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rdtq *Rent591HomeDetailTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rdtq.querySpec()
	_spec.Node.Columns = rdtq.ctx.Fields
	if len(rdtq.ctx.Fields) > 0 {
		_spec.Unique = rdtq.ctx.Unique != nil && *rdtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rdtq.driver, _spec)
}

func (rdtq *Rent591HomeDetailTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rent591homedetailtag.Table, rent591homedetailtag.Columns, sqlgraph.NewFieldSpec(rent591homedetailtag.FieldID, field.TypeInt))
	_spec.From = rdtq.sql
	if unique := rdtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rdtq.path != nil {
		_spec.Unique = true
	}
	if fields := rdtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rent591homedetailtag.FieldID)
		for i := range fields {
			if fields[i] != rent591homedetailtag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rdtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rdtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rdtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rdtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rdtq *Rent591HomeDetailTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rdtq.driver.Dialect())
	t1 := builder.Table(rent591homedetailtag.Table)
	columns := rdtq.ctx.Fields
	if len(columns) == 0 {
		columns = rent591homedetailtag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rdtq.sql != nil {
		selector = rdtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rdtq.ctx.Unique != nil && *rdtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rdtq.predicates {
		p(selector)
	}
	for _, p := range rdtq.order {
		p(selector)
	}
	if offset := rdtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rdtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Rent591HomeDetailTagGroupBy is the group-by builder for Rent591HomeDetailTag entities.
type Rent591HomeDetailTagGroupBy struct {
	selector
	build *Rent591HomeDetailTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rdtgb *Rent591HomeDetailTagGroupBy) Aggregate(fns ...AggregateFunc) *Rent591HomeDetailTagGroupBy {
	rdtgb.fns = append(rdtgb.fns, fns...)
	return rdtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rdtgb *Rent591HomeDetailTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rdtgb.build.ctx, "GroupBy")
	if err := rdtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Rent591HomeDetailTagQuery, *Rent591HomeDetailTagGroupBy](ctx, rdtgb.build, rdtgb, rdtgb.build.inters, v)
}

func (rdtgb *Rent591HomeDetailTagGroupBy) sqlScan(ctx context.Context, root *Rent591HomeDetailTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rdtgb.fns))
	for _, fn := range rdtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rdtgb.flds)+len(rdtgb.fns))
		for _, f := range *rdtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rdtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rdtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Rent591HomeDetailTagSelect is the builder for selecting fields of Rent591HomeDetailTag entities.
type Rent591HomeDetailTagSelect struct {
	*Rent591HomeDetailTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rdts *Rent591HomeDetailTagSelect) Aggregate(fns ...AggregateFunc) *Rent591HomeDetailTagSelect {
	rdts.fns = append(rdts.fns, fns...)
	return rdts
}

// Scan applies the selector query and scans the result into the given value.
func (rdts *Rent591HomeDetailTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rdts.ctx, "Select")
	if err := rdts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Rent591HomeDetailTagQuery, *Rent591HomeDetailTagSelect](ctx, rdts.Rent591HomeDetailTagQuery, rdts, rdts.inters, v)
}

func (rdts *Rent591HomeDetailTagSelect) sqlScan(ctx context.Context, root *Rent591HomeDetailTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rdts.fns))
	for _, fn := range rdts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rdts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rdts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
