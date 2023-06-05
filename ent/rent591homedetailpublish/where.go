// Code generated by ent, DO NOT EDIT.

package rent591homedetailpublish

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/programzheng/rent-house-crawler/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldID, id))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldPostID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldName, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldKey, v))
}

// PostTime applies equality check predicate on the "post_time" field. It's identical to PostTimeEQ.
func PostTime(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldPostTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldUpdateTime, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v int) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldPostID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContainsFold(FieldName, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContainsFold(FieldKey, v))
}

// PostTimeEQ applies the EQ predicate on the "post_time" field.
func PostTimeEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldPostTime, v))
}

// PostTimeNEQ applies the NEQ predicate on the "post_time" field.
func PostTimeNEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldPostTime, v))
}

// PostTimeIn applies the In predicate on the "post_time" field.
func PostTimeIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldPostTime, vs...))
}

// PostTimeNotIn applies the NotIn predicate on the "post_time" field.
func PostTimeNotIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldPostTime, vs...))
}

// PostTimeGT applies the GT predicate on the "post_time" field.
func PostTimeGT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldPostTime, v))
}

// PostTimeGTE applies the GTE predicate on the "post_time" field.
func PostTimeGTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldPostTime, v))
}

// PostTimeLT applies the LT predicate on the "post_time" field.
func PostTimeLT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldPostTime, v))
}

// PostTimeLTE applies the LTE predicate on the "post_time" field.
func PostTimeLTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldPostTime, v))
}

// PostTimeContains applies the Contains predicate on the "post_time" field.
func PostTimeContains(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContains(FieldPostTime, v))
}

// PostTimeHasPrefix applies the HasPrefix predicate on the "post_time" field.
func PostTimeHasPrefix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasPrefix(FieldPostTime, v))
}

// PostTimeHasSuffix applies the HasSuffix predicate on the "post_time" field.
func PostTimeHasSuffix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasSuffix(FieldPostTime, v))
}

// PostTimeEqualFold applies the EqualFold predicate on the "post_time" field.
func PostTimeEqualFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEqualFold(FieldPostTime, v))
}

// PostTimeContainsFold applies the ContainsFold predicate on the "post_time" field.
func PostTimeContainsFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContainsFold(FieldPostTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeContains applies the Contains predicate on the "update_time" field.
func UpdateTimeContains(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContains(FieldUpdateTime, v))
}

// UpdateTimeHasPrefix applies the HasPrefix predicate on the "update_time" field.
func UpdateTimeHasPrefix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasPrefix(FieldUpdateTime, v))
}

// UpdateTimeHasSuffix applies the HasSuffix predicate on the "update_time" field.
func UpdateTimeHasSuffix(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldHasSuffix(FieldUpdateTime, v))
}

// UpdateTimeEqualFold applies the EqualFold predicate on the "update_time" field.
func UpdateTimeEqualFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldEqualFold(FieldUpdateTime, v))
}

// UpdateTimeContainsFold applies the ContainsFold predicate on the "update_time" field.
func UpdateTimeContainsFold(v string) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(sql.FieldContainsFold(FieldUpdateTime, v))
}

// HasRent591homeDetails applies the HasEdge predicate on the "rent591home_details" edge.
func HasRent591homeDetails() predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, Rent591homeDetailsTable, Rent591homeDetailsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRent591homeDetailsWith applies the HasEdge predicate on the "rent591home_details" edge with a given conditions (other predicates).
func HasRent591homeDetailsWith(preds ...predicate.Rent591HomeDetail) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(func(s *sql.Selector) {
		step := newRent591homeDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rent591HomeDetailPublish) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rent591HomeDetailPublish) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(func(s *sql.Selector) {
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
func Not(p predicate.Rent591HomeDetailPublish) predicate.Rent591HomeDetailPublish {
	return predicate.Rent591HomeDetailPublish(func(s *sql.Selector) {
		p(s.Not())
	})
}
