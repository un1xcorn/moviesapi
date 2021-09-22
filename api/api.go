package api

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	Name        string `json:"movie_name"`
	Description string `json:"movie_desc"`
}

func GetDatabase() *Movies {

	file, err := ioutil.ReadFile("./movies.json")
	if err != nil {
		log.Fatal("Movies json file is missing.")
	}

	var movies Movies
	err = json.Unmarshal(file, &movies)
	if err != nil {
		log.Fatal("An error happened while parsing the movies file.", err)
	}
	return &movies

}

func Init() *Movies {

	return GetDatabase()

}

func GetRandomMovie(database *Movies) (string, string) {

	rdm := rand.Intn(len(database.Movies))
	mv := &database.Movies[rdm]
	return mv.Name, mv.Description

}

func GetMovieByName(database *Movies, name string) (string, string) {

	for _, mv := range database.Movies {
		if mv.Name == name {
			return mv.Name, mv.Description
		}
	}
	return "", ""

}
