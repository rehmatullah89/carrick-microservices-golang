package middlewares

import (
	"carrick-js-api/pkgs/handlers"
	"carrick-js-api/pkgs/helpers"
	"carrick-js-api/pkgs/logger"
	"carrick-js-api/pkgs/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func DomainCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerInstance()

		has_error := false

		params := mux.Vars(r)
		publisherHash := params["publisher_hash"]

		publisher, err := models.PublisherByHash(publisherHash);
		if err != nil {
			has_error = true

			response := handlers.JSONResponse{
				Data:   nil,
				Error:  fmt.Sprint(err),
				Status: false,
			}
			handlers.SendJsonResponse(w, http.StatusNotFound, response)

			return
		}

		domain, err := helpers.GetDomainFromUrl(r.Referer())
		if err == nil {
			if !models.PublisherHasDomain(domain, publisher) {
				has_error = true

				response := handlers.JSONResponse{
					Data:   nil,
					Error:  "Forbidden. Domain not exists",
					Status: false,
				}
				handlers.SendJsonResponse(w, http.StatusForbidden, response)

				logger.Warning(fmt.Sprintf("Publisher hash '%v' not access to '%v' domain",
					publisherHash, domain))

				return
			}
		} else {
			has_error = true

			response := handlers.JSONResponse{
				Data:   nil,
				Error:  "Forbidden. Domain is empty",
				Status: false,
			}
			handlers.SendJsonResponse(w, http.StatusForbidden, response)

			logger.Warning("Domain is empty")

			return
		}

		if !has_error {
			next.ServeHTTP(w, r)
		}
	})
}
