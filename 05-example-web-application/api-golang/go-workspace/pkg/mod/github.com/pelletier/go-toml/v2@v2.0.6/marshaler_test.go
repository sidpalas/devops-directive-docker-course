package toml_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	someInt := 42

	type structInline struct {
		A interface{} `toml:",inline"`
	}

	type comments struct {
		One   int
		Two   int   `comment:"Before kv"`
		Three []int `comment:"Before array"`
	}

	examples := []struct {
		desc     string
		v        interface{}
		expected string
		err      bool
	}{
		{
			desc: "simple map and string",
			v: map[string]string{
				"hello": "world",
			},
			expected: "hello = 'world'\n",
		},
		{
			desc: "map with new line in key",
			v: map[string]string{
				"hel\nlo": "world",
			},
			expected: "\"hel\\nlo\" = 'world'\n",
		},
		{
			desc: `map with " in key`,
			v: map[string]string{
				`hel"lo`: "world",
			},
			expected: "'hel\"lo' = 'world'\n",
		},
		{
			desc: "map in map and string",
			v: map[string]map[string]string{
				"table": {
					"hello": "world",
				},
			},
			expected: `[table]
hello = 'world'
`,
		},
		{
			desc: "map in map in map and string",
			v: map[string]map[string]map[string]string{
				"this": {
					"is": {
						"a": "test",
					},
				},
			},
			expected: `[this]
[this.is]
a = 'test'
`,
		},
		{
			desc: "map in map in map and string with values",
			v: map[string]interface{}{
				"this": map[string]interface{}{
					"is": map[string]string{
						"a": "test",
					},
					"also": "that",
				},
			},
			expected: `[this]
also = 'that'

[this.is]
a = 'test'
`,
		},
		{
			desc: "simple string array",
			v: map[string][]string{
				"array": {"one", "two", "three"},
			},
			expected: `array = ['one', 'two', 'three']
`,
		},
		{
			desc:     "empty string array",
			v:        map[string][]string{},
			expected: ``,
		},
		{
			desc:     "map",
			v:        map[string][]string{},
			expected: ``,
		},
		{
			desc: "nested string arrays",
			v: map[string][][]string{
				"array": {{"one", "two"}, {"three"}},
			},
			expected: `array = [['one', 'two'], ['three']]
`,
		},
		{
			desc: "mixed strings and nested string arrays",
			v: map[string][]interface{}{
				"array": {"a string", []string{"one", "two"}, "last"},
			},
			expected: `array = ['a string', ['one', 'two'], 'last']
`,
		},
		{
			desc: "array of maps",
			v: map[string][]map[string]string{
				"top": {
					{"map1.1": "v1.1"},
					{"map2.1": "v2.1"},
				},
			},
			expected: `[[top]]
'map1.1' = 'v1.1'

[[top]]
'map2.1' = 'v2.1'
`,
		},
		{
			desc: "map with two keys",
			v: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expected: `key1 = 'value1'
key2 = 'value2'
`,
		},
		{
			desc: "simple struct",
			v: struct {
				A string
			}{
				A: "foo",
			},
			expected: `A = 'foo'
`,
		},
		{
			desc: "one level of structs within structs",
			v: struct {
				A interface{}
			}{
				A: struct {
					K1 string
					K2 string
				}{
					K1: "v1",
					K2: "v2",
				},
			},
			expected: `[A]
K1 = 'v1'
K2 = 'v2'
`,
		},
		{
			desc: "structs in array with interfaces",
			v: map[string]interface{}{
				"root": map[string]interface{}{
					"nested": []interface{}{
						map[string]interface{}{"name": "Bob"},
						map[string]interface{}{"name": "Alice"},
					},
				},
			},
			expected: `[root]
[[root.nested]]
name = 'Bob'

[[root.nested]]
name = 'Alice'
`,
		},
		{
			desc: "string escapes",
			v: map[string]interface{}{
				"a": "'\b\f\r\t\"\\",
			},
			expected: `a = "'\b\f\r\t\"\\"
`,
		},
		{
			desc: "string utf8 low",
			v: map[string]interface{}{
				"a": "'Ę",
			},
			expected: `a = "'Ę"
`,
		},
		{
			desc: "string utf8 low 2",
			v: map[string]interface{}{
				"a": "'\u10A85",
			},
			expected: "a = \"'\u10A85\"\n",
		},
		{
			desc: "string utf8 low 2",
			v: map[string]interface{}{
				"a": "'\u10A85",
			},
			expected: "a = \"'\u10A85\"\n",
		},
		{
			desc: "emoji",
			v: map[string]interface{}{
				"a": "'😀",
			},
			expected: "a = \"'😀\"\n",
		},
		{
			desc: "control char",
			v: map[string]interface{}{
				"a": "'\u001A",
			},
			expected: `a = "'\u001A"
`,
		},
		{
			desc: "multi-line string",
			v: map[string]interface{}{
				"a": "hello\nworld",
			},
			expected: `a = "hello\nworld"
`,
		},
		{
			desc: "multi-line forced",
			v: struct {
				A string `toml:",multiline"`
			}{
				A: "hello\nworld",
			},
			expected: `A = """
hello
world"""
`,
		},
		{
			desc: "inline field",
			v: struct {
				A map[string]string `toml:",inline"`
				B map[string]string
			}{
				A: map[string]string{
					"isinline": "yes",
				},
				B: map[string]string{
					"isinline": "no",
				},
			},
			expected: `A = {isinline = 'yes'}

[B]
isinline = 'no'
`,
		},
		{
			desc: "mutiline array int",
			v: struct {
				A []int `toml:",multiline"`
				B []int
			}{
				A: []int{1, 2, 3, 4},
				B: []int{1, 2, 3, 4},
			},
			expected: `A = [
  1,
  2,
  3,
  4
]
B = [1, 2, 3, 4]
`,
		},
		{
			desc: "mutiline array in array",
			v: struct {
				A [][]int `toml:",multiline"`
			}{
				A: [][]int{{1, 2}, {3, 4}},
			},
			expected: `A = [
  [1, 2],
  [3, 4]
]
`,
		},
		{
			desc: "nil interface not supported at root",
			v:    nil,
			err:  true,
		},
		{
			desc: "nil interface not supported in slice",
			v: map[string]interface{}{
				"a": []interface{}{"a", nil, 2},
			},
			err: true,
		},
		{
			desc: "nil pointer in slice uses zero value",
			v: struct {
				A []*int
			}{
				A: []*int{nil},
			},
			expected: `A = [0]
`,
		},
		{
			desc: "nil pointer in slice uses zero value",
			v: struct {
				A []*int
			}{
				A: []*int{nil},
			},
			expected: `A = [0]
`,
		},
		{
			desc: "pointer in slice",
			v: struct {
				A []*int
			}{
				A: []*int{&someInt},
			},
			expected: `A = [42]
`,
		},
		{
			desc: "inline table in inline table",
			v: structInline{
				A: structInline{
					A: structInline{
						A: "hello",
					},
				},
			},
			expected: `A = {A = {A = 'hello'}}
`,
		},
		{
			desc: "empty slice in map",
			v: map[string][]string{
				"a": {},
			},
			expected: `a = []
`,
		},
		{
			desc: "map in slice",
			v: map[string][]map[string]string{
				"a": {{"hello": "world"}},
			},
			expected: `[[a]]
hello = 'world'
`,
		},
		{
			desc: "newline in map in slice",
			v: map[string][]map[string]string{
				"a\n": {{"hello": "world"}},
			},
			expected: `[["a\n"]]
hello = 'world'
`,
		},
		{
			desc: "newline in map in slice",
			v: map[string][]map[string]*customTextMarshaler{
				"a": {{"hello": &customTextMarshaler{1}}},
			},
			err: true,
		},
		{
			desc: "empty slice of empty struct",
			v: struct {
				A []struct{}
			}{
				A: []struct{}{},
			},
			expected: `A = []
`,
		},
		{
			desc: "nil field is ignored",
			v: struct {
				A interface{}
			}{
				A: nil,
			},
			expected: ``,
		},
		{
			desc: "private fields are ignored",
			v: struct {
				Public  string
				private string
			}{
				Public:  "shown",
				private: "hidden",
			},
			expected: `Public = 'shown'
`,
		},
		{
			desc: "fields tagged - are ignored",
			v: struct {
				Public  string `toml:"-"`
				private string
			}{
				Public: "hidden",
			},
			expected: ``,
		},
		{
			desc: "nil value in map is ignored",
			v: map[string]interface{}{
				"A": nil,
			},
			expected: ``,
		},
		{
			desc: "new line in table key",
			v: map[string]interface{}{
				"hello\nworld": 42,
			},
			expected: `"hello\nworld" = 42
`,
		},
		{
			desc: "new line in parent of nested table key",
			v: map[string]interface{}{
				"hello\nworld": map[string]interface{}{
					"inner": 42,
				},
			},
			expected: `["hello\nworld"]
inner = 42
`,
		},
		{
			desc: "new line in nested table key",
			v: map[string]interface{}{
				"parent": map[string]interface{}{
					"in\ner": map[string]interface{}{
						"foo": 42,
					},
				},
			},
			expected: `[parent]
[parent."in\ner"]
foo = 42
`,
		},
		{
			desc: "invalid map key",
			v:    map[int]interface{}{},
			err:  true,
		},
		{
			desc: "unhandled type",
			v: struct {
				A chan int
			}{
				A: make(chan int),
			},
			err: true,
		},
		{
			desc: "time",
			v: struct {
				T time.Time
			}{
				T: time.Time{},
			},
			expected: `T = 0001-01-01T00:00:00Z
`,
		},
		{
			desc: "time nano",
			v: struct {
				T time.Time
			}{
				T: time.Date(1979, time.May, 27, 0, 32, 0, 999999000, time.UTC),
			},
			expected: `T = 1979-05-27T00:32:00.999999Z
`,
		},
		{
			desc: "bool",
			v: struct {
				A bool
				B bool
			}{
				A: false,
				B: true,
			},
			expected: `A = false
B = true
`,
		},
		{
			desc: "numbers",
			v: struct {
				A float32
				B uint64
				C uint32
				D uint16
				E uint8
				F uint
				G int64
				H int32
				I int16
				J int8
				K int
				L float64
			}{
				A: 1.1,
				B: 42,
				C: 42,
				D: 42,
				E: 42,
				F: 42,
				G: 42,
				H: 42,
				I: 42,
				J: 42,
				K: 42,
				L: 2.2,
			},
			expected: `A = 1.1
B = 42
C = 42
D = 42
E = 42
F = 42
G = 42
H = 42
I = 42
J = 42
K = 42
L = 2.2
`,
		},
		{
			desc: "comments",
			v: struct {
				Table comments `comment:"Before table"`
			}{
				Table: comments{
					One:   1,
					Two:   2,
					Three: []int{1, 2, 3},
				},
			},
			expected: `# Before table
[Table]
One = 1
# Before kv
Two = 2
# Before array
Three = [1, 2, 3]
`,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			b, err := toml.Marshal(e.v)
			if e.err {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, e.expected, string(b))

			// make sure the output is always valid TOML
			defaultMap := map[string]interface{}{}
			err = toml.Unmarshal(b, &defaultMap)
			require.NoError(t, err)

			testWithAllFlags(t, func(t *testing.T, flags int) {
				t.Helper()

				var buf bytes.Buffer
				enc := toml.NewEncoder(&buf)
				setFlags(enc, flags)

				err := enc.Encode(e.v)
				require.NoError(t, err)

				inlineMap := map[string]interface{}{}
				err = toml.Unmarshal(buf.Bytes(), &inlineMap)
				require.NoError(t, err)

				require.Equal(t, defaultMap, inlineMap)
			})
		})
	}
}

type flagsSetters []struct {
	name string
	f    func(enc *toml.Encoder, flag bool) *toml.Encoder
}

var allFlags = flagsSetters{
	{"arrays-multiline", (*toml.Encoder).SetArraysMultiline},
	{"tables-inline", (*toml.Encoder).SetTablesInline},
	{"indent-tables", (*toml.Encoder).SetIndentTables},
}

func setFlags(enc *toml.Encoder, flags int) {
	for i := 0; i < len(allFlags); i++ {
		enabled := flags&1 > 0
		allFlags[i].f(enc, enabled)
	}
}

func testWithAllFlags(t *testing.T, testfn func(t *testing.T, flags int)) {
	t.Helper()
	testWithFlags(t, 0, allFlags, testfn)
}

func testWithFlags(t *testing.T, flags int, setters flagsSetters, testfn func(t *testing.T, flags int)) {
	t.Helper()

	if len(setters) == 0 {
		testfn(t, flags)

		return
	}

	s := setters[0]

	for _, enabled := range []bool{false, true} {
		name := fmt.Sprintf("%s=%t", s.name, enabled)
		newFlags := flags << 1

		if enabled {
			newFlags++
		}

		t.Run(name, func(t *testing.T) {
			testWithFlags(t, newFlags, setters[1:], testfn)
		})
	}
}

func TestMarshalFloats(t *testing.T) {
	v := map[string]float32{
		"nan":  float32(math.NaN()),
		"+inf": float32(math.Inf(1)),
		"-inf": float32(math.Inf(-1)),
	}

	expected := `'+inf' = inf
-inf = -inf
nan = nan
`

	actual, err := toml.Marshal(v)
	require.NoError(t, err)
	require.Equal(t, expected, string(actual))

	v64 := map[string]float64{
		"nan":  math.NaN(),
		"+inf": math.Inf(1),
		"-inf": math.Inf(-1),
	}

	actual, err = toml.Marshal(v64)
	require.NoError(t, err)
	require.Equal(t, expected, string(actual))
}

//nolint:funlen
func TestMarshalIndentTables(t *testing.T) {
	examples := []struct {
		desc     string
		v        interface{}
		expected string
	}{
		{
			desc: "one kv",
			v: map[string]interface{}{
				"foo": "bar",
			},
			expected: `foo = 'bar'
`,
		},
		{
			desc: "one level table",
			v: map[string]map[string]string{
				"foo": {
					"one": "value1",
					"two": "value2",
				},
			},
			expected: `[foo]
  one = 'value1'
  two = 'value2'
`,
		},
		{
			desc: "two levels table",
			v: map[string]interface{}{
				"root": "value0",
				"level1": map[string]interface{}{
					"one": "value1",
					"level2": map[string]interface{}{
						"two": "value2",
					},
				},
			},
			expected: `root = 'value0'

[level1]
  one = 'value1'

  [level1.level2]
    two = 'value2'
`,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			var buf strings.Builder
			enc := toml.NewEncoder(&buf)
			enc.SetIndentTables(true)
			err := enc.Encode(e.v)
			require.NoError(t, err)
			assert.Equal(t, e.expected, buf.String())
		})
	}
}

type customTextMarshaler struct {
	value int64
}

func (c *customTextMarshaler) MarshalText() ([]byte, error) {
	if c.value == 1 {
		return nil, fmt.Errorf("cannot represent 1 because this is a silly test")
	}
	return []byte(fmt.Sprintf("::%d", c.value)), nil
}

func TestMarshalTextMarshaler_NoRoot(t *testing.T) {
	c := customTextMarshaler{}
	_, err := toml.Marshal(&c)
	require.Error(t, err)
}

func TestMarshalTextMarshaler_Error(t *testing.T) {
	m := map[string]interface{}{"a": &customTextMarshaler{value: 1}}
	_, err := toml.Marshal(m)
	require.Error(t, err)
}

func TestMarshalTextMarshaler_ErrorInline(t *testing.T) {
	type s struct {
		A map[string]interface{} `inline:"true"`
	}

	d := s{
		A: map[string]interface{}{"a": &customTextMarshaler{value: 1}},
	}

	_, err := toml.Marshal(d)
	require.Error(t, err)
}

func TestMarshalTextMarshaler(t *testing.T) {
	m := map[string]interface{}{"a": &customTextMarshaler{value: 2}}
	r, err := toml.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, "a = '::2'\n", string(r))
}

type brokenWriter struct{}

func (b *brokenWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("dead")
}

func TestEncodeToBrokenWriter(t *testing.T) {
	w := brokenWriter{}
	enc := toml.NewEncoder(&w)
	err := enc.Encode(map[string]string{"hello": "world"})
	require.Error(t, err)
}

func TestEncoderSetIndentSymbol(t *testing.T) {
	var w strings.Builder
	enc := toml.NewEncoder(&w)
	enc.SetIndentTables(true)
	enc.SetIndentSymbol(">>>")
	err := enc.Encode(map[string]map[string]string{"parent": {"hello": "world"}})
	require.NoError(t, err)
	expected := `[parent]
>>>hello = 'world'
`
	assert.Equal(t, expected, w.String())
}

func TestEncoderOmitempty(t *testing.T) {
	type doc struct {
		String  string            `toml:",omitempty,multiline"`
		Bool    bool              `toml:",omitempty,multiline"`
		Int     int               `toml:",omitempty,multiline"`
		Int8    int8              `toml:",omitempty,multiline"`
		Int16   int16             `toml:",omitempty,multiline"`
		Int32   int32             `toml:",omitempty,multiline"`
		Int64   int64             `toml:",omitempty,multiline"`
		Uint    uint              `toml:",omitempty,multiline"`
		Uint8   uint8             `toml:",omitempty,multiline"`
		Uint16  uint16            `toml:",omitempty,multiline"`
		Uint32  uint32            `toml:",omitempty,multiline"`
		Uint64  uint64            `toml:",omitempty,multiline"`
		Float32 float32           `toml:",omitempty,multiline"`
		Float64 float64           `toml:",omitempty,multiline"`
		MapNil  map[string]string `toml:",omitempty,multiline"`
		Slice   []string          `toml:",omitempty,multiline"`
		Ptr     *string           `toml:",omitempty,multiline"`
		Iface   interface{}       `toml:",omitempty,multiline"`
		Struct  struct{}          `toml:",omitempty,multiline"`
	}

	d := doc{}

	b, err := toml.Marshal(d)
	require.NoError(t, err)

	expected := ``

	assert.Equal(t, expected, string(b))
}

func TestEncoderTagFieldName(t *testing.T) {
	type doc struct {
		String string `toml:"hello"`
		OkSym  string `toml:"#"`
		Bad    string `toml:"\"`
	}

	d := doc{String: "world"}

	b, err := toml.Marshal(d)
	require.NoError(t, err)

	expected := `hello = 'world'
'#' = ''
Bad = ''
`

	assert.Equal(t, expected, string(b))
}

func TestIssue436(t *testing.T) {
	data := []byte(`{"a": [ { "b": { "c": "d" } } ]}`)

	var v interface{}
	err := json.Unmarshal(data, &v)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = toml.NewEncoder(&buf).Encode(v)
	require.NoError(t, err)

	expected := `[[a]]
[a.b]
c = 'd'
`
	assert.Equal(t, expected, buf.String())
}

func TestIssue424(t *testing.T) {
	type Message1 struct {
		Text string
	}

	type Message2 struct {
		Text string `multiline:"true"`
	}

	msg1 := Message1{"Hello\\World"}
	msg2 := Message2{"Hello\\World"}

	toml1, err := toml.Marshal(msg1)
	require.NoError(t, err)

	toml2, err := toml.Marshal(msg2)
	require.NoError(t, err)

	msg1parsed := Message1{}
	err = toml.Unmarshal(toml1, &msg1parsed)
	require.NoError(t, err)
	require.Equal(t, msg1, msg1parsed)

	msg2parsed := Message2{}
	err = toml.Unmarshal(toml2, &msg2parsed)
	require.NoError(t, err)
	require.Equal(t, msg2, msg2parsed)
}

func TestIssue567(t *testing.T) {
	var m map[string]interface{}
	err := toml.Unmarshal([]byte("A = 12:08:05"), &m)
	require.NoError(t, err)
	require.IsType(t, m["A"], toml.LocalTime{})
}

func TestIssue590(t *testing.T) {
	type CustomType int
	var cfg struct {
		Option CustomType `toml:"option"`
	}
	err := toml.Unmarshal([]byte("option = 42"), &cfg)
	require.NoError(t, err)
}

func TestIssue571(t *testing.T) {
	type Foo struct {
		Float32 float32
		Float64 float64
	}

	const closeEnough = 1e-9

	foo := Foo{
		Float32: 42,
		Float64: 43,
	}
	b, err := toml.Marshal(foo)
	require.NoError(t, err)

	var foo2 Foo
	err = toml.Unmarshal(b, &foo2)
	require.NoError(t, err)

	assert.InDelta(t, 42, foo2.Float32, closeEnough)
	assert.InDelta(t, 43, foo2.Float64, closeEnough)
}

func TestIssue678(t *testing.T) {
	type Config struct {
		BigInt big.Int
	}

	cfg := &Config{
		BigInt: *big.NewInt(123),
	}

	out, err := toml.Marshal(cfg)
	require.NoError(t, err)
	assert.Equal(t, "BigInt = '123'\n", string(out))

	cfg2 := &Config{}
	err = toml.Unmarshal(out, cfg2)
	require.NoError(t, err)
	require.Equal(t, cfg, cfg2)
}

func TestIssue752(t *testing.T) {
	type Fooer interface {
		Foo() string
	}

	type Container struct {
		Fooer
	}

	c := Container{}

	out, err := toml.Marshal(c)
	require.NoError(t, err)
	require.Equal(t, "", string(out))
}

func TestIssue768(t *testing.T) {
	type cfg struct {
		Name string `comment:"This is a multiline comment.\nThis is line 2."`
	}

	out, err := toml.Marshal(&cfg{})
	require.NoError(t, err)

	expected := `# This is a multiline comment.
# This is line 2.
Name = ''
`

	require.Equal(t, expected, string(out))
}

func TestIssue786(t *testing.T) {
	type Dependencies struct {
		Dependencies         []string `toml:"dependencies,multiline,omitempty"`
		BuildDependencies    []string `toml:"buildDependencies,multiline,omitempty"`
		OptionalDependencies []string `toml:"optionalDependencies,multiline,omitempty"`
	}

	type Test struct {
		Dependencies Dependencies `toml:"dependencies,omitempty"`
	}

	x := Test{}
	b, err := toml.Marshal(x)
	require.NoError(t, err)

	require.Equal(t, "", string(b))

	type General struct {
		From      string `toml:"from,omitempty" json:"from,omitempty" comment:"from in graphite-web format, the local TZ is used"`
		Randomize bool   `toml:"randomize" json:"randomize" comment:"randomize starting time with [0,step)"`
	}

	type Custom struct {
		Name string `toml:"name" json:"name,omitempty" comment:"names for generator, braces are expanded like in shell"`
		Type string `toml:"type,omitempty" json:"type,omitempty" comment:"type of generator"`
		General
	}
	type Config struct {
		General
		Custom []Custom `toml:"custom,omitempty" json:"custom,omitempty" comment:"generators with custom parameters can be specified separately"`
	}

	buf := new(bytes.Buffer)
	config := &Config{General: General{From: "-2d", Randomize: true}}
	config.Custom = []Custom{{Name: "omit", General: General{Randomize: false}}}
	config.Custom = append(config.Custom, Custom{Name: "present", General: General{From: "-2d", Randomize: true}})
	encoder := toml.NewEncoder(buf)
	encoder.Encode(config)

	expected := `# from in graphite-web format, the local TZ is used
from = '-2d'
# randomize starting time with [0,step)
randomize = true

# generators with custom parameters can be specified separately
[[custom]]
# names for generator, braces are expanded like in shell
name = 'omit'
# randomize starting time with [0,step)
randomize = false

[[custom]]
# names for generator, braces are expanded like in shell
name = 'present'
# from in graphite-web format, the local TZ is used
from = '-2d'
# randomize starting time with [0,step)
randomize = true
`

	require.Equal(t, expected, buf.String())
}

func TestMarshalNestedAnonymousStructs(t *testing.T) {
	type Embedded struct {
		Value string `toml:"value" json:"value"`
		Top   struct {
			Value string `toml:"value" json:"value"`
		} `toml:"top" json:"top"`
	}

	type Named struct {
		Value string `toml:"value" json:"value"`
	}

	var doc struct {
		Embedded
		Named     `toml:"named" json:"named"`
		Anonymous struct {
			Value string `toml:"value" json:"value"`
		} `toml:"anonymous" json:"anonymous"`
	}

	expected := `value = ''

[top]
value = ''

[named]
value = ''

[anonymous]
value = ''
`

	result, err := toml.Marshal(doc)
	require.NoError(t, err)
	require.Equal(t, expected, string(result))
}

func TestMarshalNestedAnonymousStructs_DuplicateField(t *testing.T) {
	type Embedded struct {
		Value string `toml:"value" json:"value"`
		Top   struct {
			Value string `toml:"value" json:"value"`
		} `toml:"top" json:"top"`
	}

	var doc struct {
		Value string `toml:"value" json:"value"`
		Embedded
	}
	doc.Embedded.Value = "shadowed"
	doc.Value = "shadows"

	expected := `value = 'shadows'

[top]
value = ''
`

	result, err := toml.Marshal(doc)
	require.NoError(t, err)
	require.NoError(t, err)
	require.Equal(t, expected, string(result))
}

func TestLocalTime(t *testing.T) {
	v := map[string]toml.LocalTime{
		"a": {
			Hour:       1,
			Minute:     2,
			Second:     3,
			Nanosecond: 4,
		},
	}

	expected := `a = 01:02:03.000000004
`

	out, err := toml.Marshal(v)
	require.NoError(t, err)
	require.Equal(t, expected, string(out))
}

func TestMarshalUint64Overflow(t *testing.T) {
	// The TOML spec only requires implementation to provide support for the
	// int64 range. To avoid generating TOML documents that would not be
	// supported by standard-compliant parsers, uint64 > max int64 cannot be
	// marshaled.
	x := map[string]interface{}{
		"foo": uint64(math.MaxInt64) + 1,
	}

	_, err := toml.Marshal(x)
	require.Error(t, err)
}

func ExampleMarshal() {
	type MyConfig struct {
		Version int
		Name    string
		Tags    []string
	}

	cfg := MyConfig{
		Version: 2,
		Name:    "go-toml",
		Tags:    []string{"go", "toml"},
	}

	b, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// Version = 2
	// Name = 'go-toml'
	// Tags = ['go', 'toml']
}
