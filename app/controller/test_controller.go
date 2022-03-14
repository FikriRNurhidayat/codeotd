package controller

import (
  "net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
  "github.com/fikrirnurhidayat/codeot-svc/app/service"
  "github.com/fikrirnurhidayat/codeot-svc/app/dao"
)

type TestController interface {
  HandlePostTest(w http.ResponseWriter, r *http.Request)
  HandlePutTest(w http.ResponseWriter, r *http.Request)
  HandleDeleteTest(w http.ResponseWriter, r *http.Request)
  HandleGetTest(w http.ResponseWriter, r *http.Request)
  HandleGetTests(w http.ResponseWriter, r *http.Request)
}

type testController struct {
  testService service.TestService
}

func NewTestController(testService service.TestService) *testController {
  return &testController{
    testService: testService,
  }
}

// POST /api/v1/tests
func (c *testController) HandlePostTest(w http.ResponseWriter, r *http.Request) {
  var test dao.Test
  var err error

  w.Header().Set("Content-Type", "application/json")

  if err = decodeJSONBody(w, r, &test); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  test, err = c.testService.CreateTest(r.Context(), &test)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusCreated)
  encodeJSONBody(w, ok(test, nil))
}

// GET /api/v1/tests
func (c *testController) HandleGetTests(w http.ResponseWriter, r *http.Request) {
  var tests []dao.Test
  var err error

  w.Header().Set("Content-Type", "application/json")

  tests, err = c.testService.ListTests(r.Context())
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(tests, nil))
}

// GET /api/v1/tests/{id}
func (c *testController) HandleGetTest(w http.ResponseWriter, r *http.Request) {
  var test dao.Test
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  test, err = c.testService.GetTest(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(test, nil))
}

// PUT /api/v1/tests/{id}
func (c *testController) HandlePutTest(w http.ResponseWriter, r *http.Request) {
  var test dao.Test
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  test.ID, err = uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  if err = decodeJSONBody(w, r, &test); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  test, err = c.testService.UpdateTest(r.Context(), &test)

  if err != nil {
    w.WriteHeader(http.StatusUnprocessableEntity)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusOK)
  encodeJSONBody(w, ok(test, nil))
}

// DELETE /api/v1/tests/{id}
func (c *testController) HandleDeleteTest(w http.ResponseWriter, r *http.Request) {
  var err error

  w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)
  id, err := uuid.Parse(params["id"])

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    encodeJSONBody(w, fail(err))
    return
  }

  err = c.testService.DeleteTest(r.Context(), id)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    encodeJSONBody(w, fail(err))
    return
  }

  w.WriteHeader(http.StatusNoContent)
}
