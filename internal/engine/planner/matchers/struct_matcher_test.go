// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package matchers

import (
	"testing"

	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/stretchr/testify/require"
)

func TestStructMatcher(t *testing.T) {
	tests := []struct {
		expr string
		res  bool
	}{
		{
			expr: "{\"a\": 3}[R.attr.Id] == 4",
			res:  true,
		},
		{
			expr: "4 == {\"a\": 3}[R.attr.Id]", // can't swap args
		},
		{
			expr: "P.attr.anyMap[R.attr.Id] == R.attr.value",
		},
		{
			expr: "P.attr.anyMap[R.attr.Id][R.attr.value]",
		},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			ast, issues := conditions.StdEnv.Compile(test.expr)
			require.Nil(t, issues.Err())
			s := NewStructMatcher()
			res, _, err := s.Process(ast.Expr())
			require.NoError(t, err)
			require.Equal(t, res, test.res)
		})
	}
}