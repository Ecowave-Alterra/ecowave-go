package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			Name:            "User 1",
			Point:           0,
			Phone:           "08917283129283",
			ProfilePhotoUrl: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:          2,
		},
		{
			Name:            "User 2",
			Point:           0,
			Phone:           "0851728392716",
			ProfilePhotoUrl: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:          3,
		},
	}

	return userDetail
}