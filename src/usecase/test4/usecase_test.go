package test4_test

import (
	"back-end/src/usecase/test4"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountSmiles(t *testing.T) {

	c := test4.CountSmiles([]string{":)", ";(", ";}", ":-D"})
	fmt.Println("c : ", c)
	testCase := []struct {
		name string
        smiles []string
        expected int
	}{
		{
			name: "Case 1", 
			smiles: []string{":)", ";(", ";}", ":-D"},
			expected: 2,
		},
        {
			name: "Case 2",
			smiles: []string{";D", ":-(", ":-)", ";~)"},
			expected: 3,
		},
		{
			name: "Case 3",
			smiles: []string{";]", ":[", ";*", ":$", ";-D"},
			expected: 1,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
            actual := test4.CountSmiles(tc.smiles)
            assert.Equal(t, actual, tc.expected)
        })
	}

}