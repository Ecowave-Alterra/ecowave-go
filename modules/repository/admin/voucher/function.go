package voucher

import (
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (vr *voucherRepo) CreateVoucher(voucher *ve.Voucher) error {
	if err := vr.db.Create(voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) GetAllVoucher(offset, pageSize int) (*[]ve.Voucher, int64, error) {
	var vouchers []ve.Voucher
	var count int64
	if err := vr.db.Model(&ve.Voucher{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := vr.db.Offset(offset).Limit(pageSize).Preload("VoucherType").Find(&vouchers).Error; err != nil {
		return nil, 0, err
	}

	return &vouchers, count, nil
}

func (vr *voucherRepo) UpdateVoucher(voucherID string, voucher *ve.Voucher) error {
	if err := vr.db.Model(&ve.Voucher{}).Where("id = ?", voucherID).Updates(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) DeleteVoucher(voucherID string, voucher *ve.Voucher) error {
	if err := vr.db.Where("id = ?", voucherID).Delete(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (vr *voucherRepo) FilterVouchersByType(voucherType string, vouchers *[]ve.Voucher) ([]ve.Voucher, error) {
	if err := vr.db.Preload("VoucherType").Where("voucher_type_id IN (SELECT id FROM voucher_types WHERE type = ?)", voucherType).Find(&vouchers).Error; err != nil {
		return nil, err
	}

	return *vouchers, nil
}
