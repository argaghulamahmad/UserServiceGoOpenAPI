// This file contains types that are used in the repository layer.
package repository

type GetProfileInput struct {
	Phone string
}

type GetProfileOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}

type UpdateProfileInput struct {
	Phone string
}

type UpdateProfileOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}

type InsertProfileInput struct {
	FullName string
	Phone    string
	Password string
}

type InsertProfileOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}
