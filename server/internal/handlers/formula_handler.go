package handlers

import (
	"context"

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
	Uuid uuid.UUID `path:"id" format:"uuid"`
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
		return nil, err
	}
	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) GetFormula(ctx context.Context, input *FormulaIDInput) (*FormulaResponseOutput, error) {
	modelResponse, err := h.service.GetFormula(ctx, input.Uuid)
	if err != nil {
		return nil, err
	}
	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) UpdateFormula(ctx context.Context, input *UpdateFormulaInput) (*FormulaResponseOutput, error) {
	updateRequest := &models.UpdateFormulaRequest{
		Title:       &input.Body.Title,
		Content:     &input.Body.Content,
		Description: &input.Body.Description,
	}
	modelResponse, err := h.service.UpdateFormula(ctx, input.Uuid, updateRequest)
	if err != nil {
		return nil, err
	}
	return &FormulaResponseOutput{Body: modelResponse}, nil
}

func (h *FormulaHandler) DeleteFormula(ctx context.Context, input *FormulaIDInput) (*struct{}, error) {
	err := h.service.DeleteFormula(ctx, input.Uuid)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *FormulaHandler) GetAllFormulas(ctx context.Context, input *FormulaSearchByTitleInput) (*FormulasReponseOutput, error) {
	formulas, err := h.service.GetAllFormulas(ctx, input.Title)
	if err != nil {
		return nil, err
	}
	return &FormulasReponseOutput{Body: formulas}, nil
}
