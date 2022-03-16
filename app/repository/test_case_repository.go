package repository

import (
  "context"

	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeotd/app/dao"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
)


type TestCaseRepository interface {
  GetTestCase(context.Context, uuid.UUID) (entity.TestCase, error) 
  ListTestCases(ctx context.Context, challengeId uuid.UUID, offset entity.Offset) ([]entity.TestCase, int32, error) 
  CreateTestCase(context.Context, *entity.TestCase) (error)
  UpdateTestCase(context.Context, uuid.UUID, *entity.TestCase) (error)
  DeleteTestCase(context.Context, uuid.UUID) (error)
}

type testCaseRepository struct {
  dao dao.Querier
}

func NewTestCaseRepository(dao dao.Querier) *testCaseRepository {
  return &testCaseRepository{
    dao: dao,
  }
}

func (r *testCaseRepository) GetTestCase(ctx context.Context, id uuid.UUID) (entity.TestCase, error) {
  var testCase entity.TestCase
  tc, err := r.dao.GetTestCase(ctx, id)

  if err != nil {
    return testCase, err
  }

  testCase.ID = tc.ID
  testCase.Name = tc.Name
  testCase.Hidden = tc.Hidden
  testCase.Input = tc.Input
  testCase.Output = tc.Output
  testCase.ChallengeID = tc.ChallengeID
  testCase.CreatedAt = tc.CreatedAt
  testCase.UpdatedAt = tc.UpdatedAt

  return testCase, nil
}

func (r *testCaseRepository) CreateTestCase(ctx context.Context, testCase *entity.TestCase) (error) {
  tc, err := r.dao.CreateTestCase(ctx, dao.CreateTestCaseParams{
    Name: testCase.Name,
    Input: testCase.Input,
    Output: testCase.Output,
    Hidden: testCase.Hidden,
  })

  if err != nil {
    return err
  }

  testCase.ID = tc.ID
  testCase.Name = tc.Name
  testCase.Input = tc.Input
  testCase.Output = tc.Output
  testCase.CreatedAt = tc.CreatedAt
  testCase.UpdatedAt = tc.UpdatedAt

  return nil
}

func (r *testCaseRepository) UpdateTestCase(ctx context.Context, id uuid.UUID, testCase *entity.TestCase) (error) {
  tc, err := r.dao.UpdateTestCase(ctx, dao.UpdateTestCaseParams{
    ID: id,
    Name: testCase.Name,
    Input: testCase.Input,
    Output: testCase.Output,
    Hidden: testCase.Hidden,
  })

  if err != nil {
    return err
  }

  testCase.ID = tc.ID
  testCase.Name = tc.Name
  testCase.Input = tc.Input
  testCase.Output = tc.Output
  testCase.CreatedAt = tc.CreatedAt
  testCase.UpdatedAt = tc.UpdatedAt

  return nil
}

func (r *testCaseRepository) DeleteTestCase(ctx context.Context, id uuid.UUID) (error) {
  return r.dao.DeleteTestCase(ctx, id)
}

func (r *testCaseRepository) ListTestCases(ctx context.Context, challengeId uuid.UUID, offset entity.Offset) ([]entity.TestCase, int32, error) {
  var testCases []entity.TestCase
  var err error
  var count int32

  var tcs []dao.ListTestCasesRow
  tcs, err = r.dao.ListTestCases(ctx, dao.ListTestCasesParams{
    ChallengeID: challengeId,
    Offset: offset.Offset,
    Limit: offset.Limit,
  })

  if err != nil {
    return testCases, count, err
  }

  c, err := r.dao.CountTestCases(ctx)

  if err != nil {
    return testCases, count, err
  }

  count = int32(c)

  if err == nil {
    for i := 0; i < len(tcs); i++ {
      tc := tcs[i]
      testCase := entity.TestCase{
        ID: tc.ID, 
        Name: tc.Name,
        Hidden: tc.Hidden,
        Input: tc.Input,
        Output: tc.Output,
        ChallengeID: tc.ChallengeID,
        CreatedAt: tc.CreatedAt,
        UpdatedAt: tc.UpdatedAt,
      }

      testCases = append(testCases, testCase)
    }
  }

  if err != nil {
    return testCases, count, err
  }
  
  return testCases, count, nil
}
