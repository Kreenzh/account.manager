package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"tasks.go/encrypter"
	"tasks.go/files"
	"tasks.go/output"
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
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) (*VaultwithDb, error) {
	file, err := db.Read()
	if err != nil {
		return &VaultwithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
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
			db:  db,
			enc: enc,
		}, fmt.Errorf("%w", err)
	}
	return &VaultwithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}, nil
}

func (v *VaultwithDb) AddAccount(acc Account) error {
	db := files.NewJsonDb("data.vault")
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
func (v *VaultwithDb) FindAcc() (Account, error) {
	var smthToFind string
	fmt.Println("Input login or url to find:")
	_, err := fmt.Scanln(&smthToFind)
	if err != nil {
		return Account{}, fmt.Errorf("failed to scan account url: %w", err)
	}
	for _, acc := range v.Accounts {
		if strings.Contains(acc.Url, smthToFind) {
			return acc, nil
		}
	}

	return Account{}, fmt.Errorf("failed to find account by %s", smthToFind)

}
func (v *VaultwithDb) save() {
	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToBytes()
	encData := v.enc.Encrypt(data)
	if err != nil {
		output.PrintErr(err)
	}

	v.db.Write(encData)
}
