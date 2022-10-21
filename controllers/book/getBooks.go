package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/models"
	"gorm.io/gorm"
)

func GetBooks(context *fiber.Ctx, db *gorm.DB) error {
	bookModels := &[]models.Books{}

	err := db.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not find books"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Success",
			"data":    bookModels,
		},
	)

	return nil
}
