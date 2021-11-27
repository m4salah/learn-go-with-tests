package contextt

import (
	"fmt"
	"net/http"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type StubStore struct {
	response string
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
func Testheelo() {
	ch := make(chan int)
	ch <- 7
	val := <-ch
	fmt.Println(val)
}
