// Code generated by counterfeiter. DO NOT EDIT.
package storefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/internal/models"
	"github.com/weaveworks/weave-gitops-enterprise/pkg/query/store"
)

type FakeStoreWriter struct {
	DeleteObjectsStub        func(context.Context, []models.Object) error
	deleteObjectsMutex       sync.RWMutex
	deleteObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Object
	}
	deleteObjectsReturns struct {
		result1 error
	}
	deleteObjectsReturnsOnCall map[int]struct {
		result1 error
	}
	StoreObjectsStub        func(context.Context, []models.Object) error
	storeObjectsMutex       sync.RWMutex
	storeObjectsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Object
	}
	storeObjectsReturns struct {
		result1 error
	}
	storeObjectsReturnsOnCall map[int]struct {
		result1 error
	}
	StoreRoleBindingsStub        func(context.Context, []models.RoleBinding) error
	storeRoleBindingsMutex       sync.RWMutex
	storeRoleBindingsArgsForCall []struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}
	storeRoleBindingsReturns struct {
		result1 error
	}
	storeRoleBindingsReturnsOnCall map[int]struct {
		result1 error
	}
	StoreRolesStub        func(context.Context, []models.Role) error
	storeRolesMutex       sync.RWMutex
	storeRolesArgsForCall []struct {
		arg1 context.Context
		arg2 []models.Role
	}
	storeRolesReturns struct {
		result1 error
	}
	storeRolesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStoreWriter) DeleteObjects(arg1 context.Context, arg2 []models.Object) error {
	var arg2Copy []models.Object
	if arg2 != nil {
		arg2Copy = make([]models.Object, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteObjectsMutex.Lock()
	ret, specificReturn := fake.deleteObjectsReturnsOnCall[len(fake.deleteObjectsArgsForCall)]
	fake.deleteObjectsArgsForCall = append(fake.deleteObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Object
	}{arg1, arg2Copy})
	stub := fake.DeleteObjectsStub
	fakeReturns := fake.deleteObjectsReturns
	fake.recordInvocation("DeleteObjects", []interface{}{arg1, arg2Copy})
	fake.deleteObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreWriter) DeleteObjectsCallCount() int {
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	return len(fake.deleteObjectsArgsForCall)
}

func (fake *FakeStoreWriter) DeleteObjectsCalls(stub func(context.Context, []models.Object) error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = stub
}

func (fake *FakeStoreWriter) DeleteObjectsArgsForCall(i int) (context.Context, []models.Object) {
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	argsForCall := fake.deleteObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStoreWriter) DeleteObjectsReturns(result1 error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = nil
	fake.deleteObjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) DeleteObjectsReturnsOnCall(i int, result1 error) {
	fake.deleteObjectsMutex.Lock()
	defer fake.deleteObjectsMutex.Unlock()
	fake.DeleteObjectsStub = nil
	if fake.deleteObjectsReturnsOnCall == nil {
		fake.deleteObjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteObjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreObjects(arg1 context.Context, arg2 []models.Object) error {
	var arg2Copy []models.Object
	if arg2 != nil {
		arg2Copy = make([]models.Object, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeObjectsMutex.Lock()
	ret, specificReturn := fake.storeObjectsReturnsOnCall[len(fake.storeObjectsArgsForCall)]
	fake.storeObjectsArgsForCall = append(fake.storeObjectsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Object
	}{arg1, arg2Copy})
	stub := fake.StoreObjectsStub
	fakeReturns := fake.storeObjectsReturns
	fake.recordInvocation("StoreObjects", []interface{}{arg1, arg2Copy})
	fake.storeObjectsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreWriter) StoreObjectsCallCount() int {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	return len(fake.storeObjectsArgsForCall)
}

func (fake *FakeStoreWriter) StoreObjectsCalls(stub func(context.Context, []models.Object) error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = stub
}

func (fake *FakeStoreWriter) StoreObjectsArgsForCall(i int) (context.Context, []models.Object) {
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	argsForCall := fake.storeObjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStoreWriter) StoreObjectsReturns(result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	fake.storeObjectsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreObjectsReturnsOnCall(i int, result1 error) {
	fake.storeObjectsMutex.Lock()
	defer fake.storeObjectsMutex.Unlock()
	fake.StoreObjectsStub = nil
	if fake.storeObjectsReturnsOnCall == nil {
		fake.storeObjectsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeObjectsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreRoleBindings(arg1 context.Context, arg2 []models.RoleBinding) error {
	var arg2Copy []models.RoleBinding
	if arg2 != nil {
		arg2Copy = make([]models.RoleBinding, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeRoleBindingsMutex.Lock()
	ret, specificReturn := fake.storeRoleBindingsReturnsOnCall[len(fake.storeRoleBindingsArgsForCall)]
	fake.storeRoleBindingsArgsForCall = append(fake.storeRoleBindingsArgsForCall, struct {
		arg1 context.Context
		arg2 []models.RoleBinding
	}{arg1, arg2Copy})
	stub := fake.StoreRoleBindingsStub
	fakeReturns := fake.storeRoleBindingsReturns
	fake.recordInvocation("StoreRoleBindings", []interface{}{arg1, arg2Copy})
	fake.storeRoleBindingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreWriter) StoreRoleBindingsCallCount() int {
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	return len(fake.storeRoleBindingsArgsForCall)
}

func (fake *FakeStoreWriter) StoreRoleBindingsCalls(stub func(context.Context, []models.RoleBinding) error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = stub
}

func (fake *FakeStoreWriter) StoreRoleBindingsArgsForCall(i int) (context.Context, []models.RoleBinding) {
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	argsForCall := fake.storeRoleBindingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStoreWriter) StoreRoleBindingsReturns(result1 error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = nil
	fake.storeRoleBindingsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreRoleBindingsReturnsOnCall(i int, result1 error) {
	fake.storeRoleBindingsMutex.Lock()
	defer fake.storeRoleBindingsMutex.Unlock()
	fake.StoreRoleBindingsStub = nil
	if fake.storeRoleBindingsReturnsOnCall == nil {
		fake.storeRoleBindingsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeRoleBindingsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreRoles(arg1 context.Context, arg2 []models.Role) error {
	var arg2Copy []models.Role
	if arg2 != nil {
		arg2Copy = make([]models.Role, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.storeRolesMutex.Lock()
	ret, specificReturn := fake.storeRolesReturnsOnCall[len(fake.storeRolesArgsForCall)]
	fake.storeRolesArgsForCall = append(fake.storeRolesArgsForCall, struct {
		arg1 context.Context
		arg2 []models.Role
	}{arg1, arg2Copy})
	stub := fake.StoreRolesStub
	fakeReturns := fake.storeRolesReturns
	fake.recordInvocation("StoreRoles", []interface{}{arg1, arg2Copy})
	fake.storeRolesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreWriter) StoreRolesCallCount() int {
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	return len(fake.storeRolesArgsForCall)
}

func (fake *FakeStoreWriter) StoreRolesCalls(stub func(context.Context, []models.Role) error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = stub
}

func (fake *FakeStoreWriter) StoreRolesArgsForCall(i int) (context.Context, []models.Role) {
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	argsForCall := fake.storeRolesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStoreWriter) StoreRolesReturns(result1 error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = nil
	fake.storeRolesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) StoreRolesReturnsOnCall(i int, result1 error) {
	fake.storeRolesMutex.Lock()
	defer fake.storeRolesMutex.Unlock()
	fake.StoreRolesStub = nil
	if fake.storeRolesReturnsOnCall == nil {
		fake.storeRolesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeRolesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteObjectsMutex.RLock()
	defer fake.deleteObjectsMutex.RUnlock()
	fake.storeObjectsMutex.RLock()
	defer fake.storeObjectsMutex.RUnlock()
	fake.storeRoleBindingsMutex.RLock()
	defer fake.storeRoleBindingsMutex.RUnlock()
	fake.storeRolesMutex.RLock()
	defer fake.storeRolesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStoreWriter) recordInvocation(key string, args []interface{}) {
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

var _ store.StoreWriter = new(FakeStoreWriter)
