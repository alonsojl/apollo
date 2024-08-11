package server

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayFunc func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type Server struct {
	productHandler *ProductHandler
}

type Option func(*Server)

func WithProductHandler(productHandler *ProductHandler) Option {
	return func(s *Server) {
		s.productHandler = productHandler
	}
}

func New(options ...Option) *Server {
	s := &Server{}

	for _, option := range options {
		option(s)
	}

	return s
}
