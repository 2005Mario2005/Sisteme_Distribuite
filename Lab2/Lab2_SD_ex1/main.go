// lab 2, ex 1
package main
import "fmt"

func main(){
	ch := make(chan string)

	go func(){
		ch <- "salut"
	}()

	fmt.Println(<-ch)
}