package middleware

import (
	"gitlab.sudovi.me/erp/procurements-api/data"

	"github.com/oykos-development-hub/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
