package product

import (
	"github.com/berrylradianh/ecowave-go/helper/randomid"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pc *productUseCase) CreateProduct(product *ep.Product) error {
	for {
		productId := randomid.GenerateRandomNumber()

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

func (pc *productUseCase) CreateProductImage(productImage *ep.ProductImage) error {
	return pc.productRepo.CreateProductImage(productImage)
}

func (pc *productUseCase) GetAllProduct(products *[]ep.Product) ([]ep.Product, error) {
	return pc.productRepo.GetAllProduct(products)
}

func (pc *productUseCase) GetProductByID(productId string, product *ep.Product) (ep.Product, error) {
	return pc.productRepo.GetProductByID(productId, product)
}

func (pc *productUseCase) GetProductImageURLById(productId string, productImage *ep.ProductImage) ([]ep.ProductImage, error) {
	return pc.productRepo.GetProductImageURLById(productId, productImage)
}

func (pc *productUseCase) UpdateProduct(productId string, productRequest *ep.ProductRequest) error {
	return pc.productRepo.UpdateProduct(productId, productRequest)
}

func (pc *productUseCase) UpdateProductStock(productId string, stock uint) error {
	return pc.productRepo.UpdateProductStock(productId, stock)
}

func (pc *productUseCase) DeleteProduct(productId string, product *ep.Product) error {
	return pc.productRepo.DeleteProduct(productId, product)
}

func (pc *productUseCase) DeleteProductImage(productID string, productImages *[]ep.ProductImage) error {
	return pc.productRepo.DeleteProductImage(productID, productImages)
}

func (pc productUseCase) DeleteProductImageByID(ProductImageID string, productImage *ep.ProductImage) error {
	return pc.productRepo.DeleteProductImageByID(ProductImageID, productImage)
}

func (pc *productUseCase) SearchProductByID(productID string, product *ep.Product) (ep.Product, error) {
	return pc.productRepo.SearchProductByID(productID, product)
}

func (pc *productUseCase) SearchProductByName(name string, product *[]ep.Product) ([]ep.Product, error) {
	return pc.productRepo.SearchProductByName(name, product)
}

func (pc *productUseCase) SearchProductByCategory(category string, product *[]ep.Product) ([]ep.Product, error) {
	return pc.productRepo.SearchProductByCategory(category, product)
}

func (pc *productUseCase) FilterProductByStatus(status string, product *[]ep.Product) ([]ep.Product, error) {
	return pc.productRepo.FilterProductByStatus(status, product)
}
