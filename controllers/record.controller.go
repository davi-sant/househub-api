package controllers

import (
	"fmt"
	"net/http"

	"github.com/davi-sant/househub-go/helpers/helpErrors"
	"github.com/davi-sant/househub-go/helpers/validations"
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

	validationErrors := validations.ValidateRequest(input, r.validator)

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
	input := models.RecordById{ID: c.Params.ByName("id")}
	validationErrors := validations.ValidateRequest(input, r.validator)

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

func (rc *RecordController) Update(gc *gin.Context) {
	var inputBody models.RecordUpdate
	inputParams := models.RecordById{ID: gc.Param("id")}

	if err := gc.ShouldBindJSON(&inputBody); err != nil {
		respError := models.RecordResponseError{
			Status:   "erro",
			Mensagem: "erro ao fazer parse do JSON",
			Erros:    []models.ErrorItem{},
		}
		gc.JSON(http.StatusBadRequest, respError)
		return
	}
	vParamsBody := validations.ValidateRequest(inputParams, rc.validator)
	if vParamsBody != nil {
		respError := models.RecordResponseError{
			Status:   "erro",
			Mensagem: "erro ao validar paramatros da URL.",
			Erros:    helpErrors.ErrorItems(vParamsBody),
		}
		gc.JSON(http.StatusBadRequest, respError)
		return
	}
	vParams := validations.ValidateRequest(inputBody, rc.validator)
	if vParams != nil {
		respError := models.RecordResponseError{
			Status:   "erro",
			Mensagem: "erro de validação na atulizaçao de registros",
			Erros:    helpErrors.ErrorItems(vParams),
		}
		gc.JSON(http.StatusBadRequest, respError)
		return
	}

	record, err := rc.service.Update(gc.Request.Context(), inputParams, inputBody)

	if err != nil {
		respError := models.RecordResponseError{
			Status:   "erro",
			Mensagem: fmt.Sprintf("erro ao gravar atualizações de registro %s", err),
			Erros:    []models.ErrorItem{},
		}
		gc.JSON(http.StatusInternalServerError, respError)
		return
	}

	gc.JSON(http.StatusOK, record)
}
