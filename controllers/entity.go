package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sync"
	"treasury-parser/config"
	"treasury-parser/models"
	"treasury-parser/services/parser"
	"treasury-parser/services/search"
)

type BaseHandler struct {
	db         *gorm.DB
	inProgress bool
	mutex      sync.Mutex
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{db: db}
}

func (h *BaseHandler) IsImportInProgress() bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return h.inProgress
}

func (h *BaseHandler) SetImportInProgress(status bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.inProgress = status
}

func (h *BaseHandler) Update(c *fiber.Ctx) error {
	if h.IsImportInProgress() {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"result": false,
			"info":   "already started",
		})
	}

	h.SetImportInProgress(true)

	go func() {
		data := parser.FetchData(config.Config("FILE_URL"))
		parser.Import(h.db, parser.Parse(data))
		h.SetImportInProgress(false)
	}()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": true,
		"info":   "",
		"code":   "200",
	})
}

func (h *BaseHandler) State(c *fiber.Ctx) error {
	if h.IsImportInProgress() {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"result": false,
			"info":   "updating",
		})

	}

	entities, err := models.FetchAll(h.db)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"result": false,
			"info":   "fetch entities error",
			"code":   "500",
		})
	}

	if len(entities) <= 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": false,
			"info":   "empty",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": false,
		"info":   "ok",
	})
}

func (h *BaseHandler) GetNames(c *fiber.Ctx) error {
	t := c.Query("type")
	searchText := c.Query("name")

	if len(searchText) <= 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"result": false,
			"info":   "",
		})
	}

	return c.JSON(search.DoSearch(h.db, t, searchText))
}
