package offer_2

import (
	"sync"
)

type singletonLazyAndLock struct{}

var insLL *singletonLazyAndLock
var mu sync.Mutex

func GetInsLL() *singletonLazyAndLock {
	mu.Lock()
	defer mu.Unlock()

	if insLL == nil {
		insLL = &singletonLazyAndLock{}
	}
	return insLL
}
