/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package services

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/utils"
)

// TestOktaImportRuleUnmarshal verifies an Okta import rule resource can be unmarshaled.
func TestOktaImportRuleUnmarshal(t *testing.T) {
	expected, err := types.NewOktaImportRule(
		types.Metadata{
			Name: "test-rule",
		},
		types.OktaImportRuleSpecV1{
			Mappings: []*types.OktaImportRuleMappingV1{
				{
					Match: []*types.OktaImportRuleMatchV1{
						{
							AppIDs: []string{"yes"},
						},
					},
					AddLabels: map[string]string{
						"label1": "value1",
					},
				},
				{
					Match: []*types.OktaImportRuleMatchV1{
						{
							GroupIDs: []string{"yes"},
						},
					},
					AddLabels: map[string]string{
						"label1": "value1",
					},
				},
			},
		},
	)
	require.NoError(t, err)
	data, err := utils.ToJSON([]byte(oktaImportRuleYAML))
	require.NoError(t, err)
	actual, err := UnmarshalOktaImportRule(data)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

// TestOktaImportRuleMarshal verifies a marshaled Okta import rule resource can be unmarshaled back.
func TestOktaImportRuleMarshal(t *testing.T) {
	expected, err := types.NewOktaImportRule(
		types.Metadata{
			Name: "test-rule",
		},
		types.OktaImportRuleSpecV1{
			Mappings: []*types.OktaImportRuleMappingV1{
				{
					Match: []*types.OktaImportRuleMatchV1{
						{
							AppIDs: []string{"yes"},
						},
					},
					AddLabels: map[string]string{
						"label1": "value1",
					},
				},
				{
					Match: []*types.OktaImportRuleMatchV1{
						{
							GroupIDs: []string{"yes"},
						},
					},
					AddLabels: map[string]string{
						"label1": "value1",
					},
				},
			},
		},
	)
	require.NoError(t, err)
	data, err := MarshalOktaImportRule(expected)
	require.NoError(t, err)
	actual, err := UnmarshalOktaImportRule(data)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

// TestOktaAssignmentUnmarshal verifies an Okta assignment resource can be unmarshaled.
func TestOktaAssignmentUnmarshal(t *testing.T) {
	expected, err := types.NewOktaAssignment(
		types.Metadata{
			Name: "test-assignment",
		},
		types.OktaAssignmentSpecV1{
			User: "test-user@test.user",
			Targets: []*types.OktaAssignmentTargetV1{
				{
					Type: types.OktaAssignmentTargetV1_APPLICATION,
					Id:   "123456",
				},
				{
					Type: types.OktaAssignmentTargetV1_GROUP,
					Id:   "234567",
				},
			},
			Status: types.OktaAssignmentSpecV1_PENDING,
		},
	)
	require.NoError(t, err)
	data, err := utils.ToJSON([]byte(oktaAssignmentYAML))
	require.NoError(t, err)
	actual, err := UnmarshalOktaAssignment(data)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

// TestOktaAssignmentMarshal verifies a marshaled Okta assignment resource can be unmarshaled back.
func TestOktaAssignmentMarshal(t *testing.T) {
	expected, err := types.NewOktaAssignment(
		types.Metadata{
			Name: "test-assignment",
		},
		types.OktaAssignmentSpecV1{
			User: "test-user@test.user",
			Targets: []*types.OktaAssignmentTargetV1{
				{
					Type: types.OktaAssignmentTargetV1_APPLICATION,
					Id:   "123456",
				},
				{
					Type: types.OktaAssignmentTargetV1_GROUP,
					Id:   "234567",
				},
			},

			Status: types.OktaAssignmentSpecV1_PENDING,
		},
	)
	require.NoError(t, err)
	data, err := MarshalOktaAssignment(expected)
	require.NoError(t, err)
	actual, err := UnmarshalOktaAssignment(data)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

var (
	oktaImportRuleYAML = `---
kind: okta_import_rule
version: v1
metadata:
  name: test-rule
spec:
  mappings:
  - match:
    - app_ids: ["yes"]
    add_labels:
      label1: value1
  - match:
    - group_ids: ["yes"]
    add_labels:
      label1: value1
`

	oktaAssignmentYAML = `---
kind: okta_assignment
version: v1
metadata:
  name: test-assignment
spec:
  user: test-user@test.user
  targets:
  - type: 1
    id: "123456"
  - type: 2
    id: "234567"
  status: 1
`
)
