// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/davidgoodaragepi/api/light"
)

type FakeHandler struct {
	HandleGetStub        func(w http.ResponseWriter, r *http.Request)
	handleGetMutex       sync.RWMutex
	handleGetArgsForCall []struct {
		w http.ResponseWriter
		r *http.Request
	}
	HandleSetStub        func(w http.ResponseWriter, r *http.Request)
	handleSetMutex       sync.RWMutex
	handleSetArgsForCall []struct {
		w http.ResponseWriter
		r *http.Request
	}
	DiscoverLightStateStub        func() (*light.LightState, error)
	discoverLightStateMutex       sync.RWMutex
	discoverLightStateArgsForCall []struct{}
	discoverLightStateReturns     struct {
		result1 *light.LightState
		result2 error
	}
}

func (fake *FakeHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	fake.handleGetMutex.Lock()
	fake.handleGetArgsForCall = append(fake.handleGetArgsForCall, struct {
		w http.ResponseWriter
		r *http.Request
	}{w, r})
	fake.handleGetMutex.Unlock()
	if fake.HandleGetStub != nil {
		fake.HandleGetStub(w, r)
	}
}

func (fake *FakeHandler) HandleGetCallCount() int {
	fake.handleGetMutex.RLock()
	defer fake.handleGetMutex.RUnlock()
	return len(fake.handleGetArgsForCall)
}

func (fake *FakeHandler) HandleGetArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.handleGetMutex.RLock()
	defer fake.handleGetMutex.RUnlock()
	return fake.handleGetArgsForCall[i].w, fake.handleGetArgsForCall[i].r
}

func (fake *FakeHandler) HandleSet(w http.ResponseWriter, r *http.Request) {
	fake.handleSetMutex.Lock()
	fake.handleSetArgsForCall = append(fake.handleSetArgsForCall, struct {
		w http.ResponseWriter
		r *http.Request
	}{w, r})
	fake.handleSetMutex.Unlock()
	if fake.HandleSetStub != nil {
		fake.HandleSetStub(w, r)
	}
}

func (fake *FakeHandler) HandleSetCallCount() int {
	fake.handleSetMutex.RLock()
	defer fake.handleSetMutex.RUnlock()
	return len(fake.handleSetArgsForCall)
}

func (fake *FakeHandler) HandleSetArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.handleSetMutex.RLock()
	defer fake.handleSetMutex.RUnlock()
	return fake.handleSetArgsForCall[i].w, fake.handleSetArgsForCall[i].r
}

func (fake *FakeHandler) DiscoverLightState() (*light.LightState, error) {
	fake.discoverLightStateMutex.Lock()
	fake.discoverLightStateArgsForCall = append(fake.discoverLightStateArgsForCall, struct{}{})
	fake.discoverLightStateMutex.Unlock()
	if fake.DiscoverLightStateStub != nil {
		return fake.DiscoverLightStateStub()
	} else {
		return fake.discoverLightStateReturns.result1, fake.discoverLightStateReturns.result2
	}
}

func (fake *FakeHandler) DiscoverLightStateCallCount() int {
	fake.discoverLightStateMutex.RLock()
	defer fake.discoverLightStateMutex.RUnlock()
	return len(fake.discoverLightStateArgsForCall)
}

func (fake *FakeHandler) DiscoverLightStateReturns(result1 *light.LightState, result2 error) {
	fake.DiscoverLightStateStub = nil
	fake.discoverLightStateReturns = struct {
		result1 *light.LightState
		result2 error
	}{result1, result2}
}

var _ light.Handler = new(FakeHandler)
