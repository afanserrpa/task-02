package main

import (
	"flag"
	"fmt"

	. "github.com/afanserrpa/task-02/internal/handlers"
	. "github.com/afanserrpa/task-02/internal/models"
	. "github.com/afanserrpa/task-02/internal/storage"
	. "github.com/afanserrpa/task-02/internal/utils"
)

func main() {
	// Объявление флагов
	userId := flag.Int("user", 0, "ID пользователя для поиска")
	productId := flag.Int("item", 0, "ID товара для поиска")
	showHelp := flag.Bool("help", false, "Показать справку")

	// Парсим аргументы
	flag.Parse()

	// Создать базу товаров и добавить данные
	productStorage := CreateProductStorage()

	productsData := []Product{
		{Id: 1, Name: "Dildo", Price: 199.99},
		{Id: 2, Name: "Xbox Series X Gay Edition", Price: 499.99},
		{Id: 4, Name: "Nvidia RTX 5090 Ti Super Ultra", Price: 1999.99},
		{Id: 7, Name: "MacBook Air M4 for 13-inch", Price: 999.99},
		{Id: 9, Name: "iPhuck 10", Price: 9999.99},
	}

	for _, p := range productsData {
		err := AddProduct(productStorage, p)
		if err != nil {
			fmt.Println("Ошибка добавления товара:", err)
		}
	}

	// Создать базу пользователей и добавить данные
	userStorage := CreateUserStorage()

	usersData := []User{
		{Id: 1, Name: "Seryoja", Birthday: CreateSimpleDate(1991, 7, 18), Active: false, ProductIds: []int{7, 9}},
		{Id: 2, Name: "Dimas", Birthday: CreateSimpleDate(1991, 7, 18), Active: true, ProductIds: []int{2, 7}},
		{Id: 3, Name: "Ange", Birthday: CreateSimpleDate(1991, 6, 28), Active: true, ProductIds: []int{4, 9}},
	}

	for _, u := range usersData {
		err := AddUser(userStorage, u)
		if err != nil {
			fmt.Println("Ошибка добавления пользователя:", err)
		}
	}

	// Обработка флагов
	switch {
	case *showHelp:
		flag.PrintDefaults()
		return
	case *userId > 0 && *productId > 0:
		HandleUserProducts(productStorage, *productId, userStorage, *userId)
	case *userId > 0:
		HandleUser(productStorage, userStorage, *userId)
	case *productId > 0:
		HandleProduct(productStorage, *productId)
	}
}
