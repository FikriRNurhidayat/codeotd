package service

import (
  "context"
  "strings"

  "golang.org/x/net/html"
	"github.com/google/uuid"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
  "github.com/fikrirnurhidayat/codeotd/app/repository"
)

type ChallengeService interface {
  ListChallenges(context.Context, *entity.Pagination) ([]entity.Challenge, error)
  CreateChallenge(context.Context, *entity.Challenge) (error)
  UpdateChallenge(context.Context, *entity.Challenge) (error)
  DeleteChallenge(context.Context, uuid.UUID) (error)
  GetChallenge(context.Context, uuid.UUID) (entity.Challenge, error)
}

type challengeService struct {
  challengeRepository repository.ChallengeRepository
  paginationService PaginationService
}

type ChallengeValidationError struct {
  msg string
}

func (e *ChallengeValidationError) Error() string {
  return e.msg
}

func NewChallengeService(challengeRepository repository.ChallengeRepository, paginationService PaginationService) *challengeService {
  return &challengeService{
    challengeRepository: challengeRepository,
    paginationService: paginationService,
  }
}

func (s *challengeService) CreateChallenge(ctx context.Context, challenge *entity.Challenge) (error) {
  if challenge.Title == "" {
    return &ChallengeValidationError{msg: "title is required"} 
  }

  if challenge.Body == "" {
    return &ChallengeValidationError{msg: "body is required"} 
  }

  if _, err := html.Parse(strings.NewReader(challenge.Body)); err != nil {
    return &ChallengeValidationError{msg: "body is not a valid HTML"}
  }

  return s.challengeRepository.CreateChallenge(ctx, challenge)
}

func (s *challengeService) GetChallenge(ctx context.Context, id uuid.UUID) (entity.Challenge, error) {
  return s.challengeRepository.GetChallenge(ctx, id)
}

func (s *challengeService) ListChallenges(ctx context.Context, pagination *entity.Pagination) ([]entity.Challenge, error) {
  var challenges []entity.Challenge
  offset, err := s.paginationService.ConvertToOffset(pagination)

  if err != nil {
    return challenges, err
  }

  challenges, count, err := s.challengeRepository.ListChallenges(ctx, offset)

  if err != nil {
    return challenges, err
  }

  s.paginationService.ConvertToPagination(offset, count, pagination)

  return challenges, err
}

func (s *challengeService) DeleteChallenge(ctx context.Context, id uuid.UUID) (error) {
  return s.challengeRepository.DeleteChallenge(ctx, id)
}

func (s *challengeService) UpdateChallenge(ctx context.Context, challenge *entity.Challenge) (error) {
  return s.challengeRepository.UpdateChallenge(ctx, challenge.ID, challenge)
}
