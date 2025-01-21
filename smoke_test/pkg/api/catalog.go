package api

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type CatalogItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

func GetCatalog(c echo.Context) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "falha ao conectar ao banco de dados"})
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, type FROM catalog")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "falha ao consultar o banco de dados"})
	}
	defer rows.Close()

	var catalog []CatalogItem
	for rows.Next() {
		var item CatalogItem
		if err := rows.Scan(&item.ID, &item.Title, &item.Type); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "falha ao ler os dados do banco"})
		}
		catalog = append(catalog, item)
	}

	return c.JSON(http.StatusOK, catalog)
}
