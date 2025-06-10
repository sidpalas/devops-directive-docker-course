package urn

import (
	"fmt"
	"testing"
)

var benchs = []testCase{
	genericTestCases[14],
	genericTestCases[2],
	genericTestCases[6],
	genericTestCases[10],
	genericTestCases[11],
	genericTestCases[13],
	genericTestCases[20],
	genericTestCases[23],
	genericTestCases[33],
	genericTestCases[45],
	genericTestCases[47],
	genericTestCases[48],
	genericTestCases[50],
	genericTestCases[52],
	genericTestCases[53],
	genericTestCases[57],
	genericTestCases[62],
	genericTestCases[63],
	genericTestCases[67],
	genericTestCases[60],
}

// This is here to avoid compiler optimizations that
// could remove the actual call we are benchmarking
// during benchmarks
var benchParseResult *URN

func BenchmarkParse(b *testing.B) {
	for ii, tt := range benchs {
		tt := tt
		outcome := (map[bool]string{true: "ok", false: "no"})[tt.ok]
		b.Run(
			fmt.Sprintf("%s/%02d/%s/", outcome, ii, rxpad(string(tt.in), 45)),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					benchParseResult, _ = Parse(tt.in)
				}
			},
		)
	}
}
