package runtime

import (
	"github.com/HaziqAliSohail/ApniUniversity/db"
	"github.com/HaziqAliSohail/ApniUniversity/db/mongo"
	"github.com/HaziqAliSohail/ApniUniversity/service"
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
