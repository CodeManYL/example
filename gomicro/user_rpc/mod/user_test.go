package mod

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//func BenchmarkTest2(b *testing.B){
	//c := 0
	//for i := 0;i < 100000000000; i ++ {
	//	c += i
	//}
//}

//
func BenchmarkTest1(b *testing.B){
	a := make(map[string][]string)
	a["1"] = []string{"1"}

	locak := sync.RWMutex{}

	go func() {

		for i := 0; i < 100; i ++ {
			locak.Lock()
			a[fmt.Sprintf("%v",i)] = []string{fmt.Sprintf("%v",i)}
			locak.Unlock()
		}

	}()

	for i := 0; i < 1000000; i ++ {
		locak.RLock()
		_ = a["1"]
		locak.RUnlock()
	}

}

func BenchmarkTest2(b *testing.B) {
	a := make(map[string][]string)
	a["1"] = []string{"1"}

	var c atomic.Value
	c.Store(a)

	go func() {
		for i := 0; i < 100; i ++ {
			newMap := make(map[string][]string)
			newMap[fmt.Sprintf("%v",i)] = []string{fmt.Sprintf("%v",i)}
			c.Store(newMap)
		}
	}()

	for i := 0; i < 1000000; i ++ {
		res := c.Load().(map[string][]string)
		_ = res["1"]
	}

}