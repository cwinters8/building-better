package main

import (
	"fmt"

	"github.com/cwinters8/building-better/lesson"
)

func main() {
	l := lesson.New("hello")
	fmt.Printf("lesson: %s\n", l.Name)
}
