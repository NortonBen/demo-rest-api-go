package entities

import "time"

type User struct {
	tableName    struct{}    `pg:"users"`
	Id           int64       `json:"id" pg:",pk"`
	Username     string      `json:"username,omitempty" pg:",unique"  validate:"required"`
	Password     string      `json:"password,omitempty"  validate:"required"`
	FirstName    string      `json:"first_name,omitempty"  validate:"required"`
	LastName     string      `json:"last_name,omitempty"  validate:"required"`
	Email        string      `json:"email,omitempty" pg:",unique"  validate:"required"`
	Address      string      `json:"address,omitempty"`
	Phone        string      `json:"phone,omitempty"`
	TenantId     int64       `json:"tenant_id,omitempty"  validate:"required"`
	DepartmentId int64       `json:"department_id,omitempty"  validate:"required"`
	Permissions  []string    `json:"permissions,omitempty" pg:",array"`
	Image        string      `json:"image,omitempty"`
	CreatedAt    *time.Time  `json:"created_at,omitempty" pg:"default:now()"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty" pg:"default:now()"`
	Tenant       *Tenant     `json:"tenant,omitempty" pg:"rel:has-one,fk:tenant_id"`
	Department   *Department `json:"department,omitempty" pg:"rel:has-one,fk:department_id"`
}
