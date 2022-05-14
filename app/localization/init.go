package localization

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.quick.start/cache"
	"golang.org/x/text/language"
)

func SetLocale(lang string) {
	cache.Set("locale", lang, 0)
}

func GetLocale(lang string) string {
	return cache.Get("locale")
}

func initLocalizer(langs ...string) *i18n.Localizer {
	// Create a new i18n bundle with English as default language.
	bundle := i18n.NewBundle(language.English)

	// Register a json unmarshal function for i18n bundle.
	// This is to enable usage of json format
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load source language
	bundle.LoadMessageFile("./locales/en.json")
	bundle.LoadMessageFile("./locales/es.json")

	// Initialize localizer which will look for phrase keys in passed languages
	// in a strict order (first language is searched first)
	// When no key in any of the languages is found, it fallbacks to default - English language
	if langs == nil {
		cache.Set("locale", "en", 0)
	}
	localizer := i18n.NewLocalizer(bundle, langs...)

	return localizer
}

func Trans(key string, vars string) string {
	localize := initLocalizer(GetLocale("en"))
	var x map[string]interface{}

	json.Unmarshal([]byte(vars), &x)
	simpleMessage, _ := localize.Localize(&i18n.LocalizeConfig{
		MessageID:    key, // source key identifier
		TemplateData: x,
	})
	fmt.Println(simpleMessage)
	return simpleMessage
}

func SetPlural(key string, pluralCount int) string {
	localize := initLocalizer(GetLocale("en"))

	pluralMessage, _ := localize.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: key}, // another source key identifier
		PluralCount:    pluralCount,            // would use "one" variant if the count was 1
	})

	fmt.Println(pluralMessage)
	return pluralMessage
}
