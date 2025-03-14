package genericstack


type esvStack[T any] struct {
    Push func(T)
    Pop func() T
}

// EsvStack creates a generic stack of type T.
//
// It returns a struct of type esvStack[T] with two methods:
// - Push, which appends an element of type T to the stack.
// - Pop, which removes and returns the top element of the stack.
//
// The stack is implemented as a slice.
//
// Returns a stack of type T.
func EsvStack[T any]() esvStack[T] {
    slice := make([]T, 0)
    return esvStack[T]{
        Push: func(i T) {
            slice = append(slice, i)
        },
        Pop: func() T {
			if len(slice) == 0 {
				var empty T
				return empty
			}
            result := slice[len(slice)-1]
            slice = slice[:len(slice)-1]
            return result
        },
    }
}