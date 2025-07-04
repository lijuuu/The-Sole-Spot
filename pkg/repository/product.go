package repository

import (
	"ecommerce_clean_arch/pkg/domain"
	"ecommerce_clean_arch/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: DB,
	}
}

func (p *ProductRepository) AddProduct(product models.AddProduct) (models.ProductResponse, error) {
	var productResponse models.ProductResponse

	err := p.DB.Raw(
		`INSERT INTO products (category_id, name, stock, quantity, price, offer_price) 
        VALUES (?, ?, ?, ?, ?, ?) 
        RETURNING id, category_id, name, stock, quantity, price, offer_price`,
		product.CategoryID, product.Name, product.Stock, product.Quantity, product.Price, product.OfferPrice).Scan(&productResponse).Error

	if err != nil {
		return models.ProductResponse{}, err
	}
	return productResponse, nil
}

func (p *ProductRepository) UpdateProduct(product models.ProductResponse, productID int) (models.ProductResponse, error) {
	var updatedProduct models.ProductResponse
	err := p.DB.Raw(
		`UPDATE products SET 
            category_id = $1, 
            name = $2, 
            stock = $3, 
            quantity = $4, 
            price = $5, 
            offer_price = $6 
        WHERE id = $7 
        RETURNING id, category_id, name, stock, quantity, price, offer_price`,
		product.Category_Id,
		product.Name,
		product.Stock,
		product.Quantity,
		product.Price,
		product.OfferPrice,
		productID,
	).Scan(&updatedProduct).Error

	if err != nil {
		return models.ProductResponse{}, fmt.Errorf("error updating product: %w", err)
	}

	return updatedProduct, nil
}

func (p *ProductRepository) DeleteProduct(productID int) error {
	var products domain.Products
	err := p.DB.Where("id =?", productID).Delete(&products)
	if err.RowsAffected < 1 {
		return errors.New("the id is not existing")
	}
	return nil
}

func (p *ProductRepository) GetProductByID(productID int) (models.ProductResponse, error) {

	var productResponse models.ProductResponse
	err := p.DB.Raw("SELECT * FROM products WHERE id = ?", productID).Scan(&productResponse).Error
	if err != nil {
		return models.ProductResponse{}, err
	}
	return productResponse, nil
}

func (p *ProductRepository) UpdateStock(productID, qty int) error {
	return p.DB.Model(&models.ProductResponse{}).
		Where("id = ?", productID).
		UpdateColumn("stock", gorm.Expr("stock - ?", qty)).Error
}

func (p *ProductRepository) GetAllProducts(showOutOfStock bool) ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	query := p.DB
	if !showOutOfStock {
		query = query.Where("stock > 0")
	}
	err := query.Find(&products).Error
	return products, err
}

func (p *ProductRepository) GetProductsByCategory(categoryID string, sortBy string) ([]domain.Products, error) {
	var products []domain.Products
	query := p.DB.Model(&domain.Products{}).Where("category_id = ? AND  stock > 0", categoryID)

	switch sortBy {
	case "price_H-L":
		query = query.Order("price desc")
	case "price_L-H":
		query = query.Order("price asc")
	case "newest":
		query = query.Order("created_at desc")
	case "alphabetic":
		query = query.Order("LOWER(name) ASC")
	default:
		query = query.Order("created_at desc")
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
