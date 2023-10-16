package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetPlansInputDTO struct {
	Page           *int    `json:"page"`
	Size           *int    `json:"size"`
	IsPreBudget    *bool   `json:"is_pre_budget"`
	Year           *string `json:"year"`
	TargetBudgetID *int    `json:"target_budget_id"`
}

type PlanDTO struct {
	Year             string     `json:"year" validate:"required"`
	Title            string     `json:"title" validate:"required"`
	Active           bool       `json:"active"`
	SerialNumber     *string    `json:"serial_number"`
	DateOfPublishing *time.Time `json:"date_of_publishing"`
	DateOfClosing    *time.Time `json:"date_of_closing"`
	PreBudgetID      *int       `json:"pre_budget_id"`
	IsPreBudget      bool       `json:"is_pre_budget"`
	FileID           *int       `json:"file_id"`
}

type PlanResponseDTO struct {
	ID               int        `json:"id"`
	Year             string     `json:"year"`
	Title            string     `json:"title"`
	Active           bool       `json:"active"`
	SerialNumber     *string    `json:"serial_number"`
	DateOfPublishing *time.Time `json:"date_of_publishing"`
	DateOfClosing    *time.Time `json:"date_of_closing"`
	PreBudgetID      *int       `json:"pre_budget_id"`
	IsPreBudget      bool       `json:"is_pre_budget"`
	FileID           *int       `json:"file_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func (dto PlanDTO) ToPlan() *data.Plan {
	return &data.Plan{
		Title:            dto.Title,
		Year:             dto.Year,
		Active:           dto.Active,
		SerialNumber:     dto.SerialNumber,
		DateOfPublishing: dto.DateOfPublishing,
		DateOfClosing:    dto.DateOfClosing,
		PreBudgetID:      dto.PreBudgetID,
		IsPreBudget:      dto.IsPreBudget,
		FileID:           dto.FileID,
	}
}

func ToPlanResponseDTO(data data.Plan) PlanResponseDTO {
	return PlanResponseDTO{
		ID:               data.ID,
		Title:            data.Title,
		Year:             data.Year,
		Active:           data.Active,
		SerialNumber:     data.SerialNumber,
		DateOfPublishing: data.DateOfPublishing,
		DateOfClosing:    data.DateOfClosing,
		PreBudgetID:      data.PreBudgetID,
		IsPreBudget:      data.IsPreBudget,
		FileID:           data.FileID,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func ToPlanListResponseDTO(plans []*data.Plan) []PlanResponseDTO {
	dtoList := make([]PlanResponseDTO, len(plans))
	for i, x := range plans {
		dtoList[i] = ToPlanResponseDTO(*x)
	}
	return dtoList
}
