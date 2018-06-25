package tree

import (
	"fmt"
	"sort"
	"strings"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ById []Record

func (record ById) Len() int           { return len(record) }
func (record ById) Swap(i, j int)      { record[i], record[j] = record[j], record[i] }
func (record ById) Less(i, j int) bool { return record[i].ID < record[j].ID }

type ByParent []Record

func (record ByParent) Len() int           { return len(record) }
func (record ByParent) Swap(i, j int)      { record[i], record[j] = record[j], record[i] }
func (record ByParent) Less(i, j int) bool { return record[i].Parent < record[j].Parent }

type ByParentAndId []Record

func (record ByParentAndId) Len() int      { return len(record) }
func (record ByParentAndId) Swap(i, j int) { record[i], record[j] = record[j], record[i] }
func (record ByParentAndId) Less(i, j int) bool {
	return record[i].Parent < record[j].Parent || record[i].Parent == record[j].Parent && record[i].ID < record[j].ID
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// presorting
	sort.Sort(ByParentAndId(records))

	currentParent := &Node{}
	referenceMap := make(map[int]*Node)

	for _, record := range records {

		// sanity checks
		// fmt.Printf("[%d, %d] => ", record.ID, record.Parent)

		if record.ID == record.Parent {
			switch {
			case record.ID != 0:
				return nil, fmt.Errorf("Invalid record: parent to itself")
			case referenceMap[0] != nil:
				return nil, fmt.Errorf("Invalid record: duplicate root")
			default: // [0, 0] is permissible, should be the first element
				referenceMap[0] = currentParent
				// fmt.Printf("ROOT: %v\n", referenceMap)
				continue
			}
		}

		if referenceMap[record.Parent] == nil {
			// fmt.Printf("ERROR MAP: %+v\n", referenceMap)
			return nil, fmt.Errorf("Invalid record: no known parent for record: [%d, %d]", record.ID, record.Parent)
		}

		child := Node{ID: record.ID}

		switch {
		case record.Parent == currentParent.ID:
			// fmt.Printf("Matching parent %d\n", record.Parent)
			// fmt.Printf("%v => ", referenceMap)
			currentParent.Children = append(currentParent.Children, &child)
			referenceMap[record.ID] = &child
			// fmt.Printf("%v\n", referenceMap)
		case record.Parent > currentParent.ID:
			// fmt.Printf("Changing parent %d\n", record.Parent)
			if currentParent, ok := referenceMap[record.Parent]; ok {
				// fmt.Printf("%v => ", referenceMap)
				currentParent.Children = append(currentParent.Children, &child)
				referenceMap[record.ID] = &child
				// fmt.Printf("%v\n", referenceMap)
			} else {
				return nil, fmt.Errorf("Invalid record: Parent %d does not exist (current parent %d)", record.Parent, currentParent.ID)
			}
		default: //impossible behavior
			return nil, fmt.Errorf("Invalid record: Parent %d not processed earlier (current parent %d); suspecting failed presort", record.Parent, currentParent.ID)
			// 	if len(fastMap[record.Parent] == nil {
			// 		fmt.Println("------")
			// 		displayTree(root, 0)
			// 		fmt.Println("===================")
			// 		return nil, fmt.Errorf("Can't find a child with id: %d for %d", record.Parent, record.ID)
			// 	} else {
			// 		current.ID = record.Parent
			// 	}
		}
		// nodeMap[record.Parent] = append(nodeMap[record.Parent], &newChild)
		// fmt.Printf("%v\n", referenceMap)
	}

	// fmt.Println("------")
	// displayTree(root, 0)
	// fmt.Println("===================")
	// fmt.Printf("%#v", nodeMap)

	// fmt.Printf("\nFINAL REF MAP:\n%#v\n\n", referenceMap)
	return referenceMap[0], nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}

// func displayMap(m map[int]*Node) {
// 	for _, parent := range m {
// 		fmt.Printf("%d: \n")
// 	}
// }

func displayTree(root *Node, indent int) {
	if indent > 0 {
		fmt.Printf("%s+-", strings.Repeat("  ", indent-1))
	}
	fmt.Printf("%d\n", root.ID)
	for _, child := range root.Children {
		displayTree(child, indent+1)
	}
}

func childByID(root *Node, childId int) *Node {
	// fmt.Printf("CHILD BY ID, CHECKING FOR RECORD WITH CHILD ID %d\n", childId)
	displayTree(root, 0)

	for _, child := range root.Children {
		if child.ID == childId {
			// fmt.Printf("FOUND CHILD WITH ID %d\n", child.ID)
			// displayTree(child, 0)
			// fmt.Printf("\nDONE\n")
			return child
		}
	}
	// fmt.Printf("NOT FOUND\n")
	return nil
}
