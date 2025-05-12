package API

import (
	DB "HotelReservationBackend/Store"
	"HotelReservationBackend/Types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db DB.UserStore
}

func NewUserHandler(db DB.UserStore) *UserHandler {
	return &UserHandler{db: db}
}

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := uh.db.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uh.db.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (uh *UserHandler) PostUser(c *fiber.Ctx) error {
	var puser Types.CreateUserParams
	err := c.BodyParser(&puser)
	if err != nil {
		return err
	}
	user, err := Types.NewUserFromParams(puser)
	if err != nil {
		return err
	}
	insertedUser, err := uh.db.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}
