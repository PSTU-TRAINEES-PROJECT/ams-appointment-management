package cmd

import (
	"ams-appointment-management/app/common/logger"
	"ams-appointment-management/app/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	//connections "test/app/conn/db"
	"ams-appointment-management/app/routes"
	"ams-appointment-management/app/server"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// init logger
	logger.NewLogger()

	// Config
	config.InitConfig()
	var cfg = config.GetConfig()

	// Connections
	//db := connections.NewGormDb(cfg)
	//connections.Migration(db)

	var framework = echo.New()
	var pingRoutes = routes.NewPingRoutes(framework)
	pingRoutes.InitPingRoute()

	var Server = server.New(cfg, framework)
	Server.Start()
}
