// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/mochadwi/go-article/models"

// RatingUsecase is an autogenerated mock type for the RatingUsecase type
type RatingUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *RatingUsecase) Create(_a0 context.Context, _a1 *models.Rating) (*models.Rating, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Rating
	if rf, ok := ret.Get(0).(func(context.Context, *models.Rating) *models.Rating); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rating)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Rating) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, lessonId
func (_m *RatingUsecase) GetByID(ctx context.Context, lessonId int64) (*models.Rating, error) {
	ret := _m.Called(ctx, lessonId)

	var r0 *models.Rating
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Rating); ok {
		r0 = rf(ctx, lessonId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rating)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, lessonId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
