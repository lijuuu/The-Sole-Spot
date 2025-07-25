package usecase

import (
	"ecommerce_clean_arch/pkg/domain"
	"ecommerce_clean_arch/pkg/repository/interfaces"
	"ecommerce_clean_arch/pkg/utils/models"
	"errors"
)

type ProductUseCase struct {
	ProductRepository interfaces.ProductRepository
}

func NewProductUseCase(productRepo interfaces.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *ProductUseCase) AddProduct(product models.AddProduct) (models.ProductResponse, error) {
	if product.Price < 0 || product.Quantity < 0 {
		return models.ProductResponse{}, errors.New("invalid quantity or price")
	}
	products, err := p.ProductRepository.AddProduct(product)
	if err != nil {
		return models.ProductResponse{}, err
	}
	productResponse := models.ProductResponse{
		ID:          products.ID,
		Category_Id: products.Category_Id,
		Name:        products.Name,
		Stock:       products.Stock,
		Price:       products.Price,
		Quantity:    products.Quantity,
		OfferPrice:  product.OfferPrice,
	}

	return productResponse, nil
}

func (p *ProductUseCase) UpdateProduct(products models.ProductResponse, productID int) (models.ProductResponse, error) {

	if products.Price < 0 || products.Quantity < 0 {
		return models.ProductResponse{}, errors.New("invalid quantity or price")
	}

	updateProduct, err := p.ProductRepository.UpdateProduct(products, productID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return updateProduct, nil
}

func (p *ProductUseCase) DeleteProduct(productID int) error {
	err := p.ProductRepository.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUseCase) SearchProduct(categoryID string, sortBy string) ([]domain.Products, error) {
	return p.ProductRepository.GetProductsByCategory(categoryID, sortBy)
}
