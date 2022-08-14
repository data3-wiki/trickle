// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dereference-xyz/trickle/node (interfaces: SolanaNode)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	rpc "github.com/gagliardetto/solana-go/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockSolanaNode is a mock of SolanaNode interface.
type MockSolanaNode struct {
	ctrl     *gomock.Controller
	recorder *MockSolanaNodeMockRecorder
}

// MockSolanaNodeMockRecorder is the mock recorder for MockSolanaNode.
type MockSolanaNodeMockRecorder struct {
	mock *MockSolanaNode
}

// NewMockSolanaNode creates a new mock instance.
func NewMockSolanaNode(ctrl *gomock.Controller) *MockSolanaNode {
	mock := &MockSolanaNode{ctrl: ctrl}
	mock.recorder = &MockSolanaNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSolanaNode) EXPECT() *MockSolanaNodeMockRecorder {
	return m.recorder
}

// GetProgramAccounts mocks base method.
func (m *MockSolanaNode) GetProgramAccounts(arg0 string) (rpc.GetProgramAccountsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProgramAccounts", arg0)
	ret0, _ := ret[0].(rpc.GetProgramAccountsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgramAccounts indicates an expected call of GetProgramAccounts.
func (mr *MockSolanaNodeMockRecorder) GetProgramAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgramAccounts", reflect.TypeOf((*MockSolanaNode)(nil).GetProgramAccounts), arg0)
}
