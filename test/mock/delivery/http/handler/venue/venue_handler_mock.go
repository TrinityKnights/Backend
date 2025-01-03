// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/delivery/http/handler/venue/venue_handler.go
//
// Generated by this command:
//
//	mockgen -source=./internal/delivery/http/handler/venue/venue_handler.go -destination=test/mock/delivery/http/handler/venue/venue_handler_mock.go
//

// Package mock_venue is a generated GoMock package.
package mock_venue

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockVenueHandler is a mock of VenueHandler interface.
type MockVenueHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVenueHandlerMockRecorder
	isgomock struct{}
}

// MockVenueHandlerMockRecorder is the mock recorder for MockVenueHandler.
type MockVenueHandlerMockRecorder struct {
	mock *MockVenueHandler
}

// NewMockVenueHandler creates a new mock instance.
func NewMockVenueHandler(ctrl *gomock.Controller) *MockVenueHandler {
	mock := &MockVenueHandler{ctrl: ctrl}
	mock.recorder = &MockVenueHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVenueHandler) EXPECT() *MockVenueHandlerMockRecorder {
	return m.recorder
}

// CreateVenue mocks base method.
func (m *MockVenueHandler) CreateVenue(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVenue", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVenue indicates an expected call of CreateVenue.
func (mr *MockVenueHandlerMockRecorder) CreateVenue(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVenue", reflect.TypeOf((*MockVenueHandler)(nil).CreateVenue), ctx)
}

// GetAllVenues mocks base method.
func (m *MockVenueHandler) GetAllVenues(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVenues", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllVenues indicates an expected call of GetAllVenues.
func (mr *MockVenueHandlerMockRecorder) GetAllVenues(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVenues", reflect.TypeOf((*MockVenueHandler)(nil).GetAllVenues), ctx)
}

// GetVenueByID mocks base method.
func (m *MockVenueHandler) GetVenueByID(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVenueByID", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetVenueByID indicates an expected call of GetVenueByID.
func (mr *MockVenueHandlerMockRecorder) GetVenueByID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVenueByID", reflect.TypeOf((*MockVenueHandler)(nil).GetVenueByID), ctx)
}

// SearchVenues mocks base method.
func (m *MockVenueHandler) SearchVenues(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchVenues", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchVenues indicates an expected call of SearchVenues.
func (mr *MockVenueHandlerMockRecorder) SearchVenues(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchVenues", reflect.TypeOf((*MockVenueHandler)(nil).SearchVenues), ctx)
}

// UpdateVenue mocks base method.
func (m *MockVenueHandler) UpdateVenue(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVenue", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVenue indicates an expected call of UpdateVenue.
func (mr *MockVenueHandlerMockRecorder) UpdateVenue(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVenue", reflect.TypeOf((*MockVenueHandler)(nil).UpdateVenue), ctx)
}
