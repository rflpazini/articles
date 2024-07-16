package book

import (
	"net/http"
	"strconv"

	"book-store/internal/book"
	"book-store/pkg/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *pgxpool.Pool) {
	repo := book.NewRepository(db)
	service := book.NewService(repo)

	e.GET("/books", func(c echo.Context) error {
		return GetBooks(c, service)
	})
	e.GET("/books/:id", func(c echo.Context) error {
		return GetBook(c, service)
	})
	e.POST("/books", func(c echo.Context) error {
		return CreateBook(c, service)
	})
	e.PUT("/books/:id", func(c echo.Context) error {
		return UpdateBook(c, service)
	})
	e.DELETE("/books/:id", func(c echo.Context) error {
		return DeleteBook(c, service)
	})
}

func GetBooks(c echo.Context, service *book.Service) error {
	books, err := service.GetAllBooks()
	if err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.RespondWithJSON(c, http.StatusOK, books)
}

func GetBook(c echo.Context, service *book.Service) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := service.GetBookByID(id)
	if err != nil {
		return utils.RespondWithError(c, http.StatusNotFound, "Book not found")
	}
	return utils.RespondWithJSON(c, http.StatusOK, book)
}

func CreateBook(c echo.Context, service *book.Service) error {
	var b book.Book
	if err := c.Bind(&b); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
	}
	if err := service.CreateBook(&b); err != nil {
		return utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.RespondWithJSON(c, http.StatusCreated, b)
}

func UpdateBook(c echo.Context, service *book.Service) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var b book.Book
	if err := c.Bind(&b); err != nil {
		return utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
	}
	if err := service.UpdateBook(id, &b); err != nil {
		return utils.RespondWithError(c, http.StatusNotFound, "Book not found")
	}
	return utils.RespondWithJSON(c, http.StatusOK, b)
}

func DeleteBook(c echo.Context, service *book.Service) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteBook(id); err != nil {
		return utils.RespondWithError(c, http.StatusNotFound, "Book not found")
	}
	return utils.RespondWithJSON(c, http.StatusOK, map[string]string{"result": "success"})
}
