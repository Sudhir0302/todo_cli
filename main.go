package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Sudhir0302/todo_cli/config"
	"github.com/Sudhir0302/todo_cli/handlers"
)

func main() {

	//connect to db and create table
	config.Conn()
	handlers.Init()

	todoHand := handlers.TodoHandler{DB: config.DB}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("todo_cli")
	for {
		fmt.Println("Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-clear, 6-exit")
		var op int
		fmt.Scan(&op)
		reader.ReadString('\n')
		switch op {
		case 1:
			fmt.Println("Todo's: ")
			todoHand.View()
		case 2:
			fmt.Print("Enter a task to add: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			// fmt.Println(task)
			res := todoHand.Add(task)
			fmt.Println(res)
		case 3:
			fmt.Print("Enter a task to update: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			// fmt.Println(task)
			res := todoHand.Update(task)
			fmt.Println(res)
		case 4:
			// fmt.Println("delete")
			fmt.Print("Enter a task to delete: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			// fmt.Println(task)
			res := todoHand.Delete(task)
			fmt.Println(res)
		case 5:
			var cmd1 *exec.Cmd
			if runtime.GOOS == "windows" {
				cmd1 = exec.Command("cmd", "/c", "cls")
			} else {
				cmd1 = exec.Command("clear")
			}
			cmd1.Stdout = os.Stdout
			cmd1.Run()
		case 6:
			fmt.Println("exit-ingg...")
			time.Sleep(2 * time.Second)
			return
		default:
			fmt.Println("enter a valid number")
		}
	}
}
