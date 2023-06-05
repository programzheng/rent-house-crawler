// Code generated by ent, DO NOT EDIT.

package rent591homedetailinfo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldName, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldValue, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldContainsFold(FieldName, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldContainsFold(FieldValue, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v int) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(sql.FieldLTE(FieldKey, v))
}

// HasRent591homeDetails applies the HasEdge predicate on the "rent591home_details" edge.
func HasRent591homeDetails() predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, Rent591homeDetailsTable, Rent591homeDetailsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRent591homeDetailsWith applies the HasEdge predicate on the "rent591home_details" edge with a given conditions (other predicates).
func HasRent591homeDetailsWith(preds ...predicate.Rent591HomeDetail) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(func(s *sql.Selector) {
		step := newRent591homeDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rent591HomeDetailInfo) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rent591HomeDetailInfo) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(func(s *sql.Selector) {
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
func Not(p predicate.Rent591HomeDetailInfo) predicate.Rent591HomeDetailInfo {
	return predicate.Rent591HomeDetailInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}