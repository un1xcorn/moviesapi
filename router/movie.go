package router

import (
	"encoding/json"
	"net/http"

	"github.com/un1xcorn/moviesapi/api"
)

func MovieHandler(w http.ResponseWriter, r *http.Request) {

	db := api.GetDatabase()
	movie_name := r.URL.Query().Get("movie_name")
	if movie_name == "" {
		name, desc := api.GetRandomMovie(db)
		json, err := json.Marshal(map[string]string{"movie_name": name, "movie_desc": desc})
		if err != nil {
			json = nil
		}
		w.Write([]byte(json))
	} else {
		name, desc := api.GetMovieByName(db, movie_name)
		if name != "" {
			json, err := json.Marshal(map[string]string{"movie_name": name, "movie_desc": desc})
			if err != nil {
				json = nil
			}
			w.Write([]byte(json))
		}
	}

}
