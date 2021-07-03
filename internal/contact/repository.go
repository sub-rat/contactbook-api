package contact

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(offset, limit int, query string) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(c *Contact) error
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
