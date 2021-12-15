package user

import (
	"context"
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, createUserDTO CreateUserDTO) (u User, err error) {
	// TODO for next one
	return u, err
}
