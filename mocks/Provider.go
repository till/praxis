// Code generated by mockery v1.0.0
package mocks

import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"
import types "github.com/convox/praxis/types"

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// AppCreate provides a mock function with given fields: name
func (_m *Provider) AppCreate(name string) (*types.App, error) {
	ret := _m.Called(name)

	var r0 *types.App
	if rf, ok := ret.Get(0).(func(string) *types.App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
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

// AppGet provides a mock function with given fields: name
func (_m *Provider) AppGet(name string) (*types.App, error) {
	ret := _m.Called(name)

	var r0 *types.App
	if rf, ok := ret.Get(0).(func(string) *types.App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppList provides a mock function with given fields:
func (_m *Provider) AppList() (types.Apps, error) {
	ret := _m.Called()

	var r0 types.Apps
	if rf, ok := ret.Get(0).(func() types.Apps); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Apps)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppLogs provides a mock function with given fields: app, opts
func (_m *Provider) AppLogs(app string, opts types.LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(app, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, types.LogsOptions) io.ReadCloser); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.LogsOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppRegistry provides a mock function with given fields: app
func (_m *Provider) AppRegistry(app string) (*types.Registry, error) {
	ret := _m.Called(app)

	var r0 *types.Registry
	if rf, ok := ret.Get(0).(func(string) *types.Registry); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildCreate provides a mock function with given fields: app, url, opts
func (_m *Provider) BuildCreate(app string, url string, opts types.BuildCreateOptions) (*types.Build, error) {
	ret := _m.Called(app, url, opts)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string, types.BuildCreateOptions) *types.Build); ok {
		r0 = rf(app, url, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.BuildCreateOptions) error); ok {
		r1 = rf(app, url, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildGet provides a mock function with given fields: app, id
func (_m *Provider) BuildGet(app string, id string) (*types.Build, error) {
	ret := _m.Called(app, id)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string) *types.Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
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

// BuildList provides a mock function with given fields: app
func (_m *Provider) BuildList(app string) (types.Builds, error) {
	ret := _m.Called(app)

	var r0 types.Builds
	if rf, ok := ret.Get(0).(func(string) types.Builds); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Builds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildLogs provides a mock function with given fields: app, id
func (_m *Provider) BuildLogs(app string, id string) (io.ReadCloser, error) {
	ret := _m.Called(app, id)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
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

// BuildUpdate provides a mock function with given fields: app, id, opts
func (_m *Provider) BuildUpdate(app string, id string, opts types.BuildUpdateOptions) (*types.Build, error) {
	ret := _m.Called(app, id, opts)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string, types.BuildUpdateOptions) *types.Build); ok {
		r0 = rf(app, id, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.BuildUpdateOptions) error); ok {
		r1 = rf(app, id, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheFetch provides a mock function with given fields: app, cache, key
func (_m *Provider) CacheFetch(app string, cache string, key string) (map[string]string, error) {
	ret := _m.Called(app, cache, key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string, string) map[string]string); ok {
		r0 = rf(app, cache, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(app, cache, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheStore provides a mock function with given fields: app, cache, key, attrs, opts
func (_m *Provider) CacheStore(app string, cache string, key string, attrs map[string]string, opts types.CacheStoreOptions) error {
	ret := _m.Called(app, cache, key, attrs, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]string, types.CacheStoreOptions) error); ok {
		r0 = rf(app, cache, key, attrs, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilesDelete provides a mock function with given fields: app, pid, files
func (_m *Provider) FilesDelete(app string, pid string, files []string) error {
	ret := _m.Called(app, pid, files)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(app, pid, files)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilesUpload provides a mock function with given fields: app, pid, r
func (_m *Provider) FilesUpload(app string, pid string, r io.Reader) error {
	ret := _m.Called(app, pid, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(app, pid, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageCreate provides a mock function with given fields: name, url, opts
func (_m *Provider) ImageCreate(name string, url string, opts types.ImageCreateOptions) (*types.Image, error) {
	ret := _m.Called(name, url, opts)

	var r0 *types.Image
	if rf, ok := ret.Get(0).(func(string, string, types.ImageCreateOptions) *types.Image); ok {
		r0 = rf(name, url, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.ImageCreateOptions) error); ok {
		r1 = rf(name, url, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageList provides a mock function with given fields:
func (_m *Provider) ImageList() (types.Images, error) {
	ret := _m.Called()

	var r0 types.Images
	if rf, ok := ret.Get(0).(func() types.Images); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Images)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyDecrypt provides a mock function with given fields: app, key, data
func (_m *Provider) KeyDecrypt(app string, key string, data []byte) ([]byte, error) {
	ret := _m.Called(app, key, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, []byte) []byte); ok {
		r0 = rf(app, key, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []byte) error); ok {
		r1 = rf(app, key, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyEncrypt provides a mock function with given fields: app, key, data
func (_m *Provider) KeyEncrypt(app string, key string, data []byte) ([]byte, error) {
	ret := _m.Called(app, key, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, []byte) []byte); ok {
		r0 = rf(app, key, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []byte) error); ok {
		r1 = rf(app, key, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectExists provides a mock function with given fields: app, key
func (_m *Provider) ObjectExists(app string, key string) (bool, error) {
	ret := _m.Called(app, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(app, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectFetch provides a mock function with given fields: app, key
func (_m *Provider) ObjectFetch(app string, key string) (io.ReadCloser, error) {
	ret := _m.Called(app, key)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectStore provides a mock function with given fields: app, key, r, opts
func (_m *Provider) ObjectStore(app string, key string, r io.Reader, opts types.ObjectStoreOptions) (*types.Object, error) {
	ret := _m.Called(app, key, r, opts)

	var r0 *types.Object
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, types.ObjectStoreOptions) *types.Object); ok {
		r0 = rf(app, key, r, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader, types.ObjectStoreOptions) error); ok {
		r1 = rf(app, key, r, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessExec provides a mock function with given fields: app, pid, command, opts
func (_m *Provider) ProcessExec(app string, pid string, command string, opts types.ProcessExecOptions) (int, error) {
	ret := _m.Called(app, pid, command, opts)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, types.ProcessExecOptions) int); ok {
		r0 = rf(app, pid, command, opts)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, types.ProcessExecOptions) error); ok {
		r1 = rf(app, pid, command, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessGet provides a mock function with given fields: app, pid
func (_m *Provider) ProcessGet(app string, pid string) (*types.Process, error) {
	ret := _m.Called(app, pid)

	var r0 *types.Process
	if rf, ok := ret.Get(0).(func(string, string) *types.Process); ok {
		r0 = rf(app, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Process)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessList provides a mock function with given fields: app, opts
func (_m *Provider) ProcessList(app string, opts types.ProcessListOptions) (types.Processes, error) {
	ret := _m.Called(app, opts)

	var r0 types.Processes
	if rf, ok := ret.Get(0).(func(string, types.ProcessListOptions) types.Processes); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessLogs provides a mock function with given fields: app, pid, opts
func (_m *Provider) ProcessLogs(app string, pid string, opts types.LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(app, pid, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, types.LogsOptions) io.ReadCloser); ok {
		r0 = rf(app, pid, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.LogsOptions) error); ok {
		r1 = rf(app, pid, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessRun provides a mock function with given fields: app, opts
func (_m *Provider) ProcessRun(app string, opts types.ProcessRunOptions) (int, error) {
	ret := _m.Called(app, opts)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, types.ProcessRunOptions) int); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessRunOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStart provides a mock function with given fields: app, opts
func (_m *Provider) ProcessStart(app string, opts types.ProcessRunOptions) (string, error) {
	ret := _m.Called(app, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, types.ProcessRunOptions) string); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessRunOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStop provides a mock function with given fields: app, pid
func (_m *Provider) ProcessStop(app string, pid string) error {
	ret := _m.Called(app, pid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Proxy provides a mock function with given fields: app, pid, port, in
func (_m *Provider) Proxy(app string, pid string, port int, in io.Reader) (io.ReadCloser, error) {
	ret := _m.Called(app, pid, port, in)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, int, io.Reader) io.ReadCloser); ok {
		r0 = rf(app, pid, port, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, io.Reader) error); ok {
		r1 = rf(app, pid, port, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueFetch provides a mock function with given fields: app, queue, opts
func (_m *Provider) QueueFetch(app string, queue string, opts types.QueueFetchOptions) (map[string]string, error) {
	ret := _m.Called(app, queue, opts)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string, types.QueueFetchOptions) map[string]string); ok {
		r0 = rf(app, queue, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.QueueFetchOptions) error); ok {
		r1 = rf(app, queue, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueStore provides a mock function with given fields: app, queue, attrs
func (_m *Provider) QueueStore(app string, queue string, attrs map[string]string) error {
	ret := _m.Called(app, queue, attrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) error); ok {
		r0 = rf(app, queue, attrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryAdd provides a mock function with given fields: server, username, password
func (_m *Provider) RegistryAdd(server string, username string, password string) (*types.Registry, error) {
	ret := _m.Called(server, username, password)

	var r0 *types.Registry
	if rf, ok := ret.Get(0).(func(string, string, string) *types.Registry); ok {
		r0 = rf(server, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(server, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryList provides a mock function with given fields:
func (_m *Provider) RegistryList() (types.Registries, error) {
	ret := _m.Called()

	var r0 types.Registries
	if rf, ok := ret.Get(0).(func() types.Registries); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Registries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryRemove provides a mock function with given fields: server
func (_m *Provider) RegistryRemove(server string) error {
	ret := _m.Called(server)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseCreate provides a mock function with given fields: app, opts
func (_m *Provider) ReleaseCreate(app string, opts types.ReleaseCreateOptions) (*types.Release, error) {
	ret := _m.Called(app, opts)

	var r0 *types.Release
	if rf, ok := ret.Get(0).(func(string, types.ReleaseCreateOptions) *types.Release); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ReleaseCreateOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseGet provides a mock function with given fields: app, id
func (_m *Provider) ReleaseGet(app string, id string) (*types.Release, error) {
	ret := _m.Called(app, id)

	var r0 *types.Release
	if rf, ok := ret.Get(0).(func(string, string) *types.Release); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Release)
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

// ReleaseList provides a mock function with given fields: app, opts
func (_m *Provider) ReleaseList(app string, opts types.ReleaseListOptions) (types.Releases, error) {
	ret := _m.Called(app, opts)

	var r0 types.Releases
	if rf, ok := ret.Get(0).(func(string, types.ReleaseListOptions) types.Releases); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ReleaseListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseLogs provides a mock function with given fields: app, id, opts
func (_m *Provider) ReleaseLogs(app string, id string, opts types.LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(app, id, opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, types.LogsOptions) io.ReadCloser); ok {
		r0 = rf(app, id, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.LogsOptions) error); ok {
		r1 = rf(app, id, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleasePromote provides a mock function with given fields: app, id
func (_m *Provider) ReleasePromote(app string, id string) error {
	ret := _m.Called(app, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceList provides a mock function with given fields: app
func (_m *Provider) ResourceList(app string) (types.Resources, error) {
	ret := _m.Called(app)

	var r0 types.Resources
	if rf, ok := ret.Get(0).(func(string) types.Resources); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceList provides a mock function with given fields: app
func (_m *Provider) ServiceList(app string) (types.Services, error) {
	ret := _m.Called(app)

	var r0 types.Services
	if rf, ok := ret.Get(0).(func(string) types.Services); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemGet provides a mock function with given fields:
func (_m *Provider) SystemGet() (*types.System, error) {
	ret := _m.Called()

	var r0 *types.System
	if rf, ok := ret.Get(0).(func() *types.System); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.System)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemInstall provides a mock function with given fields: name, opts
func (_m *Provider) SystemInstall(name string, opts types.SystemInstallOptions) (string, error) {
	ret := _m.Called(name, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, types.SystemInstallOptions) string); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.SystemInstallOptions) error); ok {
		r1 = rf(name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemLogs provides a mock function with given fields: opts
func (_m *Provider) SystemLogs(opts types.LogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(opts)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(types.LogsOptions) io.ReadCloser); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.LogsOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemOptions provides a mock function with given fields:
func (_m *Provider) SystemOptions() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemUninstall provides a mock function with given fields: name, opts
func (_m *Provider) SystemUninstall(name string, opts types.SystemInstallOptions) error {
	ret := _m.Called(name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.SystemInstallOptions) error); ok {
		r0 = rf(name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemUpdate provides a mock function with given fields: opts
func (_m *Provider) SystemUpdate(opts types.SystemUpdateOptions) error {
	ret := _m.Called(opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.SystemUpdateOptions) error); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableGet provides a mock function with given fields: app, table
func (_m *Provider) TableGet(app string, table string) (*types.Table, error) {
	ret := _m.Called(app, table)

	var r0 *types.Table
	if rf, ok := ret.Get(0).(func(string, string) *types.Table); ok {
		r0 = rf(app, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Table)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableList provides a mock function with given fields: app
func (_m *Provider) TableList(app string) (types.Tables, error) {
	ret := _m.Called(app)

	var r0 types.Tables
	if rf, ok := ret.Get(0).(func(string) types.Tables); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Tables)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableQuery provides a mock function with given fields: app, table, query
func (_m *Provider) TableQuery(app string, table string, query string) (types.TableRows, error) {
	ret := _m.Called(app, table, query)

	var r0 types.TableRows
	if rf, ok := ret.Get(0).(func(string, string, string) types.TableRows); ok {
		r0 = rf(app, table, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TableRows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(app, table, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableTruncate provides a mock function with given fields: app, table
func (_m *Provider) TableTruncate(app string, table string) error {
	ret := _m.Called(app, table)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, table)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Provider) WithContext(ctx context.Context) types.Provider {
	ret := _m.Called(ctx)

	var r0 types.Provider
	if rf, ok := ret.Get(0).(func(context.Context) types.Provider); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Provider)
		}
	}

	return r0
}

// Workers provides a mock function with given fields:
func (_m *Provider) Workers() {
	_m.Called()
}
