package main

import (
	"bufio"
	"fmt"
	"os"
)

var io = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

var numOfCourses = 20
var gradeMap = map[string]float64{
	"A+": 4.5,
	"A0": 4,
	"B+": 3.5,
	"B0": 3,
	"C+": 2.5,
	"C0": 2,
	"D+": 1.5,
	"D0": 1,
	"F":  0,
}

func main() {
	var course, grade string
	var credit, gradePoint float64
	var sumOfGradePoint, sumOfCredit float64
	for range 20 {
		fmt.Fscanln(io, &course, &credit, &grade)
		if grade == "P" {
			continue
		}
		gradePoint = gradeMap[grade]
		sumOfGradePoint += credit * gradePoint
		sumOfCredit += credit
	}
	result := sumOfGradePoint / sumOfCredit
	fmt.Fprintln(io, result)
	io.Flush()
}
