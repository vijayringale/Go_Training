package main

import (
	"fmt"

	"example.com/fake/users/db"
)

// “fake” database that we import from "example.com/fake/users/db"
func findUser(username string) (*db.User, error) {
	return db.Find(username)
}

func SetUserAge(u *db.User, age int) error {
	return db.SetAge(u, age)
}

func findAndSetUserAge(username string, age int) error {
	var user *db.User
	var err error

	user, err = findUser(username)
	if err != nil {
		return err
	}
	if err = SetUserAge(user, age); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := findAndSetUserAge("b0b@example.com", 21); err != nil {
		fmt.Println("faild finding or updating user :%s", err)
		return
	}
	fmt.Println("successfully updated user's age")
}
