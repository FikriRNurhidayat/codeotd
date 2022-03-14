package app

import (
  "net/http"
  "github.com/fikrirnurhidayat/codeot-svc/app/config"
  "github.com/fikrirnurhidayat/codeot-svc/app/controller"
  "github.com/fikrirnurhidayat/codeot-svc/app/repository"
  "github.com/fikrirnurhidayat/codeot-svc/app/service"
  "github.com/fikrirnurhidayat/codeot-svc/app/dao"
  "github.com/fikrirnurhidayat/codeot-svc/app/driver/postgres"
  "github.com/gorilla/mux"
)

type Backend struct {
  router *mux.Router
  dao *dao.Queries

  testController controller.TestController
  testService    service.TestService
  testRepository repository.TestRepository
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
  b.testController = controller.NewTestController(b.testService)
}

func (b *Backend) registerRepository(dao *dao.Queries) {
  b.testRepository = repository.NewTestRepository(dao)
}

func (b *Backend) registerRouter() {
  router := mux.NewRouter()

  // Tests Service
  router.HandleFunc("/api/v1/tests", b.testController.HandlePostTest).Methods("POST")
  router.HandleFunc("/api/v1/tests", b.testController.HandleGetTests).Methods("GET")
  router.HandleFunc("/api/v1/tests/{id}", b.testController.HandleGetTest).Methods("GET")
  router.HandleFunc("/api/v1/tests/{id}", b.testController.HandlePutTest).Methods("PUT")
  router.HandleFunc("/api/v1/tests/{id}", b.testController.HandleDeleteTest).Methods("DELETE")

  b.router = router
}

func (b *Backend) registerService() {
  b.testService = service.NewTestService(b.testRepository)
}

func (b *Backend) ServeHTTP(port string) (error) {
  return http.ListenAndServe(port, b.router)
}
