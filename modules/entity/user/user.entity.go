package user

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model   `json:"-"`
	Email         string           `json:"Email" form:"Email" validate:"required,email"`
	Username      string           `json:"Username" form:"Username" validate:"required"`
	GoogleId      string           `json:"GoogleId" form:"GoogleId"`
	Password      string           `json:"Password" form:"Password" validate:"required"`
	RoleId        uint             `json:"RoleId" form:"RoleId"`
	UserDetail    UserDetail       `gorm:"foreignKey:UserId"`
	UserAddresses []UserAddress    `gorm:"foreignKey:UserId"`
	Transactions  []et.Transaction `gorm:"foreignKey:UserId"`
}

type RegisterRequest struct {
	Name     string `json:"Name" form:"Name" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Username string `json:"Username" form:"Username" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type AuthResponse struct {
	ID    int    `json:"Id" form:"Id"`
	Email string `json:"Email" form:"Email"`
	Token string `json:"Token" form:"Token"`
}
