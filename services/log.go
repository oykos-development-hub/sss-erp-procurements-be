package services

import (
	"fmt"

	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type LogServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Log
}

func NewLogServiceImpl(app *celeritas.Celeritas, repo data.Log) LogService {
	return &LogServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *LogServiceImpl) CreateLog(input dto.LogDTO) (*dto.LogResponseDTO, error) {
	data := input.ToLog()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log get")
	}

	res := dto.ToLogResponseDTO(*data)

	return &res, nil
}

func (h *LogServiceImpl) UpdateLog(id int, input dto.LogDTO) (*dto.LogResponseDTO, error) {
	data := input.ToLog()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log get")
	}

	response := dto.ToLogResponseDTO(*data)

	return &response, nil
}

func (h *LogServiceImpl) DeleteLog(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo log delete")
	}

	return nil
}

func (h *LogServiceImpl) GetLog(id int) (*dto.LogResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log get")
	}
	response := dto.ToLogResponseDTO(*data)

	return &response, nil
}

func (h *LogServiceImpl) GetLogList(filter dto.LogFilterDTO) ([]dto.LogResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	var orders []interface{}

	if filter.Entity != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"entity": *filter.Entity})
	}

	if filter.Operation != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"operation": *filter.Operation})
	}

	if filter.UserID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"user_id": *filter.UserID})
	}

	if filter.ItemID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"item_id": *filter.ItemID})
	}

	if filter.Search != nil {
		likeCondition := fmt.Sprintf("%%%s%%", *filter.Search)
		conditionAndExp = up.And(
			up.Or(
				up.Cond{"old_state ::text ILIKE": likeCondition},
				up.Cond{"new_state ::text ILIKE": likeCondition},
			),
		)
	}

	if filter.SortByTitle != nil {
		if *filter.SortByTitle == "asc" {
			orders = append(orders, "-title")
		} else {
			orders = append(orders, "title")
		}
	}

	orders = append(orders, "-id")

	data, total, err := h.repo.GetAll(filter.Page, filter.Size, conditionAndExp, orders)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo log get all")
	}
	response := dto.ToLogListResponseDTO(data)

	return response, total, nil
}
