package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/models"
	"gorm.io/gorm"
)

func GetBookByID(context *fiber.Ctx, db *gorm.DB) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "id can't be empty",
			},
		)
	}

	bookModel := &models.Books{}

	err := db.Find(bookModel, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not find book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Success",
			"data":    bookModel,
		},
	)

	return nil
}
