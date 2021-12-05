// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "github.com/rayl17/academy-go-q42021/model"
	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// GetPokemonByID provides a mock function with given fields: id
func (_m *ServiceInterface) GetPokemonByID(id string) (*model.Pokemon, error) {
	ret := _m.Called(id)

	var r0 *model.Pokemon
	if rf, ok := ret.Get(0).(func(string) *model.Pokemon); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pokemon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPokemonByName provides a mock function with given fields: name
func (_m *ServiceInterface) GetPokemonByName(name string) (*model.Pokemon, error) {
	ret := _m.Called(name)

	var r0 *model.Pokemon
	if rf, ok := ret.Get(0).(func(string) *model.Pokemon); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pokemon)
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

// SearchPokemonApi provides a mock function with given fields: name
func (_m *ServiceInterface) SearchPokemonApi(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getPokemons provides a mock function with given fields:
func (_m *ServiceInterface) getPokemons() []model.Pokemon {
	ret := _m.Called()

	var r0 []model.Pokemon
	if rf, ok := ret.Get(0).(func() []model.Pokemon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Pokemon)
		}
	}

	return r0
}