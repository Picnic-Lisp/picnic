# picnic


A lightweight and concurrent Lisp-dialect with a Golang back-end. Inspired mainly by Emacs-Lisp and Open Dylan (of course also Golang).

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

*Basic syntax:* arithmethics;

```lisp
+(1 1)
-(1 1)
*(1 1)
/(1 1)
```

Simple right? Now lets go ahead with builtins...

```lisp
(print "hello, world!")
(length [0, 1, 2, 3, 4])
```

Print is used to print value to the screen. Length is used to detect the length of an array, map or hash. Of course there are many more builtins but I won’t explain them all here. Lets head over data types:

```lisp
Bool: true, false
Int: 12
String: "string"
Array: [a, b, c]
Null: Null
```

The null value is thought to represent a non-existing or undefined value. It’s really simple as long as you’ll get into it... now functions.:

```lisp
(defun greet (name)
  (print +("Hello, " name)))

greet(timo)
```

Functions in picnic are defined with a name and a body which contains its functionality. This function is named greet and will display hello, timo. Lets start explaining macros...

```lisp
(defmacro greet (name)
  (print +("hello, " name)))

(greet timo)
```
Macros are almost similar to functions. Instead defining a functional algorithm, a macros does define a part of syntax. The result of this program is the same as the previous one. Lets head on the Go interop System...:

```lisp
(package main)
(setq fmt  (go:import fmt))
(setq time (go:import time))

(defun main ()
  (go:fmt.println "hello from go!"))
```

Picnic comes served with a fully working and powerful Go-interop system that lets users allow to embed go code in picnic and evaluate them as native picnic code. Now we’re heading to Channels and Coroutines...:

```lisp
(package main)
(setq fmt (go:import fmt))

(defun main ()
  (define (ch (go:make-chan)))
  (go:chan.send "Hello channels")
(print (ch (go:chan.rcve)
```

Channels are Go-powered lightweight multiple Goroutines that interact and cummicate with each other. Channels is the real way to express concurrency in Picnic. You can compare the system with Crystals channeling system. Channels and Goroutines are one of the most important component in picnic. They control all executed code and take part in case of error exceptions.

More docs are in progress...

### Contributors

-Timo Sarkar lead developer



