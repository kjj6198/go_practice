package main

type Movie struct {
	Title string,
	Year int `json:"released"`
	Color bool `json:"colo, omitempty`
	Actors []string
}