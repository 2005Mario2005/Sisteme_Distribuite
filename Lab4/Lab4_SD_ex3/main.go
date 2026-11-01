package main

import (
	"fmt"
	"sync"
)

func nodFloodingDecizieTareConex(index int, vot float64, votari int, inCh []<-chan float64, outCh []chan<- float64, wg *sync.WaitGroup){
	defer wg.Done()

	votInit := vot
	di := float64(len(inCh))

	for t := 1; t <= votari; t++{
		for _, ch := range outCh{
			ch <- votInit
		}

		sumaVecini := 0.0
		for _, ch := range inCh{
			sumaVecini += <-ch
		}

		votInit = (votInit + sumaVecini) / (1.0 + di)
	}

	if index == 0{
		votMajoritar := 0
		if votInit >= 0.5{
			votMajoritar = 1
		}
		fmt.Printf("Votul final este %d\n", votMajoritar)
	}

}

func main(){
	// acelasi graf ca la ex 2, refolosire cod
	// indici:    		 0    1    2    3    4    5    6    7    8    9
	voturi := []float64{1.0, 0.0, 1.0, 1.0, 0.0, 1.0, 1.0, 0.0, 1.0, 1.0}
	n := len(voturi)
	
	edges := [][2]int{ // graful din curs pentru ca imi era lene sa fiu creativ
    	{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 0}, {9, 2}, {1, 4}, {2, 5}, {3, 0}, {3, 6}, {4, 7}, {5, 8}, {6, 9}, {7, 0}, {8, 1},
	}

	inChs := make([][]<-chan float64, n)
	outChs := make([][]chan<- float64, n)

	for _, edge := range edges {
		start := edge[0] 
		destinatie := edge[1] 

		ch := make(chan float64, 1)

		outChs[start] = append(outChs[start], ch)

		inChs[destinatie] = append(inChs[destinatie], ch)
	}


	var wg sync.WaitGroup

	for i := 0; i < n; i++{
		wg.Add(1)

		go nodFloodingDecizieTareConex(i, voturi[i], 50, inChs[i], outChs[i], &wg)
	}

	wg.Wait()
}