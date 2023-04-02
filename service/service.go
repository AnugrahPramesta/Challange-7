package service

import "challange-8/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	BooksService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
