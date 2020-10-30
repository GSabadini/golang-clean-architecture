package infrastructure

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/GSabadini/go-challenge/adapter/api/handler"
	adapterlogger "github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/adapter/repository"
	"github.com/GSabadini/go-challenge/infrastructure/database"
	"github.com/GSabadini/go-challenge/infrastructure/logger"
	"github.com/GSabadini/go-challenge/infrastructure/router"
	"github.com/GSabadini/go-challenge/usecase"
)

// HTTPServer define an application structure
type HTTPServer struct {
	database *database.MongoHandler
	logger   adapterlogger.Logger
	router   router.Router
}

// NewHTTPServer creates new HTTPServer with its dependencies
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		database: database.NewMongoHandler(),
		logger:   logger.NewLogrus(),
		router:   router.NewMux(),
	}
}

// Start run the application
func (a HTTPServer) Start() {
	a.router.GET("/health", healthCheck)

	a.router.POST("/users", a.createUserHandler())
	a.router.GET("/users/{user_id}", a.findUserByIDHandler())

	a.router.SERVE(os.Getenv("APP_PORT"))
}

func (a HTTPServer) createUserHandler() http.HandlerFunc {
	uc := usecase.NewCreateUserInteractor(
		repository.NewCreateUserRepository(a.database),
		presenter.NewCreateUserPresenter())

	return handler.NewCreateUserHandler(uc, a.logger).Handle
}

func (a HTTPServer) findUserByIDHandler() http.HandlerFunc {
	uc := usecase.NewFindUserByIDInteractor(
		repository.NewFindUserByIDUserRepository(a.database),
		presenter.NewFindUserByIDPresenter())

	return handler.NewFindUserByIDHandler(uc, a.logger).Handle
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: http.StatusText(http.StatusOK)})
}
