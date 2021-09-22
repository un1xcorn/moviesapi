package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/un1xcorn/moviesapi/api"
)

func MoviesHandler(w http.ResponseWriter, r *http.Request) {

	db := api.GetDatabase()
	max_movies := r.URL.Query().Get("max_movies")
	if max_movies == "" {
		json, err := json.Marshal(db.Movies)
		if err != nil {
			json = nil
		}
		w.Write([]byte(json))
	} else {
		max_movies_int, err := strconv.Atoi(max_movies)
		if err != nil {
			w.Write([]byte(nil))
		} else {
			json, err := json.Marshal(db.Movies[:max_movies_int])
			if err != nil {
				json = nil
			}
			w.Write([]byte(json))
		}
	}

}
