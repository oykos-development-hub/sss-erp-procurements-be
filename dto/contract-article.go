package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetContractArticlesInputDTO struct {
	ContractID *int `json:"contract_id"`
}

type ContractArticleDTO struct {
	ArticleID  int     `json:"public_procurement_article_id" validate:"required"`
	ContractID int     `json:"public_procurement_contract_id" validate:"required"`
	Amount     int     `json:"amount" validate:"required"`
	NetValue   *string `json:"net_value"`
	GrossValue *string `json:"gross_value"`
}

type ContractArticleResponseDTO struct {
	ID         int       `json:"id"`
	ArticleID  int       `json:"public_procurement_article_id"`
	ContractID int       `json:"public_procurement_contract_id"`
	Amount     int       `json:"amount"`
	NetValue   *string   `json:"net_value"`
	GrossValue *string   `json:"gross_value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (dto ContractArticleDTO) ToContractArticle() *data.ContractArticle {
	return &data.ContractArticle{
		ArticleID:  dto.ArticleID,
		ContractID: dto.ContractID,
		Amount:     dto.Amount,
		NetValue:   dto.NetValue,
		GrossValue: dto.GrossValue,
	}
}

func ToContractArticleResponseDTO(data data.ContractArticle) ContractArticleResponseDTO {
	return ContractArticleResponseDTO{
		ID:         data.ID,
		ArticleID:  data.ArticleID,
		ContractID: data.ContractID,
		Amount:     data.Amount,
		NetValue:   data.NetValue,
		GrossValue: data.GrossValue,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func ToContractArticleListResponseDTO(contractArticles []*data.ContractArticle) []ContractArticleResponseDTO {
	dtoList := make([]ContractArticleResponseDTO, len(contractArticles))
	for i, x := range contractArticles {
		dtoList[i] = ToContractArticleResponseDTO(*x)
	}
	return dtoList
}
