package structs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type model struct {
	Name string `mapkey:"name"`
	Age  int    `mapkey:"age"`
}

func TestStructToMapString(t *testing.T) {
	stc := model{
		Name: "it's my name",
		Age:  17,
	}
	m, err := StructToMapString(stc)
	assert.Nil(t, err)
	assert.Equal(t, m["name"], stc.Name)
	assert.Equal(t, m["age"], strconv.Itoa(stc.Age))
}
