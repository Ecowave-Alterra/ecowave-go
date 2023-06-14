package transaction

import (
	"errors"
	"log"

	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"github.com/labstack/echo/v4"
)

func (tr *transactionRepo) CreateTransaction(transaction *et.Transaction) error {

	err := tr.db.Create(&transaction).Error
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	//update stock
	for _, val := range transaction.TransactionDetails {
		var product ep.Product
		err := tr.db.Select("stock").Where("product_id = ?", val.ProductId).First(&product).Error
		if err != nil {
			return echo.NewHTTPError(404, err)
		}

		stock := product.Stock - val.Qty
		err = tr.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Update("stock", stock).Error
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
	}
	return nil
}
func (tr *transactionRepo) UpdateTransaction(updateData et.Transaction) error {

	result := tr.db.Model(&et.Transaction{}).Where("transaction_id = ?", updateData.TransactionId).Updates(&updateData)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("err")
	}
	return nil
}

func (tr *transactionRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail

	if err := tr.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, echo.NewHTTPError(404, err)
	}
	point := userDetail.Point

	return point, nil

}

func (tr *transactionRepo) GetVoucherUser(id uint, offset int, pageSize int) ([]ev.Voucher, int64, error) {
	type IdVoucher struct {
		Id int
	}
	var idVoucher []IdVoucher
	// var result []ev.VoucherUserResponse
	var voucher []ev.Voucher
	var count int64

	subquery := tr.db.Model(&ev.Voucher{}).Select(`vouchers.id, voucher_types.type, vouchers.end_date, voucher_types.photo_url,
	vouchers.minimum_purchase,vouchers.max_claim_limit, count(*) user_claim`).Joins("left join transactions on transactions.voucher_id = vouchers.id").Joins("left join voucher_types on voucher_types.id = transactions.voucher_id").Where("transactions.user_id = ?", id).Group("vouchers.id")

	err := tr.db.Select(`id`).Table("(?) as sub", subquery).Where("user_claim > max_claim_limit").Scan(&idVoucher).Error
	if err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	err = tr.db.Not(idVoucher).Find(&voucher).Count(&count).Error

	// subquery2 := tr.db.Select(`id`).Table("(?) as sub", subquery).Where("user_claim > max_claim_limit")

	log.Println(idVoucher)

	// db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&voucher)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// SELECT * FROM vouchers WHERE id NOT IN
	// (SELECT ProductName FROM Supplier WHERE SupplierName = "ABC");

	//
	// err := tr.db.Table("(?) as sub", subquery).Where("user_claim < max_claim_limit").Count(&count).Error
	if err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	err = tr.db.Preload("VoucherType").Offset(offset).Limit(pageSize).Not(idVoucher).Find(&voucher).Error

	if err != nil {
		return voucher, 0, echo.NewHTTPError(404, err)
	}

	return voucher, count, nil

}
func (tr *transactionRepo) CountVoucherUser(idUser uint, idVoucher uint) (uint, error) {
	var count int64

	err := tr.db.Model(&et.Transaction{}).Where("user_id = ? AND voucher_id = ?", idUser, idVoucher).Count(&count).Error
	if err != nil {
		return 0, echo.NewHTTPError(404, err)
	}

	return uint(count), nil

}
func (tr *transactionRepo) ClaimVoucher(id uint) (ev.Voucher, error) {
	var voucher ev.Voucher

	if err := tr.db.Where("id = ?", id).First(&voucher).Error; err != nil {
		return voucher, echo.NewHTTPError(404, err)
	}

	return voucher, nil

}
func (tr *transactionRepo) DetailVoucher(id uint) (ev.Voucher, error) {
	var voucher ev.Voucher

	if err := tr.db.Preload("VoucherType").Where("id = ?", id).First(&voucher).Error; err != nil {
		return voucher, echo.NewHTTPError(404, err)
	}

	return voucher, nil
}
