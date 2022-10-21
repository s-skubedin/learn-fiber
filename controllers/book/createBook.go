package books

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/types"
	"gorm.io/gorm"
)

func CreateBook(context *fiber.Ctx, DB *gorm.DB) error {
	book := bookTypes.Book{}

	err := context.BodyParser(&book)

	if err != nil {
		fmt.Println("err", err)
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"},
		)
		return err
	}

	err = DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not crate book"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Book has been added",
		},
	)

	return nil
}
