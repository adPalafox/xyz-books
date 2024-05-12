package entity

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Isbn10          *string   `json:"isbn_10"`
	Isbn13          *string   `json:"isbn_13"`
	ListPrice       float32   `json:"list_price"`
	PublicationYear int       `json:"publication_year"`
	PublisherID     int       `json:"publisher_id"`
	Publisher       string    `json:"publisher"`
	ImageUrl        *string   `json:"image_url"`
	Edition         *string   `json:"edition"`
	Authors         *[]Author `json:"authors"`
}

type Author struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}
