package server

import (
	"apollo/internal/domain"
	"apollo/internal/models"
	"apollo/internal/types"
	"context"
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ProductHandler struct {
	logger         *slog.Logger
	productService domain.ProductService
}

func NewProductHandler(logger *slog.Logger, productService domain.ProductService) *ProductHandler {
	return &ProductHandler{
		logger:         logger,
		productService: productService,
	}
}

// @Summary 	Create product.
// @Description
// @Tags 		Product
// @Router 		/products [post]
// @Accept 		json
// @Produce 	json
// @Security    JWT
// @Param	    request body models.ProductCreateRequest true "Datos necesarios"
// @Success     200	{object} models.ProductResponse "Success"
// @Failure     400	{object} server.errorResponse "Bad Request"
// @Failure     401	{object} server.errorResponse "Unauthorized"
// @Failure     412	{object} server.errorResponse "Precondition"
// @Failure     500	{object} server.errorResponse "Error Internal Server"
func (h *ProductHandler) HandleCreateProduct(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request models.ProductCreateRequest

	if err := json.Unmarshal([]byte(event.Body), &request); err != nil {
		h.logger.Error("invalid product request", "error", err)
		return ErrorHandler(types.ErrInvalidBody)
	}

	if err := request.Validate(); err != nil {
		h.logger.Error("invalid product params", "error", err)
		return ErrorHandler(types.ErrInvalidParams)
	}

	image, err := base64.StdEncoding.DecodeString(request.Image)
	if err != nil {
		h.logger.Error("error decoding image", "error", err)
		return ErrorHandler(types.ErrInvalidBody)
	}

	product, err := h.productService.CreateProduct(ctx, &types.ProductParams{
		Name:         request.Name,
		Price:        request.Price,
		QRCode:       request.QRCode,
		Image:        image,
		CategoryUUID: request.CategoryUUID,
	})

	if err != nil {
		h.logger.Error("error creating product", "error", err)
		return ErrorHandler(err)
	}

	var response = models.ProductResponse{
		TemplateResponse: models.NewTemplateResponse(),
		Product: &models.CreatedProduct{
			UUID:         product.UUID,
			Name:         product.Name,
			Price:        product.Price,
			QRCode:       product.QRCode,
			Location:     product.Location,
			CategoryUUID: product.CategoryUUID,
			CreatedAt:    product.CreatedAt,
		},
	}

	return JSON(response, http.StatusOK)
}
