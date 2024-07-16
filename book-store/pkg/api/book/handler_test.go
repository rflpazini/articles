package book

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"book-store/internal/book"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	dbName = "bookstore"
	dbUser = "user"
	dbPass = "S3cret"
)

func setupTestContainer(t *testing.T) (*pgxpool.Pool, func()) {
	ctx := context.Background()

	postgresC, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPass),
		postgres.BasicWaitStrategies(),
		postgres.WithSQLDriver("pgx"),
	)
	if err != nil {
		t.Fatal(err)
	}

	dbURI, err := postgresC.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		t.Fatal(err)
	}

	// Create tables
	_, err = db.Exec(ctx, `
        CREATE TABLE books (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            author VARCHAR(255) NOT NULL,
            isbn VARCHAR(20) NOT NULL
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		db.Close()
		postgresC.Terminate(ctx)
	}

	return db, teardown
}

func TestHandlers(t *testing.T) {
	db, teardown := setupTestContainer(t)
	defer teardown()

	e := echo.New()
	RegisterRoutes(e, db)

	t.Run("Create and Get Book", func(t *testing.T) {
		// Create book
		mockedBook := `{"title":"Test Book","author":"Author","isbn":"123-4567891234"}`
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(mockedBook))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBook(c, book.NewService(book.NewRepository(db)))) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			var createdBook book.Book
			json.Unmarshal(rec.Body.Bytes(), &createdBook)
			assert.NotEqual(t, 0, createdBook.ID)

			// Get book
			req = httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(createdBook.ID), nil)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(createdBook.ID))

			if assert.NoError(t, GetBook(c, book.NewService(book.NewRepository(db)))) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var fetchedBook book.Book
				json.Unmarshal(rec.Body.Bytes(), &fetchedBook)
				assert.Equal(t, createdBook.Title, fetchedBook.Title)
				assert.Equal(t, createdBook.Author, fetchedBook.Author)
				assert.Equal(t, createdBook.ISBN, fetchedBook.ISBN)
			}
		}
	})

	t.Run("Update Book", func(t *testing.T) {
		// Create book
		mockedBook := `{"title":"Another Book","author":"Another Author","isbn":"123-4567891235"}`
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(mockedBook))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBook(c, book.NewService(book.NewRepository(db)))) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			var createdBook book.Book
			json.Unmarshal(rec.Body.Bytes(), &createdBook)

			// Update book
			updatedBook := `{"title":"Updated Book","author":"Another Author","isbn":"123-4567891235"}`
			req = httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(createdBook.ID), strings.NewReader(updatedBook))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(createdBook.ID))

			if assert.NoError(t, UpdateBook(c, book.NewService(book.NewRepository(db)))) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var updatedBookResponse book.Book
				json.Unmarshal(rec.Body.Bytes(), &updatedBookResponse)
				assert.Equal(t, "Updated Book", updatedBookResponse.Title)

				// Get updated book
				req = httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(createdBook.ID), nil)
				rec = httptest.NewRecorder()
				c = e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues(strconv.Itoa(createdBook.ID))

				if assert.NoError(t, GetBook(c, book.NewService(book.NewRepository(db)))) {
					assert.Equal(t, http.StatusOK, rec.Code)
					var fetchedBook book.Book
					json.Unmarshal(rec.Body.Bytes(), &fetchedBook)
					assert.Equal(t, "Updated Book", fetchedBook.Title)
				}
			}
		}
	})

	t.Run("Delete Book", func(t *testing.T) {
		// Create book
		mockedBook := `{"title":"Book to Delete","author":"Author","isbn":"123-4567891236"}`
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(mockedBook))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBook(c, book.NewService(book.NewRepository(db)))) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			var createdBook book.Book
			json.Unmarshal(rec.Body.Bytes(), &createdBook)

			// Delete book
			req = httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(createdBook.ID), nil)
			rec = httptest.NewRecorder()
			c = e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(createdBook.ID))

			if assert.NoError(t, DeleteBook(c, book.NewService(book.NewRepository(db)))) {
				assert.Equal(t, http.StatusOK, rec.Code)

				// Try to get deleted book
				req = httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(createdBook.ID), nil)
				rec = httptest.NewRecorder()
				c = e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues(strconv.Itoa(createdBook.ID))

				if assert.NoError(t, GetBook(c, book.NewService(book.NewRepository(db)))) {
					assert.Equal(t, http.StatusNotFound, rec.Code)
				}
			}
		}
	})
}
