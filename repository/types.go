// This file contains types that are used in the repository layer.
package repository

type GetUserOutput struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
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
	ID int
}
