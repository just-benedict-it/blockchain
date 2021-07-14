package main

import (
	"os"

	"github.com/just-benedict-it/blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
