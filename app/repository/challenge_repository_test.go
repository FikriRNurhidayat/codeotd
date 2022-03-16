package repository

import (
  "context"
  "errors"
  "testing"
  "time"

  "github.com/fikrirnurhidayat/codeotd/app/entity"
  "github.com/fikrirnurhidayat/codeotd/app/dao"
  "github.com/fikrirnurhidayat/codeotd/mocks/dao"

	"github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

type mockContext struct {}

func TestCreateChallenge(t *testing.T) {
  var ctx context.Context
  challenge := &entity.Challenge{}
  mockDao := new(mocks.Querier)
  mockCreateChallengeRow := dao.CreateChallengeRow{
    ID: uuid.New(),
    Title: "Something",
    Body: "Something",
    Description: "Something",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  mockDao.On("CreateChallenge", ctx, mock.Anything).Return(mockCreateChallengeRow, nil)

  challengeRepository := NewChallengeRepository(mockDao)

  err := challengeRepository.CreateChallenge(ctx, challenge)

  assert.Nil(t, err)
  assert.Equal(t, challenge.ID, mockCreateChallengeRow.ID)
  assert.Equal(t, challenge.Title, mockCreateChallengeRow.Title)
  assert.Equal(t, challenge.Description, mockCreateChallengeRow.Description)
  assert.Equal(t, challenge.Body, mockCreateChallengeRow.Body)
  assert.Equal(t, challenge.CreatedAt, mockCreateChallengeRow.CreatedAt)
  assert.Equal(t, challenge.UpdatedAt, mockCreateChallengeRow.UpdatedAt)
}

func TestGetChallenge(t *testing.T) {
  var ctx context.Context
  mockDao := new(mocks.Querier)
  mockGetChallengeRow := dao.GetChallengeRow{
    ID: uuid.New(),
    Title: "Something",
    Body: "Something",
    Description: "Something",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  mockDao.On("GetChallenge", ctx, mockGetChallengeRow.ID).Return(mockGetChallengeRow, nil)

  challengeRepository := NewChallengeRepository(mockDao)

  c, err := challengeRepository.GetChallenge(ctx, mockGetChallengeRow.ID)

  assert.Nil(t, err)
  assert.Equal(t, c.ID, mockGetChallengeRow.ID)
  assert.Equal(t, c.Title, mockGetChallengeRow.Title)
  assert.Equal(t, c.Description, mockGetChallengeRow.Description)
  assert.Equal(t, c.Body, mockGetChallengeRow.Body)
  assert.Equal(t, c.CreatedAt, mockGetChallengeRow.CreatedAt)
  assert.Equal(t, c.UpdatedAt, mockGetChallengeRow.UpdatedAt)
}

func TestErrorGetChallenge(t *testing.T) {
  var ctx context.Context
  mockDao := new(mocks.Querier)
  mockGetChallengeRow := dao.GetChallengeRow{
    ID: uuid.New(),
    Title: "Something",
    Body: "Something",
    Description: "Something",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  mockDao.On("GetChallenge", ctx, mockGetChallengeRow.ID).Return(mockGetChallengeRow, errors.New("Something is error!"))

  challengeRepository := NewChallengeRepository(mockDao)

  _, err := challengeRepository.GetChallenge(ctx, mockGetChallengeRow.ID)

  assert.NotNil(t, err)
}

func TestListChallenges(t *testing.T) {
  var ctx context.Context
  var offset entity.Offset
  var count int64 = 1
  mockDao := new(mocks.Querier)
  mcs := []dao.ListChallengesRow{
    {
      ID: uuid.New(),
      Title: "Something",
      Description: "Something",
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    },
  }

  mockDao.On("ListChallenges", ctx, mock.Anything).Return(mcs, nil)
  mockDao.On("CountChallenges", ctx).Return(count, nil)

  challengeRepository := NewChallengeRepository(mockDao)

  _, cnt, err := challengeRepository.ListChallenges(ctx, offset)

  assert.Nil(t, err)
  assert.Equal(t, cnt, int32(count))
}

func TestListChallengesErrorDaoListChallenges(t *testing.T) {
  var ctx context.Context
  var offset entity.Offset
  var count int64 = 1
  mockDao := new(mocks.Querier)
  mcs := []dao.ListChallengesRow{}

  mockDao.On("ListChallenges", ctx, mock.Anything).Return(mcs, errors.New("failed to retrieve challenges"))
  mockDao.On("CountChallenges", ctx).Return(count, nil)

  challengeRepository := NewChallengeRepository(mockDao)

  _, _, err := challengeRepository.ListChallenges(ctx, offset)

  assert.NotNil(t, err)
}

func TestListChallengesErrorDaoCountChallenges(t *testing.T) {
  var ctx context.Context
  var offset entity.Offset
  var count int64 = 1
  mockDao := new(mocks.Querier)
  mcs := []dao.ListChallengesRow{}

  mockDao.On("ListChallenges", ctx, mock.Anything).Return(mcs, nil)
  mockDao.On("CountChallenges", ctx).Return(count, errors.New("failed to count challenges"))

  challengeRepository := NewChallengeRepository(mockDao)

  _, _, err := challengeRepository.ListChallenges(ctx, offset)

  assert.NotNil(t, err)
}
