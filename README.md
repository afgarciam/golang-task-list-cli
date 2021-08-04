# Golang task list CLI

Simple golang cli, to manage list of task 

## Usage

- In your $GOPATH/src, clone this repo

    `git clone git@github.com:afgarciam/golang-task-list-cli.git`

- Move to folder
  
    `cd ./golang-task-list-cli`

- Get modules
  
    `go get`

- Test work

    `go run todolist.go`

- Add first task

    `go run todolist.go a "{task to add}"`

- List taks 

    `go run todolist.go l`

- Help

    `go run todolist.go -h`

## Install

Place on source folder to next steps

- Install on $GOPATH/bin

    `go intall`

- __Linux__ install on bin folder

    `sudo /usr/local/go/bin/go  build -o /usr/bin/todolist`


## Dependencies

The next list of project, are dependencies for task CLI.

- To create CLI interface [cli](https://pkg.go.dev/github.com/urfave/cli/v2@v2.3.0)
- For database [go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3@v1.14.8)
- Use go-pretty project to print table [go-pretty](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6@v6.2.4)