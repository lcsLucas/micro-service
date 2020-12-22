package usuarios

import "context"

type Usuario struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Nome  string `gorm:"size:255; not null" json:"nome"`
	Email string `gorm:"size:255; not null" json:"email`
	Err   string `json:"err,omitempty"`
}

type Repository interface {
	Get(ctx context.Context, id int) (Usuario, error)
}
