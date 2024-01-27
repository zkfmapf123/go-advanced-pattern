package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_A(t *testing.T) {
	a, b := 1, 2

	assert.Equal(t, a+b, 3)
}

func Test_Required(t *testing.T) {

	r := Required()

	assert.Equal(t, r.Name, "required")
}
