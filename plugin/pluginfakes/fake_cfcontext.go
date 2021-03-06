// This file was generated by counterfeiter
package pluginfakes

import (
	"sync"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/models"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type FakeCFContext struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct{}
	aPIVersionReturns     struct {
		result1 string
	}
	aPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	APIEndpointStub        func() string
	aPIEndpointMutex       sync.RWMutex
	aPIEndpointArgsForCall []struct{}
	aPIEndpointReturns     struct {
		result1 string
	}
	aPIEndpointReturnsOnCall map[int]struct {
		result1 string
	}
	HasAPIEndpointStub        func() bool
	hasAPIEndpointMutex       sync.RWMutex
	hasAPIEndpointArgsForCall []struct{}
	hasAPIEndpointReturns     struct {
		result1 bool
	}
	hasAPIEndpointReturnsOnCall map[int]struct {
		result1 bool
	}
	DopplerEndpointStub        func() string
	dopplerEndpointMutex       sync.RWMutex
	dopplerEndpointArgsForCall []struct{}
	dopplerEndpointReturns     struct {
		result1 string
	}
	dopplerEndpointReturnsOnCall map[int]struct {
		result1 string
	}
	UAAEndpointStub        func() string
	uAAEndpointMutex       sync.RWMutex
	uAAEndpointArgsForCall []struct{}
	uAAEndpointReturns     struct {
		result1 string
	}
	uAAEndpointReturnsOnCall map[int]struct {
		result1 string
	}
	IsLoggedInStub        func() bool
	isLoggedInMutex       sync.RWMutex
	isLoggedInArgsForCall []struct{}
	isLoggedInReturns     struct {
		result1 bool
	}
	isLoggedInReturnsOnCall map[int]struct {
		result1 bool
	}
	UsernameStub        func() string
	usernameMutex       sync.RWMutex
	usernameArgsForCall []struct{}
	usernameReturns     struct {
		result1 string
	}
	usernameReturnsOnCall map[int]struct {
		result1 string
	}
	UserEmailStub        func() string
	userEmailMutex       sync.RWMutex
	userEmailArgsForCall []struct{}
	userEmailReturns     struct {
		result1 string
	}
	userEmailReturnsOnCall map[int]struct {
		result1 string
	}
	UserGUIDStub        func() string
	userGUIDMutex       sync.RWMutex
	userGUIDArgsForCall []struct{}
	userGUIDReturns     struct {
		result1 string
	}
	userGUIDReturnsOnCall map[int]struct {
		result1 string
	}
	UAATokenStub        func() string
	uAATokenMutex       sync.RWMutex
	uAATokenArgsForCall []struct{}
	uAATokenReturns     struct {
		result1 string
	}
	uAATokenReturnsOnCall map[int]struct {
		result1 string
	}
	UAARefreshTokenStub        func() string
	uAARefreshTokenMutex       sync.RWMutex
	uAARefreshTokenArgsForCall []struct{}
	uAARefreshTokenReturns     struct {
		result1 string
	}
	uAARefreshTokenReturnsOnCall map[int]struct {
		result1 string
	}
	RefreshUAATokenStub        func() (string, error)
	refreshUAATokenMutex       sync.RWMutex
	refreshUAATokenArgsForCall []struct{}
	refreshUAATokenReturns     struct {
		result1 string
		result2 error
	}
	refreshUAATokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CurrentOrganizationStub        func() models.OrganizationFields
	currentOrganizationMutex       sync.RWMutex
	currentOrganizationArgsForCall []struct{}
	currentOrganizationReturns     struct {
		result1 models.OrganizationFields
	}
	currentOrganizationReturnsOnCall map[int]struct {
		result1 models.OrganizationFields
	}
	HasTargetedOrganizationStub        func() bool
	hasTargetedOrganizationMutex       sync.RWMutex
	hasTargetedOrganizationArgsForCall []struct{}
	hasTargetedOrganizationReturns     struct {
		result1 bool
	}
	hasTargetedOrganizationReturnsOnCall map[int]struct {
		result1 bool
	}
	CurrentSpaceStub        func() models.SpaceFields
	currentSpaceMutex       sync.RWMutex
	currentSpaceArgsForCall []struct{}
	currentSpaceReturns     struct {
		result1 models.SpaceFields
	}
	currentSpaceReturnsOnCall map[int]struct {
		result1 models.SpaceFields
	}
	HasTargetedSpaceStub        func() bool
	hasTargetedSpaceMutex       sync.RWMutex
	hasTargetedSpaceArgsForCall []struct{}
	hasTargetedSpaceReturns     struct {
		result1 bool
	}
	hasTargetedSpaceReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFContext) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	ret, specificReturn := fake.aPIVersionReturnsOnCall[len(fake.aPIVersionArgsForCall)]
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.aPIVersionReturns.result1
}

func (fake *FakeCFContext) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeCFContext) APIVersionReturns(result1 string) {
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) APIVersionReturnsOnCall(i int, result1 string) {
	fake.APIVersionStub = nil
	if fake.aPIVersionReturnsOnCall == nil {
		fake.aPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.aPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) APIEndpoint() string {
	fake.aPIEndpointMutex.Lock()
	ret, specificReturn := fake.aPIEndpointReturnsOnCall[len(fake.aPIEndpointArgsForCall)]
	fake.aPIEndpointArgsForCall = append(fake.aPIEndpointArgsForCall, struct{}{})
	fake.recordInvocation("APIEndpoint", []interface{}{})
	fake.aPIEndpointMutex.Unlock()
	if fake.APIEndpointStub != nil {
		return fake.APIEndpointStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.aPIEndpointReturns.result1
}

func (fake *FakeCFContext) APIEndpointCallCount() int {
	fake.aPIEndpointMutex.RLock()
	defer fake.aPIEndpointMutex.RUnlock()
	return len(fake.aPIEndpointArgsForCall)
}

func (fake *FakeCFContext) APIEndpointReturns(result1 string) {
	fake.APIEndpointStub = nil
	fake.aPIEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) APIEndpointReturnsOnCall(i int, result1 string) {
	fake.APIEndpointStub = nil
	if fake.aPIEndpointReturnsOnCall == nil {
		fake.aPIEndpointReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.aPIEndpointReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) HasAPIEndpoint() bool {
	fake.hasAPIEndpointMutex.Lock()
	ret, specificReturn := fake.hasAPIEndpointReturnsOnCall[len(fake.hasAPIEndpointArgsForCall)]
	fake.hasAPIEndpointArgsForCall = append(fake.hasAPIEndpointArgsForCall, struct{}{})
	fake.recordInvocation("HasAPIEndpoint", []interface{}{})
	fake.hasAPIEndpointMutex.Unlock()
	if fake.HasAPIEndpointStub != nil {
		return fake.HasAPIEndpointStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hasAPIEndpointReturns.result1
}

func (fake *FakeCFContext) HasAPIEndpointCallCount() int {
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	return len(fake.hasAPIEndpointArgsForCall)
}

func (fake *FakeCFContext) HasAPIEndpointReturns(result1 bool) {
	fake.HasAPIEndpointStub = nil
	fake.hasAPIEndpointReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) HasAPIEndpointReturnsOnCall(i int, result1 bool) {
	fake.HasAPIEndpointStub = nil
	if fake.hasAPIEndpointReturnsOnCall == nil {
		fake.hasAPIEndpointReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasAPIEndpointReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) DopplerEndpoint() string {
	fake.dopplerEndpointMutex.Lock()
	ret, specificReturn := fake.dopplerEndpointReturnsOnCall[len(fake.dopplerEndpointArgsForCall)]
	fake.dopplerEndpointArgsForCall = append(fake.dopplerEndpointArgsForCall, struct{}{})
	fake.recordInvocation("DopplerEndpoint", []interface{}{})
	fake.dopplerEndpointMutex.Unlock()
	if fake.DopplerEndpointStub != nil {
		return fake.DopplerEndpointStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dopplerEndpointReturns.result1
}

func (fake *FakeCFContext) DopplerEndpointCallCount() int {
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	return len(fake.dopplerEndpointArgsForCall)
}

func (fake *FakeCFContext) DopplerEndpointReturns(result1 string) {
	fake.DopplerEndpointStub = nil
	fake.dopplerEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) DopplerEndpointReturnsOnCall(i int, result1 string) {
	fake.DopplerEndpointStub = nil
	if fake.dopplerEndpointReturnsOnCall == nil {
		fake.dopplerEndpointReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dopplerEndpointReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAAEndpoint() string {
	fake.uAAEndpointMutex.Lock()
	ret, specificReturn := fake.uAAEndpointReturnsOnCall[len(fake.uAAEndpointArgsForCall)]
	fake.uAAEndpointArgsForCall = append(fake.uAAEndpointArgsForCall, struct{}{})
	fake.recordInvocation("UAAEndpoint", []interface{}{})
	fake.uAAEndpointMutex.Unlock()
	if fake.UAAEndpointStub != nil {
		return fake.UAAEndpointStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uAAEndpointReturns.result1
}

func (fake *FakeCFContext) UAAEndpointCallCount() int {
	fake.uAAEndpointMutex.RLock()
	defer fake.uAAEndpointMutex.RUnlock()
	return len(fake.uAAEndpointArgsForCall)
}

func (fake *FakeCFContext) UAAEndpointReturns(result1 string) {
	fake.UAAEndpointStub = nil
	fake.uAAEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAAEndpointReturnsOnCall(i int, result1 string) {
	fake.UAAEndpointStub = nil
	if fake.uAAEndpointReturnsOnCall == nil {
		fake.uAAEndpointReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uAAEndpointReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) IsLoggedIn() bool {
	fake.isLoggedInMutex.Lock()
	ret, specificReturn := fake.isLoggedInReturnsOnCall[len(fake.isLoggedInArgsForCall)]
	fake.isLoggedInArgsForCall = append(fake.isLoggedInArgsForCall, struct{}{})
	fake.recordInvocation("IsLoggedIn", []interface{}{})
	fake.isLoggedInMutex.Unlock()
	if fake.IsLoggedInStub != nil {
		return fake.IsLoggedInStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isLoggedInReturns.result1
}

func (fake *FakeCFContext) IsLoggedInCallCount() int {
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	return len(fake.isLoggedInArgsForCall)
}

func (fake *FakeCFContext) IsLoggedInReturns(result1 bool) {
	fake.IsLoggedInStub = nil
	fake.isLoggedInReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) IsLoggedInReturnsOnCall(i int, result1 bool) {
	fake.IsLoggedInStub = nil
	if fake.isLoggedInReturnsOnCall == nil {
		fake.isLoggedInReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isLoggedInReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) Username() string {
	fake.usernameMutex.Lock()
	ret, specificReturn := fake.usernameReturnsOnCall[len(fake.usernameArgsForCall)]
	fake.usernameArgsForCall = append(fake.usernameArgsForCall, struct{}{})
	fake.recordInvocation("Username", []interface{}{})
	fake.usernameMutex.Unlock()
	if fake.UsernameStub != nil {
		return fake.UsernameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.usernameReturns.result1
}

func (fake *FakeCFContext) UsernameCallCount() int {
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	return len(fake.usernameArgsForCall)
}

func (fake *FakeCFContext) UsernameReturns(result1 string) {
	fake.UsernameStub = nil
	fake.usernameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UsernameReturnsOnCall(i int, result1 string) {
	fake.UsernameStub = nil
	if fake.usernameReturnsOnCall == nil {
		fake.usernameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.usernameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UserEmail() string {
	fake.userEmailMutex.Lock()
	ret, specificReturn := fake.userEmailReturnsOnCall[len(fake.userEmailArgsForCall)]
	fake.userEmailArgsForCall = append(fake.userEmailArgsForCall, struct{}{})
	fake.recordInvocation("UserEmail", []interface{}{})
	fake.userEmailMutex.Unlock()
	if fake.UserEmailStub != nil {
		return fake.UserEmailStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.userEmailReturns.result1
}

func (fake *FakeCFContext) UserEmailCallCount() int {
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	return len(fake.userEmailArgsForCall)
}

func (fake *FakeCFContext) UserEmailReturns(result1 string) {
	fake.UserEmailStub = nil
	fake.userEmailReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UserEmailReturnsOnCall(i int, result1 string) {
	fake.UserEmailStub = nil
	if fake.userEmailReturnsOnCall == nil {
		fake.userEmailReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.userEmailReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UserGUID() string {
	fake.userGUIDMutex.Lock()
	ret, specificReturn := fake.userGUIDReturnsOnCall[len(fake.userGUIDArgsForCall)]
	fake.userGUIDArgsForCall = append(fake.userGUIDArgsForCall, struct{}{})
	fake.recordInvocation("UserGUID", []interface{}{})
	fake.userGUIDMutex.Unlock()
	if fake.UserGUIDStub != nil {
		return fake.UserGUIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.userGUIDReturns.result1
}

func (fake *FakeCFContext) UserGUIDCallCount() int {
	fake.userGUIDMutex.RLock()
	defer fake.userGUIDMutex.RUnlock()
	return len(fake.userGUIDArgsForCall)
}

func (fake *FakeCFContext) UserGUIDReturns(result1 string) {
	fake.UserGUIDStub = nil
	fake.userGUIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UserGUIDReturnsOnCall(i int, result1 string) {
	fake.UserGUIDStub = nil
	if fake.userGUIDReturnsOnCall == nil {
		fake.userGUIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.userGUIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAAToken() string {
	fake.uAATokenMutex.Lock()
	ret, specificReturn := fake.uAATokenReturnsOnCall[len(fake.uAATokenArgsForCall)]
	fake.uAATokenArgsForCall = append(fake.uAATokenArgsForCall, struct{}{})
	fake.recordInvocation("UAAToken", []interface{}{})
	fake.uAATokenMutex.Unlock()
	if fake.UAATokenStub != nil {
		return fake.UAATokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uAATokenReturns.result1
}

func (fake *FakeCFContext) UAATokenCallCount() int {
	fake.uAATokenMutex.RLock()
	defer fake.uAATokenMutex.RUnlock()
	return len(fake.uAATokenArgsForCall)
}

func (fake *FakeCFContext) UAATokenReturns(result1 string) {
	fake.UAATokenStub = nil
	fake.uAATokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAATokenReturnsOnCall(i int, result1 string) {
	fake.UAATokenStub = nil
	if fake.uAATokenReturnsOnCall == nil {
		fake.uAATokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uAATokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAARefreshToken() string {
	fake.uAARefreshTokenMutex.Lock()
	ret, specificReturn := fake.uAARefreshTokenReturnsOnCall[len(fake.uAARefreshTokenArgsForCall)]
	fake.uAARefreshTokenArgsForCall = append(fake.uAARefreshTokenArgsForCall, struct{}{})
	fake.recordInvocation("UAARefreshToken", []interface{}{})
	fake.uAARefreshTokenMutex.Unlock()
	if fake.UAARefreshTokenStub != nil {
		return fake.UAARefreshTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uAARefreshTokenReturns.result1
}

func (fake *FakeCFContext) UAARefreshTokenCallCount() int {
	fake.uAARefreshTokenMutex.RLock()
	defer fake.uAARefreshTokenMutex.RUnlock()
	return len(fake.uAARefreshTokenArgsForCall)
}

func (fake *FakeCFContext) UAARefreshTokenReturns(result1 string) {
	fake.UAARefreshTokenStub = nil
	fake.uAARefreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) UAARefreshTokenReturnsOnCall(i int, result1 string) {
	fake.UAARefreshTokenStub = nil
	if fake.uAARefreshTokenReturnsOnCall == nil {
		fake.uAARefreshTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uAARefreshTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCFContext) RefreshUAAToken() (string, error) {
	fake.refreshUAATokenMutex.Lock()
	ret, specificReturn := fake.refreshUAATokenReturnsOnCall[len(fake.refreshUAATokenArgsForCall)]
	fake.refreshUAATokenArgsForCall = append(fake.refreshUAATokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshUAAToken", []interface{}{})
	fake.refreshUAATokenMutex.Unlock()
	if fake.RefreshUAATokenStub != nil {
		return fake.RefreshUAATokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.refreshUAATokenReturns.result1, fake.refreshUAATokenReturns.result2
}

func (fake *FakeCFContext) RefreshUAATokenCallCount() int {
	fake.refreshUAATokenMutex.RLock()
	defer fake.refreshUAATokenMutex.RUnlock()
	return len(fake.refreshUAATokenArgsForCall)
}

func (fake *FakeCFContext) RefreshUAATokenReturns(result1 string, result2 error) {
	fake.RefreshUAATokenStub = nil
	fake.refreshUAATokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFContext) RefreshUAATokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.RefreshUAATokenStub = nil
	if fake.refreshUAATokenReturnsOnCall == nil {
		fake.refreshUAATokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.refreshUAATokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFContext) CurrentOrganization() models.OrganizationFields {
	fake.currentOrganizationMutex.Lock()
	ret, specificReturn := fake.currentOrganizationReturnsOnCall[len(fake.currentOrganizationArgsForCall)]
	fake.currentOrganizationArgsForCall = append(fake.currentOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("CurrentOrganization", []interface{}{})
	fake.currentOrganizationMutex.Unlock()
	if fake.CurrentOrganizationStub != nil {
		return fake.CurrentOrganizationStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.currentOrganizationReturns.result1
}

func (fake *FakeCFContext) CurrentOrganizationCallCount() int {
	fake.currentOrganizationMutex.RLock()
	defer fake.currentOrganizationMutex.RUnlock()
	return len(fake.currentOrganizationArgsForCall)
}

func (fake *FakeCFContext) CurrentOrganizationReturns(result1 models.OrganizationFields) {
	fake.CurrentOrganizationStub = nil
	fake.currentOrganizationReturns = struct {
		result1 models.OrganizationFields
	}{result1}
}

func (fake *FakeCFContext) CurrentOrganizationReturnsOnCall(i int, result1 models.OrganizationFields) {
	fake.CurrentOrganizationStub = nil
	if fake.currentOrganizationReturnsOnCall == nil {
		fake.currentOrganizationReturnsOnCall = make(map[int]struct {
			result1 models.OrganizationFields
		})
	}
	fake.currentOrganizationReturnsOnCall[i] = struct {
		result1 models.OrganizationFields
	}{result1}
}

func (fake *FakeCFContext) HasTargetedOrganization() bool {
	fake.hasTargetedOrganizationMutex.Lock()
	ret, specificReturn := fake.hasTargetedOrganizationReturnsOnCall[len(fake.hasTargetedOrganizationArgsForCall)]
	fake.hasTargetedOrganizationArgsForCall = append(fake.hasTargetedOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("HasTargetedOrganization", []interface{}{})
	fake.hasTargetedOrganizationMutex.Unlock()
	if fake.HasTargetedOrganizationStub != nil {
		return fake.HasTargetedOrganizationStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hasTargetedOrganizationReturns.result1
}

func (fake *FakeCFContext) HasTargetedOrganizationCallCount() int {
	fake.hasTargetedOrganizationMutex.RLock()
	defer fake.hasTargetedOrganizationMutex.RUnlock()
	return len(fake.hasTargetedOrganizationArgsForCall)
}

func (fake *FakeCFContext) HasTargetedOrganizationReturns(result1 bool) {
	fake.HasTargetedOrganizationStub = nil
	fake.hasTargetedOrganizationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) HasTargetedOrganizationReturnsOnCall(i int, result1 bool) {
	fake.HasTargetedOrganizationStub = nil
	if fake.hasTargetedOrganizationReturnsOnCall == nil {
		fake.hasTargetedOrganizationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasTargetedOrganizationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) CurrentSpace() models.SpaceFields {
	fake.currentSpaceMutex.Lock()
	ret, specificReturn := fake.currentSpaceReturnsOnCall[len(fake.currentSpaceArgsForCall)]
	fake.currentSpaceArgsForCall = append(fake.currentSpaceArgsForCall, struct{}{})
	fake.recordInvocation("CurrentSpace", []interface{}{})
	fake.currentSpaceMutex.Unlock()
	if fake.CurrentSpaceStub != nil {
		return fake.CurrentSpaceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.currentSpaceReturns.result1
}

func (fake *FakeCFContext) CurrentSpaceCallCount() int {
	fake.currentSpaceMutex.RLock()
	defer fake.currentSpaceMutex.RUnlock()
	return len(fake.currentSpaceArgsForCall)
}

func (fake *FakeCFContext) CurrentSpaceReturns(result1 models.SpaceFields) {
	fake.CurrentSpaceStub = nil
	fake.currentSpaceReturns = struct {
		result1 models.SpaceFields
	}{result1}
}

func (fake *FakeCFContext) CurrentSpaceReturnsOnCall(i int, result1 models.SpaceFields) {
	fake.CurrentSpaceStub = nil
	if fake.currentSpaceReturnsOnCall == nil {
		fake.currentSpaceReturnsOnCall = make(map[int]struct {
			result1 models.SpaceFields
		})
	}
	fake.currentSpaceReturnsOnCall[i] = struct {
		result1 models.SpaceFields
	}{result1}
}

func (fake *FakeCFContext) HasTargetedSpace() bool {
	fake.hasTargetedSpaceMutex.Lock()
	ret, specificReturn := fake.hasTargetedSpaceReturnsOnCall[len(fake.hasTargetedSpaceArgsForCall)]
	fake.hasTargetedSpaceArgsForCall = append(fake.hasTargetedSpaceArgsForCall, struct{}{})
	fake.recordInvocation("HasTargetedSpace", []interface{}{})
	fake.hasTargetedSpaceMutex.Unlock()
	if fake.HasTargetedSpaceStub != nil {
		return fake.HasTargetedSpaceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hasTargetedSpaceReturns.result1
}

func (fake *FakeCFContext) HasTargetedSpaceCallCount() int {
	fake.hasTargetedSpaceMutex.RLock()
	defer fake.hasTargetedSpaceMutex.RUnlock()
	return len(fake.hasTargetedSpaceArgsForCall)
}

func (fake *FakeCFContext) HasTargetedSpaceReturns(result1 bool) {
	fake.HasTargetedSpaceStub = nil
	fake.hasTargetedSpaceReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) HasTargetedSpaceReturnsOnCall(i int, result1 bool) {
	fake.HasTargetedSpaceStub = nil
	if fake.hasTargetedSpaceReturnsOnCall == nil {
		fake.hasTargetedSpaceReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasTargetedSpaceReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCFContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.aPIEndpointMutex.RLock()
	defer fake.aPIEndpointMutex.RUnlock()
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	fake.uAAEndpointMutex.RLock()
	defer fake.uAAEndpointMutex.RUnlock()
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	fake.userGUIDMutex.RLock()
	defer fake.userGUIDMutex.RUnlock()
	fake.uAATokenMutex.RLock()
	defer fake.uAATokenMutex.RUnlock()
	fake.uAARefreshTokenMutex.RLock()
	defer fake.uAARefreshTokenMutex.RUnlock()
	fake.refreshUAATokenMutex.RLock()
	defer fake.refreshUAATokenMutex.RUnlock()
	fake.currentOrganizationMutex.RLock()
	defer fake.currentOrganizationMutex.RUnlock()
	fake.hasTargetedOrganizationMutex.RLock()
	defer fake.hasTargetedOrganizationMutex.RUnlock()
	fake.currentSpaceMutex.RLock()
	defer fake.currentSpaceMutex.RUnlock()
	fake.hasTargetedSpaceMutex.RLock()
	defer fake.hasTargetedSpaceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCFContext) recordInvocation(key string, args []interface{}) {
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

var _ plugin.CFContext = new(FakeCFContext)
