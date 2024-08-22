package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PingRoutes struct {
	echo *echo.Echo
}

func NewPingRoutes(echo *echo.Echo) *PingRoutes {
	return &PingRoutes{
		echo: echo,
	}
}

func (r *PingRoutes) InitPingRoute() {
	e := r.echo
	r.registerPingRoutes(e)
}

func (r *PingRoutes) registerPingRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, "pong")
	})
	//v1.GET("/db/ping", func(c echo.Context) error {
	//	if err := connections.DbPing(); err != nil {
	//		return c.JSON(http.StatusAccepted, "Error_message:Db connection lost")
	//	}
	//	return c.JSON(http.StatusAccepted, "pong")
	//})
}
