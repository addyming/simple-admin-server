// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emfinancialindicator"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
)

// EmFinancialIndicatorUpdate is the builder for updating EmFinancialIndicator entities.
type EmFinancialIndicatorUpdate struct {
	config
	hooks    []Hook
	mutation *EmFinancialIndicatorMutation
}

// Where appends a list predicates to the EmFinancialIndicatorUpdate builder.
func (efiu *EmFinancialIndicatorUpdate) Where(ps ...predicate.EmFinancialIndicator) *EmFinancialIndicatorUpdate {
	efiu.mutation.Where(ps...)
	return efiu
}

// SetFinancialIndicatorsKey sets the "financial_indicators_key" field.
func (efiu *EmFinancialIndicatorUpdate) SetFinancialIndicatorsKey(s string) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetFinancialIndicatorsKey(s)
	return efiu
}

// SetNillableFinancialIndicatorsKey sets the "financial_indicators_key" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableFinancialIndicatorsKey(s *string) *EmFinancialIndicatorUpdate {
	if s != nil {
		efiu.SetFinancialIndicatorsKey(*s)
	}
	return efiu
}

// ClearFinancialIndicatorsKey clears the value of the "financial_indicators_key" field.
func (efiu *EmFinancialIndicatorUpdate) ClearFinancialIndicatorsKey() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearFinancialIndicatorsKey()
	return efiu
}

// SetFinancialIndicatorsValue sets the "financial_indicators_value" field.
func (efiu *EmFinancialIndicatorUpdate) SetFinancialIndicatorsValue(s string) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetFinancialIndicatorsValue(s)
	return efiu
}

// SetNillableFinancialIndicatorsValue sets the "financial_indicators_value" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableFinancialIndicatorsValue(s *string) *EmFinancialIndicatorUpdate {
	if s != nil {
		efiu.SetFinancialIndicatorsValue(*s)
	}
	return efiu
}

// ClearFinancialIndicatorsValue clears the value of the "financial_indicators_value" field.
func (efiu *EmFinancialIndicatorUpdate) ClearFinancialIndicatorsValue() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearFinancialIndicatorsValue()
	return efiu
}

// SetFinancialIndicatorsName sets the "financial_indicators_name" field.
func (efiu *EmFinancialIndicatorUpdate) SetFinancialIndicatorsName(s string) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetFinancialIndicatorsName(s)
	return efiu
}

// SetNillableFinancialIndicatorsName sets the "financial_indicators_name" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableFinancialIndicatorsName(s *string) *EmFinancialIndicatorUpdate {
	if s != nil {
		efiu.SetFinancialIndicatorsName(*s)
	}
	return efiu
}

// ClearFinancialIndicatorsName clears the value of the "financial_indicators_name" field.
func (efiu *EmFinancialIndicatorUpdate) ClearFinancialIndicatorsName() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearFinancialIndicatorsName()
	return efiu
}

// SetCreateTime sets the "create_time" field.
func (efiu *EmFinancialIndicatorUpdate) SetCreateTime(t time.Time) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetCreateTime(t)
	return efiu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableCreateTime(t *time.Time) *EmFinancialIndicatorUpdate {
	if t != nil {
		efiu.SetCreateTime(*t)
	}
	return efiu
}

// ClearCreateTime clears the value of the "create_time" field.
func (efiu *EmFinancialIndicatorUpdate) ClearCreateTime() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearCreateTime()
	return efiu
}

// SetRemarks sets the "remarks" field.
func (efiu *EmFinancialIndicatorUpdate) SetRemarks(s string) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetRemarks(s)
	return efiu
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableRemarks(s *string) *EmFinancialIndicatorUpdate {
	if s != nil {
		efiu.SetRemarks(*s)
	}
	return efiu
}

// ClearRemarks clears the value of the "remarks" field.
func (efiu *EmFinancialIndicatorUpdate) ClearRemarks() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearRemarks()
	return efiu
}

// SetStatus sets the "status" field.
func (efiu *EmFinancialIndicatorUpdate) SetStatus(s string) *EmFinancialIndicatorUpdate {
	efiu.mutation.SetStatus(s)
	return efiu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (efiu *EmFinancialIndicatorUpdate) SetNillableStatus(s *string) *EmFinancialIndicatorUpdate {
	if s != nil {
		efiu.SetStatus(*s)
	}
	return efiu
}

// ClearStatus clears the value of the "status" field.
func (efiu *EmFinancialIndicatorUpdate) ClearStatus() *EmFinancialIndicatorUpdate {
	efiu.mutation.ClearStatus()
	return efiu
}

// Mutation returns the EmFinancialIndicatorMutation object of the builder.
func (efiu *EmFinancialIndicatorUpdate) Mutation() *EmFinancialIndicatorMutation {
	return efiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (efiu *EmFinancialIndicatorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, efiu.sqlSave, efiu.mutation, efiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (efiu *EmFinancialIndicatorUpdate) SaveX(ctx context.Context) int {
	affected, err := efiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (efiu *EmFinancialIndicatorUpdate) Exec(ctx context.Context) error {
	_, err := efiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efiu *EmFinancialIndicatorUpdate) ExecX(ctx context.Context) {
	if err := efiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (efiu *EmFinancialIndicatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(emfinancialindicator.Table, emfinancialindicator.Columns, sqlgraph.NewFieldSpec(emfinancialindicator.FieldID, field.TypeInt32))
	if ps := efiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efiu.mutation.FinancialIndicatorsKey(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsKey, field.TypeString, value)
	}
	if efiu.mutation.FinancialIndicatorsKeyCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsKey, field.TypeString)
	}
	if value, ok := efiu.mutation.FinancialIndicatorsValue(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsValue, field.TypeString, value)
	}
	if efiu.mutation.FinancialIndicatorsValueCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsValue, field.TypeString)
	}
	if value, ok := efiu.mutation.FinancialIndicatorsName(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsName, field.TypeString, value)
	}
	if efiu.mutation.FinancialIndicatorsNameCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsName, field.TypeString)
	}
	if value, ok := efiu.mutation.CreateTime(); ok {
		_spec.SetField(emfinancialindicator.FieldCreateTime, field.TypeTime, value)
	}
	if efiu.mutation.CreateTimeCleared() {
		_spec.ClearField(emfinancialindicator.FieldCreateTime, field.TypeTime)
	}
	if value, ok := efiu.mutation.Remarks(); ok {
		_spec.SetField(emfinancialindicator.FieldRemarks, field.TypeString, value)
	}
	if efiu.mutation.RemarksCleared() {
		_spec.ClearField(emfinancialindicator.FieldRemarks, field.TypeString)
	}
	if value, ok := efiu.mutation.Status(); ok {
		_spec.SetField(emfinancialindicator.FieldStatus, field.TypeString, value)
	}
	if efiu.mutation.StatusCleared() {
		_spec.ClearField(emfinancialindicator.FieldStatus, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, efiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emfinancialindicator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	efiu.mutation.done = true
	return n, nil
}

// EmFinancialIndicatorUpdateOne is the builder for updating a single EmFinancialIndicator entity.
type EmFinancialIndicatorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmFinancialIndicatorMutation
}

// SetFinancialIndicatorsKey sets the "financial_indicators_key" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetFinancialIndicatorsKey(s string) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetFinancialIndicatorsKey(s)
	return efiuo
}

// SetNillableFinancialIndicatorsKey sets the "financial_indicators_key" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableFinancialIndicatorsKey(s *string) *EmFinancialIndicatorUpdateOne {
	if s != nil {
		efiuo.SetFinancialIndicatorsKey(*s)
	}
	return efiuo
}

// ClearFinancialIndicatorsKey clears the value of the "financial_indicators_key" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearFinancialIndicatorsKey() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearFinancialIndicatorsKey()
	return efiuo
}

// SetFinancialIndicatorsValue sets the "financial_indicators_value" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetFinancialIndicatorsValue(s string) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetFinancialIndicatorsValue(s)
	return efiuo
}

// SetNillableFinancialIndicatorsValue sets the "financial_indicators_value" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableFinancialIndicatorsValue(s *string) *EmFinancialIndicatorUpdateOne {
	if s != nil {
		efiuo.SetFinancialIndicatorsValue(*s)
	}
	return efiuo
}

// ClearFinancialIndicatorsValue clears the value of the "financial_indicators_value" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearFinancialIndicatorsValue() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearFinancialIndicatorsValue()
	return efiuo
}

// SetFinancialIndicatorsName sets the "financial_indicators_name" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetFinancialIndicatorsName(s string) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetFinancialIndicatorsName(s)
	return efiuo
}

// SetNillableFinancialIndicatorsName sets the "financial_indicators_name" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableFinancialIndicatorsName(s *string) *EmFinancialIndicatorUpdateOne {
	if s != nil {
		efiuo.SetFinancialIndicatorsName(*s)
	}
	return efiuo
}

// ClearFinancialIndicatorsName clears the value of the "financial_indicators_name" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearFinancialIndicatorsName() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearFinancialIndicatorsName()
	return efiuo
}

// SetCreateTime sets the "create_time" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetCreateTime(t time.Time) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetCreateTime(t)
	return efiuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableCreateTime(t *time.Time) *EmFinancialIndicatorUpdateOne {
	if t != nil {
		efiuo.SetCreateTime(*t)
	}
	return efiuo
}

// ClearCreateTime clears the value of the "create_time" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearCreateTime() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearCreateTime()
	return efiuo
}

// SetRemarks sets the "remarks" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetRemarks(s string) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetRemarks(s)
	return efiuo
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableRemarks(s *string) *EmFinancialIndicatorUpdateOne {
	if s != nil {
		efiuo.SetRemarks(*s)
	}
	return efiuo
}

// ClearRemarks clears the value of the "remarks" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearRemarks() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearRemarks()
	return efiuo
}

// SetStatus sets the "status" field.
func (efiuo *EmFinancialIndicatorUpdateOne) SetStatus(s string) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.SetStatus(s)
	return efiuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (efiuo *EmFinancialIndicatorUpdateOne) SetNillableStatus(s *string) *EmFinancialIndicatorUpdateOne {
	if s != nil {
		efiuo.SetStatus(*s)
	}
	return efiuo
}

// ClearStatus clears the value of the "status" field.
func (efiuo *EmFinancialIndicatorUpdateOne) ClearStatus() *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.ClearStatus()
	return efiuo
}

// Mutation returns the EmFinancialIndicatorMutation object of the builder.
func (efiuo *EmFinancialIndicatorUpdateOne) Mutation() *EmFinancialIndicatorMutation {
	return efiuo.mutation
}

// Where appends a list predicates to the EmFinancialIndicatorUpdate builder.
func (efiuo *EmFinancialIndicatorUpdateOne) Where(ps ...predicate.EmFinancialIndicator) *EmFinancialIndicatorUpdateOne {
	efiuo.mutation.Where(ps...)
	return efiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (efiuo *EmFinancialIndicatorUpdateOne) Select(field string, fields ...string) *EmFinancialIndicatorUpdateOne {
	efiuo.fields = append([]string{field}, fields...)
	return efiuo
}

// Save executes the query and returns the updated EmFinancialIndicator entity.
func (efiuo *EmFinancialIndicatorUpdateOne) Save(ctx context.Context) (*EmFinancialIndicator, error) {
	return withHooks(ctx, efiuo.sqlSave, efiuo.mutation, efiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (efiuo *EmFinancialIndicatorUpdateOne) SaveX(ctx context.Context) *EmFinancialIndicator {
	node, err := efiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (efiuo *EmFinancialIndicatorUpdateOne) Exec(ctx context.Context) error {
	_, err := efiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efiuo *EmFinancialIndicatorUpdateOne) ExecX(ctx context.Context) {
	if err := efiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (efiuo *EmFinancialIndicatorUpdateOne) sqlSave(ctx context.Context) (_node *EmFinancialIndicator, err error) {
	_spec := sqlgraph.NewUpdateSpec(emfinancialindicator.Table, emfinancialindicator.Columns, sqlgraph.NewFieldSpec(emfinancialindicator.FieldID, field.TypeInt32))
	id, ok := efiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmFinancialIndicator.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := efiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emfinancialindicator.FieldID)
		for _, f := range fields {
			if !emfinancialindicator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emfinancialindicator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := efiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efiuo.mutation.FinancialIndicatorsKey(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsKey, field.TypeString, value)
	}
	if efiuo.mutation.FinancialIndicatorsKeyCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsKey, field.TypeString)
	}
	if value, ok := efiuo.mutation.FinancialIndicatorsValue(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsValue, field.TypeString, value)
	}
	if efiuo.mutation.FinancialIndicatorsValueCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsValue, field.TypeString)
	}
	if value, ok := efiuo.mutation.FinancialIndicatorsName(); ok {
		_spec.SetField(emfinancialindicator.FieldFinancialIndicatorsName, field.TypeString, value)
	}
	if efiuo.mutation.FinancialIndicatorsNameCleared() {
		_spec.ClearField(emfinancialindicator.FieldFinancialIndicatorsName, field.TypeString)
	}
	if value, ok := efiuo.mutation.CreateTime(); ok {
		_spec.SetField(emfinancialindicator.FieldCreateTime, field.TypeTime, value)
	}
	if efiuo.mutation.CreateTimeCleared() {
		_spec.ClearField(emfinancialindicator.FieldCreateTime, field.TypeTime)
	}
	if value, ok := efiuo.mutation.Remarks(); ok {
		_spec.SetField(emfinancialindicator.FieldRemarks, field.TypeString, value)
	}
	if efiuo.mutation.RemarksCleared() {
		_spec.ClearField(emfinancialindicator.FieldRemarks, field.TypeString)
	}
	if value, ok := efiuo.mutation.Status(); ok {
		_spec.SetField(emfinancialindicator.FieldStatus, field.TypeString, value)
	}
	if efiuo.mutation.StatusCleared() {
		_spec.ClearField(emfinancialindicator.FieldStatus, field.TypeString)
	}
	_node = &EmFinancialIndicator{config: efiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, efiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emfinancialindicator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	efiuo.mutation.done = true
	return _node, nil
}
