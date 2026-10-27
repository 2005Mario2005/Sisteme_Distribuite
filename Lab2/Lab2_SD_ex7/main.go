// lab 2, ex 7
package main
import (
	"fmt"
	"time"
)

type Coord struct{
	x int
	y int
}

func robot(id int, traseu []Coord, chOut chan<- Coord, chIn <-chan Coord, chDone chan<- bool){
	go func(){
		for _, poz := range traseu{
			time.Sleep(250 * time.Millisecond)
			chOut <- poz
		}
		close(chOut)
	}()

	for pozRobo := range chIn{
		fmt.Printf("Robotul %d a primit pozitia celuilalt robot: %v\n", id, pozRobo)
	}

	fmt.Printf("Robotul %d a terminat traseul\n", id)
	chDone <- true
}

func main(){
	chRobo1to2 := make(chan Coord)
	chRobo2to1 := make(chan Coord)
	chDone := make(chan bool)

	traseuRobo1 := []Coord{{0,0}, {1,0}, {2,1}, {3,1}, {4,4}, {4, 5}, {5, 5}}
	traseuRobo2 := []Coord{{0,0}, {2,2}, {5,3}, {6,4}, {5, 5}}

	go robot(1, traseuRobo1, chRobo1to2, chRobo2to1, chDone)
	go robot(2, traseuRobo2, chRobo2to1, chRobo1to2, chDone)

	<-chDone
	<-chDone

}