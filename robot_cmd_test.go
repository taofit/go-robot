package main

import "testing"

type cmdTest struct {
	table    Rect
	robot    RobotStruct
	action   string
	expected Pos
}

var cmdTests = []cmdTest{
	{Rect{5, 5}, RobotStruct{3, 2, N}, "1,4,1,3,2,3,2,4,1,1,0", Pos{1, 0}},
	{Rect{10, 5}, RobotStruct{3, 2, N}, "1,4,2,3,2,3,1,4,1,4,0", Pos{5, 1}},
	{Rect{13, 6}, RobotStruct{4, 5, N}, "1,4,2,3,2,3,3,4,1,4,0", Pos{6, 5}},
	{Rect{4, 4}, RobotStruct{2, 2, N}, "1,4,1,3,2,3,2,4,1,0", Pos{0, 1}},
}

func TestAdd(t *testing.T) {

	for _, test := range cmdTests {
		if x, y := runCmd(test.table, test.robot, test.action); x != test.expected.Easting || y != test.expected.Northing {
			t.Errorf("Output (%d, %d) not equal to expected (%d, %d)", x, y, test.expected.Easting, test.expected.Northing)
		}
	}
}
