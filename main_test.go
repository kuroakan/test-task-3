package main

import (
	"fmt"
	"testing"
)

func TestFillUpLevel(t *testing.T) {
	requireNilError := func(t *testing.T, err error) {
		if err != nil {
			t.Fatalf("expected nil\ngot: %v", err)
		}
	}

	testCases := []struct {
		name     string
		info     []byte
		expected func(t *testing.T, err error)
	}{
		{
			name:     "Correct Float",
			info:     []byte{1, 65, 40, 0, 0},
			expected: requireNilError,
		},
		{
			name:     "Correct Int",
			info:     []byte{0, 0, 0, 1, 1},
			expected: requireNilError,
		},
		{
			name: "Incorrect Incoming Data",
			info: []byte("1Hello Worldqwert"),
			expected: func(t *testing.T, err error) {
				if err == nil {
					t.Fatal("expected error\ngot: nil")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fl, err := FillUpLevel(tc.info)
			tc.expected(t, err)
			fmt.Println(fl)
		})
	}

}
