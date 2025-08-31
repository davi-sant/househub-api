package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/davi-sant/househub-go/helpers/helpErrors"
	"github.com/davi-sant/househub-go/models"
	"github.com/davi-sant/househub-go/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RecordController struct {
	service   *services.RecordService
	validator *validator.Validate
}

func NewRecordController(service *services.RecordService) *RecordController {
	v := validator.New()
	return &RecordController{
		service:   service,
		validator: v,
	}
}

func validateRequest(input any, v *validator.Validate) []error {
	var errorMessages []error

	if err := v.Struct(input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, validationError := range validationErrors {
				switch validationError.Tag() {
				case "required":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s é obrigatório", validationError.Field()))
				case "max":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s excede o tamanho permitido de %s caracteres", validationError.Field(), validationError.Param()))
				case "min":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s deve ter pelo menos %s caracteres", validationError.Field(), validationError.Param()))
				case "email":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s deve ser um email válido", validationError.Field()))
				default:
					errorMessages = append(errorMessages, fmt.Errorf("erro no campo %s: %s", validationError.Field(), validationError.Tag()))
				}
			}
		} else {

			errorMessages = append(errorMessages, err)
		}
		return errorMessages
	}
	return nil
}

func (r *RecordController) Create(c *gin.Context) {
	var input models.RecordCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		respError := models.RecordResponseError{
			Status:   "error",
			Mensagem: fmt.Sprintf("Erro ao fazer parse do JSON: %v", err),
			Erros:    helpErrors.ErrorItems(nil, err),
		}
		c.JSON(http.StatusBadRequest, respError)
		return
	}

	validationErrors := validateRequest(input, r.validator)

	if validationErrors != nil {

		c.JSON(http.StatusBadRequest, models.RecordResponseError{
			Status:   "error",
			Mensagem: "Erro de validação",
			Erros:    helpErrors.ErrorItems(validationErrors, nil),
		})
		return
	}

	record, err := r.service.Create(c.Request.Context(), input)

	if err != nil {

		respError := models.RecordResponseError{
			Status:   "error",
			Mensagem: fmt.Sprintf("Erro ao gravar registro: %v", err),
			Erros:    helpErrors.ErrorItems(nil, err),
		}
		c.JSON(http.StatusInternalServerError, respError)
		return
	}

	c.JSON(http.StatusCreated, models.RecordResponse{
		Status:   "success",
		Mensagem: "Registro gravado com sucesso",
		Record:   record,
	})
}

func (r *RecordController) FindAll(c *gin.Context) {
	records, err := r.service.FindAll(c.Request.Context())
	if len(records) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	if err != nil {
		respError := models.RecordResponseError{
			Status:   "error",
			Mensagem: fmt.Sprintf("Erro ao fazer parse do JSON: %v", err),
			Erros:    helpErrors.ErrorItems(nil, err),
		}
		c.JSON(http.StatusInternalServerError, respError)
		return
	}

	c.JSON(http.StatusOK, records)
}

func (r *RecordController) FindById(c *gin.Context) {

	input := models.FindRecordById{ID: c.Params.ByName("id")}

	validationErrors := validateRequest(input, r.validator)

	if validationErrors != nil {
		c.JSON(http.StatusBadRequest, models.RecordResponseError{
			Status:   "error",
			Mensagem: "Erro de validação",
			Erros:    helpErrors.ErrorItems(validationErrors),
		})
		return
	}

	record, err := r.service.FindById(c.Request.Context(), input)
	if record == nil {
		respRecord := models.RecordResponse{
			Status:   "Sem dados",
			Mensagem: "registro não encontrado",
			Record:   &models.Record{},
		}
		c.JSON(http.StatusOK, respRecord)
		return
	}

	if err != nil {

		respErr := models.RecordResponseError{
			Status:   "erro",
			Mensagem: fmt.Sprintf("Erro ao buscar registro por ID: Erro -> %v", err),
			Erros:    helpErrors.ErrorItems(nil, err),
		}

		c.JSON(http.StatusBadRequest, respErr)
		return
	}
	c.JSON(http.StatusOK, models.RecordResponse{
		Status:   "success",
		Mensagem: "Registro encontrado com sucesso",
		Record:   record,
	})

}
