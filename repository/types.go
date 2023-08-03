// This file contains types that are used in the repository layer.
package repository

type GetUserInput struct {
	Phone string
}

type GetUserOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}

type UpdateUserInput struct {
	FullName string
	Phone    string
}

type UpdateUserOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}

type InsertUserInput struct {
	FullName string
	Phone    string
	Password string
}

type InsertUserOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}

type CheckUsernamePasswordProfileInput struct {
	Phone    string
	Password string
}

type CheckUsernamePasswordProfileOutput struct {
	FullName string
	Phone    string
	IsExist  bool `json:"isExist"`
}
