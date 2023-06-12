package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	informationHandlerAdmin "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/information"
	informationRepoAdmin "github.com/berrylradianh/ecowave-go/modules/repository/admin/information"
	informationUsecaseAdmin "github.com/berrylradianh/ecowave-go/modules/usecase/admin/information"

	productCategoryHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	productCategoryRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin/product_category"
	productCategoryUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product_category"

	informationHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/information"
	informationRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/information"
	informationUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/information"

	transactionHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/transaction"
	transactionRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
	transactionUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/transaction"

	orderHandlerUser "github.com/berrylradianh/ecowave-go/modules/handler/api/user/order"
	orderRepoUser "github.com/berrylradianh/ecowave-go/modules/repository/user/order"
	orderUsecaseUser "github.com/berrylradianh/ecowave-go/modules/usecase/user/order"

	authHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	authRepo "github.com/berrylradianh/ecowave-go/modules/repository/auth"
	authUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/auth"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	informationRepoAdmin := informationRepoAdmin.New(mysql.DB)
	informationUsecaseAdmin := informationUsecaseAdmin.New(informationRepoAdmin)
	informationHandlerAdmin := informationHandlerAdmin.New(informationUsecaseAdmin)

	productCategoryRepo := productCategoryRepo.New(mysql.DB)
	productCategoryUsecase := productCategoryUsecase.New(productCategoryRepo)
	productCategoryHandler := productCategoryHandler.New(productCategoryUsecase)

	informationRepoUser := informationRepoUser.New(mysql.DB)
	informationUsecaseUser := informationUsecaseUser.New(informationRepoUser)
	informationHandlerUser := informationHandlerUser.New(informationUsecaseUser)

	transactionRepoUser := transactionRepoUser.New(mysql.DB)
	transactionUsecaseUser := transactionUsecaseUser.New(transactionRepoUser)
	transactionHandlerUser := transactionHandlerUser.New(transactionUsecaseUser)

	orderRepoUser := orderRepoUser.New(mysql.DB)
	orderUsecaseUser := orderUsecaseUser.New(orderRepoUser)
	orderHandlerUser := orderHandlerUser.New(orderUsecaseUser)

	handler := common.Handler{
		AuthHandler:             authHandler,
		InformationHandlerAdmin: informationHandlerAdmin,
		InformationHandlerUser:  informationHandlerUser,
		TransactionHandlerUser:  transactionHandlerUser,
		OrderHandlerUser:        orderHandlerUser,
		ProductCategoryHandler:  productCategoryHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
