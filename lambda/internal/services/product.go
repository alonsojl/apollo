package services

import (
	"apollo/internal/domain"
	"apollo/internal/types"
	"context"
	"log/slog"
)

type ProductService struct {
	logger              *slog.Logger
	productDBRepository domain.ProductDBRepository
	productS3Repository domain.ProductS3Repository
}

func NewProductService(logger *slog.Logger, productRepository domain.ProductDBRepository, productS3Repository domain.ProductS3Repository) *ProductService {
	return &ProductService{
		logger:              logger,
		productDBRepository: productRepository,
		productS3Repository: productS3Repository,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, params *types.ProductParams) (*domain.Product, error) {
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
