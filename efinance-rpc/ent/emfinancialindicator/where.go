// Code generated by ent, DO NOT EDIT.

package emfinancialindicator

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldID, id))
}

// FinancialIndicatorsKey applies equality check predicate on the "financial_indicators_key" field. It's identical to FinancialIndicatorsKeyEQ.
func FinancialIndicatorsKey(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsValue applies equality check predicate on the "financial_indicators_value" field. It's identical to FinancialIndicatorsValueEQ.
func FinancialIndicatorsValue(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsName applies equality check predicate on the "financial_indicators_name" field. It's identical to FinancialIndicatorsNameEQ.
func FinancialIndicatorsName(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsName, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldCreateTime, v))
}

// Remarks applies equality check predicate on the "remarks" field. It's identical to RemarksEQ.
func Remarks(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldRemarks, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldStatus, v))
}

// FinancialIndicatorsKeyEQ applies the EQ predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyNEQ applies the NEQ predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyNEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyIn applies the In predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldFinancialIndicatorsKey, vs...))
}

// FinancialIndicatorsKeyNotIn applies the NotIn predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyNotIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldFinancialIndicatorsKey, vs...))
}

// FinancialIndicatorsKeyGT applies the GT predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyGT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyGTE applies the GTE predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyGTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyLT applies the LT predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyLT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyLTE applies the LTE predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyLTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyContains applies the Contains predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyContains(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContains(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyHasPrefix applies the HasPrefix predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyHasPrefix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasPrefix(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyHasSuffix applies the HasSuffix predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyHasSuffix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasSuffix(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyIsNil applies the IsNil predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldFinancialIndicatorsKey))
}

// FinancialIndicatorsKeyNotNil applies the NotNil predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldFinancialIndicatorsKey))
}

// FinancialIndicatorsKeyEqualFold applies the EqualFold predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyEqualFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEqualFold(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsKeyContainsFold applies the ContainsFold predicate on the "financial_indicators_key" field.
func FinancialIndicatorsKeyContainsFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContainsFold(FieldFinancialIndicatorsKey, v))
}

// FinancialIndicatorsValueEQ applies the EQ predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueNEQ applies the NEQ predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueNEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueIn applies the In predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldFinancialIndicatorsValue, vs...))
}

// FinancialIndicatorsValueNotIn applies the NotIn predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueNotIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldFinancialIndicatorsValue, vs...))
}

// FinancialIndicatorsValueGT applies the GT predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueGT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueGTE applies the GTE predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueGTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueLT applies the LT predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueLT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueLTE applies the LTE predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueLTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueContains applies the Contains predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueContains(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContains(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueHasPrefix applies the HasPrefix predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueHasPrefix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasPrefix(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueHasSuffix applies the HasSuffix predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueHasSuffix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasSuffix(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueIsNil applies the IsNil predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldFinancialIndicatorsValue))
}

// FinancialIndicatorsValueNotNil applies the NotNil predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldFinancialIndicatorsValue))
}

// FinancialIndicatorsValueEqualFold applies the EqualFold predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueEqualFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEqualFold(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsValueContainsFold applies the ContainsFold predicate on the "financial_indicators_value" field.
func FinancialIndicatorsValueContainsFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContainsFold(FieldFinancialIndicatorsValue, v))
}

// FinancialIndicatorsNameEQ applies the EQ predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameNEQ applies the NEQ predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameNEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameIn applies the In predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldFinancialIndicatorsName, vs...))
}

// FinancialIndicatorsNameNotIn applies the NotIn predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameNotIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldFinancialIndicatorsName, vs...))
}

// FinancialIndicatorsNameGT applies the GT predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameGT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameGTE applies the GTE predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameGTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameLT applies the LT predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameLT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameLTE applies the LTE predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameLTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameContains applies the Contains predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameContains(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContains(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameHasPrefix applies the HasPrefix predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameHasPrefix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasPrefix(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameHasSuffix applies the HasSuffix predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameHasSuffix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasSuffix(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameIsNil applies the IsNil predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldFinancialIndicatorsName))
}

// FinancialIndicatorsNameNotNil applies the NotNil predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldFinancialIndicatorsName))
}

// FinancialIndicatorsNameEqualFold applies the EqualFold predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameEqualFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEqualFold(FieldFinancialIndicatorsName, v))
}

// FinancialIndicatorsNameContainsFold applies the ContainsFold predicate on the "financial_indicators_name" field.
func FinancialIndicatorsNameContainsFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContainsFold(FieldFinancialIndicatorsName, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldCreateTime))
}

// RemarksEQ applies the EQ predicate on the "remarks" field.
func RemarksEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldRemarks, v))
}

// RemarksNEQ applies the NEQ predicate on the "remarks" field.
func RemarksNEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldRemarks, v))
}

// RemarksIn applies the In predicate on the "remarks" field.
func RemarksIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldRemarks, vs...))
}

// RemarksNotIn applies the NotIn predicate on the "remarks" field.
func RemarksNotIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldRemarks, vs...))
}

// RemarksGT applies the GT predicate on the "remarks" field.
func RemarksGT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldRemarks, v))
}

// RemarksGTE applies the GTE predicate on the "remarks" field.
func RemarksGTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldRemarks, v))
}

// RemarksLT applies the LT predicate on the "remarks" field.
func RemarksLT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldRemarks, v))
}

// RemarksLTE applies the LTE predicate on the "remarks" field.
func RemarksLTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldRemarks, v))
}

// RemarksContains applies the Contains predicate on the "remarks" field.
func RemarksContains(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContains(FieldRemarks, v))
}

// RemarksHasPrefix applies the HasPrefix predicate on the "remarks" field.
func RemarksHasPrefix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasPrefix(FieldRemarks, v))
}

// RemarksHasSuffix applies the HasSuffix predicate on the "remarks" field.
func RemarksHasSuffix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasSuffix(FieldRemarks, v))
}

// RemarksIsNil applies the IsNil predicate on the "remarks" field.
func RemarksIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldRemarks))
}

// RemarksNotNil applies the NotNil predicate on the "remarks" field.
func RemarksNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldRemarks))
}

// RemarksEqualFold applies the EqualFold predicate on the "remarks" field.
func RemarksEqualFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEqualFold(FieldRemarks, v))
}

// RemarksContainsFold applies the ContainsFold predicate on the "remarks" field.
func RemarksContainsFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContainsFold(FieldRemarks, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.FieldContainsFold(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmFinancialIndicator) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmFinancialIndicator) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmFinancialIndicator) predicate.EmFinancialIndicator {
	return predicate.EmFinancialIndicator(sql.NotPredicates(p))
}
