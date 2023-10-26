package dto

import (
	"fmt"
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
	NetValue            *float32   `json:"net_value"`
	GrossValue          *float32   `json:"gross_value"`
	VatValue            *float32   `json:"vat_value"`
	FileID              *int       `json:"file_id"`
}

type ContractResponseDTO struct {
	ID                  int        `json:"id"`
	PublicProcurementID int        `json:"public_procurement_id"`
	SupplierID          int        `json:"supplier_id"`
	SerialNumber        string     `json:"serial_number"`
	DateOfSigning       time.Time  `json:"date_of_signing"`
	DateOfExpiry        *time.Time `json:"date_of_expiry"`
	NetValue            *float32   `json:"net_value"`
	GrossValue          *float32   `json:"gross_value"`
	VatValue            *float32   `json:"vat_value"`
	FileID              *int       `json:"file_id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (dto ContractDTO) ToContract() *data.Contract {
	var netValue, grossValue, vatValue *int
	if dto.NetValue != nil {
		netValueInt := int(*dto.NetValue * 100)
		netValue = &netValueInt
	}

	if dto.GrossValue != nil {
		fmt.Printf("Original GrossValue: %v\n", *dto.GrossValue) // Debug line
		grossValueInt := int(*dto.GrossValue * 100)
		grossValue = &grossValueInt
	}
	if dto.VatValue != nil {
		fmt.Printf("Original VatValue: %v\n", *dto.VatValue) // Debug line
		vatValueInt := int(*dto.VatValue * 100)
		vatValue = &vatValueInt
	}
	fmt.Printf("VatValue: %v\n, Gross value: %v\n", *grossValue, *vatValue) // Debug line
	return &data.Contract{
		PublicProcurementID: dto.PublicProcurementID,
		SupplierID:          dto.SupplierID,
		SerialNumber:        dto.SerialNumber,
		DateOfSigning:       dto.DateOfSigning,
		DateOfExpiry:        dto.DateOfExpiry,
		NetValue:            netValue,
		GrossValue:          grossValue,
		VatValue:            vatValue,
		FileID:              dto.FileID,
	}
}

func ToContractResponseDTO(data data.Contract) ContractResponseDTO {
	var netValue, grossValue, vatValue *float32
	if data.NetValue != nil {
		netValueFloat := float32(*data.NetValue) / 100.0
		netValue = &netValueFloat
	}
	if data.GrossValue != nil {
		grossValueFloat := float32(*data.GrossValue) / 100.0
		grossValue = &grossValueFloat
	}
	if data.VatValue != nil {
		vatValueFloat := float32(*data.VatValue) / 100.0
		vatValue = &vatValueFloat
	}
	return ContractResponseDTO{
		ID:                  data.ID,
		PublicProcurementID: data.PublicProcurementID,
		SupplierID:          data.SupplierID,
		SerialNumber:        data.SerialNumber,
		DateOfSigning:       data.DateOfSigning,
		DateOfExpiry:        data.DateOfExpiry,
		NetValue:            netValue,
		GrossValue:          grossValue,
		VatValue:            vatValue,
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
