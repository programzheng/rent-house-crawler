// Code generated by ent, DO NOT EDIT.

package rent591homedetailbrowse

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLTE(FieldID, id))
}

// Pc applies equality check predicate on the "pc" field. It's identical to PcEQ.
func Pc(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldPc, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldMobile, v))
}

// PcEQ applies the EQ predicate on the "pc" field.
func PcEQ(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldPc, v))
}

// PcNEQ applies the NEQ predicate on the "pc" field.
func PcNEQ(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNEQ(FieldPc, v))
}

// PcIn applies the In predicate on the "pc" field.
func PcIn(vs ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldIn(FieldPc, vs...))
}

// PcNotIn applies the NotIn predicate on the "pc" field.
func PcNotIn(vs ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNotIn(FieldPc, vs...))
}

// PcGT applies the GT predicate on the "pc" field.
func PcGT(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGT(FieldPc, v))
}

// PcGTE applies the GTE predicate on the "pc" field.
func PcGTE(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGTE(FieldPc, v))
}

// PcLT applies the LT predicate on the "pc" field.
func PcLT(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLT(FieldPc, v))
}

// PcLTE applies the LTE predicate on the "pc" field.
func PcLTE(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLTE(FieldPc, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v int) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(sql.FieldLTE(FieldMobile, v))
}

// HasRent591homeDetails applies the HasEdge predicate on the "rent591home_details" edge.
func HasRent591homeDetails() predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, Rent591homeDetailsTable, Rent591homeDetailsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRent591homeDetailsWith applies the HasEdge predicate on the "rent591home_details" edge with a given conditions (other predicates).
func HasRent591homeDetailsWith(preds ...predicate.Rent591HomeDetail) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(func(s *sql.Selector) {
		step := newRent591homeDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rent591HomeDetailBrowse) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rent591HomeDetailBrowse) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(func(s *sql.Selector) {
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
func Not(p predicate.Rent591HomeDetailBrowse) predicate.Rent591HomeDetailBrowse {
	return predicate.Rent591HomeDetailBrowse(func(s *sql.Selector) {
		p(s.Not())
	})
}
