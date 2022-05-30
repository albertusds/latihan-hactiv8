package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []int{10, 20, 30}

	totalExpected := 60

	result := Sum(nums...)
	if totalExpected != result {
		t.Fail()
		t.Error("Error total not suitable")
	}
}

func TestSumByZero(t *testing.T) {
	nums := []int{0, 0, 0}
	expected := 0
	result := Sum(nums...)
	if expected != result {
		t.Fail()
		t.Error("Error")
	}
}

func TestSumTestify(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := 67

	require.Equal(t, expected, Sum(nums...), "result not equal")
}

func TestSumTableTes(t *testing.T) {

	var nums = []struct {
		req []int
		res int
	}{
		{
			req: []int{10, 20, 30},
			res: 60,
		},
		{
			req: []int{10, 30, 30},
			res: 71,
		},
	}

	// nums := []int{1, 2, 3}
	// expected := 67

	for _, v := range nums {
		assert.Equal(t, v.res, Sum(v.req...), "result not equal")
	}
}
