package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ContractServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Contract
}

func NewContractServiceImpl(app *celeritas.Celeritas, repo data.Contract) ContractService {
	return &ContractServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ContractServiceImpl) CreateContract(input dto.ContractDTO) (*dto.ContractResponseDTO, error) {
	data := input.ToContract()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToContractResponseDTO(*data)

	return &res, nil
}

func (h *ContractServiceImpl) UpdateContract(id int, input dto.ContractDTO) (*dto.ContractResponseDTO, error) {
	data := input.ToContract()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToContractResponseDTO(*data)

	return &response, nil
}

func (h *ContractServiceImpl) DeleteContract(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ContractServiceImpl) GetContract(id int) (*dto.ContractResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToContractResponseDTO(*data)

	return &response, nil
}

func (h *ContractServiceImpl) GetContractList(input dto.GetContractsInputDTO) ([]dto.ContractResponseDTO, *uint64, error) {
	cond := up.Cond{}
	var orders []interface{}

	if input.ProcurementID != nil {
		cond["public_procurement_id"] = input.ProcurementID
	}
	if input.SupplierID != nil {
		cond["supplier_id"] = input.SupplierID
	}
	if input.SortDateOfExpiry != nil {
		if *input.SortDateOfExpiry == "asc" {
			orders = append(orders, "-date_of_expiry")
		} else {
			orders = append(orders, "date_of_expiry")
		}
	}
	orders = append(orders, "-created_at")

	res, total, err := h.repo.GetAll(input.Page, input.Size, &cond, orders)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToContractListResponseDTO(res)

	return response, total, nil
}
