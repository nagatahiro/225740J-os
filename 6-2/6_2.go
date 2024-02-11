package main




import (
	"fmt"
	"sync"
	//"math/rand"

)
type Semaphore struct {
	ch chan struct{}
}
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, n),
	}
}
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}
func (s *Semaphore) Release() {
	<-s.ch
}


func main() {
	semaphore := NewSemaphore(3) 
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ { 
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Requesting access to semaphore\n", id)
			semaphore.Acquire() // リクエスト
			fmt.Printf("Goroutine %d: Access to semaphore acquired\n", id)
			semaphore.Release() // 終了
			fmt.Printf("Goroutine %d: Released access to semaphore\n", id)
		}(i)
	}
	wg.Wait()
}
