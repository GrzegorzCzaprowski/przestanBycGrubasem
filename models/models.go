package models

import (
	"errors"
)

type Weights map[string]int

type User struct {
	Email    string
	Password string
	Weights  map[string]int
}

type Model int //do modelu dodał połączenie z bd, ale zamockować jakąś gówno baze ze zmiennej na początek

func (model Model) PostUser(user User) (User, error) {

	return user, errors.New("eror zastepczy")
}

func (model Model) LogUser(user User) (User, error) {

	return user, errors.New("eror zastepczy")
}
