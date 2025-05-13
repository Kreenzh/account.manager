package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}

}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, fmt.Errorf("reading failure: %w", err)
	}
	return data, nil
}
func (db *JsonDb) Write(content []byte) error {
	file, err := os.Create(db.filename)
	if err != nil {
		return fmt.Errorf("creation failure: %w", err)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("writing failure: %w", err)
	}

	return nil

}
