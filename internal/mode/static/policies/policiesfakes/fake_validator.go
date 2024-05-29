// Code generated by counterfeiter. DO NOT EDIT.
package policiesfakes

import (
	"sync"

	"github.com/nginxinc/nginx-gateway-fabric/internal/framework/conditions"
	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/policies"
)

type FakeValidator struct {
	ConflictsStub        func(policies.Policy, policies.Policy) bool
	conflictsMutex       sync.RWMutex
	conflictsArgsForCall []struct {
		arg1 policies.Policy
		arg2 policies.Policy
	}
	conflictsReturns struct {
		result1 bool
	}
	conflictsReturnsOnCall map[int]struct {
		result1 bool
	}
	ValidateStub        func(policies.Policy, *policies.GlobalSettings) []conditions.Condition
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 policies.Policy
		arg2 *policies.GlobalSettings
	}
	validateReturns struct {
		result1 []conditions.Condition
	}
	validateReturnsOnCall map[int]struct {
		result1 []conditions.Condition
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeValidator) Conflicts(arg1 policies.Policy, arg2 policies.Policy) bool {
	fake.conflictsMutex.Lock()
	ret, specificReturn := fake.conflictsReturnsOnCall[len(fake.conflictsArgsForCall)]
	fake.conflictsArgsForCall = append(fake.conflictsArgsForCall, struct {
		arg1 policies.Policy
		arg2 policies.Policy
	}{arg1, arg2})
	stub := fake.ConflictsStub
	fakeReturns := fake.conflictsReturns
	fake.recordInvocation("Conflicts", []interface{}{arg1, arg2})
	fake.conflictsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeValidator) ConflictsCallCount() int {
	fake.conflictsMutex.RLock()
	defer fake.conflictsMutex.RUnlock()
	return len(fake.conflictsArgsForCall)
}

func (fake *FakeValidator) ConflictsCalls(stub func(policies.Policy, policies.Policy) bool) {
	fake.conflictsMutex.Lock()
	defer fake.conflictsMutex.Unlock()
	fake.ConflictsStub = stub
}

func (fake *FakeValidator) ConflictsArgsForCall(i int) (policies.Policy, policies.Policy) {
	fake.conflictsMutex.RLock()
	defer fake.conflictsMutex.RUnlock()
	argsForCall := fake.conflictsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeValidator) ConflictsReturns(result1 bool) {
	fake.conflictsMutex.Lock()
	defer fake.conflictsMutex.Unlock()
	fake.ConflictsStub = nil
	fake.conflictsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeValidator) ConflictsReturnsOnCall(i int, result1 bool) {
	fake.conflictsMutex.Lock()
	defer fake.conflictsMutex.Unlock()
	fake.ConflictsStub = nil
	if fake.conflictsReturnsOnCall == nil {
		fake.conflictsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.conflictsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeValidator) Validate(arg1 policies.Policy, arg2 *policies.GlobalSettings) []conditions.Condition {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 policies.Policy
		arg2 *policies.GlobalSettings
	}{arg1, arg2})
	stub := fake.ValidateStub
	fakeReturns := fake.validateReturns
	fake.recordInvocation("Validate", []interface{}{arg1, arg2})
	fake.validateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeValidator) ValidateCalls(stub func(policies.Policy, *policies.GlobalSettings) []conditions.Condition) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *FakeValidator) ValidateArgsForCall(i int) (policies.Policy, *policies.GlobalSettings) {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeValidator) ValidateReturns(result1 []conditions.Condition) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []conditions.Condition
	}{result1}
}

func (fake *FakeValidator) ValidateReturnsOnCall(i int, result1 []conditions.Condition) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 []conditions.Condition
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 []conditions.Condition
	}{result1}
}

func (fake *FakeValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.conflictsMutex.RLock()
	defer fake.conflictsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeValidator) recordInvocation(key string, args []interface{}) {
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

var _ policies.Validator = new(FakeValidator)
