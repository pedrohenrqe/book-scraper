package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/pedrohenrqe/book-scraper/internal/logger"
	"github.com/pedrohenrqe/book-scraper/internal/models"
)

const baseURL = "https://books.toscrape.com/catalogue/page-%d.html"

func ScrapeBooks() ([]models.Book, error) {
	var books []models.Book

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	for page := 1; page <= 50; page++ {
		url := fmt.Sprintf(baseURL, page)

		logger.Logger.Printf("Scraping page %d\n", page)

		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			if err := resp.Body.Close(); err != nil {
				return nil, err
			}
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if errClose := resp.Body.Close(); errClose != nil {
			return nil, errClose
		}

		if err != nil {
			return nil, err
		}

		doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
			title, _ := s.Find("h3 a").Attr("title")

			price := strings.TrimSpace(
				s.Find(".price_color").Text(),
			)

			availability := strings.TrimSpace(
				s.Find(".availability").Text(),
			)

			ratingClass, _ := s.Find(".star-rating").Attr("class")
			rating := strings.Replace(ratingClass, "star-rating ", "", 1)

			productURL, _ := s.Find("h3 a").Attr("href")
			imageURL, _ := s.Find("img").Attr("src")

			book := models.Book{
				Title:        title,
				Price:        price,
				Availability: availability,
				Rating:       rating,
				ProductURL:   "https://books.toscrape.com/catalogue/" + productURL,
				ImageURL:     "https://books.toscrape.com/" + imageURL,
			}

			books = append(books, book)
		})
	}

	return books, nil
}
