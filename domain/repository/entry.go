package repository

import "onion/domain/model"

type EntryRepository interface {
	Search(id string) (model.Entry, error)
	Create(entry model.Entry) (model.Entry, error)
}
