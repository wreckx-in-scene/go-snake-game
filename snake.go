package main

//point struct for direction
type Point struct {
	X int
	Y int
}

type Snake struct {
	Body      []Point
	Direction Point
}

type Food struct {
	Position Point
}

type Game struct {
	Snake     Snake
	Food      Food
	Score     int
	Highscore int
	Width     int
	Height    int
}
