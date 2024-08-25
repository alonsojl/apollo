package domain

import (
	"apollo/internal/types"
	"context"
)

type Product struct {
	UUID         string
	Name         string
	Price        float64
	QRCode       string
	Location     string
	CategoryUUID string
	CreatedAt    string
	UpdatedAt    string
}

type ProductService interface {
	CreateProduct(ctx context.Context, params *types.ProductParams) (*Product, error)
}

type ProductDBRepository interface {
	CreateProduct(ctx context.Context, params *types.ProductParams) (*Product, error)
}

type ProductS3Repository interface {
	UploadProductImage(ctx context.Context, imageBuffer []byte) (string, error)
}
