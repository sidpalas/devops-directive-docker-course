package unstable

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_AST_Numbers(t *testing.T) {
	examples := []struct {
		desc  string
		input string
		kind  Kind
		err   bool
	}{
		{
			desc:  "integer just digits",
			input: `1234`,
			kind:  Integer,
		},
		{
			desc:  "integer zero",
			input: `0`,
			kind:  Integer,
		},
		{
			desc:  "integer sign",
			input: `+99`,
			kind:  Integer,
		},
		{
			desc:  "integer hex uppercase",
			input: `0xDEADBEEF`,
			kind:  Integer,
		},
		{
			desc:  "integer hex lowercase",
			input: `0xdead_beef`,
			kind:  Integer,
		},
		{
			desc:  "integer octal",
			input: `0o01234567`,
			kind:  Integer,
		},
		{
			desc:  "integer binary",
			input: `0b11010110`,
			kind:  Integer,
		},
		{
			desc:  "float zero",
			input: `0.0`,
			kind:  Float,
		},
		{
			desc:  "float positive zero",
			input: `+0.0`,
			kind:  Float,
		},
		{
			desc:  "float negative zero",
			input: `-0.0`,
			kind:  Float,
		},
		{
			desc:  "float pi",
			input: `3.1415`,
			kind:  Float,
		},
		{
			desc:  "float negative",
			input: `-0.01`,
			kind:  Float,
		},
		{
			desc:  "float signed exponent",
			input: `5e+22`,
			kind:  Float,
		},
		{
			desc:  "float exponent lowercase",
			input: `1e06`,
			kind:  Float,
		},
		{
			desc:  "float exponent uppercase",
			input: `-2E-2`,
			kind:  Float,
		},
		{
			desc:  "float fractional with exponent",
			input: `6.626e-34`,
			kind:  Float,
		},
		{
			desc:  "float underscores",
			input: `224_617.445_991_228`,
			kind:  Float,
		},
		{
			desc:  "inf",
			input: `inf`,
			kind:  Float,
		},
		{
			desc:  "inf negative",
			input: `-inf`,
			kind:  Float,
		},
		{
			desc:  "inf positive",
			input: `+inf`,
			kind:  Float,
		},
		{
			desc:  "nan",
			input: `nan`,
			kind:  Float,
		},
		{
			desc:  "nan negative",
			input: `-nan`,
			kind:  Float,
		},
		{
			desc:  "nan positive",
			input: `+nan`,
			kind:  Float,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			p := Parser{}
			p.Reset([]byte(`A = ` + e.input))
			p.NextExpression()
			err := p.Error()
			if e.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				expected := astNode{
					Kind: KeyValue,
					Children: []astNode{
						{Kind: e.kind, Data: []byte(e.input)},
						{Kind: Key, Data: []byte(`A`)},
					},
				}
				compareNode(t, expected, p.Expression())
			}
		})
	}
}

type (
	astNode struct {
		Kind     Kind
		Data     []byte
		Children []astNode
	}
)

func compareNode(t *testing.T, e astNode, n *Node) {
	t.Helper()
	require.Equal(t, e.Kind, n.Kind)
	require.Equal(t, e.Data, n.Data)

	compareIterator(t, e.Children, n.Children())
}

func compareIterator(t *testing.T, expected []astNode, actual Iterator) {
	t.Helper()
	idx := 0

	for actual.Next() {
		n := actual.Node()

		if idx >= len(expected) {
			t.Fatal("extra child in actual tree")
		}
		e := expected[idx]

		compareNode(t, e, n)

		idx++
	}

	if idx < len(expected) {
		t.Fatal("missing children in actual", "idx =", idx, "expected =", len(expected))
	}
}

//nolint:funlen
func TestParser_AST(t *testing.T) {
	examples := []struct {
		desc  string
		input string
		ast   astNode
		err   bool
	}{
		{
			desc:  "simple string assignment",
			input: `A = "hello"`,
			ast: astNode{
				Kind: KeyValue,
				Children: []astNode{
					{
						Kind: String,
						Data: []byte(`hello`),
					},
					{
						Kind: Key,
						Data: []byte(`A`),
					},
				},
			},
		},
		{
			desc:  "simple bool assignment",
			input: `A = true`,
			ast: astNode{
				Kind: KeyValue,
				Children: []astNode{
					{
						Kind: Bool,
						Data: []byte(`true`),
					},
					{
						Kind: Key,
						Data: []byte(`A`),
					},
				},
			},
		},
		{
			desc:  "array of strings",
			input: `A = ["hello", ["world", "again"]]`,
			ast: astNode{
				Kind: KeyValue,
				Children: []astNode{
					{
						Kind: Array,
						Children: []astNode{
							{
								Kind: String,
								Data: []byte(`hello`),
							},
							{
								Kind: Array,
								Children: []astNode{
									{
										Kind: String,
										Data: []byte(`world`),
									},
									{
										Kind: String,
										Data: []byte(`again`),
									},
								},
							},
						},
					},
					{
						Kind: Key,
						Data: []byte(`A`),
					},
				},
			},
		},
		{
			desc:  "array of arrays of strings",
			input: `A = ["hello", "world"]`,
			ast: astNode{
				Kind: KeyValue,
				Children: []astNode{
					{
						Kind: Array,
						Children: []astNode{
							{
								Kind: String,
								Data: []byte(`hello`),
							},
							{
								Kind: String,
								Data: []byte(`world`),
							},
						},
					},
					{
						Kind: Key,
						Data: []byte(`A`),
					},
				},
			},
		},
		{
			desc:  "inline table",
			input: `name = { first = "Tom", last = "Preston-Werner" }`,
			ast: astNode{
				Kind: KeyValue,
				Children: []astNode{
					{
						Kind: InlineTable,
						Children: []astNode{
							{
								Kind: KeyValue,
								Children: []astNode{
									{Kind: String, Data: []byte(`Tom`)},
									{Kind: Key, Data: []byte(`first`)},
								},
							},
							{
								Kind: KeyValue,
								Children: []astNode{
									{Kind: String, Data: []byte(`Preston-Werner`)},
									{Kind: Key, Data: []byte(`last`)},
								},
							},
						},
					},
					{
						Kind: Key,
						Data: []byte(`name`),
					},
				},
			},
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			p := Parser{}
			p.Reset([]byte(e.input))
			p.NextExpression()
			err := p.Error()
			if e.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				compareNode(t, e.ast, p.Expression())
			}
		})
	}
}

func BenchmarkParseBasicStringWithUnicode(b *testing.B) {
	p := &Parser{}
	b.Run("4", func(b *testing.B) {
		input := []byte(`"\u1234\u5678\u9ABC\u1234\u5678\u9ABC"`)
		b.ReportAllocs()
		b.SetBytes(int64(len(input)))

		for i := 0; i < b.N; i++ {
			p.parseBasicString(input)
		}
	})
	b.Run("8", func(b *testing.B) {
		input := []byte(`"\u12345678\u9ABCDEF0\u12345678\u9ABCDEF0"`)
		b.ReportAllocs()
		b.SetBytes(int64(len(input)))

		for i := 0; i < b.N; i++ {
			p.parseBasicString(input)
		}
	})
}

func BenchmarkParseBasicStringsEasy(b *testing.B) {
	p := &Parser{}

	for _, size := range []int{1, 4, 8, 16, 21} {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			input := []byte(`"` + strings.Repeat("A", size) + `"`)

			b.ReportAllocs()
			b.SetBytes(int64(len(input)))

			for i := 0; i < b.N; i++ {
				p.parseBasicString(input)
			}
		})
	}
}

func TestParser_AST_DateTimes(t *testing.T) {
	examples := []struct {
		desc  string
		input string
		kind  Kind
		err   bool
	}{
		{
			desc:  "offset-date-time with delim 'T' and UTC offset",
			input: `2021-07-21T12:08:05Z`,
			kind:  DateTime,
		},
		{
			desc:  "offset-date-time with space delim and +8hours offset",
			input: `2021-07-21 12:08:05+08:00`,
			kind:  DateTime,
		},
		{
			desc:  "local-date-time with nano second",
			input: `2021-07-21T12:08:05.666666666`,
			kind:  LocalDateTime,
		},
		{
			desc:  "local-date-time",
			input: `2021-07-21T12:08:05`,
			kind:  LocalDateTime,
		},
		{
			desc:  "local-date",
			input: `2021-07-21`,
			kind:  LocalDate,
		},
	}

	for _, e := range examples {
		e := e
		t.Run(e.desc, func(t *testing.T) {
			p := Parser{}
			p.Reset([]byte(`A = ` + e.input))
			p.NextExpression()
			err := p.Error()
			if e.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				expected := astNode{
					Kind: KeyValue,
					Children: []astNode{
						{Kind: e.kind, Data: []byte(e.input)},
						{Kind: Key, Data: []byte(`A`)},
					},
				}
				compareNode(t, expected, p.Expression())
			}
		})
	}
}

func ExampleParser() {
	doc := `
	hello = "world"
	value = 42
	`
	p := Parser{}
	p.Reset([]byte(doc))
	for p.NextExpression() {
		e := p.Expression()
		fmt.Printf("Expression: %s\n", e.Kind)
		value := e.Value()
		it := e.Key()
		k := it.Node() // shortcut: we know there is no dotted key in the example
		fmt.Printf("%s -> (%s) %s\n", k.Data, value.Kind, value.Data)
	}

	// Output:
	// Expression: KeyValue
	// hello -> (String) world
	// Expression: KeyValue
	// value -> (Integer) 42
}
