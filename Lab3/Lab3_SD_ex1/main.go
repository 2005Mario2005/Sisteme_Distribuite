package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// (a) Completați funcția estimate folosind Algoritmul lui Cristian. 
func estimate(t1, t2, t3, t4 time.Time) (theta , delta time.Duration) { 
	delta = t4.Sub(t1) - t3.Sub(t2)

	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	
	return theta, delta
}

func measure(c_clock , s_clock *Clock) (theta , delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()
	
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()
	
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	clock := Clock{offset: 0}

	offsetInit := 100 * time.Millisecond

	// (a) Simulați apoi în main o cerere și un răspuns folosind funcția measure și verificați dacă offset-ul inițial a scăzut.
	clientA := Clock{offset: offsetInit}
	thetaA, _ := measure(&clientA, &clock)
	clientA.offset += thetaA

	fmt.Printf("Offset initial: %v\n", offsetInit)
	fmt.Printf("Offset dupa ajustare: %v \n\n", clientA.offset)

	// (b) Repetați simularea de 10 ori ( cu același offset inițial) și calculați δ și θ minime. 

	clientB := Clock{offset: offsetInit}
	var minDelta time.Duration
	var bestTheta time.Duration
	prima := true

	for i := 0; i < 10; i++ {
		t, d := measure(&clientB, &clock)

		if prima || d < minDelta{
			minDelta = d
			bestTheta = t
			prima = false
		}
	}

	clientB.offset += bestTheta

	fmt.Printf("Cel mai bun Round-trip: %v\n", minDelta)
	fmt.Printf("Decalaj corespunzator: %v\n", bestTheta)
	fmt.Printf("Offset dupa ajustare: %v\n", clientB.offset)
}