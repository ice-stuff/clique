// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/ice-stuff/clique/transfer"
)

type FakeTransferReceiver struct {
	ReceiveTransferStub        func(conn io.ReadWriter) (transfer.TransferResults, error)
	receiveTransferMutex       sync.RWMutex
	receiveTransferArgsForCall []struct {
		conn io.ReadWriter
	}
	receiveTransferReturns struct {
		result1 transfer.TransferResults
		result2 error
	}
}

func (fake *FakeTransferReceiver) ReceiveTransfer(conn io.ReadWriter) (transfer.TransferResults, error) {
	fake.receiveTransferMutex.Lock()
	fake.receiveTransferArgsForCall = append(fake.receiveTransferArgsForCall, struct {
		conn io.ReadWriter
	}{conn})
	fake.receiveTransferMutex.Unlock()
	if fake.ReceiveTransferStub != nil {
		return fake.ReceiveTransferStub(conn)
	} else {
		return fake.receiveTransferReturns.result1, fake.receiveTransferReturns.result2
	}
}

func (fake *FakeTransferReceiver) ReceiveTransferCallCount() int {
	fake.receiveTransferMutex.RLock()
	defer fake.receiveTransferMutex.RUnlock()
	return len(fake.receiveTransferArgsForCall)
}

func (fake *FakeTransferReceiver) ReceiveTransferArgsForCall(i int) io.ReadWriter {
	fake.receiveTransferMutex.RLock()
	defer fake.receiveTransferMutex.RUnlock()
	return fake.receiveTransferArgsForCall[i].conn
}

func (fake *FakeTransferReceiver) ReceiveTransferReturns(result1 transfer.TransferResults, result2 error) {
	fake.ReceiveTransferStub = nil
	fake.receiveTransferReturns = struct {
		result1 transfer.TransferResults
		result2 error
	}{result1, result2}
}

var _ transfer.TransferReceiver = new(FakeTransferReceiver)
