package main

import (
	"fmt"

	"tasks.go/account"
)

func main() {
Menu:
	for {
		userChoice := mainMenu()
		switch {
		case userChoice == 1:
			err := createAccount()
			if err != nil {
				break Menu
			}
		case userChoice == 2:
			findAccount()
		case userChoice == 3:
			deleteAccount()
		case userChoice == 4:
			break Menu

		}

	}

}
func mainMenu() int {
	var userChoice int
	fmt.Println("___MENU___")
	fmt.Println("1. Create new acc")
	fmt.Println("2. Find acc")
	fmt.Println("3. Delete acc")
	fmt.Println("4. Exit")
	fmt.Scanln(&userChoice)
	return userChoice

}
func createAccount() error {
	login := promptData("Введите логин")
	password := promptData("Ваш пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}
	vault, err := account.NewVault()
	if err != nil {
		return fmt.Errorf("failed to add account to vault: %w", err)
	}

	vault.AddAccount(*myAccount)
	return nil

}
func promptData(prompt string) string {
	var res string
	fmt.Println(prompt + ": ")
	fmt.Scanln(&res)
	return res
}
func findAccount() {

}
func deleteAccount() {

}
