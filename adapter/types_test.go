package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockJson(t *testing.T) {
	b := Block{
		Transactions:[]Transaction{
			{
				Operations: []Operation{
					{},
				},
			},
		},
	}

	j, err := json.Marshal(b)
	assert.NoError(t, err)

	fmt.Println(string(j))
}
