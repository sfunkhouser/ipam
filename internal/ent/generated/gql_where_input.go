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
	"errors"
	"fmt"
	"time"

	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// IPAddressWhereInput represents a where input for filtering IPAddress queries.
type IPAddressWhereInput struct {
	Predicates []predicate.IPAddress  `json:"-"`
	Not        *IPAddressWhereInput   `json:"not,omitempty"`
	Or         []*IPAddressWhereInput `json:"or,omitempty"`
	And        []*IPAddressWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNEQ   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGT    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGTE   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLT    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLTE   *gidx.PrefixedID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "IP" field predicates.
	IP             *string  `json:"ip,omitempty"`
	IPNEQ          *string  `json:"ipNEQ,omitempty"`
	IPIn           []string `json:"ipIn,omitempty"`
	IPNotIn        []string `json:"ipNotIn,omitempty"`
	IPGT           *string  `json:"ipGT,omitempty"`
	IPGTE          *string  `json:"ipGTE,omitempty"`
	IPLT           *string  `json:"ipLT,omitempty"`
	IPLTE          *string  `json:"ipLTE,omitempty"`
	IPContains     *string  `json:"ipContains,omitempty"`
	IPHasPrefix    *string  `json:"ipHasPrefix,omitempty"`
	IPHasSuffix    *string  `json:"ipHasSuffix,omitempty"`
	IPEqualFold    *string  `json:"ipEqualFold,omitempty"`
	IPContainsFold *string  `json:"ipContainsFold,omitempty"`

	// "reserved" field predicates.
	Reserved    *bool `json:"reserved,omitempty"`
	ReservedNEQ *bool `json:"reservedNEQ,omitempty"`

	// "ip_block" edge predicates.
	HasIPBlock     *bool                `json:"hasIPBlock,omitempty"`
	HasIPBlockWith []*IPBlockWhereInput `json:"hasIPBlockWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *IPAddressWhereInput) AddPredicates(predicates ...predicate.IPAddress) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the IPAddressWhereInput filter on the IPAddressQuery builder.
func (i *IPAddressWhereInput) Filter(q *IPAddressQuery) (*IPAddressQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyIPAddressWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyIPAddressWhereInput is returned in case the IPAddressWhereInput is empty.
var ErrEmptyIPAddressWhereInput = errors.New("generated: empty predicate IPAddressWhereInput")

// P returns a predicate for filtering ipaddresses.
// An error is returned if the input is empty or invalid.
func (i *IPAddressWhereInput) P() (predicate.IPAddress, error) {
	var predicates []predicate.IPAddress
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, ipaddress.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.IPAddress, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, ipaddress.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.IPAddress, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, ipaddress.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, ipaddress.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, ipaddress.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, ipaddress.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, ipaddress.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, ipaddress.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, ipaddress.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, ipaddress.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, ipaddress.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, ipaddress.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, ipaddress.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, ipaddress.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, ipaddress.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, ipaddress.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, ipaddress.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, ipaddress.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, ipaddress.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, ipaddress.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, ipaddress.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, ipaddress.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, ipaddress.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, ipaddress.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, ipaddress.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, ipaddress.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, ipaddress.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.IP != nil {
		predicates = append(predicates, ipaddress.IPEQ(*i.IP))
	}
	if i.IPNEQ != nil {
		predicates = append(predicates, ipaddress.IPNEQ(*i.IPNEQ))
	}
	if len(i.IPIn) > 0 {
		predicates = append(predicates, ipaddress.IPIn(i.IPIn...))
	}
	if len(i.IPNotIn) > 0 {
		predicates = append(predicates, ipaddress.IPNotIn(i.IPNotIn...))
	}
	if i.IPGT != nil {
		predicates = append(predicates, ipaddress.IPGT(*i.IPGT))
	}
	if i.IPGTE != nil {
		predicates = append(predicates, ipaddress.IPGTE(*i.IPGTE))
	}
	if i.IPLT != nil {
		predicates = append(predicates, ipaddress.IPLT(*i.IPLT))
	}
	if i.IPLTE != nil {
		predicates = append(predicates, ipaddress.IPLTE(*i.IPLTE))
	}
	if i.IPContains != nil {
		predicates = append(predicates, ipaddress.IPContains(*i.IPContains))
	}
	if i.IPHasPrefix != nil {
		predicates = append(predicates, ipaddress.IPHasPrefix(*i.IPHasPrefix))
	}
	if i.IPHasSuffix != nil {
		predicates = append(predicates, ipaddress.IPHasSuffix(*i.IPHasSuffix))
	}
	if i.IPEqualFold != nil {
		predicates = append(predicates, ipaddress.IPEqualFold(*i.IPEqualFold))
	}
	if i.IPContainsFold != nil {
		predicates = append(predicates, ipaddress.IPContainsFold(*i.IPContainsFold))
	}
	if i.Reserved != nil {
		predicates = append(predicates, ipaddress.ReservedEQ(*i.Reserved))
	}
	if i.ReservedNEQ != nil {
		predicates = append(predicates, ipaddress.ReservedNEQ(*i.ReservedNEQ))
	}

	if i.HasIPBlock != nil {
		p := ipaddress.HasIPBlock()
		if !*i.HasIPBlock {
			p = ipaddress.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasIPBlockWith) > 0 {
		with := make([]predicate.IPBlock, 0, len(i.HasIPBlockWith))
		for _, w := range i.HasIPBlockWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasIPBlockWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, ipaddress.HasIPBlockWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyIPAddressWhereInput
	case 1:
		return predicates[0], nil
	default:
		return ipaddress.And(predicates...), nil
	}
}

// IPBlockWhereInput represents a where input for filtering IPBlock queries.
type IPBlockWhereInput struct {
	Predicates []predicate.IPBlock  `json:"-"`
	Not        *IPBlockWhereInput   `json:"not,omitempty"`
	Or         []*IPBlockWhereInput `json:"or,omitempty"`
	And        []*IPBlockWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNEQ   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGT    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGTE   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLT    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLTE   *gidx.PrefixedID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "prefix" field predicates.
	Prefix             *string  `json:"prefix,omitempty"`
	PrefixNEQ          *string  `json:"prefixNEQ,omitempty"`
	PrefixIn           []string `json:"prefixIn,omitempty"`
	PrefixNotIn        []string `json:"prefixNotIn,omitempty"`
	PrefixGT           *string  `json:"prefixGT,omitempty"`
	PrefixGTE          *string  `json:"prefixGTE,omitempty"`
	PrefixLT           *string  `json:"prefixLT,omitempty"`
	PrefixLTE          *string  `json:"prefixLTE,omitempty"`
	PrefixContains     *string  `json:"prefixContains,omitempty"`
	PrefixHasPrefix    *string  `json:"prefixHasPrefix,omitempty"`
	PrefixHasSuffix    *string  `json:"prefixHasSuffix,omitempty"`
	PrefixEqualFold    *string  `json:"prefixEqualFold,omitempty"`
	PrefixContainsFold *string  `json:"prefixContainsFold,omitempty"`

	// "allow_auto_subnet" field predicates.
	AllowAutoSubnet    *bool `json:"allowAutoSubnet,omitempty"`
	AllowAutoSubnetNEQ *bool `json:"allowAutoSubnetNEQ,omitempty"`

	// "allow_auto_allocate" field predicates.
	AllowAutoAllocate    *bool `json:"allowAutoAllocate,omitempty"`
	AllowAutoAllocateNEQ *bool `json:"allowAutoAllocateNEQ,omitempty"`

	// "ip_block_type" edge predicates.
	HasIPBlockType     *bool                    `json:"hasIPBlockType,omitempty"`
	HasIPBlockTypeWith []*IPBlockTypeWhereInput `json:"hasIPBlockTypeWith,omitempty"`

	// "ip_address" edge predicates.
	HasIPAddress     *bool                  `json:"hasIPAddress,omitempty"`
	HasIPAddressWith []*IPAddressWhereInput `json:"hasIPAddressWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *IPBlockWhereInput) AddPredicates(predicates ...predicate.IPBlock) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the IPBlockWhereInput filter on the IPBlockQuery builder.
func (i *IPBlockWhereInput) Filter(q *IPBlockQuery) (*IPBlockQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyIPBlockWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyIPBlockWhereInput is returned in case the IPBlockWhereInput is empty.
var ErrEmptyIPBlockWhereInput = errors.New("generated: empty predicate IPBlockWhereInput")

// P returns a predicate for filtering ipblocks.
// An error is returned if the input is empty or invalid.
func (i *IPBlockWhereInput) P() (predicate.IPBlock, error) {
	var predicates []predicate.IPBlock
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, ipblock.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.IPBlock, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, ipblock.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.IPBlock, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, ipblock.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, ipblock.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, ipblock.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, ipblock.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, ipblock.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, ipblock.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, ipblock.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, ipblock.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, ipblock.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, ipblock.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, ipblock.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, ipblock.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, ipblock.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, ipblock.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, ipblock.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, ipblock.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, ipblock.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, ipblock.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, ipblock.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, ipblock.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, ipblock.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, ipblock.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, ipblock.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, ipblock.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, ipblock.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Prefix != nil {
		predicates = append(predicates, ipblock.PrefixEQ(*i.Prefix))
	}
	if i.PrefixNEQ != nil {
		predicates = append(predicates, ipblock.PrefixNEQ(*i.PrefixNEQ))
	}
	if len(i.PrefixIn) > 0 {
		predicates = append(predicates, ipblock.PrefixIn(i.PrefixIn...))
	}
	if len(i.PrefixNotIn) > 0 {
		predicates = append(predicates, ipblock.PrefixNotIn(i.PrefixNotIn...))
	}
	if i.PrefixGT != nil {
		predicates = append(predicates, ipblock.PrefixGT(*i.PrefixGT))
	}
	if i.PrefixGTE != nil {
		predicates = append(predicates, ipblock.PrefixGTE(*i.PrefixGTE))
	}
	if i.PrefixLT != nil {
		predicates = append(predicates, ipblock.PrefixLT(*i.PrefixLT))
	}
	if i.PrefixLTE != nil {
		predicates = append(predicates, ipblock.PrefixLTE(*i.PrefixLTE))
	}
	if i.PrefixContains != nil {
		predicates = append(predicates, ipblock.PrefixContains(*i.PrefixContains))
	}
	if i.PrefixHasPrefix != nil {
		predicates = append(predicates, ipblock.PrefixHasPrefix(*i.PrefixHasPrefix))
	}
	if i.PrefixHasSuffix != nil {
		predicates = append(predicates, ipblock.PrefixHasSuffix(*i.PrefixHasSuffix))
	}
	if i.PrefixEqualFold != nil {
		predicates = append(predicates, ipblock.PrefixEqualFold(*i.PrefixEqualFold))
	}
	if i.PrefixContainsFold != nil {
		predicates = append(predicates, ipblock.PrefixContainsFold(*i.PrefixContainsFold))
	}
	if i.AllowAutoSubnet != nil {
		predicates = append(predicates, ipblock.AllowAutoSubnetEQ(*i.AllowAutoSubnet))
	}
	if i.AllowAutoSubnetNEQ != nil {
		predicates = append(predicates, ipblock.AllowAutoSubnetNEQ(*i.AllowAutoSubnetNEQ))
	}
	if i.AllowAutoAllocate != nil {
		predicates = append(predicates, ipblock.AllowAutoAllocateEQ(*i.AllowAutoAllocate))
	}
	if i.AllowAutoAllocateNEQ != nil {
		predicates = append(predicates, ipblock.AllowAutoAllocateNEQ(*i.AllowAutoAllocateNEQ))
	}

	if i.HasIPBlockType != nil {
		p := ipblock.HasIPBlockType()
		if !*i.HasIPBlockType {
			p = ipblock.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasIPBlockTypeWith) > 0 {
		with := make([]predicate.IPBlockType, 0, len(i.HasIPBlockTypeWith))
		for _, w := range i.HasIPBlockTypeWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasIPBlockTypeWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, ipblock.HasIPBlockTypeWith(with...))
	}
	if i.HasIPAddress != nil {
		p := ipblock.HasIPAddress()
		if !*i.HasIPAddress {
			p = ipblock.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasIPAddressWith) > 0 {
		with := make([]predicate.IPAddress, 0, len(i.HasIPAddressWith))
		for _, w := range i.HasIPAddressWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasIPAddressWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, ipblock.HasIPAddressWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyIPBlockWhereInput
	case 1:
		return predicates[0], nil
	default:
		return ipblock.And(predicates...), nil
	}
}

// IPBlockTypeWhereInput represents a where input for filtering IPBlockType queries.
type IPBlockTypeWhereInput struct {
	Predicates []predicate.IPBlockType  `json:"-"`
	Not        *IPBlockTypeWhereInput   `json:"not,omitempty"`
	Or         []*IPBlockTypeWhereInput `json:"or,omitempty"`
	And        []*IPBlockTypeWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNEQ   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGT    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGTE   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLT    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLTE   *gidx.PrefixedID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "ip_block" edge predicates.
	HasIPBlock     *bool                `json:"hasIPBlock,omitempty"`
	HasIPBlockWith []*IPBlockWhereInput `json:"hasIPBlockWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *IPBlockTypeWhereInput) AddPredicates(predicates ...predicate.IPBlockType) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the IPBlockTypeWhereInput filter on the IPBlockTypeQuery builder.
func (i *IPBlockTypeWhereInput) Filter(q *IPBlockTypeQuery) (*IPBlockTypeQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyIPBlockTypeWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyIPBlockTypeWhereInput is returned in case the IPBlockTypeWhereInput is empty.
var ErrEmptyIPBlockTypeWhereInput = errors.New("generated: empty predicate IPBlockTypeWhereInput")

// P returns a predicate for filtering ipblocktypes.
// An error is returned if the input is empty or invalid.
func (i *IPBlockTypeWhereInput) P() (predicate.IPBlockType, error) {
	var predicates []predicate.IPBlockType
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, ipblocktype.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.IPBlockType, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, ipblocktype.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.IPBlockType, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, ipblocktype.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, ipblocktype.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, ipblocktype.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, ipblocktype.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, ipblocktype.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, ipblocktype.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, ipblocktype.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, ipblocktype.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, ipblocktype.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, ipblocktype.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, ipblocktype.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, ipblocktype.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, ipblocktype.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, ipblocktype.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, ipblocktype.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, ipblocktype.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, ipblocktype.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, ipblocktype.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, ipblocktype.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, ipblocktype.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, ipblocktype.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, ipblocktype.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, ipblocktype.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, ipblocktype.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, ipblocktype.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, ipblocktype.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, ipblocktype.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, ipblocktype.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, ipblocktype.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, ipblocktype.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, ipblocktype.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, ipblocktype.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, ipblocktype.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasIPBlock != nil {
		p := ipblocktype.HasIPBlock()
		if !*i.HasIPBlock {
			p = ipblocktype.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasIPBlockWith) > 0 {
		with := make([]predicate.IPBlock, 0, len(i.HasIPBlockWith))
		for _, w := range i.HasIPBlockWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasIPBlockWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, ipblocktype.HasIPBlockWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyIPBlockTypeWhereInput
	case 1:
		return predicates[0], nil
	default:
		return ipblocktype.And(predicates...), nil
	}
}