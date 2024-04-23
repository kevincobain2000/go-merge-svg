package svg

// write tests for all methods in float.go

import (
	"testing"
)

func TestParseFloat(t *testing.T) {
	t.Run("valid float", func(t *testing.T) {
		str := "10.5"
		f := ParseFloat(str)
		if f != 10.5 {
			t.Errorf("Expected float to be 10.5, got %f", f)
		}
	})

	t.Run("invalid float", func(t *testing.T) {
		str := "10.5a"
		f := ParseFloat(str)
		if f != 0 {
			t.Errorf("Expected float to be 0, got %f", f)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("a < b", func(t *testing.T) {
		a, b := 10.5, 20.5
		min := Min(a, b)
		if min != 10.5 {
			t.Errorf("Expected min to be 10.5, got %f", min)
		}
	})

	t.Run("a > b", func(t *testing.T) {
		a, b := 20.5, 10.5
		min := Min(a, b)
		if min != 10.5 {
			t.Errorf("Expected min to be 10.5, got %f", min)
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("a > b", func(t *testing.T) {
		a, b := 20.5, 10.5
		max := Max(a, b)
		if max != 20.5 {
			t.Errorf("Expected max to be 20.5, got %f", max)
		}
	})

	t.Run("a < b", func(t *testing.T) {
		a, b := 10.5, 20.5
		max := Max(a, b)
		if max != 20.5 {
			t.Errorf("Expected max to be 20.5, got %f", max)
		}
	})
}
