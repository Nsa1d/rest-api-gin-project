package models

type Book struct {
	ID     int     `json:"id"`
	Name   *string  `json:"name"`
	Author *string  `json:"author"`
	Year   *int     `json:"year"`
	Genre  *string  `json:"genre"`
	Rating *int `json:"rating"`
}

type BookList struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type Review struct {
	ID         int    `json:"id"`
	BookID     int    `json:"bookID"`
	ReaderName *string `json:"readerName"`
	Rating     *int    `json:"rating"`
	Comment    *string `json:"comment"`
}
