package entities

import "time"

type Department struct {
	tableName struct{}   `pg:"departments"`
	Id        int64      `json:"id" pg:",pk"`
	Name      string     `json:"name"`
	TenantId  int64      `json:"tenant_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Tenant *Tenant `json:"tenant,omitempty" pg:"rel:has-one,fk:tenant_id"`
}
