package database

import (
	"farisjr/project/config"
	"farisjr/project/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(id int) (models.Book, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(create_book models.Book) (interface{}, error) {
	if err := config.DB.Save(&create_book).Error; err != nil {
		return nil, err
	}
	return create_book, nil
}

func DeleteBookById(delete_book models.Book) (interface{}, error) {
	if err := config.DB.Delete(&delete_book).Error; err != nil {
		return nil, err
	}
	return delete_book, nil
}

func UpdateBookById(update_book models.Book) (interface{}, error) {

	if err := config.DB.Save(&update_book).Error; err != nil {
		return nil, err
	}
	return update_book, nil
}

func GetDetailBooks(bookId int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, bookId).Error; err != nil {
		return nil, err
	}

	return book, nil
}
