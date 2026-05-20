package exporter

import (
	"encoding/json"
	"os"

	"github.com/pedrohenrqe/book-scraper/internal/models"
)

func ExportJSON(books []models.Book, filename string) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}