package luhns

import (
	"testing"
)

func TestValidator(t *testing.T) {
	tables := []struct {
		input string
		valid bool
	}{
		{"79927398713", true},
		{"7992739873", false},
		{"799279873", false},
	}

	for _, table := range tables {
		valid, _ := Validator(table.input)
		if valid != table.valid {
			t.Errorf("Input : %s. Expected = %v, Got %v",
				table.input, table.valid, valid)
		}
	}
}
