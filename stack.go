package main

type Stack []Target

// Size of the stack
func (s *Stack) Size() int {
	return len(*s)
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Clear the Stack
func (s *Stack) Clear() {
	*s = nil
}

// Peek at top value
func (s *Stack) Peek() (Target, error) {
	if s.IsEmpty() {
		return Target{}, errEmptyHistory
	} else {
		index := len(*s) - 1
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, nil
	}
}

// Push a new value onto the stack
func (s *Stack) Push(tar Target) {
	*s = append(*s, tar) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (Target, error) {
	if s.IsEmpty() {
		return Target{}, errEmptyHistory
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, nil
	}
}

// Specific to dd
func (s *Stack) AsString() string {
	str := "Hist: "
	for _, past := range *s {
		str = str + past.Login + "@" + past.Address + ">> "
	}
	return str
}
