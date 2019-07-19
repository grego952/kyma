// Code generated by mockery v1.0.0
package automock

import bundle "github.com/kyma-project/kyma/components/helm-broker/internal/bundle"

import mock "github.com/stretchr/testify/mock"

// BundleProvider is an autogenerated mock type for the BundleProvider type
type BundleProvider struct {
	mock.Mock
}

// GetIndex provides a mock function with given fields: _a0
func (_m *BundleProvider) GetIndex(_a0 string) (*bundle.IndexDTO, error) {
	ret := _m.Called(_a0)

	var r0 *bundle.IndexDTO
	if rf, ok := ret.Get(0).(func(string) *bundle.IndexDTO); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundle.IndexDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadCompleteBundle provides a mock function with given fields: _a0
func (_m *BundleProvider) LoadCompleteBundle(_a0 bundle.EntryDTO) (bundle.CompleteBundle, error) {
	ret := _m.Called(_a0)

	var r0 bundle.CompleteBundle
	if rf, ok := ret.Get(0).(func(bundle.EntryDTO) bundle.CompleteBundle); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bundle.CompleteBundle)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bundle.EntryDTO) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
