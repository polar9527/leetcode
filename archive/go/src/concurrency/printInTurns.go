package test

import (
	"fmt"
	"sync"
)

// 给定用3个 goroutine 依次轮流打印 N 个数，并依次结束打印

// @lc code=start
func printInTurns(count int) {
	var wg sync.WaitGroup
	defer wg.Wait()
	worker := 3
	wg.Add(worker)

	ch := make([]chan int, worker)
	for i := range ch {
		ch[i] = make(chan int, 1)
	}

	chEnd := make([]chan interface{}, worker)
	for i := range ch {
		chEnd[i] = make(chan interface{}, 1)
	}

	for index := 0; index < worker; index++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case num := <-ch[i]:
					if num > count {
						chEnd[(i+1)%3] <- struct{}{}
						fmt.Printf("worker %d exit\n", i)
						return
					}
					fmt.Printf("worker %d process num %d\n", i, num)
					ch[(i+1)%3] <- num + 1
				case <-chEnd[i]:
					fmt.Printf("worker %d exit\n", i)
					chEnd[(i+1)%3] <- struct{}{}
					return
				}
			}
		}(index)
	}
	// fire in the hole
	go func() {
		ch[0] <- 1
	}()
}

// @lc code=end
