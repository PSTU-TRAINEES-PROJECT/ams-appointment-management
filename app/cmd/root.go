package cmd

import (
	"ams-appointment-management/app/common/logger"
	"ams-appointment-management/app/config"
	"ams-appointment-management/app/controller"
	"ams-appointment-management/app/service"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	//connections "test/app/conn/db"
	connections "ams-appointment-management/app/conn/db"
	"ams-appointment-management/app/domain/repository"
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
	db := connections.NewGormDb(cfg)
	connections.Migration(db)

	var framework = echo.New()
	var pingRoutes = routes.NewPingRoutes(framework)
	pingRoutes.InitPingRoute()

	//appointment segment
	var appointmentRepository = repository.NewAppointmentRepository(db)
	var appointmentService = service.NewAppointmentService(appointmentRepository)
	var appointmentController = controller.NewAppointmentController(appointmentService)

	var appointmentRoutes = routes.NewAppointmentRoutes(framework, &appointmentController)
	appointmentRoutes.InitAppointmentRoute()

	var Server = server.New(cfg, framework)
	Server.Start()
}
