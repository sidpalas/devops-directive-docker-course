package wo_SN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type wo_SN struct {
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

// New returns a new instance of translator for the 'wo_SN' locale
func New() locales.Translator {
	return &wo_SN{
		locale:                 "wo_SN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Sam", "Few", "Mar", "Awr", "Mee", "Suw", "Sul", "Ut", "Sàt", "Okt", "Now", "Des"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "Samwiyee", "Fewriyee", "Mars", "Awril", "Mee", "Suwe", "Sulet", "Ut", "Sàttumbar", "Oktoobar", "Nowàmbar", "Desàmbar"},
		daysAbbreviated:        []string{"Dib", "Alt", "Tal", "Àla", "Alx", "Àjj", "Ase"},
		daysNarrow:             []string{"Dib", "Alt", "Tal", "Àla", "Alx", "Àjj", "Ase"},
		daysShort:              []string{"Dib", "Alt", "Tal", "Àla", "Alx", "Àjj", "Ase"},
		daysWide:               []string{"Dibéer", "Altine", "Talaata", "Àlarba", "Alxamis", "Àjjuma", "Aseer"},
		periodsAbbreviated:     []string{"Sub", "Ngo"},
		periodsNarrow:          []string{"Sub", "Ngo"},
		periodsWide:            []string{"Sub", "Ngo"},
		erasAbbreviated:        []string{"JC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"av. JC", "AD"},
		timezones:              map[string]string{"ACDT": "ACDT", "ACST": "ACST", "ACWDT": "ACWDT", "ACWST": "ACWST", "ADT": "ADT (waxtu bëccëgu atlàntik)", "AEDT": "AEDT", "AEST": "AEST", "AKDT": "AKDT", "AKST": "AKST", "ARST": "ARST", "ART": "ART", "AST": "AST (waxtu estàndaaru penku)", "AWDT": "AWDT", "AWST": "AWST", "BOT": "BOT", "BT": "BT", "CAT": "CAT", "CDT": "CDT (waxtu bëccëgu sàntaraal", "CHADT": "CHADT", "CHAST": "CHAST", "CLST": "CLST", "CLT": "CLT", "COST": "COST", "COT": "COT", "CST": "CST (waxtu estàndaaru sàntaraal)", "ChST": "ChST", "EAT": "EAT", "ECT": "ECT", "EDT": "EDT (waxtu bëccëgu penku)", "EST": "EST (waxtu estàndaaru penku)", "GFT": "GFT", "GMT": "GMT (waxtu Greenwich)", "GST": "GST", "GYT": "GYT", "HADT": "HADT", "HAST": "HAST", "HAT": "HAT", "HECU": "HECU", "HEEG": "HEEG", "HENOMX": "HENOMX", "HEOG": "HEOG", "HEPM": "HEPM", "HEPMX": "HEPMX", "HKST": "HKST", "HKT": "HKT", "HNCU": "HNCU", "HNEG": "HNEG", "HNNOMX": "HNNOMX", "HNOG": "HNOG", "HNPM": "HNPM", "HNPMX": "HNPMX", "HNT": "HNT", "IST": "IST", "JDT": "JDT", "JST": "JST", "LHDT": "LHDT", "LHST": "LHST", "MDT": "MDT (waxtu bëccëgu tundu)", "MESZ": "CEST (waxtu ete wu ëroop sàntaraal)", "MEZ": "CEST (waxtu estàndaaru ëroop sàntaraal)", "MST": "MST (waxtu estàndaaru tundu)", "MYT": "MYT", "NZDT": "NZDT", "NZST": "NZST", "OESZ": "EEST (waxtu ete wu ëroop u penku)", "OEZ": "EEST (waxtu estàndaaru ëroop u penku)", "PDT": "PDT (waxtu bëccëgu pasifik)", "PST": "PST (waxtu estàndaaru pasifik)", "SAST": "SAST", "SGT": "SGT", "SRT": "SRT", "TMST": "TMST", "TMT": "TMT", "UYST": "UYST", "UYT": "UYT", "VET": "VET", "WARST": "WARST", "WART": "WART", "WAST": "WAST", "WAT": "WAT", "WESZ": "WEST (waxtu ete wu ëroop u sowwu-jant)", "WEZ": "WEST (waxtu estàndaaru ëroop u sowwu-jant)", "WIB": "WIB", "WIT": "WIT", "WITA": "WITA", "∅∅∅": "∅∅∅"},
	}
}

// Locale returns the current translators string locale
func (wo *wo_SN) Locale() string {
	return wo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'wo_SN'
func (wo *wo_SN) PluralsCardinal() []locales.PluralRule {
	return wo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'wo_SN'
func (wo *wo_SN) PluralsOrdinal() []locales.PluralRule {
	return wo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'wo_SN'
func (wo *wo_SN) PluralsRange() []locales.PluralRule {
	return wo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'wo_SN'
func (wo *wo_SN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'wo_SN'
func (wo *wo_SN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'wo_SN'
func (wo *wo_SN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (wo *wo_SN) MonthAbbreviated(month time.Month) string {
	return wo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (wo *wo_SN) MonthsAbbreviated() []string {
	return wo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (wo *wo_SN) MonthNarrow(month time.Month) string {
	return wo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (wo *wo_SN) MonthsNarrow() []string {
	return wo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (wo *wo_SN) MonthWide(month time.Month) string {
	return wo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (wo *wo_SN) MonthsWide() []string {
	return wo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (wo *wo_SN) WeekdayAbbreviated(weekday time.Weekday) string {
	return wo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (wo *wo_SN) WeekdaysAbbreviated() []string {
	return wo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (wo *wo_SN) WeekdayNarrow(weekday time.Weekday) string {
	return wo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (wo *wo_SN) WeekdaysNarrow() []string {
	return wo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (wo *wo_SN) WeekdayShort(weekday time.Weekday) string {
	return wo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (wo *wo_SN) WeekdaysShort() []string {
	return wo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (wo *wo_SN) WeekdayWide(weekday time.Weekday) string {
	return wo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (wo *wo_SN) WeekdaysWide() []string {
	return wo.daysWide
}

// Decimal returns the decimal point of number
func (wo *wo_SN) Decimal() string {
	return wo.decimal
}

// Group returns the group of number
func (wo *wo_SN) Group() string {
	return wo.group
}

// Group returns the minus sign of number
func (wo *wo_SN) Minus() string {
	return wo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'wo_SN' and handles both Whole and Real numbers based on 'v'
func (wo *wo_SN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, wo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, wo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'wo_SN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (wo *wo_SN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, wo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, wo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'wo_SN'
func (wo *wo_SN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := wo.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, wo.group[0])
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

	for j := len(wo.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, wo.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, wo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, wo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'wo_SN'
// in accounting notation.
func (wo *wo_SN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := wo.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, wo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, wo.group[0])
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

		for j := len(wo.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, wo.currencyNegativePrefix[j])
		}

		b = append(b, wo.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(wo.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, wo.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, wo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, wo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, wo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'wo_SN'
func (wo *wo_SN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, wo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := wo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
