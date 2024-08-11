package services

import (
	"apolo/internal"
	"context"
	"log/slog"
)

type ProductDBRepository interface {
	CreateProduct(ctx context.Context, params *internal.ProductParams) (*internal.Product, error)
}

type ProductS3Repository interface {
	UploadProductImage(ctx context.Context, imageBuffer []byte) (string, error)
}

type ProductService struct {
	logger              *slog.Logger
	productDBRepository ProductDBRepository
	productS3Repository ProductS3Repository
}

func NewProductService(logger *slog.Logger, productRepository ProductDBRepository, productS3Repository ProductS3Repository) *ProductService {
	return &ProductService{
		logger:              logger,
		productDBRepository: productRepository,
		productS3Repository: productS3Repository,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, params *internal.ProductParams) (*internal.Product, error) {
	location, err := s.productS3Repository.UploadProductImage(ctx, params.Image)
	if err != nil {
		s.logger.Error("error uploading product image", "error", err)
		return nil, err
	}

	params.Location = location
	product, err := s.productDBRepository.CreateProduct(ctx, params)
	if err != nil {
		s.logger.Error("error creating product", "error", err)
		return nil, err
	}

	return product, nil
}
