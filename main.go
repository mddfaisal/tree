package main

import (
	"fmt"
	"tree/utils"
)

func CreateTree(root *utils.Node) {
	root.Element = "/"
	root.Node = []utils.Node{}

	for _, e := range []string{"A", "B"} {
		root.Node = append(root.Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}

	// insert B, C, D into A & M into B
	root.Node[1].Node = append(root.Node[1].Node, utils.Node{Element: "M", Node: []utils.Node{}})
	root.Node[1].Node[0].Node = append(root.Node[1].Node[0].Node, utils.Node{
		Element: "N",
		Node:    []utils.Node{},
	})

	for _, e := range []string{"B", "C", "D"} {
		root.Node[0].Node = append(root.Node[0].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}

	for _, e := range []string{"F", "G"} {
		root.Node[0].Node[0].Node = append(root.Node[0].Node[0].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}

	for _, e := range []string{"O", "P", "Q", "R"} {
		root.Node[0].Node[0].Node[1].Node = append(root.Node[0].Node[0].Node[1].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}
	root.Node[0].Node[0].Node[0].Node = append(root.Node[0].Node[0].Node[0].Node, utils.Node{
		Element: "N",
		Node:    []utils.Node{},
	})

	for _, e := range []string{"H", "I", "J"} {
		root.Node[0].Node[1].Node = append(root.Node[0].Node[1].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}
	root.Node[0].Node[1].Node[0].Node = append(root.Node[0].Node[1].Node[0].Node, utils.Node{
		Element: "S",
		Node:    []utils.Node{},
	})

	root.Node[0].Node[1].Node[1].Node = append(root.Node[0].Node[1].Node[1].Node, utils.Node{
		Element: "T",
		Node:    []utils.Node{},
	})
	for _, e := range []string{"K", "M", "L"} {
		root.Node[0].Node[2].Node = append(root.Node[0].Node[2].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}

	for _, e := range []string{"U", "V"} {
		root.Node[0].Node[2].Node[0].Node = append(root.Node[0].Node[2].Node[0].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}
	for _, e := range []string{"X", "Y"} {
		root.Node[0].Node[2].Node[1].Node = append(root.Node[0].Node[2].Node[1].Node, utils.Node{
			Element: e,
			Node:    []utils.Node{},
		})
	}
	root.Node[0].Node[2].Node[1].Node[1].Node = append(root.Node[0].Node[2].Node[1].Node[0].Node, utils.Node{
		Element: "Z",
		Node:    []utils.Node{},
	})
	root.Node[0].Node[2].Node[2].Node = append(root.Node[0].Node[2].Node[2].Node, utils.Node{
		Element: "W",
		Node:    []utils.Node{},
	})
}

func main() {
	r1 := new(utils.Node)
	CreateTree(r1)
	search := utils.Exp[4]
	fmt.Println("searching: ", search)
	found := utils.Search(r1, search)
	fmt.Println("match: ", found)
}
