package toml_test

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"
)

func TestFastSimple(t *testing.T) {
	m := map[string]int64{}
	err := toml.Unmarshal([]byte(`a = 42`), &m)
	require.NoError(t, err)
	require.Equal(t, map[string]int64{"a": 42}, m)
}

func TestFastSimpleString(t *testing.T) {
	m := map[string]string{}
	err := toml.Unmarshal([]byte(`a = "hello"`), &m)
	require.NoError(t, err)
	require.Equal(t, map[string]string{"a": "hello"}, m)
}

func TestFastSimpleInterface(t *testing.T) {
	m := map[string]interface{}{}
	err := toml.Unmarshal([]byte(`
	a = "hello"
	b = 42`), &m)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{
		"a": "hello",
		"b": int64(42),
	}, m)
}

func TestFastMultipartKeyInterface(t *testing.T) {
	m := map[string]interface{}{}
	err := toml.Unmarshal([]byte(`
	a.interim = "test"
	a.b.c = "hello"
	b = 42`), &m)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{
		"a": map[string]interface{}{
			"interim": "test",
			"b": map[string]interface{}{
				"c": "hello",
			},
		},
		"b": int64(42),
	}, m)
}

func TestFastExistingMap(t *testing.T) {
	m := map[string]interface{}{
		"ints": map[string]int{},
	}
	err := toml.Unmarshal([]byte(`
	ints.one = 1
	ints.two = 2
	strings.yo = "hello"`), &m)
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{
		"ints": map[string]interface{}{
			"one": int64(1),
			"two": int64(2),
		},
		"strings": map[string]interface{}{
			"yo": "hello",
		},
	}, m)
}

func TestFastArrayTable(t *testing.T) {
	b := []byte(`
	[root]
	[[root.nested]]
	name = 'Bob'
	[[root.nested]]
	name = 'Alice'
	`)

	m := map[string]interface{}{}

	err := toml.Unmarshal(b, &m)
	require.NoError(t, err)

	require.Equal(t, map[string]interface{}{
		"root": map[string]interface{}{
			"nested": []interface{}{
				map[string]interface{}{
					"name": "Bob",
				},
				map[string]interface{}{
					"name": "Alice",
				},
			},
		},
	}, m)
}
