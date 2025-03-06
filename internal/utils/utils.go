// Общие утилиты для работы с данными
package utils

import (
	"strings"
	"time"

	. "github.com/afanserrpa/task-02/internal/models"
)

func CreateSimpleDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func JoinProductNames(products []Product) string {
	var names []string
	for _, p := range products {
		names = append(names, p.Name)
	}

	return strings.Join(names, ", ")
}

func CalculateProductSum(products []Product) float64 {
	var sum float64
	for _, p := range products {
		sum += p.Price
	}

	return sum
}

func IsProductPurchased(products []Product, productId int) bool {
	for _, p := range products {
		if productId == p.Id {
			return true
		}
	}

	return false
}
