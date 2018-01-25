// This file was generated by counterfeiter
package scriptfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agent/script"
)

type FakeScript struct {
	TagStub        func() string
	tagMutex       sync.RWMutex
	tagArgsForCall []struct{}
	tagReturns     struct {
		result1 string
	}
	tagReturnsOnCall map[int]struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	ExistsStub        func() bool
	existsMutex       sync.RWMutex
	existsArgsForCall []struct{}
	existsReturns     struct {
		result1 bool
	}
	existsReturnsOnCall map[int]struct {
		result1 bool
	}
	RunStub        func() error
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns     struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScript) Tag() string {
	fake.tagMutex.Lock()
	ret, specificReturn := fake.tagReturnsOnCall[len(fake.tagArgsForCall)]
	fake.tagArgsForCall = append(fake.tagArgsForCall, struct{}{})
	fake.recordInvocation("Tag", []interface{}{})
	fake.tagMutex.Unlock()
	if fake.TagStub != nil {
		return fake.TagStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tagReturns.result1
}

func (fake *FakeScript) TagCallCount() int {
	fake.tagMutex.RLock()
	defer fake.tagMutex.RUnlock()
	return len(fake.tagArgsForCall)
}

func (fake *FakeScript) TagReturns(result1 string) {
	fake.TagStub = nil
	fake.tagReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) TagReturnsOnCall(i int, result1 string) {
	fake.TagStub = nil
	if fake.tagReturnsOnCall == nil {
		fake.tagReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.tagReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pathReturns.result1
}

func (fake *FakeScript) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeScript) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) PathReturnsOnCall(i int, result1 string) {
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeScript) Exists() bool {
	fake.existsMutex.Lock()
	ret, specificReturn := fake.existsReturnsOnCall[len(fake.existsArgsForCall)]
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct{}{})
	fake.recordInvocation("Exists", []interface{}{})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.existsReturns.result1
}

func (fake *FakeScript) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeScript) ExistsReturns(result1 bool) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeScript) ExistsReturnsOnCall(i int, result1 bool) {
	fake.ExistsStub = nil
	if fake.existsReturnsOnCall == nil {
		fake.existsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.existsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeScript) Run() error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeScript) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeScript) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScript) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScript) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tagMutex.RLock()
	defer fake.tagMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScript) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ script.Script = new(FakeScript)
