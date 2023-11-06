package dto

import (
	"math"
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetContractArticlesInputDTO struct {
	ContractID *int `json:"contract_id"`
}

type ContractArticleDTO struct {
	ArticleID  int      `json:"public_procurement_article_id" validate:"required"`
	ContractID int      `json:"public_procurement_contract_id" validate:"required"`
	NetValue   *float32 `json:"net_value"`
	GrossValue *float32 `json:"gross_value"`
}

type ContractArticleResponseDTO struct {
	ID         int       `json:"id"`
	ArticleID  int       `json:"public_procurement_article_id"`
	ContractID int       `json:"public_procurement_contract_id"`
	NetValue   *float32  `json:"net_value"`
	GrossValue *float32  `json:"gross_value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (dto ContractArticleDTO) ToContractArticle() *data.ContractArticle {
	data := &data.ContractArticle{
		ArticleID:  dto.ArticleID,
		ContractID: dto.ContractID,
	}

	if dto.NetValue != nil {
		net := int(math.Round(float64(*dto.NetValue) * 100))
		data.NetValue = &net
	}
	if dto.GrossValue != nil {
		gross := int(math.Round(float64(*dto.GrossValue) * 100))
		data.GrossValue = &gross
	}

	return data
}

func ToContractArticleResponseDTO(data data.ContractArticle) ContractArticleResponseDTO {
	res := ContractArticleResponseDTO{
		ID:         data.ID,
		ArticleID:  data.ArticleID,
		ContractID: data.ContractID,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}

	if data.NetValue != nil {
		net := float32(*data.NetValue) / 100.0 // converting cents back to float
		res.NetValue = &net
	}
	if data.GrossValue != nil {
		gross := float32(*data.GrossValue) / 100.0 // converting cents back to float
		res.GrossValue = &gross
	}

	return res
}

func ToContractArticleListResponseDTO(contractArticles []*data.ContractArticle) []ContractArticleResponseDTO {
	dtoList := make([]ContractArticleResponseDTO, len(contractArticles))
	for i, x := range contractArticles {
		dtoList[i] = ToContractArticleResponseDTO(*x)
	}
	return dtoList
}
