package contract

import (
	"einfach-msg/model"
)

type Collection interface {
	Insert(request model.Model) error
	Get() ([]*model.Model, error)
}
