package model

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float32
}

type BookReq struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func (b *BookReq) MapperToBook() *Book {
	return &Book{
		Title:  b.Title,
		Author: b.Author,
		Price:  b.Price,
	}
}
