package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() { l.counter += rand.Intn(3) + 1 }
func (l *LamportClock) Send() int { 
	l.Tick()
	return l.counter
}
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct{ sender , ts int }

type Event struct {
	process int
	internal_counter int
	kind string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool { return a.internal_counter < b.internal_counter }

func main() {
	ch12 := make(chan Message, 1)
	ch23 := make(chan Message, 1)
	ch31 := make(chan Message, 1)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var ev []Event

	logEvent := func(process int, counter int, kind string) {
		mu.Lock()
		ev = append(ev, Event{process: process, internal_counter: counter, kind: kind})
		mu.Unlock()
	}

	wg.Add(3)

	go func() {
		defer wg.Done()

		clock := LamportClock{0}
		ts := clock.Send()
		logEvent(1, ts, "send")

		ch12 <- Message{sender: 1, ts: ts}
		msg := <-ch31
		clock.Receive(msg.ts)
		logEvent(1, clock.counter, "recv")
	}()

	go func(){
		defer wg.Done()
		clock := LamportClock{0}

		msg := <-ch12
		clock.Receive(msg.ts)
		logEvent(2, clock.counter, "recv")

		time.Sleep(time.Millisecond * 10) 
		ts := clock.Send()
		logEvent(2, ts, "send")
		ch23 <- Message{sender: 2, ts: ts}
	}()

	go func() {
		defer wg.Done()
		clock := LamportClock{0}

		// P3 primește de la P2
		msg := <-ch23
		clock.Receive(msg.ts)
		logEvent(3, clock.counter, "recv")

		// P3 trimite la P1
		time.Sleep(time.Millisecond * 10)
		ts := clock.Send()
		logEvent(3, ts, "send")
		ch31 <- Message{sender: 3, ts: ts}
	}()

	wg.Wait()

	sort.Slice(ev, func(i, j int) bool{
		return less(ev[i], ev[j])
	})

	for _, e := range ev {
		fmt.Printf("proces P%d, timp lamport: %d tip: %s\n", e.process, e.internal_counter, e.kind)
	}
}