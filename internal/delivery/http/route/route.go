package route

import (
	"github.com/TrinityKnights/Backend/internal/delivery/http/handler/event"
	"github.com/TrinityKnights/Backend/internal/delivery/http/handler/user"
	"github.com/TrinityKnights/Backend/pkg/route"
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
)

type RouteConfig interface {
	PublicRoute() []route.Route
	PrivateRoute() []route.Route
	SwaggerRoutes()
}

type Config struct {
	App         *echo.Echo
	UserHandler *user.UserHandlerImpl
	EventHandler *event.EventHandlerImpl
}

func (c *Config) PublicRoute() []route.Route {
	return []route.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: c.UserHandler.Register,
		},
		{
			Method:  echo.POST,
			Path:    "/users/login",
			Handler: c.UserHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/users/refresh",
			Handler: c.UserHandler.RefreshToken,
		},
		{
			Method:  echo.GET,
			Path:    "/event/:id",
			Handler: c.EventHandler.GetEventWithDetails,
		},
	}
}

func (c *Config) PrivateRoute() []route.Route {
	return []route.Route{
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: c.UserHandler.Profile,
		},
		{
			Method:  echo.PUT,
			Path:    "/users",
			Handler: c.UserHandler.Update,
		},
	}
}

func (c *Config) SwaggerRoutes() {
	c.App.GET("/swagger/*", swagger.WrapHandler)

	c.App.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(301, "/swagger/index.html")
	})
}
