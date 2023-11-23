package handlers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"gitlab.sudovi.me/erp/procurements-api/dto"
	"gitlab.sudovi.me/erp/procurements-api/errors"
	"gitlab.sudovi.me/erp/procurements-api/services"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// ContractArticleHandler is a concrete type that implements ContractHandler
type contractArticleHandlerImpl struct {
	App            *celeritas.Celeritas
	service        services.ContractArticleService
	articleService services.ArticleService
}

// NewContractArticleHandler initializes a new ContractArticleHandler with its dependencies
func NewContractArticleHandler(app *celeritas.Celeritas, contractArticleService services.ContractArticleService, articleService services.ArticleService) ContractArticleHandler {
	return &contractArticleHandlerImpl{
		App:            app,
		service:        contractArticleService,
		articleService: articleService,
	}
}

func (h *contractArticleHandlerImpl) CreateContractArticle(w http.ResponseWriter, r *http.Request) {
	var input dto.ContractArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateContractArticle(input)
	if err != nil {
		h.App.ErrorLog.Printf("Error creating contract article: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticle created successfuly", res)
}

func (h *contractArticleHandlerImpl) UpdateContractArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.ContractArticleDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateContractArticle(id, input)
	if err != nil {
		h.App.ErrorLog.Printf("Error updating contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "ContractArticle updated successfuly", res)
}

func (h *contractArticleHandlerImpl) DeleteContractArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteContractArticle(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error deleting contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "ContractArticle deleted successfuly")
}

func (h *contractArticleHandlerImpl) GetContractArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetContractArticle(id)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract article with ID %d: %v", id, err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *contractArticleHandlerImpl) GetContractArticleList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetContractArticlesInputDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetContractArticleList(&input)
	if err != nil {
		h.App.ErrorLog.Printf("Error fetching contract article list: %v", err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}

func (h *contractArticleHandlerImpl) ReadTemplate(w http.ResponseWriter, r *http.Request) {
	maxFileSize := int64(100 * 1024 * 1024) // file maximum 100 MB

	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "File is not valid", response)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Error during fetching file", response)
		return
	}
	defer file.Close()

	procurementID := r.FormValue("public_procurement_id")

	publicProcurementID, err := strconv.Atoi(procurementID)

	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "You must provide a valid public_procurement_id", response)
		return
	}

	contractid := r.FormValue("contract_id")

	contractID, err := strconv.Atoi(contractid)

	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "You must provide a valid public_procurement_id", response)
		return
	}

	// Save the file to disk
	tempFile, err := os.CreateTemp("", "uploaded-file-")
	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusInternalServerError, "Error during opening file", response)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusInternalServerError, "Error during reading file", response)
		return
	}

	// Now you can open the saved file using its path
	xlsFile, err := excelize.OpenFile(tempFile.Name())

	if err != nil {
		response := dto.ArticleResponse{
			Status: "failed",
		}
		_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Error during opening file", response)
		return
	}

	// Iterating through the Excel sheet and reading data
	var articles []dto.ContractArticleResponseDTO

	// Accessing sheets in the Excel file
	sheetMap := xlsFile.GetSheetMap()

	for _, sheetName := range sheetMap {
		rows, err := xlsFile.Rows(sheetName)
		if err != nil {
			response := dto.ArticleResponse{
				Status: "failed",
			}
			_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Error during reading file rows!", response)
			return
		}

		rowindex := 0

		for rows.Next() {
			if rowindex == 0 {
				rowindex++
				continue
			}

			cols := rows.Columns()
			if err != nil {
				response := dto.ArticleResponse{
					Status: "failed",
				}
				_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Error during reading column value", response)
				return
			}

			var article dto.ContractArticleResponseDTO
			var title, description string
			var price float32
			for cellIndex, cellValue := range cols {
				value := cellValue
				switch cellIndex {
				case 0:
					title = value
				case 1:
					description = value
				case 2:
					if value == "" {
						break
					}

					floatValue, err := strconv.ParseFloat(value, 32)

					if err != nil {
						response := dto.ArticleResponse{
							Status: "failed",
						}
						_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Error during converting neto price", response)
						return
					}
					price = float32(floatValue)
				}
			}

			if title == "" || description == "" || price == 0 {
				continue
			}

			input := dto.GetArticleListInput{
				Title:       &title,
				Description: &description,
				ItemID:      &publicProcurementID,
			}

			res, err := h.articleService.GetArticleList(&input)

			if err != nil {
				h.App.ErrorLog.Printf("Error fetching article: %v", err)
				_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
				return
			}

			if len(res) == 0 {
				response := dto.ArticleResponse{
					Status: "failed",
				}
				_ = h.App.WriteDataResponse(w, http.StatusBadRequest, "Artikal \""+title+"\" nije validan", response)
				return
			}

			filter := dto.GetContractArticlesInputDTO{
				ArticleID: &res[0].ID,
			}

			contractArticle, _, err := h.service.GetContractArticleList(&filter)

			if err != nil {
				h.App.ErrorLog.Printf("Error fetching contract article: %v", err)
				_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
				return
			}

			if len(contractArticle) > 0 {
				article.ID = contractArticle[0].ID
			}

			vatPercentage, _ := strconv.ParseFloat(res[0].VATPercentage, 32)
			vatFloat32 := float32(vatPercentage)
			article.ArticleID = res[0].ID
			grossValue := price + price*vatFloat32/100
			article.NetValue = &price
			article.GrossValue = &grossValue
			article.ContractID = contractID

			articles = append(articles, article)
		}
	}

	response := dto.ArticleResponse{
		Data:   articles,
		Status: "success",
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "File readed successfully", response)
}
