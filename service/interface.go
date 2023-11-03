package service

import "ApniUniversity/db"

type Service struct {
	db db.DataStore
}

func NewService(dataStore db.DataStore) *Service {
	return &Service{db: dataStore}
}
