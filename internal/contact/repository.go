package contact

import "gorm.io/gorm"

type Repository interface {
	Query(offset, limit int, query string) ([]Contact, error)
}

type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) Repository {
	return repository{db}
}

// select * from contacts where name like %gol%
func (repository repository) Query(offset, limit int, query string) ([]Contact, error) {
	var contact []Contact
	err := repository.db.
		Where("lower(name) like lower(?)", "%"+query+"%").
		Order("id").
		Offset(offset).
		Limit(limit).
		Find(&contact).Error
	return contact, err
}
