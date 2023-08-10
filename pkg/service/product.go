package service

import (
	"gorm.io/gorm"

	"df-ecomm/pkg/model"
)

func (r *Repository) GetAll() (products []Product, err error) {
	if err = r.Db.Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *Repository) Create(newProduct Product) (createdProduct Product, err error) {
	createdProduct = Product{
		Name:  newProduct.Name,
		Price: newProduct.Price,
	}

	if err = r.Db.Create(&createdProduct).Error; err != nil {
		return Product{}, err
	}

	return createdProduct, nil
}

func (r *Repository) UpdateById(productIdToUpdate int, productToUpdate Product) (updatedProduct Product, err error) {
	if err = r.Db.First(&updatedProduct, productIdToUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Product{}, model.ErrProductNotFound
		}
		return Product{}, err
	}

	if productToUpdate.Name != "" {
		updatedProduct.Name = productToUpdate.Name
	}

	if productToUpdate.Price != 0 {
		updatedProduct.Price = productToUpdate.Price
	}

	if err = r.Db.Save(&updatedProduct).Error; err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func (r *Repository) DeleteById(productIdToDelete int) (err error) {
	return r.Db.Delete(&Product{}, productIdToDelete).Error
}
