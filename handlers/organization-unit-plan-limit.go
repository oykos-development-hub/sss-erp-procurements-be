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

// OrganizationUnitPlanLimitHandler is a concrete type that implements OrganizationUnitPlanLimitHandler
type organizationunitplanlimitHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.OrganizationUnitPlanLimitService
}

// NewOrganizationUnitPlanLimitHandler initializes a new OrganizationUnitPlanLimitHandler with its dependencies
func NewOrganizationUnitPlanLimitHandler(app *celeritas.Celeritas, organizationunitplanlimitService services.OrganizationUnitPlanLimitService) OrganizationUnitPlanLimitHandler {
	return &organizationunitplanlimitHandlerImpl{
		App:     app,
		service: organizationunitplanlimitService,
	}
}

func (h *organizationunitplanlimitHandlerImpl) CreateOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request) {
	var input dto.OrganizationUnitPlanLimitDTO
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

	res, err := h.service.CreateOrganizationUnitPlanLimit(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error organization unit plan limit: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrganizationUnitPlanLimit created successfuly", res)
}

func (h *organizationunitplanlimitHandlerImpl) UpdateOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.OrganizationUnitPlanLimitDTO
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

	res, err := h.service.UpdateOrganizationUnitPlanLimit(id, input)
	if err != nil {
		h.App.ErrorLog.Printf("Error updating organization unit plan limit with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrganizationUnitPlanLimit updated successfuly", res)
}

func (h *organizationunitplanlimitHandlerImpl) DeleteOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteOrganizationUnitPlanLimit(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error deleting organization unit plan limit with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "OrganizationUnitPlanLimit deleted successfuly")
}

func (h *organizationunitplanlimitHandlerImpl) GetOrganizationUnitPlanLimitById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetOrganizationUnitPlanLimit(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching organization unit plan limit with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *organizationunitplanlimitHandlerImpl) GetOrganizationUnitPlanLimitList(w http.ResponseWriter, r *http.Request) {
	var input dto.OrganizationUnitPlanLimitInputDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.service.GetOrganizationUnitPlanLimitList(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching organization unit plan limit list: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
