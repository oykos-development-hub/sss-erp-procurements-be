package main

import (
	"gitlab.sudovi.me/erp/procurements-api/handlers"
	"gitlab.sudovi.me/erp/procurements-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/oykos-development-hub/celeritas"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	r := chi.NewRouter()

	// Konfigurišite CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)

	r.Route("/api", func(rt chi.Router) {
		rt.Post("/plans", handlers.PlanHandler.CreatePlan)
		rt.Get("/plans/{id}", handlers.PlanHandler.GetPlanById)
		rt.Get("/plans", handlers.PlanHandler.GetPlanList)
		rt.Put("/plans/{id}", handlers.PlanHandler.UpdatePlan)
		rt.Delete("/plans/{id}", handlers.PlanHandler.DeletePlan)

		rt.Post("/items", handlers.ItemHandler.CreateItem)
		rt.Get("/items/{id}", handlers.ItemHandler.GetItemById)
		rt.Get("/items", handlers.ItemHandler.GetItemList)
		rt.Put("/items/{id}", handlers.ItemHandler.UpdateItem)
		rt.Delete("/items/{id}", handlers.ItemHandler.DeleteItem)

		rt.Post("/articles", handlers.ArticleHandler.CreateArticle)
		rt.Get("/articles/{id}", handlers.ArticleHandler.GetArticleById)
		rt.Get("/articles", handlers.ArticleHandler.GetArticleList)
		rt.Put("/articles/{id}", handlers.ArticleHandler.UpdateArticle)
		rt.Delete("/articles/{id}", handlers.ArticleHandler.DeleteArticle)

		rt.Post("/contracts", handlers.ContractHandler.CreateContract)
		rt.Get("/contracts/{id}", handlers.ContractHandler.GetContractById)
		rt.Get("/contracts", handlers.ContractHandler.GetContractList)
		rt.Put("/contracts/{id}", handlers.ContractHandler.UpdateContract)
		rt.Delete("/contracts/{id}", handlers.ContractHandler.DeleteContract)

		rt.Post("/organization-unit-articles", handlers.OrganizationUnitArticleHandler.CreateOrganizationUnitArticle)
		rt.Get("/organization-unit-articles/{id}", handlers.OrganizationUnitArticleHandler.GetOrganizationUnitArticleById)
		rt.Get("/organization-unit-articles", handlers.OrganizationUnitArticleHandler.GetOrganizationUnitArticleList)
		rt.Put("/organization-unit-articles/{id}", handlers.OrganizationUnitArticleHandler.UpdateOrganizationUnitArticle)
		rt.Delete("/organization-unit-articles/{id}", handlers.OrganizationUnitArticleHandler.DeleteOrganizationUnitArticle)

		rt.Post("/organization-unit-plan-limits", handlers.OrganizationUnitPlanLimitHandler.CreateOrganizationUnitPlanLimit)
		rt.Get("/organization-unit-plan-limits/{id}", handlers.OrganizationUnitPlanLimitHandler.GetOrganizationUnitPlanLimitById)
		rt.Get("/organization-unit-plan-limits", handlers.OrganizationUnitPlanLimitHandler.GetOrganizationUnitPlanLimitList)
		rt.Put("/organization-unit-plan-limits/{id}", handlers.OrganizationUnitPlanLimitHandler.UpdateOrganizationUnitPlanLimit)
		rt.Delete("/organization-unit-plan-limits/{id}", handlers.OrganizationUnitPlanLimitHandler.DeleteOrganizationUnitPlanLimit)

		rt.Post("/contract-articles", handlers.ContractArticleHandler.CreateContractArticle)
		rt.Get("/contract-articles/{id}", handlers.ContractArticleHandler.GetContractArticleById)
		rt.Get("/contract-articles", handlers.ContractArticleHandler.GetContractArticleList)
		rt.Put("/contract-articles/{id}", handlers.ContractArticleHandler.UpdateContractArticle)
		rt.Delete("/contract-articles/{id}", handlers.ContractArticleHandler.DeleteContractArticle)
		rt.Post("/read-template-articles", handlers.ContractArticleHandler.ReadTemplate)

		rt.Post("/contract-article-overages", handlers.ContractArticleOverageHandler.CreateContractArticleOverage)
		rt.Get("/contract-article-overages/{id}", handlers.ContractArticleOverageHandler.GetContractArticleOverageById)
		rt.Get("/contract-article-overages", handlers.ContractArticleOverageHandler.GetContractArticleOverageList)
		rt.Put("/contract-article-overages/{id}", handlers.ContractArticleOverageHandler.UpdateContractArticleOverage)
		rt.Delete("/contract-article-overages/{id}", handlers.ContractArticleOverageHandler.DeleteContractArticleOverage)
	})

	return r
}
