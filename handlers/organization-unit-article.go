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

// OrganizationUnitArticleHandler is a concrete type that implements OrganizationUnitArticleHandler
type organizationunitarticleHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.OrganizationUnitArticleService
}

// NewOrganizationUnitArticleHandler initializes a new OrganizationUnitArticleHandler with its dependencies
func NewOrganizationUnitArticleHandler(app *celeritas.Celeritas, organizationunitarticleService services.OrganizationUnitArticleService) OrganizationUnitArticleHandler {
	return &organizationunitarticleHandlerImpl{
		App:     app,
		service: organizationunitarticleService,
	}
}

func (h *organizationunitarticleHandlerImpl) CreateOrganizationUnitArticle(w http.ResponseWriter, r *http.Request) {
	var input dto.OrganizationUnitArticleDTO
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

	res, err := h.service.CreateOrganizationUnitArticle(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrganizationUnitArticle created successfuly", res)
}

func (h *organizationunitarticleHandlerImpl) UpdateOrganizationUnitArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.OrganizationUnitArticleDTO
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

	res, err := h.service.UpdateOrganizationUnitArticle(id, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "OrganizationUnitArticle updated successfuly", res)
}

func (h *organizationunitarticleHandlerImpl) DeleteOrganizationUnitArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteOrganizationUnitArticle(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "OrganizationUnitArticle deleted successfuly")
}

func (h *organizationunitarticleHandlerImpl) GetOrganizationUnitArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetOrganizationUnitArticle(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *organizationunitarticleHandlerImpl) GetOrganizationUnitArticleList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetOrganizationUnitArticleListInputDTO
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

	res, err := h.service.GetOrganizationUnitArticleList(input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
