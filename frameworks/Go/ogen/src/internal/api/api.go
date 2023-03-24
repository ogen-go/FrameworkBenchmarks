package api

import (
	"context"

	"ogen/app/internal/oas"
)

type Server struct {
	oas.UnimplementedHandler
}

func (s Server) JSON(ctx context.Context) (*oas.HelloWorld, error) {
	return &oas.HelloWorld{Message: "Hello, World!"}, nil
}

var _ oas.Handler = (*Server)(nil)
