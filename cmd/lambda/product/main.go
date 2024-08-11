package main

import (
	"apolo/internal/dynamodb"
	"apolo/internal/s3"
	"apolo/internal/server"
	"apolo/internal/services"
	"apolo/pkg/env"
	"log"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

// @title 			Go RESTful API
// @version 		1.0
// @description 	Testing Swagger APIs.
// @termsOfService 	http://swagger.io/terms/

// @schemes 		https
// @host            traqlss9y8.execute-api.us-east-1.amazonaws.com
// @basePath 		/Dev/v1/cybersource

// @contact.name 	API Support
// @contact.url		http://www.swagger.io/support
// @contact.email 	support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apiKey JWT
// @in	 header
// @name Authorization
func main() {
	svr, err := run()
	if err != nil {
		log.Fatalf("error running lambda server: %v", err)
	}
	lambda.Start(svr.Routes())
}

func run() (*server.Server, error) {
	if err := env.Validate(getEnvs()); err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	db, err := dynamodb.Connection()
	if err != nil {
		return nil, err
	}

	s3Client, err := s3.Connection()
	if err != nil {
		return nil, err
	}

	productS3Repository := s3.NewProductRepository(logger, s3Client)
	productDBRepository := dynamodb.NewProductRepository(db, logger)
	productService := services.NewProductService(logger, productDBRepository, productS3Repository)
	productHandler := server.NewProductHandler(logger, productService)

	svr := server.New(
		server.WithProductHandler(productHandler),
	)

	return svr, nil
}

func getEnvs() []string {
	return []string{
		"LOG_LEVEL",
	}
}
