package handlers

import (
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"
	"gitlab.sudovi.me/erp/procurements-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// ContractArticleHandler is a concrete type that implements ContractHandler
type contractArticleHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.ContractArticleService
}

// NewContractArticleHandler initializes a new ContractArticleHandler with its dependencies
func NewContractArticleHandler(app *celeritas.Celeritas, contractArticleService services.ContractArticleService) ContractArticleHandler {
	return &contractArticleHandlerImpl{
		App:     app,
		service: contractArticleService,
	}
}

func (h *contractArticleHandlerImpl) CreateContractArticle(w http.ResponseWriter, r *http.Request) {
	var input dto.ContractArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateContractArticle(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error creating contract article: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticle created successfuly", res)
}

func (h *contractArticleHandlerImpl) UpdateContractArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.ContractArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateContractArticle(id, input)
	if err != nil {
		h.App.ErrorLog.Printf("Error updating contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticle updated successfuly", res)
}

func (h *contractArticleHandlerImpl) DeleteContractArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteContractArticle(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error deleting contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "ContractArticle deleted successfuly")
}

func (h *contractArticleHandlerImpl) GetContractArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetContractArticle(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *contractArticleHandlerImpl) GetContractArticleList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetContractArticlesInputDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetContractArticleList(&input)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract article list: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
