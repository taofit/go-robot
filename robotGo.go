package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N Dir = iota
	E
	S
	W
)

const direction_max = 4

// var directions = [direction_max]string{"North", "East", "South", "West"}
var moveValues = map[Dir]int{N: -1, S: 1, E: 1, W: -1}

// func Right() {
// 	Robot.Dir = rotate(3, Robot.Dir)
// }

// func Left() {
// 	Robot.Dir = rotate(4, Robot.Dir)
// }

func rotate(order int, aspect Dir) Dir {
	switch order {
	case 3:
		return (aspect + 1) % direction_max
	case 4:
		return (direction_max + (aspect - 1)) % direction_max
	default:
		panic("invalid order")
	}
}

// func Advance() {
// 	switch Robot.Dir {
// 	case N, S:
// 		Robot.Y = move(Robot.Dir, Robot.X, Robot.Y)
// 	case E, W:
// 		Robot.X = move(Robot.Dir, Robot.X, Robot.Y)
// 	}
// }

func move(aspect Dir, X int, Y int) int {
	switch aspect {
	case N, S:
		return Y + moveValues[aspect]
	case E, W:
		return X + moveValues[aspect]
	default:
		panic("invalid aspect")
	}
}

func runCmd(table Rect, robot RobotStruct, action string) {
	actionArr := strings.Split(action, ",")
	for _, order := range actionArr {
		orderInt, _ := strconv.Atoi(order)

		switch orderInt {
		case 3, 4:
			robot.Dir = rotate(orderInt, robot.Dir)
		case 1:
			if proposedPos, isValid := makeAMove(table, robot); isValid {
				robot.X = proposedPos.Easting
				robot.Y = proposedPos.Northing
			}
		case 0:
			fmt.Printf("final position: %v, %v", robot.X, robot.Y)
		default:
			panic("invalid order")
		}
	}
	fmt.Printf("final position: %v, %v", robot.X, robot.Y)
}

func makeAMove(table Rect, robot RobotStruct) (Pos, bool) {
	var proposedPos Pos
	proposedPos.Easting = robot.X
	proposedPos.Northing = robot.Y

	aspect := robot.Dir

	switch aspect {
	case N, S:
		proposedPos.Northing = move(aspect, robot.X, robot.Y)
	case E, W:
		proposedPos.Easting = move(aspect, robot.X, robot.Y)
	default:
		panic("invalid aspect")
	}

	return proposedPos, table.contains(proposedPos)
}

func (t Rect) contains(pos Pos) bool {
	return pos.Easting <= t.Width && pos.Northing <= t.Height
}

func initInput(action string) (RobotStruct, Rect) {
	inputArr := strings.Split(action, ",")

	var table Rect
	var robot RobotStruct

	table.Width, _ = strconv.Atoi(inputArr[0])
	table.Height, _ = strconv.Atoi(inputArr[1])

	robot.Dir = N
	var err error
	robot.X, _ = strconv.Atoi(inputArr[2])
	robot.Y, err = strconv.Atoi(inputArr[3])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(robot.Dir, robot.X, robot.Y)
	return robot, table
}

func main() {
	fmt.Printf("Enter initial data input separated by comma: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	initData, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	initData = strings.TrimSpace(initData)
	// fmt.Println(initData)
	robot, table := initInput(initData)

	fmt.Print("Enter commands separated by comma: ")
	reader = bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	action, err := reader.ReadString('\n')
	action = strings.Replace(action, "2", "3,3,1", -1)
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	action = strings.TrimSpace(action)
	fmt.Println(action)
	runCmd(table, robot, action)
}
