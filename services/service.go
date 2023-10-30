package services

import (
	"gitlab.sudovi.me/erp/procurements-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type PlanService interface {
	CreatePlan(input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	UpdatePlan(id int, input dto.PlanDTO) (*dto.PlanResponseDTO, error)
	DeletePlan(id int) error
	GetPlan(id int) (*dto.PlanResponseDTO, error)
	GetPlanList(input dto.GetPlansInputDTO) ([]dto.PlanResponseDTO, *uint64, error)
}

type ItemService interface {
	CreateItem(input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	UpdateItem(id int, input dto.ItemDTO) (*dto.ItemResponseDTO, error)
	DeleteItem(id int) error
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
	CreateContract(input dto.ContractDTO) (*dto.ContractResponseDTO, error)
	UpdateContract(id int, input dto.ContractDTO) (*dto.ContractResponseDTO, error)
	DeleteContract(id int) error
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
