package internal

import (
	"cloud.google.com/go/bigtable"
)

type BigTable struct {
	client *bigtable.Client
	table  *bigtable.Table
}

func (bt BigTable) ReadRow(string key) ([]byte, error) {
	row, err := bt.table.ReadRow(key)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (bt BigTable) UpdateRow(key string, value []byte) error {

}

func (bt BigTable) CreateRow(key string, value []byte) error {

}