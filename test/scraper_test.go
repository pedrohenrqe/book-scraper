package test

import (
	"testing"

	"github.com/pedrohenrqe/book-scraper/internal/scraper"
)

func TestScrapeBooks(t *testing.T) {
	books, err := scraper.ScrapeBooks()

	if err != nil {
		t.Fatalf("error scraping books: %v", err)
	}

	if len(books) == 0 {
		t.Fatal("expected books but got none")
	}
}