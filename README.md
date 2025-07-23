# Go Todo CLI

A command-line to-do application built with **Go** and **SQLite** for managing tasks. This application allows you to add, update, delete, and view tasks directly from the terminal. It uses an SQLite database to store the tasks persistently.

## Features

* Add a new task.
* View all tasks.
* Update task status (mark as complete).
* Delete a task.
* Clear the terminal screen.
* Exit the application.

## Prerequisites

Before running the application, make sure you have the following installed on your system:

* [Go](https://golang.org/dl/) (version 1.16 or later)
* [SQLite](https://www.sqlite.org/download.html) (SQLite is included via the `modernc.org/sqlite` package for Go, but you may want the SQLite CLI for manual database management)

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/todo_cli.git
cd todo_cli
```

### 2. Install Go dependencies:

```bash
go mod tidy
```

### 3. Run the application:

```bash
go run main.go
```

The app will automatically connect to an SQLite database (`todo_cli.db`) and create a `todo` table to store your tasks.

## Usage

Once the application is running, you will be presented with the following options:

* **View all tasks**: Displays a list of all tasks in the database.
* **Add a new task**: Allows you to add a new task to the list.
* **Update task status**: Marks a specific task as completed.
* **Delete a task**: Removes a task from the list.
* **Clear terminal**: Clears the terminal screen.
* **Exit**: Exits the application.

### Example:

```bash
Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-clear, 6-exit
```

### Commands:

* **1**: View all tasks.
* **2**: Add a new task.
* **3**: Update a task (mark as completed).
* **4**: Delete a task.
* **5**: Clear the terminal screen.
* **6**: Exit the application.

### Example Interaction:

```bash
Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-clear, 6-exit
2
Enter a task to add: Buy groceries
success

Enter a number to continue: 1 - view todos, 2-add, 3-update, 4-delete, 5-clear, 6-exit
1
Todo's:
{1 Buy groceries false}
```

## How It Works

* The tasks are stored in an SQLite database (`todo_cli.db`).
* A `todo` table is created with columns for `id`, `task`, and `status`.
* The status column indicates whether the task is completed (`true`) or pending (`false`).

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-name`).
6. Open a Pull Request.

## Go Reference

[![Go Reference](https://pkg.go.dev/badge/github.com/Sudhir0302/todo_cli.svg)](https://pkg.go.dev/github.com/Sudhir0302/todo_cli)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---