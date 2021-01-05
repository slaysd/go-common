package channel

import (
	"math/rand"
	"testing"
	"time"
)

func makeDataset(n int) map[int]bool {
	data := make(map[int]bool)
	for i := 0; i < n; i++ {
		v := rand.Int() % 10000
		data[v] = true
	}
	return data
}

func TestMergeChannel(t *testing.T) {
	in1, in2, out := make(chan int), make(chan int), make(chan int)

	data := makeDataset(100)
	go func() {
		for k, _ := range data {
			if rand.Int() % 2 == 0 {
				in1 <- k
			} else {
				in2 <- k
			}

		}
	}()
	MergeChannel(out, in1, in2)

	select {
	case res := <-out:
		if _, ok := data[res]; !ok {
			t.Error("Wrong results")
		}
	case <-time.After(time.Second * 2):
		t.Error("Timeout")
	}
}

func BenchmarkMergeChannel(b *testing.B) {
	in1, in2, out := make(chan int), make(chan int), make(chan int)

	data := makeDataset(b.N)
	go func() {
		for k, _ := range data {
			if rand.Int() % 2 == 0 {
				in1 <- k
			} else {
				in2 <- k
			}
		}
		close(in1)
		close(in2)
	}()
	MergeChannel(out, in1, in2)

	for i := range out {
		_ = i
	}
}