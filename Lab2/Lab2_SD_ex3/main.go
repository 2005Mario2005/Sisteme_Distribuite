// lab 2, ex 3
package main
import "fmt"

func main(){
	ch := make(chan string, 2)

	go func(){
		ch <- "rezolvare"
		ch <- "ex3"
		close(ch)
	}()

	for elm := range ch{
		fmt.Println(elm)
	}
}