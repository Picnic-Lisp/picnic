/*
* This file is Part of the Picnic-Lisp research project
*
* Copyright (c) 2020 Timo Sarkar
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to
* deal in the Software without restriction, including without limitation the
* rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
* sell copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in
* all copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
* FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
* IN THE SOFTWARE.
*/

package main

import (
	"encoding/json"
	"bufio"
	"io"
	"io/ioutil"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"

	"github.com/mattn/go-isatty"
	"github.com/timo-cmd/picnic"
)



const (

	HEADER    = "\033[95m"

	INFO      = "\033[94m"

	SUCCESS   = "\033[92m"

	WARNING   = "\033[93m"

	ERROR     = "\033[91m"

	ENDC      = "\033[0m"

	BOLD      = "\033[1m"

	UNDERLINE = "\033[4m"



	MODULE_DIR = ""



	HELP_MESSAGE = "picnic init, initialization \n" +

		"picnic -g, use system GOPATH \n" +

		"picnic -v, show version \n" +

		"picnic -i, install dependencies \n" +

		"picnic -i <name>, install package \n" +

		"picnic -i -s <name>, install package and save \n" +

		"picnic -e <command>, execute installed go binary \n" +

		"picnic -c <command>, run system command 'ex. picnic -c go build'"



	picnic_VERSION = "0.0.2"

)



type picnic struct{}



type packageJson struct {

	Name         string

	Description  string

	Version      string

	Dependencies []string

	Author       string

}



type picnicInterface interface {

	init()



	getGoRoot() string

	getGoPath() string

	getVersion() string



	setGoPathTmp(path string) error



	execCmd(cmd string, wg *sync.WaitGroup) []byte

	execCmdAsync(cmd string)

	runBinary(name string)

	runBinaryAsync(name string)

	isInStrings(str string, list []string) bool

	removeFromSilce(str string, list []string) []string

	checkErr(err error)



	install(name string)

	installDependencies()

	saveDependency(name string)



	headerMessage(message string)

	successMessage(message string)

	infoMessage(message string)

	boldMessage(message string)

	underlineMessage(message string)

	warningMessage(message string)

	errorMessage(err error)

}



func (g picnic) headerMessage(message string) {

	fmt.Print(HEADER, message, ENDC)

}



func (g picnic) successMessage(message string) {

	fmt.Print(SUCCESS, message, ENDC)

}



func (g picnic) infoMessage(message string) {

	fmt.Print(INFO, message, ENDC)

}



func (g picnic) boldMessage(message string) {

	fmt.Print(BOLD, message, ENDC)

}



func (g picnic) underlineMessage(message string) {

	fmt.Print(UNDERLINE, message, ENDC)

}



func (g picnic) warningMessage(message string) {

	fmt.Print(WARNING, "warning:", message, ENDC)

}



func (g picnic) errorMessage(err error) {

	fmt.Print(ERROR, "error: ", err, ENDC, "\n")

}



func (g picnic) checkErr(err error) {

	if err != nil {

		g.errorMessage(err)

		os.Exit(0)

	}

}



func (g picnic) getGoPath() string {

	path := os.Getenv("GOPATH")



	if len(path) > 0 && path[len(path)-1] == '/' {

		path = path[:len(path)-1]

	}



	return path

}



func (g picnic) getGoRoot() string {

	return os.Getenv("GOROOT")

}



func (g picnic) getVersion() string {

	wg := new(sync.WaitGroup)

	wg.Add(1)

	goVersion := g.execCmd("go version", wg)

	wg.Wait()

	return "picnic version " + picnic_VERSION + "\n" + string(goVersion)

}



func (g picnic) setGoPathTmp(path string) error {

	err := os.Setenv("GOPATH", path)

	g.checkErr(err)

	return err

}



func (g picnic) execCmd(cmd string, wg *sync.WaitGroup) []byte {

	parts := strings.Fields(cmd)

	head := parts[0]

	parts = parts[1:len(parts)]



	out, err := exec.Command(head, parts...).Output()

	g.checkErr(err)



	wg.Done()

	return out

}



func (g picnic) execCmdAsync(cmd string) {

	parts := strings.Fields(cmd)

	head := parts[0]

	parts = parts[1:len(parts)]



	command := exec.Command(head, parts...)



	stdout, err := command.StdoutPipe()

	g.checkErr(err)

	stderr, err := command.StderrPipe()

	g.checkErr(err)



	err = command.Start()

	g.checkErr(err)



	defer command.Wait()



	go io.Copy(os.Stdout, stdout)

	go io.Copy(os.Stderr, stderr)

}



func (g picnic) runBinary(name string) {

	wg := new(sync.WaitGroup)

	wg.Add(1)

	out := g.execCmd(g.getGoPath()+"/bin/"+name, wg)

	g.successMessage(string(out))

	wg.Wait()

}



func (g picnic) runBinaryAsync(name string) {

	g.execCmdAsync(g.getGoPath() + "/bin/" + name)

}



func (g picnic) isInStrings(str string, list []string) bool {

	for _, part := range list {

		if part == str {

			return true

		}

	}

	return false

}



func (g picnic) removeFromSilce(str string, list []string) []string {

	s := list

	for i, part := range list {

		if part == str {

			s = append(s[:i], s[i+1:]...)

		}

	}

	return s

}



func (g picnic) install(name string) {

	g.execCmdAsync("go get " + name)

}



func (g picnic) installDependencies() {

	var packageFile packageJson



	file, err := ioutil.ReadFile("package.json")

	g.checkErr(err)



	err = json.Unmarshal(file, &packageFile)

	g.checkErr(err)



	for _, el := range packageFile.Dependencies {

		g.install(el)

	}

}



func (g picnic) saveDependency(name string) {

	var packageFile packageJson



	file, err := ioutil.ReadFile("package.json")

	g.checkErr(err)



	err = json.Unmarshal(file, &packageFile)

	g.checkErr(err)



	if !g.isInStrings(name, packageFile.Dependencies) {

		packageFile.Dependencies = append(packageFile.Dependencies, name)

	}



	packageByte, err := json.MarshalIndent(packageFile, "", "\t")

	g.checkErr(err)



	err = ioutil.WriteFile("package.json", packageByte, 0644)

	g.checkErr(err)

}



func (g picnic) init() {

	g.headerMessage("Press ^C at any time to quit.\n")



	var packageFile packageJson



	pwd, err := os.Getwd()

	g.checkErr(err)

	pwdSlice := strings.Split(pwd, "/")

	packageFile.Name = pwdSlice[len(pwdSlice)-1]

	g.infoMessage("Name: (" + packageFile.Name + ") ")

	fmt.Scanln(&packageFile.Name)



	file, err := ioutil.ReadFile("README.md")

	// g.checkErr(err)

	packageFile.Description = string(file)



	packageFile.Version = "0.0.1"

	g.infoMessage("Version: (" + packageFile.Version + ") ")

	fmt.Scanln(&packageFile.Version)



	usr, err := user.Current()

	g.checkErr(err)

	packageFile.Author = usr.Username

	g.infoMessage("Author: (" + packageFile.Author + ") ")

	fmt.Scanln(&packageFile.Author)



L1:

	isOk := "yes"

	g.boldMessage("Is this ok?: (" + isOk + ") ")

	fmt.Scanln(&isOk)



	if isOk == "no" {

		g.headerMessage("\nAborted")

		os.Exit(0)

	} else if isOk != "no" && isOk != "yes" {

		goto L1

	}



	packageByte, err := json.MarshalIndent(packageFile, "", "\t")

	g.checkErr(err)



	err = ioutil.WriteFile("package.json", packageByte, 0644)

	g.checkErr(err)

	g.successMessage("Done.\n")

}

func repl() {
	env := picnic.NewEnv(nil)
	err := picnic.LoadLib(env)
	if err != nil {

	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		parser := picnic.NewParser(strings.NewReader(scanner.Text()))
		node, err := parser.Parse()
		if err != nil {
			log.Fatal(err)
		}

		ret, err := env.Eval(node)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ret)
	}
}

func main() {
	var picnic picnicInterface = new(picnic)



	args := os.Args[1:]



	if !picnic.isInStrings("-g", args) {

		pwd, err := os.Getwd()

		picnic.checkErr(err)

		picnic.setGoPathTmp(pwd + MODULE_DIR)

	}



	args = picnic.removeFromSilce("-g", args)



	if picnic.isInStrings("-c", args) && len(args) != 1 {

		picnic.execCmdAsync(strings.Join(args[1:], " "))

		os.Exit(0)

	}



	if picnic.isInStrings("init", args) {

		picnic.init()

		os.Exit(0)

	}



	if picnic.isInStrings("-i", args) && !picnic.isInStrings("-s", args) && len(args) == 1 {

		picnic.infoMessage("Installing dependencies...\n")

		picnic.installDependencies()

		picnic.successMessage("Done.\n")

		os.Exit(0)

	}



	if picnic.isInStrings("-i", args) && !picnic.isInStrings("-s", args) && len(args) == 2 {

		args = picnic.removeFromSilce("-i", args)

		picnic.infoMessage("Installing package " + args[0] + "...\n")

		picnic.install(args[0])

		picnic.successMessage("Done.\n")

		os.Exit(0)

	}



	if picnic.isInStrings("-i", args) && picnic.isInStrings("-s", args) && len(args) == 3 {

		args = picnic.removeFromSilce("-i", args)

		args = picnic.removeFromSilce("-s", args)



		picnic.infoMessage("Installing " + args[0] + "...\n")

		picnic.install(args[0])



		picnic.saveDependency(args[0])

		picnic.infoMessage("Package " + args[0] + " added in package.json.\n")



		picnic.successMessage("Done.\n")

		os.Exit(0)

	}



	args = picnic.removeFromSilce("-i", args)



	if picnic.isInStrings("-e", args) {

		args = picnic.removeFromSilce("-e", args)

		picnic.runBinaryAsync(strings.Join(args, " "))

		os.Exit(0)

	}



	if picnic.isInStrings("-v", args) {

		args = picnic.removeFromSilce("-v", args)

		picnic.infoMessage(picnic.getVersion())

		os.Exit(0)

	}



	picnic.headerMessage(HELP_MESSAGE + "\n")

	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(2)
	}

	var f *os.File
	var err error

	if flag.NArg() == 0 {
		if isatty.IsTerminal(os.Stdin.Fd()) {
			repl()
			return
		}
		f = os.Stdin
	}

	if flag.NArg() == 1 {
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	parser := picnic.NewParser(f)
	node, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	env := picnic.NewEnv(nil)
	err = picnic.LoadLib(env)
	if err != nil {
		log.Fatal(err)
	}
	_, err = env.Eval(node)
	if err != nil {
		log.Fatal(err)
	}
}
