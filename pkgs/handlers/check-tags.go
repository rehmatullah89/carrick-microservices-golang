package handlers

import (
	"carrick-js-api/pkgs/db"
	"carrick-js-api/pkgs/models"
	"github.com/gorilla/mux"
	"net/http"
)

func CheckTagsHandler() JSONHandler {
	return JSONHandlerFunc(func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		db := db.GetDBInstance().GetDB()

		params := mux.Vars(r)
		publisherHash := params["publisher_hash"]

		publisher, err := models.PublisherByHash(publisherHash)
		if err != nil {
			return 0, http.StatusNotFound, err
		}

		sqlQuery := `select count(t.id) as tags_count
			from tags t
			where t.publisher_id=@publisher_id
				and t.used=false`

		sqlParams := map[string]interface{}{
			"publisher_id": publisher.ID,
		}

		type Result struct {
			Tags_Count int
		}
		var result Result
		db.Raw(sqlQuery, sqlParams).First(&result)

		return result.Tags_Count, http.StatusOK, nil
	})
}