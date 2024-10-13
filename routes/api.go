package routes

import (
	"amartha-billing-app/common"
	billing_handler "amartha-billing-app/handlers/billing"

	"github.com/labstack/echo/v4"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []common.BaseHandler{
		billing_handler.New(),
	}
	var routes []common.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Routes()...)
	}
	api := e.Group("/api/")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middlewares...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middlewares...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middlewares...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middlewares...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middlewares...)
				break
			}
		}
	}
}
