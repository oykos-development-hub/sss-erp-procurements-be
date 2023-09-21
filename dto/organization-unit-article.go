package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetOrganizationUnitArticleListInputDTO struct {
	OrganizationUnitID *int `json:"organization_unit_id" validate:"omitempty"`
	ArticleID          *int `json:"article_id" validate:"omitempty"`
}

type OrganizationUnitArticleDTO struct {
	ArticleID           int     `json:"public_procurement_article_id" validate:"required"`
	OrganizationUnitID  int     `json:"organization_unit_id" validate:"required"`
	Amount              int     `json:"amount" validate:"required"`
	Status              string  `json:"status"`
	IsRejected          bool    `json:"is_rejected"`
	RejectedDescription *string `json:"rejected_description"`
}

type OrganizationUnitArticleResponseDTO struct {
	ID                  int       `json:"id"`
	ArticleID           int       `json:"public_procurement_article_id"`
	OrganizationUnitID  int       `json:"organization_unit_id"`
	Amount              int       `json:"amount"`
	Status              string    `json:"status"`
	IsRejected          bool      `json:"is_rejected"`
	RejectedDescription *string   `json:"rejected_description"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (dto OrganizationUnitArticleDTO) ToOrganizationUnitArticle() *data.OrganizationUnitArticle {
	return &data.OrganizationUnitArticle{
		ArticleID:           dto.ArticleID,
		OrganizationUnitID:  dto.OrganizationUnitID,
		Amount:              dto.Amount,
		Status:              dto.Status,
		IsRejected:          dto.IsRejected,
		RejectedDescription: dto.RejectedDescription,
	}
}

func ToOrganizationUnitArticleResponseDTO(data data.OrganizationUnitArticle) OrganizationUnitArticleResponseDTO {
	return OrganizationUnitArticleResponseDTO{
		ID:                  data.ID,
		ArticleID:           data.ArticleID,
		OrganizationUnitID:  data.OrganizationUnitID,
		Amount:              data.Amount,
		Status:              data.Status,
		IsRejected:          data.IsRejected,
		RejectedDescription: data.RejectedDescription,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
}

func ToOrganizationUnitArticleListResponseDTO(organizationunitarticles []*data.OrganizationUnitArticle) []OrganizationUnitArticleResponseDTO {
	dtoList := make([]OrganizationUnitArticleResponseDTO, len(organizationunitarticles))
	for i, x := range organizationunitarticles {
		dtoList[i] = ToOrganizationUnitArticleResponseDTO(*x)
	}
	return dtoList
}
