package breadthFirst

import (
	"reflect"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
    testCases := []struct {
        name         string
        f             func(item string) []string
        worklist     []string
        expectedSeen map[string]bool
    }{
        {
            name:         "Simple case",
            f:             func(item string) []string { return []string{"a", "b"} },
            worklist:     []string{"c"},
            expectedSeen: map[string]bool{"c": true, "a": true, "b": true},
        },
        {
            name:         "Circular dependency",
            f:             func(item string) []string { return []string{item} },
            worklist:     []string{"a"},
            expectedSeen: map[string]bool{"a": true},
        },
        {
            name:         "Empty worklist",
            f:             func(item string) []string { return []string{"a", "b"} },
            worklist:     []string{},
            expectedSeen: map[string]bool{},
        },
    }

    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            seen := make(map[string]bool)
            BreadthFirst(testCase.f, testCase.worklist)

            if !reflect.DeepEqual(seen, testCase.expectedSeen) {
                t.Errorf("Expected seen: %v, got: %v", testCase.expectedSeen, seen)
            }
        })
    }
}
/*
var seen map[string]bool

func TestBreadthFirst(t *testing.T) {
	// Define a simple graph as an adjacency list
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B", "F"},
		"F": {"C", "E"},
	}

	// Define the function to pass to BreadthFirst
	f := func(item string) []string {
		return graph[item]
	}

	// Call BreadthFirst with a starting node
	BreadthFirst(f, []string{"A"})

	// Check that all nodes were visited
	if !reflect.DeepEqual(seen, map[string]bool{
		"A": true,
		"B": true,
		"C": true,
		"D": true,
		"E": true,
		"F": true,
	}) {
		t.Errorf("Not all nodes were visited")
	}
}
*/