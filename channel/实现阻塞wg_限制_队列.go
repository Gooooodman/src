package main

import (
	"fmt"
	"sync"
	"time"
)

//LimitRate 限速
type LimitRate struct {
	rate       int
	interval   time.Duration
	lastAction time.Time
	lock       sync.Mutex
}

//Limit 限速
func (l *LimitRate) Limit() bool {
	result := false
	for {
		l.lock.Lock()
		//判断最后一次执行的时间与当前的时间间隔是否大于限速速率
		if time.Now().Sub(l.lastAction) > l.interval {
			l.lastAction = time.Now()
			result = true
		}
		l.lock.Unlock()
		if result {
			return result
		}
		time.Sleep(l.interval)
	}
}

//SetRate 设置Rate
func (l *LimitRate) SetRate(r int) {
	l.rate = r
	l.interval = time.Microsecond * time.Duration(1000*1000/l.rate)
}

//GetRate 获取Rate
func (l *LimitRate) GetRate() int {
	return l.rate
}

func main() {
	var wg sync.WaitGroup
	var lr LimitRate
	lr.SetRate(15) // 限制并发数

	b := time.Now()
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(i int) {
			if lr.Limit() {
				//业务代码
				a := 1
				for i := 0; i < 1000000000; i++ {
					a += i
				}
				fmt.Println(i, ": Counting", a)
				//业务代码
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(b))
}
