// Code generated by MockGen. DO NOT EDIT.
// Source: api/accountservice/v1/service.pb.go

// Package apimock is a generated GoMock package.
package apimock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/temporalio/saas-proto/protogen/api/accountservice/v1"
	grpc "google.golang.org/grpc"
)

// MockAccountServiceClient is a mock of AccountServiceClient interface.
type MockAccountServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceClientMockRecorder
}

// MockAccountServiceClientMockRecorder is the mock recorder for MockAccountServiceClient.
type MockAccountServiceClientMockRecorder struct {
	mock *MockAccountServiceClient
}

// NewMockAccountServiceClient creates a new mock instance.
func NewMockAccountServiceClient(ctrl *gomock.Controller) *MockAccountServiceClient {
	mock := &MockAccountServiceClient{ctrl: ctrl}
	mock.recorder = &MockAccountServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceClient) EXPECT() *MockAccountServiceClientMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountServiceClient) GetAccount(ctx context.Context, in *v1.GetAccountRequest, opts ...grpc.CallOption) (*v1.GetAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*v1.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountServiceClientMockRecorder) GetAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountServiceClient)(nil).GetAccount), varargs...)
}

// GetAccountUsage mocks base method.
func (m *MockAccountServiceClient) GetAccountUsage(ctx context.Context, in *v1.GetAccountUsageRequest, opts ...grpc.CallOption) (*v1.GetAccountUsageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountUsage", varargs...)
	ret0, _ := ret[0].(*v1.GetAccountUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountUsage indicates an expected call of GetAccountUsage.
func (mr *MockAccountServiceClientMockRecorder) GetAccountUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountUsage", reflect.TypeOf((*MockAccountServiceClient)(nil).GetAccountUsage), varargs...)
}

// GetNamespaceUsage mocks base method.
func (m *MockAccountServiceClient) GetNamespaceUsage(ctx context.Context, in *v1.GetNamespaceUsageRequest, opts ...grpc.CallOption) (*v1.GetNamespaceUsageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamespaceUsage", varargs...)
	ret0, _ := ret[0].(*v1.GetNamespaceUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceUsage indicates an expected call of GetNamespaceUsage.
func (mr *MockAccountServiceClientMockRecorder) GetNamespaceUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceUsage", reflect.TypeOf((*MockAccountServiceClient)(nil).GetNamespaceUsage), varargs...)
}

// GetNamespacesUsage mocks base method.
func (m *MockAccountServiceClient) GetNamespacesUsage(ctx context.Context, in *v1.GetNamespacesUsageRequest, opts ...grpc.CallOption) (*v1.GetNamespacesUsageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamespacesUsage", varargs...)
	ret0, _ := ret[0].(*v1.GetNamespacesUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespacesUsage indicates an expected call of GetNamespacesUsage.
func (mr *MockAccountServiceClientMockRecorder) GetNamespacesUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespacesUsage", reflect.TypeOf((*MockAccountServiceClient)(nil).GetNamespacesUsage), varargs...)
}

// GetRegions mocks base method.
func (m *MockAccountServiceClient) GetRegions(ctx context.Context, in *v1.GetRegionsRequest, opts ...grpc.CallOption) (*v1.GetRegionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegions", varargs...)
	ret0, _ := ret[0].(*v1.GetRegionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegions indicates an expected call of GetRegions.
func (mr *MockAccountServiceClientMockRecorder) GetRegions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegions", reflect.TypeOf((*MockAccountServiceClient)(nil).GetRegions), varargs...)
}

// UpdateAccount mocks base method.
func (m *MockAccountServiceClient) UpdateAccount(ctx context.Context, in *v1.UpdateAccountRequest, opts ...grpc.CallOption) (*v1.UpdateAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAccount", varargs...)
	ret0, _ := ret[0].(*v1.UpdateAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountServiceClientMockRecorder) UpdateAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountServiceClient)(nil).UpdateAccount), varargs...)
}

// MockAccountServiceServer is a mock of AccountServiceServer interface.
type MockAccountServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceServerMockRecorder
}

// MockAccountServiceServerMockRecorder is the mock recorder for MockAccountServiceServer.
type MockAccountServiceServerMockRecorder struct {
	mock *MockAccountServiceServer
}

// NewMockAccountServiceServer creates a new mock instance.
func NewMockAccountServiceServer(ctrl *gomock.Controller) *MockAccountServiceServer {
	mock := &MockAccountServiceServer{ctrl: ctrl}
	mock.recorder = &MockAccountServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceServer) EXPECT() *MockAccountServiceServerMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountServiceServer) GetAccount(arg0 context.Context, arg1 *v1.GetAccountRequest) (*v1.GetAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountServiceServerMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountServiceServer)(nil).GetAccount), arg0, arg1)
}

// GetAccountUsage mocks base method.
func (m *MockAccountServiceServer) GetAccountUsage(arg0 context.Context, arg1 *v1.GetAccountUsageRequest) (*v1.GetAccountUsageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountUsage", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetAccountUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountUsage indicates an expected call of GetAccountUsage.
func (mr *MockAccountServiceServerMockRecorder) GetAccountUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountUsage", reflect.TypeOf((*MockAccountServiceServer)(nil).GetAccountUsage), arg0, arg1)
}

// GetNamespaceUsage mocks base method.
func (m *MockAccountServiceServer) GetNamespaceUsage(arg0 context.Context, arg1 *v1.GetNamespaceUsageRequest) (*v1.GetNamespaceUsageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceUsage", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetNamespaceUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceUsage indicates an expected call of GetNamespaceUsage.
func (mr *MockAccountServiceServerMockRecorder) GetNamespaceUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceUsage", reflect.TypeOf((*MockAccountServiceServer)(nil).GetNamespaceUsage), arg0, arg1)
}

// GetNamespacesUsage mocks base method.
func (m *MockAccountServiceServer) GetNamespacesUsage(arg0 context.Context, arg1 *v1.GetNamespacesUsageRequest) (*v1.GetNamespacesUsageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespacesUsage", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetNamespacesUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespacesUsage indicates an expected call of GetNamespacesUsage.
func (mr *MockAccountServiceServerMockRecorder) GetNamespacesUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespacesUsage", reflect.TypeOf((*MockAccountServiceServer)(nil).GetNamespacesUsage), arg0, arg1)
}

// GetRegions mocks base method.
func (m *MockAccountServiceServer) GetRegions(arg0 context.Context, arg1 *v1.GetRegionsRequest) (*v1.GetRegionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegions", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetRegionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegions indicates an expected call of GetRegions.
func (mr *MockAccountServiceServerMockRecorder) GetRegions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegions", reflect.TypeOf((*MockAccountServiceServer)(nil).GetRegions), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockAccountServiceServer) UpdateAccount(arg0 context.Context, arg1 *v1.UpdateAccountRequest) (*v1.UpdateAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(*v1.UpdateAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountServiceServerMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountServiceServer)(nil).UpdateAccount), arg0, arg1)
}
