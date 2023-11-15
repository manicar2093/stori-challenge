package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/manicar2093/filestores"
	"github.com/manicar2093/stori-challenge/cmd/api/controllers"
	"github.com/manicar2093/stori-challenge/internal/txanalizer"
	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/connections"
	"github.com/manicar2093/stori-challenge/pkg/logger"
	"github.com/manicar2093/stori-challenge/pkg/validator"

	"github.com/labstack/echo/v4"
)

var (
	port = flag.String("port", "5001", "app port to use")
)

//	@title		stori_challenge
//	@BasePath	/api/v1
//	@version	0.0.0
func main() {
	var (
		echoInstance       = echo.New()
		baseEndpoint       = "/api/v1"
		gookitValidator    = validator.NewGooKitValidator()
		initialController  = controllers.NewInitial()
		dbConn             = connections.GetGormConnection()
		analyzerController = controllers.NewTransactionAnalyzer(
			txanalizer.NewDefaultService(
				txanalizer.NewTursoRepository(dbConn),
				getNotificatorImpl(config.Instance.Environment),
				filestores.NewFileSystem("./files"),
				uuid.New,
			),
		)
		server = NewServer(echoInstance, gookitValidator, baseEndpoint, initialController, analyzerController)
	)

	flag.Parse()

	doMigrations(dbConn)
	fmt.Println("Environment:", config.Instance.Environment)
	logger.GetLogger().Fatal(server.Start(fmt.Sprintf(":%s", *port)))
}

func getNotificatorImpl(env string) txanalizer.Notificator {
	if env == "prod" {
		return txanalizer.NewMailgun(txanalizer.DefaultSender, txanalizer.MailgunConfig{
			EmailTo:      config.Instance.EmailTo,
			EmailFrom:    config.Instance.EmailFrom,
			SmtpServer:   config.Instance.SmtpServer,
			SmtpAddr:     config.Instance.SmtpAddr,
			SmtpUser:     config.Instance.SmtpUser,
			SmtpPassword: config.Instance.SmtpPassword,
			SmtpHost:     config.Instance.SmtpHost,
		})
	}
	return txanalizer.NewEmailMock()
}

func doMigrations(conn *connections.DBWPaginator) {
	if err := conn.AutoMigrate(&txanalizer.Transaction{}); err != nil {
		logger.GetLogger().Panicln(err)
	}
}
