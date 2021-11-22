package entities

import "time"

type Tenant struct {
	tableName struct{}  `pg:"tenants"`
	Id        int64     `json:"id" pg:",pk"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty" pg:"default:now()"`
	UpdatedAt time.Time `json:"updated_at,omitempty" pg:"default:now()"`
}
