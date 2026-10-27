// lab 2, ex 5
package main
import "fmt"

func send(ch chan<- string){
	ch <- "rezolvare ex 5"
}

func receive(ch <-chan string) string{
	return <-ch
}

func main(){
	ch := make(chan string)
	

	go send(ch)
	msjRez := receive(ch)

	fmt.Println(msjRez)

}