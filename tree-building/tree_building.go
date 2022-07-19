package tree

import (
	"errors"
	"sort"
)

const root = 0

// Record is a struct containing int fields ID and Parent
type Record struct {
	ID, Parent int
}

// Node is a struct containing int field ID and []*Node field Children.
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs a Node from a list Record containing nested child nodes
func Build(records []Record) (*Node, error) {
	nodes := map[int]*Node{}
	// Sort the records to align indices with the record ids
	// First record.ID should always == root
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {
		switch {
		case r.ID == r.Parent && r.ID != root:
			return nil, errors.New("cyclical id/parent id")
		case r.Parent > r.ID:
			return nil, errors.New("parent id cannot higher than node id")
		case i != r.ID:
			return nil, errors.New("duplicate or non-contiguous ids")
		default:
			nodes[r.ID] = &Node{ID: r.ID}
			if r.ID != root {
				nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
			}
		}
	}

	return nodes[root], nil
}
