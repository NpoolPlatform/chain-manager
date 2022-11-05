// Code generated by ent, DO NOT EDIT.

package coinextra

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// HomePage applies equality check predicate on the "home_page" field. It's identical to HomePageEQ.
func HomePage(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHomePage), v))
	})
}

// Specs applies equality check predicate on the "specs" field. It's identical to SpecsEQ.
func Specs(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpecs), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// HomePageEQ applies the EQ predicate on the "home_page" field.
func HomePageEQ(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHomePage), v))
	})
}

// HomePageNEQ applies the NEQ predicate on the "home_page" field.
func HomePageNEQ(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHomePage), v))
	})
}

// HomePageIn applies the In predicate on the "home_page" field.
func HomePageIn(vs ...string) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHomePage), v...))
	})
}

// HomePageNotIn applies the NotIn predicate on the "home_page" field.
func HomePageNotIn(vs ...string) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHomePage), v...))
	})
}

// HomePageGT applies the GT predicate on the "home_page" field.
func HomePageGT(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHomePage), v))
	})
}

// HomePageGTE applies the GTE predicate on the "home_page" field.
func HomePageGTE(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHomePage), v))
	})
}

// HomePageLT applies the LT predicate on the "home_page" field.
func HomePageLT(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHomePage), v))
	})
}

// HomePageLTE applies the LTE predicate on the "home_page" field.
func HomePageLTE(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHomePage), v))
	})
}

// HomePageContains applies the Contains predicate on the "home_page" field.
func HomePageContains(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHomePage), v))
	})
}

// HomePageHasPrefix applies the HasPrefix predicate on the "home_page" field.
func HomePageHasPrefix(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHomePage), v))
	})
}

// HomePageHasSuffix applies the HasSuffix predicate on the "home_page" field.
func HomePageHasSuffix(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHomePage), v))
	})
}

// HomePageIsNil applies the IsNil predicate on the "home_page" field.
func HomePageIsNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHomePage)))
	})
}

// HomePageNotNil applies the NotNil predicate on the "home_page" field.
func HomePageNotNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHomePage)))
	})
}

// HomePageEqualFold applies the EqualFold predicate on the "home_page" field.
func HomePageEqualFold(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHomePage), v))
	})
}

// HomePageContainsFold applies the ContainsFold predicate on the "home_page" field.
func HomePageContainsFold(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHomePage), v))
	})
}

// SpecsEQ applies the EQ predicate on the "specs" field.
func SpecsEQ(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpecs), v))
	})
}

// SpecsNEQ applies the NEQ predicate on the "specs" field.
func SpecsNEQ(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpecs), v))
	})
}

// SpecsIn applies the In predicate on the "specs" field.
func SpecsIn(vs ...string) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSpecs), v...))
	})
}

// SpecsNotIn applies the NotIn predicate on the "specs" field.
func SpecsNotIn(vs ...string) predicate.CoinExtra {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSpecs), v...))
	})
}

// SpecsGT applies the GT predicate on the "specs" field.
func SpecsGT(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpecs), v))
	})
}

// SpecsGTE applies the GTE predicate on the "specs" field.
func SpecsGTE(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpecs), v))
	})
}

// SpecsLT applies the LT predicate on the "specs" field.
func SpecsLT(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpecs), v))
	})
}

// SpecsLTE applies the LTE predicate on the "specs" field.
func SpecsLTE(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpecs), v))
	})
}

// SpecsContains applies the Contains predicate on the "specs" field.
func SpecsContains(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpecs), v))
	})
}

// SpecsHasPrefix applies the HasPrefix predicate on the "specs" field.
func SpecsHasPrefix(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpecs), v))
	})
}

// SpecsHasSuffix applies the HasSuffix predicate on the "specs" field.
func SpecsHasSuffix(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpecs), v))
	})
}

// SpecsIsNil applies the IsNil predicate on the "specs" field.
func SpecsIsNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpecs)))
	})
}

// SpecsNotNil applies the NotNil predicate on the "specs" field.
func SpecsNotNil() predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpecs)))
	})
}

// SpecsEqualFold applies the EqualFold predicate on the "specs" field.
func SpecsEqualFold(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpecs), v))
	})
}

// SpecsContainsFold applies the ContainsFold predicate on the "specs" field.
func SpecsContainsFold(v string) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpecs), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CoinExtra) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CoinExtra) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
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
func Not(p predicate.CoinExtra) predicate.CoinExtra {
	return predicate.CoinExtra(func(s *sql.Selector) {
		p(s.Not())
	})
}
