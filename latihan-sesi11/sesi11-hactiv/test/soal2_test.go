package test

import (
	"fmt"
	"testing"
	"unit-test/helpers"

	"github.com/stretchr/testify/assert"
)

func TestFight(t *testing.T) {
	var reqData = []struct {
		enemyPower   []int
		powerGained  []int
		initialPower int
		expected     int
	}{
		{
			enemyPower:   []int{5, 3, 9, 8},
			powerGained:  []int{2, 2, 3, 1},
			initialPower: 3,
			expected:     7,
		},
		{
			enemyPower:   []int{2, 6, 3, 9},
			powerGained:  []int{2, 2, 3, 5},
			initialPower: 2,
			expected:     14,
		},
	}

	for i, rq := range reqData {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, rq.expected, helpers.Fight(rq.enemyPower, rq.powerGained, rq.initialPower), "final power is not correct")
		})
	}
}
