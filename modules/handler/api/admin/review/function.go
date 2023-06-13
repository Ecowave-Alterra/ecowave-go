package review

import (
	"fmt"
	"net/http"

	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) GetAllReview(c echo.Context) error {
	var products []pe.Product
	products, err := rh.reviewUsecase.GetAllProducts(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data produk",
		})
	}

	count := 0
	var transactionDetails []te.TransactionDetail
	// var review re.Review
	var reviewResponses []re.GetAllReviewResponse
	for _, product := range products {
		transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(product.ID), &transactionDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data transaksi detail produk",
			})
		}
		for _, td := range transactionDetails {
			if fmt.Sprint(td.RatingProductId) != "" {
				count++
			}
		}
		reviewResponse := re.GetAllReviewResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.ProductCategory.Category,
			ReviewQty: uint(count),
		}

		reviewResponses = append(reviewResponses, reviewResponse)
		count = 0
	}

	if reviewResponses == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Belum ada list ulasan",
			"Status":  http.StatusOK,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil mengambil data review produk",
		"Reviews": reviewResponses,
	})
}

func (rh *ReviewHandler) GetReviewByProductID(c echo.Context) error {
	productID := c.Param("id")

	var transactionDetails []te.TransactionDetail
	transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(productID), &transactionDetails)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data transaksi detail produk",
		})
	}

	var review re.Review
	var reviewResponses []re.ReviewResponse
	for _, td := range transactionDetails {
		if fmt.Sprint(td.RatingProductId) != "" {
			review, err = rh.reviewUsecase.GetAllReviewByID(fmt.Sprint(td.RatingProductId), &review)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data ulasan produk",
					"Status":  http.StatusInternalServerError,
				})
			}

			reviewResponse := re.ReviewResponse{
				TransactionID: td.TransactionId,
				Rating:        review.Rating,
				Comment:       review.Comment,
				CommentAdmin:  review.CommentAdmin,
				PhotoUrl:      review.PhotoUrl,
				VideoUrl:      review.VideoUrl,
			}

			reviewResponses = append(reviewResponses, reviewResponse)
		}
	}

	if reviewResponses == nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data ulasan produk",
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil mengambil data review produk",
		"Reviews": reviewResponses,
	})
}
