package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/s-skubedin/todo-fiber/controllers/book"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	return books.CreateBook(context, r.DB)
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	return books.DeleteBook(context, r.DB)
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	return books.GetBooks(context, r.DB)
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	return books.GetBookByID(context, r.DB)
}

func (r *Repository) SetupBooksRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/books", r.CreateBook)
	api.Get("/books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
	api.Delete("/books", r.DeleteBook)
}
