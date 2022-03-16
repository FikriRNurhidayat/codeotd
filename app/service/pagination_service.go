package service

import (
  "errors"
  "math"

  "github.com/fikrirnurhidayat/codeotd/app/entity"
)

type PaginationService interface {
  ConvertToOffset(*entity.Pagination) (entity.Offset, error)
  ConvertToPagination(entity.Offset, int32, *entity.Pagination) (error)
}

type paginationService struct {}

func NewPaginationService() *paginationService {
  return &paginationService{}
}

func (s *paginationService) ConvertToOffset(pagination *entity.Pagination) (entity.Offset, error) {
  var o entity.Offset

  if pagination.Page < 0 {
    return o, errors.New("page must be positive")
  }

  if pagination.PageSize < 0 {
    return o, errors.New("page_size must be positive")
  }

  if pagination.Page == 0 {
    pagination.Page = 1
  }

  if pagination.PageSize == 0 {
    pagination.PageSize = 10
  }

  o.Limit = pagination.PageSize
  o.Offset = (pagination.Page - 1) * pagination.PageSize

  return o, nil
}

func (s *paginationService) ConvertToPagination(offset entity.Offset, count int32, pagination *entity.Pagination) (error) {
  pagination.Page = int32(math.Ceil(float64(offset.Offset) / float64(offset.Limit))) + 1
  pagination.PageSize = offset.Limit 
  pagination.PageCount = int32(math.Ceil(float64(count) / float64(offset.Limit)))
  pagination.Total = count
  return nil
}
