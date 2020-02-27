package offer_2

import (
	"sync"
)

type singletonOnce struct{}

var insOnce *singletonOnce
var once sync.Once

func GetInsOnce() *singletonOnce {
	once.Do(func() {
		insOnce = &singletonOnce{}
	})
	return insOnce
}
