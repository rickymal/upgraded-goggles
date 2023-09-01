
package database

import (
	"encoding/json"
	"io/ioutil"

)

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

var DataMovies = []Movie{}
var index = len(DataMovies)

func InitializeDatabase() {
	file, err := ioutil.ReadFile("data/movies.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &DataMovies)
	index = len(DataMovies)
}

func takeSnapshot() {

	file, err := json.Marshal(DataMovies)
	if err != nil {
		return
	}

	
	ioutil.WriteFile("data/movies.json", file, 0644)
}

func AddMovie(movie Movie) int {

	index++
	movie.Id = index
	DataMovies = append(DataMovies, movie)
	takeSnapshot()
	
	return index
}


func DeleteMovie(idx int) {
	DataMovies = append(DataMovies[:idx], DataMovies[idx+1:]...)
	takeSnapshot()
}

func UpdateMovie(idx int, movie Movie) {

	// para manter o indice
	lastIndex := DataMovies[idx].Id
	DataMovies[idx] = movie
	DataMovies[idx].Id =lastIndex
	takeSnapshot()
}