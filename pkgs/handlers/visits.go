package handlers

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/helpers"
	"carrick-js-api/pkgs/logger"
	"carrick-js-api/pkgs/models"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type VisitJsonRequest struct {
	T       string `validate:"omitempty,max=50"`   // tag
	Ci_T    string `validate:"omitempty,max=10"`   // client identifier type
	Ci_V    string `validate:"omitempty,max=200"`  // client identifier value
	R       string `validate:"omitempty,max=2500"` // referrer
	A_U     string `validate:"required,max=2500"`  // current page url
	U_Agent string `validate:"omitempty,max=400"`  // user agent
	Amp     int    `validate:"omitempty,numeric"`
}

func SaveVisitHandler() JSONHandler {
	return JSONHandlerFunc(func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		db := db.GetDBInstance().GetDB()
		logger := logger.GetLoggerInstance()

		validate := validator.New()

		params := mux.Vars(r)
		publisherHash := params["publisher_hash"]

		var jsonRequest VisitJsonRequest
		if err := json.NewDecoder(r.Body).Decode(&jsonRequest); err != nil {
			logger.Error(err)
			return nil, http.StatusBadRequest, errors.New("Bad Request")
		}

		if err := validate.Struct(jsonRequest); err != nil {
			logger.Error(err)
			return nil, http.StatusBadRequest, errors.New("Bad Request")
		}

		ip, _ := helpers.GetIP(r)

		visitItem := models.Visit{
			Publisher_Hash:         publisherHash,
			Tag:                    sql.NullString{String: jsonRequest.T, Valid: jsonRequest.T != ""},
			Click_Identifier_Type:  sql.NullString{String: jsonRequest.Ci_T, Valid: jsonRequest.Ci_T != ""},
			Click_Identifier_Value: sql.NullString{String: jsonRequest.Ci_V, Valid: jsonRequest.Ci_V != ""},
			Traffic_Source_Url:     sql.NullString{String: jsonRequest.R, Valid: jsonRequest.R != ""},
			Publisher_Url:          jsonRequest.A_U,
			Device_Type:            helpers.GetDeviceType(jsonRequest.U_Agent),
			User_Agent:             jsonRequest.U_Agent,
			Is_Amp:                 jsonRequest.Amp != 0,
			Ip:                     sql.NullString{String: ip, Valid: ip != ""},
		}

		if err := db.Create(&visitItem).Error; err != nil {
			logger.Error(err)
			return nil, http.StatusInternalServerError, errors.New("Create error")
		}

		return nil, http.StatusOK, nil
	})
}
