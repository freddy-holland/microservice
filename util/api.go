package util

import (
	"net/http"

	"fholl.net/microservice-base/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GetByIDConfig struct {
	Model    interface{}
	Preloads []string
}

func GetModelByID(c echo.Context, cfg GetByIDConfig) error {
	id := c.Param("id")

	var q *gorm.DB = database.DB
	for _, pl := range cfg.Preloads {
		q = q.Preload(pl)
	}

	if err := q.Where("id = ?", id).First(cfg.Model).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not download model",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, cfg.Model)
}

func CreateModel(c echo.Context, model interface{}) error {

	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "could not create new model",
			"error":   err,
		})
	}

	if err := database.DB.Create(model).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not upload model to database",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, model)
}

func DeleteModel(c echo.Context, model interface{}) error {
	id := c.Param("id")

	if err := database.DB.Delete(model, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not delete model from database",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "model deleted successfully",
		"id":      id,
	})
}
