package ig_NG

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ig_NG struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	percentSuffix      string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'ig_NG' locale
func New() locales.Translator {
	return &ig_NG{
		locale:             "ig_NG",
		pluralsCardinal:    []locales.PluralRule{6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            "٫",
		group:              "٬",
		minus:              "‏-",
		percent:            "٪‏",
		perMille:           "؉",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:      " ",
		monthsAbbreviated:  []string{"", "Jen", "Feb", "Maa", "Epr", "Mee", "Juu", "Jul", "Ọgọ", "Sep", "Ọkt", "Nov", "Dis"},
		monthsNarrow:       []string{"", "J", "F", "M", "E", "M", "J", "J", "Ọ", "S", "Ọ", "N", "D"},
		monthsWide:         []string{"", "Jenụwarị", "Febrụwarị", "Maachị", "Epreel", "Mee", "Juun", "Julaị", "Ọgọọst", "Septemba", "Ọktoba", "Novemba", "Disemba"},
		daysAbbreviated:    []string{"Ụka", "Mọn", "Tiu", "Wen", "Tọọ", "Fraị", "Sat"},
		daysShort:          []string{"Sọn", "Mọn", "Tiu", "Wen", "Tọọ", "Fraị", "Sat"},
		daysWide:           []string{"Sọndee", "Mọnde", "Tiuzdee", "Wenezdee", "Tọọzdee", "Fraịdee", "Satọdee"},
		periodsAbbreviated: []string{"A.M.", "P.M."},
		periodsNarrow:      []string{"A.M.", "P.M."},
		periodsWide:        []string{"N’ụtụtụ", "N’abali"},
		erasAbbreviated:    []string{"T.K.", "A.K."},
		erasNarrow:         []string{"T.K.", "A.K."},
		erasWide:           []string{"Tupu Kraist", "Afọ Kraịst"},
		timezones:          map[string]string{"ACDT": "Oge Ihe Etiti Australia", "ACST": "Oge Izugbe Etiti Australia", "ACWDT": "Oge Ihe Mpaghara Ọdịda Anyanwụ Etiti Australia", "ACWST": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Etiti Australia", "ADT": "Oge Ihe Mpaghara Atlantic", "AEDT": "Oge Ihe Mpaghara Ọwụwa Anyanwụ Australia", "AEST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Australia", "AKDT": "Oge Ihe Alaska", "AKST": "Oge Izugbe Alaska", "ARST": "Oge Okpomọkụ Argentina", "ART": "Oge Izugbe Argentina", "AST": "Oge Izugbe Mpaghara Atlantic", "AWDT": "Oge Ihe Mpaghara Ọdịda Anyanwụ Australia", "AWST": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Australia", "BOT": "Oge Bolivia", "BT": "Oge Bhutan", "CAT": "Oge Etiti Afrịka", "CDT": "Oge Ihe Mpaghara Etiti", "CHADT": "Oge Ihe Chatham", "CHAST": "Oge Izugbe Chatham", "CLST": "Oge Okpomọkụ Chile", "CLT": "Oge Izugbe Chile", "COST": "Oge Okpomọkụ Columbia", "COT": "Oge Izugbe Columbia", "CST": "Oge Izugbe Mpaghara Etiti", "ChST": "Oge Izugbe Chamorro", "EAT": "Oge Mpaghara Ọwụwa Anyanwụ Afrịka", "ECT": "Oge Ecuador", "EDT": "Oge Ihe Mpaghara Ọwụwa Anyanwụ", "EST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ", "GFT": "Oge French Guiana", "GMT": "Oge Mpaghara Greemwich Mean", "GST": "Oge Izugbe Gulf", "GYT": "Oge Guyana", "HADT": "Oge Ihe Hawaii-Aleutian", "HAST": "Oge Izugbe Hawaii-Aleutian", "HAT": "Oge Ihe Newfoundland", "HECU": "Oge Ihe Mpaghara Cuba", "HEEG": "Oge Okpomọkụ Mpaghara Ọwụwa Anyanwụ Greenland", "HENOMX": "Oge Ihe Northwest Mexico", "HEOG": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Greenland", "HEPM": "Oge Ihe St. Pierre & Miquelon", "HEPMX": "Oge Ihe Mexican Pacific", "HKST": "Oge Okpomọkụ Hong Kong", "HKT": "Oge Izugbe Hong Kong", "HNCU": "Oge Izugbe Cuba", "HNEG": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Greenland", "HNNOMX": "Oge Izugbe Northwest Mexico", "HNOG": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Greenland", "HNPM": "Oge Izugbe St. Pierre & Miquelon", "HNPMX": "Oge Izugbe Mexican Pacific", "HNT": "Oge Izugbe Newfoundland", "IST": "Oge Izugbe India", "JDT": "Oge Ihe Japan", "JST": "Oge Izugbe Japan", "LHDT": "Oge Ihe Lord Howe", "LHST": "Oge Izugbe Lord Howe", "MDT": "Oge Ihe Mpaghara Ugwu", "MESZ": "Oge Okpomọkụ Mpaghara Etiti Europe", "MEZ": "Oge Izugbe Mpaghara Etiti Europe", "MST": "Oge Izugbe Mpaghara Ugwu", "MYT": "Oge Malaysia", "NZDT": "Oge Ihe New Zealand", "NZST": "Oge Izugbe New Zealand", "OESZ": "Oge Okpomọkụ Mpaghara Ọwụwa Anyanwụ Europe", "OEZ": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Europe", "PDT": "Oge Ihe Mpaghara Pacific", "PST": "Oge Izugbe Mpaghara Pacific", "SAST": "Oge Izugbe Mpaghara Mgbada Ugwu Afrịka", "SGT": "Oge Izugbe Singapore", "SRT": "Oge Suriname", "TMST": "Oge Okpomọkụ Turkmenist", "TMT": "Oge Izugbe Turkmenist", "UYST": "Oge Okpomọkụ Uruguay", "UYT": "Oge Izugbe Uruguay", "VET": "Oge Venezuela", "WARST": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Argentina", "WART": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Argentina", "WAST": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Afrịka", "WAT": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Afrịka", "WESZ": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Europe", "WEZ": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Europe", "WIB": "Oge Mpaghara Ọdịda Anyanwụ Indonesia", "WIT": "Oge Mpaghara Ọwụwa Anyanwụ Indonesia", "WITA": "Oge Etiti Indonesia", "∅∅∅": "Oge Okpomọkụ Azores"},
	}
}

// Locale returns the current translators string locale
func (ig *ig_NG) Locale() string {
	return ig.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsCardinal() []locales.PluralRule {
	return ig.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsOrdinal() []locales.PluralRule {
	return ig.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsRange() []locales.PluralRule {
	return ig.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ig_NG'
func (ig *ig_NG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ig *ig_NG) MonthAbbreviated(month time.Month) string {
	return ig.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ig *ig_NG) MonthsAbbreviated() []string {
	return ig.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ig *ig_NG) MonthNarrow(month time.Month) string {
	return ig.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ig *ig_NG) MonthsNarrow() []string {
	return ig.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ig *ig_NG) MonthWide(month time.Month) string {
	return ig.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ig *ig_NG) MonthsWide() []string {
	return ig.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayAbbreviated(weekday time.Weekday) string {
	return ig.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ig *ig_NG) WeekdaysAbbreviated() []string {
	return ig.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayNarrow(weekday time.Weekday) string {
	return ig.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ig *ig_NG) WeekdaysNarrow() []string {
	return ig.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayShort(weekday time.Weekday) string {
	return ig.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ig *ig_NG) WeekdaysShort() []string {
	return ig.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayWide(weekday time.Weekday) string {
	return ig.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ig *ig_NG) WeekdaysWide() []string {
	return ig.daysWide
}

// Decimal returns the decimal point of number
func (ig *ig_NG) Decimal() string {
	return ig.decimal
}

// Group returns the group of number
func (ig *ig_NG) Group() string {
	return ig.group
}

// Group returns the minus sign of number
func (ig *ig_NG) Minus() string {
	return ig.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
func (ig *ig_NG) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ig *ig_NG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 13
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ig.decimal) - 1; j >= 0; j-- {
				b = append(b, ig.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ig.percentSuffix...)

	b = append(b, ig.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ig.decimal) - 1; j >= 0; j-- {
				b = append(b, ig.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ig.group) - 1; j >= 0; j-- {
					b = append(b, ig.group[j])
				}
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

	if num < 0 {
		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
// in accounting notation.
func (ig *ig_NG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ig.decimal) - 1; j >= 0; j-- {
				b = append(b, ig.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ig.group) - 1; j >= 0; j-- {
					b = append(b, ig.group[j])
				}
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

		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ig.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ig.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
