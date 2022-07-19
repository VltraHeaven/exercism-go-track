package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type mutexMap struct {
	freqMap FreqMap
	mtx     sync.Mutex
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	mm := mutexMap{freqMap: map[rune]int{}}
	var wg sync.WaitGroup
	for i := range l {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for k, v := range Frequency(s) {
				mm.mtx.Lock()
				mm.freqMap[k] += v
				mm.mtx.Unlock()
			}
		}(l[i])
	}
	wg.Wait()
	return mm.freqMap
}
