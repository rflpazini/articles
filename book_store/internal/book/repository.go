package book

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllBooks() ([]Book, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, title, author, isbn FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) GetBookByID(id int) (*Book, error) {
	var book Book
	err := r.db.QueryRow(context.Background(), "SELECT id, title, author, isbn FROM books WHERE id=$1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

func (r *Repository) CreateBook(book *Book) error {
	err := r.db.QueryRow(context.Background(), "INSERT INTO books (title, author, isbn) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.ISBN).Scan(&book.ID)
	return err
}

func (r *Repository) UpdateBook(id int, book *Book) error {
	commandTag, err := r.db.Exec(context.Background(), "UPDATE books SET title=$1, author=$2, isbn=$3 WHERE id=$4",
		book.Title, book.Author, book.ISBN, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("book not found")
	}
	book.ID = id
	return nil
}

func (r *Repository) DeleteBook(id int) error {
	commandTag, err := r.db.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("book not found")
	}
	return nil
}
