package contact

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sub-rat/contactbook-api/internal/model"
)

// offset, limit => pagination
type ServiceInterface interface {
	Query(offset, limit int, query string) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(req Contact) (Contact, error)
	Update(id uint, input Contact) (Contact, error)
	Delete(id uint) (Contact, error)
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

func (c Contact) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&c.PhoneNo, validation.Required, validation.Length(10, 10)),
	)
}

func (s service) Get(id uint) (Contact, error) {
	contact, err := s.repo.Get(id)
	if err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func (s service) Create(req Contact) (Contact, error) {
	if err := req.Validate(); err != nil {
		return Contact{}, err
	}
	err := s.repo.Create(&req)
	if err != nil {
		return Contact{}, err
	}
	return s.Get(req.ID)
}

// Update updates the Contact with the specified ID.
func (s service) Update(id uint, req Contact) (Contact, error) {

	contact, err := s.Get(id)
	if err != nil {
		return Contact{}, err
	}
	if req.Name != contact.Name && req.Name != "" {
		contact.Name = req.Name
	}

	if req.Email != contact.Email && req.Email != "" {
		contact.Email = req.Email
	}

	if req.PhoneNo != contact.PhoneNo && req.PhoneNo != "" {
		contact.PhoneNo = req.PhoneNo
	}

	contact.UpdatedAt = time.Now()

	if err := s.repo.Update(&contact); err != nil {
		return Contact{}, err
	}
	return contact, nil
}

// Delete deletes the Contact with the specified ID.
func (s service) Delete(id uint) (Contact, error) {
	contact, err := s.Get(id)
	if err != nil {
		return Contact{}, err
	}
	if err = s.repo.Delete(id); err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func (s service) Query(offset, limit int, query string) ([]Contact, error) {
	items, err := s.repo.Query(offset, limit, query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
