package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetArticleListInput struct {
	ItemID *int `json:"public_procurement_id"`
}

type ArticleDTO struct {
	BudgetID      int    `json:"budget_indent_id"`
	ItemID        int    `json:"public_procurement_id"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	NetPrice      string `json:"net_price" validate:"required"`
	VATPercentage string `json:"vat_percentage" validate:"required"`
}

type ArticleResponseDTO struct {
	ID            int       `json:"id"`
	BudgetID      int       `json:"budget_indent_id"`
	ItemID        int       `json:"public_procurement_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	NetPrice      string    `json:"net_price"`
	VATPercentage string    `json:"vat_percentage"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (dto ArticleDTO) ToArticle() *data.Article {
	return &data.Article{
		Title:         dto.Title,
		BudgetID:      dto.BudgetID,
		ItemID:        dto.ItemID,
		Description:   dto.Description,
		NetPrice:      dto.NetPrice,
		VATPercentage: dto.VATPercentage,
	}
}

func ToArticleResponseDTO(data data.Article) ArticleResponseDTO {
	return ArticleResponseDTO{
		ID:            data.ID,
		BudgetID:      data.BudgetID,
		ItemID:        data.ItemID,
		Title:         data.Title,
		Description:   data.Description,
		NetPrice:      data.NetPrice,
		VATPercentage: data.VATPercentage,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func ToArticleListResponseDTO(articles []*data.Article) []ArticleResponseDTO {
	dtoList := make([]ArticleResponseDTO, len(articles))
	for i, x := range articles {
		dtoList[i] = ToArticleResponseDTO(*x)
	}
	return dtoList
}
