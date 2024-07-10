package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/procurements-api/data"
	"gitlab.sudovi.me/erp/procurements-api/handlers"
	"gitlab.sudovi.me/erp/procurements-api/middleware"

	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/procurements-api/services"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/procurements-api"

	models := data.New(cel.DB.Pool)

	ErrorLogService := services.NewErrorLogServiceImpl(cel, models.ErrorLog)
	ErrorLogHandler := handlers.NewErrorLogHandler(cel, ErrorLogService)

	PlanService := services.NewPlanServiceImpl(cel, models.Plan)
	PlanHandler := handlers.NewPlanHandler(cel, PlanService, ErrorLogService)

	ItemService := services.NewItemServiceImpl(cel, models.Item)
	ItemHandler := handlers.NewItemHandler(cel, ItemService, ErrorLogService)

	ArticleService := services.NewArticleServiceImpl(cel, models.Article)
	ArticleHandler := handlers.NewArticleHandler(cel, ArticleService, ErrorLogService)

	ContractService := services.NewContractServiceImpl(cel, models.Contract)
	ContractHandler := handlers.NewContractHandler(cel, ContractService, ErrorLogService)

	OrganizationUnitArticleService := services.NewOrganizationUnitArticleServiceImpl(cel, models.OrganizationUnitArticle)
	OrganizationUnitArticleHandler := handlers.NewOrganizationUnitArticleHandler(cel, OrganizationUnitArticleService, ErrorLogService)

	OrganizationUnitPlanLimitService := services.NewOrganizationUnitPlanLimitServiceImpl(cel, models.OrganizationUnitPlanLimit)
	OrganizationUnitPlanLimitHandler := handlers.NewOrganizationUnitPlanLimitHandler(cel, OrganizationUnitPlanLimitService, ErrorLogService)

	ContractArticleService := services.NewContractArticleServiceImpl(cel, models.ContractArticle)
	ContractArticleHandler := handlers.NewContractArticleHandler(cel, ContractArticleService, ArticleService, ErrorLogService)

	ContractArticleOverageService := services.NewContractArticleOverageServiceImpl(cel, models.ContractArticleOverage)
	ContractArticleOverageHandler := handlers.NewContractArticleOverageHandler(cel, ContractArticleOverageService, ErrorLogService)

	LogService := services.NewLogServiceImpl(cel, models.Log)
	LogHandler := handlers.NewLogHandler(cel, LogService, ErrorLogService)

	myHandlers := &handlers.Handlers{
		PlanHandler:                      PlanHandler,
		ItemHandler:                      ItemHandler,
		ArticleHandler:                   ArticleHandler,
		OrganizationUnitArticleHandler:   OrganizationUnitArticleHandler,
		ContractHandler:                  ContractHandler,
		OrganizationUnitPlanLimitHandler: OrganizationUnitPlanLimitHandler,
		ContractArticleHandler:           ContractArticleHandler,
		ContractArticleOverageHandler:    ContractArticleOverageHandler,
		LogHandler:                       LogHandler,
		ErrorLogHandler:                  ErrorLogHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
