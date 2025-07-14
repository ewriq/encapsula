package main

import (
	"encapsula/cli"
	"encapsula/handler"
)

func main() {
	cli.Execute()
	handler.BuildWSLRDP()
}