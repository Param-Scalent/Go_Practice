package something_test

import (
	"testing"
)

func TestSomething(t *testing.T) {
	testTable := []struct {
		name string
		// input, output
	}{
		{
			name: "happy happy",
			// input output
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			// call function or method under test
			// ...
			// check response
		})
	}
}
