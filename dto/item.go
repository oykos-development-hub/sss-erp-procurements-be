package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type ItemDTO struct {
	Title            string     `json:"title" validate:"required"`
	BudgetIndentID   int        `json:"budget_indent_id" validate:"required"`
	PlanID           int        `json:"plan_id" validate:"required"`
	IsOpenProcurment *bool      `json:"is_open_procurement" validate:"required"`
	ArticleType      string     `json:"article_type" validate:"required"`
	Status           *string    `json:"status"`
	SerialNumber     *string    `json:"serial_number"`
	DateOfPublishing *time.Time `json:"date_of_publishing"`
	DateOfAwarding   *time.Time `json:"date_of_awarding"`
	FileID           *int       `json:"file_id"`
}

type ItemResponseDTO struct {
	ID               int        `json:"id"`
	Title            string     `json:"title"`
	BudgetIndentID   int        `json:"budget_indent_id"`
	PlanID           int        `json:"plan_id"`
	IsOpenProcurment bool       `json:"is_open_procurement"`
	ArticleType      string     `json:"article_type"`
	Status           *string    `json:"status"`
	SerialNumber     *string    `json:"serial_number"`
	DateOfPublishing *time.Time `json:"date_of_publishing"`
	DateOfAwarding   *time.Time `json:"date_of_awarding"`
	FileID           *int       `json:"file_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type GetItemsInputDTO struct {
	Page                   *int    `json:"page" validate:"omitempty"`
	Size                   *int    `json:"size" validate:"omitempty"`
	PlanID                 *int    `json:"plan_id" validate:"omitempty"`
	SortByTitle            *string `json:"sort_by_title"`
	SortBySerialNumber     *string `json:"sort_by_serial_number"`
	SortByDateOfPublishing *string `json:"sort_by_date_of_publishing"`
	SortByDateOfAwarding   *string `json:"sort_by_date_of_awarding"`
}

func (dto ItemDTO) ToItem() *data.Item {
	return &data.Item{
		Title:            dto.Title,
		BudgetIndentID:   dto.BudgetIndentID,
		PlanID:           dto.PlanID,
		IsOpenProcurment: *dto.IsOpenProcurment,
		ArticleType:      dto.ArticleType,
		Status:           dto.Status,
		SerialNumber:     dto.SerialNumber,
		DateOfPublishing: dto.DateOfPublishing,
		DateOfAwarding:   dto.DateOfAwarding,
		FileID:           dto.FileID,
	}
}

func ToItemResponseDTO(data data.Item) ItemResponseDTO {
	return ItemResponseDTO{
		ID:               data.ID,
		Title:            data.Title,
		BudgetIndentID:   data.BudgetIndentID,
		PlanID:           data.PlanID,
		IsOpenProcurment: data.IsOpenProcurment,
		ArticleType:      data.ArticleType,
		Status:           data.Status,
		SerialNumber:     data.SerialNumber,
		DateOfPublishing: data.DateOfPublishing,
		DateOfAwarding:   data.DateOfAwarding,
		FileID:           data.FileID,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func ToItemListResponseDTO(items []*data.Item) []ItemResponseDTO {
	dtoList := make([]ItemResponseDTO, len(items))
	for i, x := range items {
		dtoList[i] = ToItemResponseDTO(*x)
	}
	return dtoList
}
