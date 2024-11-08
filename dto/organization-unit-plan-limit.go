package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type OrganizationUnitPlanLimitDTO struct {
	ItemID             int `json:"public_procurement_id" validate:"required"`
	OrganizationUnitID int `json:"organization_unit_id" validate:"required"`
	Limit              int `json:"limit"`
}

type OrganizationUnitPlanLimitInputDTO struct {
	ItemID *int `json:"procurement_id" validate:"omitempty"`
}

type OrganizationUnitPlanLimitResponseDTO struct {
	ID                 int       `json:"id"`
	ItemID             int       `json:"public_procurement_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Limit              int       `json:"limit"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (dto OrganizationUnitPlanLimitDTO) ToOrganizationUnitPlanLimit() *data.OrganizationUnitPlanLimit {
	return &data.OrganizationUnitPlanLimit{
		ItemID:             dto.ItemID,
		OrganizationUnitID: dto.OrganizationUnitID,
		Limit:              dto.Limit,
	}
}

func ToOrganizationUnitPlanLimitResponseDTO(data data.OrganizationUnitPlanLimit) OrganizationUnitPlanLimitResponseDTO {
	return OrganizationUnitPlanLimitResponseDTO{
		ID:                 data.ID,
		ItemID:             data.ItemID,
		OrganizationUnitID: data.OrganizationUnitID,
		Limit:              data.Limit,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToOrganizationUnitPlanLimitListResponseDTO(organizationunitplanlimits []*data.OrganizationUnitPlanLimit) []OrganizationUnitPlanLimitResponseDTO {
	dtoList := make([]OrganizationUnitPlanLimitResponseDTO, len(organizationunitplanlimits))
	for i, x := range organizationunitplanlimits {
		dtoList[i] = ToOrganizationUnitPlanLimitResponseDTO(*x)
	}
	return dtoList
}
