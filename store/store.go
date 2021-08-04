package store

import (
	"github.com/YogeshTembe/golang_project/logwrapper"
	"gorm.io/gorm"
)

type Store struct {
	Db     *gorm.DB
	Logger *logwrapper.StandardLogger
}

func NewStore(db *gorm.DB, logger *logwrapper.StandardLogger) *Store {
	return &Store{db, logger}
}
