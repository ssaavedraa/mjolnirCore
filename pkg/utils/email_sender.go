package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/interfaces"
	"net/smtp"
	"text/template"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var auth smtp.Auth

type EmailSenderImpl struct {
	Config config.Config
}

func NewEmailSender(
	config config.Config,
) interfaces.EmailSender {
	return &EmailSenderImpl{
		Config: config,
	}
}

func (es *EmailSenderImpl) Send(email interfaces.Email) error {
	if auth == nil {
		auth = smtp.PlainAuth(
			"",
			es.Config.GetEnv("SMTP_USERNAME"),
			es.Config.GetEnv("SMTP_PASSWORD"),
			es.Config.GetEnv("SMTP_SERVER"),
		)
	}

	body, err := es.getEmailBody(email.TemplateName, email.Locale, email.TemplateData)

	if err != nil {
		return err
	}

	msg := []byte(
		"From: " + email.SenderAddress + "\r\n" +
			"To: " + email.ReceiverAddress + "\r\n" +
			"Subject: " + email.Subject + "\r\n" +
			"MIME-version: 1.0;\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
			body + "",
	)

	smtpAddress := fmt.Sprintf(
		"%s:%s",
		es.Config.GetEnv("SMTP_SERVER"),
		es.Config.GetEnv("SMTP_PORT"),
	)

	err = smtp.SendMail(
		smtpAddress,
		auth,
		email.SenderAddress,
		[]string{email.ReceiverAddress},
		msg,
	)

	if err != nil {
		return err
	}

	return nil
}

func (es *EmailSenderImpl) getEmailBody(templateName, locale string, templateData map[string]string) (string, error) {
	localizationKeys, err := es.getLocalizationKeys(templateName)

	if err != nil {
		return "", err
	}

	localizer, err := es.getI18nLocalizer(templateName, locale)

	if err != nil {
		return "", nil
	}

	templateValues := make(map[string]string)

	for field, key := range templateData {
		templateValues[field] = key
	}

	for field, key := range localizationKeys {
		localizedString := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: key})

		templateValues[field] = localizedString
	}

	templatePath := fmt.Sprintf("templates/%s/index.html", templateName)
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, templateValues); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (es *EmailSenderImpl) getLocalizationKeys(templateName string) (map[string]string, error) {
	switch templateName {
	case "user_invite_mvp":
		templateFields := map[string]string{
			"Subject":        "WelcomeSubject",
			"WelcomeMessage": "WelcomeMessage",
			"Greeting":       "Greeting",
			"Intro":          "Intro",
			"Participation":  "Participation",
			"Expectations":   "Expectations",
			"Benefit1":       "Benefit1",
			"Benefit2":       "Benefit2",
			"Benefit3":       "Benefit3",
			"Benefit4":       "Benefit4",
			"ReachOut":       "ReachOut",
			"WelcomeAboard":  "WelcomeAboard",
			"BestRegards":    "BestRegards",
			"Founders":       "Founders",
			"FoundersTitle":  "FoundersTitle",
			"JoinUs":         "JoinUs",
		}

		return templateFields, nil

	default:
		return map[string]string{}, errors.New("invalid email template")
	}
}

func (es *EmailSenderImpl) getI18nLocalizer(templateName, locale string) (*i18n.Localizer, error) {
	var i18nLocale language.Tag

	if locale == "es" {
		i18nLocale = language.Spanish
	} else {
		i18nLocale = language.English
	}

	bundle := i18n.NewBundle(i18nLocale)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	i18nFilePath := fmt.Sprintf("i18n/%s/%s.json", templateName, &i18nLocale)
	_, err := bundle.LoadMessageFile(i18nFilePath)

	if err != nil {
		return &i18n.Localizer{}, err
	}

	localizer := i18n.NewLocalizer(bundle, locale)

	return localizer, nil
}
