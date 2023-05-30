package product

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/berrylradianh/ecowave-go/helper/cloudstorage"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	productCategoryIDstr := c.FormValue("Product_category_id")
	productCategoryID, err := strconv.ParseUint(productCategoryIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid product category ID",
			"error":   err,
		})
	}
	name := c.FormValue("Name")
	stockStr := c.FormValue("Stock")
	stock, err := strconv.ParseUint(stockStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid stock",
			"error":   err,
		})
	}
	priceStr := c.FormValue("Price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid price",
			"error":   err,
		})
	}
	description := c.FormValue("Description")

	status := "tersedia"
	if stock == 0 {
		status = "habis"
	}

	product := ep.Product{
		Product_category_id: uint(productCategoryID),
		Name:                name,
		Stock:               uint(stock),
		Price:               price,
		Status:              status,
		Rating:              0.00,
		Description:         description,
	}
	err = h.productUseCase.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create product",
			"error":   err,
		})
	}

	for i := 1; i <= 10; i++ {
		fileHeader, err := c.FormFile(fmt.Sprintf("PhotoContentUrl%d", i))
		if fileHeader != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, Anda harus mengunggah foto",
				})
			}
			fileExtension := filepath.Ext(fileHeader.Filename)
			allowedExtensions := map[string]bool{
				".png":  true,
				".jpeg": true,
				".jpg":  true,
			}
			if !allowedExtensions[fileExtension] {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, format file yang Anda unggah tidak sesuai",
				})
			}
			maxFileSize := 4 * 1024 * 1024
			fileSize := fileHeader.Size
			if fileSize > int64(maxFileSize) {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"Message": "Mohon maaf, ukuran file Anda melebihi batas maksimum 4MB",
				})
			}

			PhotoUrl, _ := cloudstorage.UploadToBucket(c.Request().Context(), fileHeader)

			log.Println(PhotoUrl)

			productImage := ep.ProductImage{
				Product_id:        product.ID,
				Product_image_url: PhotoUrl,
			}
			err = h.productUseCase.CreateProductImage(&productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed to create product image",
					"error":   err,
				})
			}

		} else {
			if err != nil {
				i = 1000
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully create new product",
	})
}

func (h *ProductHandler) GetAllProduct(c echo.Context) error {
	var products []ep.Product

	products, err := h.productUseCase.GetAllProduct(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get product datas",
			"error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Description,
			Product_image_url: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get all product datas",
		"products": productResponses,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	var product ep.Product
	productID := c.Param("id")

	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	var productImage ep.ProductImage
	productImages, err := h.productUseCase.GetProductImageURLById(productID, &productImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product images",
			"error":   err,
		})
	}

	var imageURLs []string
	for _, image := range productImages {
		imageURLs = append(imageURLs, image.Product_image_url)
	}

	productResponse := ep.ProductResponse{
		Product_id:        product.ID,
		Name:              product.Name,
		Category:          product.Product_Category.Category,
		Stock:             product.Stock,
		Price:             product.Price,
		Status:            product.Status,
		Rating:            product.Rating,
		Description:       product.Description,
		Product_image_url: imageURLs,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get product data by id",
		"products": productResponse,
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	productID := c.Param("id")
	var req ep.ProductRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to bind data",
			"error":   err,
		})
	}

	status := "tersedia"
	if req.Stock == 0 {
		status = "habis"
	}

	req = ep.ProductRequest{
		Product_category_id: req.Product_category_id,
		Name:                req.Name,
		Stock:               req.Stock,
		Price:               req.Price,
		Description:         req.Description,
		Status:              status,
		Product_image_url:   req.Product_image_url,
	}

	var product ep.Product
	product, err = h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	err = h.productUseCase.UpdateProduct(productID, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update product",
	})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	productID := c.Param("id")

	var product ep.Product
	product, err := h.productUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get product",
			"error":   err,
		})
	}

	var productImages []ep.ProductImage
	err = h.productUseCase.DeleteProductImage(productID, &productImages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete product image",
			"error":   err,
		})
	}

	err = h.productUseCase.DeleteProduct(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete product",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully delete product",
	})
}

func (h *ProductHandler) SearchProduct(c echo.Context) error {
	param := c.QueryParam("param")

	switch param {
	case "id":
		var product ep.Product
		productID := c.QueryParam("id")
		product, err := h.productUseCase.SearchProductByID(productID, &product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product",
				"error":   err,
			})
		}

		var productImage ep.ProductImage
		productImages, err := h.productUseCase.GetProductImageURLById(productID, &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Description,
			Product_image_url: imageURLs,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product data by id",
			"products": productResponse,
		})
	case "name":
		var products []ep.Product
		name := c.QueryParam("name")
		products, err := h.productUseCase.SearchProductByName(name, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product",
				"error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Failed to get product images",
					"error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.Product_image_url)
			}

			productResponse := ep.ProductResponse{
				Product_id:        product.ID,
				Name:              product.Name,
				Category:          product.Product_Category.Category,
				Stock:             product.Stock,
				Price:             product.Price,
				Status:            product.Status,
				Rating:            product.Rating,
				Description:       product.Description,
				Product_image_url: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Product not found",
				"product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product by name",
			"products": productResponses,
		})
	case "category":
		var products []ep.Product
		category := c.QueryParam("category")
		products, err := h.productUseCase.SearchProductByCategory(category, &products)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product by category",
				"error":   err,
			})
		}

		var productResponses []ep.ProductResponse
		var productImage ep.ProductImage
		for _, product := range products {
			productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Failed to get product images",
					"error":   err,
				})
			}

			var imageURLs []string
			for _, image := range productImages {
				imageURLs = append(imageURLs, image.Product_image_url)
			}

			productResponse := ep.ProductResponse{
				Product_id:        product.ID,
				Name:              product.Name,
				Category:          product.Product_Category.Category,
				Stock:             product.Stock,
				Price:             product.Price,
				Status:            product.Status,
				Rating:            product.Rating,
				Description:       product.Description,
				Product_image_url: imageURLs,
			}

			productResponses = append(productResponses, productResponse)
		}

		if len(productResponses) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Product not found",
				"product": productResponses,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Successfully get product data by category",
			"products": productResponses,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "Invalid search parameter",
	})
}

func (h *ProductHandler) FilterProductByStatus(c echo.Context) error {
	status := c.QueryParam("status")

	var product []ep.Product
	products, err := h.productUseCase.FilterProductByStatus(status, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to filter product",
			"error":   err,
		})
	}

	var productResponses []ep.ProductResponse
	var productImage ep.ProductImage
	for _, product := range products {
		productImages, err := h.productUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to get product images",
				"error":   err,
			})
		}

		var imageURLs []string
		for _, image := range productImages {
			imageURLs = append(imageURLs, image.Product_image_url)
		}

		productResponse := ep.ProductResponse{
			Product_id:        product.ID,
			Name:              product.Name,
			Category:          product.Product_Category.Category,
			Stock:             product.Stock,
			Price:             product.Price,
			Status:            product.Status,
			Rating:            product.Rating,
			Description:       product.Description,
			Product_image_url: imageURLs,
		}

		productResponses = append(productResponses, productResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Successfully get product based on stock status",
		"products": productResponses,
	})
}
