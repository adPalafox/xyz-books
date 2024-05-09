package entity

type Book struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Isbn10          *string   `json:"isbn_10"`
	Isbn13          *string   `json:"isbn_13"`
	ListPrice       int       `json:"list_price"`
	PublicationYear int       `json:"publication_year"`
	Publisher       string    `json:"publisher"`
	ImageUrl        *string   `json:"image_url"`
	Edition         *string   `json:"edition"`
	Authors         *[]Author `json:"authors"`
}

type Author struct {
	Id         int     `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}
