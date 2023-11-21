package breadthFirst

import (
	"reflect"
	"testing"
)

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
