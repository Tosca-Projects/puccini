package main

import (
	"github.com/tebeka/atexit"
	"github.com/tliron/puccini/puccini-clout/commands"
)

func main() {
	commands.Execute()
	atexit.Exit(0)
}
