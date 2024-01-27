package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch_1(t *testing.T) {

	v, err := 10, ""

	Match(v,
		func(v int) {
			if v > 10 {
				err = "over 10"
			}
		},
		func(v int) {
			if v < 10 {
				err = "under 10"
			}
		},
		func(v int) {
			if v == 10 {
				err = "match 10"
			}
		})

	assert.Equal(t, err, "match 10")
}

func TestMatch_2(t *testing.T) {

	v, err := 11, ""

	Match(v,
		func(v int) {
			if v > 10 {
				err = "over 10"
			}
		},
		func(v int) {
			if v < 10 {
				err = "under 10"
			}
		},
		func(v int) {
			if v == 10 {
				err = "match 10"
			}
		})

	assert.Equal(t, err, "over 10")
}

func TestMatch_3(t *testing.T) {

	v, err := 9, ""

	Match(v,
		func(v int) {
			if v > 10 {
				err = "over 10"
			}
		},
		func(v int) {
			if v < 10 {
				err = "under 10"
			}
		},
		func(v int) {
			if v == 10 {
				err = "match 10"
			}
		})

	assert.Equal(t, err, "under 10")
}
