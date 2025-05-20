package adminDto

type RoleDTO struct {
	Role_Name string `json:"role_name" validate:"required"`
	
}

type RoleResponse struct {
	Role_Name      string `json:"role_name"`
	CreatedAT string `json:"created_at"`
}
