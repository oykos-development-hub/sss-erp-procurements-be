package dto

import (
	"math"
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetArticleListInput struct {
	ItemID *int `json:"public_procurement_id"`
}

type ArticleDTO struct {
	ItemID        int     `json:"public_procurement_id"`
	Title         string  `json:"title" validate:"required"`
	Description   string  `json:"description"`
	NetPrice      float32 `json:"net_price" validate:"required"`
	VATPercentage string  `json:"vat_percentage" validate:"required"`
	Manufacturer  *string `json:"manufacturer"`
	Amount        *int    `json:"amount"`
}

type ArticleResponseDTO struct {
	ID            int       `json:"id"`
	ItemID        int       `json:"public_procurement_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	NetPrice      float32   `json:"net_price"`
	VATPercentage string    `json:"vat_percentage"`
	Manufacturer  *string   `json:"manufacturer"`
	Amount        *int      `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (dto ArticleDTO) ToArticle() *data.Article {
	return &data.Article{
		Title:         dto.Title,
		ItemID:        dto.ItemID,
		Description:   dto.Description,
		NetPrice:      int(math.Round(float64(dto.NetPrice) * 100)),
		VATPercentage: dto.VATPercentage,
		Manufacturer:  dto.Manufacturer,
		Amount:        dto.Amount,
	}
}

func ToArticleResponseDTO(data data.Article) ArticleResponseDTO {
	return ArticleResponseDTO{
		ID:            data.ID,
		ItemID:        data.ItemID,
		Title:         data.Title,
		Description:   data.Description,
		NetPrice:      float32(data.NetPrice) / 100.0,
		VATPercentage: data.VATPercentage,
		Manufacturer:  data.Manufacturer,
		Amount:        data.Amount,
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
