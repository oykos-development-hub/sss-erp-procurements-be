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

// PlanHandler is a concrete type that implements PlanHandler
type planHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.PlanService
}

// NewPlanHandler initializes a new PlanHandler with its dependencies
func NewPlanHandler(app *celeritas.Celeritas, planService services.PlanService) PlanHandler {
	return &planHandlerImpl{
		App:     app,
		service: planService,
	}
}

func (h *planHandlerImpl) CreatePlan(w http.ResponseWriter, r *http.Request) {
	var input dto.PlanDTO
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

	res, err := h.service.CreatePlan(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error creating plan: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Plan created successfuly", res)
}

func (h *planHandlerImpl) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.PlanDTO
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

	res, err := h.service.UpdatePlan(id, input)
	if err != nil {
		h.App.ErrorLog.Printf("Error updating plan: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Plan updated successfuly", res)
}

func (h *planHandlerImpl) DeletePlan(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeletePlan(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error deleting plan: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Plan deleted successfuly")
}

func (h *planHandlerImpl) GetPlanById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetPlan(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error getting plan: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *planHandlerImpl) GetPlanList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetPlansInputDTO
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

	res, total, err := h.service.GetPlanList(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error getting plans: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
