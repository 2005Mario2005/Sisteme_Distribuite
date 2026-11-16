package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	Inbox     chan Memory
	Crashed   bool
}

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func main() {
	data, err := os.ReadFile("date_input_lab6_2.txt")
	if err != nil {
		fmt.Println("Eroare la citirea fisierului:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var n, s int
	fmt.Sscan(lines[0], &n, &s)

	tokens := make([]int, n)
	for i, f := range strings.Fields(lines[1]) {
		fmt.Sscan(f, &tokens[i])
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for _, f := range strings.Fields(lines[i+2]) {
			var v int
			fmt.Sscan(f, &v)
			adj[i] = append(adj[i], v)
		}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex.")
		return
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		m := make(Memory)
		m[Info{ID: i, Value: tokens[i]}] = struct{}{}
		nodes[i] = &Node{
			ID:        i,
			Neighbors: adj[i],
			M:         m,
			Inbox:     make(chan Memory, n),
		}
	}

	done := make(chan bool, n)
	crashedCount := 0
	t := 0

	for {
		if t > 0 && crashedCount < s && rand.Intn(2) == 1 {
			idx := rand.Intn(n)
			if !nodes[idx].Crashed && crashedCount < n-1 {
				nodes[idx].Crashed = true
				crashedCount++
			}
		}

		for i := 0; i < n; i++ {
			go func(node *Node) {
				if !node.Crashed {
					for _, neigh := range node.Neighbors {
						nodes[neigh].Inbox <- node.M
					}
				}
				done <- true
			}(nodes[i])
		}
		for i := 0; i < n; i++ {
			<-done
		}

		for i := 0; i < n; i++ {
			go func(node *Node) {
				nextM := make(Memory)
				for k := range node.M {
					nextM[k] = struct{}{}
				}
				for len(node.Inbox) > 0 {
					msg := <-node.Inbox
					if !node.Crashed {
						for k := range msg {
							nextM[k] = struct{}{}
						}
					}
				}
				node.M = nextM
				done <- true
			}(nodes[i])
		}
		for i := 0; i < n; i++ {
			<-done
		}

		if t > (s+1)*diam {
			break
		}
		t++
	}

	var correctNode *Node
	for i := 0; i < n; i++ {
		if !nodes[i].Crashed {
			correctNode = nodes[i]
			break
		}
	}

	for id := 0; id < n; id++ {
		for info := range correctNode.M {
			if info.ID == id {
				fmt.Printf("(%d, %d) ", info.ID, info.Value)
			}
		}
	}
	fmt.Println()
}
