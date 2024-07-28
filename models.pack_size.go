package main

import (
	"sort"
)

// assuming these are in descending order
var packSizesDefaultData = []int{
	5000,
	2000,
	1000,
	500,
	250,
}

var packSizesData = packSizesDefaultData

func packSizes() []int {
	return packSizesData
}

func setPackSizes(newPackSizes []int) bool {
	if len(newPackSizes) == 0 {
		return false
	}
	// array must be in descending order for pack info logic to work
	sort.Sort(sort.Reverse(sort.IntSlice(newPackSizes)))
	packSizesData = newPackSizes
	return true
}

func setPackSizesToDefault() {
	packSizesData = packSizesDefaultData
}
