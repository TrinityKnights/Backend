package config

import (
	"github.com/TrinityKnights/Backend/internal/builder"
	handlerUser "github.com/TrinityKnights/Backend/internal/delivery/http/handler/user"
	handlerEvent "github.com/TrinityKnights/Backend/internal/delivery/http/handler/event"
	"github.com/TrinityKnights/Backend/internal/delivery/http/middleware"
	"github.com/TrinityKnights/Backend/internal/delivery/http/route"
	repositoryUser "github.com/TrinityKnights/Backend/internal/repository/user"
	repositoryEvent "github.com/TrinityKnights/Backend/internal/repository/event"
	serviceUser "github.com/TrinityKnights/Backend/internal/service/user"
	serviceEvent "github.com/TrinityKnights/Backend/internal/service/event"
	"github.com/TrinityKnights/Backend/pkg/cache"
	"github.com/TrinityKnights/Backend/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	Cache    *cache.ImplCache
	App      *echo.Echo
	Log      *logrus.Logger
	Validate *validator.Validate
	JWT      *jwt.JWTConfig
}

func Bootstrap(config *BootstrapConfig) error {
	// Initialize JWT service
	jwtService := jwt.NewJWTService(config.JWT)

	// Initialize repository
	userRepository := repositoryUser.NewUserRepository(config.DB, config.Log)
	eventRepository := repositoryEvent.NewEventRepository(config.DB)

	// Initialize service
	userService := serviceUser.NewUserServiceImpl(config.DB, config.Log, config.Validate, userRepository, jwtService)
	eventService := serviceEvent.NewEventServiceImpl(config.DB, config.Log, config.Validate, eventRepository, jwtService)


	// Initialize handler
	userHandler := handlerUser.NewUserHandler(config.Log, userService)
	eventHandler := handlerEvent.NewEventHandler(eventService)

	// Initialize graphql

	// Initialize middleware
	authMiddleware := middleware.AuthMiddleware(jwtService)

	// Initialize route
	routeConfig := route.Config{
		App:         config.App,
		UserHandler: userHandler,
		EventHandler: eventHandler,  
	}

	// Build routes
	b := builder.Config{
		App:            config.App,
		UserHandler:    userHandler,
		EventHandler:   eventHandler,
		AuthMiddleware: authMiddleware,
		Routes:         &routeConfig,
	}
	b.BuildRoutes()

	config.Log.Infof("Application is ready")
	return nil
}
