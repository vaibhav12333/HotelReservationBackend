package Types

import "golang.org/x/crypto/bcrypt"

const (
	bcryptCost = 12
)

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	Id                int    `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"first_name" json:"firstname"`
	LastName          string `bson:"last_name" json:"lastname"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encPass" json:"password"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		EncryptedPassword: string(pass),
		Email:             params.Email,
	}, nil
}
