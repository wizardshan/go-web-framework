// Code generated by ent, DO NOT EDIT.

package user

import (
	"app/repository/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// HashID applies equality check predicate on the "hash_id" field. It's identical to HashIDEQ.
func HashID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHashID, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// Bio applies equality check predicate on the "bio" field. It's identical to BioEQ.
func Bio(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBio, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdateTime, v))
}

// HashIDEQ applies the EQ predicate on the "hash_id" field.
func HashIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHashID, v))
}

// HashIDNEQ applies the NEQ predicate on the "hash_id" field.
func HashIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHashID, v))
}

// HashIDIn applies the In predicate on the "hash_id" field.
func HashIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHashID, vs...))
}

// HashIDNotIn applies the NotIn predicate on the "hash_id" field.
func HashIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHashID, vs...))
}

// HashIDGT applies the GT predicate on the "hash_id" field.
func HashIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHashID, v))
}

// HashIDGTE applies the GTE predicate on the "hash_id" field.
func HashIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHashID, v))
}

// HashIDLT applies the LT predicate on the "hash_id" field.
func HashIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHashID, v))
}

// HashIDLTE applies the LTE predicate on the "hash_id" field.
func HashIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHashID, v))
}

// HashIDContains applies the Contains predicate on the "hash_id" field.
func HashIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHashID, v))
}

// HashIDHasPrefix applies the HasPrefix predicate on the "hash_id" field.
func HashIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHashID, v))
}

// HashIDHasSuffix applies the HasSuffix predicate on the "hash_id" field.
func HashIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHashID, v))
}

// HashIDEqualFold applies the EqualFold predicate on the "hash_id" field.
func HashIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHashID, v))
}

// HashIDContainsFold applies the ContainsFold predicate on the "hash_id" field.
func HashIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHashID, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickname, v))
}

// BioEQ applies the EQ predicate on the "bio" field.
func BioEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBio, v))
}

// BioNEQ applies the NEQ predicate on the "bio" field.
func BioNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBio, v))
}

// BioIn applies the In predicate on the "bio" field.
func BioIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBio, vs...))
}

// BioNotIn applies the NotIn predicate on the "bio" field.
func BioNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBio, vs...))
}

// BioGT applies the GT predicate on the "bio" field.
func BioGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBio, v))
}

// BioGTE applies the GTE predicate on the "bio" field.
func BioGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBio, v))
}

// BioLT applies the LT predicate on the "bio" field.
func BioLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBio, v))
}

// BioLTE applies the LTE predicate on the "bio" field.
func BioLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBio, v))
}

// BioContains applies the Contains predicate on the "bio" field.
func BioContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBio, v))
}

// BioHasPrefix applies the HasPrefix predicate on the "bio" field.
func BioHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBio, v))
}

// BioHasSuffix applies the HasSuffix predicate on the "bio" field.
func BioHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBio, v))
}

// BioEqualFold applies the EqualFold predicate on the "bio" field.
func BioEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBio, v))
}

// BioContainsFold applies the ContainsFold predicate on the "bio" field.
func BioContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBio, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
