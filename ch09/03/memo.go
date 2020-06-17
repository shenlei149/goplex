package memo

import (
	"errors"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(string, <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string, done <-chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		cancel := make(chan struct{})
		doneByF := make(chan struct{})

		go func() {
			e.res.value, e.res.err = memo.f(key, cancel)
			doneByF <- struct{}{}
		}()

		select {
		case <-done:
			cancel <- struct{}{}
			delete(memo.cache, key)
			e = nil
		case <-doneByF:
		}
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		select {
		case <-done:
			e = nil
		case <-e.ready: // wait for ready condition
		}
	}

	if e != nil {
		return e.res.value, e.res.err
	}

	return nil, errors.New("Cancelled")
}
