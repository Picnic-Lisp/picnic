## twingo

A lightweight and concurrent Lisp-dialect with a Golang back-end. Inspired mainly by Emacs-Lisp and Open Dylan (of course also Golang).

Submission project for a talk at the International Lisp Symposium.

### Motivation

I finally stumbled on Golang repo and tried to learn the language. Since I saw it has great reputation and is easy to use I tried to come up with a little week-end project. I then come up with a little Lisp-1 dialect that I call twingo. My first thought was to make simple but powerful language that covers a lot of features such as channels, go interop, a macro
-system and more. 



### Installation 

The twingo programming language relies on a golang backend. In order to get started you should firstly install the Go programming language. After that you simply can install twingo as a global environment by typing this command in a local shell.:

```go
go get github.com/timo-cmd/twingo-Lisp/twingo/cmd/twingo
```

After that you should be able to call the twingo environment from any shell.



## Getting started

After all installation, we want to get more informations about the syntax and about Twingo‘s environment.

Twingo comes served with a simple but powerful REPL and file executor. A REPL is invoked when you simply type the ```twingo``` command with no additional flags. It will start an interactive execution environment where you are able to type twingo code and it will be executed.

With the provided environment you might even want to place your twingo code in separate script files such as: ```file.twingo```. 
Lets get started by creating a little twingo program. First create a new file named ```hello.twingo``` and place this code as its content..:

```Lisp
(package main)
(setq base  (import base))
(setq crypt (import crypt))

(define-function greet (name)
  (let* ch (go:chan-make))
  (go:chan-send "Hello, " name)
  (go:chan-rcve ch)
(greet (twingo))
```
To execute this script file you must open a local shell and type...

```bash
twingo hello.twingo
```

And it will greet with:

```
Hello, twingo
```

### Contributors

-Timo Sarkar lead developer


