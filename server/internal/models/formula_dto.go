package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type CreateFormulaRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

type UpdateFormulaRequest struct {
	Title       *string `json:"title"`
	Content     *string `json:"content"`
	Description *string `json:"description"`
}

type FormulaResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
