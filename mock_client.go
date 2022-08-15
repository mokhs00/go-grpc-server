// Code generated by MockGen. DO NOT EDIT.
// Source: protos/apis/v1/content/content_grpc.pb.go

// Package go_grpc_server is a generated GoMock package.
package go_grpc_server

import (
	context "context"
	go_grpc_server "go-grpc-server/protos/apis/v1/content"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockContentClient is a mock of ContentClient interface.
type MockContentClient struct {
	ctrl     *gomock.Controller
	recorder *MockContentClientMockRecorder
}

// MockContentClientMockRecorder is the mock recorder for MockContentClient.
type MockContentClientMockRecorder struct {
	mock *MockContentClient
}

// NewMockContentClient creates a new mock instance.
func NewMockContentClient(ctrl *gomock.Controller) *MockContentClient {
	mock := &MockContentClient{ctrl: ctrl}
	mock.recorder = &MockContentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContentClient) EXPECT() *MockContentClientMockRecorder {
	return m.recorder
}

// ListContent mocks base method.
func (m *MockContentClient) ListContent(ctx context.Context, in *go_grpc_server.ListContentRequest, opts ...grpc.CallOption) (*go_grpc_server.ListContentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContent", varargs...)
	ret0, _ := ret[0].(*go_grpc_server.ListContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContent indicates an expected call of ListContent.
func (mr *MockContentClientMockRecorder) ListContent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContent", reflect.TypeOf((*MockContentClient)(nil).ListContent), varargs...)
}

// MockContentServer is a mock of ContentServer interface.
type MockContentServer struct {
	ctrl     *gomock.Controller
	recorder *MockContentServerMockRecorder
}

// MockContentServerMockRecorder is the mock recorder for MockContentServer.
type MockContentServerMockRecorder struct {
	mock *MockContentServer
}

// NewMockContentServer creates a new mock instance.
func NewMockContentServer(ctrl *gomock.Controller) *MockContentServer {
	mock := &MockContentServer{ctrl: ctrl}
	mock.recorder = &MockContentServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContentServer) EXPECT() *MockContentServerMockRecorder {
	return m.recorder
}

// ListContent mocks base method.
func (m *MockContentServer) ListContent(arg0 context.Context, arg1 *go_grpc_server.ListContentRequest) (*go_grpc_server.ListContentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContent", arg0, arg1)
	ret0, _ := ret[0].(*go_grpc_server.ListContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContent indicates an expected call of ListContent.
func (mr *MockContentServerMockRecorder) ListContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContent", reflect.TypeOf((*MockContentServer)(nil).ListContent), arg0, arg1)
}

// mustEmbedUnimplementedContentServer mocks base method.
func (m *MockContentServer) mustEmbedUnimplementedContentServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContentServer")
}

// mustEmbedUnimplementedContentServer indicates an expected call of mustEmbedUnimplementedContentServer.
func (mr *MockContentServerMockRecorder) mustEmbedUnimplementedContentServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContentServer", reflect.TypeOf((*MockContentServer)(nil).mustEmbedUnimplementedContentServer))
}

// MockUnsafeContentServer is a mock of UnsafeContentServer interface.
type MockUnsafeContentServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeContentServerMockRecorder
}

// MockUnsafeContentServerMockRecorder is the mock recorder for MockUnsafeContentServer.
type MockUnsafeContentServerMockRecorder struct {
	mock *MockUnsafeContentServer
}

// NewMockUnsafeContentServer creates a new mock instance.
func NewMockUnsafeContentServer(ctrl *gomock.Controller) *MockUnsafeContentServer {
	mock := &MockUnsafeContentServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeContentServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeContentServer) EXPECT() *MockUnsafeContentServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedContentServer mocks base method.
func (m *MockUnsafeContentServer) mustEmbedUnimplementedContentServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedContentServer")
}

// mustEmbedUnimplementedContentServer indicates an expected call of mustEmbedUnimplementedContentServer.
func (mr *MockUnsafeContentServerMockRecorder) mustEmbedUnimplementedContentServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedContentServer", reflect.TypeOf((*MockUnsafeContentServer)(nil).mustEmbedUnimplementedContentServer))
}
