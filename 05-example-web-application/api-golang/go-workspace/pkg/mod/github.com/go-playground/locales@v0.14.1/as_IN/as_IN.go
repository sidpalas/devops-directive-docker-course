package as_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type as_IN struct {
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

// New returns a new instance of translator for the 'as_IN' locale
func New() locales.Translator {
	return &as_IN{
		locale:                 "as_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
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
		monthsAbbreviated:      []string{"", "জানু", "ফেব্ৰু", "মাৰ্চ", "এপ্ৰিল", "মে’", "জুন", "জুলাই", "আগ", "ছেপ্তে", "অক্টো", "নৱে", "ডিচে"},
		monthsNarrow:           []string{"", "জ", "ফ", "ম", "এ", "ম", "জ", "জ", "আ", "ছ", "অ", "ন", "ড"},
		monthsWide:             []string{"", "জানুৱাৰী", "ফেব্ৰুৱাৰী", "মাৰ্চ", "এপ্ৰিল", "মে’", "জুন", "জুলাই", "আগষ্ট", "ছেপ্তেম্বৰ", "অক্টোবৰ", "নৱেম্বৰ", "ডিচেম্বৰ"},
		daysAbbreviated:        []string{"দেও", "সোম", "মঙ্গল", "বুধ", "বৃহ", "শুক্ৰ", "শনি"},
		daysNarrow:             []string{"দ", "স", "ম", "ব", "ব", "শ", "শ"},
		daysShort:              []string{"দেও", "সোম", "মঙ্গল", "বুধ", "বৃহ", "শুক্ৰ", "শনি"},
		daysWide:               []string{"দেওবাৰ", "সোমবাৰ", "মঙ্গলবাৰ", "বুধবাৰ", "বৃহস্পতিবাৰ", "শুক্ৰবাৰ", "শনিবাৰ"},
		periodsAbbreviated:     []string{"পূৰ্বাহ্ন", "অপৰাহ্ন"},
		periodsNarrow:          []string{"পূৰ্বাহ্ন", "অপৰাহ্ন"},
		periodsWide:            []string{"পূৰ্বাহ্ন", "অপৰাহ্ন"},
		erasAbbreviated:        []string{"খ্ৰীঃ পূঃ", "খ্ৰীঃ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"খ্ৰীষ্টপূৰ্ব", "খ্ৰীষ্টাব্দ"},
		timezones:              map[string]string{"ACDT": "অষ্ট্ৰেলিয়াৰ কেন্দ্ৰীয় ডেলাইট সময়", "ACST": "অষ্ট্ৰেলিয়াৰ কেন্দ্ৰীয় মান সময়", "ACWDT": "অষ্ট্ৰেলিয়াৰ কেন্দ্ৰীয় পাশ্চাত্য ডেলাইট সময়", "ACWST": "অষ্ট্ৰেলিয়াৰ কেন্দ্ৰীয় পাশ্চাত্য মান সময়", "ADT": "আটলাণ্টিক ডেলাইট সময়", "AEDT": "অষ্ট্ৰেলিয়াৰ প্ৰাচ্য ডেলাইট সময়", "AEST": "অষ্ট্ৰেলিয়াৰ প্ৰাচ্য মান সময়", "AKDT": "আলাস্কাৰ ডেলাইট সময়", "AKST": "আলাস্কাৰ মান সময়", "ARST": "আৰ্জেণ্টিনাৰ গ্ৰীষ্মকালীন সময়", "ART": "আৰ্জেণ্টিনাৰ মান সময়", "AST": "আটলাণ্টিক মান সময়", "AWDT": "অষ্ট্ৰেলিয়াৰ পাশ্চাত্য ডেলাইট সময়", "AWST": "অষ্ট্ৰেলিয়াৰ পাশ্চাত্য মান সময়", "BOT": "বলিভিয়াৰ সময়", "BT": "ভুটানৰ সময়", "CAT": "মধ্য আফ্ৰিকাৰ সময়", "CDT": "উত্তৰ আমেৰিকাৰ কেন্দ্ৰীয় ডেলাইট সময়", "CHADT": "চাথামৰ ডেলাইট সময়", "CHAST": "চাথামৰ মান সময়", "CLST": "চিলিৰ গ্ৰীষ্মকালীন সময়", "CLT": "চিলিৰ মান সময়", "COST": "কলম্বিয়াৰ গ্ৰীষ্মকালীন সময়", "COT": "কলম্বিয়াৰ মান সময়", "CST": "উত্তৰ আমেৰিকাৰ কেন্দ্ৰীয় মান সময়", "ChST": "চামোৰোৰ মান সময়", "EAT": "পূব আফ্ৰিকাৰ সময়", "ECT": "ইকুৱেডৰৰ সময়", "EDT": "উত্তৰ আমেৰিকাৰ প্ৰাচ্য ডেলাইট সময়", "EST": "উত্তৰ আমেৰিকাৰ প্ৰাচ্য মান সময়", "GFT": "ফ্ৰান্স গয়ানাৰ সময়", "GMT": "গ্ৰীণউইচ মান সময়", "GST": "উপসাগৰীয় মান সময়", "GYT": "গায়ানাৰ সময়", "HADT": "হাৱাই-এলিউশ্বনৰ ডেলাইট সময়", "HAST": "হাৱাই-এলিউশ্বনৰ মান সময়", "HAT": "নিউফাউণ্ডলেণ্ডৰ ডেলাইট সময়", "HECU": "কিউবাৰ ডেলাইট সময়", "HEEG": "পূব গ্ৰীণলেণ্ডৰ গ্ৰীষ্মকালীন সময়", "HENOMX": "উত্তৰ-পশ্চিম মেক্সিকোৰ ডেলাইট সময়", "HEOG": "পশ্চিম গ্ৰীণলেণ্ডৰ গ্ৰীষ্মকালীন সময়", "HEPM": "ছেইণ্ট পিয়েৰে আৰু মিকিউৱেলনৰ ডেলাইট সময়", "HEPMX": "মেক্সিকোৰ প্ৰশান্ত ডেলাইট সময়", "HKST": "হং কঙৰ গ্ৰীষ্মকালীন সময়", "HKT": "হং কঙৰ মান সময়", "HNCU": "কিউবাৰ মান সময়", "HNEG": "পূব গ্ৰীণলেণ্ডৰ মান সময়", "HNNOMX": "উত্তৰ-পশ্চিম মেক্সিকোৰ মান সময়", "HNOG": "পশ্চিম গ্ৰীণলেণ্ডৰ মান সময়", "HNPM": "ছেইণ্ট পিয়েৰে আৰু মিকিউৱেলনৰ মান সময়", "HNPMX": "মেক্সিকোৰ প্ৰশান্ত মান সময়", "HNT": "নিউফাউণ্ডলেণ্ডৰ মান সময়", "IST": "ভাৰতীয় মান সময়", "JDT": "জাপানৰ ডেলাইট সময়", "JST": "জাপানৰ মান সময়", "LHDT": "লৰ্ড হাওৰ ডেলাইট সময়", "LHST": "লৰ্ড হাওৰ মান সময়", "MDT": "উত্তৰ আমেৰিকাৰ পৰ্ব্বতীয় ডেলাইট সময়", "MESZ": "মধ্য ইউৰোপীয় গ্ৰীষ্মকালীন সময়", "MEZ": "মধ্য ইউৰোপীয় মান সময়", "MST": "উত্তৰ আমেৰিকাৰ পৰ্ব্বতীয় মান সময়", "MYT": "মালয়েচিয়াৰ সময়", "NZDT": "নিউজিলেণ্ডৰ ডেলাইট সময়", "NZST": "নিউজিলেণ্ডৰ মান সময়", "OESZ": "প্ৰাচ্য ইউৰোপীয় গ্ৰীষ্মকালীন সময়", "OEZ": "প্ৰাচ্য ইউৰোপীয় মান সময়", "PDT": "উত্তৰ আমেৰিকাৰ ডেলাইট সময়", "PST": "উত্তৰ আমেৰিকাৰ প্ৰশান্ত মান সময়", "SAST": "দক্ষিণ আফ্ৰিকাৰ মান সময়", "SGT": "ছিংগাপুৰৰ মান সময়", "SRT": "ছুৰিনামৰ সময়", "TMST": "তুৰ্কমেনিস্তানৰ গ্ৰীষ্মকালীন সময়", "TMT": "তুৰ্কমেনিস্তানৰ মান সময়", "UYST": "উৰুগুৱেৰ গ্ৰীষ্মকালীন সময়", "UYT": "উৰুগুৱেৰ মান সময়", "VET": "ভেনিজুৱেলাৰ সময়", "WARST": "পাশ্চাত্য আৰ্জেণ্টিনাৰ গ্ৰীষ্মকালীন সময়", "WART": "পাশ্চাত্য আৰ্জেণ্টিনাৰ মান সময়", "WAST": "পশ্চিম আফ্ৰিকাৰ গ্ৰীষ্মকালীন সময়", "WAT": "পশ্চিম আফ্ৰিকাৰ মান সময়", "WESZ": "পাশ্চাত্য ইউৰোপীয় গ্ৰীষ্মকালীন সময়", "WEZ": "পাশ্চাত্য ইউৰোপীয় মান সময়", "WIB": "পাশ্চাত্য ইণ্ডোনেচিয়াৰ সময়", "WIT": "প্ৰাচ্য ইণ্ডোনেচিয়াৰ সময়", "WITA": "মধ্য ইণ্ডোনেচিয়াৰ সময়", "∅∅∅": "আজোৰেছৰ গ্ৰীষ্মকালীন সময়"},
	}
}

// Locale returns the current translators string locale
func (as *as_IN) Locale() string {
	return as.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'as_IN'
func (as *as_IN) PluralsCardinal() []locales.PluralRule {
	return as.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'as_IN'
func (as *as_IN) PluralsOrdinal() []locales.PluralRule {
	return as.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'as_IN'
func (as *as_IN) PluralsRange() []locales.PluralRule {
	return as.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'as_IN'
func (as *as_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'as_IN'
func (as *as_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 5 || n == 7 || n == 8 || n == 9 || n == 10 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'as_IN'
func (as *as_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := as.CardinalPluralRule(num1, v1)
	end := as.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (as *as_IN) MonthAbbreviated(month time.Month) string {
	return as.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (as *as_IN) MonthsAbbreviated() []string {
	return as.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (as *as_IN) MonthNarrow(month time.Month) string {
	return as.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (as *as_IN) MonthsNarrow() []string {
	return as.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (as *as_IN) MonthWide(month time.Month) string {
	return as.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (as *as_IN) MonthsWide() []string {
	return as.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (as *as_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return as.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (as *as_IN) WeekdaysAbbreviated() []string {
	return as.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (as *as_IN) WeekdayNarrow(weekday time.Weekday) string {
	return as.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (as *as_IN) WeekdaysNarrow() []string {
	return as.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (as *as_IN) WeekdayShort(weekday time.Weekday) string {
	return as.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (as *as_IN) WeekdaysShort() []string {
	return as.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (as *as_IN) WeekdayWide(weekday time.Weekday) string {
	return as.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (as *as_IN) WeekdaysWide() []string {
	return as.daysWide
}

// Decimal returns the decimal point of number
func (as *as_IN) Decimal() string {
	return as.decimal
}

// Group returns the group of number
func (as *as_IN) Group() string {
	return as.group
}

// Group returns the minus sign of number
func (as *as_IN) Minus() string {
	return as.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'as_IN' and handles both Whole and Real numbers based on 'v'
func (as *as_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, as.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, as.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, as.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'as_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (as *as_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, as.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, as.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, as.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'as_IN'
func (as *as_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := as.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, as.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, as.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(as.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, as.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, as.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, as.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'as_IN'
// in accounting notation.
func (as *as_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := as.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, as.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, as.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(as.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, as.currencyNegativePrefix[j])
		}

		b = append(b, as.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(as.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, as.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, as.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'as_IN'
func (as *as_IN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'as_IN'
func (as *as_IN) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'as_IN'
func (as *as_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, as.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'as_IN'
func (as *as_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, as.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, as.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'as_IN'
func (as *as_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'as_IN'
func (as *as_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'as_IN'
func (as *as_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'as_IN'
func (as *as_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := as.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
