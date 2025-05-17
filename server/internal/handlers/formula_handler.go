package handlers

import (
	"context"
	"fmt"

	"github.com/broot5/formula-place/server/internal/models"
	"github.com/broot5/formula-place/server/internal/services"

	"github.com/gofrs/uuid/v5"
)

type CreateFormulaInput struct {
	Body struct {
		Title       string `json:"title" minLength:"1" maxLength:"255"`
		Content     string `json:"content"`
		Description string `json:"description,omitempty"`
	}
}

type UpdateFormulaInput struct {
	FormulaIDInput
	Body struct {
		Title       string `json:"title,omitempty" minLength:"1" maxLength:"255"`
		Content     string `json:"content,omitempty"`
		Description string `json:"description,omitempty"`
	}
}

type FormulaIDInput struct {
	UUID uuid.UUID `path:"id" format:"uuid"`
}

type FormulaSearchByTitleInput struct {
	Title string `query:"title"`
}

type FormulaResponseOutput struct {
	Body *models.FormulaResponse
}

type FormulasReponseOutput struct {
	Body []models.FormulaResponse
}

type FormulaHandler struct {
	service services.FormulaService
}

func NewFormulaHandler(service services.FormulaService) *FormulaHandler {
	return &FormulaHandler{
		service: service,
	}
}

func (h *FormulaHandler) CreateFormula(ctx context.Context, input *CreateFormulaInput) (*FormulaResponseOutput, error) {
	modelResponse, err := h.service.CreateFormula(ctx, (*models.CreateFormulaRequest)(&input.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to create formula in handler: %w", err)
	}

	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) GetFormula(ctx context.Context, input *FormulaIDInput) (*FormulaResponseOutput, error) {
	modelResponse, err := h.service.GetFormula(ctx, input.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get formula in handler: %w", err)
	}

	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) UpdateFormula(ctx context.Context, input *UpdateFormulaInput) (*FormulaResponseOutput, error) {
	updateRequest := &models.UpdateFormulaRequest{
		Title:       &input.Body.Title,
		Content:     &input.Body.Content,
		Description: &input.Body.Description,
	}
	modelResponse, err := h.service.UpdateFormula(ctx, input.UUID, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to update formula in handler: %w", err)
	}

	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) DeleteFormula(ctx context.Context, input *FormulaIDInput) (*struct{}, error) {
	err := h.service.DeleteFormula(ctx, input.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete formula in handler: %w", err)
	}

	return nil, nil
}

func (h *FormulaHandler) GetAllFormulas(ctx context.Context, input *FormulaSearchByTitleInput) (*FormulasReponseOutput, error) {
	formulas, err := h.service.GetAllFormulas(ctx, input.Title)
	if err != nil {
		return nil, fmt.Errorf("failed to get formulas in handler: %w", err)
	}

	return &FormulasReponseOutput{Body: formulas}, nil
}
