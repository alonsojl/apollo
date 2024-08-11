package server

import (
	"apolo/internal"
	"apolo/internal/models"
	"apolo/pkg/errorx"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type errorResponse struct {
	models.TemplateResponse
	Message string            `json:"message"`
	Errors  validation.Errors `json:"errors,omitempty" swaggertype:"object"`
}

func ErrorHandler(err error) (events.APIGatewayProxyResponse, error) {
	var (
		verr     validation.Errors
		werr     errorx.Error
		response errorResponse
	)

	response.Datetime = time.Now().Format(time.DateTime)
	response.Timestamp = time.Now().Unix()

	if !errors.As(err, &werr) {
		response.Status = "error"
		response.Message = "internal server error"
		response.Code = http.StatusInternalServerError
	} else {
		response.Status = "fail"
		response.Message = werr.Message()

		switch werr.Code() {
		case internal.CodeInvalidArgument:
			response.Code = http.StatusBadRequest
			if errors.As(werr, &verr) {
				response.Errors = verr
			}
		case internal.CodeUnauthorized:
			response.Code = http.StatusUnauthorized
		case internal.CodeNotFound:
			response.Code = http.StatusNotFound
		}
	}
	return JSON(response, response.Code)
}

func JSON(response interface{}, statusCode int) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
