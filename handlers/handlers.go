package handlers

import (
	"net/http"
)

type Handlers struct {
	PlanHandler                      PlanHandler
	ItemHandler                      ItemHandler
	ArticleHandler                   ArticleHandler
	ContractHandler                  ContractHandler
	OrganizationUnitArticleHandler   OrganizationUnitArticleHandler
	OrganizationUnitPlanLimitHandler OrganizationUnitPlanLimitHandler
	ContractArticleHandler           ContractArticleHandler
}

type PlanHandler interface {
	CreatePlan(w http.ResponseWriter, r *http.Request)
	UpdatePlan(w http.ResponseWriter, r *http.Request)
	DeletePlan(w http.ResponseWriter, r *http.Request)
	GetPlanById(w http.ResponseWriter, r *http.Request)
	GetPlanList(w http.ResponseWriter, r *http.Request)
}

type ItemHandler interface {
	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	GetItemById(w http.ResponseWriter, r *http.Request)
	GetItemList(w http.ResponseWriter, r *http.Request)
}

type ArticleHandler interface {
	CreateArticle(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
	GetArticleById(w http.ResponseWriter, r *http.Request)
	GetArticleList(w http.ResponseWriter, r *http.Request)
}

type ContractHandler interface {
	CreateContract(w http.ResponseWriter, r *http.Request)
	UpdateContract(w http.ResponseWriter, r *http.Request)
	DeleteContract(w http.ResponseWriter, r *http.Request)
	GetContractById(w http.ResponseWriter, r *http.Request)
	GetContractList(w http.ResponseWriter, r *http.Request)
}

type OrganizationUnitArticleHandler interface {
	CreateOrganizationUnitArticle(w http.ResponseWriter, r *http.Request)
	UpdateOrganizationUnitArticle(w http.ResponseWriter, r *http.Request)
	DeleteOrganizationUnitArticle(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitArticleById(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitArticleList(w http.ResponseWriter, r *http.Request)
}

type OrganizationUnitPlanLimitHandler interface {
	CreateOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request)
	UpdateOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request)
	DeleteOrganizationUnitPlanLimit(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitPlanLimitById(w http.ResponseWriter, r *http.Request)
	GetOrganizationUnitPlanLimitList(w http.ResponseWriter, r *http.Request)
}

type ContractArticleHandler interface {
	CreateContractArticle(w http.ResponseWriter, r *http.Request)
	UpdateContractArticle(w http.ResponseWriter, r *http.Request)
	DeleteContractArticle(w http.ResponseWriter, r *http.Request)
	GetContractArticleById(w http.ResponseWriter, r *http.Request)
	GetContractArticleList(w http.ResponseWriter, r *http.Request)
}
