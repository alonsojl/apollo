package models

import (
	"apollo/internal"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ProductCreateRequest struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	QRCode       string  `json:"qr_code"`
	Image        string  `json:"image"`
	CategoryUUID string  `json:"category_uuid"`
}

func (p ProductCreateRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name,
			validation.Required,
			validation.Length(1, 100),
			internal.Alphanumeric,
		),
		validation.Field(&p.Price,
			validation.Required,
			validation.Min(0.01),
		),
		validation.Field(&p.QRCode,
			validation.Required,
			validation.Length(1, 50),
			internal.Alphanumeric,
		),
		validation.Field(&p.Image,
			validation.Required,
			is.Base64,
		),
		validation.Field(&p.CategoryUUID,
			validation.Required,
			is.UUID,
		),
	)
}

type ProductResponse struct {
	TemplateResponse
	Product *CreatedProduct
}

type CreatedProduct struct {
	UUID         string  `json:"uuid"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	QRCode       string  `json:"qr_code"`
	Location     string  `json:"location"`
	CategoryUUID string  `json:"category_uuid"`
	CreatedAt    string  `json:"created_at"`
}
