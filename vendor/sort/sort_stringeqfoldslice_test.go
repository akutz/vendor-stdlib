package sort_test

import (
	"sort"
	"testing"
)

func TestStringsEqFold(t *testing.T) {
	szs := []string{"b", "c", "A"}
	sort.StringsEqFold(szs)
	assertEl(t, szs, 0, "A")
	assertEl(t, szs, 1, "b")
	assertEl(t, szs, 2, "c")

	szs = []string{"Ardvark", "ARDVARKS", "2ardvark"}
	sort.StringsEqFold(szs)
	assertEl(t, szs, 0, "2ardvark")
	assertEl(t, szs, 1, "Ardvark")
	assertEl(t, szs, 2, "ARDVARKS")
}

func assertEl(t *testing.T, szs []string, idx int, exp string) {
	if act := szs[idx]; act != exp {
		t.Errorf("szs[%d] != exp: %s", act)
		t.Fail()
	}
}

func TestStringsEqFoldAreSorted(t *testing.T) {
	if !sort.StringsEqFoldAreSorted([]string{"A", "b", "c"}) {
		t.Error("alphabet is sorted")
		t.Fail()
	}
	if sort.StringsEqFoldAreSorted(
		[]string{"Ardvark", "ARDVARKS", "2ardvark"}) {
		t.Error("ardvarks are not sorted")
		t.Fail()
	}
}
