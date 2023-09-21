package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type PlanServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Plan
}

func NewPlanServiceImpl(app *celeritas.Celeritas, repo data.Plan) PlanService {
	return &PlanServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *PlanServiceImpl) CreatePlan(input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToPlanResponseDTO(*data)

	return &res, nil
}

func (h *PlanServiceImpl) UpdatePlan(id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) DeletePlan(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *PlanServiceImpl) GetPlan(id int) (*dto.PlanResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) GetPlanList(input dto.GetPlansInputDTO) ([]dto.PlanResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.IsPreBudget != nil {
		if *input.IsPreBudget {
			cond["pre_budget_id"] = up.IsNull()
		} else {
			cond["pre_budget_id"] = up.IsNotNull()
		}
	}
	if input.Year != nil {
		cond["year"] = *input.Year
	}
	if input.TargetBudgetID != nil {
		cond["pre_budget_id"] = input.TargetBudgetID
	}

	res, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToPlanListResponseDTO(res)

	return response, total, nil
}
