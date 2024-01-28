package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("TOtal CPU", totalCPU)
	
	runtime.GOMAXPROCS(8)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)
	
	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("TOtal GOROUtine", totalGoRoutine)
	
	group.Wait()
}