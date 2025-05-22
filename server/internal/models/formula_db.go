package models

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Formula struct {
	ID          uuid.UUID          `json:"id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	Content     string             `json:"content"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
