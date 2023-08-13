package utils

import (
	"fmt"
	"testing"
)

type ContainsTest struct {
	slice    []string
	city     string
	expected bool
}

func TestContains(t *testing.T) {
	tests := []ContainsTest{
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Istanbul", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Ankara", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Izmir", true},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "Antalya", false},
		{[]string{"Istanbul", "Ankara", "Izmir", "Manisa"}, "", false},
		{nil, "Istanbul", false},
		{nil, "", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Contains(%v, %s) should be %v", test.slice, test.city, test.expected), func(t *testing.T) {
			result := Contains(test.slice, test.city)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
