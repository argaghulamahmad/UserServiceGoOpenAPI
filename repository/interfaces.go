// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	IsPhonePasswordUserExist(ctx context.Context, input CheckUsernamePasswordProfileInput) (output CheckUsernamePasswordProfileOutput, err error)
	GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (output UpdateUserOutput, err error)
	InsertUser(ctx context.Context, input InsertUserInput) (output InsertUserOutput, err error)
}
