// lab 2, ex 2
package main

import (
	"fmt"
	"sync"
)

func afisareSalut(ch <- chan string, wg *sync.WaitGroup){
	defer wg.Done()
	strSalut := <-ch
	fmt.Println(strSalut)

}

func main(){
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go afisareSalut(ch, &wg)
	
	ch <- "salut"

	wg.Wait()
}