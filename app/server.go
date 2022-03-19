package app

import (
  "net/http"
  "github.com/fikrirnurhidayat/codeotd/app/config"
  "github.com/fikrirnurhidayat/codeotd/app/controller"
  "github.com/fikrirnurhidayat/codeotd/app/repository"
  "github.com/fikrirnurhidayat/codeotd/app/service"
  "github.com/fikrirnurhidayat/codeotd/app/dao"
  "github.com/fikrirnurhidayat/codeotd/app/driver/postgres"
  "github.com/gorilla/mux"
)

type Backend struct {
  router               *mux.Router
  dao                  *dao.Queries

  challengeController  controller.ChallengeController
  challengeService     service.ChallengeService
  challengeRepository  repository.ChallengeRepository

  paginationService    service.PaginationService

  testCaseController   controller.TestCaseController
  testCaseService      service.TestCaseService
  testCaseRepository   repository.TestCaseRepository
}

func New() (*Backend, error) {
  backend := &Backend{}
  dao, err := registerDatabase()

  if err != nil {
    return nil, err
  }

  backend.registerRepository(dao)
  backend.registerService()
  backend.registerController()
  backend.registerRouter()

  return backend, nil
}

func registerDatabase() (*dao.Queries, error) {
  dbConnStr := config.GetDatabaseConnectionString()
  conn, err := postgres.ConnectDB(dbConnStr)
  
  if err != nil {
    return nil, err
  }

  return dao.New(conn), nil
}

func (b *Backend) registerController() {
  b.challengeController = controller.NewChallengeController(b.challengeService)
  b.testCaseController = controller.NewTestCaseController(b.testCaseService)
}

func (b *Backend) registerRepository(dao *dao.Queries) {
  b.challengeRepository = repository.NewChallengeRepository(dao)
  b.testCaseRepository = repository.NewTestCaseRepository(dao)
}

func (b *Backend) registerRouter() {
  router := mux.NewRouter()

  // Meta Service
  router.HandleFunc("/", controller.HandleGetRoot).Methods("GET")

  // Challenges Service
  router.HandleFunc("/api/v1/challenges", b.challengeController.HandlePostChallenge).Methods("POST")
  router.HandleFunc("/api/v1/challenges", b.challengeController.HandleGetChallenges).Methods("GET")
  router.HandleFunc("/api/v1/challenges/{id}", b.challengeController.HandleGetChallenge).Methods("GET")
  router.HandleFunc("/api/v1/challenges/{id}", b.challengeController.HandlePutChallenge).Methods("PUT")
  router.HandleFunc("/api/v1/challenges/{id}", b.challengeController.HandleDeleteChallenge).Methods("DELETE")

  // Test Cases Service
  router.HandleFunc("/api/v1/challenges/{id}/test-cases", b.testCaseController.HandleGetTestCases).Methods("GET")
  router.HandleFunc("/api/v1/test-cases", b.testCaseController.HandlePostTestCase).Methods("POST")
  router.HandleFunc("/api/v1/test-cases/{id}", b.testCaseController.HandleGetTestCase).Methods("GET")
  router.HandleFunc("/api/v1/test-cases/{id}", b.testCaseController.HandlePutTestCase).Methods("PUT")
  router.HandleFunc("/api/v1/test-cases/{id}", b.testCaseController.HandleDeleteTestCase).Methods("DELETE")

  b.router = router
}

func (b *Backend) registerService() {
  b.paginationService = service.NewPaginationService()
  b.challengeService = service.NewChallengeService(b.challengeRepository, b.paginationService)
  b.testCaseService = service.NewTestCaseService(b.testCaseRepository, b.paginationService)
}

func (b *Backend) ServeHTTP(port string) (error) {
  return http.ListenAndServe(port, controller.LogRequest(b.router))
}
