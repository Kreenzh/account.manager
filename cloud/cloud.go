package cloud

import (
	"fmt"
	"net/url"
)

type CloudDb struct {
	url string
}

func NewCloudDb(urlString string) (*CloudDb, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("INVALID_URL:%w", err)
	}

	return &CloudDb{
		url: urlString,
	}, nil

}
func (db *CloudDb) Read() ([]byte, error) {
	return nil, nil
}
func (db *CloudDb) Write(content []byte) error {
	return nil
}
