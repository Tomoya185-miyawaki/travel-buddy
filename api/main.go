package main

import (
	"github.com/Tomoya185-miyawaki/travel-buddy/db"
	"github.com/Tomoya185-miyawaki/travel-buddy/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation/adapter"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation/validation"
	"github.com/Tomoya185-miyawaki/travel-buddy/usecase"
)

func main() {
	db := db.NewDB()
	userValidator := validation.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := adapter.NewUserController(userUsecase)
	presentation.Route(userController)
}
