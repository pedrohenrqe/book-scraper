package main

import (
	"os"

	"github.com/pedrohenrqe/book-scraper/internal/exporter"
	"github.com/pedrohenrqe/book-scraper/internal/logger"
	"github.com/pedrohenrqe/book-scraper/internal/scraper"
)

func main() {
	logger.Logger.Println("Starting scraper")

	books, err := scraper.ScrapeBooks()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	err = os.MkdirAll("output", os.ModePerm)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	if err := exporter.ExportJSON(books, "output/books.json"); err != nil {
		logger.Logger.Fatal(err)
	}

	if err := exporter.ExportCSV(books, "output/books.csv"); err != nil {
		logger.Logger.Fatal(err)
	}

	logger.Logger.Printf("Successfully scraped %d books\n", len(books))
}