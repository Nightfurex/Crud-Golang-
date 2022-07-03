package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
 Isbn string `json:"isbn`
	Title string `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct {
 Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies []Movie

func getmovies(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/json")     // tells the client what type content its receving 
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			//json.NewEncoder(w).Encode(message:"The Movie has been deleted")
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range movies {
		 if item.ID == params["id"] {
			   json.NewEncoder(w).Encode(item)
			   return
		}
	}
}

func createmovies(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movies)

}

func updatemovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			movie.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies,movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies,Movie{ID : "1",Isbn : "438227",Title: "Movie One", Director : &Director{Firstname : "Vishnu",Lastname: "Nair"}})
	movies = append(movies , Movie{ID : "2",Isbn : "438228",Title : "Movie Two", Director : &Director{Firstname : "John", Lastname: "Doe"}})
	movies = append (movies ,Movie{ID : "3",Isbn : "438229",Title : "Movie Three", Director : &Director{Firstname: "Manu",Lastname : "Nair"}})
	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",createmovies).Methods("POST")
	r.HandleFunc("/update/{id}",updatemovie).Methods("PUT")
	r.HandleFunc("/delete/{id}",deletemovie).Methods("DELETE")
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}
