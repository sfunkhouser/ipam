// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package pubsubhooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/ent/generated/hook"
	"go.infratographer.com/ipam-api/internal/ent/schema"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
)

func IPAddressHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPAddressFunc(func(ctx context.Context, m *generated.IPAddressMutation) (ent.Value, error) {
					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_IP := ""
					IP, ok := m.IP()

					if ok {
						cv_IP = fmt.Sprintf("%s", fmt.Sprint(IP))
						pv_IP := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldIP(ctx)
							if err != nil {
								pv_IP = "<unknown>"
							} else {
								pv_IP = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "IP",
							PreviousValue: pv_IP,
							CurrentValue:  cv_IP,
						})
					}

					cv_block_id := ""
					block_id, ok := m.BlockID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						block_id, err = m.OldBlockID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, block_id)

					if ok {
						cv_block_id = fmt.Sprintf("%s", fmt.Sprint(block_id))
						pv_block_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldBlockID(ctx)
							if err != nil {
								pv_block_id = "<unknown>"
							} else {
								pv_block_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "block_id",
							PreviousValue: pv_block_id,
							CurrentValue:  cv_block_id,
						})
					}

					cv_node_id := ""
					node_id, ok := m.NodeID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						node_id, err = m.OldNodeID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, node_id)

					if ok {
						cv_node_id = fmt.Sprintf("%s", fmt.Sprint(node_id))
						pv_node_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldNodeID(ctx)
							if err != nil {
								pv_node_id = "<unknown>"
							} else {
								pv_node_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "node_id",
							PreviousValue: pv_node_id,
							CurrentValue:  cv_node_id,
						})
					}

					cv_node_owner_id := ""
					node_owner_id, ok := m.NodeOwnerID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						node_owner_id, err = m.OldNodeOwnerID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, node_owner_id)

					if ok {
						cv_node_owner_id = fmt.Sprintf("%s", fmt.Sprint(node_owner_id))
						pv_node_owner_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldNodeOwnerID(ctx)
							if err != nil {
								pv_node_owner_id = "<unknown>"
							} else {
								pv_node_owner_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "node_owner_id",
							PreviousValue: pv_node_owner_id,
							CurrentValue:  cv_node_owner_id,
						})
					}

					cv_reserved := ""
					reserved, ok := m.Reserved()

					if ok {
						cv_reserved = fmt.Sprintf("%s", fmt.Sprint(reserved))
						pv_reserved := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldReserved(ctx)
							if err != nil {
								pv_reserved = "<unknown>"
							} else {
								pv_reserved = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "reserved",
							PreviousValue: pv_reserved,
							CurrentValue:  cv_reserved,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPAddressFunc(func(ctx context.Context, m *generated.IPAddressMutation) (ent.Value, error) {
					// queueName := ""
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().IPAddress.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for pubsub, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.BlockID)

					additionalSubjects = append(additionalSubjects, dbObj.NodeID)

					additionalSubjects = append(additionalSubjects, dbObj.NodeOwnerID)

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func IPBlockHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPBlockFunc(func(ctx context.Context, m *generated.IPBlockMutation) (ent.Value, error) {
					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_prefix := ""
					prefix, ok := m.Prefix()

					if ok {
						cv_prefix = fmt.Sprintf("%s", fmt.Sprint(prefix))
						pv_prefix := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldPrefix(ctx)
							if err != nil {
								pv_prefix = "<unknown>"
							} else {
								pv_prefix = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "prefix",
							PreviousValue: pv_prefix,
							CurrentValue:  cv_prefix,
						})
					}

					cv_block_type_id := ""
					block_type_id, ok := m.BlockTypeID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						block_type_id, err = m.OldBlockTypeID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, block_type_id)

					if ok {
						cv_block_type_id = fmt.Sprintf("%s", fmt.Sprint(block_type_id))
						pv_block_type_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldBlockTypeID(ctx)
							if err != nil {
								pv_block_type_id = "<unknown>"
							} else {
								pv_block_type_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "block_type_id",
							PreviousValue: pv_block_type_id,
							CurrentValue:  cv_block_type_id,
						})
					}

					cv_location_id := ""
					location_id, ok := m.LocationID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						location_id, err = m.OldLocationID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, location_id)

					if ok {
						cv_location_id = fmt.Sprintf("%s", fmt.Sprint(location_id))
						pv_location_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldLocationID(ctx)
							if err != nil {
								pv_location_id = "<unknown>"
							} else {
								pv_location_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "location_id",
							PreviousValue: pv_location_id,
							CurrentValue:  cv_location_id,
						})
					}

					cv_parent_block_id := ""
					parent_block_id, ok := m.ParentBlockID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						parent_block_id, err = m.OldParentBlockID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, parent_block_id)

					if ok {
						cv_parent_block_id = fmt.Sprintf("%s", fmt.Sprint(parent_block_id))
						pv_parent_block_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldParentBlockID(ctx)
							if err != nil {
								pv_parent_block_id = "<unknown>"
							} else {
								pv_parent_block_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "parent_block_id",
							PreviousValue: pv_parent_block_id,
							CurrentValue:  cv_parent_block_id,
						})
					}

					cv_allow_auto_subnet := ""
					allow_auto_subnet, ok := m.AllowAutoSubnet()

					if ok {
						cv_allow_auto_subnet = fmt.Sprintf("%s", fmt.Sprint(allow_auto_subnet))
						pv_allow_auto_subnet := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldAllowAutoSubnet(ctx)
							if err != nil {
								pv_allow_auto_subnet = "<unknown>"
							} else {
								pv_allow_auto_subnet = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "allow_auto_subnet",
							PreviousValue: pv_allow_auto_subnet,
							CurrentValue:  cv_allow_auto_subnet,
						})
					}

					cv_allow_auto_allocate := ""
					allow_auto_allocate, ok := m.AllowAutoAllocate()

					if ok {
						cv_allow_auto_allocate = fmt.Sprintf("%s", fmt.Sprint(allow_auto_allocate))
						pv_allow_auto_allocate := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldAllowAutoAllocate(ctx)
							if err != nil {
								pv_allow_auto_allocate = "<unknown>"
							} else {
								pv_allow_auto_allocate = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "allow_auto_allocate",
							PreviousValue: pv_allow_auto_allocate,
							CurrentValue:  cv_allow_auto_allocate,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPBlockFunc(func(ctx context.Context, m *generated.IPBlockMutation) (ent.Value, error) {
					// queueName := ""
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().IPBlock.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for pubsub, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.BlockTypeID)

					additionalSubjects = append(additionalSubjects, dbObj.LocationID)

					additionalSubjects = append(additionalSubjects, dbObj.ParentBlockID)

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
func IPBlockTypeHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPBlockTypeFunc(func(ctx context.Context, m *generated.IPBlockTypeMutation) (ent.Value, error) {
					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_name := ""
					name, ok := m.Name()

					if ok {
						cv_name = fmt.Sprintf("%s", fmt.Sprint(name))
						pv_name := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldName(ctx)
							if err != nil {
								pv_name = "<unknown>"
							} else {
								pv_name = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "name",
							PreviousValue: pv_name,
							CurrentValue:  cv_name,
						})
					}

					cv_owner_id := ""
					owner_id, ok := m.OwnerID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						owner_id, err = m.OldOwnerID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, owner_id)

					if ok {
						cv_owner_id = fmt.Sprintf("%s", fmt.Sprint(owner_id))
						pv_owner_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldOwnerID(ctx)
							if err != nil {
								pv_owner_id = "<unknown>"
							} else {
								pv_owner_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "owner_id",
							PreviousValue: pv_owner_id,
							CurrentValue:  cv_owner_id,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.IPBlockTypeFunc(func(ctx context.Context, m *generated.IPBlockTypeMutation) (ent.Value, error) {
					// queueName := ""
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().IPBlockType.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for pubsub, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.OwnerID)

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					m.EventsPublisher.PublishChange(ctx, eventSubject(objID), msg)

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

func PubsubHooks(c *generated.Client) {
	c.IPAddress.Use(IPAddressHooks()...)

	c.IPBlock.Use(IPBlockHooks()...)

	c.IPBlockType.Use(IPBlockTypeHooks()...)

}

func eventType(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return "create"
	case ent.OpUpdate, ent.OpUpdateOne:
		return "update"
	case ent.OpDelete, ent.OpDeleteOne:
		return "delete"
	default:
		return "unknown"
	}
}

func eventSubject(objID gidx.PrefixedID) string {
	switch objID.Prefix() {
	case schema.IPBlockTypePrefix:
		return "ipam-block-type"
	case schema.IPBlockPrefix:
		return "ipam-block"
	case schema.IPAddressPrefix:
		return "ipam-ip-address"
	default:
		return "unknown"
	}
}
