package ja

import (
	"testing"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

func TestLocale(t *testing.T) {

	trans := New()
	expected := "ja"

	if trans.Locale() != expected {
		t.Errorf("Expected '%s' Got '%s'", expected, trans.Locale())
	}
}

func TestPluralsRange(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsRange()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestPluralsOrdinal(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsOrdinal()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestPluralsCardinal(t *testing.T) {

	trans := New()

	tests := []struct {
		expected locales.PluralRule
	}{
		{
			expected: locales.PluralRuleOther,
		},
	}

	rules := trans.PluralsCardinal()
	expected := 1
	if len(rules) != expected {
		t.Errorf("Expected '%d' Got '%d'", expected, len(rules))
	}

	for _, tt := range tests {

		r := locales.PluralRuleUnknown

		for i := 0; i < len(rules); i++ {
			if rules[i] == tt.expected {
				r = rules[i]
				break
			}
		}
		if r == locales.PluralRuleUnknown {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, r)
		}
	}
}

func TestRangePlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num1     float64
		v1       uint64
		num2     float64
		v2       uint64
		expected locales.PluralRule
	}{
		{
			num1:     1,
			v1:       1,
			num2:     2,
			v2:       2,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.RangePluralRule(tt.num1, tt.v1, tt.num2, tt.v2)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestOrdinalPlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected locales.PluralRule
	}{
		{
			num:      1,
			v:        0,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.OrdinalPluralRule(tt.num, tt.v)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestCardinalPlurals(t *testing.T) {

	trans := New()

	tests := []struct {
		num      float64
		v        uint64
		expected locales.PluralRule
	}{
		{
			num:      1,
			v:        0,
			expected: locales.PluralRuleOther,
		},
	}

	for _, tt := range tests {
		rule := trans.CardinalPluralRule(tt.num, tt.v)
		if rule != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, rule)
		}
	}
}

func TestDaysAbbreviated(t *testing.T) {

	trans := New()
	days := trans.WeekdaysAbbreviated()

	for i, day := range days {
		s := trans.WeekdayAbbreviated(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "日",
		},
		{
			idx:      1,
			expected: "月",
		},
		{
			idx:      2,
			expected: "火",
		},
		{
			idx:      3,
			expected: "水",
		},
		{
			idx:      4,
			expected: "木",
		},
		{
			idx:      5,
			expected: "金",
		},
		{
			idx:      6,
			expected: "土",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayAbbreviated(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysNarrow(t *testing.T) {

	trans := New()
	days := trans.WeekdaysNarrow()

	for i, day := range days {
		s := trans.WeekdayNarrow(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", string(day), s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "日",
		},
		{
			idx:      1,
			expected: "月",
		},
		{
			idx:      2,
			expected: "火",
		},
		{
			idx:      3,
			expected: "水",
		},
		{
			idx:      4,
			expected: "木",
		},
		{
			idx:      5,
			expected: "金",
		},
		{
			idx:      6,
			expected: "土",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayNarrow(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysShort(t *testing.T) {

	trans := New()
	days := trans.WeekdaysShort()

	for i, day := range days {
		s := trans.WeekdayShort(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "日",
		},
		{
			idx:      1,
			expected: "月",
		},
		{
			idx:      2,
			expected: "火",
		},
		{
			idx:      3,
			expected: "水",
		},
		{
			idx:      4,
			expected: "木",
		},
		{
			idx:      5,
			expected: "金",
		},
		{
			idx:      6,
			expected: "土",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayShort(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestDaysWide(t *testing.T) {

	trans := New()
	days := trans.WeekdaysWide()

	for i, day := range days {
		s := trans.WeekdayWide(time.Weekday(i))
		if s != day {
			t.Errorf("Expected '%s' Got '%s'", day, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      0,
			expected: "日曜日",
		},
		{
			idx:      1,
			expected: "月曜日",
		},
		{
			idx:      2,
			expected: "火曜日",
		},
		{
			idx:      3,
			expected: "水曜日",
		},
		{
			idx:      4,
			expected: "木曜日",
		},
		{
			idx:      5,
			expected: "金曜日",
		},
		{
			idx:      6,
			expected: "土曜日",
		},
	}

	for _, tt := range tests {
		s := trans.WeekdayWide(time.Weekday(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsAbbreviated(t *testing.T) {

	trans := New()
	months := trans.MonthsAbbreviated()

	for i, month := range months {
		s := trans.MonthAbbreviated(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "1月",
		},
		{
			idx:      2,
			expected: "2月",
		},
		{
			idx:      3,
			expected: "3月",
		},
		{
			idx:      4,
			expected: "4月",
		},
		{
			idx:      5,
			expected: "5月",
		},
		{
			idx:      6,
			expected: "6月",
		},
		{
			idx:      7,
			expected: "7月",
		},
		{
			idx:      8,
			expected: "8月",
		},
		{
			idx:      9,
			expected: "9月",
		},
		{
			idx:      10,
			expected: "10月",
		},
		{
			idx:      11,
			expected: "11月",
		},
		{
			idx:      12,
			expected: "12月",
		},
	}

	for _, tt := range tests {
		s := trans.MonthAbbreviated(time.Month(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsNarrow(t *testing.T) {

	trans := New()
	months := trans.MonthsNarrow()

	for i, month := range months {
		s := trans.MonthNarrow(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "1",
		},
		{
			idx:      2,
			expected: "2",
		},
		{
			idx:      3,
			expected: "3",
		},
		{
			idx:      4,
			expected: "4",
		},
		{
			idx:      5,
			expected: "5",
		},
		{
			idx:      6,
			expected: "6",
		},
		{
			idx:      7,
			expected: "7",
		},
		{
			idx:      8,
			expected: "8",
		},
		{
			idx:      9,
			expected: "9",
		},
		{
			idx:      10,
			expected: "10",
		},
		{
			idx:      11,
			expected: "11",
		},
		{
			idx:      12,
			expected: "12",
		},
	}

	for _, tt := range tests {
		s := trans.MonthNarrow(time.Month(tt.idx))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestMonthsWide(t *testing.T) {

	trans := New()
	months := trans.MonthsWide()

	for i, month := range months {
		s := trans.MonthWide(time.Month(i + 1))
		if s != month {
			t.Errorf("Expected '%s' Got '%s'", month, s)
		}
	}

	tests := []struct {
		idx      int
		expected string
	}{
		{
			idx:      1,
			expected: "1月",
		},
		{
			idx:      2,
			expected: "2月",
		},
		{
			idx:      3,
			expected: "3月",
		},
		{
			idx:      4,
			expected: "4月",
		},
		{
			idx:      5,
			expected: "5月",
		},
		{
			idx:      6,
			expected: "6月",
		},
		{
			idx:      7,
			expected: "7月",
		},
		{
			idx:      8,
			expected: "8月",
		},
		{
			idx:      9,
			expected: "9月",
		},
		{
			idx:      10,
			expected: "10月",
		},
		{
			idx:      11,
			expected: "11月",
		},
		{
			idx:      12,
			expected: "12月",
		},
	}

	for _, tt := range tests {
		s := string(trans.MonthWide(time.Month(tt.idx)))
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeFull(t *testing.T) {

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	fixed := time.FixedZone("OTHER", -4)

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, loc),
			expected: "9時05分01秒 日本標準時",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 5, 1, 0, fixed),
			expected: "20時05分01秒 OTHER",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeLong(t *testing.T) {

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Errorf("Expected '<nil>' Got '%s'", err)
	}

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, loc),
			expected: "9:05:01 JST",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 5, 1, 0, loc),
			expected: "20:05:01 JST",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeMedium(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, time.UTC),
			expected: "9:05:01",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 5, 1, 0, time.UTC),
			expected: "20:05:01",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtTimeShort(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 5, 1, 0, time.UTC),
			expected: "9:05",
		},
		{
			t:        time.Date(2016, 02, 03, 20, 5, 1, 0, time.UTC),
			expected: "20:05",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtTimeShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateFull(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "2016年2月3日水曜日",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateFull(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateLong(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "2016年2月3日",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateLong(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateMedium(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "2016/02/03",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateMedium(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtDateShort(t *testing.T) {

	tests := []struct {
		t        time.Time
		expected string
	}{
		{
			t:        time.Date(2016, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "2016/02/03",
		},
		{
			t:        time.Date(-500, 02, 03, 9, 0, 1, 0, time.UTC),
			expected: "500/02/03",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtDateShort(tt.t)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtNumber(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			expected: "1,123,456.56",
		},
		{
			num:      1123456.5643,
			v:        1,
			expected: "1,123,456.6",
		},
		{
			num:      221123456.5643,
			v:        3,
			expected: "221,123,456.564",
		},
		{
			num:      -221123456.5643,
			v:        3,
			expected: "-221,123,456.564",
		},
		{
			num:      -221123456.5643,
			v:        3,
			expected: "-221,123,456.564",
		},
		{
			num:      0,
			v:        2,
			expected: "0.00",
		},
		{
			num:      -0,
			v:        2,
			expected: "0.00",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtNumber(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtCurrency(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			currency: currency.JPY,
			expected: "￥1,123,456.56",
		},
		{
			num:      1123456.5643,
			v:        1,
			currency: currency.JPY,
			expected: "￥1,123,456.60",
		},
		{
			num:      221123456.5643,
			v:        3,
			currency: currency.JPY,
			expected: "￥221,123,456.564",
		},
		{
			num:      -221123456.5643,
			v:        3,
			currency: currency.JPY,
			expected: "-￥221,123,456.564",
		},
		{
			num:      0,
			v:        2,
			currency: currency.JPY,
			expected: "￥0.00",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.JPY,
			expected: "￥0.00",
		},
		{
			num:      1.23,
			v:        0,
			currency: currency.JPY,
			expected: "￥1.00",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtCurrency(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtAccounting(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		currency currency.Type
		expected string
	}{
		{
			num:      1123456.5643,
			v:        2,
			currency: currency.JPY,
			expected: "￥1,123,456.56",
		},
		{
			num:      1123456.5643,
			v:        1,
			currency: currency.JPY,
			expected: "￥1,123,456.60",
		},
		{
			num:      221123456.5643,
			v:        3,
			currency: currency.JPY,
			expected: "￥221,123,456.564",
		},
		{
			num:      -221123456.5643,
			v:        3,
			currency: currency.JPY,
			expected: "(￥221,123,456.564)",
		},
		{
			num:      -0,
			v:        2,
			currency: currency.JPY,
			expected: "￥0.00",
		},
		{
			num:      1.23,
			v:        0,
			currency: currency.JPY,
			expected: "￥1.00",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtAccounting(tt.num, tt.v, tt.currency)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}

func TestFmtPercent(t *testing.T) {

	tests := []struct {
		num      float64
		v        uint64
		expected string
	}{
		{
			num:      15,
			v:        0,
			expected: "15%",
		},
		{
			num:      15,
			v:        2,
			expected: "15.00%",
		},
		{
			num:      434.45,
			v:        0,
			expected: "434%",
		},
		{
			num:      34.4,
			v:        2,
			expected: "34.40%",
		},
		{
			num:      -34,
			v:        0,
			expected: "-34%",
		},
	}

	trans := New()

	for _, tt := range tests {
		s := trans.FmtPercent(tt.num, tt.v)
		if s != tt.expected {
			t.Errorf("Expected '%s' Got '%s'", tt.expected, s)
		}
	}
}
