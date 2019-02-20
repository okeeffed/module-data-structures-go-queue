package queue

type testCase struct {
	description string
	input       int
	inputQ      intQueue
	expected    int
	expectedQ   intQueue
}

var enqueueTestCases = []testCase{
	{
		description: "should enqueue correctly",
		input:       8,
		inputQ:      intQueue{[]int{3, 5, 3, 6}},
		expected:    5,
		expectedQ:   intQueue{[]int{3, 5, 3, 6, 8}},
	},
}

var dequeueTestCases = []testCase{
	{
		description: "should dequeue correctly",
		inputQ:      intQueue{[]int{3, 5, 3, 6}},
		expected:    6,
		expectedQ:   intQueue{[]int{3, 5, 3}},
	},
}

var dequeueErrorTestCases = []testCase{
	{
		description: "should throw dequeue error correctly",
		inputQ:      intQueue{[]int{}},
		expected:    0,
		expectedQ:   intQueue{[]int{}},
	},
}

var peekTestCases = []testCase{
	{
		description: "should peek correctly",
		inputQ:      intQueue{[]int{3, 5, 3, 6}},
		expected:    6,
		expectedQ:   intQueue{[]int{3, 5, 3, 6}},
	},
}
