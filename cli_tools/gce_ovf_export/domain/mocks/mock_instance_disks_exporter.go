// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_ovf_export/domain (interfaces: InstanceDisksExporter)

// Package mocks is a generated GoMock package.
package mocks

import (
	ovfexportdomain "github.com/GoogleCloudPlatform/compute-image-tools/cli_tools/gce_ovf_export/domain"
	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	reflect "reflect"
)

// MockInstanceDisksExporter is a mock of InstanceDisksExporter interface
type MockInstanceDisksExporter struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceDisksExporterMockRecorder
}

// MockInstanceDisksExporterMockRecorder is the mock recorder for MockInstanceDisksExporter
type MockInstanceDisksExporterMockRecorder struct {
	mock *MockInstanceDisksExporter
}

// NewMockInstanceDisksExporter creates a new mock instance
func NewMockInstanceDisksExporter(ctrl *gomock.Controller) *MockInstanceDisksExporter {
	mock := &MockInstanceDisksExporter{ctrl: ctrl}
	mock.recorder = &MockInstanceDisksExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstanceDisksExporter) EXPECT() *MockInstanceDisksExporterMockRecorder {
	return m.recorder
}

// Cancel mocks base method
func (m *MockInstanceDisksExporter) Cancel(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockInstanceDisksExporterMockRecorder) Cancel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockInstanceDisksExporter)(nil).Cancel), arg0)
}

// Export mocks base method
func (m *MockInstanceDisksExporter) Export(arg0 *compute.Instance, arg1 *ovfexportdomain.OVFExportArgs) ([]*ovfexportdomain.ExportedDisk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0, arg1)
	ret0, _ := ret[0].([]*ovfexportdomain.ExportedDisk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Export indicates an expected call of Export
func (mr *MockInstanceDisksExporterMockRecorder) Export(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockInstanceDisksExporter)(nil).Export), arg0, arg1)
}
