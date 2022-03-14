package repository

import (
  "context"
  "errors"

	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeot-svc/app/dao"
)

type TestRepository interface {
  ListsTest(context.Context) ([]dao.Test, error)
  CreateTest(context.Context, *dao.Test) (dao.Test, error)
  UpdateTest(context.Context, *dao.Test) (dao.Test, error)
  GetTest(context.Context, uuid.UUID) (dao.Test, error)
  DeleteTest(context.Context, uuid.UUID) (error)
}

type testRepository struct {
  dao *dao.Queries
}

func NewTestRepository(dao *dao.Queries) *testRepository {
  return &testRepository{
    dao: dao,
  }
}

func (r *testRepository) CreateTest(ctx context.Context, test *dao.Test) (dao.Test, error) {
  return r.dao.CreateTest(ctx, dao.CreateTestParams{
    Title: test.Title,
    Body: test.Body,
  })
}

func (r *testRepository) UpdateTest(ctx context.Context, test *dao.Test) (dao.Test, error) {
  return r.dao.UpdateTest(ctx, dao.UpdateTestParams{
    ID: test.ID,
    Title: test.Title,
    Body: test.Body,
  })
}

func (r *testRepository) DeleteTest(ctx context.Context, id uuid.UUID) (error) {
  return r.dao.DeleteTest(ctx, id)
}

func (r *testRepository) GetTest(ctx context.Context, id uuid.UUID) (dao.Test, error) {
  t, err := r.dao.GetTest(ctx, id)

  if err != nil {
    return t, errors.New("test not found")
  }

  return t, err
}

func (r *testRepository) ListsTest(ctx context.Context) ([]dao.Test, error) {
  return r.dao.ListTests(ctx)
}
