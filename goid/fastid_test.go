package goid

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGenID(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)                        // using 100 goroutine to generate 10000 ids
	results := make(chan int64, 10000) // store result
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				id := CommonConfig.GenInt64ID()
				t.Logf("id: %b \t %x \t %d", id, id, id)
				results <- id
			}
			wg.Done()
		}()
	}
	wg.Wait()

	m := make(map[int64]bool)
	for i := 0; i < 10000; i++ {
		select {
		case id := <-results:
			if _, ok := m[id]; ok {
				t.Errorf("Found duplicated id: %x", id)
				// return
			} else {
				m[id] = true
			}
		case <-time.After(2 * time.Second):
			t.Errorf("Expect 10000 ids in results, but got %d", i)
			return
		}
	}
}
func TestGenID2(t *testing.T) {
	g := ConstructConfigWithMachineID(40, 7, 11, 10)
	g2 := ConstructConfigWithMachineID(40, 7, 12, 10)
	// g2 := ConstructConfigWithMachineID(40, 11, 12, 2)

	for i := 0; i < 100; i++ {
		id := g.GenInt64ID()
		id2 := g2.GenInt64ID()
		fmt.Printf("id1:%v\tid2:%v\n", id, id2)
	}
}
func ExampleConfig_recommendedSettings() {
	id := CommonConfig.GenInt64ID()
	id2 := CommonConfig.GenInt64ID()
	fmt.Printf("id generated: %v %v\n", id, id2)
}

func ExampleConfig_customizedSettings() {
	var config = ConstructConfigWithMachineID(40, 11, 12, 2)
	id := config.GenInt64ID()
	fmt.Printf("id generated: %v\n", id)
}

func BenchmarkGenID(b *testing.B) {
	for i := 0; i < 10; i++ {
		ExampleConfig_customizedSettings()
		// BenchmarkConfig.GenInt64ID()
	}
}

func BenchmarkGenIDP(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			BenchmarkConfig.GenInt64ID()
		}
	})
}
