package profile

import (
	p "github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (pc *profileUsecase) GetUserProfile(user *ut.User, id int) error {
	return pc.profileRepo.GetUserProfile(user, id)
}

func (pc *profileUsecase) GetUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	return pc.profileRepo.GetUserDetailProfile(userDetail, id)
}

func (pc *profileUsecase) UpdateUserProfile(user *ut.User, id int) error {
	return pc.profileRepo.UpdateUserProfile(user, id)
}

func (pc *profileUsecase) UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	return pc.profileRepo.UpdateUserDetailProfile(userDetail, id)
}

func (pc *profileUsecase) CreateAddressProfile(address *ut.UserAddress) error {
	return pc.profileRepo.CreateAddressProfile(address)
}

func (pc *profileUsecase) GetAllAddressProfile(address *[]ut.UserAddress, idUser int) error {
	return pc.profileRepo.GetAllAddressProfile(address, idUser)
}

func (pc *profileUsecase) GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	return pc.profileRepo.GetAddressByIdProfile(address, idUser, idAddress)
}

func (pc *profileUsecase) UpdateAddressProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	return pc.profileRepo.UpdateAddressProfile(address, idUser, idAddress)
}

func (pc *profileUsecase) UpdatePasswordProfile(user *ut.User, oldPassword string, newPassword string, id int) (string, error) {
	if err := pc.GetUserProfile(user, id); err != nil {
		return "", err
	}

	if err := p.VerifyPassword(user.Password, oldPassword); err != nil {
		return "password salah", err
	}

	hashNewPassword, err := p.HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	return "", pc.profileRepo.UpdatePasswordProfile(string(hashNewPassword), id)
}
