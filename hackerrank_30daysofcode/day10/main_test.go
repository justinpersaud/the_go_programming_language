package main

import (
	"testing"
)

func TestConsecutiveOne(t *testing.T) {
	tables := []struct {
		x string
		y int
	}{
		{"5", 1},
		{"13", 2},
		{"6", 2},
		{"439", 3}, // this one is failing
	}
	for _, table := range tables {
		result := ConsecutiveOne(table.x)
		if result != table.y {
			t.Errorf("Got: %d, Expected %d", result, table.y)
		} else {
			t.Logf("Got: %d, Expected %d", result, table.y)
		}
	}
}
