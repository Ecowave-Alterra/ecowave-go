package product

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (pr *productRepo) CreateProduct(product *pe.Product) error {
	if err := pr.db.Save(&product).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) CheckProductExist(productId string) (bool, error) {
	var count int64
	result := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Count(&count)
	if result.Error != nil {
		return false, echo.NewHTTPError(500, result.Error)
	}

	exists := count > 0
	return exists, nil
}

func (pr *productRepo) CreateProductImage(productImage *pe.ProductImage) error {
	if err := pr.db.Save(&productImage).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) GetAllProduct(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	var count int64
	if err := pr.db.Model(&pe.Product{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := pr.db.
		Preload("ProductCategory").Preload("ProductImages").
		Offset(offset).
		Limit(pageSize).
		Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *products, count, nil
}

func (pr *productRepo) GetAllProductNoPagination(products *[]pe.Product) ([]pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").
		Find(&products).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return *products, nil
}

func (pr *productRepo) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	if err := pr.db.
		Preload("ProductCategory").Preload("ProductImages").
		Where("product_id = ?", productId).
		First(&product).Error; err != nil {
		return *product, echo.NewHTTPError(404, err)
	}

	return *product, nil
}

func (pr *productRepo) GetProductImageURLById(productId string, productImage *pe.ProductImage) ([]pe.ProductImage, error) {
	var productImages []pe.ProductImage
	if err := pr.db.Model(&pe.ProductImage{}).Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		return productImages, echo.NewHTTPError(404, err)
	}
	return productImages, nil
}

func (pr *productRepo) UpdateProduct(productId string, req *pe.ProductRequest) error {
	if err := pr.db.Model(&pe.Product{}).Where("product_id = ?", productId).Updates(pe.Product{ProductCategoryId: req.ProductCategoryId, Name: req.Name, Price: req.Price, Status: req.Status, Description: req.Description}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) UpdateProductStock(productId string, stock uint) error {
	if err := pr.db.Exec("UPDATE products SET stock = ? WHERE product_id = ?", stock, productId).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) DeleteProduct(productId string, product *pe.Product) error {
	if err := pr.db.Where("product_id = ?", productId).Delete(&product).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) DeleteProductImage(productID string, productImages *[]pe.ProductImage) error {
	if err := pr.db.Where("product_id = ?", productID).Delete(&productImages).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) DeleteProductImageByID(ProductImageID uint, productImage *pe.ProductImage) error {
	if err := pr.db.Where("id = ?", ProductImageID).Delete(productImage).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (pr *productRepo) SearchProduct(search, filter string, offset, pageSize int) (*[]pe.Product, int64, error) {
	var products []pe.Product
	var count int64

	if err := pr.db.Model(&pe.Product{}).
		Where("name LIKE ? OR product_id LIKE ? OR product_category_id IN (?)",
			"%"+search+"%",
			"%"+search+"%",
			pr.db.Model(&pe.ProductCategory{}).Select("id").Where("category LIKE ?", "%"+search+"%")).
		Where("status LIKE ?", "%"+filter+"%").
		Preload("ProductCategory").Preload("ProductImages").
		Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := pr.db.Model(&pe.Product{}).
		Where("name LIKE ? OR product_id LIKE ? OR product_category_id IN (?)",
			"%"+search+"%",
			"%"+search+"%",
			pr.db.Model(&pe.ProductCategory{}).Select("id").Where("category LIKE ?", "%"+search+"%")).
		Where("status LIKE ?", "%"+filter+"%").
		Preload("ProductCategory").Preload("ProductImages").
		Offset(offset).Limit(pageSize).
		Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return &products, count, nil
}