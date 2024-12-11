package handlers

import (
	"carrick-js-api/pkgs/cache"
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/helpers"
	"carrick-js-api/pkgs/logger"
	"carrick-js-api/pkgs/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

func tagForOrganicTraffic(publisher models.Publisher, publisherUrl string, referrerDomain string) (models.TagResult, error) {
	db := db.GetDBInstance().GetDB()
	logger := logger.GetLoggerInstance()
	cache := cache.GetRedisCacheInstance()

	var (
		tagResult models.TagResult
	)

	trafficSource, err := models.GetTrafficSourceByDomain(referrerDomain)
	if err != nil {
		return models.TagResult{}, err
	}

	sqlQuery := `SELECT pu.tag
		FROM new_publisher_url pu
		left join urls u on pu.url_id = u.id
		where pu.publisher_id=@publisher_id
			and pu.traffic_source_id=@traffic_source_id
			and u.url_path=@publisher_url
			and pu.is_paid = false
		limit 1
		`
	sqlParams := map[string]interface{}{
		"publisher_id":      publisher.ID,
		"publisher_url":     publisherUrl,
		"traffic_source_id": trafficSource.ID,
	}
	err = db.Raw(sqlQuery, sqlParams).First(&tagResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err)
		return models.TagResult{}, errors.New("DB Error")
	}

	if tagResult.Tag == "" {

		if publisher.ID != 0 {
			newPublisherCacheKey := fmt.Sprintf("publisher-%d", publisher.ID)

			if err := cache.Get(newPublisherCacheKey, &publisher); err != nil || publisher.ID == 0 {
				db.First(&publisher, publisher.ID)

				if publisher.ID != 0 {
					cache.Set(newPublisherCacheKey, publisher, config.AppConfig.CacheTTL)
				}
			}
		}

		if trafficSource.ID != 0 {
			trafficSourceCacheKey := fmt.Sprintf("traffic_source-%d", trafficSource.ID)

			if err := cache.Get(trafficSourceCacheKey, &trafficSource); err != nil || trafficSource.ID == 0 {
				db.First(&trafficSource, trafficSource.ID)

				if trafficSource.ID != 0 {
					cache.Set(trafficSourceCacheKey, trafficSource, config.AppConfig.CacheTTL)
				}
			}
		}

		if publisher.ID != 0 && trafficSource.ID != 0 {
			tagResult, _ = models.AttachTagToUrlAndTrafficSource(publisher, trafficSource, publisherUrl)
		}
	}

	return tagResult, nil
}

func GetTagHandler() JSONHandler {
	return JSONHandlerFunc(func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		var tagResult models.TagResult
		validate := validator.New()
		logger := logger.GetLoggerInstance()

		type query struct {
			T_Type string `validate:"required,oneof=organic paid"`
			R      string
			A_u    string
		}

		params := mux.Vars(r)
		publisherHash := params["publisher_hash"]

		publisher, err := models.PublisherByHash(publisherHash)
		if err != nil {
			return nil, http.StatusNotFound, errors.New("Publisher not found")
		}

		query_vars := r.URL.Query()
		queryParams := query{
			T_Type: query_vars.Get("t_type"),
			R:      query_vars.Get("r"),
			A_u:    query_vars.Get("a_u"),
		}

		if err := validate.Struct(queryParams); err != nil {
			logger.Error(err)
			return nil, http.StatusBadRequest, errors.New("Bad Request. Traffic type not found or invalid")
		}

		trafficType := queryParams.T_Type

		switch trafficType {
		case "organic":
			// organic traffic
			domain, err := helpers.GetDomainFromUrl(queryParams.R)
			if err != nil {
				return nil, http.StatusBadRequest, err
			}

			publisherUrl, err := helpers.GetPathFromUrl(queryParams.A_u)
			if err != nil {
				return nil, http.StatusBadRequest, err
			} else {
				tagResult, _ = tagForOrganicTraffic(publisher, publisherUrl, domain)
			}
		case "paid":
			// paid traffic
			tagResult = models.GetTagForTrafficType(publisher, "paid")
		}

		if tagResult.Tag == "" {
			return nil, http.StatusNotFound, errors.New("Tag not found")
		}

		return tagResult.Tag, http.StatusOK, nil
	})
}
