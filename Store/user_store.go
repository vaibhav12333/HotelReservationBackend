package DB

import (
	constants "HotelReservationBackend/Constants"
	"HotelReservationBackend/Types"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*Types.User, error)
	GetUsers(ctx context.Context) ([]*Types.User, error)
	CreateUser(ctx context.Context, user *Types.User) (*Types.User, error)
}

type MongoUserStore struct {
	db *sql.DB
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*Types.User, error) {
	user := &Types.User{}
	resp := s.db.QueryRow("SELECT id,email,firstname,lastname,password FROM users WHERE id=$1", id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.EncryptedPassword)
	if resp == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*Types.User, error) {

	return nil, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *Types.User) (*Types.User, error) {

	return nil, nil
}
func NewMongoStore() *MongoUserStore {
	connStr := constants.PsqlConnect

	db, err := sql.Open(constants.DriverName, connStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to PostgresSql")
	defer db.Close()
	return &MongoUserStore{}
}
