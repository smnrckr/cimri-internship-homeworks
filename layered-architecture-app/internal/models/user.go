package models

import "errors"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u UserCreateRequest) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Age < 0 {
		return errors.New("age must be greater than 0")
	}

	return nil
}

type UserUpdateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
