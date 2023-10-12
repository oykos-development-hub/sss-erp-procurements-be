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

// ContractHandler is a concrete type that implements ContractHandler
type contractHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.ContractService
}

// NewContractHandler initializes a new ContractHandler with its dependencies
func NewContractHandler(app *celeritas.Celeritas, contractService services.ContractService) ContractHandler {
	return &contractHandlerImpl{
		App:     app,
		service: contractService,
	}
}

func (h *contractHandlerImpl) CreateContract(w http.ResponseWriter, r *http.Request) {
	var input dto.ContractDTO
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

	res, err := h.service.CreateContract(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error creating contract: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Contract created successfuly", res)
}

func (h *contractHandlerImpl) UpdateContract(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.ContractDTO
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

	res, err := h.service.UpdateContract(id, input)
	if err != nil {
		h.App.ErrorLog.Printf("Error updating contract with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Contract updated successfuly", res)
}

func (h *contractHandlerImpl) DeleteContract(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteContract(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error deleting contract with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Contract deleted successfuly")
}

func (h *contractHandlerImpl) GetContractById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetContract(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *contractHandlerImpl) GetContractList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetContractsInputDTO
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

	res, total, err := h.service.GetContractList(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract list: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
