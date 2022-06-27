package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	password string
}

func TestHashPassword(t *testing.T) {
	password := "123456"
	result, _ := HashPassword(password)
	assert.NotEqual(t, password, result, "Expected hash to be created")
}

func TestHashTable(t *testing.T) {
	var testTb = []testTable{
		{
			password: "qwertyuiop",
		},
		{
			password: "123456",
		},
	}

	for _, val := range testTb {
		t.Run(val.password, func(t *testing.T) {
			result, _ := HashPassword(val.password)
			assert.NotEqual(t, val.password, result, "Expected hash to be created")
		})
	}
}

func TestCheckPassword(t *testing.T) {
	password := "123456"
	hash, _ := HashPassword(password)
	result := CheckPassword(hash, password)
	assert.Equal(t, true, result, "Expected password to be correct")
}
