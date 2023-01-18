package factor

import (
	"api-examples/api"
)

type Service interface {
	GetSomeFactorInfo()
}

//func NewService(userService user.Service) Service {
//	return &serviceImpl{
//		userService: userService,
//	}
//}

func NewService() Service {
	s := &serviceImpl{}
	api.GetSomeFactorInfo = s.GetSomeFactorInfo
	return s
}

type serviceImpl struct {
	//userService user.Service
}

//func (s serviceImpl) GetSomeFactorInfo() {
//	// do sth.
//	s.userService.GetSomeUserInfo()
//	// do sth.
//}

func (s serviceImpl) GetSomeFactorInfo() {
	// do sth.
	api.GetSomeUserInfo()
	// do sth.
}
