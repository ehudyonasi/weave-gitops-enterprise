// Code generated by counterfeiter. DO NOT EDIT.
package accesscheckerfakes

import (
	"sync"

	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/accesschecker"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/internal/models"
	"github.com/weaveworks/weave-gitops/pkg/server/auth"
)

type FakeChecker struct {
	HasAccessStub        func(*auth.UserPrincipal, models.Object, []models.AccessRule) (bool, error)
	hasAccessMutex       sync.RWMutex
	hasAccessArgsForCall []struct {
		arg1 *auth.UserPrincipal
		arg2 models.Object
		arg3 []models.AccessRule
	}
	hasAccessReturns struct {
		result1 bool
		result2 error
	}
	hasAccessReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RelevantRulesForUserStub        func(*auth.UserPrincipal, []models.AccessRule) []models.AccessRule
	relevantRulesForUserMutex       sync.RWMutex
	relevantRulesForUserArgsForCall []struct {
		arg1 *auth.UserPrincipal
		arg2 []models.AccessRule
	}
	relevantRulesForUserReturns struct {
		result1 []models.AccessRule
	}
	relevantRulesForUserReturnsOnCall map[int]struct {
		result1 []models.AccessRule
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChecker) HasAccess(arg1 *auth.UserPrincipal, arg2 models.Object, arg3 []models.AccessRule) (bool, error) {
	var arg3Copy []models.AccessRule
	if arg3 != nil {
		arg3Copy = make([]models.AccessRule, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.hasAccessMutex.Lock()
	ret, specificReturn := fake.hasAccessReturnsOnCall[len(fake.hasAccessArgsForCall)]
	fake.hasAccessArgsForCall = append(fake.hasAccessArgsForCall, struct {
		arg1 *auth.UserPrincipal
		arg2 models.Object
		arg3 []models.AccessRule
	}{arg1, arg2, arg3Copy})
	stub := fake.HasAccessStub
	fakeReturns := fake.hasAccessReturns
	fake.recordInvocation("HasAccess", []interface{}{arg1, arg2, arg3Copy})
	fake.hasAccessMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeChecker) HasAccessCallCount() int {
	fake.hasAccessMutex.RLock()
	defer fake.hasAccessMutex.RUnlock()
	return len(fake.hasAccessArgsForCall)
}

func (fake *FakeChecker) HasAccessCalls(stub func(*auth.UserPrincipal, models.Object, []models.AccessRule) (bool, error)) {
	fake.hasAccessMutex.Lock()
	defer fake.hasAccessMutex.Unlock()
	fake.HasAccessStub = stub
}

func (fake *FakeChecker) HasAccessArgsForCall(i int) (*auth.UserPrincipal, models.Object, []models.AccessRule) {
	fake.hasAccessMutex.RLock()
	defer fake.hasAccessMutex.RUnlock()
	argsForCall := fake.hasAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeChecker) HasAccessReturns(result1 bool, result2 error) {
	fake.hasAccessMutex.Lock()
	defer fake.hasAccessMutex.Unlock()
	fake.HasAccessStub = nil
	fake.hasAccessReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) HasAccessReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasAccessMutex.Lock()
	defer fake.hasAccessMutex.Unlock()
	fake.HasAccessStub = nil
	if fake.hasAccessReturnsOnCall == nil {
		fake.hasAccessReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasAccessReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) RelevantRulesForUser(arg1 *auth.UserPrincipal, arg2 []models.AccessRule) []models.AccessRule {
	var arg2Copy []models.AccessRule
	if arg2 != nil {
		arg2Copy = make([]models.AccessRule, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.relevantRulesForUserMutex.Lock()
	ret, specificReturn := fake.relevantRulesForUserReturnsOnCall[len(fake.relevantRulesForUserArgsForCall)]
	fake.relevantRulesForUserArgsForCall = append(fake.relevantRulesForUserArgsForCall, struct {
		arg1 *auth.UserPrincipal
		arg2 []models.AccessRule
	}{arg1, arg2Copy})
	stub := fake.RelevantRulesForUserStub
	fakeReturns := fake.relevantRulesForUserReturns
	fake.recordInvocation("RelevantRulesForUser", []interface{}{arg1, arg2Copy})
	fake.relevantRulesForUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeChecker) RelevantRulesForUserCallCount() int {
	fake.relevantRulesForUserMutex.RLock()
	defer fake.relevantRulesForUserMutex.RUnlock()
	return len(fake.relevantRulesForUserArgsForCall)
}

func (fake *FakeChecker) RelevantRulesForUserCalls(stub func(*auth.UserPrincipal, []models.AccessRule) []models.AccessRule) {
	fake.relevantRulesForUserMutex.Lock()
	defer fake.relevantRulesForUserMutex.Unlock()
	fake.RelevantRulesForUserStub = stub
}

func (fake *FakeChecker) RelevantRulesForUserArgsForCall(i int) (*auth.UserPrincipal, []models.AccessRule) {
	fake.relevantRulesForUserMutex.RLock()
	defer fake.relevantRulesForUserMutex.RUnlock()
	argsForCall := fake.relevantRulesForUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeChecker) RelevantRulesForUserReturns(result1 []models.AccessRule) {
	fake.relevantRulesForUserMutex.Lock()
	defer fake.relevantRulesForUserMutex.Unlock()
	fake.RelevantRulesForUserStub = nil
	fake.relevantRulesForUserReturns = struct {
		result1 []models.AccessRule
	}{result1}
}

func (fake *FakeChecker) RelevantRulesForUserReturnsOnCall(i int, result1 []models.AccessRule) {
	fake.relevantRulesForUserMutex.Lock()
	defer fake.relevantRulesForUserMutex.Unlock()
	fake.RelevantRulesForUserStub = nil
	if fake.relevantRulesForUserReturnsOnCall == nil {
		fake.relevantRulesForUserReturnsOnCall = make(map[int]struct {
			result1 []models.AccessRule
		})
	}
	fake.relevantRulesForUserReturnsOnCall[i] = struct {
		result1 []models.AccessRule
	}{result1}
}

func (fake *FakeChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hasAccessMutex.RLock()
	defer fake.hasAccessMutex.RUnlock()
	fake.relevantRulesForUserMutex.RLock()
	defer fake.relevantRulesForUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChecker) recordInvocation(key string, args []interface{}) {
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

var _ accesschecker.Checker = new(FakeChecker)