package model

type AIImage struct {
	ProductID int64  `json:"productId" validate:"required"`
	ImageURL  string `json:"imageUrl" validate:"required,url,max=2000"`
	Order     int64  `json:"order" validate:"oneof=0 1"`
	Category  string `json:"category" validate:"required"`
	ScaleAB   int64  `json:"scale_AB" validate:"required,max=300,min=1"`
	ScaleBD   int64  `json:"scale_BD" validate:"required,max=300,min=1"`
	ScaleDE   int64  `json:"scale_DE" validate:"required,max=300,min=1"`
	ScaleEA   int64  `json:"scale_EA" validate:"required,max=300,min=1"`
}
