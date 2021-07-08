package contact

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(offset, limit int, query string) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(c *Contact) error
	// Update updates the contact with given ID in the storage.
	Update(contact *Contact) error
	// Delete removes the contact with given ID from the storage.
	Delete(id uint) error
}

type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) RepositoryInterface {
	return repository{db}
}

// SELECT * FROM contacts where id = 1 limit 1;
func (repository repository) Get(id uint) (Contact, error) {
	var c Contact
	err := repository.db.First(&c, "id = ?", id).Error
	return c, err
}

// Create
func (repository repository) Create(c *Contact) error {
	return repository.db.Create(&c).Error
}

// Update saves the changes to an contact in the database.
func (r repository) Update(contact *Contact) error {
	return r.db.Where("id = ?", contact.ID).UpdateColumns(&contact).Error
}

// Delete deletes an contact with the specified ID from the database.
func (r repository) Delete(id uint) error {
	contact, err := r.Get(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&contact).Error
}

// select * from contacts where name like %gol%
func (repository repository) Query(offset, limit int, query string) ([]Contact, error) {
	var contactList []Contact
	err := repository.db.
		Where("lower(name) like lower(?)", "%"+query+"%").
		Order("id").
		Offset(offset).
		Limit(limit).
		Find(&contactList).Error
	return contactList, err
}
