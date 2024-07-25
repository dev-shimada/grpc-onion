package persistence

import (
	"onion/domain/model"
	"onion/domain/repository"

	"gorm.io/gorm"
)

type entryPersistence struct {
	db *gorm.DB
}

func NewEntryPersistence(db *gorm.DB) repository.EntryRepository {
	return &entryPersistence{db}
}

func (ep *entryPersistence) Search(id string) (model.Entry, error) {
	entry := model.Entry{ID: id}
	res := ep.db.First(&entry)
	if res.Error != nil {
		return model.Entry{}, res.Error
	}
	return entry, nil
}

func (ep *entryPersistence) Create(entry model.Entry) (model.Entry, error) {
	e := model.Entry{}
	res := ep.db.Create(&entry)
	if res.Error != nil {
		return e, res.Error
	}
	return entry, nil
}
