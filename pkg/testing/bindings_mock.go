package testing

import (
	"github.com/dapr/components-contrib/bindings"
	mock "github.com/stretchr/testify/mock"
)

// MockBinding is a mock input/output component object
type MockBinding struct {
	mock.Mock
}

// Init is a mock initialization method
func (m *MockBinding) Init(metadata bindings.Metadata) error {
	return nil
}

// Read is a mock read method
func (m *MockBinding) Read(handler func(*bindings.ReadResponse) error) error {
	args := m.Called(handler)
	return args.Error(0)
}

// Invoke is a mock invoke method
func (m *MockBinding) Invoke(req *bindings.InvokeRequest) (*bindings.InvokeResponse, error) {
	args := m.Called(req)
	return nil, args.Error(0)
}

// Operations is a mock operations method
func (m *MockBinding) Operations() []bindings.OperationKind {
	return []bindings.OperationKind{bindings.CreateOperation}
}
