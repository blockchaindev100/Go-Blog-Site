package handlers

import (
	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Overview of the blog site
// @Schemes http
// @Description Get Overview of the blog site
// @Tags Admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.Overview
// @Router /admin/overview [get]
func (h *Handlers) Overview(c *fiber.Ctx) error {
	var overview models.Overview
	if val, err := h.Repo.TotalPostCount(); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	} else {
		overview.Total_Posts = val
	}
	if val, err := h.Repo.TotalCommands(); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	} else {
		overview.Total_Commands = val
	}
	if val, err := h.Repo.FirstPost(); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	} else {
		overview.First_Blog = val.Created_at
	}
	return c.JSON(&overview)
}
