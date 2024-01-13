package messages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
  bundle.LoadMessageFile("locale/messages/messages.en.json")
  bundle.LoadMessageFile("locale/messages/messages.pt_BR.json")
}

func GetMessage(key string, lang string, args ...string) (string, error) {
	arguments := make(map[string]string)
	
	for i, arg := range args {
		arguments[fmt.Sprintf("Arg%d", i)] = arg
	}

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: key,
		TemplateData: arguments,
	}
	localizer = i18n.NewLocalizer(bundle, lang)  
	localizedMessage, err := localizer.Localize(&localizeConfigWelcome)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return localizedMessage, nil
}