// This file was generated by counterfeiter
package fakes

import (
	"sync"

	routing_api "github.com/cloudfoundry-incubator/routing-api"
	"github.com/cloudfoundry-incubator/routing-api/handlers"
	"github.com/cloudfoundry-incubator/routing-api/models"
)

type FakeRouteValidator struct {
	ValidateCreateStub        func(routes []models.Route, maxTTL int) *routing_api.Error
	validateCreateMutex       sync.RWMutex
	validateCreateArgsForCall []struct {
		routes []models.Route
		maxTTL int
	}
	validateCreateReturns struct {
		result1 *routing_api.Error
	}
	ValidateDeleteStub        func(routes []models.Route) *routing_api.Error
	validateDeleteMutex       sync.RWMutex
	validateDeleteArgsForCall []struct {
		routes []models.Route
	}
	validateDeleteReturns struct {
		result1 *routing_api.Error
	}
	ValidateCreateTcpRouteMappingStub        func(tcpRouteMappings []models.TcpRouteMapping, routerGroups models.RouterGroups) *routing_api.Error
	validateCreateTcpRouteMappingMutex       sync.RWMutex
	validateCreateTcpRouteMappingArgsForCall []struct {
		tcpRouteMappings []models.TcpRouteMapping
		routerGroups     models.RouterGroups
	}
	validateCreateTcpRouteMappingReturns struct {
		result1 *routing_api.Error
	}
	ValidateDeleteTcpRouteMappingStub        func(tcpRouteMappings []models.TcpRouteMapping) *routing_api.Error
	validateDeleteTcpRouteMappingMutex       sync.RWMutex
	validateDeleteTcpRouteMappingArgsForCall []struct {
		tcpRouteMappings []models.TcpRouteMapping
	}
	validateDeleteTcpRouteMappingReturns struct {
		result1 *routing_api.Error
	}
}

func (fake *FakeRouteValidator) ValidateCreate(routes []models.Route, maxTTL int) *routing_api.Error {
	fake.validateCreateMutex.Lock()
	fake.validateCreateArgsForCall = append(fake.validateCreateArgsForCall, struct {
		routes []models.Route
		maxTTL int
	}{routes, maxTTL})
	fake.validateCreateMutex.Unlock()
	if fake.ValidateCreateStub != nil {
		return fake.ValidateCreateStub(routes, maxTTL)
	} else {
		return fake.validateCreateReturns.result1
	}
}

func (fake *FakeRouteValidator) ValidateCreateCallCount() int {
	fake.validateCreateMutex.RLock()
	defer fake.validateCreateMutex.RUnlock()
	return len(fake.validateCreateArgsForCall)
}

func (fake *FakeRouteValidator) ValidateCreateArgsForCall(i int) ([]models.Route, int) {
	fake.validateCreateMutex.RLock()
	defer fake.validateCreateMutex.RUnlock()
	return fake.validateCreateArgsForCall[i].routes, fake.validateCreateArgsForCall[i].maxTTL
}

func (fake *FakeRouteValidator) ValidateCreateReturns(result1 *routing_api.Error) {
	fake.ValidateCreateStub = nil
	fake.validateCreateReturns = struct {
		result1 *routing_api.Error
	}{result1}
}

func (fake *FakeRouteValidator) ValidateDelete(routes []models.Route) *routing_api.Error {
	fake.validateDeleteMutex.Lock()
	fake.validateDeleteArgsForCall = append(fake.validateDeleteArgsForCall, struct {
		routes []models.Route
	}{routes})
	fake.validateDeleteMutex.Unlock()
	if fake.ValidateDeleteStub != nil {
		return fake.ValidateDeleteStub(routes)
	} else {
		return fake.validateDeleteReturns.result1
	}
}

func (fake *FakeRouteValidator) ValidateDeleteCallCount() int {
	fake.validateDeleteMutex.RLock()
	defer fake.validateDeleteMutex.RUnlock()
	return len(fake.validateDeleteArgsForCall)
}

func (fake *FakeRouteValidator) ValidateDeleteArgsForCall(i int) []models.Route {
	fake.validateDeleteMutex.RLock()
	defer fake.validateDeleteMutex.RUnlock()
	return fake.validateDeleteArgsForCall[i].routes
}

func (fake *FakeRouteValidator) ValidateDeleteReturns(result1 *routing_api.Error) {
	fake.ValidateDeleteStub = nil
	fake.validateDeleteReturns = struct {
		result1 *routing_api.Error
	}{result1}
}

func (fake *FakeRouteValidator) ValidateCreateTcpRouteMapping(tcpRouteMappings []models.TcpRouteMapping, routerGroups models.RouterGroups) *routing_api.Error {
	fake.validateCreateTcpRouteMappingMutex.Lock()
	fake.validateCreateTcpRouteMappingArgsForCall = append(fake.validateCreateTcpRouteMappingArgsForCall, struct {
		tcpRouteMappings []models.TcpRouteMapping
		routerGroups     models.RouterGroups
	}{tcpRouteMappings, routerGroups})
	fake.validateCreateTcpRouteMappingMutex.Unlock()
	if fake.ValidateCreateTcpRouteMappingStub != nil {
		return fake.ValidateCreateTcpRouteMappingStub(tcpRouteMappings, routerGroups)
	} else {
		return fake.validateCreateTcpRouteMappingReturns.result1
	}
}

func (fake *FakeRouteValidator) ValidateCreateTcpRouteMappingCallCount() int {
	fake.validateCreateTcpRouteMappingMutex.RLock()
	defer fake.validateCreateTcpRouteMappingMutex.RUnlock()
	return len(fake.validateCreateTcpRouteMappingArgsForCall)
}

func (fake *FakeRouteValidator) ValidateCreateTcpRouteMappingArgsForCall(i int) ([]models.TcpRouteMapping, models.RouterGroups) {
	fake.validateCreateTcpRouteMappingMutex.RLock()
	defer fake.validateCreateTcpRouteMappingMutex.RUnlock()
	return fake.validateCreateTcpRouteMappingArgsForCall[i].tcpRouteMappings, fake.validateCreateTcpRouteMappingArgsForCall[i].routerGroups
}

func (fake *FakeRouteValidator) ValidateCreateTcpRouteMappingReturns(result1 *routing_api.Error) {
	fake.ValidateCreateTcpRouteMappingStub = nil
	fake.validateCreateTcpRouteMappingReturns = struct {
		result1 *routing_api.Error
	}{result1}
}

func (fake *FakeRouteValidator) ValidateDeleteTcpRouteMapping(tcpRouteMappings []models.TcpRouteMapping) *routing_api.Error {
	fake.validateDeleteTcpRouteMappingMutex.Lock()
	fake.validateDeleteTcpRouteMappingArgsForCall = append(fake.validateDeleteTcpRouteMappingArgsForCall, struct {
		tcpRouteMappings []models.TcpRouteMapping
	}{tcpRouteMappings})
	fake.validateDeleteTcpRouteMappingMutex.Unlock()
	if fake.ValidateDeleteTcpRouteMappingStub != nil {
		return fake.ValidateDeleteTcpRouteMappingStub(tcpRouteMappings)
	} else {
		return fake.validateDeleteTcpRouteMappingReturns.result1
	}
}

func (fake *FakeRouteValidator) ValidateDeleteTcpRouteMappingCallCount() int {
	fake.validateDeleteTcpRouteMappingMutex.RLock()
	defer fake.validateDeleteTcpRouteMappingMutex.RUnlock()
	return len(fake.validateDeleteTcpRouteMappingArgsForCall)
}

func (fake *FakeRouteValidator) ValidateDeleteTcpRouteMappingArgsForCall(i int) []models.TcpRouteMapping {
	fake.validateDeleteTcpRouteMappingMutex.RLock()
	defer fake.validateDeleteTcpRouteMappingMutex.RUnlock()
	return fake.validateDeleteTcpRouteMappingArgsForCall[i].tcpRouteMappings
}

func (fake *FakeRouteValidator) ValidateDeleteTcpRouteMappingReturns(result1 *routing_api.Error) {
	fake.ValidateDeleteTcpRouteMappingStub = nil
	fake.validateDeleteTcpRouteMappingReturns = struct {
		result1 *routing_api.Error
	}{result1}
}

var _ handlers.RouteValidator = new(FakeRouteValidator)
