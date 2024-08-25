package dynamodb

import (
	"apollo/internal/domain"
	"apollo/internal/types"
	"context"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type PruductSchema struct {
	UUID         string  `dynamodbav:"uuid"`
	Name         string  `dynamodbav:"name"`
	Price        float64 `dynamodbav:"price"`
	QRCode       string  `dynamodbav:"qrcode"`
	Location     string  `dynamodbav:"location"`
	CategoryUUID string  `dynamodbav:"category_uuid"`
	CreatedAt    string  `dynamodbav:"created_at"`
	UpdatedAt    string  `dynamodbav:"updated_at"`
}

type ProductRepository struct {
	db     *dynamodb.Client
	logger *slog.Logger
	table  string
}

func NewProductRepository(db *dynamodb.Client, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{
		db:     db,
		logger: logger,
		table:  "products",
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, params *types.ProductParams) (*domain.Product, error) {
	var (
		uuid      = uuid.New().String()
		createdAt = time.Now().UTC().Format(time.DateTime)
	)
	item, err := attributevalue.MarshalMap(PruductSchema{
		UUID:         uuid,
		Name:         params.Name,
		Price:        params.Price,
		QRCode:       params.QRCode,
		Location:     params.Location,
		CategoryUUID: params.CategoryUUID,
		CreatedAt:    createdAt,
	})
	if err != nil {
		r.logger.Error("error marshaling product", "error", err)
		return nil, err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      item,
	})

	if err != nil {
		r.logger.Error("error creating product", "error", err)
		return nil, err
	}

	return &domain.Product{
		UUID:         uuid,
		Name:         params.Name,
		Price:        params.Price,
		QRCode:       params.QRCode,
		Location:     params.Location,
		CategoryUUID: params.CategoryUUID,
		CreatedAt:    createdAt,
	}, nil
}
