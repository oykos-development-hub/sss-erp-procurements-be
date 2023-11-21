package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ContractArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.ContractArticle
}

func NewContractArticleServiceImpl(app *celeritas.Celeritas, repo data.ContractArticle) ContractArticleService {
	return &ContractArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ContractArticleServiceImpl) CreateContractArticle(input dto.ContractArticleDTO) (*dto.ContractArticleResponseDTO, error) {
	data := input.ToContractArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToContractArticleResponseDTO(*data)

	return &res, nil
}

func (h *ContractArticleServiceImpl) UpdateContractArticle(id int, input dto.ContractArticleDTO) (*dto.ContractArticleResponseDTO, error) {
	data := input.ToContractArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToContractArticleResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleServiceImpl) DeleteContractArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ContractArticleServiceImpl) GetContractArticle(id int) (*dto.ContractArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToContractArticleResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleServiceImpl) GetContractArticleList(input *dto.GetContractArticlesInputDTO) ([]dto.ContractArticleResponseDTO, *uint64, error) {
	cond := up.Cond{}

	if input.ContractID != nil {
		cond["contract_id"] = *input.ContractID
	}

	if input.ArticleID != nil {
		cond["article_id"] = *input.ArticleID
	}

	data, total, err := h.repo.GetAll(&cond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToContractArticleListResponseDTO(data)

	return response, total, nil
}
