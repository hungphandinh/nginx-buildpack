// Code generated by MockGen. DO NOT EDIT.
// Source: supply.go

package supply_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return _m.recorder
}

// InstallOnlyVersion mocks base method
func (_m *MockManifest) InstallOnlyVersion(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "InstallOnlyVersion", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallOnlyVersion indicates an expected call of InstallOnlyVersion
func (_mr *MockManifestMockRecorder) InstallOnlyVersion(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallOnlyVersion", reflect.TypeOf((*MockManifest)(nil).InstallOnlyVersion), arg0, arg1)
}

// DefaultVersion mocks base method
func (_m *MockManifest) DefaultVersion(depName string) (libbuildpack.Dependency, error) {
	ret := _m.ctrl.Call(_m, "DefaultVersion", depName)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (_mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DefaultVersion", reflect.TypeOf((*MockManifest)(nil).DefaultVersion), arg0)
}

// AllDependencyVersions mocks base method
func (_m *MockManifest) AllDependencyVersions(_param0 string) []string {
	ret := _m.ctrl.Call(_m, "AllDependencyVersions", _param0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllDependencyVersions indicates an expected call of AllDependencyVersions
func (_mr *MockManifestMockRecorder) AllDependencyVersions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AllDependencyVersions", reflect.TypeOf((*MockManifest)(nil).AllDependencyVersions), arg0)
}

// InstallDependency mocks base method
func (_m *MockManifest) InstallDependency(dep libbuildpack.Dependency, outputDir string) error {
	ret := _m.ctrl.Call(_m, "InstallDependency", dep, outputDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (_mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallDependency", reflect.TypeOf((*MockManifest)(nil).InstallDependency), arg0, arg1)
}

// RootDir mocks base method
func (_m *MockManifest) RootDir() string {
	ret := _m.ctrl.Call(_m, "RootDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// RootDir indicates an expected call of RootDir
func (_mr *MockManifestMockRecorder) RootDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RootDir", reflect.TypeOf((*MockManifest)(nil).RootDir))
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStager) EXPECT() *MockStagerMockRecorder {
	return _m.recorder
}

// AddBinDependencyLink mocks base method
func (_m *MockStager) AddBinDependencyLink(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AddBinDependencyLink", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinDependencyLink indicates an expected call of AddBinDependencyLink
func (_mr *MockStagerMockRecorder) AddBinDependencyLink(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddBinDependencyLink", reflect.TypeOf((*MockStager)(nil).AddBinDependencyLink), arg0, arg1)
}

// DepDir mocks base method
func (_m *MockStager) DepDir() string {
	ret := _m.ctrl.Call(_m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (_mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}

// BuildDir mocks base method
func (_m *MockStager) BuildDir() string {
	ret := _m.ctrl.Call(_m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (_mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}
