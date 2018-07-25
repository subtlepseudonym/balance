package balance

import (
	"testing"
)

// TestPush inserts a number of elements into a BraceStack, then tests the length of the
// BraceStack's internal slice to ensure that it matches the number of elements pushed
func TestPush(t *testing.T) {
	testCases := []int{
		10,
		100,
		1000,
		10000,
	}

	for i, numElem := range testCases {
		stack := BraceStack{}
		for i := 0; i < numElem; i++ {
			stack.Push(1)
		}
		if len(stack.braces) != numElem {
			t.Errorf("test %d: stack size \"%d\" != expected size \"%d\"", i, len(stack.braces), numElem)
			t.Fail()
		}
	}
}

// TestPop inserts a number of elements into a BraceStack manually (not using Push), then
// calls Pop as many times as elements entered. If the BraceStack is empty early, the test
// fails
func TestPop(t *testing.T) {
	testCases := []int{
		10,
		100,
		1000,
		10000,
	}

	for i, numElem := range testCases {
		stack := BraceStack{}
		for i := 0; i < numElem; i++ {
			stack.braces = append(stack.braces, 1) // we don't care what goes into the stack
		}
		for i := 0; i < numElem; i++ {
			ok := stack.Pop()
			if !ok {
				t.Errorf("test %d: empty stack too soon: %d", i, len(stack.braces))
				t.FailNow()
			}
		}
		if len(stack.braces) != 0 {
			t.Errorf("test %d: still %d elements left in stack", i, len(stack.braces))
			t.Fail()
		}
	}
}

// TestLength inserts a number of elements into a BrackStack manually, then ensures that
// the length of the BraceStack's internal slice matches that number
func TestLength(t *testing.T) {
	testCases := []int{
		10,
		100,
		1000,
		10000,
	}

	for i, numElem := range testCases {
		stack := BraceStack{}
		for i := 0; i < numElem; i++ {
			stack.braces = append(stack.braces, 1)
		}
		if stack.Length() != numElem {
			t.Errorf("test %d: stack size \"%d\" != expected size \"%d\"", i, len(stack.braces), numElem)
			t.Fail()
		}
	}
}

// TestBalance sends input (where the balanced state and index of first unbalanced brace
// is known) to Balance() and compares the returned output against the expected output
func TestBalance(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"hello world", -1},    // no braces
		{"{}", -1},             // base case
		{"{{{foo();}}}{}", -1}, // valid with other
		{"{{}{}}", -1},         // valid
		{"valid {} case", -1},  // valid with preceeding and succeeding other

		{"{I", 0},             // orphan open
		{"{{used{to}", 0},     // orphan open with suceeding orphan open
		{"{be}{an", 2},        // orphan open with preceeding valid
		{"{{adventurer}", 0},  // orphan open with succeeding valid
		{"{like}{you}{{}", 4}, // orphan open with preceeding and succeeding valid

		{"}But", 0},             // orphan close
		{"}then}}", 0},          // orphan close with succeeding orphan close
		{"{I}{took}{}an}", 6},   // orphan close with preceeding valid
		{"}{arrow}{}to", 0},     // orphan close with succeeding valid
		{"{{the}} knee} {}", 4}, // orphan close with preceeding and succeeding valid
	}

	for i, test := range testCases {
		if output := Balance(test.input); output != test.expectedOutput {
			t.Errorf("test %d: output \"%d\" != expected \"%d\"", i, output, test.expectedOutput)
			t.Fail()
		}
	}
}

// TestFastBalance sends input (where the balanced state and index of first unbalanced brace
// is known) to Balance() and compares the returned output against the expected output
func TestFastBalance(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"hello world", -1},    // no braces
		{"{}", -1},             // base case
		{"{{{foo();}}}{}", -1}, // valid with other
		{"{{}{}}", -1},         // valid
		{"valid {} case", -1},  // valid with preceeding and succeeding other

		{"{I", 0},             // orphan open
		{"{{used{to}", 0},     // orphan open with suceeding orphan open
		{"{be}{an", 2},        // orphan open with preceeding valid
		{"{{adventurer}", 0},  // orphan open with succeeding valid
		{"{like}{you}{{}", 4}, // orphan open with preceeding and succeeding valid

		{"}But", 0},             // orphan close
		{"}then}}", 0},          // orphan close with succeeding orphan close
		{"{I}{took}{}an}", 6},   // orphan close with preceeding valid
		{"}{arrow}{}to", 0},     // orphan close with succeeding valid
		{"{{the}} knee} {}", 4}, // orphan close with preceeding and succeeding valid
	}

	for i, test := range testCases {
		if output := FastBalance(test.input); output != test.expectedOutput {
			t.Errorf("test %d: output \"%d\" != expected \"%d\"", i, output, test.expectedOutput)
			t.Fail()
		}
	}
}

// BenchmarkBalance runs a few varieties of test cases against Balance
func BenchmarkBalance(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"text", "hello world"},
		{"rand", "{{{}}}{}{{}{{{}}{}}{}}{{{}}{{}}{}{}{{{{}{}{{{}}}}}{}{}"},
		{"close", "}"},
		{"open", "{"},
		{"flat", "{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}"},
		{"deep", "{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}"},
		// long is 1000 runes long and balanced
		{"long", "{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}"},
		{"tough", "{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}"},
	}
	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Balance(bench.input)
			}
		})
	}
}

// BenchmarkFastBalance runs a few varieties of test cases against FastBalance
func BenchmarkFastBalance(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"text", "hello world"},
		{"rand", "{{{}}}{}{{}{{{}}{}}{}}{{{}}{{}}{}{}{{{{}{}{{{}}}}}{}{}"},
		{"close", "}"},
		{"open", "{"},
		{"flat", "{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}"},
		{"deep", "{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}"},
		// long and tough are 1000 runes long and balanced
		{"long", "{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}"},
		{"tough", "{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}{}"},
	}
	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FastBalance(bench.input)
			}
		})
	}
}
