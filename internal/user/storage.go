package user

import "context"

// Storage An enumeration of the methods to be implemented.
type Storage interface {
	Create(ctx context.Context, user User) (string, error)
	FindOne(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
