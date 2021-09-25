package book

type BookBuilder struct {
	id    int
	name  string
	price float64
}

func NewBookBuilder(id int) *BookBuilder {
	return &BookBuilder{id: id}
}

// Build 构建
func (bb *BookBuilder) Build() *Book {
	book := &Book{ID: bb.id}
	if bb.price > 0 {
		book.Price = bb.price
	}
	return book
}

func (bb *BookBuilder) SetPrice(price float64) *BookBuilder {
	bb.price = price
	return bb
}
