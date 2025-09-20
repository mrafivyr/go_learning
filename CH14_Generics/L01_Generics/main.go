package main

/*
// Explicit creation of variable to return its zero value when slice is empty
func getLast[T any](s []T) T {
	var genZero T
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return genZero
}
// Named return will create implicit zero value of a type
func getLast[T any](s []T) (result T) {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return
}
*/

// new(T) creates a pointer to a new value of type T, initialized to its zero value.
// *new(T) dereferences that pointer, returning the zero value directly.
// This approach is a single, compact expression, which some developers prefer.
func getLast[T any](s []T) T {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return *new(T)
}
