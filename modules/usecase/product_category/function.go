package productcategory

import pct "github.com/berrylradianh/ecowave-go/modules/entity/product"

func (pcc *productCategoryUsecase) CreateProductCategory(productCategory *pct.ProductCategory) error {
	return pcc.productCategoryRepo.CreateProductCategory(productCategory)
}

func (pcc *productCategoryUsecase) UpdateProductCategory(productCategory *pct.ProductCategory, id int) error {
	return pcc.productCategoryRepo.UpdateProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) DeleteProductCategory(productCategory *pct.ProductCategory, id int) error {
	return pcc.productCategoryRepo.DeleteProductCategory(productCategory, id)
}

func (pcc *productCategoryUsecase) GetAllProductCategory(productCategory *[]pct.ProductCategory) error {
	return pcc.productCategoryRepo.GetAllProductCategory(productCategory)
}

func (pcc *productCategoryUsecase) SearchingProductCategoyByName(productCategory *pct.ProductCategory, name string) error {
	return pcc.productCategoryRepo.SearchingProductCategoyByName(productCategory, name)
}
