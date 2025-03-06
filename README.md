# Установка
### 1. Клонировать репозиторий
```sh
git clone https://https://github.com/afanserrpa/task-02.git
```
### 2. Собрать бинарник
```sh
cd ./task-02/bin
go build -o myapp.exe ./bin
```

# Использование
### Показать справку
```sh
./myapp.exe -help
```
### Поиск товара
```sh
./myapp.exe -item=7
```
### Поиск пользователя
```sh
./myapp.exe -user=2
```
### Поиск товара у пользователя
```sh
./myapp.exe -user=2 -item=7
```
