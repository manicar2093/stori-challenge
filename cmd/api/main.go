package main

import (
	"flag"
	"fmt"

	"github.com/manicar2093/stori-challenge/cmd/api/controllers"
	"github.com/manicar2093/stori-challenge/pkg/config"
	"github.com/manicar2093/stori-challenge/pkg/logger"
	"github.com/manicar2093/stori-challenge/pkg/validator"

	"github.com/labstack/echo/v4"
)

var (
	port = flag.String("port", "5001", "app port to use")
)

// @title		stori_challenge
// @version	0.0.0
func main() {
	var (
		echoInstance      = echo.New()
		baseEndpoint      = "/api/v1"
		gookitValidator   = validator.NewGooKitValidator()
		initialController = controllers.NewInitial()
		server            = NewServer(echoInstance, gookitValidator, baseEndpoint, initialController)
	)

	flag.Parse()

	fmt.Println("Environment:", config.Instance.Environment)
	logger.GetLogger().Fatal(server.Start(fmt.Sprintf(":%s", *port)))
}
