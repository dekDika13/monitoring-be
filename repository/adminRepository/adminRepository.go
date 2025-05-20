package adminRepository

import (
	"backend/dto/adminDto"
	"backend/model"
	"errors"
	"time"
	"gorm.io/gorm"
)

type AdminRepository interface {
	// TODO AUTH
	RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error)
	LoginAdmin(payloads adminDto.LoginDTO) (model.Admin, error)

	// TODO ROLES
	CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error)

}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

// TODO ADMIN REPOSITORY HERE

// TODO LOGIN ADMIN
func (u *adminRepository) LoginAdmin(payloads adminDto.LoginDTO) (model.Admin, error) {
	var admin model.Admin

	query := u.db.Where("username = ?", payloads.Username).First(&admin)
	if query.Error != nil {
		return admin, query.Error
	}

	if query.RowsAffected < 1 {
		return admin, errors.New("username is incorrect")
	}

	return admin, nil
}

// TODO REGISTER ADMIN
func (u *adminRepository) RegisterAdmin(payloads adminDto.RegisterAdminDto) (adminDto.RegisterAdminDto, error) {

	if err := u.db.Create(&model.Admin{
		RoleId:             payloads.RoleId,
		Username:           payloads.Username,
		Password:           payloads.Password,
		CreatedAT:          time.Now(),
	}).Error; err != nil {
		return payloads, err
	}

	return payloads, nil
}
