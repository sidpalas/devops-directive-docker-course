package urn

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultPrefixWhenString(t *testing.T) {
	u := &URN{
		ID: "pippo",
		SS: "pluto",
	}

	assert.Equal(t, "urn:pippo:pluto", u.String())
}

func TestParseSignature(t *testing.T) {
	urn, ok := Parse([]byte(``))
	assert.Nil(t, urn)
	assert.False(t, ok)
}

func TestLexicalEquivalence(t *testing.T) {
	for ii, tt := range equivalenceTests {
		urnlx, oklx := Parse(tt.lx)
		urnrx, okrx := Parse(tt.rx)

		if oklx && okrx {

			assert.True(t, urnlx.Equal(urnlx))
			assert.True(t, urnrx.Equal(urnrx))

			if tt.eq {
				assert.True(t, urnlx.Equal(urnrx), ierror(ii))
				assert.True(t, urnrx.Equal(urnlx), ierror(ii))
			} else {
				assert.False(t, urnlx.Equal(urnrx), ierror(ii))
				assert.False(t, urnrx.Equal(urnlx), ierror(ii))
			}
		} else {
			t.Log("Something wrong in the testing table ...")
		}
	}
}

func TestJSONMarshaling(t *testing.T) {
	t.Run("roundtrip", func(t *testing.T) {
		// Marshal
		expected := URN{ID: "oid", SS: "1.2.3.4"}
		bytes, err := json.Marshal(expected)
		if !assert.NoError(t, err) {
			return
		}
		// Unmarshal
		var actual URN
		err = json.Unmarshal(bytes, &actual)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expected.String(), actual.String())
	})
	t.Run("invalid URN", func(t *testing.T) {
		var actual URN
		err := json.Unmarshal([]byte(`"not a URN"`), &actual)
		assert.EqualError(t, err, "invalid URN: not a URN")
	})
}