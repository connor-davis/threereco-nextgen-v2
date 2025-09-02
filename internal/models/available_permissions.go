package models

type AvailablePermissionsGroup struct {
	Name        string                `json:"name"`
	Permissions []AvailablePermission `json:"permissions"`
}

type AvailablePermission struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}
