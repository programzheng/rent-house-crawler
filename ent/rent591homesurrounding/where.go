// Code generated by ent, DO NOT EDIT.

package rent591homesurrounding

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldType, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldDesc, v))
}

// Distance applies equality check predicate on the "distance" field. It's identical to DistanceEQ.
func Distance(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldDistance, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldCreatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContainsFold(FieldType, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContainsFold(FieldDesc, v))
}

// DistanceEQ applies the EQ predicate on the "distance" field.
func DistanceEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldDistance, v))
}

// DistanceNEQ applies the NEQ predicate on the "distance" field.
func DistanceNEQ(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNEQ(FieldDistance, v))
}

// DistanceIn applies the In predicate on the "distance" field.
func DistanceIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldIn(FieldDistance, vs...))
}

// DistanceNotIn applies the NotIn predicate on the "distance" field.
func DistanceNotIn(vs ...string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNotIn(FieldDistance, vs...))
}

// DistanceGT applies the GT predicate on the "distance" field.
func DistanceGT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGT(FieldDistance, v))
}

// DistanceGTE applies the GTE predicate on the "distance" field.
func DistanceGTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGTE(FieldDistance, v))
}

// DistanceLT applies the LT predicate on the "distance" field.
func DistanceLT(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLT(FieldDistance, v))
}

// DistanceLTE applies the LTE predicate on the "distance" field.
func DistanceLTE(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLTE(FieldDistance, v))
}

// DistanceContains applies the Contains predicate on the "distance" field.
func DistanceContains(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContains(FieldDistance, v))
}

// DistanceHasPrefix applies the HasPrefix predicate on the "distance" field.
func DistanceHasPrefix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasPrefix(FieldDistance, v))
}

// DistanceHasSuffix applies the HasSuffix predicate on the "distance" field.
func DistanceHasSuffix(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldHasSuffix(FieldDistance, v))
}

// DistanceEqualFold applies the EqualFold predicate on the "distance" field.
func DistanceEqualFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEqualFold(FieldDistance, v))
}

// DistanceContainsFold applies the ContainsFold predicate on the "distance" field.
func DistanceContainsFold(v string) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldContainsFold(FieldDistance, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(sql.FieldLTE(FieldCreatedAt, v))
}

// HasRent591homes applies the HasEdge predicate on the "rent591homes" edge.
func HasRent591homes() predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, Rent591homesTable, Rent591homesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRent591homesWith applies the HasEdge predicate on the "rent591homes" edge with a given conditions (other predicates).
func HasRent591homesWith(preds ...predicate.Rent591Home) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(func(s *sql.Selector) {
		step := newRent591homesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rent591HomeSurrounding) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rent591HomeSurrounding) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(func(s *sql.Selector) {
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
func Not(p predicate.Rent591HomeSurrounding) predicate.Rent591HomeSurrounding {
	return predicate.Rent591HomeSurrounding(func(s *sql.Selector) {
		p(s.Not())
	})
}
