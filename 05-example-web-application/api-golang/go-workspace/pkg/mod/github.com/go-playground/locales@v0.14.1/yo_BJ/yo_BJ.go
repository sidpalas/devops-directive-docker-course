package yo_BJ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type yo_BJ struct {
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

// New returns a new instance of translator for the 'yo_BJ' locale
func New() locales.Translator {
	return &yo_BJ{
		locale:                 "yo_BJ",
		pluralsCardinal:        []locales.PluralRule{6},
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
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Shɛ́r", "Èrèl", "Ɛrɛ̀n", "Ìgb", "Ɛ̀bi", "Òkú", "Agɛ", "Ògú", "Owe", "Ɔ̀wà", "Bél", "Ɔ̀pɛ"},
		monthsNarrow:           []string{"", "S", "È", "Ɛ", "Ì", "Ɛ̀", "Ò", "A", "Ò", "O", "Ɔ̀", "B", "Ɔ̀"},
		monthsWide:             []string{"", "Oshù Shɛ́rɛ́", "Oshù Èrèlè", "Oshù Ɛrɛ̀nà", "Oshù Ìgbé", "Oshù Ɛ̀bibi", "Oshù Òkúdu", "Oshù Agɛmɔ", "Oshù Ògún", "Oshù Owewe", "Oshù Ɔ̀wàrà", "Oshù Bélú", "Oshù Ɔ̀pɛ̀"},
		daysAbbreviated:        []string{"Àìk", "Aj", "Ìsɛ́g", "Ɔjɔ́r", "Ɔjɔ́b", "Ɛt", "Àbám"},
		daysNarrow:             []string{"À", "A", "Ì", "Ɔ", "Ɔ", "Ɛ", "À"},
		daysShort:              []string{"Àìkú", "Ajé", "Ìsɛ́gun", "Ɔjɔ́rú", "Ɔjɔ́bɔ", "Ɛtì", "Àbámɛ́ta"},
		daysWide:               []string{"Ɔjɔ́ Àìkú", "Ɔjɔ́ Ajé", "Ɔjɔ́ Ìsɛ́gun", "Ɔjɔ́rú", "Ɔjɔ́bɔ", "Ɔjɔ́ Ɛtì", "Ɔjɔ́ Àbámɛ́ta"},
		periodsAbbreviated:     []string{"Àárɔ̀", "Ɔ̀sán"},
		periodsNarrow:          []string{"Àárɔ̀", "Ɔ̀sán"},
		periodsWide:            []string{"Àárɔ̀", "Ɔ̀sán"},
		erasAbbreviated:        []string{"BCE", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Saju Kristi", "Lehin Kristi"},
		timezones:              map[string]string{"ACDT": "Australian Central Daylight Time", "ACST": "Australian Central Standard Time", "ACWDT": "Australian Central Western Daylight Time", "ACWST": "Australian Central Western Standard Time", "ADT": "Àkókò Ìyálɛta Àtìláńtíìkì", "AEDT": "Australian Eastern Daylight Time", "AEST": "Australian Eastern Standard Time", "AKDT": "Alaska Daylight Time", "AKST": "Alaska Standard Time", "ARST": "Argentina Summer Time", "ART": "Argentina Standard Time", "AST": "Àkókò àsikò Àtìláńtíìkì", "AWDT": "Australian Western Daylight Time", "AWST": "Australian Western Standard Time", "BOT": "Bolivia Time", "BT": "Bhutan Time", "CAT": "Central Africa Time", "CDT": "Akókò àárín gbùngbùn ojúmɔmɔ", "CHADT": "Chatham Daylight Time", "CHAST": "Chatham Standard Time", "CLST": "Chile Summer Time", "CLT": "Chile Standard Time", "COST": "Colombia Summer Time", "COT": "Colombia Standard Time", "CST": "àkókò asiko àárín gbùngbùn", "ChST": "Chamorro Standard Time", "EAT": "East Africa Time", "ECT": "Ecuador Time", "EDT": "Àkókò ojúmɔmɔ Ìhà Ìlà Oòrun", "EST": "Akókò Àsikò Ìha Ìla Oòrùn", "GFT": "French Guiana Time", "GMT": "Greenwich Mean Time", "GST": "Gulf Standard Time [translation hint: translate as just \"Gulf Time\"]", "GYT": "Guyana Time", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "Hong Kong Summer Time", "HKT": "Hong Kong Standard Time", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "India Standard Time", "JDT": "Japan Daylight Time", "JST": "Japan Standard Time", "LHDT": "Lord Howe Daylight Time", "LHST": "Lord Howe Standard Time", "MDT": "Àkókò ojúmɔmɔ Ori-òkè", "MESZ": "Àkókò Àárin Sɔmà Europe", "MEZ": "Àkókò Àárin àsikò Europe", "MST": "Àkókò asiko òkè", "MYT": "Malaysia Time", "NZDT": "New Zealand Daylight Time", "NZST": "New Zealand Standard Time", "OESZ": "Àkókò Sɔmà Ìha Ìlà Oòrùn Europe", "OEZ": "Àkókò àsikò Ìhà Ìlà Oòrùn Europe", "PDT": "Àkókò Ìyálɛta Pàsífíìkì", "PST": "Àkókò àsikò Pàsífíìkì", "SAST": "South Africa Standard Time", "SGT": "Singapore Standard Time", "SRT": "Suriname Time", "TMST": "Turkmenistan Summer Time", "TMT": "Turkmenistan Standard Time", "UYST": "Uruguay Summer Time", "UYT": "Uruguay Standard Time", "VET": "Venezuela Time", "WARST": "Western Argentina Summer Time", "WART": "Western Argentina Standard Time", "WAST": "West Africa Summer Time", "WAT": "West Africa Standard Time", "WESZ": "Àkókò Sɔmà Ìhà Ìwɔ Oòrùn Europe", "WEZ": "Àkókò àsikò Ìwɔ Oòrùn Europe", "WIB": "WIB", "WIT": "Eastern Indonesia Time", "WITA": "WITA", "∅∅∅": "Brasilia Summer Time"},
	}
}

// Locale returns the current translators string locale
func (yo *yo_BJ) Locale() string {
	return yo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsCardinal() []locales.PluralRule {
	return yo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsOrdinal() []locales.PluralRule {
	return yo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yo_BJ'
func (yo *yo_BJ) PluralsRange() []locales.PluralRule {
	return yo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yo_BJ'
func (yo *yo_BJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yo *yo_BJ) MonthAbbreviated(month time.Month) string {
	return yo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yo *yo_BJ) MonthsAbbreviated() []string {
	return yo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yo *yo_BJ) MonthNarrow(month time.Month) string {
	return yo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yo *yo_BJ) MonthsNarrow() []string {
	return yo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yo *yo_BJ) MonthWide(month time.Month) string {
	return yo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yo *yo_BJ) MonthsWide() []string {
	return yo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayAbbreviated(weekday time.Weekday) string {
	return yo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yo *yo_BJ) WeekdaysAbbreviated() []string {
	return yo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayNarrow(weekday time.Weekday) string {
	return yo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yo *yo_BJ) WeekdaysNarrow() []string {
	return yo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayShort(weekday time.Weekday) string {
	return yo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yo *yo_BJ) WeekdaysShort() []string {
	return yo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yo *yo_BJ) WeekdayWide(weekday time.Weekday) string {
	return yo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yo *yo_BJ) WeekdaysWide() []string {
	return yo.daysWide
}

// Decimal returns the decimal point of number
func (yo *yo_BJ) Decimal() string {
	return yo.decimal
}

// Group returns the group of number
func (yo *yo_BJ) Group() string {
	return yo.group
}

// Group returns the minus sign of number
func (yo *yo_BJ) Minus() string {
	return yo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
func (yo *yo_BJ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yo_BJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yo *yo_BJ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
func (yo *yo_BJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
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
		b = append(b, yo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yo_BJ'
// in accounting notation.
func (yo *yo_BJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yo.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yo.group[0])
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

		b = append(b, yo.currencyNegativePrefix[0])

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
			b = append(b, yo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, yo.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, yo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'yo_BJ'
func (yo *yo_BJ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
