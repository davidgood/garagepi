// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/davidgood/garagepi/web/homepage"
)

type FakeHandler struct {
	HandleStub        func(w http.ResponseWriter, r *http.Request)
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
		w http.ResponseWriter
		r *http.Request
	}
}

func (fake *FakeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
		w http.ResponseWriter
		r *http.Request
	}{w, r})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		fake.HandleStub(w, r)
	}
}

func (fake *FakeHandler) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeHandler) HandleArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return fake.handleArgsForCall[i].w, fake.handleArgsForCall[i].r
}

var _ homepage.Handler = new(FakeHandler)
