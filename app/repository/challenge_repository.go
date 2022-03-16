package repository

import (
  "context"

	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeotd/app/dao"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
)

type ChallengeRepository interface {
  ListChallenges(context.Context, entity.Offset) ([]entity.Challenge, int32, error)
  GetChallenge(context.Context, uuid.UUID) (entity.Challenge, error)
  CreateChallenge(context.Context, *entity.Challenge) (error)
  UpdateChallenge(context.Context, uuid.UUID, *entity.Challenge) (error)
  DeleteChallenge(context.Context, uuid.UUID) (error)
}

type challengeRepository struct {
  dao dao.Querier
}

func NewChallengeRepository(dao dao.Querier) *challengeRepository {
  return &challengeRepository{
    dao: dao,
  }
}

func (r *challengeRepository) ListChallenges(ctx context.Context, offset entity.Offset) ([]entity.Challenge, int32, error) {
  var challenges []entity.Challenge
  var count int32

  cs, err := r.dao.ListChallenges(ctx, dao.ListChallengesParams{
    Limit: offset.Limit,
    Offset: offset.Offset,
  })

  if err != nil {
    return challenges, count, err
  }

  c, err := r.dao.CountChallenges(ctx)

  if err != nil {
    return challenges, count, err
  }

  count = int32(c)

  for _, c := range cs {
    i := entity.Challenge{
      ID: c.ID,
      Title: c.Title,
      Description: c.Description,
      CreatedAt: c.CreatedAt,
      UpdatedAt: c.UpdatedAt,
    }

    challenges = append(challenges, i)
  }

  return challenges, count, nil
}

func (r *challengeRepository) GetChallenge(ctx context.Context, id uuid.UUID) (entity.Challenge, error) {
  var challenge entity.Challenge

  c, err := r.dao.GetChallenge(ctx, id)

  if err != nil {
    return challenge, err
  }

  challenge.ID = c.ID
  challenge.Title = c.Title
  challenge.Description = c.Description
  challenge.Body = c.Body
  challenge.CreatedAt = c.CreatedAt
  challenge.UpdatedAt = c.UpdatedAt

  return challenge, nil
}

func (r *challengeRepository) CreateChallenge(ctx context.Context, challenge *entity.Challenge) (error) {
  c, err := r.dao.CreateChallenge(ctx, dao.CreateChallengeParams{
    Title: challenge.Title,
    Body: challenge.Body,
  })

  if err != nil {
    return err
  }

  challenge.ID = c.ID
  challenge.Title = c.Title
  challenge.Description = c.Description
  challenge.Body = c.Body
  challenge.CreatedAt = c.CreatedAt
  challenge.UpdatedAt = c.UpdatedAt

  return nil
}

func (r *challengeRepository) UpdateChallenge(ctx context.Context, id uuid.UUID, challenge *entity.Challenge) (error) {
  c, err := r.dao.UpdateChallenge(ctx, dao.UpdateChallengeParams{
    ID: id,
    Title: challenge.Title,
    Description: challenge.Description,
    Body: challenge.Body,
  })

  if err != nil {
    return err
  }

  challenge.ID = c.ID
  challenge.Title = c.Title
  challenge.Description = c.Description
  challenge.Body = c.Body
  challenge.CreatedAt = c.CreatedAt
  challenge.UpdatedAt = c.UpdatedAt

  return nil
}

func (r *challengeRepository) DeleteChallenge(ctx context.Context, id uuid.UUID) (error) {
  return r.dao.DeleteChallenge(ctx, id)
}
