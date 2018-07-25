package balance

// BraceStack is a specialized struct on which to store brace structs to determine if the
// input string is balanced
type BraceStack struct {
	braces []int
}

// Push puts the provided brace index on top of the BraceStack
func (s *BraceStack) Push(b int) {
	s.braces = append([]int{b}, s.braces...)
}

// Pop removes the topmost entry from the BraceStack and returns true or, if the BraceStack
// is empty, simply returns false
func (s *BraceStack) Pop() bool {
	if len(s.braces) == 0 {
		return false
	}
	s.braces = s.braces[1:]
	return true
}

// Length returns the number of entries in the BraceStack
func (s *BraceStack) Length() int {
	return len(s.braces)
}

// Balance takes an input string and returns the index of the first unbalanced brace,
// 0 indexed. If the input string is balanced, a -1 is returned
// This function is O(n) compute and O(n) memory
func Balance(in string) int {
	var stack BraceStack
	var braceIndex int
	for _, s := range []rune(in) {
		switch s {
		default:
			continue
		case '{':
			stack.Push(braceIndex)
			braceIndex += 1
		case '}':
			ok := stack.Pop()
			if !ok {
				return braceIndex
			}
			braceIndex += 1
		}
	}

	if stack.Length() == 0 {
		return -1
	}
	return stack.braces[len(stack.braces)-1]
}

// FastBalance is a faster and more memory efficient, but less readable, implementation
// of Balance
// This function is O(n) compute and O(1) memory, though the computation time is
// considerably faster than Balance. For small cases, FastBalance is ~3.5x faster and
// for longer cases, can be anywhere between from 5x to 30x faster, depending upon how
// flat (JSON-wise) the object formed by the input string is. Deeper objects see a greater
// compute time improvement over Balance
func FastBalance(in string) int {
	var openCount, braceIndex int
	leftmost := -1
	for _, s := range []rune(in) {
		switch s {
		default:
			continue
		case '{':
			if openCount == 0 {
				leftmost = braceIndex
			}
			openCount++
			braceIndex++
		case '}':
			if openCount == 1 {
				leftmost = -1
			} else if openCount == 0 {
				return braceIndex
			}
			openCount--
			braceIndex++
		}
	}

	return leftmost
}
