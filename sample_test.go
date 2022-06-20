package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name    string
	params  string
	returns string
}

func TestHelloName(t *testing.T) {
	result := HelloName("john doe")

	if result != "Hello john doe" {
		t.Fatal("return error")
	}
}

func TestHelloNames(t *testing.T) {
	result := HelloName("john doe")
	assert.Equal(t, "Hello john doe", result, "salah")
}

func TestHelloSubtest(t *testing.T) {
	t.Run("params jd", func(t *testing.T) {
		result := HelloName("john doe")
		assert.Equal(t, "Hello john doe", result, "salah")
	})
	t.Run("params ebi", func(t *testing.T) {
		result := HelloName("ebi")
		assert.Equal(t, "Hello ebi", result, "salah")
	})
}

func TestHelloTable(t *testing.T) {
	var testTb = []testTable{
		{
			name:    "john doe",
			params:  "john doe",
			returns: "Hello john doe",
		},
		{
			name:    "ebi",
			params:  "ebi",
			returns: "Hello ebi",
		},
	}

	for _, val := range testTb {
		t.Run(val.name, func(t *testing.T) {
			result := HelloName(val.params)
			assert.Equal(t, val.returns, result, "salah")
		})
	}
}
