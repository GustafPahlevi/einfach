package contract

import (
	"github.com/GustafPahlevi/go-simple-svc/model"
)

// Collector is an interface that MUST comply with mongodb collection
type Collector interface {
	Insert(request model.Model) error
	Get() ([]*model.Model, error)
}
