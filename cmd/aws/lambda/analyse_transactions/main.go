package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/manicar2093/filestores"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/connections"
	"github.com/manicar2093/stori-challenge/pkg/converters"
	"github.com/manicar2093/stori-challenge/pkg/logger"
	"github.com/manicar2093/stori-challenge/pkg/validator"

	awsConf "github.com/aws/aws-sdk-go-v2/config"
)

var (
	gookitValidator = validator.NewGooKitValidator()
	dbConn          = connections.GetGormConnection()
	awsConfig       = converters.Must(awsConf.LoadDefaultConfig(context.TODO()))
	service         = txanalizer.NewDefaultService(
		txanalizer.NewTursoRepository(dbConn),
		txanalizer.NewMailgun(txanalizer.DefaultSender, txanalizer.MailgunConfig{
			EmailTo:      config.Instance.EmailTo,
			EmailFrom:    config.Instance.EmailFrom,
			SmtpServer:   config.Instance.SmtpServer,
			SmtpAddr:     config.Instance.SmtpAddr,
			SmtpUser:     config.Instance.SmtpUser,
			SmtpPassword: config.Instance.SmtpPassword,
			SmtpHost:     config.Instance.SmtpHost,
		}),
		filestores.NewAwsBucket(config.Instance.TransactionsFilesBucketName, awsConfig),
		uuid.New,
	)
)

func main() {
	lambda.Start(handler)
}

func handler(
	ctx context.Context,
	event *txanalizer.AnalyzeAccountTransactionsInput,
) error {
	logger.GetLogger().WithField("data", event).Info("Initialize")
	if err := gookitValidator.ValidateStruct(event); err != nil {
		logger.GetLogger().WithField("validations", err).Error("request does not fullfil requirements")
		return err
	}
	event.TransactionsFilePath = config.Instance.TransactionsFilePath
	if err := service.AnalyzeAccountTransactions(*event); err != nil {
		logger.GetLogger().WithField("message", err).Error()
		return err
	}
	logger.GetLogger().Info("Done")
	return nil
}
