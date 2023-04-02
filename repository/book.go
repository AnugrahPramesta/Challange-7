package repository

import (
	"challange-8/model"
)

// clean architectures -> handler->service->repo

// interface employee
type BookRepo interface {
	Addbook(in model.Books) (res model.Books, err error)
	GetBookByID(id int64) (res model.Books, err error)
	GetBooks(data []model.Books) (res []model.Books, err error)
	UpdateBook(in model.Books) (res model.Books, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) Addbook(in model.Books) (res model.Books, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookByID(id int64) (res model.Books, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBooks(data []model.Books) (res []model.Books, err error) {
	err = r.gorm.Find(&data).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r Repo) UpdateBook(in model.Books) (res model.Books, err error) {
	err = r.gorm.Model(&res).Where("id = ?", in.ID).First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Model(&res).Updates(model.Books{
		NameBook: in.NameBook,
		Author:   in.Author,
	}).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	var book model.Books
	err = r.gorm.Model(&book).Where("id = ?", id).First(&book).Error
	if err != nil {
		return err
	}
	r.gorm.Delete(&book)
	return nil
}
