// Copyright 2025 The Contributors to go-hedera-app
// This file is part of the go-hedera-app project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the License.
//
// For a list of contributors, refer to the CONTRIBUTORS file or the
// repository's commit history.

package wallet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClientPrivKey(t *testing.T) {
	testcases := []struct {
		input string
		want  error
	}{
		{
			input: "abc",
			want:  ErrInvalidPrivKeyStringFormat,
		},
		{
			input: "Ox1qa",
			want:  ErrInvalidPrivKeyStringFormat,
		},
	}

	for i, tc := range testcases {
		_, got := CreateClient(tc.input, "")
		assert.ErrorIs(t, got, tc.want, fmt.Sprintf("Case %d Want: %v Got: %v", i, tc.want, got))
	}
}
