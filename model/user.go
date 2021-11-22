package model

import "apm/pkg/entities"

type ListUserResult struct {
	Total   int             `json:"total"`
	Limit   int             `json:"limit"`
	Skip    int             `json:"skip"`
	Records []entities.User `json:"records"`
}

type UserCreate struct {
	Username     string   `json:"username" pg:",unique"  validate:"required"`
	Password     string   `json:"password,omitempty"  validate:"required"`
	FirstName    string   `json:"first_name"  validate:"required"`
	LastName     string   `json:"last_name"  validate:"required"`
	Email        string   `json:"email" pg:",unique"  validate:"required"`
	Address      string   `json:"address,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	Permissions  []string `json:"permissions"  validate:"required"`
	DepartmentId int64    `json:"department_id,omitempty"  validate:"required"`
}

type UserUpdate struct {
	Username     string   `json:"username" pg:",unique"  validate:"required"`
	FirstName    string   `json:"first_name"  validate:"required"`
	LastName     string   `json:"last_name"  validate:"required"`
	Email        string   `json:"email" pg:",unique"  validate:"required"`
	Address      string   `json:"address,omitempty"`
	Phone        string   `json:"phone,omitempty"`
	Permissions  []string `json:"permissions"  validate:"required"`
	DepartmentId int64    `json:"department_id,omitempty"  validate:"required"`
}
