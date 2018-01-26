package scheduler

import (
	"container/heap"
)

type decisionTree struct ***REMOVED***
	// Count of tasks for the service scheduled to this subtree
	tasks int

	// Non-leaf point to the next level of the tree. The key is the
	// value that the subtree covers.
	next map[string]*decisionTree

	// Leaf nodes contain a list of nodes
	nodeHeap nodeMaxHeap
***REMOVED***

// orderedNodes returns the nodes in this decision tree entry, sorted best
// (lowest) first according to the sorting function. Must be called on a leaf
// of the decision tree.
//
// The caller may modify the nodes in the returned slice.
func (dt *decisionTree) orderedNodes(meetsConstraints func(*NodeInfo) bool, nodeLess func(*NodeInfo, *NodeInfo) bool) []NodeInfo ***REMOVED***
	if dt.nodeHeap.length != len(dt.nodeHeap.nodes) ***REMOVED***
		// We already collapsed the heap into a sorted slice, so
		// re-heapify. There may have been modifications to the nodes
		// so we can't return dt.nodeHeap.nodes as-is. We also need to
		// reevaluate constraints because of the possible modifications.
		for i := 0; i < len(dt.nodeHeap.nodes); ***REMOVED***
			if meetsConstraints(&dt.nodeHeap.nodes[i]) ***REMOVED***
				i++
			***REMOVED*** else ***REMOVED***
				last := len(dt.nodeHeap.nodes) - 1
				dt.nodeHeap.nodes[i] = dt.nodeHeap.nodes[last]
				dt.nodeHeap.nodes = dt.nodeHeap.nodes[:last]
			***REMOVED***
		***REMOVED***
		dt.nodeHeap.length = len(dt.nodeHeap.nodes)
		heap.Init(&dt.nodeHeap)
	***REMOVED***

	// Popping every element orders the nodes from best to worst. The
	// first pop gets the worst node (since this a max-heap), and puts it
	// at position n-1. Then the next pop puts the next-worst at n-2, and
	// so on.
	for dt.nodeHeap.Len() > 0 ***REMOVED***
		heap.Pop(&dt.nodeHeap)
	***REMOVED***

	return dt.nodeHeap.nodes
***REMOVED***
