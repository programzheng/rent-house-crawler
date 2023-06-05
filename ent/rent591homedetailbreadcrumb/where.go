// Code generated by ent, DO NOT EDIT.

package rent591homedetailbreadcrumb

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldName, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldPostID, v))
}

// Query applies equality check predicate on the "query" field. It's identical to QueryEQ.
func Query(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldQuery, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldLink, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContainsFold(FieldName, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v int) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLTE(FieldPostID, v))
}

// QueryEQ applies the EQ predicate on the "query" field.
func QueryEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldQuery, v))
}

// QueryNEQ applies the NEQ predicate on the "query" field.
func QueryNEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNEQ(FieldQuery, v))
}

// QueryIn applies the In predicate on the "query" field.
func QueryIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldIn(FieldQuery, vs...))
}

// QueryNotIn applies the NotIn predicate on the "query" field.
func QueryNotIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNotIn(FieldQuery, vs...))
}

// QueryGT applies the GT predicate on the "query" field.
func QueryGT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGT(FieldQuery, v))
}

// QueryGTE applies the GTE predicate on the "query" field.
func QueryGTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGTE(FieldQuery, v))
}

// QueryLT applies the LT predicate on the "query" field.
func QueryLT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLT(FieldQuery, v))
}

// QueryLTE applies the LTE predicate on the "query" field.
func QueryLTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLTE(FieldQuery, v))
}

// QueryContains applies the Contains predicate on the "query" field.
func QueryContains(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContains(FieldQuery, v))
}

// QueryHasPrefix applies the HasPrefix predicate on the "query" field.
func QueryHasPrefix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasPrefix(FieldQuery, v))
}

// QueryHasSuffix applies the HasSuffix predicate on the "query" field.
func QueryHasSuffix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasSuffix(FieldQuery, v))
}

// QueryEqualFold applies the EqualFold predicate on the "query" field.
func QueryEqualFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEqualFold(FieldQuery, v))
}

// QueryContainsFold applies the ContainsFold predicate on the "query" field.
func QueryContainsFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContainsFold(FieldQuery, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(sql.FieldContainsFold(FieldLink, v))
}

// HasRent591homeDetails applies the HasEdge predicate on the "rent591home_details" edge.
func HasRent591homeDetails() predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, Rent591homeDetailsTable, Rent591homeDetailsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRent591homeDetailsWith applies the HasEdge predicate on the "rent591home_details" edge with a given conditions (other predicates).
func HasRent591homeDetailsWith(preds ...predicate.Rent591HomeDetail) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(func(s *sql.Selector) {
		step := newRent591homeDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rent591HomeDetailBreadcrumb) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rent591HomeDetailBreadcrumb) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Rent591HomeDetailBreadcrumb) predicate.Rent591HomeDetailBreadcrumb {
	return predicate.Rent591HomeDetailBreadcrumb(func(s *sql.Selector) {
		p(s.Not())
	})
}
