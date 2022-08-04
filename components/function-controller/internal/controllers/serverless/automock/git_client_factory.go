// Code generated by mockery v2.14.0. DO NOT EDIT.

package automock

import (
	git "github.com/kyma-project/kyma/components/function-controller/internal/git"
	mock "github.com/stretchr/testify/mock"

	zap "go.uber.org/zap"
)

// GitClientFactory is an autogenerated mock type for the GitClientFactory type
type GitClientFactory struct {
	mock.Mock
}

// GetGitClient provides a mock function with given fields: logger
func (_m *GitClientFactory) GetGitClient(logger *zap.SugaredLogger) git.GitClient {
	ret := _m.Called(logger)

	var r0 git.GitClient
	if rf, ok := ret.Get(0).(func(*zap.SugaredLogger) git.GitClient); ok {
		r0 = rf(logger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(git.GitClient)
		}
	}

	return r0
}

type mockConstructorTestingTNewGitClientFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewGitClientFactory creates a new instance of GitClientFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGitClientFactory(t mockConstructorTestingTNewGitClientFactory) *GitClientFactory {
	mock := &GitClientFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
