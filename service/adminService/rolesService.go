package adminService

import (
	"backend/dto/adminDto"
	"errors"
)

func (s *adminService) CreateRoles(payloads adminDto.RoleDTO) (adminDto.RoleDTO, error) {

	temp := adminDto.RoleDTO{
		Role_Name: payloads.Role_Name,
	}

	if temp.Role_Name == "" {
		return temp, errors.New("name cannot be empty")
	}

	res, err := s.adminRepository.CreateRoles(payloads)

	if err != nil {
		return res, err
	}

	return temp, nil
}
