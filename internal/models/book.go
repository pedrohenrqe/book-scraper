package models

type Book struct {
	Title       string `json:"title"`
	Price       string `json:"price"`
	Availability string `json:"availability"`
	Rating      string `json:"rating"`
	ProductURL  string `json:"product_url"`
	ImageURL    string `json:"image_url"`
}