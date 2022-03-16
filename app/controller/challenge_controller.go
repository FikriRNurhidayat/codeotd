package controller

import (
  "net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
  "github.com/fikrirnurhidayat/codeotd/app/service"
)

type ChallengeController interface {
  HandlePostChallenge(w http.ResponseWriter, r *http.Request)
  HandlePutChallenge(w http.ResponseWriter, r *http.Request)
  HandleDeleteChallenge(w http.ResponseWriter, r *http.Request)
  HandleGetChallenge(w http.ResponseWriter, r *http.Request)
  HandleGetChallenges(w http.ResponseWriter, r *http.Request)
}

type challengeController struct {
  challengeService service.ChallengeService
}

func NewChallengeController(challengeService service.ChallengeService) *challengeController {
  return &challengeController{
    challengeService: challengeService,
  }
}

// POST /api/v1/challenges
func (c *challengeController) HandlePostChallenge(w http.ResponseWriter, r *http.Request) {
  var challenge entity.Challenge
  var err error

  w.Header().Set("Content-Type", "application/json")

  if err = decodeJSONBody(w, r, &challenge); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.challengeService.CreateChallenge(r.Context(), &challenge)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusCreated)
  encodeJSONBody(w, ok(challenge, nil))
}

// GET /api/v1/challenges
func (c *challengeController) HandleGetChallenges(w http.ResponseWriter, r *http.Request) {
  var challenges []entity.Challenge
  var err error

  w.Header().Set("Content-Type", "application/json")

  pagination, err := getPagination(r) 

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  challenges, err = c.challengeService.ListChallenges(r.Context(), &pagination)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(challenges, map[string]*entity.Pagination{
    "pagination": &pagination,
  }))
}

// GET /api/v1/challenges/{id}
func (c *challengeController) HandleGetChallenge(w http.ResponseWriter, r *http.Request) {
  var challenge entity.Challenge
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  challenge, err = c.challengeService.GetChallenge(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(challenge, nil))
}

// PUT /api/v1/challenges/{id}
func (c *challengeController) HandlePutChallenge(w http.ResponseWriter, r *http.Request) {
  var challenge entity.Challenge
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  challenge.ID, err = uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  if err = decodeJSONBody(w, r, &challenge); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.challengeService.UpdateChallenge(r.Context(), &challenge)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(challenge, nil))
}

// DELETE /api/v1/challenges/{id}
func (c *challengeController) HandleDeleteChallenge(w http.ResponseWriter, r *http.Request) {
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.challengeService.DeleteChallenge(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusNoContent)
}
