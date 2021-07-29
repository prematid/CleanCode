package middlewares

import (
	"github.com/labstack/echo/middleware"
)

func ConfigLogger() {
	middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	})

}
func ConfigRecover() {
	middleware.Recover()
}
