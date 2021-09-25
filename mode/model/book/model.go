package book

// 书
type Book struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (b *Book) Builder(id int) *BookBuilder {
	return NewBookBuilder(id)
}