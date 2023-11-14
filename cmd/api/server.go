package main

import (
	"github.com/manicar2093/stori-challenge/pkg/apperrors"
	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/logger"
	"github.com/manicar2093/stori-challenge/pkg/validator"

	_ "github.com/manicar2093/stori-challenge/cmd/api/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manicar2093/echoroutesview"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	Server struct {
		echoInstance    *echo.Echo
		gookitValidator *validator.GooKitValidator
		baseEndpoint    *echo.Group
		controllers     []Controller
	}

	Controller interface {
		SetUpRoutes(*echo.Group)
	}
)

func NewServer(
	echoInstance *echo.Echo,
	gookitValidator *validator.GooKitValidator,
	baseEndpoint string,
	controllers ...Controller,
) *Server {
	server := &Server{
		echoInstance:    echoInstance,
		gookitValidator: gookitValidator,
		baseEndpoint:    echoInstance.Group(baseEndpoint),
		controllers:     controllers,
	}

	server.configEcho()
	server.configControllers()

	logger.GetLogger().Info(server.echoInstance.Routes())

	return server
}

func (c *Server) Start(address string) error {
	return c.echoInstance.Start(address)
}

func (c *Server) configEcho() {
	c.echoInstance.HideBanner = true
	c.echoInstance.Use(middleware.Recover())
	c.echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	c.echoInstance.Validator = c.gookitValidator
	c.echoInstance.HTTPErrorHandler = apperrors.HandlerWEcho

	c.echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)
	if config.Instance.Environment == "dev" {
		logger.GetLogger().Info("/echo endpoint registered")
		c.echoInstance.GET("/echo", func(ctx echo.Context) error {
			var body = make(map[string]interface{})
			if err := ctx.Bind(&body); err != nil {
				logger.GetLogger().WithField("message", "error on /echo handler").Error(err)
				return err
			}
			logger.GetLogger().WithFields(logrus.Fields{
				"data_received": map[string]interface{}{
					"body":         body,
					"query_params": ctx.QueryParams(),
				},
			}).Info("echoed data received")
			return nil
		})
	}
}

func (c *Server) configControllers() {
	for _, controller := range c.controllers {
		controller.SetUpRoutes(c.baseEndpoint)
	}
	echoroutesview.RegisterRoutesViewer(c.echoInstance) //nolint: errcheck
}
