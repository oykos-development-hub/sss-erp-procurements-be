package services

import (
	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/dto"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type OrganizationUnitArticleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.OrganizationUnitArticle
}

func NewOrganizationUnitArticleServiceImpl(app *celeritas.Celeritas, repo data.OrganizationUnitArticle) OrganizationUnitArticleService {
	return &OrganizationUnitArticleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *OrganizationUnitArticleServiceImpl) CreateOrganizationUnitArticle(input dto.OrganizationUnitArticleDTO) (*dto.OrganizationUnitArticleResponseDTO, error) {
	data := input.ToOrganizationUnitArticle()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article get")
	}

	res := dto.ToOrganizationUnitArticleResponseDTO(*data)

	return &res, nil
}

func (h *OrganizationUnitArticleServiceImpl) UpdateOrganizationUnitArticle(id int, input dto.OrganizationUnitArticleDTO) (*dto.OrganizationUnitArticleResponseDTO, error) {
	data := input.ToOrganizationUnitArticle()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article get")
	}

	response := dto.ToOrganizationUnitArticleResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitArticleServiceImpl) DeleteOrganizationUnitArticle(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo organization unit article delete")
	}

	return nil
}

func (h *OrganizationUnitArticleServiceImpl) GetOrganizationUnitArticle(id int) (*dto.OrganizationUnitArticleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article get")
	}
	response := dto.ToOrganizationUnitArticleResponseDTO(*data)

	return &response, nil
}

func (h *OrganizationUnitArticleServiceImpl) GetOrganizationUnitArticleList(input dto.GetOrganizationUnitArticleListInputDTO) ([]dto.OrganizationUnitArticleResponseDTO, error) {
	cond := up.Cond{}

	if input.ArticleID != nil {
		cond["article_id"] = *input.ArticleID
	}
	if input.OrganizationUnitID != nil {
		cond["organization_unit_id"] = *input.OrganizationUnitID
	}

	data, err := h.repo.GetAll(&cond)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo organization unit article get all")
	}
	response := dto.ToOrganizationUnitArticleListResponseDTO(data)

	return response, nil
}
