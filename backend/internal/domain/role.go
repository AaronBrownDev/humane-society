package domain

// Role represents a security role in the system
type Role struct {
	RoleID      int    `json:"roleId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleRepository interface {
}
