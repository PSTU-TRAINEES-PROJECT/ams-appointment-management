package routes

import (
	"ams-appointment-management/app/controller"

	"github.com/labstack/echo/v4"
)

type AppointmentRoutes struct {
	echo                  *echo.Echo
	appointmentController *controller.AppointmentController
}

func NewAppointmentRoutes(echo *echo.Echo, appointmentController *controller.AppointmentController) *AppointmentRoutes {
	return &AppointmentRoutes{
		echo:                  echo,
		appointmentController: appointmentController,
	}
}

func (r *AppointmentRoutes) InitUserRoute() {
	e := r.echo
	r.registerUserRoutes(e)
}

func (r *AppointmentRoutes) registerUserRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/appointment/create", func(c echo.Context) error {
		return r.appointmentController.Create(c)
	})
	v1.GET("/appointment/list", func(c echo.Context) error {
		return r.appointmentController.FindAll(c)
	})
	v1.GET("/appointment/view/:id", func(c echo.Context) error {
		return r.appointmentController.FindByID(c)
	})
	v1.PATCH("/appointment/update/:id", func(c echo.Context) error {
		return r.appointmentController.Update(c)
	})
	v1.DELETE("/appointment/delete/:id", func(c echo.Context) error {
		return r.appointmentController.Delete(c)
	})

}
