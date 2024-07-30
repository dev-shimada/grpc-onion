package usecase_test

import (
	"context"
	"onion/domain/model"
	"onion/domain/repository"
	"onion/usecase"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

type MockEntryRepository struct {
	repository.EntryRepository
	FakeSearch func(id string) (model.Entry, error)
	FakeCreate func(entry model.Entry) (model.Entry, error)
}

func (m MockEntryRepository) Search(id string) (model.Entry, error) {
	return m.FakeSearch(id)
}
func (m MockEntryRepository) Create(entry model.Entry) (model.Entry, error) {
	return m.FakeCreate(entry)
}

func TestSearch(t *testing.T) {
	entry := model.Entry{
		ID:        "1",
		User:      "user",
		Status:    "active",
		CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
		//zero value
		DeletedAt: time.Date(01, 01, 01, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2021, 1, 1, 0, 0, 1, 0, time.Local),
	}
	er := MockEntryRepository{
		FakeSearch: func(id string) (model.Entry, error) {
			return entry, nil
		},
	}

	eu := usecase.NewEntryUsecase(er)
	got, err := eu.Search(context.TODO(), "1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &entry
	if cmp.Equal(got, want) {
		t.Logf("ok")
	} else {
		t.Errorf("diff: %v", cmp.Diff(got, want))
	}
}

func TestCreate(t *testing.T) {
	er := MockEntryRepository{
		FakeCreate: func(entry model.Entry) (model.Entry, error) {
			return entry, nil
		},
	}
	eu := usecase.NewEntryUsecase(er)
	got, err := eu.Create(context.TODO(), "user")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check ID is UUIDv7
	uuid, err := uuid.Parse(got.ID)
	if err != nil && uuid.Version() != 7 {
		t.Errorf("unexpected error: %v", err)
	}

	// check User is "user"
	if got.User != "user" {
		t.Errorf("User is not 'user'")
	}

	// check Status is "active"
	if got.Status != "active" {
		t.Errorf("Status is not 'active'")
	}

	// check if DeletedAt is zero value
	if !got.DeletedAt.IsZero() {
		t.Errorf("DeletedAt is not zero value")
	}

	// check if CreatedAt is in local timezone
	if got.CreatedAt.Location() != time.Local {
		t.Errorf("CreatedAt is not in local timezone")
	}

	// check if the CreatedAt and UpdatedAt are the same
	if got.CreatedAt != got.UpdatedAt {
		t.Errorf("CreatedAt and UpdatedAt are different")
	}
}
