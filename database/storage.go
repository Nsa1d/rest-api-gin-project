package database

import (
	"errors"
	"rest-api-gin-project/models"
)

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

var Books = []models.Book{
	{
		ID:     1,
		Name:   strPtr("Ведьмак"),
		Author: strPtr("Анджей Сапковский"),
		Year:   intPtr(1986),
		Genre:  strPtr("Фэнтези"),
		Rating: intPtr(4),
	},
	{
		ID:     2,
		Name:   strPtr("Властелин колец"),
		Author: strPtr("Джон Рональд Руэл Толкин"),
		Year:   intPtr(1954),
		Genre:  strPtr("Фэнтези"),
		Rating: intPtr(5),
	},
}

var BookList = []models.BookList{
	{
		ID:     1,
		Name:   "Ведьмак",
		Author: "Анджей Сапковский",
	},
}

var Reviews = []models.Review{
	{
		ID:         1,
		BookID:     1,
		ReaderName: strPtr("Рахьман"),
		Rating:     intPtr(4),
		Comment:    strPtr("Топ"),
	},
}

func GetAllBooks() ([]models.BookList, error) {
	var result []models.BookList
	
	for _, b := range Books {
		var name string
		var author string
		
		if b.Name != nil {
			name = *b.Name
		}
		if b.Author != nil {
			author = *b.Author
		}
		
		result = append(result, models.BookList{
			ID:     b.ID,
			Name:   name,
			Author: author,
		})
	}
	
	return result, nil
}

func GetBookByID(id int) (models.Book, error) {
	for _, b := range Books {
		if b.ID == id {
			return b, nil
		}
	}
	return models.Book{}, errors.New("Книга не найдена")
}

func AddBook(book models.Book) (models.Book, error) {
	maxID := 0
	for _, b := range Books {
		if b.ID > maxID {
			maxID = b.ID
		}
	}
	book.ID = maxID + 1
	Books = append(Books, book)
	return book, nil
}

func UpdateBook(id int, input models.Book) (models.Book, error) {
	for i := range Books {
		if Books[i].ID == id {
			if input.Name != nil {
				Books[i].Name = input.Name
			}
			if input.Author != nil {
				Books[i].Author = input.Author
			}
			if input.Year != nil {
				Books[i].Year = input.Year
			}
			if input.Genre != nil {
				Books[i].Genre = input.Genre
			}
			if input.Rating != nil {
				Books[i].Rating = input.Rating
			}
			return Books[i], nil
		}
	}
	return models.Book{}, errors.New("книга не найдена")
}

func DeleteBook(id int) error {
	for i, b := range Books {
		if b.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return nil
		}
	}
	return errors.New("книга не найдена")
}

func GetAllReviews() ([]models.Review, error) {
	return Reviews, nil
}

func GetReviewsByBookID(bookID int) ([]models.Review, error) {
	var result []models.Review
	for _, r := range Reviews {
		if r.BookID == bookID {
			result = append(result, r)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("отзывы не найдены")
	}
	return result, nil
}