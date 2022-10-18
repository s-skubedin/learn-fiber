package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/models"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupBooksRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBook)
	api.Delete("/books", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		fmt.Println("err", err)
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"},
		)
		return err
	}

	err = r.DB.Create(&book).Error

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

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
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
	err := r.DB.Delete(bookModel, id)

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

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not find books"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"messgae": "Success",
			"data":    bookModels,
		},
	)

	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "id can't be empty",
			},
		)
	}

	bookModel := &models.Books{}

	err := r.DB.Find(bookModel, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not find book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"messgae": "Success",
			"data":    bookModel,
		},
	)

	return nil
}
