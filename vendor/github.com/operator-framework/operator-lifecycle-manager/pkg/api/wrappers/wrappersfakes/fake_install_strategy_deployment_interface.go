// Code generated by counterfeiter. DO NOT EDIT.
package wrappersfakes

import (
	"sync"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/wrappers"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/ownerutil"
	v1 "k8s.io/api/apps/v1"
	v1b "k8s.io/api/core/v1"
	v1a "k8s.io/api/rbac/v1"
	labels "k8s.io/apimachinery/pkg/labels"
)

type FakeInstallStrategyDeploymentInterface struct {
	CreateDeploymentStub        func(*v1.Deployment) (*v1.Deployment, error)
	createDeploymentMutex       sync.RWMutex
	createDeploymentArgsForCall []struct {
		arg1 *v1.Deployment
	}
	createDeploymentReturns struct {
		result1 *v1.Deployment
		result2 error
	}
	createDeploymentReturnsOnCall map[int]struct {
		result1 *v1.Deployment
		result2 error
	}
	CreateOrUpdateDeploymentStub        func(*v1.Deployment) (*v1.Deployment, error)
	createOrUpdateDeploymentMutex       sync.RWMutex
	createOrUpdateDeploymentArgsForCall []struct {
		arg1 *v1.Deployment
	}
	createOrUpdateDeploymentReturns struct {
		result1 *v1.Deployment
		result2 error
	}
	createOrUpdateDeploymentReturnsOnCall map[int]struct {
		result1 *v1.Deployment
		result2 error
	}
	CreateRoleStub        func(*v1a.Role) (*v1a.Role, error)
	createRoleMutex       sync.RWMutex
	createRoleArgsForCall []struct {
		arg1 *v1a.Role
	}
	createRoleReturns struct {
		result1 *v1a.Role
		result2 error
	}
	createRoleReturnsOnCall map[int]struct {
		result1 *v1a.Role
		result2 error
	}
	CreateRoleBindingStub        func(*v1a.RoleBinding) (*v1a.RoleBinding, error)
	createRoleBindingMutex       sync.RWMutex
	createRoleBindingArgsForCall []struct {
		arg1 *v1a.RoleBinding
	}
	createRoleBindingReturns struct {
		result1 *v1a.RoleBinding
		result2 error
	}
	createRoleBindingReturnsOnCall map[int]struct {
		result1 *v1a.RoleBinding
		result2 error
	}
	DeleteDeploymentStub        func(string) error
	deleteDeploymentMutex       sync.RWMutex
	deleteDeploymentArgsForCall []struct {
		arg1 string
	}
	deleteDeploymentReturns struct {
		result1 error
	}
	deleteDeploymentReturnsOnCall map[int]struct {
		result1 error
	}
	EnsureServiceAccountStub        func(*v1b.ServiceAccount, ownerutil.Owner) (*v1b.ServiceAccount, error)
	ensureServiceAccountMutex       sync.RWMutex
	ensureServiceAccountArgsForCall []struct {
		arg1 *v1b.ServiceAccount
		arg2 ownerutil.Owner
	}
	ensureServiceAccountReturns struct {
		result1 *v1b.ServiceAccount
		result2 error
	}
	ensureServiceAccountReturnsOnCall map[int]struct {
		result1 *v1b.ServiceAccount
		result2 error
	}
	FindAnyDeploymentsMatchingLabelsStub        func(labels.Selector) ([]*v1.Deployment, error)
	findAnyDeploymentsMatchingLabelsMutex       sync.RWMutex
	findAnyDeploymentsMatchingLabelsArgsForCall []struct {
		arg1 labels.Selector
	}
	findAnyDeploymentsMatchingLabelsReturns struct {
		result1 []*v1.Deployment
		result2 error
	}
	findAnyDeploymentsMatchingLabelsReturnsOnCall map[int]struct {
		result1 []*v1.Deployment
		result2 error
	}
	FindAnyDeploymentsMatchingNamesStub        func([]string) ([]*v1.Deployment, error)
	findAnyDeploymentsMatchingNamesMutex       sync.RWMutex
	findAnyDeploymentsMatchingNamesArgsForCall []struct {
		arg1 []string
	}
	findAnyDeploymentsMatchingNamesReturns struct {
		result1 []*v1.Deployment
		result2 error
	}
	findAnyDeploymentsMatchingNamesReturnsOnCall map[int]struct {
		result1 []*v1.Deployment
		result2 error
	}
	GetServiceAccountByNameStub        func(string) (*v1b.ServiceAccount, error)
	getServiceAccountByNameMutex       sync.RWMutex
	getServiceAccountByNameArgsForCall []struct {
		arg1 string
	}
	getServiceAccountByNameReturns struct {
		result1 *v1b.ServiceAccount
		result2 error
	}
	getServiceAccountByNameReturnsOnCall map[int]struct {
		result1 *v1b.ServiceAccount
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeployment(arg1 *v1.Deployment) (*v1.Deployment, error) {
	fake.createDeploymentMutex.Lock()
	ret, specificReturn := fake.createDeploymentReturnsOnCall[len(fake.createDeploymentArgsForCall)]
	fake.createDeploymentArgsForCall = append(fake.createDeploymentArgsForCall, struct {
		arg1 *v1.Deployment
	}{arg1})
	fake.recordInvocation("CreateDeployment", []interface{}{arg1})
	fake.createDeploymentMutex.Unlock()
	if fake.CreateDeploymentStub != nil {
		return fake.CreateDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createDeploymentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentCallCount() int {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	return len(fake.createDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentCalls(stub func(*v1.Deployment) (*v1.Deployment, error)) {
	fake.createDeploymentMutex.Lock()
	defer fake.createDeploymentMutex.Unlock()
	fake.CreateDeploymentStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentArgsForCall(i int) *v1.Deployment {
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	argsForCall := fake.createDeploymentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentReturns(result1 *v1.Deployment, result2 error) {
	fake.createDeploymentMutex.Lock()
	defer fake.createDeploymentMutex.Unlock()
	fake.CreateDeploymentStub = nil
	fake.createDeploymentReturns = struct {
		result1 *v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateDeploymentReturnsOnCall(i int, result1 *v1.Deployment, result2 error) {
	fake.createDeploymentMutex.Lock()
	defer fake.createDeploymentMutex.Unlock()
	fake.CreateDeploymentStub = nil
	if fake.createDeploymentReturnsOnCall == nil {
		fake.createDeploymentReturnsOnCall = make(map[int]struct {
			result1 *v1.Deployment
			result2 error
		})
	}
	fake.createDeploymentReturnsOnCall[i] = struct {
		result1 *v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeployment(arg1 *v1.Deployment) (*v1.Deployment, error) {
	fake.createOrUpdateDeploymentMutex.Lock()
	ret, specificReturn := fake.createOrUpdateDeploymentReturnsOnCall[len(fake.createOrUpdateDeploymentArgsForCall)]
	fake.createOrUpdateDeploymentArgsForCall = append(fake.createOrUpdateDeploymentArgsForCall, struct {
		arg1 *v1.Deployment
	}{arg1})
	fake.recordInvocation("CreateOrUpdateDeployment", []interface{}{arg1})
	fake.createOrUpdateDeploymentMutex.Unlock()
	if fake.CreateOrUpdateDeploymentStub != nil {
		return fake.CreateOrUpdateDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createOrUpdateDeploymentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentCallCount() int {
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	return len(fake.createOrUpdateDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentCalls(stub func(*v1.Deployment) (*v1.Deployment, error)) {
	fake.createOrUpdateDeploymentMutex.Lock()
	defer fake.createOrUpdateDeploymentMutex.Unlock()
	fake.CreateOrUpdateDeploymentStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentArgsForCall(i int) *v1.Deployment {
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	argsForCall := fake.createOrUpdateDeploymentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentReturns(result1 *v1.Deployment, result2 error) {
	fake.createOrUpdateDeploymentMutex.Lock()
	defer fake.createOrUpdateDeploymentMutex.Unlock()
	fake.CreateOrUpdateDeploymentStub = nil
	fake.createOrUpdateDeploymentReturns = struct {
		result1 *v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateOrUpdateDeploymentReturnsOnCall(i int, result1 *v1.Deployment, result2 error) {
	fake.createOrUpdateDeploymentMutex.Lock()
	defer fake.createOrUpdateDeploymentMutex.Unlock()
	fake.CreateOrUpdateDeploymentStub = nil
	if fake.createOrUpdateDeploymentReturnsOnCall == nil {
		fake.createOrUpdateDeploymentReturnsOnCall = make(map[int]struct {
			result1 *v1.Deployment
			result2 error
		})
	}
	fake.createOrUpdateDeploymentReturnsOnCall[i] = struct {
		result1 *v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRole(arg1 *v1a.Role) (*v1a.Role, error) {
	fake.createRoleMutex.Lock()
	ret, specificReturn := fake.createRoleReturnsOnCall[len(fake.createRoleArgsForCall)]
	fake.createRoleArgsForCall = append(fake.createRoleArgsForCall, struct {
		arg1 *v1a.Role
	}{arg1})
	fake.recordInvocation("CreateRole", []interface{}{arg1})
	fake.createRoleMutex.Unlock()
	if fake.CreateRoleStub != nil {
		return fake.CreateRoleStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRoleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleCallCount() int {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	return len(fake.createRoleArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleCalls(stub func(*v1a.Role) (*v1a.Role, error)) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleArgsForCall(i int) *v1a.Role {
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	argsForCall := fake.createRoleArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleReturns(result1 *v1a.Role, result2 error) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = nil
	fake.createRoleReturns = struct {
		result1 *v1a.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleReturnsOnCall(i int, result1 *v1a.Role, result2 error) {
	fake.createRoleMutex.Lock()
	defer fake.createRoleMutex.Unlock()
	fake.CreateRoleStub = nil
	if fake.createRoleReturnsOnCall == nil {
		fake.createRoleReturnsOnCall = make(map[int]struct {
			result1 *v1a.Role
			result2 error
		})
	}
	fake.createRoleReturnsOnCall[i] = struct {
		result1 *v1a.Role
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBinding(arg1 *v1a.RoleBinding) (*v1a.RoleBinding, error) {
	fake.createRoleBindingMutex.Lock()
	ret, specificReturn := fake.createRoleBindingReturnsOnCall[len(fake.createRoleBindingArgsForCall)]
	fake.createRoleBindingArgsForCall = append(fake.createRoleBindingArgsForCall, struct {
		arg1 *v1a.RoleBinding
	}{arg1})
	fake.recordInvocation("CreateRoleBinding", []interface{}{arg1})
	fake.createRoleBindingMutex.Unlock()
	if fake.CreateRoleBindingStub != nil {
		return fake.CreateRoleBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRoleBindingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingCallCount() int {
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	return len(fake.createRoleBindingArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingCalls(stub func(*v1a.RoleBinding) (*v1a.RoleBinding, error)) {
	fake.createRoleBindingMutex.Lock()
	defer fake.createRoleBindingMutex.Unlock()
	fake.CreateRoleBindingStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingArgsForCall(i int) *v1a.RoleBinding {
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	argsForCall := fake.createRoleBindingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingReturns(result1 *v1a.RoleBinding, result2 error) {
	fake.createRoleBindingMutex.Lock()
	defer fake.createRoleBindingMutex.Unlock()
	fake.CreateRoleBindingStub = nil
	fake.createRoleBindingReturns = struct {
		result1 *v1a.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) CreateRoleBindingReturnsOnCall(i int, result1 *v1a.RoleBinding, result2 error) {
	fake.createRoleBindingMutex.Lock()
	defer fake.createRoleBindingMutex.Unlock()
	fake.CreateRoleBindingStub = nil
	if fake.createRoleBindingReturnsOnCall == nil {
		fake.createRoleBindingReturnsOnCall = make(map[int]struct {
			result1 *v1a.RoleBinding
			result2 error
		})
	}
	fake.createRoleBindingReturnsOnCall[i] = struct {
		result1 *v1a.RoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeployment(arg1 string) error {
	fake.deleteDeploymentMutex.Lock()
	ret, specificReturn := fake.deleteDeploymentReturnsOnCall[len(fake.deleteDeploymentArgsForCall)]
	fake.deleteDeploymentArgsForCall = append(fake.deleteDeploymentArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteDeployment", []interface{}{arg1})
	fake.deleteDeploymentMutex.Unlock()
	if fake.DeleteDeploymentStub != nil {
		return fake.DeleteDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteDeploymentReturns
	return fakeReturns.result1
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentCallCount() int {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return len(fake.deleteDeploymentArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentCalls(stub func(string) error) {
	fake.deleteDeploymentMutex.Lock()
	defer fake.deleteDeploymentMutex.Unlock()
	fake.DeleteDeploymentStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentArgsForCall(i int) string {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	argsForCall := fake.deleteDeploymentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentReturns(result1 error) {
	fake.deleteDeploymentMutex.Lock()
	defer fake.deleteDeploymentMutex.Unlock()
	fake.DeleteDeploymentStub = nil
	fake.deleteDeploymentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallStrategyDeploymentInterface) DeleteDeploymentReturnsOnCall(i int, result1 error) {
	fake.deleteDeploymentMutex.Lock()
	defer fake.deleteDeploymentMutex.Unlock()
	fake.DeleteDeploymentStub = nil
	if fake.deleteDeploymentReturnsOnCall == nil {
		fake.deleteDeploymentReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteDeploymentReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccount(arg1 *v1b.ServiceAccount, arg2 ownerutil.Owner) (*v1b.ServiceAccount, error) {
	fake.ensureServiceAccountMutex.Lock()
	ret, specificReturn := fake.ensureServiceAccountReturnsOnCall[len(fake.ensureServiceAccountArgsForCall)]
	fake.ensureServiceAccountArgsForCall = append(fake.ensureServiceAccountArgsForCall, struct {
		arg1 *v1b.ServiceAccount
		arg2 ownerutil.Owner
	}{arg1, arg2})
	fake.recordInvocation("EnsureServiceAccount", []interface{}{arg1, arg2})
	fake.ensureServiceAccountMutex.Unlock()
	if fake.EnsureServiceAccountStub != nil {
		return fake.EnsureServiceAccountStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.ensureServiceAccountReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountCallCount() int {
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	return len(fake.ensureServiceAccountArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountCalls(stub func(*v1b.ServiceAccount, ownerutil.Owner) (*v1b.ServiceAccount, error)) {
	fake.ensureServiceAccountMutex.Lock()
	defer fake.ensureServiceAccountMutex.Unlock()
	fake.EnsureServiceAccountStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountArgsForCall(i int) (*v1b.ServiceAccount, ownerutil.Owner) {
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	argsForCall := fake.ensureServiceAccountArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountReturns(result1 *v1b.ServiceAccount, result2 error) {
	fake.ensureServiceAccountMutex.Lock()
	defer fake.ensureServiceAccountMutex.Unlock()
	fake.EnsureServiceAccountStub = nil
	fake.ensureServiceAccountReturns = struct {
		result1 *v1b.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) EnsureServiceAccountReturnsOnCall(i int, result1 *v1b.ServiceAccount, result2 error) {
	fake.ensureServiceAccountMutex.Lock()
	defer fake.ensureServiceAccountMutex.Unlock()
	fake.EnsureServiceAccountStub = nil
	if fake.ensureServiceAccountReturnsOnCall == nil {
		fake.ensureServiceAccountReturnsOnCall = make(map[int]struct {
			result1 *v1b.ServiceAccount
			result2 error
		})
	}
	fake.ensureServiceAccountReturnsOnCall[i] = struct {
		result1 *v1b.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabels(arg1 labels.Selector) ([]*v1.Deployment, error) {
	fake.findAnyDeploymentsMatchingLabelsMutex.Lock()
	ret, specificReturn := fake.findAnyDeploymentsMatchingLabelsReturnsOnCall[len(fake.findAnyDeploymentsMatchingLabelsArgsForCall)]
	fake.findAnyDeploymentsMatchingLabelsArgsForCall = append(fake.findAnyDeploymentsMatchingLabelsArgsForCall, struct {
		arg1 labels.Selector
	}{arg1})
	fake.recordInvocation("FindAnyDeploymentsMatchingLabels", []interface{}{arg1})
	fake.findAnyDeploymentsMatchingLabelsMutex.Unlock()
	if fake.FindAnyDeploymentsMatchingLabelsStub != nil {
		return fake.FindAnyDeploymentsMatchingLabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findAnyDeploymentsMatchingLabelsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabelsCallCount() int {
	fake.findAnyDeploymentsMatchingLabelsMutex.RLock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.RUnlock()
	return len(fake.findAnyDeploymentsMatchingLabelsArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabelsCalls(stub func(labels.Selector) ([]*v1.Deployment, error)) {
	fake.findAnyDeploymentsMatchingLabelsMutex.Lock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.Unlock()
	fake.FindAnyDeploymentsMatchingLabelsStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabelsArgsForCall(i int) labels.Selector {
	fake.findAnyDeploymentsMatchingLabelsMutex.RLock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.RUnlock()
	argsForCall := fake.findAnyDeploymentsMatchingLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabelsReturns(result1 []*v1.Deployment, result2 error) {
	fake.findAnyDeploymentsMatchingLabelsMutex.Lock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.Unlock()
	fake.FindAnyDeploymentsMatchingLabelsStub = nil
	fake.findAnyDeploymentsMatchingLabelsReturns = struct {
		result1 []*v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingLabelsReturnsOnCall(i int, result1 []*v1.Deployment, result2 error) {
	fake.findAnyDeploymentsMatchingLabelsMutex.Lock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.Unlock()
	fake.FindAnyDeploymentsMatchingLabelsStub = nil
	if fake.findAnyDeploymentsMatchingLabelsReturnsOnCall == nil {
		fake.findAnyDeploymentsMatchingLabelsReturnsOnCall = make(map[int]struct {
			result1 []*v1.Deployment
			result2 error
		})
	}
	fake.findAnyDeploymentsMatchingLabelsReturnsOnCall[i] = struct {
		result1 []*v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNames(arg1 []string) ([]*v1.Deployment, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.findAnyDeploymentsMatchingNamesMutex.Lock()
	ret, specificReturn := fake.findAnyDeploymentsMatchingNamesReturnsOnCall[len(fake.findAnyDeploymentsMatchingNamesArgsForCall)]
	fake.findAnyDeploymentsMatchingNamesArgsForCall = append(fake.findAnyDeploymentsMatchingNamesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("FindAnyDeploymentsMatchingNames", []interface{}{arg1Copy})
	fake.findAnyDeploymentsMatchingNamesMutex.Unlock()
	if fake.FindAnyDeploymentsMatchingNamesStub != nil {
		return fake.FindAnyDeploymentsMatchingNamesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findAnyDeploymentsMatchingNamesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesCallCount() int {
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	return len(fake.findAnyDeploymentsMatchingNamesArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesCalls(stub func([]string) ([]*v1.Deployment, error)) {
	fake.findAnyDeploymentsMatchingNamesMutex.Lock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.Unlock()
	fake.FindAnyDeploymentsMatchingNamesStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesArgsForCall(i int) []string {
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	argsForCall := fake.findAnyDeploymentsMatchingNamesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesReturns(result1 []*v1.Deployment, result2 error) {
	fake.findAnyDeploymentsMatchingNamesMutex.Lock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.Unlock()
	fake.FindAnyDeploymentsMatchingNamesStub = nil
	fake.findAnyDeploymentsMatchingNamesReturns = struct {
		result1 []*v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) FindAnyDeploymentsMatchingNamesReturnsOnCall(i int, result1 []*v1.Deployment, result2 error) {
	fake.findAnyDeploymentsMatchingNamesMutex.Lock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.Unlock()
	fake.FindAnyDeploymentsMatchingNamesStub = nil
	if fake.findAnyDeploymentsMatchingNamesReturnsOnCall == nil {
		fake.findAnyDeploymentsMatchingNamesReturnsOnCall = make(map[int]struct {
			result1 []*v1.Deployment
			result2 error
		})
	}
	fake.findAnyDeploymentsMatchingNamesReturnsOnCall[i] = struct {
		result1 []*v1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByName(arg1 string) (*v1b.ServiceAccount, error) {
	fake.getServiceAccountByNameMutex.Lock()
	ret, specificReturn := fake.getServiceAccountByNameReturnsOnCall[len(fake.getServiceAccountByNameArgsForCall)]
	fake.getServiceAccountByNameArgsForCall = append(fake.getServiceAccountByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceAccountByName", []interface{}{arg1})
	fake.getServiceAccountByNameMutex.Unlock()
	if fake.GetServiceAccountByNameStub != nil {
		return fake.GetServiceAccountByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceAccountByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameCallCount() int {
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	return len(fake.getServiceAccountByNameArgsForCall)
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameCalls(stub func(string) (*v1b.ServiceAccount, error)) {
	fake.getServiceAccountByNameMutex.Lock()
	defer fake.getServiceAccountByNameMutex.Unlock()
	fake.GetServiceAccountByNameStub = stub
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameArgsForCall(i int) string {
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	argsForCall := fake.getServiceAccountByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameReturns(result1 *v1b.ServiceAccount, result2 error) {
	fake.getServiceAccountByNameMutex.Lock()
	defer fake.getServiceAccountByNameMutex.Unlock()
	fake.GetServiceAccountByNameStub = nil
	fake.getServiceAccountByNameReturns = struct {
		result1 *v1b.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) GetServiceAccountByNameReturnsOnCall(i int, result1 *v1b.ServiceAccount, result2 error) {
	fake.getServiceAccountByNameMutex.Lock()
	defer fake.getServiceAccountByNameMutex.Unlock()
	fake.GetServiceAccountByNameStub = nil
	if fake.getServiceAccountByNameReturnsOnCall == nil {
		fake.getServiceAccountByNameReturnsOnCall = make(map[int]struct {
			result1 *v1b.ServiceAccount
			result2 error
		})
	}
	fake.getServiceAccountByNameReturnsOnCall[i] = struct {
		result1 *v1b.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeInstallStrategyDeploymentInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDeploymentMutex.RLock()
	defer fake.createDeploymentMutex.RUnlock()
	fake.createOrUpdateDeploymentMutex.RLock()
	defer fake.createOrUpdateDeploymentMutex.RUnlock()
	fake.createRoleMutex.RLock()
	defer fake.createRoleMutex.RUnlock()
	fake.createRoleBindingMutex.RLock()
	defer fake.createRoleBindingMutex.RUnlock()
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	fake.ensureServiceAccountMutex.RLock()
	defer fake.ensureServiceAccountMutex.RUnlock()
	fake.findAnyDeploymentsMatchingLabelsMutex.RLock()
	defer fake.findAnyDeploymentsMatchingLabelsMutex.RUnlock()
	fake.findAnyDeploymentsMatchingNamesMutex.RLock()
	defer fake.findAnyDeploymentsMatchingNamesMutex.RUnlock()
	fake.getServiceAccountByNameMutex.RLock()
	defer fake.getServiceAccountByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstallStrategyDeploymentInterface) recordInvocation(key string, args []interface{}) {
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

var _ wrappers.InstallStrategyDeploymentInterface = new(FakeInstallStrategyDeploymentInterface)