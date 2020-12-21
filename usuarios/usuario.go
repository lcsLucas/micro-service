package usuarios

import "context"

type Usuario struct {
	ID    int    `json:"id, omitempty"`
	Nome  string `json:"nome"`
	Email string `json:"email`
	Err   string `json:"err,omitempty"`
}

type Repository interface {
	Get(ctx context.Context, id int) (Usuario, error)
}
