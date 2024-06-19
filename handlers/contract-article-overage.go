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

// ContractArticleOverageHandler is a concrete type that implements ContractArticleOverageHandler
type contractarticleoverageHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.ContractArticleOverageService
}

// NewContractArticleOverageHandler initializes a new ContractArticleOverageHandler with its dependencies
func NewContractArticleOverageHandler(app *celeritas.Celeritas, contractarticleoverageService services.ContractArticleOverageService) ContractArticleOverageHandler {
	return &contractarticleoverageHandlerImpl{
		App:     app,
		service: contractarticleoverageService,
	}
}

func (h *contractarticleoverageHandlerImpl) CreateContractArticleOverage(w http.ResponseWriter, r *http.Request) {
	var input dto.ContractArticleOverageDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateContractArticleOverage(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticleOverage created successfuly", res)
}

func (h *contractarticleoverageHandlerImpl) UpdateContractArticleOverage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.ContractArticleOverageDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateContractArticleOverage(id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticleOverage updated successfuly", res)
}

func (h *contractarticleoverageHandlerImpl) DeleteContractArticleOverage(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteContractArticleOverage(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "ContractArticleOverage deleted successfuly")
}

func (h *contractarticleoverageHandlerImpl) GetContractArticleOverageById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetContractArticleOverage(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *contractarticleoverageHandlerImpl) GetContractArticleOverageList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetContractArticleOverageInputDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.GetContractArticleOverageList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
