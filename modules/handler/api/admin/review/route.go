package review

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (reviewHandler *ReviewHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	reviewGroup := e.Group("/admin/reviews")
	reviewGroup.Use(jwtMiddleware)
	reviewGroup.GET("", reviewHandler.GetAllProducts)
	reviewGroup.GET("/:id", reviewHandler.GetReviewByProductID)
	reviewGroup.GET("/search", reviewHandler.SearchReview)
}
