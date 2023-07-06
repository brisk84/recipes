// Code generated by mockery v2.30.16. DO NOT EDIT.

package usecase

import (
	context "context"
	domain "recipes/domain"

	mock "github.com/stretchr/testify/mock"
)

// storageMock is an autogenerated mock type for the storage type
type storageMock struct {
	mock.Mock
}

// DeleteRecipe provides a mock function with given fields: ctx, req
func (_m *storageMock) DeleteRecipe(ctx context.Context, req domain.ID) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindRecipe provides a mock function with given fields: ctx, req
func (_m *storageMock) FindRecipe(ctx context.Context, req domain.Query) ([]domain.Recipe, error) {
	ret := _m.Called(ctx, req)

	var r0 []domain.Recipe
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Query) ([]domain.Recipe, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Query) []domain.Recipe); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Recipe)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Query) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecipes provides a mock function with given fields: ctx
func (_m *storageMock) ListRecipes(ctx context.Context) ([]domain.RecipeForList, error) {
	ret := _m.Called(ctx)

	var r0 []domain.RecipeForList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.RecipeForList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.RecipeForList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RecipeForList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadRecipe provides a mock function with given fields: ctx, req
func (_m *storageMock) ReadRecipe(ctx context.Context, req domain.ID) (domain.Recipe, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Recipe
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) (domain.Recipe, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ID) domain.Recipe); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Recipe)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadUser provides a mock function with given fields: ctx, login
func (_m *storageMock) ReadUser(ctx context.Context, login string) (domain.User, error) {
	ret := _m.Called(ctx, login)

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(ctx, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, login)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoteRecipe provides a mock function with given fields: ctx, req
func (_m *storageMock) VoteRecipe(ctx context.Context, req domain.Vote) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Vote) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteRecipe provides a mock function with given fields: ctx, req
func (_m *storageMock) WriteRecipe(ctx context.Context, req domain.Recipe) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Recipe) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newStorageMock creates a new instance of storageMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStorageMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *storageMock {
	mock := &storageMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}