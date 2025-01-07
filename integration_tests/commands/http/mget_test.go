// This file is part of DiceDB.
// Copyright (C) 2024 DiceDB (dicedb.io).
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

//go:build ignore
// +build ignore

// Ignored as multishard commands not supported by HTTP
package http

import (
	"testing"

	"github.com/dicedb/dice/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMGET(t *testing.T) {
	exec := NewHTTPCommandExecutor()
	exec.FireCommand(HTTPCommand{
		Command: "DEL",
		Body: map[string]interface{}{
			"keys": []interface{}{"k1", "k2"},
		},
	})

	testCases := []TestCase{
		{
			name: "MGET With non-existing keys",
			commands: []HTTPCommand{
				{Command: "MGET", Body: map[string]interface{}{"keys": []interface{}{"k1", "k2"}}},
			},
			expected: []interface{}{[]interface{}{nil, nil}},
		},
		{
			name: "MGET With existing keys",
			commands: []HTTPCommand{
				{
					Command: "MSET",
					Body:    map[string]interface{}{"key_values": map[string]interface{}{"k1": "v1", "k2": "v2"}},
				},
				{Command: "MGET", Body: map[string]interface{}{"keys": []interface{}{"k1", "k2"}}},
			},
			expected: []interface{}{"OK", []interface{}{"v1", "v2"}},
		},
		{
			name: "MGET with existing and non existing keys",
			commands: []HTTPCommand{
				{Command: "SET", Body: map[string]interface{}{"key": "k1", "value": "v1"}},
				{Command: "MGET", Body: map[string]interface{}{"keys": []interface{}{"k1", "k3"}}},
			},
			expected: []interface{}{"OK", []interface{}{"v1", nil}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.commands {
				result, _ := exec.FireCommand(cmd)
				if slice, ok := tc.expected[i].([]interface{}); ok {
					assert.True(t, testutils.UnorderedEqual(slice, result))
				} else {
					assert.Equal(t, tc.expected[i], result)
				}
			}
		})
	}

	// Deleting the used keys
	exec.FireCommand(HTTPCommand{
		Command: "DEL",
		Body: map[string]interface{}{
			"keys": []interface{}{"k1", "k2"},
		},
	})
}