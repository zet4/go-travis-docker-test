package lib

import (
	"sync/atomic"
)

var counter = uint64(0)

// IncrementAndReturn will get a value, increment it and return temp.
func IncrementAndReturn() uint64 {
	temp := atomic.LoadUint64(&counter)
	atomic.AddUint64(&counter, 1)
	return temp
}

// Reset clears a counter value
func Reset() {
	atomic.StoreUint64(&counter, 0)
}
