package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type OrganizationUnitPlanLimitServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.OrganizationUnitPlanLimit
}

func NewOrganizationUnitPlanLimitServiceImpl(app *celeritas.Celeritas, repo data.OrganizationUnitPlanLimit) OrganizationUnitPlanLimitService {
	return &OrganizationUnitPlanLimitServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *OrganizationUnitPlanLimitServiceImpl) CreateOrganizationUnitPlanLimit(input dto.OrganizationUnitPlanLimitDTO) (*dto.OrganizationUnitPlanLimitResponseDTO, error) {
	data := input.ToOrganizationUnitPlanLimit()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit get")
	}

	res := dto.ToOrganizationUnitPlanLimitResponseDTO(*data)

	return &res, nil
}

func (h *OrganizationUnitPlanLimitServiceImpl) UpdateOrganizationUnitPlanLimit(id int, input dto.OrganizationUnitPlanLimitDTO) (*dto.OrganizationUnitPlanLimitResponseDTO, error) {
	data := input.ToOrganizationUnitPlanLimit()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit get")
	}

	response := dto.ToOrganizationUnitPlanLimitResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitPlanLimitServiceImpl) DeleteOrganizationUnitPlanLimit(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo organization unit plan limit delete")
	}

	return nil
}

func (h *OrganizationUnitPlanLimitServiceImpl) GetOrganizationUnitPlanLimit(id int) (*dto.OrganizationUnitPlanLimitResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit get")
	}
	response := dto.ToOrganizationUnitPlanLimitResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitPlanLimitServiceImpl) GetOrganizationUnitPlanLimitList(input dto.OrganizationUnitPlanLimitInputDTO) ([]dto.OrganizationUnitPlanLimitResponseDTO, error) {
	cond := up.Cond{}

	if input.ItemID != nil {
		cond["item_id"] = input.ItemID
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit plan limit get all")
	}
	response := dto.ToOrganizationUnitPlanLimitListResponseDTO(data)

	return response, nil
}
