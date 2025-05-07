package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"tasks.go/files"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() (*Vault, error) {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  make([]Account, 0),
			UpdatedAt: time.Now(),
		}, nil
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &vault, nil
}
func (v *Vault) AddAccount(acc Account) error {
	v.Accounts = append(v.Accounts, acc)
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	err = files.WriteFile(data, "data.json")
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json: %w", err)
	}
	return file, nil
}
func (v *Vault) FindAccByUrl() (Account, error) {
	var urlToFind string
	fmt.Println("Input url to find:")
	_, err := fmt.Scanln(&urlToFind)
	if err != nil {
		return Account{}, fmt.Errorf("failed to scan account url: %w", err)
	}
	for _, acc := range v.Accounts {
		if strings.Contains(acc.Url, urlToFind) {
			return acc, nil
		}
	}

	return Account{}, fmt.Errorf("failed to find account by %s", urlToFind)

}
