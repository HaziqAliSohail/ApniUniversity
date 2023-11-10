package service

import "github.com/HaziqAliSohail/ApniUniversity/db"

type Service struct {
	db db.DataStore
}

func NewService(dataStore db.DataStore) *Service {
	return &Service{db: dataStore}
}
