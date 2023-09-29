package main

import (
	"fmt"
)

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	prod = "product"
)

func main() {
	var command string

	UserList := []string{"User1_password1", "User2_password2"}
	ProductList := make([]string, 0, 10)

	for command != exit {
		fmt.Println("Выберите действие из списка") //сделать красивый вывод, вывести список команд на этом шаге

		fmt.Println("Exit = Выход")
		fmt.Println("Reg = Регистрация")
		fmt.Println("Auth = Вход в аккаунт")
		fmt.Println("Product = Перейти в продукты")
		fmt.Scan(&command)

		switch command {
		case exit:
			break

		case prod:
			var i int64
			fmt.Println("Выберите продукты")
			fmt.Println("1. Яблоки")
			fmt.Println("2. Персики")
			fmt.Println("3. Груши")

			fmt.Scan(&i)

			switch i {
			case 1:
				fmt.Println("Яблоки добавлены в корзину")
			case 2:
				fmt.Println("Персики добавлены в корзину")
			case 3:
				fmt.Println("Груши добавлены в корзину")
			}
			ProductList = append(ProductList, command)

			message1 := fmt.Sprint("Продукты", "добавлены в корзину")
			fmt.Println(message1)

		case reg:
			fmt.Println("Введите логин и пароль вот так login_password")
			fmt.Scan(&command) // сделать так чтобы выводил сообщение если пользователь уже есть
			UserList = append(UserList, command)
			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)

			fmt.Println(UserList)
		case auth:
			fmt.Println("Введите логин и пароль вот так: 'login_parol' ")
			fmt.Scan(&command) //

			for _, v := range UserList {
				if v == command {
					fmt.Println("Добро пожаловать в наш магазин!")
				} else {
					fmt.Println("ВЫ НЕ АВТОРИЗОВАНЫ!")
				}

			}
		}
	}
}
