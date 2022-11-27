package internal

import (
	"errors"
	"fmt"
)

type Service interface {
	Persist()
	ReadAll() (result []Comment)
}

type ServiceTypeName string
type serviceCreateFunc func(config *Config, comment *Comment) Service

const (
	LocalStoreServiceName ServiceTypeName = "local-store-service"
)

type ServiceFactory struct {
	registerServiceCreateFunc map[ServiceTypeName]serviceCreateFunc
}

var registerServiceCreateFunc = map[ServiceTypeName]serviceCreateFunc{
	LocalStoreServiceName: newLocalStoreService,
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{registerServiceCreateFunc: registerServiceCreateFunc}
}

func (factor *ServiceFactory) CreateService(config *Config, comment *Comment) (Service, error) {
	serviceName, isExisted := comment.Labels["service"]
	if !isExisted {
		return nil, errors.New("no service specified")
	}

	createFunc, isExisted := factor.registerServiceCreateFunc[ServiceTypeName(serviceName)]
	if !isExisted {
		return nil, errors.New(fmt.Sprintf("can't support %s", serviceName))
	}

	return createFunc(config, comment), nil
}
