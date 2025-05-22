package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type CreateFormulaRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type UpdateFormulaRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Content     *string `json:"content"`
}

type FormulaResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
