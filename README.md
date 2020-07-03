# picnic
A lightweight and concurrent Lisp-dialect with a Golang back-end.

Submission project for a talk at the International Lisp Conference.

### Installation 

The picnic programming language relies on a golang backend. In order to get started you should firstly install the Go programming language. After that you simply can install picnic as a global environment by typing this command in a local shell.:

```go
go get github.com/timo-cmd/picnic/cmd/picnic
```

After that you should be able to call the Picnic environment from any shell.

### Getting started

After all installation, we want to get more informations about the syntax and about Picnic‘s environment.

Picnic comes served with a simple but powerful REPL and file executor. A REPL is invoked when you simply type the ```picnic``` command with no additional flags. It will start an interactive execution environment where you are able to type picnic code and it will be executed.

With the provided environment you might even want to place your picnic code in separate script files such as: ```file.picnic```. To execute those script files you must open a local shell and type...

```bash

picnic < NameOfYourFile.picnic

```

### Syntax

Lets start explaining the syntax of picnic. As you already know, picnic is a Lisp-1 dialect and will be parenthesis based. So let’s explain the Basics of picnic.

*Basic syntax:* 
