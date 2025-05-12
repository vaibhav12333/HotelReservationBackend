package main

import (
	"HotelReservationBackend/API"
	DB "HotelReservationBackend/Store"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()
	store := DB.NewMongoStore()
	userHandler := API.NewUserHandler(store)

	//Routes
	app.Get("/user/:id", userHandler.GetUser)
	app.Get("/users", userHandler.GetUsers)

	//Server Running
	if err := app.Listen(":7878"); err != nil {
		log.Fatal(err)
	}
}
