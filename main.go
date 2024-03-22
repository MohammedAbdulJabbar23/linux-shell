package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
  "log"
  "os/user"
)

// ANSI color codes
const (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Reset   = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	welcomeMessage();
  neofetch();
  for {
    path, err := os.Getwd();
    if err != nil {
      log.Println(err)
    }
    user, err := user.Current();
    if err != nil {
      log.Fatalf(err.Error());
    }
    fmt.Print(Green +user.Username+":"+Blue+ path + "$ "+ Reset)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, Red+err.Error()+Reset)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
  case "":
    //do nothing
    return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func neofetch() {
  cmd :=exec.Command("neofetch");
  cmd.Stderr = os.Stderr;
  cmd.Stdout = os.Stdout;
  cmd.Run();
}

func welcomeMessage() {
    fmt.Println(Cyan+"***************************")
    fmt.Println("*  Welcome to Go Shell!   *")
    fmt.Println("***************************"+Reset)
}

