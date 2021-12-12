package internal

import (
	"encoding/json"
)

const currentVersion = "v1"

type dataNode struct {
	key string
	payload string
	previousID string
	nextID string
	dataSet dataSet
}

func (dn dataNode) Get() string {
	return dn.payload
}

func (dn dataNode) Send(bt BigTable) error {
	row, err := dn.toEntry().serialize()
	if err != nil {
		return err
	}

	err := bt.CreateRow(row)
	return err
}

func (dn dataNode) toEntry() entry {
	return entry{
		Version:    currentVersion,
		PreviousID: dn.previousID,
		NextID:     dn.nextID,
		Payload:    dn.payload,
	}
}

func GetDataNode(bt BigTable, id string, ds dataSet) (*dataNode, error) {
	key := Key{dataset: ds, id: id}

	// read from table
	row, err := bt.ReadRow(key.getIndex())
	if err != nil {
		return nil, err
	}

	// create dataNode
	entry, err := deserialize(row)
	if err != nil {
		return nil, err
	}

	dataNode := dataNode{
		key:        id,
		payload:    entry.Payload,
		previousID: entry.PreviousID,
		nextID:     entry.NextID,
		dataSet:    ds,
	}

	return &dataNode, nil
}

func CreateDataNode(id string, ds dataSet, pre dataNode, next dataNode, pl string) *dataNode {
	key := Key{
		dataset: ds,
		id:      id,
	}

	dataNode := dataNode{
		key:        key.getIndex(),
		payload:    pl,
		previousID: pre.key,
		nextID:     next.key,
		dataSet:    ds,
	}

	return &dataNode
}


type entry struct {
	Version    string `json:"version"`
	PreviousID string `json:"previousID"`
	NextID     string `json:"nextID"`
	Payload    string `json:"Payload"`
}

func (e entry) serialize() ([]byte, error) {
	return json.Marshal(e)
}

func deserialize(bs []byte) (*entry, error) {
	var e entry
	err := json.Unmarshal(bs, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}