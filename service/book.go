package service

import "challange-8/model"

// usecase

type BooksService interface {
	Addbook(in model.Books) (res model.Books, err error)
	GetBookByID(id int64) (res model.Books, err error)
	GetBooks(data []model.Books) (res []model.Books, err error)
	UpdateBook(in model.Books) (res model.Books, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) Addbook(in model.Books) (res model.Books, err error) {
	// call repo
	res, err = s.repo.Addbook(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBookByID(id int64) (res model.Books, err error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) GetBooks(data []model.Books) (res []model.Books, err error) {
	return s.repo.GetBooks(data)
}

func (s *Service) UpdateBook(in model.Books) (res model.Books, err error) {
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
