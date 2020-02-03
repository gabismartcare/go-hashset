package hashset

//Set interface of a Set
type Set interface {
	//Append a value to a set
	Append(s ...string)

	//Contains Returns true if the set contains the given value
	Contains(s string) bool

	// Return the current size of the set
	Len() int

	//ToSlice Return the current set as a slice
	ToSlice() []string

	//Delete Remove the given value from the set, returns true if the value was present
	Delete(s string) bool

	//IsEmpty Returns true if the current set is empty
	IsEmpty() bool
}
