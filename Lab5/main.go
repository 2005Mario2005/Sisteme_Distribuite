package main

import (
	"fmt"
	"math"
	"os"
)

type Robot struct {
	x float64
	y float64
	act bool
	id int
}

const CommunicationRadius float64 = 10.0
const Epsilon float64 = 0.0001

func distance(r1, r2 Robot) float64 {
	dx := r2.x - r1.x
	dy := r2.y - r1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func robotAgent(id int, inCh <-chan []Robot, outCh chan<- Robot) {

	for state := range inCh {
		me := state[id]
		
		if !me.act {
			outCh <- me
			continue
		}

		sumx, sumy := me.x, me.y
		neighborsCount := 1.0

		for j := 0; j < len(state); j++ {
			if id != j && state[j].act {
				dist := distance(me, state[j])
				
				if dist <= CommunicationRadius {
					sumx += state[j].x
					sumy += state[j].y
					neighborsCount++
				}
			}
		}

		me.x = sumx / neighborsCount
		me.y = sumy / neighborsCount

		outCh <- me
	}
}

func main() {
	file, err := os.Open("date_robot.txt")
	if err != nil {
		fmt.Println("Eroare la deschiderea fișierului: ", err)
		return
	}
	defer file.Close()

	var n int
	fmt.Fscanln(file, &n)

	robots := make([]Robot, 0, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscanln(file, &x, &y)
		robots = append(robots, Robot{x: x, y: y, act: true, id: i})
	}

	inChs := make([]chan []Robot, n)
	outCh := make(chan Robot, n)

	for i := 0; i < n; i++ {
		inChs[i] = make(chan []Robot, 1) 
		go robotAgent(i, inChs[i], outCh)
	}

	maxItr := 1000
	actCount := n

	for iter := 0; iter < maxItr && actCount > 1; iter++ {
		
		for i := 0; i < n; i++ {
			stateCopy := make([]Robot, len(robots))
			copy(stateCopy, robots)
			inChs[i] <- stateCopy
		}

		newPositions := make([]Robot, n)
		for i := 0; i < n; i++ {
			r := <-outCh
			newPositions[r.id] = r
		}

		for i := 0; i < n; i++{
			robots[i] = newPositions[i]
		}

		for i := 0; i < n; i++ {
			if !robots[i].act {
				continue
			}

			for j := i + 1; j <n; j++{
				if robots[j].act{
					if distance(robots[i], robots[j]) < Epsilon {
						robots[j].act = false
						actCount--
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		close(inChs[i])
	}
	
	fmt.Printf("%d roboti ramasi activi in final\n", actCount)
}