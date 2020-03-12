package model

import "github.com/omnisinc/queen/pkg/validator"

func NewAIImage(
	productID	int64,
	imageURL	string,
	order	int64,
	category	string,
	scaleAB,
	scaleBD,
	scaleDE,
	scaleEA	int64,
) (*AIImage, error){
	target := &AIImage{
		ProductID: productID,
		ImageURL: imageURL,
		Order: order,
		Category: category,
		ScaleAB: scaleAB,
		ScaleBD: scaleBD,
		ScaleDE: scaleDE,
		ScaleEA: scaleEA,
	}
	if err := validator.Validator.Struct(target);err!=nil{
		return err
	}
	return target, nil
}
