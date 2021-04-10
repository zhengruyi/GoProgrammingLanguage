package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"release"`
	Color bool `json:color,omitempty`
	Actors []string
}

func main() {
	var movie = []Movie{
		{Title:"Casablanca",Year:1942,Color: false,Actors: []string{"Humphrey Bogart"}},
		{Title:"Cool and Hake",Year:1942,Color: false,Actors: []string{"Humphrey Bogart"}},
	}
	data,err := json.MarshalIndent(movie,""," ")
	if err != nil {
		log.Fatal("JSON marshaling failed")
	}
	fmt.Printf("%s,\n",data)
}
