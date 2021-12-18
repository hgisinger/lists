package lists

import (
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	if !All(func(x int) bool { return x > 0 }, []int{1, 2, 3}) {
		t.Error("All(func(x int) bool { return x > 0 }, []int{1, 2, 3}) != true")
	}
	if All(func(x int) bool { return x > 0 }, []int{1, -2, 3}) {
		t.Error("All(func(x int) bool { return x > 0 }, []int{1, -2, 3}) != false")
	}
}

func TestAny(t *testing.T) {
	if !Any(func(x int) bool { return x > 0 }, []int{1, -2, 3}) {
		t.Error("Any(func(x int) bool { return x > 0 }, []int{1, -2, 3}) != true")
	}
	if Any(func(x int) bool { return x > 0 }, []int{-1, -2, -3}) {
		t.Error("Any(func(x int) bool { return x > 0 }, []int{-1, -2, -3}) != false")
	}
}

func TestConcat(t *testing.T) {
	result := Concat([]int{1, 2, 3}, []int{4, 5, 6})
	if result[0] != 1 ||
		result[1] != 2 ||
		result[2] != 3 ||
		result[3] != 4 ||
		result[4] != 5 ||
		result[5] != 6 {
		t.Error("Concat([]int{1, 2, 3}, []int{4, 5, 6}) != []int{1, 2, 3, 4, 5, 6}")
	}
}

func TestDelete(t *testing.T) {
	result := Delete([]int{1, 2, 3}, 2)
	if result[0] != 1 || result[1] != 3 {
		t.Error("Delete([]int{1, 2, 3}, 2) != []int{1, 3}")
	}
}

func TestDropLast(t *testing.T) {
	result, ok := DropLast([]int{1, 2, 3})
	if !ok || result[0] != 1 || result[1] != 2 {
		t.Error("DropLast([]int{1, 2, 3}) != []int{1, 2}")
	}
	result, ok = DropLast([]int{})
	if ok {
		t.Error("DropLast([]int{}) != false")
	}
}

func TestDropWhile(t *testing.T) {
	result := DropWhile(func(x int) bool { return x < 3 }, []int{1, 2, 3, 4, 5})
	if result[0] != 3 || result[1] != 4 || result[2] != 5 {
		t.Error("DropWhile(func(x int) bool { return x < 3 }, []int{1, 2, 3, 4, 5}) != []int{3, 4, 5}")
	}
}

func TestDuplicate(t *testing.T) {
	result := Duplicate(1, 3)
	if result[0] != 1 || result[1] != 1 || result[2] != 1 {
		t.Error("Duplicate([]int{1, 2, 3}) != []int{1, 1, 1}")
	}
}

func TestFilter(t *testing.T) {
	result := Filter(func(x int) bool { return x%2 != 0 }, []int{1, 2, 3, 4, 5})
	if result[0] != 1 || result[1] != 3 || result[2] != 5 {
		t.Error("Filter(func(x int) bool { return x % 2 != 0}, []int{1, 2, 3, 4, 5}) != []int{1, 3, 5}")
	}
}

func TestFilterMap(t *testing.T) {
	result := FilterMap(func(x int) (bool, int) {
		if x%2 == 0 {
			return false, x
		}
		return true, x * 2
	}, []int{1, 2, 3, 4, 5})
	if result[0] != 2 || result[1] != 6 || result[2] != 10 {
		t.Error("FilterMap(fun, []int{1, 2, 3, 4, 5}) != []int{2, 6, 10}")
	}
}

func TestFlatMap(t *testing.T) {
	result := FlatMap(func(x int) []int {
		return []int{x, x}
	}, []int{1, 2, 3})
	if result[0] != 1 || result[1] != 1 || result[2] != 2 || result[3] != 2 || result[4] != 3 || result[5] != 3 {
		t.Error("FlatMap(func(x int) []int { return []int{x, x} }, []int{1, 2, 3}) != []int{1, 1, 2, 2, 3, 3}")
	}
}

func TestFlatten(t *testing.T) {
	result := Flatten([][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}})
	if result[0] != 1 ||
		result[1] != 2 ||
		result[2] != 3 ||
		result[3] != 4 ||
		result[4] != 5 ||
		result[5] != 6 ||
		result[6] != 7 ||
		result[7] != 8 {
		t.Error("Flatten([]interface{}{1, 2, []interface{}{3, 4, []interface{}{5, 6}}, 7, 8}) != []interface{}{1, 2, 3, 4, 5, 6, 7, 8}")
	}
}

func TestFoldL(t *testing.T) {
	result := FoldL(func(x, y int) int { return x + y }, 0, []int{1, 2, 3})
	if result != 6 {
		t.Error("Fold(func(x, y int) int { return x + y }, 0, []int{1, 2, 3}) != 6")
	}
}

func TestFoldR(t *testing.T) {
	result := FoldR(func(x, y int) int { return x + y }, 0, []int{1, 2, 3})
	if result != 6 {
		t.Error("Fold(func(x, y int) int { return x + y }, 0, []int{1, 2, 3}) != 6")
	}
}

func TestForEach(t *testing.T) {
	all := ""
	ForEach(func(x string) { all += x }, []string{"1", "2", "3"})
	if all != "123" {
		t.Error(`func(x string) {all += x}, []string{"1", "2", "3"} != "123"`)
	}
}

func TestJoin(t *testing.T) {
	result := Join(",", []string{"1", "2", "3"})
	if result[0] != "1" ||
		result[1] != "," ||
		result[2] != "2" ||
		result[3] != "," ||
		result[4] != "3" {
		t.Error(`Join([]string{",", "1", "2", "3"}) != "1,2,3"`)
	}
}

func TestLast(t *testing.T) {
	result, ok := Last([]int{1, 2, 3})
	if !ok || result != 3 {
		t.Error("Last([]int{1, 2, 3}) != 3")
	}
	result, ok = Last([]int{})
	if ok {
		t.Error("Last([]int{}) != false")
	}
}

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

func TestMapFoldL(t *testing.T) {
	list, result := MapFoldL(func(a, b int) (string, int) { return "n-" + strconv.Itoa(a), a + b }, 0, []int{1, 2, 3})
	if list[0] != "n-1" || list[1] != "n-2" || list[2] != "n-3" {
		t.Error(`MapFoldL(fun, 0, []int{1, 2, 3}) != []string{"n-1", "n-2", "n-3"}`)
	}
	if result != 6 {
		t.Error("MapFoldL(fun, 0, []int{1, 2, 3}) != 6")
	}
}

func TestMapFoldR(t *testing.T) {
	list, result := MapFoldR(func(a, b int) (string, int) { return "n-" + strconv.Itoa(a), a + b }, 0, []int{1, 2, 3})
	if list[0] != "n-3" || list[1] != "n-2" || list[2] != "n-1" {
		t.Error(`MapFoldL(fun, 0, []int{1, 2, 3}) != []string{"n-3", "n-2", "n-1"}`)
	}
	if result != 6 {
		t.Error("MapFoldL(fun, 0, []int{1, 2, 3}) != 6")
	}
}

func TestMax(t *testing.T) {
	m1, _ := Max([]int{1, 2, 3})
	if m1 != 3 {
		t.Error("Max([]int{1, 2, 3}) != 3")
	}
	m2, _ := Max([]float64{1.0, 2.0, 3.0})
	if m2 != 3.0 {
		t.Error("Max([]float64{1.0, 2.0, 3.0}) != 3.0")
	}
	m3, _ := Max([]string{"a", "b", "c"})
	if m3 != "c" {
		t.Error(`Max([]string{"a", "b", "c"}) != "c"`)
	}
}

func TestMember(t *testing.T) {
	if !Member(2, []int{1, 2, 3}) {
		t.Error("Member(2, []int{1, 2, 3}) != true")
	}
	if Member(2, []int{1, 3}) {
		t.Error("Member(2, []int{1, 3}) != false")
	}
}

func TestMin(t *testing.T) {
	m1, _ := Min([]int{1, 2, 3})
	if m1 != 1 {
		t.Error("Min([]int{1, 2, 3}) != 1")
	}
	m2, _ := Min([]float64{1.0, 2.0, 3.0})
	if m2 != 1.0 {
		t.Error("Min([]float64{1.0, 2.0, 3.0}) != 1.0")
	}
	m3, _ := Min([]string{"a", "b", "c"})
	if m3 != "a" {
		t.Error(`Min([]string{"a", "b", "c"}) != "a"`)
	}
}

func TestNth(t *testing.T) {
	n1, _ := Nth(0, []int{1, 2, 3})
	if n1 != 1 {
		t.Error("Nth(0, []int{1, 2, 3}) != 1")
	}
	n2, _ := Nth(2, []int{1, 2, 3})
	if n2 != 3 {
		t.Error("Nth(2, []int{1, 2, 3}) != 3")
	}
}

func TestNthTail(t *testing.T) {
	n1, ok := NthTail(0, []int{1, 2, 3})
	if !ok || n1[0] != 2 || n1[1] != 3 {
		t.Error("NthTail(0, []int{1, 2, 3}) != []int{2, 3}")
	}
	n2, ok := NthTail(1, []int{1, 2, 3})
	if !ok || n2[0] != 3 {
		t.Error("NthTail(1, []int{1, 2, 3}) != []int{3}")
	}
	_, ok = NthTail(3, []int{1, 2, 3})
	if ok {
		t.Error("NthTail(3, []int{1, 2, 3}) != false")
	}
}

func TestPartition(t *testing.T) {
	p1, p2 := Partition(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5})
	if p1[0] != 2 || p1[1] != 4 || p2[0] != 1 || p2[1] != 3 || p2[2] != 5 {
		t.Error(`Partition(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5}) != []int{{2, 4}, []int{1, 3, 5}`)
	}
}
