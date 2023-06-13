package product

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pc *productUseCase) CreateProduct(product *pe.Product) error {
	for {
		productId := randomid.GenerateRandomID()

		exists, err := pc.productRepo.CheckProductExist(productId)
		if err != nil {
			return err
		}
		if !exists {
			product.ProductID = productId
			break
		}
	}
	return pc.productRepo.CreateProduct(product)
}

func (pc *productUseCase) CreateProductImage(productImage *pe.ProductImage) error {
	return pc.productRepo.CreateProductImage(productImage)
}

func (pc *productUseCase) GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	product, count, err := pc.productRepo.GetAllProduct(products, offset, pageSize)
	return product, count, err
}

func (pc *productUseCase) GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error) {
	return pc.productRepo.GetAllProductNoPagination(products)
}

func (pc *productUseCase) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	return pc.productRepo.GetProductByID(productId, product)
}

func (pc *productUseCase) GetProductImageURLById(productId string, productImage *pe.ProductImage) ([]pe.ProductImage, error) {
	return pc.productRepo.GetProductImageURLById(productId, productImage)
}

func (pc *productUseCase) UpdateProduct(productId string, productRequest *pe.ProductRequest) error {
	return pc.productRepo.UpdateProduct(productId, productRequest)
}

func (pc *productUseCase) UpdateProductStock(productId string, stock uint) error {
	return pc.productRepo.UpdateProductStock(productId, stock)
}

func (pc *productUseCase) DeleteProduct(productId string, product *pe.Product) error {
	return pc.productRepo.DeleteProduct(productId, product)
}

func (pc *productUseCase) DeleteProductImage(productID string, productImages *[]pe.ProductImage) error {
	return pc.productRepo.DeleteProductImage(productID, productImages)
}

func (pc *productUseCase) DeleteProductImageByID(ProductImageID string, productImage *pe.ProductImage) error {
	return pc.productRepo.DeleteProductImageByID(ProductImageID, productImage)
}

func (pc *productUseCase) SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error) {
	products, count, err := pc.productRepo.SearchProduct(search, filter, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}

// func (pc *productUseCase) SearchProductByID(productID string, product *pe.Product) (pe.Product, error) {
// 	return pc.productRepo.SearchProductByID(productID, product)
// }

// func (pc *productUseCase) SearchProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
// 	return pc.productRepo.SearchProductByName(name, product)
// }

// func (pc *productUseCase) SearchProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
// 	return pc.productRepo.SearchProductByCategory(category, product)
// }

// func (pc *productUseCase) FilterProductByStatus(status string, product *[]pe.Product) ([]pe.Product, error) {
// 	return pc.productRepo.FilterProductByStatus(status, product)
// }