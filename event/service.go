package event

import (
	"api-examples/api"
)

type Service interface {
	GetSomeEventInfo()
}

//func NewService(factorService factor.Service) Service {
//	return &serviceImpl{
//		factorService: factorService,
//	}
//}

func NewService() Service {
	s := &serviceImpl{}
	api.GetSomeEventInfo = s.GetSomeEventInfo
	return s
}

type serviceImpl struct {
	//factorService factor.Service
}

//func (s serviceImpl) GetSomeEventInfo() {
//	// do sth.
//	s.factorService.GetSomeFactorInfo()
//	// do sth.
//}

func (s serviceImpl) GetSomeEventInfo() {
	// do sth.
	api.GetSomeFactorInfo()
	// do sth.
}
