package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type call struct {
	wg  sync.WaitGroup
	val any
	err error
}
type singleflight struct {
	mu sync.Mutex
	m  map[string]*call
}

func (s *singleflight) Do(key string, action func() (any, error)) (any, error) {
	s.mu.Lock()
	if s.m == nil {
		s.m = make(map[string]*call)
	}

	res, ok := s.m[key]

	if ok {
		s.mu.Unlock()
		res.wg.Wait()
		return res.val, res.err
	}

	res = new(call)
	res.wg.Add(1)
	s.m[key] = res
	s.mu.Unlock()

	defer func() {
		res.wg.Done()

		s.mu.Lock()
		delete(s.m, key)
		s.mu.Unlock()
	}()

	res.val, res.err = action()
	return res.val, res.err

}

func getRandomNumber() (any, error) {
	number := rand.Intn(10)
	time.Sleep(time.Second * 3)
	if number == 0 {
		return nil, errors.New("the number is zero")
	}
	return number, nil
}
func main() {
	var s singleflight
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			res, err := s.Do("key", getRandomNumber)
			fmt.Printf("result %v, error %v\n", res, err)
		}()
	}
	wg.Wait()
}
