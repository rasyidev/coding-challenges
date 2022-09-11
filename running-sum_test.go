package codingchallenges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runningSum(nums ...int) []int {
	recent := 0
	result := []int{}

	for _, val := range nums {
		recent += val
		result = append(result, recent)
	}
	return result
}

func TestRunningSum(t *testing.T) {
	testCase := []struct {
		request  []int
		expected []int
	}{
		{
			request:  []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			request:  []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			request:  []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for i, tc := range testCase {
		t.Run(fmt.Sprintf("Test Case-%v", i+1), func(t *testing.T) {
			res := runningSum(tc.request...)
			assert.Equal(t, tc.expected, res)
		})
	}

}
