package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetContractArticleOverageInputDTO struct {
	ArticleID *int `json:"article_id"`
}

type ContractArticleOverageDTO struct {
	Amount    int `json:"amount"`
	ArticleID int `json:"article_id"`
}

type ContractArticleOverageResponseDTO struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	ArticleID int       `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (dto ContractArticleOverageDTO) ToContractArticleOverage() *data.ContractArticleOverage {
	return &data.ContractArticleOverage{
		ArticleID: dto.ArticleID,
		Amount:    dto.Amount,
	}
}

func ToContractArticleOverageResponseDTO(data data.ContractArticleOverage) ContractArticleOverageResponseDTO {
	return ContractArticleOverageResponseDTO{
		ID:        data.ID,
		ArticleID: data.ArticleID,
		Amount:    data.Amount,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToContractArticleOverageListResponseDTO(contractarticleoverages []*data.ContractArticleOverage) []ContractArticleOverageResponseDTO {
	dtoList := make([]ContractArticleOverageResponseDTO, len(contractarticleoverages))
	for i, x := range contractarticleoverages {
		dtoList[i] = ToContractArticleOverageResponseDTO(*x)
	}
	return dtoList
}
