package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	title string
	adj   []*node
}

type edge struct {
	dad *node
	son *node
}

type graph struct {
	V map[string]*node
	A map[string]*edge
}

func insertNode(G *graph, title string) bool {
	var nv node
	nv.title = title
	G.V[title] = &nv
	return true
}

func insertEdge(G *graph, title string) bool {
	var na edge
	d, ok := G.V[title[0:1]]
	if !ok {
		return false
	}
	s, ok := G.V[title[3:4]]
	if !ok {
		return false
	}
	na.dad = d
	na.son = s
	G.A[title] = &na
	v := G.V[title[0:1]]
	v.adj = append(v.adj, na.son)
	G.V[title[0:1]] = v
	return true
}

func newGraph(path string) graph {
	var G graph
	G.V = make(map[string]*node)
	G.A = make(map[string]*edge)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	countline := 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 1 {
			if line < "a" || line > "z" {
				log.Fatal("Invalid File! Error, line: ", countline)
				os.Exit(1)
			}
			if !insertNode(&G, line) {
				log.Fatal("Invalid File! Error, line: ", countline)
				os.Exit(1)
			}
		} else if len(line) == 4 {
			if line[0:1] < "a" || line[1:3] != "->" || line[3:4] > "z" {
				log.Fatal("Invalid File! Error, line: ", countline)
				os.Exit(1)
			}
			if !insertEdge(&G, line) {
				log.Fatal("Invalid File! Error line: ", countline)
				os.Exit(1)
			}
		} else {
			log.Fatal("Invalid File! Error line: ", countline)
			os.Exit(1)
		}
		countline++
	}
	return G
}

func main() {
	G := newGraph("sampleInput.txt")
	fmt.Println("\nGraph:")
	for key, value := range G.V {
		fmt.Print(key, " ->")
		for i := 0; i < len(value.adj); i++ {
			fmt.Print(" ", value.adj[i].title)
		}
		fmt.Println()
	}
}
