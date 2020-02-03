package hashset

import (
	"testing"
)

func Test(t *testing.T) {
	set := New()
	set.Append("string1")
	set.Append("string1")
	set.Append("string1")
	if set.Len() != 1 {
		t.Errorf("Expected %d; got %d", 1, set.Len())
	}
}

func TestHashset_Contains(t *testing.T) {
	testData := []struct {
		testStr string
		output  bool
	}{
		{"string1", true},
		{"string2", true},
		{"string1 ", false},
		{"String1", false},
		{"dsfds", false},
	}
	set := New()
	set.Append("string1")
	set.Append("string2")

	for _, data := range testData {
		err := set.Contains(data.testStr)
		if err != data.output {
			t.Errorf("Failure incorrect %s", data.testStr)
		}
	}

}

func TestHashset_Len(t *testing.T) {
	set := New()
	set.Append([]string{"string1", "string2", "string2"}...)
	if set.Len() != 2 {
		t.Errorf("Expected %d; got %d", 2, set.Len())
	}
}

func TestRemove(t *testing.T) {
	set := New()
	set.Append("toto")
	if !set.Contains("toto") {
		t.Error("Expecting toto be in the set")
	}
	contains := set.Delete("toto")
	if !contains {
		t.Error("Expecting toto be in the set")
	}
	if set.Contains("toto") {
		t.Error("Expecting toto removed from the set")
	}
}

func TestToSlice(t *testing.T) {

	initialStrings := []string{"string1", "string2", "string3"}
	set := New()
	set.Append(initialStrings...)

	if set.Len() != 3 {
		t.Errorf("Expecting 3 values in slice, got %d", set.Len())
	}
	found := false
	for _, v := range set.ToSlice() {
		found = false
		for _, v2 := range initialStrings {
			if v == v2 {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Cannot find %s in new slice", v)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	emptySet := New()
	if !emptySet.IsEmpty() {
		t.Error("Expected set to be empty")
	}

	initialStrings := []string{"string1", "string2", "string3"}
	notEmptySet := New()
	notEmptySet.Append(initialStrings...)
	if notEmptySet.IsEmpty() {
		t.Error("Expected set to be not empty")

	}
}
