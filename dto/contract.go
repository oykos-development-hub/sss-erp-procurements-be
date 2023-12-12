package dto

import (
	"math"
	"time"

	"github.com/lib/pq"
	"gitlab.sudovi.me/erp/procurements-api/data"
)

type GetContractsInputDTO struct {
	Page                *int    `json:"page" validate:"omitempty"`
	Size                *int    `json:"size" validate:"omitempty"`
	ProcurementID       *int    `json:"procurement_id"`
	SupplierID          *int    `json:"supplier_id"`
	Year                *string `json:"year"`
	SortByDateOfExpiry  *string `json:"sort_by_date_of_expiry"`
	SortByDateOfSigning *string `json:"sort_by_date_of_signing"`
	SortByGrossValue    *string `json:"sort_by_gross_value"`
	SortBySupplierID    *string `json:"sort_by_supplier_id"`
	SortBySerialNumber  *string `json:"sort_by_serial_number"`
}

type ContractDTO struct {
	PublicProcurementID int           `json:"public_procurement_id" validate:"required"`
	SupplierID          int           `json:"supplier_id"  validate:"required"`
	SerialNumber        string        `json:"serial_number"  validate:"required"`
	DateOfSigning       time.Time     `json:"date_of_signing"  validate:"required"`
	DateOfExpiry        *time.Time    `json:"date_of_expiry"`
	NetValue            *float32      `json:"net_value"`
	GrossValue          *float32      `json:"gross_value"`
	VatValue            *float32      `json:"vat_value"`
	File                pq.Int64Array `json:"file"`
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
	File                []int      `json:"file"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (dto ContractDTO) ToContract() *data.Contract {
	var netValue, grossValue, vatValue *int
	if dto.NetValue != nil {
		netValueInt := int(math.Round(float64(*dto.NetValue) * 100))
		netValue = &netValueInt
	}

	if dto.GrossValue != nil {
		grossValueInt := int(math.Round(float64(*dto.GrossValue) * 100))
		grossValue = &grossValueInt
	}
	if dto.VatValue != nil {
		vatValueInt := int(math.Round(float64(*dto.VatValue) * 100))
		vatValue = &vatValueInt
	}
	return &data.Contract{
		PublicProcurementID: dto.PublicProcurementID,
		SupplierID:          dto.SupplierID,
		SerialNumber:        dto.SerialNumber,
		DateOfSigning:       dto.DateOfSigning,
		DateOfExpiry:        dto.DateOfExpiry,
		NetValue:            netValue,
		GrossValue:          grossValue,
		VatValue:            vatValue,
		File:                dto.File,
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

	array := make([]int, len(data.File))
	for i, id := range data.File {
		array[i] = int(id)
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
		File:                array,
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
