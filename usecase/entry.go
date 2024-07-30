package usecase

import (
	"context"
	"onion/domain/model"
	"onion/domain/repository"
	"time"

	"github.com/google/uuid"
)

type EntryUsecase interface {
	Search(ctx context.Context, id string) (*model.Entry, error)
	Create(ctx context.Context, user string) (*model.Entry, error)
}

type entryUsecase struct {
	er repository.EntryRepository
}

func NewEntryUsecase(er repository.EntryRepository) EntryUsecase {
	return &entryUsecase{er}
}

func (eu *entryUsecase) Search(ctx context.Context, id string) (*model.Entry, error) {
	res, err := eu.er.Search(id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (eu *entryUsecase) Create(ctx context.Context, user string) (*model.Entry, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// truncate to second
	now := time.Now().Truncate(time.Second)
	entry := model.Entry{
		ID:        id.String(),
		User:      user,
		Status:    "active",
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err := eu.er.Create(entry)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
