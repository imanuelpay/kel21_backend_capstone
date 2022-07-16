// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	dto "backend_capstone/services/product/dto"

	mock "github.com/stretchr/testify/mock"

	models "backend_capstone/models"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ClientFindAll provides a mock function with given fields:
func (_m *Repository) ClientFindAll() (*[]dto.ProductCategory, error) {
	ret := _m.Called()

	var r0 *[]dto.ProductCategory
	if rf, ok := ret.Get(0).(func() *[]dto.ProductCategory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]dto.ProductCategory)
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

// ClientFindAllBySlug provides a mock function with given fields: slug
func (_m *Repository) ClientFindAllBySlug(slug string) (*dto.ProductCategory, error) {
	ret := _m.Called(slug)

	var r0 *dto.ProductCategory
	if rf, ok := ret.Get(0).(func(string) *dto.ProductCategory); ok {
		r0 = rf(slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ProductCategory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: params
func (_m *Repository) FindAll(params ...string) (int64, *[]dto.Product, error) {
	_va := make([]interface{}, len(params))
	for _i := range params {
		_va[_i] = params[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(...string) int64); ok {
		r0 = rf(params...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *[]dto.Product
	if rf, ok := ret.Get(1).(func(...string) *[]dto.Product); ok {
		r1 = rf(params...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*[]dto.Product)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(...string) error); ok {
		r2 = rf(params...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: id
func (_m *Repository) FindById(id string) (*models.ProductResponse, error) {
	ret := _m.Called(id)

	var r0 *models.ProductResponse
	if rf, ok := ret.Get(0).(func(string) *models.ProductResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductResponse)
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

// Insert provides a mock function with given fields: data
func (_m *Repository) Insert(data *models.Product) (*models.ProductResponse, error) {
	ret := _m.Called(data)

	var r0 *models.ProductResponse
	if rf, ok := ret.Get(0).(func(*models.Product) *models.ProductResponse); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Product) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, data
func (_m *Repository) Update(id string, data *models.Product) (*models.ProductResponse, error) {
	ret := _m.Called(id, data)

	var r0 *models.ProductResponse
	if rf, ok := ret.Get(0).(func(string, *models.Product) *models.ProductResponse); ok {
		r0 = rf(id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *models.Product) error); ok {
		r1 = rf(id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStock provides a mock function with given fields: data
func (_m *Repository) UpdateStock(data *dto.UpdateStockDTO) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto.UpdateStockDTO) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateProductBrandCategories provides a mock function with given fields: brandId, categoryId
func (_m *Repository) ValidateProductBrandCategories(brandId string, categoryId string) (string, error) {
	ret := _m.Called(brandId, categoryId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(brandId, categoryId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(brandId, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
