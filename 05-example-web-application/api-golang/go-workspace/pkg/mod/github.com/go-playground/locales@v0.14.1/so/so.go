package so

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type so struct {
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

// New returns a new instance of translator for the 'so' locale
func New() locales.Translator {
	return &so{
		locale:                 "so",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "DBB", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "S", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "L", "O", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Bisha Koobaad", "Bisha Labaad", "Bisha Saddexaad", "Bisha Afraad", "Bisha Shanaad", "Bisha Lixaad", "Bisha Todobaad", "Bisha Sideedaad", "Bisha Sagaalaad", "Bisha Tobnaad", "Bisha Kow iyo Tobnaad", "Bisha Laba iyo Tobnaad"},
		daysAbbreviated:        []string{"Axd", "Isn", "Tldo", "Arbc", "Khms", "Jmc", "Sbti"},
		daysNarrow:             []string{"A", "I", "T", "A", "Kh", "J", "S"},
		daysShort:              []string{"Axd", "Isn", "Tldo", "Arbc", "Khms", "Jmc", "Sbti"},
		daysWide:               []string{"Axad", "Isniin", "Talaado", "Arbaco", "Khamiis", "Jimco", "Sabti"},
		periodsAbbreviated:     []string{"GH", "GD"},
		periodsNarrow:          []string{"h", "d"},
		periodsWide:            []string{"GH", "GD"},
		erasAbbreviated:        []string{"CH", "CD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Ciise Hortii", "Ciise Dabadii"},
		timezones:              map[string]string{"ACDT": "Waqtiga Dharaarta ee Bartamaha Astaraaliya", "ACST": "Waqtiga Caadiga Ah ee Bartamaha Astaraaliya", "ACWDT": "Waqtiga Dharaarta Bartamaha Galbeedka Australiya", "ACWST": "Waqtiga Caadiga Ah ee Bartamaha Galbeedka Astaraaliya", "ADT": "Waqtiga Dharaarta ee Atlantika Waqooyiga Ameerika", "AEDT": "Waqtiga Dharaarta ee Bariga Astaraaliya", "AEST": "Waqtiyada Caadiga ah ee Bariga Astaraaliya", "AKDT": "Waqtiga Dharaarta ee Alaska", "AKST": "Waqtiga Caadiga Ah ee Alaska", "ARST": "Waqtiga Xagaaga ee Arjentiina", "ART": "Waqtiga Caadiga Ah ee Arjentiina", "AST": "Waqtiga Caadiga Ah ee Atlantika Waqooyiga Ameerika", "AWDT": "Waqtiga Dharaarta ee Galbeedka Astaraaliya", "AWST": "Waqtiga Caadiga Ah ee Galbeedka Astaraaliya", "BOT": "Waqtiga Boliifiya", "BT": "Waqtiga Butaan", "CAT": "Waqtiga Bartamaha Afrika", "CDT": "Waqtiga Dharaarta ee Bartamaha Waqooyiga Ameerika", "CHADT": "Waqtiga Dharaarta ee Jaatam", "CHAST": "Waqtiga Caadiga Ah ee Jaatam", "CLST": "Waqtiga Xagaaga ee Jili", "CLT": "Waqtiga Caadiga Ah ee Jili", "COST": "Waqtiga Xagaaga ee Kolambiya", "COT": "Waqtiga Caadiga Ah ee Kolambiya", "CST": "Waqtiga Caadiga Ah ee Bartamaha Waqooyiga Ameerika", "ChST": "Waqtiga Jamoro", "EAT": "Waqtiga Bariga Afrika", "ECT": "Waqtiga Ekuwadoor", "EDT": "Waqtiga Dharaarta ee Bariga Waqooyiga Ameerika", "EST": "Waqtiga Caadiga Ah ee Bariga Waqooyiga Ameerika", "GFT": "Waqtiga Ferenj Guyana", "GMT": "Waqtiga Celceliska Giriinwij", "GST": "Waqtiga Gacanka", "GYT": "Waqtiga Guyaana", "HADT": "Waqtiga Dharaarta ee Hawaay-Alutiyaan", "HAST": "Waqtiga Caadiga Ah ee Hawaay-Alutiyaan", "HAT": "Waqtiga Dharaarta ee Niyuufoonlaan", "HECU": "Waqtiga Dharaarta ee Kuuba", "HEEG": "Waqtiga Xagaaga ee Bariga Giriinlaan", "HENOMX": "Waqtiga Dharaarta ee Waqooyi-Galbeed Meksiko", "HEOG": "Waqtiga Xagaaga ee Galbeedka Giriinlaan", "HEPM": "Waqtiga Dharaarta ee St. Beere & Mikiwelon", "HEPMX": "Waqtiga Dharaarta ee Baasifikada Meksiko", "HKST": "Waqtiga Xagaaga ee Hoong Koong", "HKT": "Waqtiga Caadiga Ah ee Hoong Koong", "HNCU": "Waqtiga Caadiga Ah ee Kuuba", "HNEG": "Waqtiga Caadiga ah ee Bariga Giriinlaan", "HNNOMX": "Waqtiga Caadiga Ah ee Waqooyi-Galbeed Meksiko", "HNOG": "Waqtiga Caadiga Ah ee Galbeedka Giriinlaan", "HNPM": "Waqtiga Caadiga Ah St. Beere & Mikiwelon", "HNPMX": "Waqtiga Caadiga Ah ee Baasifikada Meksiko", "HNT": "Waqtiga Caadiga Ah ee Niyuufoonlaan", "IST": "Waqtiga Caadiga Ah ee Hindiya", "JDT": "Waqtiga Dharaarta ee Jabaan", "JST": "Waqtiga Caadiga Ah ee Jabaan", "LHDT": "Waqtiga Dharaarta ee Lod How", "LHST": "Waqtiga Caadiga Ah ee Lod How", "MDT": "Waqtiga Dharaarta ee Buurleyda Waqooyiga Ameerika", "MESZ": "Waqtiga Xagaaga ee Bartamaha Yurub", "MEZ": "Waqtiga Caadiga Ah ee Bartamaha Yurub", "MST": "Waqtiga Caadiga ah ee Buuraleyda Waqooyiga Ameerika", "MYT": "Waqtiga Maleyshiya", "NZDT": "Waqtiga Dharaarta ee Niyuu Si’laan", "NZST": "Waqtiga Caadiga Ah ee Niyuu Si’laan", "OESZ": "Waqtiga Xagaaga ee Bariga Yurub", "OEZ": "Waqtiga Caadiga Ah ee Bariga Yurub", "PDT": "Waqtiga Dharaarta ee Basifika Waqooyiga Ameerika", "PST": "Waqtiga Caadiga ah ee Basifika Waqooyiga Ameerika", "SAST": "Waqtiyada Caadiga Ah ee Koonfur Afrika", "SGT": "Waqtiga Singabuur", "SRT": "Waqtiga Surineym", "TMST": "Waqtiga Xagaaga ee Turkmenistan", "TMT": "Waqtiga Caadiga Ah ee Turkmenistan", "UYST": "Waqtiga Xagaaga ee Urugwaay", "UYT": "Waqtiga Caadiga Ah ee Urugwaay", "VET": "Waqtiga Fenezuweela", "WARST": "Waqtiga Xagaaga ee Galbeedka Arjentiina", "WART": "Waqtiga Caadiga Ah ee Galbeedka Arjentiina", "WAST": "Waqtiga Xagaaga ee Galbeedka Afrika", "WAT": "Waqtiga Caadiga Ah ee Galbeedka Afrika", "WESZ": "Waqtiga Xagaaga ee Galbeedka Yurub", "WEZ": "Waqtiga Caadiga Ah ee Galbeedka Yurub", "WIB": "Waqtiga Galbeedka Indoneeysiya", "WIT": "Waqtiga Indoneeysiya", "WITA": "Waqtiga Bartamaha Indoneeysiya", "∅∅∅": "Waqtiga Xagaaga ee Asores"},
	}
}

// Locale returns the current translators string locale
func (so *so) Locale() string {
	return so.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'so'
func (so *so) PluralsCardinal() []locales.PluralRule {
	return so.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'so'
func (so *so) PluralsOrdinal() []locales.PluralRule {
	return so.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'so'
func (so *so) PluralsRange() []locales.PluralRule {
	return so.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'so'
func (so *so) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'so'
func (so *so) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'so'
func (so *so) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (so *so) MonthAbbreviated(month time.Month) string {
	return so.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (so *so) MonthsAbbreviated() []string {
	return so.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (so *so) MonthNarrow(month time.Month) string {
	return so.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (so *so) MonthsNarrow() []string {
	return so.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (so *so) MonthWide(month time.Month) string {
	return so.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (so *so) MonthsWide() []string {
	return so.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (so *so) WeekdayAbbreviated(weekday time.Weekday) string {
	return so.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (so *so) WeekdaysAbbreviated() []string {
	return so.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (so *so) WeekdayNarrow(weekday time.Weekday) string {
	return so.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (so *so) WeekdaysNarrow() []string {
	return so.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (so *so) WeekdayShort(weekday time.Weekday) string {
	return so.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (so *so) WeekdaysShort() []string {
	return so.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (so *so) WeekdayWide(weekday time.Weekday) string {
	return so.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (so *so) WeekdaysWide() []string {
	return so.daysWide
}

// Decimal returns the decimal point of number
func (so *so) Decimal() string {
	return so.decimal
}

// Group returns the group of number
func (so *so) Group() string {
	return so.group
}

// Group returns the minus sign of number
func (so *so) Minus() string {
	return so.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'so' and handles both Whole and Real numbers based on 'v'
func (so *so) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'so' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (so *so) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'so'
func (so *so) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := so.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, so.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, so.group[0])
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
		b = append(b, so.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, so.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'so'
// in accounting notation.
func (so *so) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := so.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, so.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, so.group[0])
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

		b = append(b, so.currencyNegativePrefix[0])

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
			b = append(b, so.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, so.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'so'
func (so *so) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'so'
func (so *so) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, so.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'so'
func (so *so) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, so.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'so'
func (so *so) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, so.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, so.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'so'
func (so *so) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'so'
func (so *so) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'so'
func (so *so) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'so'
func (so *so) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := so.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
