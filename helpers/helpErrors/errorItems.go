package errors

import "github.com/davi-sant/househub-go/models"

func ErrorItems(err ...error) []models.ErrorItem {
	var items []models.ErrorItem
	for _, e := range err {
		items = append(items, models.ErrorItem{Erro: e.Error()})
	}

	return items
}
