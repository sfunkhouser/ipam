// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
)

// IPAddressDelete is the builder for deleting a IPAddress entity.
type IPAddressDelete struct {
	config
	hooks    []Hook
	mutation *IPAddressMutation
}

// Where appends a list predicates to the IPAddressDelete builder.
func (iad *IPAddressDelete) Where(ps ...predicate.IPAddress) *IPAddressDelete {
	iad.mutation.Where(ps...)
	return iad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iad *IPAddressDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, iad.sqlExec, iad.mutation, iad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (iad *IPAddressDelete) ExecX(ctx context.Context) int {
	n, err := iad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iad *IPAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ipaddress.Table, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString))
	if ps := iad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, iad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	iad.mutation.done = true
	return affected, err
}

// IPAddressDeleteOne is the builder for deleting a single IPAddress entity.
type IPAddressDeleteOne struct {
	iad *IPAddressDelete
}

// Where appends a list predicates to the IPAddressDelete builder.
func (iado *IPAddressDeleteOne) Where(ps ...predicate.IPAddress) *IPAddressDeleteOne {
	iado.iad.mutation.Where(ps...)
	return iado
}

// Exec executes the deletion query.
func (iado *IPAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := iado.iad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ipaddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iado *IPAddressDeleteOne) ExecX(ctx context.Context) {
	if err := iado.Exec(ctx); err != nil {
		panic(err)
	}
}
