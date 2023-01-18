package user

import "api-examples/event"

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
	GetSomeUserInfo = s.GetSomeUserInfo
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
	event.GetSomeEventInfo()
	// do sth.
}
