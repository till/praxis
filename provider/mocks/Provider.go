package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"
import provider "github.com/convox/praxis/provider"

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// AppCreate provides a mock function with given fields: name, opts
func (_m *Provider) AppCreate(name string, opts provider.AppCreateOptions) (*provider.App, error) {
	ret := _m.Called(name, opts)

	var r0 *provider.App
	if rf, ok := ret.Get(0).(func(string, provider.AppCreateOptions) *provider.App); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, provider.AppCreateOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppDelete provides a mock function with given fields: name
func (_m *Provider) AppDelete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobStore provides a mock function with given fields: app, key, r, opts
func (_m *Provider) BlobStore(app string, key string, r io.Reader, opts provider.BlobStoreOptions) (string, error) {
	ret := _m.Called(app, key, r, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, provider.BlobStoreOptions) string); ok {
		r0 = rf(app, key, r, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader, provider.BlobStoreOptions) error); ok {
		r1 = rf(app, key, r, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildCreate provides a mock function with given fields: app, url, opts
func (_m *Provider) BuildCreate(app string, url string, opts provider.BuildCreateOptions) (*provider.Build, error) {
	ret := _m.Called(app, url, opts)

	var r0 *provider.Build
	if rf, ok := ret.Get(0).(func(string, string, provider.BuildCreateOptions) *provider.Build); ok {
		r0 = rf(app, url, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, provider.BuildCreateOptions) error); ok {
		r1 = rf(app, url, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildLoad provides a mock function with given fields: app, id
func (_m *Provider) BuildLoad(app string, id string) (*provider.Build, error) {
	ret := _m.Called(app, id)

	var r0 *provider.Build
	if rf, ok := ret.Get(0).(func(string, string) *provider.Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildSave provides a mock function with given fields: build
func (_m *Provider) BuildSave(build *provider.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*provider.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessStart provides a mock function with given fields: app, service, opts
func (_m *Provider) ProcessStart(app string, service string, opts provider.ProcessRunOptions) (*provider.Process, error) {
	ret := _m.Called(app, service, opts)

	var r0 *provider.Process
	if rf, ok := ret.Get(0).(func(string, string, provider.ProcessRunOptions) *provider.Process); ok {
		r0 = rf(app, service, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Process)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, provider.ProcessRunOptions) error); ok {
		r1 = rf(app, service, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessWait provides a mock function with given fields: app, pid
func (_m *Provider) ProcessWait(app string, pid string) (int, error) {
	ret := _m.Called(app, pid)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableLoad provides a mock function with given fields: app, table, id
func (_m *Provider) TableLoad(app string, table string, id string) (map[string]string, error) {
	ret := _m.Called(app, table, id)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string, string) map[string]string); ok {
		r0 = rf(app, table, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(app, table, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableRemove provides a mock function with given fields: app, table, id
func (_m *Provider) TableRemove(app string, table string, id string) error {
	ret := _m.Called(app, table, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(app, table, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableSave provides a mock function with given fields: app, table, id, attrs
func (_m *Provider) TableSave(app string, table string, id string, attrs map[string]string) error {
	ret := _m.Called(app, table, id, attrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]string) error); ok {
		r0 = rf(app, table, id, attrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ provider.Provider = (*Provider)(nil)
