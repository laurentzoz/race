package fast

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Errorf("Expected 3 but got %d", res)
	}
}

func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			res := Add(tt.a, tt.b)
			if res != tt.want {
				t.Errorf("Expected %d but got %d", tt.want, res)
			}
		})
	}
}
