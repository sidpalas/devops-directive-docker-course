// Copyright (c) 2021 Uber Technologies, Inc.
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

package atomic

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestUnsafePointer(t *testing.T) {
	i := int64(42)
	j := int64(0)
	k := int64(1)

	tests := []struct {
		desc      string
		newAtomic func() *UnsafePointer
		initial   unsafe.Pointer
	}{
		{
			desc: "non-empty",
			newAtomic: func() *UnsafePointer {
				return NewUnsafePointer(unsafe.Pointer(&i))
			},
			initial: unsafe.Pointer(&i),
		},
		{
			desc: "nil",
			newAtomic: func() *UnsafePointer {
				var p UnsafePointer
				return &p
			},
			initial: unsafe.Pointer(nil),
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
				require.Equal(t, tt.initial, atom.Swap(unsafe.Pointer(&k)), "Swap didn't return the old value.")
				require.Equal(t, unsafe.Pointer(&k), atom.Load(), "Swap didn't set the correct value.")
			})

			t.Run("CAS", func(t *testing.T) {
				atom := tt.newAtomic()
				require.True(t, atom.CAS(tt.initial, unsafe.Pointer(&j)), "CAS didn't report a swap.")
				require.Equal(t, unsafe.Pointer(&j), atom.Load(), "CAS didn't set the correct value.")
			})

			t.Run("Store", func(t *testing.T) {
				atom := tt.newAtomic()
				atom.Store(unsafe.Pointer(&i))
				require.Equal(t, unsafe.Pointer(&i), atom.Load(), "Store didn't set the correct value.")
			})
		})
	}
}
