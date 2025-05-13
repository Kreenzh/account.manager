package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"tasks.go/files"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write([]byte) error
}
type Db interface {
	ByteReader
	Writer
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type VaultwithDb struct {
	Vault
	db Db
}

func NewVault(db Db) (*VaultwithDb, error) {
	file, err := db.Read()
	if err != nil {
		return &VaultwithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}, fmt.Errorf("%w", err)
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return &VaultwithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}, fmt.Errorf("%w", err)
	}
	return &VaultwithDb{
		Vault: vault,
		db:    db,
	}, nil
}

func (v *VaultwithDb) AddAccount(acc Account) error {
	db := files.NewJsonDb("data.json")
	v.Accounts = append(v.Accounts, acc)
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	err = db.Write(data)
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
func (v *VaultwithDb) FindAccByUrl() (Account, error) {
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
