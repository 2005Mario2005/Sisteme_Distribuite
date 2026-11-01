package main

import (
	"fmt"
	"sync"
)


func nodFloodingTareConex(id int, diam int, inCh []<-chan int, outCh []chan<- int, wg *sync.WaitGroup){
	defer wg.Done()

	max_id := id

	for _, ch := range outCh{
		ch <- max_id
	}

	for i := 1; i <= diam; i++{
		for _, ch := range inCh {
			receiveId := <-ch
			if receiveId > max_id{
				max_id = receiveId
			}
		}	

		if i == diam {
			if max_id == id {
				fmt.Printf("Nodul %d este liderul\n", id)
			}
		} else {
			for _, ch := range outCh {
				ch <- max_id
			}
		}
	}


}


func main(){
	// indici:  0  1  2  3   4  5  6   7  8   9
	id := []int{9, 2, 6, 14, 4, 3, 17, 8, 22, 19}
	n := len(id)
	
	edges := [][2]int{ // graful din curs pentru ca imi era lene sa fiu creativ
    	{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 0}, {9, 2}, {1, 4}, {2, 5}, {3, 0}, {3, 6}, {4, 7}, {5, 8}, {6, 9}, {7, 0}, {8, 1},
	}

	inChs := make([][]<-chan int, n)
	outChs := make([][]chan<- int, n)

	for _, edge := range edges {
		start := edge[0] 
		destinatie := edge[1] 

		ch := make(chan int, 1)

		outChs[start] = append(outChs[start], ch)

		inChs[destinatie] = append(inChs[destinatie], ch)
	}


	var wg sync.WaitGroup

	diam := n - 1

	for i := 0; i < n; i++{
		wg.Add(1)

		go nodFloodingTareConex(id[i], diam, inChs[i], outChs[i], &wg)
	}

	wg.Wait()
}