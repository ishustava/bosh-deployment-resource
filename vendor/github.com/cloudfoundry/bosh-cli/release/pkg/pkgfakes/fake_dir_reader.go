// Code generated by counterfeiter. DO NOT EDIT.
package pkgfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/release/pkg"
)

type FakeDirReader struct {
	ReadStub        func(string) (*pkg.Package, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 *pkg.Package
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 *pkg.Package
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirReader) Read(arg1 string) (*pkg.Package, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *FakeDirReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeDirReader) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].arg1
}

func (fake *FakeDirReader) ReadReturns(result1 *pkg.Package, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 *pkg.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeDirReader) ReadReturnsOnCall(i int, result1 *pkg.Package, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 *pkg.Package
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 *pkg.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeDirReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDirReader) recordInvocation(key string, args []interface{}) {
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

var _ pkg.DirReader = new(FakeDirReader)