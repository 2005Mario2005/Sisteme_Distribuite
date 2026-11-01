package main

import (
	"fmt"
	"sync"
)


func nodFlooding(id int, inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup){
	defer wg.Done()

	outCh <- id

	for{
		receiveId := <- inCh

		if receiveId == id{
			fmt.Printf("Nodul %d este liderul\n", id)
			outCh <- -1
			break
		}else if receiveId == -1{
			outCh <- -1
			break
		} else if receiveId > id{
			outCh <- receiveId
		} 
	}
}


func main(){
	id := []int{1, 4, 2, 9, 34, 6, 19, 7}
	n := len(id)

	var wg sync.WaitGroup

	inel := make([]chan int, n)
	for i := 0; i < n; i++{
		inel[i] = make(chan int, 1)
	}

	for i := 0; i < n; i++{
		wg.Add(1)

		inCh := inel[i]
		outCh := inel[(i + 1) % n]

		go nodFlooding(id[i], inCh, outCh, &wg)
	}

	wg.Wait()
}