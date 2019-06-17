package models

import (
	"errors"
	"log"
)

type Weight struct {
	Value int    `json:"value,omitempty"`
	Date  string `json:"date,omitempty"`
}

type User struct {
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	Weights  []Weight `json:"weights,omitempty"`
}

type Model struct {
	Db *[]User `json:"db,omitempty"`
} //do modelu dodac połączenie z bd, ale zamockować jakąś gówno baze ze zmiennej na początek

func (model Model) PostUser(user User) User {
	*model.Db = append(*model.Db, user)
	log.Println("add user")
	return user
}

func (model Model) GetAllUsers() []User {
	log.Println("read all users")
	return *model.Db
}

func (model Model) LoginUser(user User) bool {
	for _, u := range *model.Db {
		if u.Email == user.Email && u.Password == user.Password {
			return true
		}
	}
	return false
}

func (model Model) LogoutUser(user User) error {

	return errors.New("eror zastepczy")
}

func (model Model) AddWeight(user User, weight Weight) {
	user.Weights = append(user.Weights, weight)
}
