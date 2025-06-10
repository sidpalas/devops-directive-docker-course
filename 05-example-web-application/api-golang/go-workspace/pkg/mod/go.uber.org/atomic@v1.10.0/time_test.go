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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	start := time.Date(2021, 6, 17, 9, 10, 0, 0, time.UTC)
	atom := NewTime(start)

	require.Equal(t, start, atom.Load(), "Load didn't work")
	require.Equal(t, time.Time{}, NewTime(time.Time{}).Load(), "Default time value is wrong")
}

func TestTimeLocation(t *testing.T) {
	// Check TZ data hasn't been lost from load/store.
	ny, err := time.LoadLocation("America/New_York")
	require.NoError(t, err, "Failed to load location")
	nyTime := NewTime(time.Date(2021, 1, 1, 0, 0, 0, 0, ny))

	var atom Time
	atom.Store(nyTime.Load())

	assert.Equal(t, ny, atom.Load().Location(), "Location information is wrong")
}

func TestLargeTime(t *testing.T) {
	// Check "large/small" time that are beyond int64 ns
	// representation (< year 1678 or > year 2262) can be
	// correctly load/store'd.
	t.Parallel()

	t.Run("future", func(t *testing.T) {
		future := time.Date(2262, 12, 31, 0, 0, 0, 0, time.UTC)
		atom := NewTime(future)
		dayAfterFuture := atom.Load().AddDate(0, 1, 0)

		atom.Store(dayAfterFuture)
		assert.Equal(t, 2263, atom.Load().Year())
	})

	t.Run("past", func(t *testing.T) {
		past := time.Date(1678, 1, 1, 0, 0, 0, 0, time.UTC)
		atom := NewTime(past)
		dayBeforePast := atom.Load().AddDate(0, -1, 0)

		atom.Store(dayBeforePast)
		assert.Equal(t, 1677, atom.Load().Year())
	})
}

func TestMonotonic(t *testing.T) {
	before := NewTime(time.Now())
	time.Sleep(15 * time.Millisecond)
	after := NewTime(time.Now())

	// try loading/storing before and test monotonic clock value hasn't been lost
	bt := before.Load()
	before.Store(bt)
	d := after.Load().Sub(before.Load())
	assert.True(t, 15 <= d.Milliseconds())
}
