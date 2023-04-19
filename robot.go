package main

type RobotStruct struct {
	X, Y int
	Dir
}

// var Robot = RobotStruct{
// 	X:   0,
// 	Y:   0,
// 	Dir: 0,
// }

type Dir int
type Pos struct{ Easting, Northing int }
type Rect struct{ Width, Height int }
