package service

import "df-ecomm/pkg/model"

type Product = model.Product

var products []Product = []Product{
	{Id: 1, Name: "Product 1", Price: 100},
	{Id: 2, Name: "Product 2", Price: 200},
	{Id: 3, Name: "Product 3", Price: 300},
}

func (r *Repository) GetAll() []Product {
	return products
}

func (r *Repository) Create(product model.NewProduct) Product {
	newProduct := Product{
		Id:    products[len(products)-1].Id + 1,
		Name:  product.Name,
		Price: product.Price,
	}
	products = append(products, newProduct)
	return newProduct
}

func (r *Repository) UpdateById(productIdToUpdate int, updateProduct Product) (Product, error) {
	for index, product := range products {
		if product.Id == productIdToUpdate {
			if updateProduct.Name != "" {
				products[index].Name = updateProduct.Name
			}

			if updateProduct.Price >= 0 {
				products[index].Price = updateProduct.Price
			}

			return products[index], nil
		}
	}
	return Product{}, model.ErrNotFound
}

func (r *Repository) DeleteById(productIdToDelete int) error {
	for index, product := range products {
		if product.Id == productIdToDelete {
			products = append(products[:index], products[index+1:]...)
			return nil
		}
	}
	return model.ErrNotFound
}
