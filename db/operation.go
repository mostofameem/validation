package db

import (
	"errors"
	"fmt"
)

func Register(name, email, pass string) error {
	dbpass := GetPass(name)
	if dbpass == "" {
		err := INSERT(name, email, pass)
		return err
	}
	return fmt.Errorf("user already exists")
}
func Login(name string, pass string) error {

	dbpass := GetPass(name)
	if dbpass == pass {
		return nil
	}
	return errors.New("failed ")
}
