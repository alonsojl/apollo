package s3

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type ProductRepository struct {
	logger   *slog.Logger
	s3Client *s3.Client
	bucket   string
}

func NewProductRepository(logger *slog.Logger, s3Client *s3.Client) *ProductRepository {
	bucket := os.Getenv("BUCKET")
	return &ProductRepository{
		logger:   logger,
		s3Client: s3Client,
		bucket:   bucket,
	}
}

func (r *ProductRepository) UploadProductImage(ctx context.Context, imageBuffer []byte) (string, error) {
	var (
		filename = r.getFilename(imageBuffer)
		filepath = fmt.Sprintf("product-images/%s", filename)
	)

	uploader := manager.NewUploader(r.s3Client)
	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(filepath),
		Body:        bytes.NewReader(imageBuffer),
		ContentType: aws.String(http.DetectContentType(imageBuffer)),
	})
	if err != nil {
		r.logger.Error("error uploading product image", "error", err)
		return "", err
	}

	r.logger.Info("uploaded product image", "location", result.Location)
	return result.Location, nil

}

func (r *ProductRepository) getFilename(imageBuffer []byte) string {
	mimeType := http.DetectContentType(imageBuffer)
	extension := strings.Split(mimeType, "/")[1]
	return fmt.Sprintf("%s.%s", uuid.New().String(), extension)
}
