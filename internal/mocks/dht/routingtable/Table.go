// Code generated by mockery v2.35.2. DO NOT EDIT.

package routingtable

import (
	context "context"

	krpc "github.com/anacrolix/dht/v2/krpc"
	mock "github.com/stretchr/testify/mock"

	routingtable "github.com/bitmagnet-io/bitmagnet/internal/dht/routingtable"
)

// Table is an autogenerated mock type for the Table type
type Table struct {
	mock.Mock
}

type Table_Expecter struct {
	mock *mock.Mock
}

func (_m *Table) EXPECT() *Table_Expecter {
	return &Table_Expecter{mock: &_m.Mock}
}

// FindNode provides a mock function with given fields: _a0
func (_m *Table) FindNode(_a0 krpc.ID) krpc.CompactIPv4NodeInfo {
	ret := _m.Called(_a0)

	var r0 krpc.CompactIPv4NodeInfo
	if rf, ok := ret.Get(0).(func(krpc.ID) krpc.CompactIPv4NodeInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(krpc.CompactIPv4NodeInfo)
		}
	}

	return r0
}

// Table_FindNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindNode'
type Table_FindNode_Call struct {
	*mock.Call
}

// FindNode is a helper method to define mock.On call
//   - _a0 krpc.ID
func (_e *Table_Expecter) FindNode(_a0 interface{}) *Table_FindNode_Call {
	return &Table_FindNode_Call{Call: _e.mock.On("FindNode", _a0)}
}

func (_c *Table_FindNode_Call) Run(run func(_a0 krpc.ID)) *Table_FindNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(krpc.ID))
	})
	return _c
}

func (_c *Table_FindNode_Call) Return(_a0 krpc.CompactIPv4NodeInfo) *Table_FindNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Table_FindNode_Call) RunAndReturn(run func(krpc.ID) krpc.CompactIPv4NodeInfo) *Table_FindNode_Call {
	_c.Call.Return(run)
	return _c
}

// GetPeers provides a mock function with given fields: hash
func (_m *Table) GetPeers(hash krpc.ID) ([]krpc.NodeAddr, krpc.CompactIPv4NodeInfo) {
	ret := _m.Called(hash)

	var r0 []krpc.NodeAddr
	var r1 krpc.CompactIPv4NodeInfo
	if rf, ok := ret.Get(0).(func(krpc.ID) ([]krpc.NodeAddr, krpc.CompactIPv4NodeInfo)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(krpc.ID) []krpc.NodeAddr); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]krpc.NodeAddr)
		}
	}

	if rf, ok := ret.Get(1).(func(krpc.ID) krpc.CompactIPv4NodeInfo); ok {
		r1 = rf(hash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(krpc.CompactIPv4NodeInfo)
		}
	}

	return r0, r1
}

// Table_GetPeers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPeers'
type Table_GetPeers_Call struct {
	*mock.Call
}

// GetPeers is a helper method to define mock.On call
//   - hash krpc.ID
func (_e *Table_Expecter) GetPeers(hash interface{}) *Table_GetPeers_Call {
	return &Table_GetPeers_Call{Call: _e.mock.On("GetPeers", hash)}
}

func (_c *Table_GetPeers_Call) Run(run func(hash krpc.ID)) *Table_GetPeers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(krpc.ID))
	})
	return _c
}

func (_c *Table_GetPeers_Call) Return(_a0 []krpc.NodeAddr, _a1 krpc.CompactIPv4NodeInfo) *Table_GetPeers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Table_GetPeers_Call) RunAndReturn(run func(krpc.ID) ([]krpc.NodeAddr, krpc.CompactIPv4NodeInfo)) *Table_GetPeers_Call {
	_c.Call.Return(run)
	return _c
}

// ReceiveNodeAddr provides a mock function with given fields: _a0
func (_m *Table) ReceiveNodeAddr(_a0 ...krpc.NodeAddr) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Table_ReceiveNodeAddr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReceiveNodeAddr'
type Table_ReceiveNodeAddr_Call struct {
	*mock.Call
}

// ReceiveNodeAddr is a helper method to define mock.On call
//   - _a0 ...krpc.NodeAddr
func (_e *Table_Expecter) ReceiveNodeAddr(_a0 ...interface{}) *Table_ReceiveNodeAddr_Call {
	return &Table_ReceiveNodeAddr_Call{Call: _e.mock.On("ReceiveNodeAddr",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Table_ReceiveNodeAddr_Call) Run(run func(_a0 ...krpc.NodeAddr)) *Table_ReceiveNodeAddr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]krpc.NodeAddr, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(krpc.NodeAddr)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Table_ReceiveNodeAddr_Call) Return() *Table_ReceiveNodeAddr_Call {
	_c.Call.Return()
	return _c
}

func (_c *Table_ReceiveNodeAddr_Call) RunAndReturn(run func(...krpc.NodeAddr)) *Table_ReceiveNodeAddr_Call {
	_c.Call.Return(run)
	return _c
}

// ReceiveNodeInfo provides a mock function with given fields: _a0
func (_m *Table) ReceiveNodeInfo(_a0 ...krpc.NodeInfo) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Table_ReceiveNodeInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReceiveNodeInfo'
type Table_ReceiveNodeInfo_Call struct {
	*mock.Call
}

// ReceiveNodeInfo is a helper method to define mock.On call
//   - _a0 ...krpc.NodeInfo
func (_e *Table_Expecter) ReceiveNodeInfo(_a0 ...interface{}) *Table_ReceiveNodeInfo_Call {
	return &Table_ReceiveNodeInfo_Call{Call: _e.mock.On("ReceiveNodeInfo",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Table_ReceiveNodeInfo_Call) Run(run func(_a0 ...krpc.NodeInfo)) *Table_ReceiveNodeInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]krpc.NodeInfo, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(krpc.NodeInfo)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Table_ReceiveNodeInfo_Call) Return() *Table_ReceiveNodeInfo_Call {
	_c.Call.Return()
	return _c
}

func (_c *Table_ReceiveNodeInfo_Call) RunAndReturn(run func(...krpc.NodeInfo)) *Table_ReceiveNodeInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ReceivePeersForHash provides a mock function with given fields: hash, addrs
func (_m *Table) ReceivePeersForHash(hash krpc.ID, addrs ...krpc.NodeAddr) {
	_va := make([]interface{}, len(addrs))
	for _i := range addrs {
		_va[_i] = addrs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, hash)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Table_ReceivePeersForHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReceivePeersForHash'
type Table_ReceivePeersForHash_Call struct {
	*mock.Call
}

// ReceivePeersForHash is a helper method to define mock.On call
//   - hash krpc.ID
//   - addrs ...krpc.NodeAddr
func (_e *Table_Expecter) ReceivePeersForHash(hash interface{}, addrs ...interface{}) *Table_ReceivePeersForHash_Call {
	return &Table_ReceivePeersForHash_Call{Call: _e.mock.On("ReceivePeersForHash",
		append([]interface{}{hash}, addrs...)...)}
}

func (_c *Table_ReceivePeersForHash_Call) Run(run func(hash krpc.ID, addrs ...krpc.NodeAddr)) *Table_ReceivePeersForHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]krpc.NodeAddr, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(krpc.NodeAddr)
			}
		}
		run(args[0].(krpc.ID), variadicArgs...)
	})
	return _c
}

func (_c *Table_ReceivePeersForHash_Call) Return() *Table_ReceivePeersForHash_Call {
	_c.Call.Return()
	return _c
}

func (_c *Table_ReceivePeersForHash_Call) RunAndReturn(run func(krpc.ID, ...krpc.NodeAddr)) *Table_ReceivePeersForHash_Call {
	_c.Call.Return(run)
	return _c
}

// SampleInfoHashes provides a mock function with given fields:
func (_m *Table) SampleInfoHashes() (krpc.CompactInfohashes, krpc.CompactIPv4NodeInfo, int64) {
	ret := _m.Called()

	var r0 krpc.CompactInfohashes
	var r1 krpc.CompactIPv4NodeInfo
	var r2 int64
	if rf, ok := ret.Get(0).(func() (krpc.CompactInfohashes, krpc.CompactIPv4NodeInfo, int64)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() krpc.CompactInfohashes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(krpc.CompactInfohashes)
		}
	}

	if rf, ok := ret.Get(1).(func() krpc.CompactIPv4NodeInfo); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(krpc.CompactIPv4NodeInfo)
		}
	}

	if rf, ok := ret.Get(2).(func() int64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int64)
	}

	return r0, r1, r2
}

// Table_SampleInfoHashes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SampleInfoHashes'
type Table_SampleInfoHashes_Call struct {
	*mock.Call
}

// SampleInfoHashes is a helper method to define mock.On call
func (_e *Table_Expecter) SampleInfoHashes() *Table_SampleInfoHashes_Call {
	return &Table_SampleInfoHashes_Call{Call: _e.mock.On("SampleInfoHashes")}
}

func (_c *Table_SampleInfoHashes_Call) Run(run func()) *Table_SampleInfoHashes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Table_SampleInfoHashes_Call) Return(_a0 krpc.CompactInfohashes, _a1 krpc.CompactIPv4NodeInfo, _a2 int64) *Table_SampleInfoHashes_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Table_SampleInfoHashes_Call) RunAndReturn(run func() (krpc.CompactInfohashes, krpc.CompactIPv4NodeInfo, int64)) *Table_SampleInfoHashes_Call {
	_c.Call.Return(run)
	return _c
}

// TryEachNode provides a mock function with given fields: _a0, _a1
func (_m *Table) TryEachNode(_a0 context.Context, _a1 func(context.Context, routingtable.PeerInfo) error) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, routingtable.PeerInfo) error) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Table_TryEachNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TryEachNode'
type Table_TryEachNode_Call struct {
	*mock.Call
}

// TryEachNode is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 func(context.Context , routingtable.PeerInfo) error
func (_e *Table_Expecter) TryEachNode(_a0 interface{}, _a1 interface{}) *Table_TryEachNode_Call {
	return &Table_TryEachNode_Call{Call: _e.mock.On("TryEachNode", _a0, _a1)}
}

func (_c *Table_TryEachNode_Call) Run(run func(_a0 context.Context, _a1 func(context.Context, routingtable.PeerInfo) error)) *Table_TryEachNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context, routingtable.PeerInfo) error))
	})
	return _c
}

func (_c *Table_TryEachNode_Call) Return(_a0 error) *Table_TryEachNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Table_TryEachNode_Call) RunAndReturn(run func(context.Context, func(context.Context, routingtable.PeerInfo) error) error) *Table_TryEachNode_Call {
	_c.Call.Return(run)
	return _c
}

// WithPeer provides a mock function with given fields: _a0, _a1, _a2
func (_m *Table) WithPeer(_a0 context.Context, _a1 krpc.NodeAddr, _a2 func(context.Context) error) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, krpc.NodeAddr, func(context.Context) error) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Table_WithPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithPeer'
type Table_WithPeer_Call struct {
	*mock.Call
}

// WithPeer is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 krpc.NodeAddr
//   - _a2 func(context.Context) error
func (_e *Table_Expecter) WithPeer(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Table_WithPeer_Call {
	return &Table_WithPeer_Call{Call: _e.mock.On("WithPeer", _a0, _a1, _a2)}
}

func (_c *Table_WithPeer_Call) Run(run func(_a0 context.Context, _a1 krpc.NodeAddr, _a2 func(context.Context) error)) *Table_WithPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(krpc.NodeAddr), args[2].(func(context.Context) error))
	})
	return _c
}

func (_c *Table_WithPeer_Call) Return(_a0 error) *Table_WithPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Table_WithPeer_Call) RunAndReturn(run func(context.Context, krpc.NodeAddr, func(context.Context) error) error) *Table_WithPeer_Call {
	_c.Call.Return(run)
	return _c
}

// NewTable creates a new instance of Table. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTable(t interface {
	mock.TestingT
	Cleanup(func())
}) *Table {
	mock := &Table{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
