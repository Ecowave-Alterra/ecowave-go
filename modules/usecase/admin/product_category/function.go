package product_category

import (
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (pcc *productCategoryUsecase) CreateProductCategory(productCategory *pe.ProductCategory) (bool, error) {
	if err := vld.Validation(productCategory); err != nil {
		return false, err
	}

	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Category)
	return available, pcc.productCategoryRepo.CreateProductCategory(productCategory)
}

func (pcc *productCategoryUsecase) UpdateProductCategory(productCategory *pe.ProductCategory, id int) (bool, error) {
	available, _ := pcc.productCategoryRepo.IsProductCategoryAvailable(productCategory, productCategory.Category)
	return available, pcc.productCategoryRepo.UpdateProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) DeleteProductCategory(productCategory *pe.ProductCategory, id int) error {
	return pcc.productCategoryRepo.DeleteProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) GetAllProductCategoryNoPagination() (*[]pe.ProductCategory, error) {
	productCategories, err := pcc.productCategoryRepo.GetAllProductCategoryNoPagination()
	return productCategories, err
}

func (pcc *productCategoryUsecase) GetProductCategoryById(id int) (*pe.ProductCategory, error) {
	productCategories, err := pcc.productCategoryRepo.GetProductCategoryById(id)
	return productCategories, err
}

func (pcc *productCategoryUsecase) GetAllProductCategory(offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	productCategories, count, err := pcc.productCategoryRepo.GetAllProductCategory(offset, pageSize)
	return productCategories, count, err
}

func (pcc *productCategoryUsecase) SearchingProductCategoryByName(name string, offset, pageSize int) (*[]pe.ProductCategory, int64, error) {
	productCategories, count, err := pcc.productCategoryRepo.SearchingProductCategoryByName(name, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return productCategories, count, nil
}
