package handlers

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/helpers"
	"carrick-js-api/pkgs/logger"
	"carrick-js-api/pkgs/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type ClicksTrackingRequest struct {
	Ci_T    string `validate:"omitempty,oneof=gclid fbclid msclkid"` // click identifier type
	Ci_V    string `validate:"required_with=Ci_T,max=200"`           // click identifier value
	T       string `validate:"required,max=50"`                      // tag
	R       string `validate:"omitempty,max=2500"`                   // referrer
	A_U     string `validate:"required,max=2500"`                    // current page url
	C_U     string `validate:"required,max=2500"`                    // click url
	U_Agent string `validate:"omitempty,max=400"`                    // user agent
	Amp     int    `validate:"omitempty,numeric"`
	Anti    int    `validate:"omitempty,numeric"`
}

func SaveClicksTrackingHandler() JSONHandler {
	return JSONHandlerFunc(func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		var err error

		db := db.GetDBInstance().GetDB()
		logger := logger.GetLoggerInstance()

		validate := validator.New()

		params := mux.Vars(r)
		publisherHash := params["publisher_hash"]

		publisher, err := models.PublisherByHash(publisherHash)
		if err != nil {
			return nil, http.StatusNotFound, err
		}

		var jsonRequest ClicksTrackingRequest
		rawBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			if err := json.Unmarshal(rawBody, &jsonRequest); err != nil {
				return nil, http.StatusBadRequest, errors.New("Bad Request")
			}
		}

		//rawBody, _ := ioutil.ReadAll(r.Body)
		err = validate.Struct(jsonRequest)
		if err != nil {
			logger.Error(fmt.Sprintf("%s. Raw json: %s", err, rawBody))
			return nil, http.StatusBadRequest, errors.New("Bad Request")
		}

		if !models.PublisherHasUrl(publisherHash, jsonRequest.A_U) {
			models.FirstOrCreateUrl(publisher, jsonRequest.A_U)
		}

		ip, _ := helpers.GetIP(r)

		Anti := strconv.Itoa(jsonRequest.Anti)
		if Anti == "0" {
			Anti = ""
		}

		clickTrackingItem := models.ClicksTrackingBuffer{
			Publisher_Hash:            publisherHash,
			Tag:                       jsonRequest.T,
			Click_Identifier_Type:     sql.NullString{String: jsonRequest.Ci_T, Valid: jsonRequest.Ci_T != ""},
			Click_Identifier_Value:    sql.NullString{String: jsonRequest.Ci_V, Valid: jsonRequest.Ci_V != ""},
			Traffic_Source_Url:        sql.NullString{String: jsonRequest.R, Valid: jsonRequest.R != ""},
			Publisher_Url:             jsonRequest.A_U,
			Click_Url:                 jsonRequest.C_U,
			User_Agent:                jsonRequest.U_Agent,
			Device_Type:               helpers.GetDeviceType(jsonRequest.U_Agent),
			Is_Amp:                    jsonRequest.Amp != 0,
			Ip:                        sql.NullString{String: ip, Valid: ip != ""},
			Affiliate_Network_Type_Id: sql.NullString{String: Anti, Valid: Anti != ""},
		}

		if err := db.Create(&clickTrackingItem).Error; err != nil {
			logger.Error(err)
			return nil, http.StatusInternalServerError, errors.New("Create error")
		}

		newclickTrackingItem := models.NewClicksTracking{
			Publisher_Id:           publisher.ID,
			Tag:                    jsonRequest.T,
			Click_Identifier_Type:  sql.NullString{String: jsonRequest.Ci_T, Valid: jsonRequest.Ci_T != ""},
			Click_Identifier_Value: sql.NullString{String: jsonRequest.Ci_V, Valid: jsonRequest.Ci_V != ""},
			Publisher_Url:          jsonRequest.A_U,
			Click_Url:              jsonRequest.C_U,
			Traffic_Source_Url:     sql.NullString{String: jsonRequest.R, Valid: jsonRequest.R != ""},
			User_Agent:             jsonRequest.U_Agent,
			Device_Type:            helpers.GetDeviceType(jsonRequest.U_Agent),
			Is_Amp:                 jsonRequest.Amp != 0,
			Ip:                     sql.NullString{String: ip, Valid: ip != ""},
		}

		if err := db.Create(&newclickTrackingItem).Error; err != nil {
			logger.Error(err)
			return nil, http.StatusInternalServerError, errors.New("Create new clicks error")
		}
		return nil, http.StatusOK, nil
	})
}
