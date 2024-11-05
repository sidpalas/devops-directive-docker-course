package kok_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type kok_IN struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
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

// New returns a new instance of translator for the 'kok_IN' locale
func New() locales.Translator {
	return &kok_IN{
		locale:             "kok_IN",
		pluralsCardinal:    nil,
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		perMille:           "‰",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "जानेवारी", "फेब्रुवारी", "मार्च", "एप्रिल", "मे", "जून", "जुलाय", "आगोस्त", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
		monthsNarrow:       []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:         []string{"", "जानेवारी", "फेब्रुवारी", "मार्च", "एप्रिल", "मे", "जून", "जुलाय", "आगोस्त", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
		daysAbbreviated:    []string{"आयतार", "सोमार", "मंगळार", "बुधवार", "गुरुवार", "शुक्रार", "शेनवार"},
		daysNarrow:         []string{"आ", "सो", "मं", "बु", "गु", "शु", "शे"},
		daysShort:          []string{"आय", "सोम", "मंगळ", "बुध", "गुरु", "शुक्र", "शेन"},
		daysWide:           []string{"आयतार", "सोमार", "मंगळार", "बुधवार", "गुरुवार", "शुक्रार", "शेनवार"},
		periodsAbbreviated: []string{"AM", "PM"},
		periodsNarrow:      []string{"a", "p"},
		periodsWide:        []string{"AM", "PM"},
		erasAbbreviated:    []string{"क्रिस्तपूर्व", "क्रिस्तशखा"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"क्रिस्तपूर्व", "क्रिस्तशखा"},
		timezones:          map[string]string{"ACDT": "ऑस्ट्रेलीयन मध्य डेलायट वेळ", "ACST": "ऑस्ट्रेलीयन मध्य प्रमाणित वेळ", "ACWDT": "ऑस्ट्रेलीयन मध्य अस्तंत डेलायट वेळ", "ACWST": "ऑस्ट्रेलीयन मध्य अस्तंत प्रमाणित वेळ", "ADT": "अटलांटीक डेलायट वेळ", "AEDT": "ऑस्ट्रेलीयन उदेंत डेलायट वेळ", "AEST": "ऑस्ट्रेलीयन उदेंत प्रमाणित वेळ", "AKDT": "अलास्का डेलायट वेळ", "AKST": "अलास्का प्रमाणित वेळ", "ARST": "अर्जेंटिना ग्रीष्म वेळ", "ART": "अर्जेंटिना प्रमाणित वेळ", "AST": "अटलांटीक प्रमाणित वेळ", "AWDT": "ऑस्ट्रेलीयन अस्तंत डेलायट वेळ", "AWST": "ऑस्ट्रेलीयन अस्तंत प्रमाणित वेळ", "BOT": "बोलिव्हिया वेळ", "BT": "भूतान", "CAT": "मध्य आफ्रिका वेळ", "CDT": "मध्य डेलायट वेळ", "CHADT": "चॅथम डेलायट वेळ", "CHAST": "चॅथम प्रमाणित वेळ", "CLST": "चिली ग्रीष्म वेळ", "CLT": "चिली प्रमाणित वेळ", "COST": "कोलंबिया ग्रीष्म वेळ", "COT": "कोलंबिया प्रमाणित वेळ", "CST": "मध्य प्रमाणित वेळ", "ChST": "कॅमोरा प्रमाणित वेळ", "EAT": "उदेंत आफ्रिका वेळ", "ECT": "इक्वेडोर वेळ", "EDT": "उदेंत डेलायट वेळ", "EST": "उदेंत प्रमाणित वेळ", "GFT": "फ्रेंच गयाना वेळ", "GMT": "ग्रीनविच मध्यc वेळ", "GST": "गल्फ प्रमाणित वेळ", "GYT": "गुयाना वेळ", "HADT": "हवाई-अलेयुशिन डेलायट वेळ", "HAST": "हवाई-अलेयुशिन प्रमाणित वेळ", "HAT": "न्युफावंडलँड डेलायट वेळ", "HECU": "क्युबा डेलायट वेळ", "HEEG": "उदेंत ग्रीनलँड ग्रीष्म वेळ", "HENOMX": "वायव्य मेक्सिको डेलायट वेळ", "HEOG": "अस्तंत ग्रीनलँड ग्रीष्म वेळ", "HEPM": "सेंट पियर आनी मिकलान डेलायट वेळ", "HEPMX": "मेक्सिकन प्रशांत डेलायट वेळ", "HKST": "हाँग काँग ग्रीष्म वेळ", "HKT": "हाँग काँग प्रमाणित वेळ", "HNCU": "क्युबा प्रमाणित वेळ", "HNEG": "उदेंत ग्रीनलँड प्रमाणित वेळ", "HNNOMX": "वायव्य मेक्सिको प्रमाणित वेळ", "HNOG": "अस्तंत ग्रीनलँड प्रमाणित वेळ", "HNPM": "सेंट पियर आनी मिकलान प्रमाणित वेळ", "HNPMX": "मेक्सिकन प्रशांत प्रमाणित वेळ", "HNT": "न्युफावंडलँड प्रमाणित वेळ", "IST": "भारतीय प्रमाणित वेळ", "JDT": "जपान डेलायट वेळ", "JST": "जपान प्रमाणित वेळ", "LHDT": "लॉर्ड होवे डेलायट वेळ", "LHST": "लॉर्ड होवे प्रमाणित वेळ", "MDT": "पर्वतीय डेलायट वेळ", "MESZ": "मध्य युवरोपियन ग्रीष्म वेळ", "MEZ": "मध्य युरोपियन प्रमाणित वेळ", "MST": "पर्वतीय प्रमाणित वेळ", "MYT": "मलेशिया वेळ", "NZDT": "न्युझीलॅन्ड डेलायट वेळ", "NZST": "न्युझीलॅन्ड प्रमाणित वेळ", "OESZ": "उदेंत युरोपियन ग्रीष्म वेळ", "OEZ": "उदेंत युरोपियन प्रमाणित वेळ", "PDT": "प्रशांत डेलायट वेळ", "PST": "प्रशांत प्रमाणित वेळ", "SAST": "दक्षिण आफ्रिका प्रमाणित वेळ", "SGT": "सिंगापूर प्रमाणित वेळ", "SRT": "सुरिनाम वेळ", "TMST": "तुर्कमेनिस्तान ग्रीष्म वेळ", "TMT": "तुर्कमेनिस्तान प्रमाणित वेळ", "UYST": "उरुग्वे ग्रीष्म वेळ", "UYT": "उरुग्वे प्रमाणित वेळ", "VET": "वेनेझुएला वेळ", "WARST": "अस्तंत अर्जेंटिना ग्रीष्म वेळ", "WART": "अस्तंत अर्जेंटिना प्रमाणित वेळ", "WAST": "अस्तंत आफ्रिका ग्रीष्म वेळ", "WAT": "अस्तंत आफ्रिका प्रमाणित वेळ", "WESZ": "अस्तंत युरोपियन ग्रीष्म वेळ", "WEZ": "अस्तंत युरोपियन प्रमाणित वेळ", "WIB": "अस्तंत इंडोनेशिया वेळ", "WIT": "उदेंत इंडोनेशिया वेळ", "WITA": "मध्य इंडोनेशिया वेळ", "∅∅∅": "अझोरेस ग्रीष्म वेळ"},
	}
}

// Locale returns the current translators string locale
func (kok *kok_IN) Locale() string {
	return kok.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kok_IN'
func (kok *kok_IN) PluralsCardinal() []locales.PluralRule {
	return kok.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kok_IN'
func (kok *kok_IN) PluralsOrdinal() []locales.PluralRule {
	return kok.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kok_IN'
func (kok *kok_IN) PluralsRange() []locales.PluralRule {
	return kok.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kok_IN'
func (kok *kok_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kok_IN'
func (kok *kok_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kok_IN'
func (kok *kok_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kok *kok_IN) MonthAbbreviated(month time.Month) string {
	return kok.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kok *kok_IN) MonthsAbbreviated() []string {
	return kok.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kok *kok_IN) MonthNarrow(month time.Month) string {
	return kok.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kok *kok_IN) MonthsNarrow() []string {
	return kok.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kok *kok_IN) MonthWide(month time.Month) string {
	return kok.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kok *kok_IN) MonthsWide() []string {
	return kok.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kok *kok_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return kok.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kok *kok_IN) WeekdaysAbbreviated() []string {
	return kok.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kok *kok_IN) WeekdayNarrow(weekday time.Weekday) string {
	return kok.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kok *kok_IN) WeekdaysNarrow() []string {
	return kok.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kok *kok_IN) WeekdayShort(weekday time.Weekday) string {
	return kok.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kok *kok_IN) WeekdaysShort() []string {
	return kok.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kok *kok_IN) WeekdayWide(weekday time.Weekday) string {
	return kok.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kok *kok_IN) WeekdaysWide() []string {
	return kok.daysWide
}

// Decimal returns the decimal point of number
func (kok *kok_IN) Decimal() string {
	return kok.decimal
}

// Group returns the group of number
func (kok *kok_IN) Group() string {
	return kok.group
}

// Group returns the minus sign of number
func (kok *kok_IN) Minus() string {
	return kok.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kok_IN' and handles both Whole and Real numbers based on 'v'
func (kok *kok_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kok.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kok_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kok *kok_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kok.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kok_IN'
func (kok *kok_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kok.group[0])
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
		b = append(b, kok.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kok.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kok_IN'
// in accounting notation.
func (kok *kok_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kok.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kok.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kok.group[0])
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

		b = append(b, kok.minus[0])

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
			b = append(b, kok.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kok.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kok.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kok_IN'
func (kok *kok_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kok.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kok.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kok.periodsAbbreviated[0]...)
	} else {
		b = append(b, kok.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kok.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
