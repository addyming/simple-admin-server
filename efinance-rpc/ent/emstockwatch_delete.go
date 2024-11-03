// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockwatch"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
)

// EmStockWatchDelete is the builder for deleting a EmStockWatch entity.
type EmStockWatchDelete struct {
	config
	hooks    []Hook
	mutation *EmStockWatchMutation
}

// Where appends a list predicates to the EmStockWatchDelete builder.
func (eswd *EmStockWatchDelete) Where(ps ...predicate.EmStockWatch) *EmStockWatchDelete {
	eswd.mutation.Where(ps...)
	return eswd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eswd *EmStockWatchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eswd.sqlExec, eswd.mutation, eswd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eswd *EmStockWatchDelete) ExecX(ctx context.Context) int {
	n, err := eswd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eswd *EmStockWatchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emstockwatch.Table, sqlgraph.NewFieldSpec(emstockwatch.FieldID, field.TypeInt32))
	if ps := eswd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eswd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eswd.mutation.done = true
	return affected, err
}

// EmStockWatchDeleteOne is the builder for deleting a single EmStockWatch entity.
type EmStockWatchDeleteOne struct {
	eswd *EmStockWatchDelete
}

// Where appends a list predicates to the EmStockWatchDelete builder.
func (eswdo *EmStockWatchDeleteOne) Where(ps ...predicate.EmStockWatch) *EmStockWatchDeleteOne {
	eswdo.eswd.mutation.Where(ps...)
	return eswdo
}

// Exec executes the deletion query.
func (eswdo *EmStockWatchDeleteOne) Exec(ctx context.Context) error {
	n, err := eswdo.eswd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emstockwatch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eswdo *EmStockWatchDeleteOne) ExecX(ctx context.Context) {
	if err := eswdo.Exec(ctx); err != nil {
		panic(err)
	}
}
