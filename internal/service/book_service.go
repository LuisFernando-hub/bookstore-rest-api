package service

import (
	"bookstore-rest-api/internal/model"
	"bookstore-rest-api/internal/store"
	"errors"
)

type Logger interface {
	Log(msg, error string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store:  s,
		logger: Logger(nil),
	}
}

func (s *Service) GetAll() ([]*model.Book, error) {
	books, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) GetByID(id int) (*model.Book, error) {
	book, err := s.store.GetByID(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) Create(book *model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("Title is required")
	}

	if book.Author == "" {
		return nil, errors.New("Author is required")
	}

	return s.store.Create(book)
}

func (s *Service) Update(id int, book *model.Book) (*model.Book, error) {
	s.logger.Log("Update called in Service layer - updating book", "")
	if book.Title == "" {
		return nil, errors.New("Title is required")
	}

	if book.Author == "" {
		return nil, errors.New("Author is required")
	}

	return s.store.Update(id, book)
}

func (s *Service) Delete(id int) error {
	return s.store.Delete(id)
}
