package main

import (
	"reflect"
	"testing"
)

func TestStringSliceIntSlice(t *testing.T) {
	tables := []struct {
		x []string
		y []int
	}{
		{[]string{"0", "1", "2"}, []int{0, 1, 2}},
	}
	for _, table := range tables {
		result := StringSliceIntSlice(table.x)
		if !reflect.DeepEqual(result, table.y) { // need to use reflect to test composite types
			t.Errorf("Got: %d, type: %s, Expected: %d, type: %s", result, reflect.TypeOf(result).Elem(), table.y, reflect.TypeOf(table.y).Elem())
		} else {
			t.Logf("Tested passed: %d", result)
		}
	}
}

func TestStringsToSlice(t *testing.T) {
	tables := []struct {
		x string
		y []string
	}{
		{"1 1 1 0 0 0",
			[]string{
				"1", "1", "1", "0", "0", "0"},
		},
	}
	for _, table := range tables {
		result := StringsToSlice(table.x)
		if !reflect.DeepEqual(result, table.y) {
			t.Errorf("Got: %d, Expected %d", result, table.y)
		}
	}
}

func TestHourglassSum(t *testing.T) {
	tables := []struct {
		x [][]int
		y int
	}{
		{[][]int{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}, 19},
		{[][]int{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}, 28},
	}
	for _, table := range tables {
		result := HourglassSum(table.x)
		if result != table.y {
			t.Errorf("Got: %d, Expected %d", result, table.y)
		}
	}
}
