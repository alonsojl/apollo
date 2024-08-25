package server

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func (s *Server) Routes() APIGatewayFunc {
	return func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch event.HTTPMethod {
		case http.MethodPost:
			return s.productHandler.HandleCreateProduct(ctx, event)
		}
		return events.APIGatewayProxyResponse{
			Body:       "Not found",
			StatusCode: http.StatusNotFound,
		}, nil
	}
}
