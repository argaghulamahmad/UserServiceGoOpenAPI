// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetProfile(ctx context.Context, input GetProfileInput) (output GetProfileOutput, err error)
	UpdateProfile(ctx context.Context, input UpdateProfileInput) (output UpdateProfileOutput, err error)
	InsertProfile(ctx context.Context, input InsertProfileInput) (output InsertProfileOutput, err error)
}
