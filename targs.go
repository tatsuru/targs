package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var tmux = flag.String("tmux-command", "tmux", "tmux command name")
	var windowTitle = flag.String("t", fmt.Sprintf("targs-%d", os.Getpid()), "window title")
	var remainOnExit = flag.Bool("r", false, "remain window on exit command")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Error: Commands required.\n")
		flag.Usage()
		os.Exit(1)
	}

	command := strings.Join(flag.Args(), " ")
	bytes, _ := ioutil.ReadAll(os.Stdin)
	re := regexp.MustCompile("[^\\s]+")
	args := re.FindAllString(string(bytes), -1)

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Arguments required.\n")
		flag.Usage()
		os.Exit(1)
	}
	
	exec.Command(*tmux, "new-window", "-n", *windowTitle, fmt.Sprintf("%s %s", command, args[0])).Run()
	if *remainOnExit {
		exec.Command(*tmux, "set-option", "-t", *windowTitle, "remain-on-exit", "on").Run()
	}

	for i := 1; i < len(args); i++ {
		exec.Command(*tmux, "split-window", "-t", *windowTitle, fmt.Sprintf("%s %s", command, args[i])).Run()
		exec.Command(*tmux, "select-layout", "-t", *windowTitle, "tiled").Run()
	}

	exec.Command(*tmux, "set-window", "-t", *windowTitle, "synchronize-panes").Run()
}
