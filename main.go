package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("todo")
	for {
		fmt.Println("Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-exit")
		var op int
		fmt.Scan(&op)
		reader.ReadString('\n')
		switch op {
		case 1:
			fmt.Println("Todo's: ")
		case 2:
			fmt.Print("Enter a task to add: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			fmt.Println(task)
		case 3:
			fmt.Print("Enter a task to update: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			fmt.Println(task)
		case 4:
			// fmt.Println("delete")
			fmt.Print("Enter a task to delete: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			fmt.Println(task)
		case 5:
			fmt.Println("exit-ingg...")
			time.Sleep(2 * time.Second)
			return
		}
	}
}
