package internal

type dataSet struct {
	name string
	head *dataNode
	tail *dataNode
}

func createDataSet(name string) dataSet {
	return dataSet{
		name: name,
		head: nil,
		tail: nil,
	}
}

func (ds dataSet) getNode(index string) dataNode {


	return
}

func (ds dataSet) setNode(bt BigTable, key string, value string) error {

}