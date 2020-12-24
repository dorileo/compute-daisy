// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_ovf_export/domain (interfaces: OvfDescriptorGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	ovfexportdomain "github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_ovf_export/domain"
	pb "github.com/GoogleCloudPlatform/compute-image-tools/proto/go/pb"
	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	reflect "reflect"
)

// MockOvfDescriptorGenerator is a mock of OvfDescriptorGenerator interface
type MockOvfDescriptorGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockOvfDescriptorGeneratorMockRecorder
}

// MockOvfDescriptorGeneratorMockRecorder is the mock recorder for MockOvfDescriptorGenerator
type MockOvfDescriptorGeneratorMockRecorder struct {
	mock *MockOvfDescriptorGenerator
}

// NewMockOvfDescriptorGenerator creates a new mock instance
func NewMockOvfDescriptorGenerator(ctrl *gomock.Controller) *MockOvfDescriptorGenerator {
	mock := &MockOvfDescriptorGenerator{ctrl: ctrl}
	mock.recorder = &MockOvfDescriptorGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOvfDescriptorGenerator) EXPECT() *MockOvfDescriptorGeneratorMockRecorder {
	return m.recorder
}

// Cancel mocks base method
func (m *MockOvfDescriptorGenerator) Cancel(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockOvfDescriptorGeneratorMockRecorder) Cancel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockOvfDescriptorGenerator)(nil).Cancel), arg0)
}

// GenerateAndWriteOVFDescriptor mocks base method
func (m *MockOvfDescriptorGenerator) GenerateAndWriteOVFDescriptor(arg0 *compute.Instance, arg1 []*ovfexportdomain.ExportedDisk, arg2, arg3, arg4 string, arg5 *pb.InspectionResults) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAndWriteOVFDescriptor", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateAndWriteOVFDescriptor indicates an expected call of GenerateAndWriteOVFDescriptor
func (mr *MockOvfDescriptorGeneratorMockRecorder) GenerateAndWriteOVFDescriptor(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAndWriteOVFDescriptor", reflect.TypeOf((*MockOvfDescriptorGenerator)(nil).GenerateAndWriteOVFDescriptor), arg0, arg1, arg2, arg3, arg4, arg5)
}
