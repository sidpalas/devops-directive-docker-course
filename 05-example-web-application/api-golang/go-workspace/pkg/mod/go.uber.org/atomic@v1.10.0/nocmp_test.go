// Copyright (c) 2020 Uber Technologies, Inc.
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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNocmpComparability(t *testing.T) {
	tests := []struct {
		desc       string
		give       interface{}
		comparable bool
	}{
		{
			desc: "nocmp struct",
			give: nocmp{},
		},
		{
			desc: "struct with nocmp embedded",
			give: struct{ nocmp }{},
		},
		{
			desc:       "pointer to struct with nocmp embedded",
			give:       &struct{ nocmp }{},
			comparable: true,
		},

		// All exported types must be uncomparable.
		{desc: "Bool", give: Bool{}},
		{desc: "Duration", give: Duration{}},
		{desc: "Error", give: Error{}},
		{desc: "Float64", give: Float64{}},
		{desc: "Int32", give: Int32{}},
		{desc: "Int64", give: Int64{}},
		{desc: "String", give: String{}},
		{desc: "Uint32", give: Uint32{}},
		{desc: "Uint64", give: Uint64{}},
		{desc: "Value", give: Value{}},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			typ := reflect.TypeOf(tt.give)
			assert.Equalf(t, tt.comparable, typ.Comparable(),
				"type %v comparablity mismatch", typ)
		})
	}
}

// nocmp must not add to the size of a struct in-memory.
func TestNocmpSize(t *testing.T) {
	type x struct{ _ int }

	before := reflect.TypeOf(x{}).Size()

	type y struct {
		_ nocmp
		_ x
	}

	after := reflect.TypeOf(y{}).Size()

	assert.Equal(t, before, after,
		"expected nocmp to have no effect on struct size")
}

// This test will fail to compile if we disallow copying of nocmp.
//
// We need to allow this so that users can do,
//
//	var x atomic.Int32
//	x = atomic.NewInt32(1)
func TestNocmpCopy(t *testing.T) {
	type foo struct{ _ nocmp }

	t.Run("struct copy", func(t *testing.T) {
		a := foo{}
		b := a
		_ = b // unused
	})

	t.Run("pointer copy", func(t *testing.T) {
		a := &foo{}
		b := *a
		_ = b // unused
	})
}

// Fake go.mod with no dependencies.
const _exampleGoMod = `module example.com/nocmp`

const _badFile = `package atomic

import "fmt"

type Int64 struct {
	nocmp

	v int64
}

func shouldNotCompile() {
	var x, y Int64
	fmt.Println(x == y)
}
`

func TestNocmpIntegration(t *testing.T) {
	tempdir := t.TempDir()

	nocmp, err := os.ReadFile("nocmp.go")
	require.NoError(t, err, "unable to read nocmp.go")

	require.NoError(t,
		os.WriteFile(filepath.Join(tempdir, "go.mod"), []byte(_exampleGoMod), 0o644),
		"unable to write go.mod")

	require.NoError(t,
		os.WriteFile(filepath.Join(tempdir, "nocmp.go"), nocmp, 0o644),
		"unable to write nocmp.go")

	require.NoError(t,
		os.WriteFile(filepath.Join(tempdir, "bad.go"), []byte(_badFile), 0o644),
		"unable to write bad.go")

	var stderr bytes.Buffer
	cmd := exec.Command("go", "build")
	cmd.Dir = tempdir
	// Create a minimal build environment with only HOME set so that "go
	// build" has somewhere to put the cache and other Go files in.
	cmd.Env = []string{"HOME=" + filepath.Join(tempdir, "home")}
	cmd.Stderr = &stderr
	require.Error(t, cmd.Run(), "bad.go must not compile")

	assert.Contains(t, stderr.String(),
		"struct containing nocmp cannot be compared")
}
