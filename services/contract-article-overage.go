package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ContractArticleOverageServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.ContractArticleOverage
}

func NewContractArticleOverageServiceImpl(app *celeritas.Celeritas, repo data.ContractArticleOverage) ContractArticleOverageService {
	return &ContractArticleOverageServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ContractArticleOverageServiceImpl) CreateContractArticleOverage(input dto.ContractArticleOverageDTO) (*dto.ContractArticleOverageResponseDTO, error) {
	data := input.ToContractArticleOverage()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToContractArticleOverageResponseDTO(*data)

	return &res, nil
}

func (h *ContractArticleOverageServiceImpl) UpdateContractArticleOverage(id int, input dto.ContractArticleOverageDTO) (*dto.ContractArticleOverageResponseDTO, error) {
	data := input.ToContractArticleOverage()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToContractArticleOverageResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleOverageServiceImpl) DeleteContractArticleOverage(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ContractArticleOverageServiceImpl) GetContractArticleOverage(id int) (*dto.ContractArticleOverageResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToContractArticleOverageResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleOverageServiceImpl) GetContractArticleOverageList(input dto.GetContractArticleOverageInputDTO) ([]dto.ContractArticleOverageResponseDTO, error) {
	cond := up.Cond{}

	if input.ArticleID != nil {
		cond["article_id"] = *input.ArticleID
	}

	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = *input.OrganizationUnitID
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToContractArticleOverageListResponseDTO(data)

	return response, nil
}
