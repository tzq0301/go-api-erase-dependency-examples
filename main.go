package main

import (
	"api-examples/event"
	"api-examples/factor"
	"api-examples/user"
)

func main() {
	//userService := user.NewService()
	//factorService := factor.NewService(userService)
	//eventService := event.NewService(factorService)

	_ = user.NewService()
	_ = factor.NewService()
	_ = event.NewService()
}
