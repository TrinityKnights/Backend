// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/gomail/gomail.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/gomail/gomail.go -destination=test/mock/./pkg/gomail/gomail_mock.go
//

// Package mock_gomail is a generated GoMock package.
package mock_gomail

import (
	reflect "reflect"

	gomail "github.com/TrinityKnights/Backend/pkg/gomail"
	gomock "go.uber.org/mock/gomock"
)

// MockGomail is a mock of Gomail interface.
type MockGomail struct {
	ctrl     *gomock.Controller
	recorder *MockGomailMockRecorder
	isgomock struct{}
}

// MockGomailMockRecorder is the mock recorder for MockGomail.
type MockGomailMockRecorder struct {
	mock *MockGomail
}

// NewMockGomail creates a new mock instance.
func NewMockGomail(ctrl *gomock.Controller) *MockGomail {
	mock := &MockGomail{ctrl: ctrl}
	mock.recorder = &MockGomailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGomail) EXPECT() *MockGomailMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockGomail) SendEmail(request *gomail.SendEmail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockGomailMockRecorder) SendEmail(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockGomail)(nil).SendEmail), request)
}