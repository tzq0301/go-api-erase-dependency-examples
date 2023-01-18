package user

import (
	"api-examples/api"
)

type Service interface {
	GetSomeUserInfo()
}

//func NewService(eventService event.Service) Service {
//	return &serviceImpl{
//		eventService: eventService,
//	}
//}

func NewService() Service {
	s := &serviceImpl{}
	api.GetSomeUserInfo = s.GetSomeUserInfo
	return s
}

type serviceImpl struct {
	//eventService event.Service
}

//func (s serviceImpl) GetSomeUserInfo() {
//	// do sth.
//	s.eventService.GetSomeEventInfo()
//	// do sth.
//}

func (s serviceImpl) GetSomeUserInfo() {
	// do sth.
	api.GetSomeEventInfo()
	// do sth.
}
