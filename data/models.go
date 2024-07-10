package data

import (
	"fmt"

	db2 "github.com/upper/db/v4"
	up "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"

	"database/sql"
	"os"
)

//nolint:all
var db *sql.DB

//nolint:all
var Upper db2.Session

type Models struct {
	// any models inserted here (and in the New function)
	// are easily accessible throughout the entire application
	Plan                      Plan
	Item                      Item
	Article                   Article
	Contract                  Contract
	OrganizationUnitArticle   OrganizationUnitArticle
	OrganizationUnitPlanLimit OrganizationUnitPlanLimit
	ContractArticle           ContractArticle
	ContractArticleOverage    ContractArticleOverage
	Log                       Log
	ErrorLog ErrorLog
	}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		Upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		Upper, _ = postgresql.New(databasePool)
	default:
		// do nothing
	}

	return Models{
		Plan:                      Plan{},
		Item:                      Item{},
		Article:                   Article{},
		Contract:                  Contract{},
		OrganizationUnitArticle:   OrganizationUnitArticle{},
		OrganizationUnitPlanLimit: OrganizationUnitPlanLimit{},
		ContractArticle:           ContractArticle{},
		ContractArticleOverage:    ContractArticleOverage{},
		Log:                       Log{},
		ErrorLog: ErrorLog{},
	}
}

func getInsertId(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}

func paginateResult(res up.Result, page int, pageSize int) up.Result {
	// Calculate the offset based on the page number and page size
	offset := (page - 1) * pageSize

	// Apply pagination to the query
	res = res.Offset(offset).Limit(pageSize)

	return res
}
