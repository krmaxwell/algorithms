package main

import (
	"container/list"
	"fmt"
)

type Neighbors []string

type Graph map[string]Neighbors

func buildGraph() Graph {
	graph := Graph{
		"you":    Neighbors{"alice", "bob", "claire"},
		"bob":    Neighbors{"anuj", "peggy"},
		"alice":  Neighbors{"peggy"},
		"claire": Neighbors{"thom", "johnny"},
		"anuj":   Neighbors{},
		"peggy":  Neighbors{},
		"thom":   Neighbors{},
		"johnny": Neighbors{},
	}
	return graph
}

func isSeller(person interface{}) bool {
	name := fmt.Sprint(person)
	fmt.Printf("isSeller: checking %q\n", name)
	lastChar := name[:len(name)-1]
	return lastChar == "m"
}

func main() {

	var foundSeller bool
	graph := buildGraph()

	searchQ := list.New()

	for _, name := range graph {
		searchQ.PushBack(name)
	}

	for searchQ.Len() > 0 {
		firstElement := searchQ.Front()
		neighbors := firstElement.Value
		searchQ.Remove(firstElement)
		fmt.Printf("main: checking %q\n", neighbors)
		for _, person := range neighbors {

			if isSeller(person) {
				fmt.Printf("%q sells mangos\n", person)
				foundSeller = true

			}
		}
	}

	if !foundSeller {
		fmt.Println("Found no sellers")
	}

}
