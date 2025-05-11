package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/broot5/formula-place/server/internal/models"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FormulaRepositoryInterface interface {
	CreateFormula(ctx context.Context, formula *models.Formula) error
	GetFormulaByID(ctx context.Context, id uuid.UUID) (*models.Formula, error)
	GetAllFormulas(ctx context.Context) ([]models.Formula, error)
	UpdateFormula(ctx context.Context, formula *models.Formula) error
	DeleteFormula(ctx context.Context, id uuid.UUID) error
	SearchFormulasByTitle(ctx context.Context, title string) ([]models.Formula, error)
}

type formulaRepository struct {
	pool *pgxpool.Pool
}

func NewFormulaRepository(pool *pgxpool.Pool) FormulaRepositoryInterface {
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
		log.Printf("Error creating formula in repository: %v\n", err)
		return fmt.Errorf("could not create formula in DB: %w", err)
	}
	return nil
}

func (r *formulaRepository) GetFormulaByID(ctx context.Context, id uuid.UUID) (*models.Formula, error) {
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
			return nil, nil
		}
		log.Printf("Error scanning formula by ID in repository: %v\n", err)
		return nil, fmt.Errorf("could not get formula by ID: %w", err)
	}
	return &formula, nil
}

func (r *formulaRepository) GetAllFormulas(ctx context.Context) ([]models.Formula, error) {
	query := `
		SELECT id, title, content, description, created_at, updated_at
		FROM formulas
		ORDER BY updated_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		log.Printf("Error getting all formulas in repository: %v\n", err)
		return nil, fmt.Errorf("could not get all formulas: %w", err)
	}
	defer rows.Close()

	return scanFormulas(rows)
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
		log.Printf("Error updating formula in repository: %v\n", err)
		return fmt.Errorf("could not update formula: %w", err)
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
		log.Printf("Error deleting formula in repository: %v\n", err)
		return fmt.Errorf("could not delete formula: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
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
		log.Printf("Error searching formulas by title in repository: %v\n", err)
		return nil, fmt.Errorf("could not search formulas by title: %w", err)
	}
	defer rows.Close()

	return scanFormulas(rows)
}

func scanFormulas(rows pgx.Rows) ([]models.Formula, error) {
	var formulas []models.Formula
	for rows.Next() {
		var f models.Formula
		if err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Content,
			&f.Description,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("could not scan formula row: %w", err)
		}
		formulas = append(formulas, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating formula rows: %w", err)
	}
	return formulas, nil
}
