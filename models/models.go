package models

import (
	"errors"
)

type Weight struct {
	Value int
	Date  string
}

type User struct {
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	Weights  []Weight `json:"weights,omitempty"`
}

type Model struct {
	Db []User
} //do modelu dodac połączenie z bd, ale zamockować jakąś gówno baze ze zmiennej na początek

func (model Model) PostUser(user User) (User, error) {
	model.Db = append(model.Db, user)
	return user, errors.New("eror zastepczy")
}

func (model Model) LoginUser(user User) (User, error) {

	return user, errors.New("eror zastepczy")
}

func (model Model) LogoutUser(user User) error {

	return errors.New("eror zastepczy")
}

func (model Model) AddWeight(user User, weight Weight) error {
	return errors.New("err zastepczy")
}
