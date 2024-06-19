package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"

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
		return nil, newErrors.Wrap(err, "repo contract article insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo contract article get")
	}

	res := dto.ToContractArticleResponseDTO(*data)

	return &res, nil
}

func (h *ContractArticleServiceImpl) UpdateContractArticle(id int, input dto.ContractArticleDTO) (*dto.ContractArticleResponseDTO, error) {
	data := input.ToContractArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo contract article update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo contract article get")
	}

	response := dto.ToContractArticleResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleServiceImpl) DeleteContractArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo contract article delete")
	}

	return nil
}

func (h *ContractArticleServiceImpl) GetContractArticle(id int) (*dto.ContractArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo contract article get")
	}
	response := dto.ToContractArticleResponseDTO(*data)

	return &response, nil
}

func (h *ContractArticleServiceImpl) GetContractArticleList(input *dto.GetContractArticlesInputDTO) ([]dto.ContractArticleResponseDTO, *uint64, error) {
	cond := up.Cond{}
	var orders []interface{}

	if input.ContractID != nil {
		cond["contract_id"] = *input.ContractID
	}

	if input.ArticleID != nil {
		cond["article_id"] = *input.ArticleID
	}

	if input.SortByNetValue != nil {
		if *input.SortByNetValue == "asc" {
			orders = append(orders, "-net_value")
		} else {
			orders = append(orders, "net_value")
		}
	}

	orders = append(orders, "-created_at")

	data, total, err := h.repo.GetAll(&cond, orders)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo contract article get all")
	}
	response := dto.ToContractArticleListResponseDTO(data)

	return response, total, nil
}
