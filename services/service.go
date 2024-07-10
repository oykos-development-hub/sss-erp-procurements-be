package services

import (
	"context"

	"gitlab.sudovi.me/erp/procurements-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type PlanService interface {
	CreatePlan(ctx context.Context, input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	UpdatePlan(ctx context.Context, id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	DeletePlan(ctx context.Context, id int) error
	GetPlan(id int) (*dto.PlanResponseDTO, error)
	GetPlanList(input dto.GetPlansInputDTO) ([]dto.PlanResponseDTO, *uint64, error)
}

type ItemService interface {
	CreateItem(ctx context.Context, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	UpdateItem(ctx context.Context, id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	DeleteItem(ctx context.Context, id int) error
	GetItem(id int) (*dto.ItemResponseDTO, error)
	GetItemList(input dto.GetItemsInputDTO) ([]dto.ItemResponseDTO, *uint64, error)
}

type ArticleService interface {
	CreateArticle(input dto.ArticleDTO) (*dto.ArticleResponseDTO, error)
	UpdateArticle(id int, input dto.ArticleDTO) (*dto.ArticleResponseDTO, error)
	DeleteArticle(id int) error
	GetArticle(id int) (*dto.ArticleResponseDTO, error)
	GetArticleList(input *dto.GetArticleListInput) ([]dto.ArticleResponseDTO, error)
}

type ContractService interface {
	CreateContract(ctx context.Context, input dto.ContractDTO) (*dto.ContractResponseDTO, error)
	UpdateContract(ctx context.Context, id int, input dto.ContractDTO) (*dto.ContractResponseDTO, error)
	DeleteContract(ctx context.Context, id int) error
	GetContract(id int) (*dto.ContractResponseDTO, error)
	GetContractList(input dto.GetContractsInputDTO) ([]dto.ContractResponseDTO, *uint64, error)
}

type OrganizationUnitArticleService interface {
	CreateOrganizationUnitArticle(input dto.OrganizationUnitArticleDTO) (*dto.OrganizationUnitArticleResponseDTO, error)
	UpdateOrganizationUnitArticle(id int, input dto.OrganizationUnitArticleDTO) (*dto.OrganizationUnitArticleResponseDTO, error)
	DeleteOrganizationUnitArticle(id int) error
	GetOrganizationUnitArticle(id int) (*dto.OrganizationUnitArticleResponseDTO, error)
	GetOrganizationUnitArticleList(input dto.GetOrganizationUnitArticleListInputDTO) ([]dto.OrganizationUnitArticleResponseDTO, error)
}

type OrganizationUnitPlanLimitService interface {
	CreateOrganizationUnitPlanLimit(input dto.OrganizationUnitPlanLimitDTO) (*dto.OrganizationUnitPlanLimitResponseDTO, error)
	UpdateOrganizationUnitPlanLimit(id int, input dto.OrganizationUnitPlanLimitDTO) (*dto.OrganizationUnitPlanLimitResponseDTO, error)
	DeleteOrganizationUnitPlanLimit(id int) error
	GetOrganizationUnitPlanLimit(id int) (*dto.OrganizationUnitPlanLimitResponseDTO, error)
	GetOrganizationUnitPlanLimitList(input dto.OrganizationUnitPlanLimitInputDTO) ([]dto.OrganizationUnitPlanLimitResponseDTO, error)
}

type ContractArticleService interface {
	CreateContractArticle(input dto.ContractArticleDTO) (*dto.ContractArticleResponseDTO, error)
	UpdateContractArticle(id int, input dto.ContractArticleDTO) (*dto.ContractArticleResponseDTO, error)
	DeleteContractArticle(id int) error
	GetContractArticle(id int) (*dto.ContractArticleResponseDTO, error)
	GetContractArticleList(input *dto.GetContractArticlesInputDTO) ([]dto.ContractArticleResponseDTO, *uint64, error)
}

type ContractArticleOverageService interface {
	CreateContractArticleOverage(input dto.ContractArticleOverageDTO) (*dto.ContractArticleOverageResponseDTO, error)
	UpdateContractArticleOverage(id int, input dto.ContractArticleOverageDTO) (*dto.ContractArticleOverageResponseDTO, error)
	DeleteContractArticleOverage(id int) error
	GetContractArticleOverage(id int) (*dto.ContractArticleOverageResponseDTO, error)
	GetContractArticleOverageList(input dto.GetContractArticleOverageInputDTO) ([]dto.ContractArticleOverageResponseDTO, error)
}

type LogService interface {
	CreateLog(input dto.LogDTO) (*dto.LogResponseDTO, error)
	UpdateLog(id int, input dto.LogDTO) (*dto.LogResponseDTO, error)
	DeleteLog(id int) error
	GetLog(id int) (*dto.LogResponseDTO, error)
	GetLogList(filter dto.LogFilterDTO) ([]dto.LogResponseDTO, *uint64, error)
}

type ErrorLogService interface {
	CreateErrorLog(err error)
	UpdateErrorLog(id int, input dto.ErrorLogDTO) (*dto.ErrorLogResponseDTO, error)
	DeleteErrorLog(id int) error
	GetErrorLog(id int) (*dto.ErrorLogResponseDTO, error)
	GetErrorLogList(filter dto.ErrorLogFilterDTO) ([]dto.ErrorLogResponseDTO, *uint64, error)
}
