package information

import (
	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ir *informationRepo) GetAllInformations(offset int, pageSize int) (*[]ie.UserInformationResponse, int64, error) {
	var informations []ie.Information
	var informationsRes []ie.UserInformationResponse
	var count int64

	if err := ir.db.Model(&ie.Information{}).Where("status = ?", "Terbit").Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := ir.db.Where("status = ?", "Terbit").Find(&informations).Offset(offset).Limit(pageSize).Find(&informations).Error; err != nil {
		return nil, 0, err
	}

	for _, val := range informations {
		result := ie.UserInformationResponse{
			InformationId:   val.InformationId,
			Title:           val.Title,
			PhotoContentUrl: val.PhotoContentUrl,
			Content:         val.Content,
			Date:            val.CreatedAt,
		}
		informationsRes = append(informationsRes, result)
	}

	return &informationsRes, count, nil
}

func (ir *informationRepo) UpdatePoint(id uint, point uint) error {

	err := ir.db.Model(&eu.UserDetail{}).Where("user_id = ?", id).Update("point", point).Error
	if err != nil {
		return err
	}

	return nil
}

func (ir *informationRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail
	if err := ir.db.Where("user_id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}
	point := userDetail.Point

	return point, nil
}
