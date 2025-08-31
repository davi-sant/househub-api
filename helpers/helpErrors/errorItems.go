package helpErrors

import "github.com/davi-sant/househub-go/models"

func ErrorItems(errors []error, additionalErrors ...error) []models.ErrorItem {
	var items []models.ErrorItem

	for _, e := range errors {
		if e != nil {
			items = append(items, models.ErrorItem{Erro: e.Error()})
		}
	}

	for _, e := range additionalErrors {
		if e != nil {
			items = append(items, models.ErrorItem{Erro: e.Error()})
		}
	}

	return items
}
