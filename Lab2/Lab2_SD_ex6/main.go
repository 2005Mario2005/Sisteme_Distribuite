// lab 2, ex 6
package main
import (
	"fmt"
	"math/rand"
	"time"
)

func robot1(ch chan<- bool){
	time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	ch <- true
}

func robot2(ch chan<- bool){
	time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	ch <- true
}

func main(){
	chRobot1 := make(chan bool)
	chRobot2 := make(chan bool)

	go robot1(chRobot1)
	go robot2(chRobot2)

	select{
	case <-chRobot1:
		fmt.Println("Robot 1 a terminat")
	case <-chRobot2:
		fmt.Println("Robot 2 a terminat")
	}
}