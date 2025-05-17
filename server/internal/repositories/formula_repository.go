package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/broot5/formula-place/server/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FormulaRepository interface {
	CreateFormula(ctx context.Context, formula *models.Formula) error
	GetFormula(ctx context.Context, id uuid.UUID) (*models.Formula, error)
	UpdateFormula(ctx context.Context, formula *models.Formula) error
	DeleteFormula(ctx context.Context, id uuid.UUID) error
	GetAllFormulas(ctx context.Context) ([]models.Formula, error)
	SearchFormulasByTitle(ctx context.Context, title string) ([]models.Formula, error)
}

type formulaRepository struct {
	pool *pgxpool.Pool
}

func NewFormulaRepository(pool *pgxpool.Pool) FormulaRepository {
	return &formulaRepository{pool: pool}
}

func (r *formulaRepository) CreateFormula(ctx context.Context, formula *models.Formula) error {
	query := `
		INSERT INTO formulas (id, title, content, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.pool.Exec(ctx, query,
		formula.ID,
		formula.Title,
		formula.Content,
		formula.Description,
		formula.CreatedAt,
		formula.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert formula in repository: %w", err)
	}

	return nil
}

func (r *formulaRepository) GetFormula(ctx context.Context, id uuid.UUID) (*models.Formula, error) {
	query := `
		SELECT id, title, content, description, created_at, updated_at
		FROM formulas
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var formula models.Formula
	err := row.Scan(
		&formula.ID,
		&formula.Title,
		&formula.Content,
		&formula.Description,
		&formula.CreatedAt,
		&formula.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}

		return nil, fmt.Errorf("failed to query formula in repository: %w", err)
	}

	return &formula, nil
}

func (r *formulaRepository) UpdateFormula(ctx context.Context, formula *models.Formula) error {
	query := `
		UPDATE formulas
		SET title = $1, content = $2, description = $3, updated_at = $4
		WHERE id = $5
	`

	cmdTag, err := r.pool.Exec(ctx, query,
		formula.Title,
		formula.Content,
		formula.Description,
		formula.UpdatedAt,
		formula.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update formula in repository: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *formulaRepository) DeleteFormula(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM formulas WHERE id = $1"

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete formula in repository: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *formulaRepository) GetAllFormulas(ctx context.Context) ([]models.Formula, error) {
	query := `
		SELECT id, title, content, description, created_at, updated_at
		FROM formulas
		ORDER BY updated_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all formulas in repository: %w", err)
	}
	defer rows.Close()

	return scanFormulas(rows)
}

func (r *formulaRepository) SearchFormulasByTitle(ctx context.Context, title string) ([]models.Formula, error) {
	searchTerm := "%" + title + "%"

	query := `
        SELECT id, title, content, description, created_at, updated_at
        FROM formulas
        WHERE title ILIKE $1
        ORDER BY updated_at DESC
    `

	rows, err := r.pool.Query(ctx, query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search formulas by title in repository: %w", err)
	}
	defer rows.Close()

	return scanFormulas(rows)
}

func scanFormulas(rows pgx.Rows) ([]models.Formula, error) {
	formulas, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Formula, error) {
		var f models.Formula
		err := row.Scan(
			&f.ID,
			&f.Title,
			&f.Content,
			&f.Description,
			&f.CreatedAt,
			&f.UpdatedAt,
		)
		if err != nil {
			return models.Formula{}, fmt.Errorf("failed to scan formula data in repository: %w", err)
		}

		return f, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to collect formula rows in repository: %w", err)
	}

	return formulas, nil
}
