package mi_NZ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type mi_NZ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'mi_NZ' locale
func New() locales.Translator {
	return &mi_NZ{
		locale:                 "mi_NZ",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Kohi", "Hui", "Pou", "Pae", "Hara", "Pipi", "Hōngo", "Here", "Mahu", "Nuku", "Rangi", "Haki"},
		monthsNarrow:           []string{"", "K", "H", "P", "P", "H", "P", "H", "H", "M", "N", "R", "H"},
		monthsWide:             []string{"", "Kohitātea", "Huitanguru", "Poutūterangi", "Paengawhāwhā", "Haratua", "Pipiri", "Hōngongoi", "Hereturikōkā", "Mahuru", "Whiringa-ā-nuku", "Whiringa-ā-rangi", "Hakihea"},
		daysAbbreviated:        []string{"Tap", "Hin", "Tū", "Apa", "Par", "Mer", "Hor"},
		daysNarrow:             []string{"T", "H", "T", "A", "P", "M", "H"},
		daysShort:              []string{"Tap", "Hin", "Tū", "Apa", "Par", "Mer", "Hor"},
		daysWide:               []string{"Rātapu", "Rāhina", "Rātū", "Rāapa", "Rāpare", "Rāmere", "Rāhoroi"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"BCE", "CE"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"BCE", "CE"},
		timezones:              map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "Wā Awatea Ranatiki", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "AKDT", "AKST": "AKST", "ARST": "ARST", "ART": "ART", "AST": "Wā Arowhānui Ranatiki", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "Wā Awatea Waenga", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "Wā Arowhānui Waenga", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "Wā Awatea Rāwhiti", "EST": "Wā Arowhānui Rāwhiti", "GFT": "GFT", "GMT": "Wā Toharite Greenwich", "GST": "GST", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "HKST", "HKT": "HKT", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "Wā Awatea Maunga", "MESZ": "Wā Raumati Uropi Waenga", "MEZ": "Wā Arowhānui Uropi Waenga", "MST": "Wā Arowhānui Maunga", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "Wā Raumati Uropi Rāwhiti", "OEZ": "Wā Arowhānui Uropi Rāwhiti", "PDT": "Wā Awatea Kiwa", "PST": "Wā Arowhānui Kiwa", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "Wā Raumati Uropi Uru", "WEZ": "Wā Arowhānui Uropi Uru", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (mi *mi_NZ) Locale() string {
	return mi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mi_NZ'
func (mi *mi_NZ) PluralsCardinal() []locales.PluralRule {
	return mi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mi_NZ'
func (mi *mi_NZ) PluralsOrdinal() []locales.PluralRule {
	return mi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mi_NZ'
func (mi *mi_NZ) PluralsRange() []locales.PluralRule {
	return mi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mi_NZ'
func (mi *mi_NZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mi_NZ'
func (mi *mi_NZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mi_NZ'
func (mi *mi_NZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mi *mi_NZ) MonthAbbreviated(month time.Month) string {
	return mi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mi *mi_NZ) MonthsAbbreviated() []string {
	return mi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mi *mi_NZ) MonthNarrow(month time.Month) string {
	return mi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mi *mi_NZ) MonthsNarrow() []string {
	return mi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mi *mi_NZ) MonthWide(month time.Month) string {
	return mi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mi *mi_NZ) MonthsWide() []string {
	return mi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mi *mi_NZ) WeekdayAbbreviated(weekday time.Weekday) string {
	return mi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mi *mi_NZ) WeekdaysAbbreviated() []string {
	return mi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mi *mi_NZ) WeekdayNarrow(weekday time.Weekday) string {
	return mi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mi *mi_NZ) WeekdaysNarrow() []string {
	return mi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mi *mi_NZ) WeekdayShort(weekday time.Weekday) string {
	return mi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mi *mi_NZ) WeekdaysShort() []string {
	return mi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mi *mi_NZ) WeekdayWide(weekday time.Weekday) string {
	return mi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mi *mi_NZ) WeekdaysWide() []string {
	return mi.daysWide
}

// Decimal returns the decimal point of number
func (mi *mi_NZ) Decimal() string {
	return mi.decimal
}

// Group returns the group of number
func (mi *mi_NZ) Group() string {
	return mi.group
}

// Group returns the minus sign of number
func (mi *mi_NZ) Minus() string {
	return mi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mi_NZ' and handles both Whole and Real numbers based on 'v'
func (mi *mi_NZ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mi_NZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mi *mi_NZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mi_NZ'
func (mi *mi_NZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(mi.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, mi.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mi_NZ'
// in accounting notation.
func (mi *mi_NZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mi.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, mi.currencyNegativePrefix[j])
		}

		b = append(b, mi.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mi.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, mi.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, mi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mi_NZ'
func (mi *mi_NZ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
