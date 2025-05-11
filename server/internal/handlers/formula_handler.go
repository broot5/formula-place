package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/broot5/formula-place/server/internal/models"
	"github.com/broot5/formula-place/server/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid/v5"
)

type FormulaHandler struct {
	service  services.FormulaServiceInterface
	validate *validator.Validate
}

func NewFormulaHandler(service services.FormulaServiceInterface) *FormulaHandler {
	return &FormulaHandler{
		service:  service,
		validate: validator.New(),
	}
}

func httpJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
		}
	}
}

func httpErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Printf("HTTP Error %d: %s", statusCode, message)
	httpJSONResponse(w, statusCode, map[string]string{"error": message})
}

func formatValidationErrors(err error) map[string]string {
	errorsMap := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			fieldName := strings.ToLower(fieldErr.Field())
			switch fieldErr.Tag() {
			case "required":
				errorsMap[fieldName] = fieldName + " is required"
			case "min":
				errorsMap[fieldName] = fieldName + " must be at least " + fieldErr.Param() + " characters long"
			case "max":
				errorsMap[fieldName] = fieldName + " must be at most " + fieldErr.Param() + " characters long"
			case "uuid":
				errorsMap[fieldName] = fieldName + " must be a valid UUID"
			case "boolean":
				errorsMap[fieldName] = fieldName + " must be a boolean value"
			default:
				errorsMap[fieldName] = fieldName + " is invalid (" + fieldErr.Tag() + ")"
			}
		}
	}
	return errorsMap
}

func (h *FormulaHandler) CreateFormula(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqPayload models.CreateFormulaRequest
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		httpErrorResponse(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.validate.StructCtx(ctx, reqPayload); err != nil {
		validateErrors := formatValidationErrors(err)
		log.Printf("Validation failed for create formula: %v, errors: %v\n", err, validateErrors)
		httpJSONResponse(w, http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": validateErrors})
		return
	}

	createdFormula, err := h.service.CreateFormula(ctx, &reqPayload)
	if err != nil {
		log.Printf("Error creating formula: %v", err)
		httpErrorResponse(w, http.StatusInternalServerError, "Failed to create formula")
		return
	}

	httpJSONResponse(w, http.StatusCreated, createdFormula)
}

func (h *FormulaHandler) GetFormulaByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		httpErrorResponse(w, http.StatusBadRequest, "Missing formula ID")
		return
	}

	id, err := uuid.FromString(idParam)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, "Invalid formula ID format")
		return
	}

	formula, err := h.service.GetFormulaByID(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrFormulaNotFound) {
			httpErrorResponse(w, http.StatusNotFound, "Formula not found")
			return
		}
		log.Printf("Error getting formula by ID: %v", err)
		httpErrorResponse(w, http.StatusInternalServerError, "Failed to get formula")
		return
	}

	httpJSONResponse(w, http.StatusOK, formula)
}

func (h *FormulaHandler) GetAllFormulas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	title := r.URL.Query().Get("title")

	formulas, err := h.service.GetAllFormulas(ctx, title)
	if err != nil {
		log.Printf("Error getting formulas: %v", err)
		httpErrorResponse(w, http.StatusInternalServerError, "Failed to get formulas")
		return
	}

	httpJSONResponse(w, http.StatusOK, formulas)
}

func (h *FormulaHandler) UpdateFormula(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		httpErrorResponse(w, http.StatusBadRequest, "Missing formula ID")
		return
	}

	id, err := uuid.FromString(idParam)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, "Invalid formula ID format")
		return
	}

	var reqPayload models.UpdateFormulaRequest
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		httpErrorResponse(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.validate.StructCtx(ctx, reqPayload); err != nil {
		validateErrors := formatValidationErrors(err)
		log.Printf("Validation failed for update formula: %v, errors: %v\n", err, validateErrors)
		httpJSONResponse(w, http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": validateErrors})
		return
	}

	updatedFormula, err := h.service.UpdateFormula(ctx, id, &reqPayload)
	if err != nil {
		if errors.Is(err, services.ErrFormulaNotFound) {
			httpErrorResponse(w, http.StatusNotFound, "Formula not found")
			return
		}
		log.Printf("Error updating formula: %v", err)
		httpErrorResponse(w, http.StatusInternalServerError, "Failed to update formula")
		return
	}

	httpJSONResponse(w, http.StatusOK, updatedFormula)
}

func (h *FormulaHandler) DeleteFormula(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		httpErrorResponse(w, http.StatusBadRequest, "Missing formula ID")
		return
	}

	id, err := uuid.FromString(idParam)
	if err != nil {
		httpErrorResponse(w, http.StatusBadRequest, "Invalid formula ID format")
		return
	}

	err = h.service.DeleteFormula(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrFormulaNotFound) {
			httpErrorResponse(w, http.StatusNotFound, "Formula not found")
			return
		}
		log.Printf("Error deleting formula: %v", err)
		httpErrorResponse(w, http.StatusInternalServerError, "Failed to delete formula")
		return
	}

	httpJSONResponse(w, http.StatusNoContent, nil)
}
