package config

import (
	"github.com/spf13/viper"
)

var Instance = &config{}

type config struct {
	Environment          string
	DatabaseURL          string
	TransactionsFilePath string
	EmailTo              string
	EmailFrom            string
	SmtpServer           string
	SmtpAddr             string
	SmtpUser             string
	SmtpPassword         string
	SmtpHost             string
}

func init() {
	viper.AutomaticEnv()

	Instance.Environment = viper.GetString("ENVIRONMENT")
	Instance.DatabaseURL = viper.GetString("DATABASE_URL")
	Instance.TransactionsFilePath = "./files/account_details.csv"
	Instance.EmailTo = viper.GetString("EMAIL_TO")
	Instance.EmailFrom = viper.GetString("EMAIL_FROM")
	Instance.SmtpServer = viper.GetString("SMTP_SERVER")
	Instance.SmtpAddr = viper.GetString("SMTP_ADDR")
	Instance.SmtpUser = viper.GetString("SMTP_USER")
	Instance.SmtpPassword = viper.GetString("SMTP_PASSWORD")
	Instance.SmtpHost = viper.GetString("SMTP_HOST")
}
