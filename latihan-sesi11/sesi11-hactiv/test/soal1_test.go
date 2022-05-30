package test

import (
	"fmt"
	"testing"
	"unit-test/helpers"

	"github.com/stretchr/testify/assert"
)

func TestPria(t *testing.T) {
	var reqData = []struct {
		age      int
		gender   string
		expected bool
	}{
		{
			age:      16,
			gender:   "pria",
			expected: false,
		},
		{
			age:      17,
			gender:   "pria",
			expected: false,
		},
		{
			age:      50,
			gender:   "pria",
			expected: true,
		},
		{
			age:      60,
			gender:   "pria",
			expected: false,
		},
	}

	for i, rq := range reqData {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, rq.expected, helpers.IsAllowedToWork(rq.age, rq.gender), "Not Allowed to Work")
		})
	}

}

func TestWanita(t *testing.T) {
	var reqData = []struct {
		age      int
		gender   string
		expected bool
	}{
		{
			age:      16,
			gender:   "wanita",
			expected: false,
		},
		{
			age:      21,
			gender:   "wanita",
			expected: true,
		},
		{
			age:      50,
			gender:   "wanita",
			expected: true,
		},
		{
			age:      60,
			gender:   "wanita",
			expected: false,
		},
	}

	for i, rq := range reqData {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, rq.expected, helpers.IsAllowedToWork(rq.age, rq.gender), "Not Allowed to Work")
		})
	}

}
