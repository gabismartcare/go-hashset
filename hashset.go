// Package hashet is a simple implemetation of a set in Go

package hashset

type hashMap map[string]bool

// HashSet implementation of a set
type HashSet struct {
	set hashMap
}

//New Creates a new Set
func New() Set {
	hashset := &HashSet{make(hashMap)}
	return hashset
}

// Append a value to a set
func (h *HashSet) Append(s ...string) {
	for _, v := range s {
		h.set[v] = true
	}
}

// Contains Return true if the given value is present
func (h *HashSet) Contains(value string) bool {
	return h.set[value]
}

//IsEmpty Returns true if the set empty
func (h *HashSet) IsEmpty() bool {
	return len(h.set) == 0
}

// Len Returns the size of the current set
func (h *HashSet) Len() int {
	return len(h.set)
}

// ToSlice Returns the current
func (h *HashSet) ToSlice() []string {
	slice := make([]string, 0, h.Len())
	for k := range h.set {
		slice = append(slice, k)
	}
	return slice
}

//Delete remove the given value from the set
func (h *HashSet) Delete(id string) bool {
	defer delete(h.set, id)
	return h.set[id]
}
