package uniclock

import (
	"time"
	"sync"
)

var counter int64
var mutex sync.Mutex

func Next() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	t := time.Now().UnixNano()
	if t > counter {
		counter = t
	} else {
		counter++
	}
	return counter
}