package api

import (
	"lambda-golang-addons-boilerplate/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
}

func NewHandler(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/item", h.createItem)
	app.Get("/item/:id", h.getItem)
}

func (h *Handler) createItem(c *fiber.Ctx) error {
	item := c.FormValue("item")
	if err := h.service.CreateItem(c.Context(), item); err != nil {
		h.logger.Error("Failed to create item", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create item")
	}
	return c.SendString("Item created")
}

func (h *Handler) getItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := h.service.GetItem(c.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get item", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get item")
	}
	return c.JSON(item)
}
