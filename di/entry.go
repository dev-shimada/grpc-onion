package di

import (
	"onion/infrastructure/persistence"
	"onion/interface/handler"
	"onion/usecase"

	"gorm.io/gorm"
)

func Entry(db *gorm.DB) handler.EntryHandler {
	ep := persistence.NewEntryPersistence(db)
	eu := usecase.NewEntryUsecase(ep)
	eh := handler.NewEntryHandler(eu)
	return eh
}
