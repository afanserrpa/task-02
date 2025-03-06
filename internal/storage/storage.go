// Работа с хранилищами данных
package storage

import (
	"fmt"

	. "github.com/afanserrpa/task-02/internal/models"
)

type ProductStorage struct {
	Products map[int]Product
}

func CreateProductStorage() *ProductStorage {
	return &ProductStorage{
		Products: make(map[int]Product),
	}
}

func AddProduct(productStorage *ProductStorage, product Product) error {
	_, exists := productStorage.Products[product.Id]
	if exists {
		return fmt.Errorf("товар с ID \"%v\" уже существует", product.Id)
	}

	productStorage.Products[product.Id] = product
	return nil
}

func GetProduct(productStorage *ProductStorage, id int) (Product, error) {
	product, exists := productStorage.Products[id]
	if !exists {
		return Product{}, fmt.Errorf("товар с ID \"%v\" не найден", id)
	}

	return product, nil
}

type UserStorage struct {
	Users map[int]User
}

func CreateUserStorage() *UserStorage {
	return &UserStorage{
		Users: make(map[int]User),
	}
}

func AddUser(userStorage *UserStorage, user User) error {
	_, exists := userStorage.Users[user.Id]
	if exists {
		return fmt.Errorf("пользователь с ID \"%v\" уже существует", user.Id)
	}

	userStorage.Users[user.Id] = user
	return nil
}

func GetUser(userStorage *UserStorage, id int) (User, error) {
	user, exists := userStorage.Users[id]
	if !exists {
		return User{}, fmt.Errorf("пользователь с ID \"%v\" не найден", id)
	}

	if !user.Active {
		return User{}, fmt.Errorf("пользователь с ID \"%v\" не активен", id)
	}

	return user, nil
}

func GetUserProducts(userStorage *UserStorage, userId int, productStorage *ProductStorage) ([]Product, error) {
	user, err := GetUser(userStorage, userId)
	if err != nil {
		return nil, err
	}

	var products []Product

	for _, pid := range user.ProductIds {
		product, err := GetProduct(productStorage, pid)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
