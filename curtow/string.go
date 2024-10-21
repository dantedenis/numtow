package curtow

import (
	"github.com/dantedenis/numtow/lang"
	"github.com/dantedenis/numtow/lang/en"
	"github.com/dantedenis/numtow/lang/kz"
	"github.com/dantedenis/numtow/lang/ru"
)

// String converts amount to words
func String(amount string, language lang.Lang, options ...interface{}) (words string, err error) {
	switch language {
	case lang.KZ:
		o := kz.ParseCurrencyOpts(options...)

		return kz.CurrencyString(amount, o...)
	case lang.RU:
		o := ru.ParseCurrencyOpts(options...)

		return ru.CurrencyString(amount, o...)
	case lang.EN:
		o := en.ParseCurrencyOpts(options...)

		return en.CurrencyString(amount, o...)
	case lang.Unknown:
		return words, lang.ErrBadLanguage
	default:
		return words, lang.ErrBadLanguage
	}
}

// MustString converts amount to words, on error returns empty string.
func MustString(amount string, language lang.Lang, options ...interface{}) string {
	words, err := String(amount, language, options...)
	if err != nil {
		return ""
	}

	return words
}
