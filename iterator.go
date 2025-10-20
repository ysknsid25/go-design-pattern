package main

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Book struct {
	name string
}

type BookShelf struct {
	books []Book
	last  int
}

func NewBookShelf(size int) *BookShelf {
	return &BookShelf{
		books: make([]Book, size),
		last:  0,
	}
}

func (bs *BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book Book) {
	bs.books[bs.last] = book
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
		index:     0,
	}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (it *BookShelfIterator) HasNext() bool {
	return it.index < it.bookShelf.GetLength()
}

func (it *BookShelfIterator) Next() interface{} {
	book := it.bookShelf.GetBookAt(it.index)
	it.index++
	return book
}

func execIterator() {
	bookShelf := NewBookShelf(100)
	bookShelf.AppendBook(Book{name: "Book 1"})
	bookShelf.AppendBook(Book{name: "Book 2"})
	bookShelf.AppendBook(Book{name: "Book 3"})

	iterator := bookShelf.Iterator()
	for iterator.HasNext() {
		book := iterator.Next().(Book)
		println(book.name)
	}
}
