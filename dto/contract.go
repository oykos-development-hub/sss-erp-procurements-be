package dto

import (
	"time"

	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetContractsInputDTO struct {
	Page          *int `json:"page" validate:"omitempty"`
	Size          *int `json:"size" validate:"omitempty"`
	ProcurementID *int `json:"procurement_id"`
	SupplierID    *int `json:"supplier_id"`
}

type ContractDTO struct {
	PublicProcurementID int        `json:"public_procurement_id" validate:"required"`
	SupplierID          int        `json:"supplier_id"  validate:"required"`
	SerialNumber        string     `json:"serial_number"  validate:"required"`
	DateOfSigning       time.Time  `json:"date_of_signing"  validate:"required"`
	DateOfExpiry        *time.Time `json:"date_of_expiry"`
	NetValue            float32    `json:"net_value"`
	GrossValue          float32    `json:"gross_value"`
	VatValue            float32    `json:"vat_value"`
	FileID              *int       `json:"file_id"`
}

type ContractResponseDTO struct {
	ID                  int        `json:"id"`
	PublicProcurementID int        `json:"public_procurement_id"`
	SupplierID          int        `json:"supplier_id"`
	SerialNumber        string     `json:"serial_number"`
	DateOfSigning       time.Time  `json:"date_of_signing"`
	DateOfExpiry        *time.Time `json:"date_of_expiry"`
	NetValue            float32    `json:"net_value"`
	GrossValue          float32    `json:"gross_value"`
	VatValue            float32    `json:"vat_value"`
	FileID              *int       `json:"file_id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (dto ContractDTO) ToContract() *data.Contract {
	return &data.Contract{
		PublicProcurementID: dto.PublicProcurementID,
		SupplierID:          dto.SupplierID,
		SerialNumber:        dto.SerialNumber,
		DateOfSigning:       dto.DateOfSigning,
		DateOfExpiry:        dto.DateOfExpiry,
		NetValue:            int(dto.NetValue * 100),
		GrossValue:          int(dto.GrossValue * 100),
		VatValue:            int(dto.GrossValue * 100),
		FileID:              dto.FileID,
	}
}

func ToContractResponseDTO(data data.Contract) ContractResponseDTO {
	return ContractResponseDTO{
		ID:                  data.ID,
		PublicProcurementID: data.PublicProcurementID,
		SupplierID:          data.SupplierID,
		SerialNumber:        data.SerialNumber,
		DateOfSigning:       data.DateOfSigning,
		DateOfExpiry:        data.DateOfExpiry,
		NetValue:            float32(data.NetValue) / 100.0,
		GrossValue:          float32(data.GrossValue) / 100.0,
		VatValue:            float32(data.VatValue) / 100.0,
		FileID:              data.FileID,
		CreatedAt:           data.CreatedAt,
		UpdatedAt:           data.UpdatedAt,
	}
}

func ToContractListResponseDTO(contracts []*data.Contract) []ContractResponseDTO {
	dtoList := make([]ContractResponseDTO, len(contracts))
	for i, x := range contracts {
		dtoList[i] = ToContractResponseDTO(*x)
	}
	return dtoList
}
