package service

import (
  "context"

	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
  "github.com/fikrirnurhidayat/codeotd/app/repository"
)

type TestCaseService interface {
  CreateTestCase(context.Context, *entity.TestCase) (error)
  UpdateTestCase(context.Context, uuid.UUID, *entity.TestCase) (error)
  DeleteTestCase(context.Context, uuid.UUID) (error)
  GetTestCase(context.Context, uuid.UUID) (entity.TestCase, error)
  ListTestCases(ctx context.Context, challengeId uuid.UUID, pagination *entity.Pagination) ([]entity.TestCase, error)
}

type testCaseService struct {
  testCaseRepository repository.TestCaseRepository
  paginationService PaginationService
}

func NewTestCaseService(testCaseRepository repository.TestCaseRepository, paginationService PaginationService) *testCaseService {
  return &testCaseService{
    testCaseRepository: testCaseRepository,
    paginationService: paginationService,
  }
}

func (s *testCaseService) CreateTestCase(ctx context.Context, testCase *entity.TestCase) (error) {
  return s.testCaseRepository.CreateTestCase(ctx, testCase)
}

func (s *testCaseService) UpdateTestCase(ctx context.Context, id uuid.UUID, testCase *entity.TestCase) (error) {
  return s.testCaseRepository.UpdateTestCase(ctx, id, testCase)
}

func (s *testCaseService) GetTestCase(ctx context.Context, id uuid.UUID) (entity.TestCase, error) {
  return s.testCaseRepository.GetTestCase(ctx, id)
}

func (s *testCaseService) ListTestCases(ctx context.Context, challengeId uuid.UUID, pagination *entity.Pagination) ([]entity.TestCase, error) {
  var testCases []entity.TestCase
  offset, err := s.paginationService.ConvertToOffset(pagination)

  if err != nil {
    return testCases, err
  }

  testCases, count, err := s.testCaseRepository.ListTestCases(ctx, challengeId, offset)

  if err != nil {
    return testCases, err
  }

  s.paginationService.ConvertToPagination(offset, count, pagination)

  return testCases, err
}

func (s *testCaseService) DeleteTestCase(ctx context.Context, id uuid.UUID) (error) {
  return s.testCaseRepository.DeleteTestCase(ctx, id)
}
