package ps_AF

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ps_AF struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ps_AF' locale
func New() locales.Translator {
	return &ps_AF{
		locale:                 "ps_AF",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		monthsNarrow:           []string{"", "ج", "ف", "م", "ا", "م", "ج", "ج", "ا", "س", "ا", "ن", "د"},
		monthsWide:             []string{"", "جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysAbbreviated:        []string{"يونۍ", "دونۍ", "درېنۍ", "څلرنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		daysNarrow:             []string{"S", "M", "T", "W", "T", "F", "S"},
		daysShort:              []string{"يونۍ", "دونۍ", "درېنۍ", "څلرنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		daysWide:               []string{"يونۍ", "دونۍ", "درېنۍ", "څلرنۍ", "پينځنۍ", "جمعه", "اونۍ"},
		periodsAbbreviated:     []string{"غ.م.", "غ.و."},
		periodsNarrow:          []string{"غ.م.", "غ.و."},
		periodsWide:            []string{"غ.م.", "غ.و."},
		erasAbbreviated:        []string{"له میلاد وړاندې", "م."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"له میلاد څخه وړاندې", "له میلاد څخه وروسته"},
		timezones:              map[string]string{"ACDT": "آسترالوي مرکزي د ورځې روښانه وخت", "ACST": "آسترالوي مرکزي معياري وخت", "ACWDT": "آسترالوي مرکزي لوېديځ د ورځې روښانه وخت", "ACWST": "آسترالوي مرکزي لوېديځ معياري وخت", "ADT": "اتلانتیک د رڼا ورځې وخت", "AEDT": "آسترالوي ختيځ د ورځې روښانه وخت", "AEST": "آسترالوي ختيځ معياري وخت", "AKDT": "الاسکا د ورځې روښانه وخت", "AKST": "الاسکا معياري وخت", "ARST": "ارجنټاین اوړي وخت", "ART": "ارجنټاین معیاری وخت", "AST": "اتلانتیک معياري وخت", "AWDT": "آسترالوي لوېديځ د ورځې روښانه وخت", "AWST": "آسترالوي لوېديځ معياري وخت", "BOT": "بولیویا وخت", "BT": "بهوټان وخت", "CAT": "منځنی افريقا وخت", "CDT": "مرکزي د ورځې روښانه وخت", "CHADT": "چاتام د ورځې روښانه وخت", "CHAST": "چاتام معياري وخت", "CLST": "چلی اوړي وخت", "CLT": "چلی معیاری وخت", "COST": "کولمبیا اوړي وخت", "COT": "کولمبیا معیاری وخت", "CST": "مرکزي معياري وخت", "ChST": "چمارو معياري وخت", "EAT": "ختيځ افريقا وخت", "ECT": "د اکوادور وخت", "EDT": "ختيځ د رڼا ورځې وخت", "EST": "ختيځ معياري وخت", "GFT": "د فرانسوي ګانا وخت", "GMT": "ګرينويچ معياري وخت", "GST": "خلیج معياري وخت", "GYT": "د ګوانانا وخت", "HADT": "هوایی الیوتین رڼا ورځې وخت", "HAST": "هوایی الیوتین معیاری وخت", "HAT": "نيو فاونډلېنډ د ورځې روښانه وخت", "HECU": "کیوبا د رڼا ورځې وخت", "HEEG": "د ختیځ ګرینلینډ اوړي وخت", "HENOMX": "د شمال لویدیځ مکسیکو رڼا ورځې وخت", "HEOG": "لویدیځ ګرینلینډ اوړي وخت", "HEPM": "سینټ پییرا و ميکلين رڼا ورځې وخت", "HEPMX": "مکسیکن پیسفک رڼا ورځې وخت", "HKST": "هانګ کانګ اوړي وخت", "HKT": "هانګ کانګ معياري وخت", "HNCU": "کیوبا معياري وخت", "HNEG": "د ختیځ ګرینلینډ معياري وخت", "HNNOMX": "د شمال لویدیځ مکسیکو معیاري وخت", "HNOG": "لویدیځ ګرینلینډ معياري وخت", "HNPM": "سینټ پییرا و ميکلين معیاری وخت", "HNPMX": "مکسیکن پیسفک معیاری وخت", "HNT": "د نوي فیلډلینډ معیاری وخت", "IST": "هند معیاري وخت", "JDT": "جاپان د ورځې روښانه وخت", "JST": "جاپان معياري وخت", "LHDT": "لارډ هوي د ورځې روښانه وخت", "LHST": "لارډ هوي معياري وخت", "MDT": "د غره د رڼا ورځې وخت", "MESZ": "وسطي اروپايي د اوړي وخت", "MEZ": "د مرکزي اروپا معیاري وخت", "MST": "د غره معياري وخت", "MYT": "ملائیشیا وخت", "NZDT": "نيوزي لېنډ د ورځې روښانه وخت", "NZST": "نيوزي لېنډ معياري وخت", "OESZ": "ختيځ اروپايي اوړي وخت", "OEZ": "ختيځ اروپايي معياري وخت", "PDT": "پیسفک د رڼا ورځې وخت", "PST": "د پیسفک معياري وخت", "SAST": "جنوبي افريقا معياري وخت", "SGT": "سنګاپور معیاري وخت", "SRT": "سورینام وخت", "TMST": "ترکمنستان اوړي وخت", "TMT": "ترکمنستان معياري وخت", "UYST": "یوروګوای اوړي وخت", "UYT": "یوروګوای معياري وخت", "VET": "وینزویلا وخت", "WARST": "لوېديځ ارجنټاين اوړي وخت", "WART": "لوېديځ ارجنټاين معياري وخت", "WAST": "د افریقا افریقا لویدیځ وخت", "WAT": "لویدیځ افریقایي معیاري وخت", "WESZ": "لوېديځ اروپايي اوړي وخت", "WEZ": "لوېديځ اروپايي معياري وخت", "WIB": "لویدیځ اندونیزیا وخت", "WIT": "اندونیزیا وخت", "WITA": "مرکزي ادونيزيا وخت", "∅∅∅": "برسلیا اوړي وخت"},
	}
}

// Locale returns the current translators string locale
func (ps *ps_AF) Locale() string {
	return ps.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ps_AF'
func (ps *ps_AF) PluralsCardinal() []locales.PluralRule {
	return ps.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ps_AF'
func (ps *ps_AF) PluralsOrdinal() []locales.PluralRule {
	return ps.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ps_AF'
func (ps *ps_AF) PluralsRange() []locales.PluralRule {
	return ps.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ps_AF'
func (ps *ps_AF) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ps_AF'
func (ps *ps_AF) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ps_AF'
func (ps *ps_AF) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ps.CardinalPluralRule(num1, v1)
	end := ps.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ps *ps_AF) MonthAbbreviated(month time.Month) string {
	return ps.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ps *ps_AF) MonthsAbbreviated() []string {
	return ps.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ps *ps_AF) MonthNarrow(month time.Month) string {
	return ps.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ps *ps_AF) MonthsNarrow() []string {
	return ps.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ps *ps_AF) MonthWide(month time.Month) string {
	return ps.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ps *ps_AF) MonthsWide() []string {
	return ps.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ps *ps_AF) WeekdayAbbreviated(weekday time.Weekday) string {
	return ps.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ps *ps_AF) WeekdaysAbbreviated() []string {
	return ps.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ps *ps_AF) WeekdayNarrow(weekday time.Weekday) string {
	return ps.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ps *ps_AF) WeekdaysNarrow() []string {
	return ps.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ps *ps_AF) WeekdayShort(weekday time.Weekday) string {
	return ps.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ps *ps_AF) WeekdaysShort() []string {
	return ps.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ps *ps_AF) WeekdayWide(weekday time.Weekday) string {
	return ps.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ps *ps_AF) WeekdaysWide() []string {
	return ps.daysWide
}

// Decimal returns the decimal point of number
func (ps *ps_AF) Decimal() string {
	return ps.decimal
}

// Group returns the group of number
func (ps *ps_AF) Group() string {
	return ps.group
}

// Group returns the minus sign of number
func (ps *ps_AF) Minus() string {
	return ps.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ps_AF' and handles both Whole and Real numbers based on 'v'
func (ps *ps_AF) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ps_AF' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ps *ps_AF) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ps.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ps.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ps_AF'
func (ps *ps_AF) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ps.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ps.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ps.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ps.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ps.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ps_AF'
// in accounting notation.
func (ps *ps_AF) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ps.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ps.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ps.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ps.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ps.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ps.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ps.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ps.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0xd8, 0xaf, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd8, 0xaf, 0x20}...)
	b = append(b, ps.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ps.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0xd8, 0xaf, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd8, 0xaf, 0x20}...)
	b = append(b, ps.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ps_AF'
func (ps *ps_AF) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ps.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := ps.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
