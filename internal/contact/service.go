package contact

import "github.com/sub-rat/contactbook-api/internal/model"

// offset, limit => pagination
type ServiceInterface interface {
	Query(offset, limit int, query string) ([]Contact, error)
}

type service struct {
	repo RepositoryInterface
}

type Contact struct {
	model.Contact
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return service{repo}
}

func (s service) Query(offset, limit int, query string) ([]Contact, error) {
	items, err := s.repo.Query(offset, limit, query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
