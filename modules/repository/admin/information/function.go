package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
)

func (ir *informationRepo) GetAllInformationsNoPagination() (*[]ie.Information, error) {
	var informations []ie.Information
	if err := ir.db.Find(&informations).Error; err != nil {
		return nil, err
	}

	return &informations, nil
}

func (ir *informationRepo) GetAllInformations(offset, pageSize int) (*[]ie.Information, int64, error) {
	var informations []ie.Information
	var count int64
	if err := ir.db.Model(&ie.Information{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := ir.db.Offset(offset).Limit(pageSize).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}

func (ir *informationRepo) GetInformationById(informationId int) (*ie.Information, error) {
	var information ie.Information
	if err := ir.db.Where("information_id = ?", informationId).First(&information).Error; err != nil {
		return nil, err
	}

	return &information, nil
}

func (ir *informationRepo) CreateInformation(information *ie.Information) error {
	if err := ir.db.Create(&information).Error; err != nil {
		return err
	}

	return nil
}

func (ir *informationRepo) CheckInformationExists(informationId uint) (bool, error) {
	var count int64
	result := ir.db.Model(&ie.Information{}).Where("information_id = ?", informationId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	exists := count > 0
	return exists, nil
}

func (ir *informationRepo) UpdateInformation(informationId int, information *ie.Information) error {
	result := ir.db.Model(&information).Where("information_id = ?", informationId).Updates(&information)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ir *informationRepo) DeleteInformation(informationId int) error {
	var information *ie.Information
	if err := ir.db.Delete(&information, "information_id = ?", informationId).Error; err != nil {
		return err
	}

	return nil
}

func (ir *informationRepo) SearchInformations(search, filter string, offset, pageSize int) (*[]ie.Information, int64, error) {
	var informations []ie.Information
	var count int64

	countQuery := "SELECT COUNT(*) FROM information WHERE (title LIKE ? OR information_id LIKE ?) AND status LIKE ?"
	if err := ir.db.Raw(countQuery, "%"+search+"%", "%"+search+"%", "%"+filter+"%").Count(&count).Error; err != nil {
		return nil, 0, err
	}

	selectQuery := "SELECT * FROM information WHERE (title LIKE ? OR information_id LIKE ?) AND status LIKE ? LIMIT ? OFFSET ?"
	if err := ir.db.Raw(selectQuery, "%"+search+"%", "%"+search+"%", "%"+filter+"%", pageSize, offset).Scan(&informations).Error; err != nil {
		return nil, 0, err
	}

	return &informations, count, nil
}