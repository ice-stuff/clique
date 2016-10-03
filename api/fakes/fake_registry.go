// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/ice-stuff/clique/api"
)

type FakeRegistry struct {
	TransfersByStateStub        func(state api.TransferState) []api.Transfer
	transfersByStateMutex       sync.RWMutex
	transfersByStateArgsForCall []struct {
		state api.TransferState
	}
	transfersByStateReturns struct {
		result1 []api.Transfer
	}
	TransferResultsStub        func() []api.TransferResults
	transferResultsMutex       sync.RWMutex
	transferResultsArgsForCall []struct{}
	transferResultsReturns     struct {
		result1 []api.TransferResults
	}
	TransferResultsByIPStub        func(net.IP) []api.TransferResults
	transferResultsByIPMutex       sync.RWMutex
	transferResultsByIPArgsForCall []struct {
		arg1 net.IP
	}
	transferResultsByIPReturns struct {
		result1 []api.TransferResults
	}
}

func (fake *FakeRegistry) TransfersByState(state api.TransferState) []api.Transfer {
	fake.transfersByStateMutex.Lock()
	fake.transfersByStateArgsForCall = append(fake.transfersByStateArgsForCall, struct {
		state api.TransferState
	}{state})
	fake.transfersByStateMutex.Unlock()
	if fake.TransfersByStateStub != nil {
		return fake.TransfersByStateStub(state)
	} else {
		return fake.transfersByStateReturns.result1
	}
}

func (fake *FakeRegistry) TransfersByStateCallCount() int {
	fake.transfersByStateMutex.RLock()
	defer fake.transfersByStateMutex.RUnlock()
	return len(fake.transfersByStateArgsForCall)
}

func (fake *FakeRegistry) TransfersByStateArgsForCall(i int) api.TransferState {
	fake.transfersByStateMutex.RLock()
	defer fake.transfersByStateMutex.RUnlock()
	return fake.transfersByStateArgsForCall[i].state
}

func (fake *FakeRegistry) TransfersByStateReturns(result1 []api.Transfer) {
	fake.TransfersByStateStub = nil
	fake.transfersByStateReturns = struct {
		result1 []api.Transfer
	}{result1}
}

func (fake *FakeRegistry) TransferResults() []api.TransferResults {
	fake.transferResultsMutex.Lock()
	fake.transferResultsArgsForCall = append(fake.transferResultsArgsForCall, struct{}{})
	fake.transferResultsMutex.Unlock()
	if fake.TransferResultsStub != nil {
		return fake.TransferResultsStub()
	} else {
		return fake.transferResultsReturns.result1
	}
}

func (fake *FakeRegistry) TransferResultsCallCount() int {
	fake.transferResultsMutex.RLock()
	defer fake.transferResultsMutex.RUnlock()
	return len(fake.transferResultsArgsForCall)
}

func (fake *FakeRegistry) TransferResultsReturns(result1 []api.TransferResults) {
	fake.TransferResultsStub = nil
	fake.transferResultsReturns = struct {
		result1 []api.TransferResults
	}{result1}
}

func (fake *FakeRegistry) TransferResultsByIP(arg1 net.IP) []api.TransferResults {
	fake.transferResultsByIPMutex.Lock()
	fake.transferResultsByIPArgsForCall = append(fake.transferResultsByIPArgsForCall, struct {
		arg1 net.IP
	}{arg1})
	fake.transferResultsByIPMutex.Unlock()
	if fake.TransferResultsByIPStub != nil {
		return fake.TransferResultsByIPStub(arg1)
	} else {
		return fake.transferResultsByIPReturns.result1
	}
}

func (fake *FakeRegistry) TransferResultsByIPCallCount() int {
	fake.transferResultsByIPMutex.RLock()
	defer fake.transferResultsByIPMutex.RUnlock()
	return len(fake.transferResultsByIPArgsForCall)
}

func (fake *FakeRegistry) TransferResultsByIPArgsForCall(i int) net.IP {
	fake.transferResultsByIPMutex.RLock()
	defer fake.transferResultsByIPMutex.RUnlock()
	return fake.transferResultsByIPArgsForCall[i].arg1
}

func (fake *FakeRegistry) TransferResultsByIPReturns(result1 []api.TransferResults) {
	fake.TransferResultsByIPStub = nil
	fake.transferResultsByIPReturns = struct {
		result1 []api.TransferResults
	}{result1}
}

var _ api.Registry = new(FakeRegistry)
