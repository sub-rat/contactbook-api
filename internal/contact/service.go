package contact

import "github.com/sub-rat/contactbook-api/internal/model"

// offset, limit => pagination
type Service interface {
	Query(offset, limit int, query string) ([]Contact, error)
}

type service struct {
	repo Repository
}

type Contact struct {
	model.Contact
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Query(offset, limit int, query string) ([]Contact, error) {
	items, err := s.repo.Query(offset, limit, query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
