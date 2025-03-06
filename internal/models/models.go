// Структуры данных
package models

import (
	"time"
)

type Product struct {
	Id    int
	Name  string
	Price float64
}

type User struct {
	Id         int
	Name       string
	Birthday   time.Time
	Active     bool
	ProductIds []int
}
