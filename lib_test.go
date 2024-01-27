package main

import "testing"

func TestIntersectionWithTwoLists(t *testing.T) {
	items1 := []string{"a", "b", "c"}
	items2 := []string{"c", "d", "e"}

	intersectionItems := Intersection([][]string{items1, items2})
	want := 1
	got := len(intersectionItems)
	if got != want {
		t.Errorf("%d != %d", got, want)
	}

	wantLetter := "c"
	gotLetter := intersectionItems[0]
	if got != want {
		t.Errorf("%s != %s", gotLetter, wantLetter)
	}
}

func TestIntersectionWithThreeLists(t *testing.T) {
	items1 := []string{"a", "b", "c"}
	items2 := []string{"b", "c", "d"}
	items3 := []string{"c", "d", "e"}

	intersectionItems := Intersection([][]string{items1, items2, items3})
	want := 1
	got := len(intersectionItems)
	if got != want {
		t.Errorf("%d != %d", got, want)
	}

	wantLetter := "c"
	gotLetter := intersectionItems[0]
	if got != want {
		t.Errorf("%s != %s", gotLetter, wantLetter)
	}
}

func TestUnionWithThreeLists(t *testing.T) {
	items1 := []string{"a", "b", "c"}
	items2 := []string{"b", "c", "d"}
	items3 := []string{"c", "d", "e"}

	unionItems := Union([][]string{items1, items2, items3})
	want := 5
	got := len(unionItems)
	if got != want {
		t.Errorf("%d != %d", got, want)
	}

}
