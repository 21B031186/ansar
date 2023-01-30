package models

import (
	"fmt"
)

type User struct {
	Name string
	Age uint8
}
func (u User) ToString() string {
	return fmt.Sprintf("Name: %s, age: %d", u.Name, u.Age)
}