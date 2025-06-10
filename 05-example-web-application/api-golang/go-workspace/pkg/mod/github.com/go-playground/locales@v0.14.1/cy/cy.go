package cy

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type cy struct {
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

// New returns a new instance of translator for the 'cy' locale
func New() locales.Translator {
	return &cy{
		locale:                 "cy",
		pluralsCardinal:        []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 3, 4, 5, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ion", "Chwef", "Maw", "Ebr", "Mai", "Meh", "Gorff", "Awst", "Medi", "Hyd", "Tach", "Rhag"},
		monthsNarrow:           []string{"", "I", "Ch", "M", "E", "M", "M", "G", "A", "M", "H", "T", "Rh"},
		monthsWide:             []string{"", "Ionawr", "Chwefror", "Mawrth", "Ebrill", "Mai", "Mehefin", "Gorffennaf", "Awst", "Medi", "Hydref", "Tachwedd", "Rhagfyr"},
		daysAbbreviated:        []string{"Sul", "Llun", "Maw", "Mer", "Iau", "Gwen", "Sad"},
		daysNarrow:             []string{"S", "Ll", "M", "M", "I", "G", "S"},
		daysShort:              []string{"Su", "Ll", "Ma", "Me", "Ia", "Gw", "Sa"},
		daysWide:               []string{"Dydd Sul", "Dydd Llun", "Dydd Mawrth", "Dydd Mercher", "Dydd Iau", "Dydd Gwener", "Dydd Sadwrn"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"b", "h"},
		periodsWide:            []string{"yb", "yh"},
		erasAbbreviated:        []string{"CC", "OC"},
		erasNarrow:             []string{"C", "O"},
		erasWide:               []string{"Cyn Crist", "Oed Crist"},
		timezones:              map[string]string{"ACDT": "Amser Haf Canolbarth Awstralia", "ACST": "Amser Safonol Canolbarth Awstralia", "ACWDT": "Amser Haf Canolbarth Gorllewin Awstralia", "ACWST": "Amser Safonol Canolbarth Gorllewin Awstralia", "ADT": "Amser Haf Cefnfor yr Iwerydd", "AEDT": "Amser Haf Dwyrain Awstralia", "AEST": "Amser Safonol Dwyrain Awstralia", "AKDT": "Amser Haf Alaska", "AKST": "Amser Safonol Alaska", "ARST": "Amser Haf Ariannin", "ART": "Amser Safonol Ariannin", "AST": "Amser Safonol Cefnfor yr Iwerydd", "AWDT": "Amser Haf Gorllewin Awstralia", "AWST": "Amser Safonol Gorllewin Awstralia", "BOT": "Amser Bolivia", "BT": "Amser Bhutan", "CAT": "Amser Canolbarth Affrica", "CDT": "Amser Haf Canolbarth Gogledd America", "CHADT": "Amser Haf Chatham", "CHAST": "Amser Safonol Chatham", "CLST": "Amser Haf Chile", "CLT": "Amser Safonol Chile", "COST": "Amser Haf Colombia", "COT": "Amser Safonol Colombia", "CST": "Amser Safonol Canolbarth Gogledd America", "ChST": "Amser Chamorro", "EAT": "Amser Dwyrain Affrica", "ECT": "Amser Ecuador", "EDT": "Amser Haf Dwyrain Gogledd America", "EST": "Amser Safonol Dwyrain Gogledd America", "GFT": "Amser Guyane Ffrengig", "GMT": "Amser Safonol Greenwich", "GST": "Amser Safonol y Gwlff", "GYT": "Amser Guyana", "HADT": "Amser Haf Hawaii-Aleutian", "HAST": "Amser Safonol Hawaii-Aleutian", "HAT": "Amser Haf Newfoundland", "HECU": "Amser Haf Cuba", "HEEG": "Amser Haf Dwyrain yr Ynys Las", "HENOMX": "Amser Haf Gogledd Orllewin Mecsico", "HEOG": "Amser Haf Gorllewin yr Ynys Las", "HEPM": "Amser Haf Saint-Pierre-et-Miquelon", "HEPMX": "Amser Haf Pasiffig Mecsico", "HKST": "Amser Haf Hong Kong", "HKT": "Amser Safonol Hong Kong", "HNCU": "Amser Safonol Cuba", "HNEG": "Amser Safonol Dwyrain yr Ynys Las", "HNNOMX": "Amser Safonol Gogledd Orllewin Mecsico", "HNOG": "Amser Safonol Gorllewin yr Ynys Las", "HNPM": "Amser Safonol Saint-Pierre-et-Miquelon", "HNPMX": "Amser Safonol Pasiffig Mecsico", "HNT": "Amser Safonol Newfoundland", "IST": "Amser India", "JDT": "Amser Haf Siapan", "JST": "Amser Safonol Siapan", "LHDT": "Amser Haf yr Arglwydd Howe", "LHST": "Amser Safonol yr Arglwydd Howe", "MDT": "Amser Haf Mynyddoedd Gogledd America", "MESZ": "Amser Haf Canolbarth Ewrop", "MEZ": "Amser Safonol Canolbarth Ewrop", "MST": "Amser Safonol Mynyddoedd Gogledd America", "MYT": "Amser Malaysia", "NZDT": "Amser Haf Seland Newydd", "NZST": "Amser Safonol Seland Newydd", "OESZ": "Amser Haf Dwyrain Ewrop", "OEZ": "Amser Safonol Dwyrain Ewrop", "PDT": "Amser Haf Cefnfor Tawel Gogledd America", "PST": "Amser Safonol Cefnfor Tawel Gogledd America", "SAST": "Amser Safonol De Affrica", "SGT": "Amser Singapore", "SRT": "Amser Suriname", "TMST": "Amser Haf Tyrcmenistan", "TMT": "Amser Safonol Tyrcmenistan", "UYST": "Amser Haf Uruguay", "UYT": "Amser Safonol Uruguay", "VET": "Amser Venezuela", "WARST": "Amser Haf Gorllewin Ariannin", "WART": "Amser Safonol Gorllewin Ariannin", "WAST": "Amser Haf Gorllewin Affrica", "WAT": "Amser Safonol Gorllewin Affrica", "WESZ": "Amser Haf Gorllewin Ewrop", "WEZ": "Amser Safonol Gorllewin Ewrop", "WIB": "Amser Gorllewin Indonesia", "WIT": "Amser Dwyrain Indonesia", "WITA": "Amser Canolbarth Indonesia", "∅∅∅": "Amser Haf Amazonas"},
	}
}

// Locale returns the current translators string locale
func (cy *cy) Locale() string {
	return cy.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cy'
func (cy *cy) PluralsCardinal() []locales.PluralRule {
	return cy.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cy'
func (cy *cy) PluralsOrdinal() []locales.PluralRule {
	return cy.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cy'
func (cy *cy) PluralsRange() []locales.PluralRule {
	return cy.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cy'
func (cy *cy) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cy'
func (cy *cy) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 || n == 7 || n == 8 || n == 9 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 || n == 4 {
		return locales.PluralRuleFew
	} else if n == 5 || n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cy'
func (cy *cy) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := cy.CardinalPluralRule(num1, v1)
	end := cy.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (cy *cy) MonthAbbreviated(month time.Month) string {
	return cy.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (cy *cy) MonthsAbbreviated() []string {
	return cy.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (cy *cy) MonthNarrow(month time.Month) string {
	return cy.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (cy *cy) MonthsNarrow() []string {
	return cy.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (cy *cy) MonthWide(month time.Month) string {
	return cy.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (cy *cy) MonthsWide() []string {
	return cy.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (cy *cy) WeekdayAbbreviated(weekday time.Weekday) string {
	return cy.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (cy *cy) WeekdaysAbbreviated() []string {
	return cy.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (cy *cy) WeekdayNarrow(weekday time.Weekday) string {
	return cy.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (cy *cy) WeekdaysNarrow() []string {
	return cy.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (cy *cy) WeekdayShort(weekday time.Weekday) string {
	return cy.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (cy *cy) WeekdaysShort() []string {
	return cy.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (cy *cy) WeekdayWide(weekday time.Weekday) string {
	return cy.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (cy *cy) WeekdaysWide() []string {
	return cy.daysWide
}

// Decimal returns the decimal point of number
func (cy *cy) Decimal() string {
	return cy.decimal
}

// Group returns the group of number
func (cy *cy) Group() string {
	return cy.group
}

// Group returns the minus sign of number
func (cy *cy) Minus() string {
	return cy.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'cy' and handles both Whole and Real numbers based on 'v'
func (cy *cy) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'cy' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (cy *cy) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, cy.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'cy'
func (cy *cy) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cy.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
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
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'cy'
// in accounting notation.
func (cy *cy) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cy.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
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

		b = append(b, cy.currencyNegativePrefix[0])

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
			b = append(b, cy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, cy.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'cy'
func (cy *cy) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'cy'
func (cy *cy) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'cy'
func (cy *cy) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'cy'
func (cy *cy) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, cy.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'cy'
func (cy *cy) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cy'
func (cy *cy) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'cy'
func (cy *cy) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'cy'
func (cy *cy) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := cy.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
