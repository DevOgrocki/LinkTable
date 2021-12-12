package internal

import (
	"fmt"
)

type Key struct {
	dataset dataSet
	id string
}

const separator = "|"

func (k Key) getIndex() string {
	unHashedKey := fmt.Sprintf("%s%s%s", k.dataset.name, separator, k.id)

	// hash key
	hashedKey := unHashedKey

	return hashedKey
}

