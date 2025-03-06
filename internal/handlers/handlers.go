// Бизнес-логика
package handlers

import (
	"fmt"

	. "github.com/afanserrpa/task-02/internal/storage"
	. "github.com/afanserrpa/task-02/internal/utils"
)

func HandleUser(productStorage *ProductStorage, userStorage *UserStorage, userId int) {
	fmt.Printf("Ищем пользователя с ID: %d ...\n", userId)

	user, err := GetUser(userStorage, userId)
	if err != nil {
		panic(err)
	}

	products, err := GetUserProducts(userStorage, userId, productStorage)
	if err != nil {
		panic(err)
	}

	fmt.Printf("У пользователя %v куплены следующие товары: %v. Сумма купленных товаров: %v", user.Name, JoinProductNames(products), CalculateProductSum(products))
}

func HandleProduct(productStorage *ProductStorage, productId int) {
	fmt.Printf("Ищем товар с ID: %d ...\n", productId)

	product, err := GetProduct(productStorage, productId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Название товара: %v. Цена товара: %v", product.Name, product.Price)
}

func HandleUserProducts(productStorage *ProductStorage, productId int, userStorage *UserStorage, userId int) {
	fmt.Printf("Проверяем, куплен ли пользователем с ID: %d товар с ID: %d ...\n", userId, productId)

	user, err := GetUser(userStorage, userId)
	if err != nil {
		panic(err)
	}

	product, err := GetProduct(productStorage, productId)
	if err != nil {
		panic(err)
	}

	products, err := GetUserProducts(userStorage, userId, productStorage)
	if err != nil {
		panic(err)
	}

	if IsProductPurchased(products, productId) {
		fmt.Printf("Пользователь: %v приобрел товар: %v", user.Name, product.Name)
	} else {
		fmt.Printf("Пользователь: %v не приобретал товар: %v", user.Name, product.Name)
	}
}
