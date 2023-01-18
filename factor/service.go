package factor

import "api-examples/user"

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
	GetSomeFactorInfo = s.GetSomeFactorInfo
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
	user.GetSomeUserInfo()
	// do sth.
}
