package main

import (
	"fmt"
)

type user struct {
	ID int
	Name,
	User,
	Password string
}

func (u user) insert() error {
	// ID set to NULL so db autoincrements it

	// !!! DOES NOT WORK
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO users (ID, NAME, EMAIL, PASSWORD) VALUES (NULL, '%s', '%s', '%s');", u.Name, u.User, u.Password))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
