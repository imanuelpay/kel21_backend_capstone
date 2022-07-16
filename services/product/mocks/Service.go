// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	dto "backend_capstone/services/product/dto"

	mock "github.com/stretchr/testify/mock"

	models "backend_capstone/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// ClientGetAll provides a mock function with given fields:
func (_m *Service) ClientGetAll() ([]dto.ProductCategory, error) {
	ret := _m.Called()

	var r0 []dto.ProductCategory
	if rf, ok := ret.Get(0).(func() []dto.ProductCategory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ProductCategory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientGetAllBySlug provides a mock function with given fields: slug
func (_m *Service) ClientGetAllBySlug(slug string) (dto.ProductCategory, error) {
	ret := _m.Called(slug)

	var r0 dto.ProductCategory
	if rf, ok := ret.Get(0).(func(string) dto.ProductCategory); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(dto.ProductCategory)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: createproductDTO
func (_m *Service) Create(createproductDTO dto.CraeteProductDTO) (models.ProductResponse, error) {
	ret := _m.Called(createproductDTO)

	var r0 models.ProductResponse
	if rf, ok := ret.Get(0).(func(dto.CraeteProductDTO) models.ProductResponse); ok {
		r0 = rf(createproductDTO)
	} else {
		r0 = ret.Get(0).(models.ProductResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.CraeteProductDTO) error); ok {
		r1 = rf(createproductDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: params
func (_m *Service) GetAll(params ...string) (dto.ResponseBodyProduct, error) {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dto.ResponseBodyProduct
	if rf, ok := ret.Get(0).(func(...string) dto.ResponseBodyProduct); ok {
		r0 = rf(params...)
	} else {
		r0 = ret.Get(0).(dto.ResponseBodyProduct)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllByCategory provides a mock function with given fields: categoryId
func (_m *Service) GetAllByCategory(categoryId string) ([]models.ProductResponse, error) {
	ret := _m.Called(categoryId)

	var r0 []models.ProductResponse
	if rf, ok := ret.Get(0).(func(string) []models.ProductResponse); ok {
		r0 = rf(categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Service) GetById(id string) (models.ProductResponse, error) {
	ret := _m.Called(id)

	var r0 models.ProductResponse
	if rf, ok := ret.Get(0).(func(string) models.ProductResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.ProductResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modify provides a mock function with given fields: id, updateproductDTO
func (_m *Service) Modify(id string, updateproductDTO dto.UpdateProductDTO) (models.ProductResponse, error) {
	ret := _m.Called(id, updateproductDTO)

	var r0 models.ProductResponse
	if rf, ok := ret.Get(0).(func(string, dto.UpdateProductDTO) models.ProductResponse); ok {
		r0 = rf(id, updateproductDTO)
	} else {
		r0 = ret.Get(0).(models.ProductResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, dto.UpdateProductDTO) error); ok {
		r1 = rf(id, updateproductDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyStock provides a mock function with given fields: data
func (_m *Service) ModifyStock(data *dto.UpdateStockDTO) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto.UpdateStockDTO) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: id
func (_m *Service) Remove(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
