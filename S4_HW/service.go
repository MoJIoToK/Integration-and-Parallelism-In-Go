package main

//type BookStorage interface {
//	List() []*Book
//	Get(string) *Book
//	Update(string, Book) *Book
//	Create(Book)
//	Delete(string) *Book
//}
//
//type BookStore struct{}

//func (b BookStore) List() []*Book {
//	return books
//}
//
//func (b BookStore) Get(id string) *Book {
//	for _, book := range books {
//		if book.ID == id {
//			return book
//		}
//	}
//	return nil
//}
//
//func (b BookStore) Create(book Book) {
//	books = append(books, &book)
//}
//
//func (b BookStore) Delete(id string) *Book {
//	for i, book := range books {
//		if book.ID == id {
//			books = append(books[:i], (books)[i+1:]...)
//			return &Book{}
//		}
//	}
//	return nil
//}
//
//func (b BookStore) Update(id string, bookUpdate Book) *Book {
//	for i, book := range books {
//		if book.ID == id {
//			books[i] = &bookUpdate
//			return book
//		}
//	}
//	return nil
//}
