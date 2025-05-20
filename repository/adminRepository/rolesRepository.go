package adminRepository

import (
	"backend/dto/adminDto"
	"backend/model"
	"time"
)

func (u *adminRepository) CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error) {

	temp := adminDto.RoleDTO{
		Role_Name: payloads.Role_Name,
	}

	if err := u.db.Create(&model.Role{
		Role_Name:      temp.Role_Name,
		CreatedAT: time.Now(),
	}).Error; err != nil {
		return temp, err
	}

	return temp, nil
}
