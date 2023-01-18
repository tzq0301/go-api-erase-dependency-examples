package event

import "api-examples/factor"

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
	GetSomeEventInfo = s.GetSomeEventInfo
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
	factor.GetSomeFactorInfo()
	// do sth.
}
