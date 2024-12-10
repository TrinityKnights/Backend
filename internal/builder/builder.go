package builder

import (
	"github.com/TrinityKnights/Backend/internal/delivery/http/handler/event"
	"github.com/TrinityKnights/Backend/internal/delivery/http/handler/user"
	"github.com/TrinityKnights/Backend/internal/delivery/http/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	App            *echo.Echo
	UserHandler    *user.UserHandlerImpl
	EventHandler   *event.EventHandlerImpl
	AuthMiddleware echo.MiddlewareFunc
	Routes         *route.Config
}

func (c *Config) BuildRoutes() {
	// Global middleware
	c.App.Use(middleware.Recover())

	// API group
	g := c.App.Group("/api/v1")
	g.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(60)))

	// Public routes
	for _, r := range c.Routes.PublicRoute() {
		g.Add(r.Method, r.Path, r.Handler)
	}

	// Private routes with auth middleware
	privateGroup := g.Group("/api/v1", c.AuthMiddleware)
	for _, r := range c.Routes.PrivateRoute() {
		privateGroup.Add(r.Method, r.Path, r.Handler)
	}

	// Swagger routes
	c.Routes.SwaggerRoutes()
}
