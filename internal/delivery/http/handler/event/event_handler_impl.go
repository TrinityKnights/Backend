package event

import (
	"net/http"
	"strconv"

	"github.com/TrinityKnights/Backend/internal/delivery/http/handler"
	"github.com/TrinityKnights/Backend/internal/domain/model"
	"github.com/TrinityKnights/Backend/internal/service/event"
	"github.com/labstack/echo/v4"
)

type EventHandlerImpl struct {
	Event event.EventService
}

func NewEventHandler(eventService event.EventService) *EventHandlerImpl {
	return &EventHandlerImpl{Event: eventService}
}

func (h *EventHandlerImpl) GetEventWithDetails(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return handler.HandleError(ctx, 400, err)
	}
	response, err := h.Event.GetEventWithDetails(ctx.Request().Context(), uint(id))
	if err != nil {
		return handler.HandleError(ctx, 500, err)
	}
	return ctx.JSON(http.StatusOK, model.NewResponse(response, nil))
}
