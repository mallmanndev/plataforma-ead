// Code generated by MockGen. DO NOT EDIT.
// Source: files_service.go

// Package mock_ports is a generated GoMock package.
package mocks

import (
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
        ports "github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

// MockFilesService is a mock of FilesService interface.
type MockFilesService struct {
        ctrl     *gomock.Controller
        recorder *MockFilesServiceMockRecorder
}

// MockFilesServiceMockRecorder is the mock recorder for MockFilesService.
type MockFilesServiceMockRecorder struct {
        mock *MockFilesService
}

// NewMockFilesService creates a new mock instance.
func NewMockFilesService(ctrl *gomock.Controller) *MockFilesService {
        mock := &MockFilesService{ctrl: ctrl}
        mock.recorder = &MockFilesServiceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilesService) EXPECT() *MockFilesServiceMockRecorder {
        return m.recorder
}

// Close mocks base method.
func (m *MockFilesService) Close() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Close")
        ret0, _ := ret[0].(error)
        return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFilesServiceMockRecorder) Close() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFilesService)(nil).Close))
}

// CreateFile mocks base method.
func (m *MockFilesService) CreateFile(File ports.FileInfo) (ports.FilesService, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateFile", File)
        ret0, _ := ret[0].(ports.FilesService)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockFilesServiceMockRecorder) CreateFile(File interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockFilesService)(nil).CreateFile), File)
}

// GetResolution mocks base method.
func (m *MockFilesService) GetResolution(Url string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetResolution", Url)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetResolution indicates an expected call of GetResolution.
func (mr *MockFilesServiceMockRecorder) GetResolution(Url interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolution", reflect.TypeOf((*MockFilesService)(nil).GetResolution), Url)
}

// ProcessVideo mocks base method.
func (m *MockFilesService) ProcessVideo(InputUrl, OutputPath, Resolution string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ProcessVideo", InputUrl, OutputPath, Resolution)
        ret0, _ := ret[0].(error)
        return ret0
}

// ProcessVideo indicates an expected call of ProcessVideo.
func (mr *MockFilesServiceMockRecorder) ProcessVideo(InputUrl, OutputPath, Resolution interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessVideo", reflect.TypeOf((*MockFilesService)(nil).ProcessVideo), InputUrl, OutputPath, Resolution)
}

// Remove mocks base method.
func (m *MockFilesService) Remove() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Remove")
        ret0, _ := ret[0].(error)
        return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockFilesServiceMockRecorder) Remove() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFilesService)(nil).Remove))
}

// SendChunk mocks base method.
func (m *MockFilesService) SendChunk(chunk []byte) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendChunk", chunk)
        ret0, _ := ret[0].(error)
        return ret0
}

// SendChunk indicates an expected call of SendChunk.
func (mr *MockFilesServiceMockRecorder) SendChunk(chunk interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChunk", reflect.TypeOf((*MockFilesService)(nil).SendChunk), chunk)
}

// WriteString mocks base method.
func (m *MockFilesService) WriteString(content string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "WriteString", content)
        ret0, _ := ret[0].(error)
        return ret0
}

// WriteString indicates an expected call of WriteString.
func (mr *MockFilesServiceMockRecorder) WriteString(content interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockFilesService)(nil).WriteString), content)
}
