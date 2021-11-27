package contextt

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "Hello, world!"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, world!"
		store := &SpyStore{response: data}
		svr := Server(&StubStore{data})
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}

func TestHeelo(t *testing.T) {
	ch := make(chan int)
	ch <- 7
	val := <-ch

	if val != 7 {
		t.Error("store was not told to cancel %v", val)
	}
	fmt.Println(val)
}
