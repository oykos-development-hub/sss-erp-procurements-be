package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ItemServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Item
}

func NewItemServiceImpl(app *celeritas.Celeritas, repo data.Item) ItemService {
	return &ItemServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ItemServiceImpl) CreateItem(input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToItemResponseDTO(*data)

	return &res, nil
}

func (h *ItemServiceImpl) UpdateItem(id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error) {
	data := input.ToItem()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToItemResponseDTO(*data)

	return &response, nil
}

func (h *ItemServiceImpl) DeleteItem(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ItemServiceImpl) GetItem(id int) (*dto.ItemResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToItemResponseDTO(*data)

	return &response, nil
}

func (h *ItemServiceImpl) GetItemList(input dto.GetItemsInputDTO) ([]dto.ItemResponseDTO, *uint64, error) {
	cond := up.Cond{}
	if input.PlanID != nil {
		cond["plan_id"] = input.PlanID
	}

	res, total, err := h.repo.GetAll(input.Page, input.Size, &cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToItemListResponseDTO(res)

	return response, total, nil
}
