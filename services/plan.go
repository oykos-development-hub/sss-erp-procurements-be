package services

import (
	"context"

	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"

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

func (h *PlanServiceImpl) CreatePlan(ctx context.Context, input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}

	res := dto.ToPlanResponseDTO(*data)

	return &res, nil
}

func (h *PlanServiceImpl) UpdatePlan(ctx context.Context, id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error) {
	data := input.ToPlan()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}

	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) DeletePlan(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo plan delete")
	}

	return nil
}

func (h *PlanServiceImpl) GetPlan(id int) (*dto.PlanResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo plan get")
	}
	response := dto.ToPlanResponseDTO(*data)

	return &response, nil
}

func (h *PlanServiceImpl) GetPlanList(input dto.GetPlansInputDTO) ([]dto.PlanResponseDTO, *uint64, error) {
	cond := up.Cond{}
	var orders []interface{}

	if input.IsPreBudget != nil {
		if *input.IsPreBudget {
			cond["is_pre_budget"] = true
		} else {
			cond["is_pre_budget"] = false
		}
	}
	if input.Year != nil {
		cond["year"] = *input.Year
	}
	if input.TargetBudgetID != nil {
		cond["pre_budget_id"] = input.TargetBudgetID
	}

	if input.SortByDateOfPublishing != nil {
		if *input.SortByDateOfPublishing == "asc" {
			orders = append(orders, "-date_of_publishing")
		} else {
			orders = append(orders, "date_of_publishing")
		}
	}

	if input.SortByTitle != nil {
		if *input.SortByTitle == "asc" {
			orders = append(orders, "-title")
		} else {
			orders = append(orders, "title")
		}
	}

	if input.SortByYear != nil {
		if *input.SortByYear == "asc" {
			orders = append(orders, "-year")
		} else {
			orders = append(orders, "year")
		}
	}

	orders = append(orders, "-created_at")

	res, total, err := h.repo.GetAll(input.Page, input.Size, &cond, orders)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo plan get all")
	}
	response := dto.ToPlanListResponseDTO(res)

	return response, total, nil
}
