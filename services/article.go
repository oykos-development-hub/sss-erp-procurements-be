package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Article
}

func NewArticleServiceImpl(app *celeritas.Celeritas, repo data.Article) ArticleService {
	return &ArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ArticleServiceImpl) CreateArticle(input dto.ArticleDTO) (*dto.ArticleResponseDTO, error) {
	data := input.ToArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToArticleResponseDTO(*data)

	return &res, nil
}

func (h *ArticleServiceImpl) UpdateArticle(id int, input dto.ArticleDTO) (*dto.ArticleResponseDTO, error) {
	data := input.ToArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToArticleResponseDTO(*data)

	return &response, nil
}

func (h *ArticleServiceImpl) DeleteArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}
	return nil
}

func (h *ArticleServiceImpl) GetArticle(id int) (*dto.ArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToArticleResponseDTO(*data)

	return &response, nil
}

func (h *ArticleServiceImpl) GetArticleList(input *dto.GetArticleListInput) ([]dto.ArticleResponseDTO, error) {
	cond := up.Cond{}
	var orders []interface{}

	if input.ItemID != nil {
		cond["item_id"] = *input.ItemID
	}

	if input.Title != nil {
		cond["title"] = *input.Title
	}

	if input.Description != nil {
		cond["description"] = *input.Description
	}

	if input.VisibilityType != nil {
		cond["visibility_type"] = *input.VisibilityType
	}

	if input.SortByPrice != nil {
		if *input.SortByPrice == "asc" {
			orders = append(orders, "-price")
		} else {
			orders = append(orders, "price")
		}
	}

	if input.SortByTitle != nil {
		if *input.SortByTitle == "asc" {
			orders = append(orders, "-title")
		} else {
			orders = append(orders, "title")
		}
	}

	orders = append(orders, "-created_at")

	data, err := h.repo.GetAll(&cond, orders)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToArticleListResponseDTO(data)

	return response, nil
}
