package exporter

import (
	"encoding/csv"
	"os"

	"github.com/pedrohenrqe/book-scraper/internal/models"
)

func ExportCSV(books []models.Book, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"title",
		"price",
		"availability",
		"rating",
		"product_url",
		"image_url",
	}

	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, book := range books {
		record := []string{
			book.Title,
			book.Price,
			book.Availability,
			book.Rating,
			book.ProductURL,
			book.ImageURL,
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}