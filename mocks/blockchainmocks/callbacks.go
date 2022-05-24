// Code generated by mockery v1.0.0. DO NOT EDIT.

package blockchainmocks

import (
	blockchain "github.com/hyperledger/firefly/pkg/blockchain"
	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// BatchPinComplete provides a mock function with given fields: batch, signingKey
func (_m *Callbacks) BatchPinComplete(batch *blockchain.BatchPin, signingKey *core.VerifierRef) error {
	ret := _m.Called(batch, signingKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blockchain.BatchPin, *core.VerifierRef) error); ok {
		r0 = rf(batch, signingKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockchainEvent provides a mock function with given fields: event
func (_m *Callbacks) BlockchainEvent(event *blockchain.EventWithSubscription) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blockchain.EventWithSubscription) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockchainOpUpdate provides a mock function with given fields: plugin, operationID, txState, blockchainTXID, errorMessage, opOutput
func (_m *Callbacks) BlockchainOpUpdate(plugin blockchain.Plugin, operationID *fftypes.UUID, txState core.OpStatus, blockchainTXID string, errorMessage string, opOutput fftypes.JSONObject) {
	_m.Called(plugin, operationID, txState, blockchainTXID, errorMessage, opOutput)
}

// BlockchainOperatorAction provides a mock function with given fields: action, event, signingKey
func (_m *Callbacks) BlockchainOperatorAction(action string, event *blockchain.Event, signingKey *core.VerifierRef) error {
	ret := _m.Called(action, event, signingKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *blockchain.Event, *core.VerifierRef) error); ok {
		r0 = rf(action, event, signingKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
