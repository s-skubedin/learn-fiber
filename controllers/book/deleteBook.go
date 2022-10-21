package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/models"
	"gorm.io/gorm"
)

func DeleteBook(context *fiber.Ctx, db *gorm.DB) error {
	bookModel := models.Books{}

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id can't be empty",
			},
		)
		return nil
	}
	err := db.Delete(bookModel, id)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not delete book"},
		)
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book has been deleted",
			"data":    map[string](string){"id": id},
		},
	)

	return nil
}
