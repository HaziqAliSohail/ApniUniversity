package runtime

import (
	"ApniUniversity/db"
	"ApniUniversity/db/mongo"
	"ApniUniversity/service"
)

type Runtime struct {
	service *service.Service
}

func NewRuntime() (*Runtime, error) {
	dataStore, err := mongo.NewClient(db.Options{})
	if err != nil {

		return nil, err
	}

	return &Runtime{service: service.NewService(dataStore)}, err
}

func (r Runtime) Service() *service.Service {
	return r.service
}
