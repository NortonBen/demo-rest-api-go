package entities

import "time"

type File struct {
	tableName struct{}   `pg:"files"`
	Id        string     `json:"id" pg:"type:uuid,pk"`
	Name      string     `json:"name" pg:",notnull"`
	FileExt   string     `json:"file_ext"`
	Type      string     `json:"type" pg:",notnull"`
	Note      string     `json:"note"`
	Size      int64      `json:"size"`
	FilePath  string     `json:"file_path" pg:",notnull"  swaggerignore:"true"`
	UploadBy  int64      `json:"upload_by"  swaggerignore:"true"`
	TenantId  int64      `json:"tenant_id"  swaggerignore:"true"`
	Keep      bool       `json:"keep"  swaggerignore:"true"`
	Time      *time.Time `json:"time,omitempty"`

	Tenant *Tenant `json:"tenant,omitempty" pg:"rel:has-one,fk:tenant_id" swaggerignore:"true"`
	User   *User   `json:"tenant,omitempty" pg:"rel:has-one,fk:upload_by" swaggerignore:"true"`
}
