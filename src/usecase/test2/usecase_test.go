package test2_test

import (
	"back-end/src/usecase/test2"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPermutations(t *testing.T) {

	
	testCase := []struct {
		name string
        word string
        expected []string
	}{
		{
			name: "Case 1", 
			word: "a",
			expected: []string{"a"},
		},
        {
			name: "Case 2",
			word: "ab", 
			expected: []string{"ab", "ba"}},
		{
			name: "Case 3",
			word: "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{
			name: "Case 4", 
			word: "aabb", 
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
            actual := test2.Permutations(tc.word)
            assert.Equal(t, actual, tc.expected)
        })
	}
}