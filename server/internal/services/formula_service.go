package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/broot5/formula-place/server/internal/models"
	"github.com/broot5/formula-place/server/internal/repositories"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrFormulaNotFound = errors.New("formula not found")
)

type FormulaServiceInterface interface {
	CreateFormula(ctx context.Context, req *models.CreateFormulaRequest) (*models.FormulaResponse, error)
	GetFormulaByID(ctx context.Context, id uuid.UUID) (*models.FormulaResponse, error)
	GetAllFormulas(ctx context.Context, title string) ([]models.FormulaResponse, error)
	UpdateFormula(ctx context.Context, id uuid.UUID, req *models.UpdateFormulaRequest) (*models.FormulaResponse, error)
	DeleteFormula(ctx context.Context, id uuid.UUID) error
}

type formulaService struct {
	repo repositories.FormulaRepositoryInterface
}

func NewFormulaService(repo repositories.FormulaRepositoryInterface) FormulaServiceInterface {
	return &formulaService{repo: repo}
}

func (s *formulaService) CreateFormula(ctx context.Context, req *models.CreateFormulaRequest) (*models.FormulaResponse, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	now := time.Now().UTC()

	formula := &models.Formula{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}

	if err := s.repo.CreateFormula(ctx, formula); err != nil {
		return nil, fmt.Errorf("failed to create formula: %w", err)
	}

	return dbModelToResponse(formula), nil
}

func (s *formulaService) GetFormulaByID(ctx context.Context, id uuid.UUID) (*models.FormulaResponse, error) {
	formula, err := s.repo.GetFormulaByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get formula by ID: %w", err)
	}

	if formula == nil {
		return nil, ErrFormulaNotFound
	}

	return dbModelToResponse(formula), nil
}

func (s *formulaService) GetAllFormulas(ctx context.Context, title string) ([]models.FormulaResponse, error) {
	var formulas []models.Formula
	var err error

	if title != "" {
		formulas, err = s.repo.SearchFormulasByTitle(ctx, title)
	} else {
		formulas, err = s.repo.GetAllFormulas(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get formulas: %w", err)
	}

	responses := make([]models.FormulaResponse, len(formulas))
	for i, formula := range formulas {
		responses[i] = *dbModelToResponse(&formula)
	}

	return responses, nil
}

func (s *formulaService) UpdateFormula(ctx context.Context, id uuid.UUID, req *models.UpdateFormulaRequest) (*models.FormulaResponse, error) {
	formula, err := s.repo.GetFormulaByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get formula for update: %w", err)
	}

	if formula == nil {
		return nil, ErrFormulaNotFound
	}

	if req.Title != nil {
		formula.Title = *req.Title
	}
	if req.Content != nil {
		formula.Content = *req.Content
	}
	if req.Description != nil {
		formula.Description = pgtype.Text{
			String: *req.Description,
			Valid:  true,
		}
	}

	formula.UpdatedAt = pgtype.Timestamptz{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	if err := s.repo.UpdateFormula(ctx, formula); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrFormulaNotFound
		}
		return nil, fmt.Errorf("failed to update formula: %w", err)
	}

	return dbModelToResponse(formula), nil
}

func (s *formulaService) DeleteFormula(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteFormula(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrFormulaNotFound
		}
		return fmt.Errorf("failed to delete formula: %w", err)
	}

	return nil
}

func dbModelToResponse(formula *models.Formula) *models.FormulaResponse {
	description := ""
	if formula.Description.Valid {
		description = formula.Description.String
	}

	createdAt := time.Time{}
	if formula.CreatedAt.Valid {
		createdAt = formula.CreatedAt.Time
	}

	updatedAt := time.Time{}
	if formula.UpdatedAt.Valid {
		updatedAt = formula.UpdatedAt.Time
	}

	return &models.FormulaResponse{
		ID:          formula.ID,
		Title:       formula.Title,
		Content:     formula.Content,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
