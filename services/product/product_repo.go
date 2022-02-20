package product

import "gorm.io/gorm"

type ProductRepo interface {
	Find() (Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Find() (Product, error) {
	var product Product
	err := r.db.First(&product).Error
	return product, err
}
