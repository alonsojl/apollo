package internal

type ProductParams struct {
	Name         string
	Price        float64
	QRCode       string
	Image        []byte
	Location     string
	CategoryUUID string
}

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
