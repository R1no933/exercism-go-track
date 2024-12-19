package tree

import "errors"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodes := make([]*Node, len(records))
	for _, record := range records {
		id := record.ID
		switch {
		case id <= record.Parent && id != 0:
			return nil, errors.New("higher id parent of lower id")
		case id == 0 && record.Parent != 0:
			return nil, errors.New("root node has parent")
		case id >= len(records):
			return nil, errors.New("non-continuous")
		case nodes[id] != nil:
			return nil, errors.New("duplicate node")
		}
		nodes[id] = &Node{ID: id}
	}
	for i := 0; i < len(records); i++ {
		id := records[i].ID
		records[i], records[id] = records[id], records[i]
	}
	for _, record := range records {
		if record.ID == 0 {
			continue
		}
		parent := nodes[record.Parent]
		parent.Children = append(parent.Children, nodes[record.ID])
	}
	return nodes[0], nil
}
