// lab 2, ex 4
package main
import "fmt"

func main(){
	done := make(chan bool)

	go func(){
		fmt.Println("rez ex4, inceput")
		fmt.Println("rez ex4, final")
		done <- true
	}()

	fmt.Println(<-done)
}