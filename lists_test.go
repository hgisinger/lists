package lists

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	doubles := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	if doubles[0] != 2 || doubles[1] != 4 || doubles[2] != 6 {
		t.Error("Map([]int{1, 2, 3}, func(x int) int { return x * 2 }) != []int{2, 4, 6}")
	}
	toString := Map([]int{1, 2, 3}, strconv.Itoa)
	if toString[0] != "1" || toString[1] != "2" || toString[2] != "3" {
		t.Error(`Map([]int{1, 2, 3}, strconv.Itoa) != []string{"1", "2", "3"}`)
	}
}
