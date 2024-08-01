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

	// Configura o contêiner com a imagem Docker da versão que queremos utilizar,
	// nome do banco de dados, usuário e senha, e o driver de comunicação.
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

	// Obtém a URI de conexão diretamente do contêiner criado.
	dbURI, err := postgresC.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Cria a conexão utilizando o driver PGX.
	db, err := pgxpool.Connect(ctx, dbURI)
	if err != nil {
		t.Fatal(err)
	}

	// Cria a tabela "books" no banco de dados.
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
		if err := postgresC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}

	return db, teardown
}

func TestHandlers(t *testing.T) {
	db, teardown := setupTestContainer(t)
	defer teardown()

	e := echo.New()
	RegisterRoutes(e, db)

	tests := []struct {
		setupFunc  func(*testing.T, *pgxpool.Pool)
		assertFunc func(*testing.T, *httptest.ResponseRecorder)
		name       string
		method     string
		url        string
		body       string
	}{
		{
			name:   "Create and Get Book",
			method: http.MethodPost,
			url:    "/books",
			body:   `{"title":"Test Book","author":"Author","isbn":"123-4567891234"}`,
			assertFunc: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				var createdBook book.Book
				json.Unmarshal(rec.Body.Bytes(), &createdBook)
				assert.NotEqual(t, 0, createdBook.ID)

				// Get book
				req := httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(createdBook.ID), nil)
				rec = httptest.NewRecorder()
				c := e.NewContext(req, rec)
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
			},
		},
		{
			name:   "Update Book",
			method: http.MethodPut,
			url:    "/books/1",
			body:   `{"title":"Updated Book","author":"Another Author","isbn":"123-4567891235"}`,
			setupFunc: func(t *testing.T, db *pgxpool.Pool) {
				_, err := db.Exec(context.Background(), `INSERT INTO books (title, author, isbn) VALUES ('Another Book', 'Another Author', '123-4567891235')`)
				assert.NoError(t, err)
			},
			assertFunc: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var updatedBook book.Book
				json.Unmarshal(rec.Body.Bytes(), &updatedBook)
				assert.Equal(t, "Updated Book", updatedBook.Title)
			},
		},
		{
			name:   "Delete Book",
			method: http.MethodDelete,
			url:    "/books/1",
			setupFunc: func(t *testing.T, db *pgxpool.Pool) {
				_, err := db.Exec(context.Background(), `INSERT INTO books (title, author, isbn) VALUES ('Book to Delete', 'Author', '123-4567891236')`)
				assert.NoError(t, err)
			},
			assertFunc: func(t *testing.T, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, rec.Code)

				// Try to get deleted book
				req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
				rec = httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues("1")

				if assert.NoError(t, GetBook(c, book.NewService(book.NewRepository(db)))) {
					assert.Equal(t, http.StatusNotFound, rec.Code)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				tt.setupFunc(t, db)
			}

			req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			switch tt.method {
			case http.MethodPost:
				assert.NoError(t, CreateBook(c, book.NewService(book.NewRepository(db))))
			case http.MethodPut:
				c.SetParamNames("id")
				c.SetParamValues("1")
				assert.NoError(t, UpdateBook(c, book.NewService(book.NewRepository(db))))
			case http.MethodDelete:
				c.SetParamNames("id")
				c.SetParamValues("1")
				assert.NoError(t, DeleteBook(c, book.NewService(book.NewRepository(db))))
			}

			tt.assertFunc(t, rec)
		})
	}
}
