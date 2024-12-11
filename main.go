package main

import (
	"carrick-js-api/pkgs/cache"
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/handlers"
	"carrick-js-api/pkgs/logger"
	"carrick-js-api/pkgs/middlewares"
	"carrick-js-api/pkgs/queue"
	"github.com/gorilla/mux"
	"net/http"
)

func initDB() {
	db.GetDBInstance()
}

func initCache() {
	cache.GetRedisCacheInstance()
}

func main() {
	initDB()
	initCache()

	logger := logger.GetLoggerInstance()
	queue := queue.GetRabbitMQInstance()
	defer queue.Close()

	// TODO: get publisher by hash and save globally in middleware

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/", handlers.IndexHandler)
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	router.Handle("/check-tags/{publisher_hash}", handlers.ServeJSON(handlers.CheckTagsHandler()))
	router.Handle("/get-tag/{publisher_hash}", handlers.Middleware(
		handlers.ServeJSON(handlers.GetTagHandler()),
		middlewares.DomainCheckMiddleware,
	)).Methods(http.MethodGet)
	router.Handle("/send-tracking/{publisher_hash}", handlers.Middleware(
		handlers.ServeJSON(handlers.SaveClicksTrackingHandler()),
		middlewares.DomainCheckMiddleware,
	)).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/visit/{publisher_hash}", handlers.Middleware(
		handlers.ServeJSON(handlers.SaveVisitHandler()),
		middlewares.DomainCheckMiddleware,
	)).Methods(http.MethodPost, http.MethodOptions)

	err := http.ListenAndServe(":5000", router)
	logger.Info("👍 API started. Listening for requests on port 5000...")

	if err != nil {
		logger.Fatal(err)
	}
}
