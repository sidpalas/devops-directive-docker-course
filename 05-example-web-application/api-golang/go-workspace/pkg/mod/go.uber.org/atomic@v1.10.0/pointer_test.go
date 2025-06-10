// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//go:build go1.18
// +build go1.18

package atomic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointer(t *testing.T) {
	type foo struct{ v int }

	i := foo{42}
	j := foo{0}
	k := foo{1}

	tests := []struct {
		desc      string
		newAtomic func() *Pointer[foo]
		initial   *foo
	}{
		{
			desc: "New",
			newAtomic: func() *Pointer[foo] {
				return NewPointer(&i)
			},
			initial: &i,
		},
		{
			desc: "New/nil",
			newAtomic: func() *Pointer[foo] {
				return NewPointer[foo](nil)
			},
			initial: nil,
		},
		{
			desc: "zero value",
			newAtomic: func() *Pointer[foo] {
				var p Pointer[foo]
				return &p
			},
			initial: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Run("Load", func(t *testing.T) {
				atom := tt.newAtomic()
				require.Equal(t, tt.initial, atom.Load(), "Load should report nil.")
			})

			t.Run("Swap", func(t *testing.T) {
				atom := tt.newAtomic()
				require.Equal(t, tt.initial, atom.Swap(&k), "Swap didn't return the old value.")
				require.Equal(t, &k, atom.Load(), "Swap didn't set the correct value.")
			})

			t.Run("CAS", func(t *testing.T) {
				atom := tt.newAtomic()
				require.True(t, atom.CompareAndSwap(tt.initial, &j), "CAS didn't report a swap.")
				require.Equal(t, &j, atom.Load(), "CAS didn't set the correct value.")
			})

			t.Run("Store", func(t *testing.T) {
				atom := tt.newAtomic()
				atom.Store(&i)
				require.Equal(t, &i, atom.Load(), "Store didn't set the correct value.")
			})
		})
	}
}
