package main

import (
	"fmt"
	"tasks.go/account"
	"tasks.go/files"
)

func main() {
	v, err := createVault()
	if err != nil {
		return
	}

Menu:
	for {
		userChoice := mainMenu()
		switch {
		case userChoice == 1:
			err := createAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}

		case userChoice == 2:
			acc, err := findAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}
			acc.Output()

		case userChoice == 3:
			s, err := deleteAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}
			fmt.Println(s)
		case userChoice == 4:
			break Menu

		}
		data, _ := v.ToBytes()
		files.WriteFile(data, "data.json")

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
func createAccount(v *account.Vault) error {
	login := promptData("Введите логин")
	password := promptData("Ваш пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	err = v.AddAccount(*myAccount)
	if err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}

	return nil

}
func createVault() (*account.Vault, error) {
	Vault, err := account.NewVault()
	if err != nil {
		return &account.Vault{}, fmt.Errorf("failed to add account to vault: %w", err)
	}
	return Vault, nil
}
func promptData(prompt string) string {
	var res string
	fmt.Println(prompt + ": ")
	fmt.Scanln(&res)
	return res
}
func findAccount(v *account.Vault) (account.Account, error) {
	// scan url to find
	//method to vault to find acc using url(strings.contain)
	// output acc data (few acc?)
	acc, err := (*account.Vault).FindAccByUrl(v)
	if err != nil {
		return account.Account{}, fmt.Errorf("failed to find account: %w", err)
	}

	return acc, nil
}

func deleteAccount(v *account.Vault) (string, error) {
	//	URl
	//	method to vault to delete
	//	deleted or not found
	acc, err := (*account.Vault).FindAccByUrl(v)
	if err != nil {

		return "not found acc", fmt.Errorf("failed to find account: %w", err)
	}
	for index, allAccs := range v.Accounts {
		if allAccs == acc {
			v.Accounts = append(v.Accounts[:index], v.Accounts[index+1:]...)
			fmt.Println("deleted")
			break
		}
	}
	return "", nil
}
