package lists

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	doubles := Map(func(x int) int { return x * 2 }, []int{1, 2, 3})
	if doubles[0] != 2 || doubles[1] != 4 || doubles[2] != 6 {
		t.Error("Map(func(x int) int { return x * 2 }, []int{1, 2, 3}) != []int{2, 4, 6}")
	}
	toString := Map(strconv.Itoa, []int{1, 2, 3})
	if toString[0] != "1" || toString[1] != "2" || toString[2] != "3" {
		t.Error(`Map(strconv.Itoa, []int{1, 2, 3}) != []string{"1", "2", "3"}`)
	}
}
