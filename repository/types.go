// This file contains types that are used in the repository layer.
package repository

type GetProfileByPhoneNumberInput struct {
	PhoneNumber string
}

type GetProfileByPhoneNumberOutput struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
}
