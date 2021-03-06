package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/pkg/database"
)

const (
	apiPrefix = "/api/v1"
)

var readWriteUpdateMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
var readWriteMethods = []string{"GET", "POST", "OPTIONS"}
var readMethods = []string{"GET", "OPTIONS"}
var writeMethods = []string{"POST", "OPTIONS"}

// TODO: Implement logging

// RegisterRoutes registers all the routes into the mux
func RegisterRoutes(mux *http.ServeMux, db database.Database) {
	mux.Handle(apiPrefix+"/projects", apiContext{db: db, apiHandler: projectsRoute, supportedMethods: readWriteUpdateMethods})
	mux.Handle(apiPrefix+"/project", apiContext{db: db, apiHandler: projectRoute, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/tasks", apiContext{db: db, apiHandler: tasksRoute, supportedMethods: readWriteUpdateMethods})
	mux.Handle(apiPrefix+"/iteration", apiContext{db: db, apiHandler: iterationsRoute, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/currentiteration", apiContext{db: db, apiHandler: currentIterationRoute, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/accounts", apiContext{db: db, apiHandler: accountsRoute, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/login", apiContext{db: db, apiHandler: loginRoute, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/logout", apiContext{db: db, apiHandler: logoutRoute, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/session", apiContext{db: db, apiHandler: sessionRoute, supportedMethods: readMethods})
	// TODO: Non-API routes should register somewhere else
	mux.Handle("/health", http.HandlerFunc(healthCheckHandler))
}
