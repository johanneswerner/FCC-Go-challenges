package c147

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{2024, true},
		{2023, false},
		{2100, false},
		{2000, true},
		{1999, false},
		{2040, true},
		{2026, false},
	}
	for _, tt := range tests {
		if got := IsLeapYear(tt.year); got != tt.want {
			t.Errorf("IsLeapYear(%d) = %t, want %t", tt.year, got, tt.want)
		}
	}
}
