package offer_2

import (
	"sync"
)

type singletonLazyAndDoubleLock struct{}

var insLDL *singletonLazyAndDoubleLock
var muLDL sync.Mutex

func GetIninsLDL() *singletonLazyAndDoubleLock {

	if insLDL == nil {
		muLDL.Lock()
		defer muLDL.Unlock()
		if insLDL == nil {
			insLDL = &singletonLazyAndDoubleLock{}
		}
	}

	return insLDL
}
