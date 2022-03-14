package service

import (
  "context"
  "strings"

  "golang.org/x/net/html"
	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeot-svc/app/dao"
  "github.com/fikrirnurhidayat/codeot-svc/app/repository"
)

type TestService interface {
  ListTests(context.Context) ([]dao.Test, error)
  CreateTest(context.Context, *dao.Test) (dao.Test, error)
  UpdateTest(context.Context, *dao.Test) (dao.Test, error)
  DeleteTest(context.Context, uuid.UUID) (error)
  GetTest(context.Context, uuid.UUID) (dao.Test, error)
}

type testService struct {
  testRepository repository.TestRepository
}

type TestValidationError struct {
  msg string
}

func (e *TestValidationError) Error() string {
  return e.msg
}

func NewTestService(testRepository repository.TestRepository) *testService {
  return &testService{
    testRepository: testRepository,
  }
}

func (s *testService) CreateTest(ctx context.Context, test *dao.Test) (dao.Test, error) {
  var t dao.Test

  if test.Title == "" {
    return t, &TestValidationError{msg: "title is required"} 
  }

  if test.Body == "" {
    return t, &TestValidationError{msg: "body is required"} 
  }

  if _, err := html.Parse(strings.NewReader(test.Body)); err != nil {
    return t, &TestValidationError{msg: "body is not a valid HTML"}
  }


  return s.testRepository.CreateTest(ctx, test)
}

func (s *testService) GetTest(ctx context.Context, id uuid.UUID) (dao.Test, error) {
  return s.testRepository.GetTest(ctx, id)
}

func (s *testService) ListTests(ctx context.Context) ([]dao.Test, error) {
  return s.testRepository.ListsTest(ctx)
}

func (s *testService) DeleteTest(ctx context.Context, id uuid.UUID) (error) {
  return s.testRepository.DeleteTest(ctx, id)
}

func (s *testService) UpdateTest(ctx context.Context, test *dao.Test) (dao.Test, error) {
  return s.testRepository.UpdateTest(ctx, test)
}
