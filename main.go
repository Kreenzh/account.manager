package main

import (
	"fmt"

	"github.com/lpernett/godotenv"
	"tasks.go/account"
	"tasks.go/encrypter"
	"tasks.go/files"
	"tasks.go/output"
)

// var menu = map[string]func(*account.VaultwithDb){
// 	"1": createAccount,
// 	"2": findAccount,
// 	"3": deleteAccount,
// }

func main() {

	err := godotenv.Load("env.env")
	if err != nil {
		output.PrintErr("failed to read .env file")
	}
	db := files.NewJsonDb("data.vault")
	v, err := createVault()
	if err != nil {
		return
	}

Menu:
	for {
		userChoice := promptData([]string{
			"1. Create new acc",
			"2. Find acc",
			"3. Delete acc",
			"4. Exit",
			"Choose variant",
		})
		switch {
		case userChoice == "1":
			err := createAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}

		case userChoice == "2":
			acc, err := findAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}
			acc.Output()

		case userChoice == "3":
			s, err := deleteAccount(v)
			if err != nil {
				fmt.Print(err.Error())
				break Menu
			}
			fmt.Println(s)
		case userChoice == "4":
			break Menu

		}
		data, _ := v.ToBytes()
		db.Write(data)

	}

}

func createAccount(v *account.VaultwithDb) error {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Ваш пароль"})
	url := promptData([]string{"Введите URL"})
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
func createVault() (*account.VaultwithDb, error) {
	Vault, err := account.NewVault(files.NewJsonDb("data.json"), *encrypter.NewEncrypter())
	if err != nil {
		return &account.VaultwithDb{}, fmt.Errorf("failed to add account to vault: %w", err)
	}
	return Vault, nil
}

// get slice any type
// output elems by string
// last elem - printf
// append : to last elem
func promptData[T any](prompt []T) string {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
func findAccount(v *account.VaultwithDb) (account.Account, error) {
	// scan url to find
	//method to vault to find acc using url(strings.contain)
	// output acc data (few acc?)
	acc, err := (*account.VaultwithDb).FindAcc(v)
	if err != nil {
		return account.Account{}, fmt.Errorf("failed to find account: %w", err)
	}

	return acc, nil
}

func deleteAccount(v *account.VaultwithDb) (string, error) {
	//	URl
	//	method to vault to delete
	//	deleted or not found
	acc, err := (*account.VaultwithDb).FindAcc(v)
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
