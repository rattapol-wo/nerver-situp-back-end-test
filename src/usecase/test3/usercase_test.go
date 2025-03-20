package test3_test

import (
	"back-end/src/usecase/test3"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	testCase := []struct {
		name string
        nums []int
        expected int
	}{
		{
			name: "Case 1", 
			nums: []int{7},
			expected: 7,
		},
        {
			name: "Case 2",
			nums: []int{0},
			expected: 0,
		},
		{
			name: "Case 3",
			nums: []int{1,1,2},
			expected: 2,
		},
		{
			name: "Case 4", 
			nums: []int{0,1,0,1,0},
			expected: 0,
		},
		{
			name: "Case 5", 
			nums: []int{1,2,2,3,3,3,4,3,3,3,2,2,1},
			expected: 4,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
            actual := test3.FindOddInt(tc.nums)
            assert.Equal(t, actual, tc.expected)
        })
	}
	
}