package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func main() {
	root := new(Student)
	root.Name = "root"
	root.Age = 18
	root.Score = 100

	left1 := new(Student)
	left1.Name = "left1"
	left1.Age = 18
	left1.Score = 100

	right1 := new(Student)
	right1.Name = "right1"
	right1.Age = 18
	right1.Score = 100

	left2 := new(Student)
	left2.Name = "left2"
	left2.Age = 18
	left2.Score = 100

	root.left = left1
	root.right = right1
	left1.left = left2

	fmt.Println(root)
}