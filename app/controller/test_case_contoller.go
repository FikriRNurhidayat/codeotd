package controller

import (
  "net/http"

	"github.com/google/uuid"
  "github.com/gorilla/mux"
  "github.com/fikrirnurhidayat/codeotd/app/entity"
  "github.com/fikrirnurhidayat/codeotd/app/service"
)

type TestCaseController interface {
  HandlePostTestCase(w http.ResponseWriter, r *http.Request)
  HandlePutTestCase(w http.ResponseWriter, r *http.Request)
  HandleDeleteTestCase(w http.ResponseWriter, r *http.Request)
  HandleGetTestCase(w http.ResponseWriter, r *http.Request)
  HandleGetTestCases(w http.ResponseWriter, r *http.Request)
}

type testCaseController struct {
  testCaseService service.TestCaseService
} 

func NewTestCaseController(testCaseService service.TestCaseService) *testCaseController {
  return &testCaseController{
    testCaseService: testCaseService,
  }
}

// POST /api/v1/test-cases
func (c *testCaseController) HandlePostTestCase(w http.ResponseWriter, r *http.Request) {
  var testCase entity.TestCase
  var err error

  w.Header().Set("Content-Type", "application/json")

  if err = decodeJSONBody(w, r, &testCase); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.testCaseService.CreateTestCase(r.Context(), &testCase)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusCreated)
  encodeJSONBody(w, ok(testCase, nil))
}

// PUT /api/v1/test-cases/{id}
func (c *testCaseController) HandlePutTestCase(w http.ResponseWriter, r *http.Request) {
  var testCase entity.TestCase
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  if err = decodeJSONBody(w, r, &testCase); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.testCaseService.UpdateTestCase(r.Context(), id, &testCase)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(testCase, nil))
}

// DELETE /api/v1/test-cases/{id}
func (c *testCaseController) HandleDeleteTestCase(w http.ResponseWriter, r *http.Request) {
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.testCaseService.DeleteTestCase(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusNoContent)
}

// GET /api/v1/test-cases/{id}
func (c *testCaseController) HandleGetTestCase(w http.ResponseWriter, r *http.Request) {
  var testCase entity.TestCase
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  testCase, err = c.testCaseService.GetTestCase(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(testCase, nil))
}

// GET /api/v1/tests/{id}/test-cases
func (c *testCaseController) HandleGetTestCases(w http.ResponseWriter, r *http.Request) {
  var res []entity.TestCase
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  challengeId, err := uuid.Parse(params["id"])

  pagination, err := getPagination(r) 

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  res, err = c.testCaseService.ListTestCases(r.Context(), challengeId, &pagination)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(res, map[string]*entity.Pagination{
    "pagination": &pagination,
  }))
}
