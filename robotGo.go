package main

import (
	"bufio"
	"fmt"
	"log"
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

var moveValues = getMoveValues()

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

func getMoveValues() map[string]int {
	var moveValues = make(map[string]int)

	key := fmt.Sprintf("%d-%d", N, 1)
	moveValues[key] = -1
	key = fmt.Sprintf("%d-%d", N, 2)
	moveValues[key] = 1

	key = fmt.Sprintf("%d-%d", S, 1)
	moveValues[key] = 1
	key = fmt.Sprintf("%d-%d", S, 2)
	moveValues[key] = -1

	key = fmt.Sprintf("%d-%d", E, 1)
	moveValues[key] = 1
	key = fmt.Sprintf("%d-%d", E, 2)
	moveValues[key] = -1

	key = fmt.Sprintf("%d-%d", W, 1)
	moveValues[key] = -1
	key = fmt.Sprintf("%d-%d", W, 2)
	moveValues[key] = 1

	return moveValues
}

func runCmd(table Rect, robot RobotStruct, action string) (int, int) {
	actionArr := strings.Split(action, ",")
	for _, order := range actionArr {
		tempOrder, _ := strconv.Atoi(order)

		switch tempOrder {
		case 3, 4:
			robot.Dir = rotate(tempOrder, robot.Dir)
		case 1, 2:
			if proposedPos, isValid := makeAMove(table, robot, tempOrder); isValid {
				robot.X = proposedPos.Easting
				robot.Y = proposedPos.Northing
			}
		case 0:
			// fmt.Printf("final position: %v, %v", robot.X, robot.Y)
			break
		default:
			log.Fatal("invalid order")
		}
	}

	return robot.X, robot.Y
}

func makeAMove(table Rect, robot RobotStruct, forOrBackWard int) (Pos, bool) {
	var proposedPos Pos
	proposedPos.Easting = robot.X
	proposedPos.Northing = robot.Y

	aspect := robot.Dir

	switch aspect {
	case N, S:
		proposedPos.Northing = move(robot, forOrBackWard)
	case E, W:
		proposedPos.Easting = move(robot, forOrBackWard)
	default:
		panic("invalid aspect")
	}

	return proposedPos, table.contains(proposedPos)
}

func (t Rect) contains(pos Pos) bool {
	return 0 <= pos.Easting && pos.Easting <= t.Width && 0 <= pos.Northing && pos.Northing <= t.Height
}

func move(robot RobotStruct, forOrBackWard int) int {
	aspect := robot.Dir
	key := fmt.Sprintf("%d-%d", aspect, forOrBackWard)
	Y := robot.Y
	X := robot.X

	switch aspect {
	case N, S:
		return Y + moveValues[key]
	case E, W:
		return X + moveValues[key]
	default:
		panic("invalid aspect")
	}
}

func initInput(action string) (RobotStruct, Rect) {
	inputArr := strings.Split(action, ",")
	if len(inputArr) != 4 {
		log.Fatal("inital data length is wrong")
	}
	var table Rect
	var robot RobotStruct

	var err error
	table.Width, err = strconv.Atoi(inputArr[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	table.Height, err = strconv.Atoi(inputArr[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	robot.Dir = N
	robot.X, err = strconv.Atoi(inputArr[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	robot.Y, err = strconv.Atoi(inputArr[3])
	if err != nil {
		log.Fatal(err.Error())
	}

	checkInitInput(table, robot)

	return robot, table
}

func checkInitInput(table Rect, robot RobotStruct) {
	if table.Width <= 0 || table.Height <= 0 {
		log.Fatal("table size is wrong")
	}
	if robot.X < 0 || robot.Y < 0 {
		log.Fatal("robot inital location is wrong")
	}

	if table.Width <= robot.X || table.Height <= robot.Y {
		log.Fatal("robot is not in table")
	}
}

func getRobotAndTable() (RobotStruct, Rect) {
	fmt.Printf("Enter table size and robot initial position separated by comma: ")
	reader := bufio.NewReader(os.Stdin)
	initData, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	initData = strings.TrimSpace(initData)
	robot, table := initInput(initData)

	return robot, table
}

func executeCmd(table Rect, robot RobotStruct) {
	fmt.Print("Enter commands separated by comma: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	action, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	action = strings.TrimSpace(action)
	x, y := runCmd(table, robot, action)
	fmt.Printf("final position: %v, %v", x, y)
}

func main() {
	robot, table := getRobotAndTable()
	executeCmd(table, robot)
}
