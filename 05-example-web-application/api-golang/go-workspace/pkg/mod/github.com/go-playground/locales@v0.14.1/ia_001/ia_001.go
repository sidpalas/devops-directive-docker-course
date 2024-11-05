package ia_001

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ia_001 struct {
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
	currencyNegativeSuffix string
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

// New returns a new instance of translator for the 'ia_001' locale
func New() locales.Translator {
	return &ia_001{
		locale:                 "ia_001",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan", "feb", "mar", "apr", "mai", "jun", "jul", "aug", "sep", "oct", "nov", "dec"},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januario", "februario", "martio", "april", "maio", "junio", "julio", "augusto", "septembre", "octobre", "novembre", "decembre"},
		daysAbbreviated:        []string{"dom", "lun", "mar", "mer", "jov", "ven", "sab"},
		daysNarrow:             []string{"d", "l", "m", "m", "j", "v", "s"},
		daysShort:              []string{"do", "lu", "ma", "me", "jo", "ve", "sa"},
		daysWide:               []string{"dominica", "lunedi", "martedi", "mercuridi", "jovedi", "venerdi", "sabbato"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"a.Chr.", "p.Chr."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ante Christo", "post Christo"},
		timezones:              map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "hora estive atlantic", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "hora estive de Alaska", "AKST": "hora normal de Alaska", "ARST": "ARST", "ART": "ART", "AST": "hora normal atlantic", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "hora estive central", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "hora normal central", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "hora estive del est", "EST": "hora normal del est", "GFT": "GFT", "GMT": "hora medie de Greenwich", "GST": "GST", "GYT": "GYT", "HADT": "hora estive de Hawaii-Aleutianas", "HAST": "hora normal de Hawaii-Aleutianas", "HAT": "hora estive de Terranova", "HECU": "hora estive de Cuba", "HEEG": "hora estive de Groenlandia oriental", "HENOMX": "hora estive del nordwest de Mexico", "HEOG": "hora estive de Groenlandia occidental", "HEPM": "hora estive de Saint-Pierre e Miquelon", "HEPMX": "hora estive del Pacifico mexican", "HKST": "HKST", "HKT": "HKT", "HNCU": "hora normal de Cuba", "HNEG": "hora normal de Groenlandia oriental", "HNNOMX": "hora normal del nordwest de Mexico", "HNOG": "hora normal de Groenlandia occidental", "HNPM": "hora normal de Saint-Pierre e Miquelon", "HNPMX": "hora normal del Pacifico mexican", "HNT": "hora normal de Terranova", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "hora estive del montanias", "MESZ": "hora estive de Europa central", "MEZ": "hora normal de Europa central", "MST": "hora normal del montanias", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "hora estive de Europa oriental", "OEZ": "hora normal de Europa oriental", "PDT": "hora estive pacific", "PST": "hora normal pacific", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "hora estive de Europa occidental", "WEZ": "hora normal de Europa occidental", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (ia *ia_001) Locale() string {
	return ia.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ia_001'
func (ia *ia_001) PluralsCardinal() []locales.PluralRule {
	return ia.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ia_001'
func (ia *ia_001) PluralsOrdinal() []locales.PluralRule {
	return ia.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ia_001'
func (ia *ia_001) PluralsRange() []locales.PluralRule {
	return ia.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ia_001'
func (ia *ia_001) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ia_001'
func (ia *ia_001) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ia_001'
func (ia *ia_001) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ia *ia_001) MonthAbbreviated(month time.Month) string {
	return ia.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ia *ia_001) MonthsAbbreviated() []string {
	return ia.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ia *ia_001) MonthNarrow(month time.Month) string {
	return ia.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ia *ia_001) MonthsNarrow() []string {
	return ia.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ia *ia_001) MonthWide(month time.Month) string {
	return ia.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ia *ia_001) MonthsWide() []string {
	return ia.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ia *ia_001) WeekdayAbbreviated(weekday time.Weekday) string {
	return ia.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ia *ia_001) WeekdaysAbbreviated() []string {
	return ia.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ia *ia_001) WeekdayNarrow(weekday time.Weekday) string {
	return ia.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ia *ia_001) WeekdaysNarrow() []string {
	return ia.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ia *ia_001) WeekdayShort(weekday time.Weekday) string {
	return ia.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ia *ia_001) WeekdaysShort() []string {
	return ia.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ia *ia_001) WeekdayWide(weekday time.Weekday) string {
	return ia.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ia *ia_001) WeekdaysWide() []string {
	return ia.daysWide
}

// Decimal returns the decimal point of number
func (ia *ia_001) Decimal() string {
	return ia.decimal
}

// Group returns the group of number
func (ia *ia_001) Group() string {
	return ia.group
}

// Group returns the minus sign of number
func (ia *ia_001) Minus() string {
	return ia.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ia_001' and handles both Whole and Real numbers based on 'v'
func (ia *ia_001) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ia_001' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ia *ia_001) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ia.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ia_001'
func (ia *ia_001) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ia.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
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

	for j := len(ia.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ia.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ia.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ia.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ia_001'
// in accounting notation.
func (ia *ia_001) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ia.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ia.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ia.group[0])
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

		for j := len(ia.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ia.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ia.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ia.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ia.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ia.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ia_001'
func (ia *ia_001) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ia.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0x6c, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, ia.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ia_001'
func (ia *ia_001) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ia.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ia.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
